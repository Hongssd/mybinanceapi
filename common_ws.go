package mybinanceapi

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/gorilla/websocket"
)

const (
	BINANCE_API_SPOT_WS_STREAM        = "stream.binance.com:9443"
	BINANCE_API_SWAP_WS_STREAM        = "dstream.binance.com"
	BINANCE_API_FUTURE_WS_STREAM      = "fstream.binance.com"
	BINANCE_API_SWAP_WS_STREAM_GZIP   = "sdstream.binance.com"
	BINANCE_API_FUTURE_WS_STREAM_GZIP = "sfstream.binance.com"

	BINANCE_API_SPOT_WS_API   = "ws-api.binance.com:9443"
	BINANCE_API_FUTURE_WS_API = "ws-fapi.binance.com"

	TEST_BINANCE_API_SPOT_WS_STREAM   = "testnet.binance.vision"
	TEST_BINANCE_API_FUTURE_WS_STREAM = "fstream.binancefuture.com"
	TEST_BINANCE_API_SWAP_WS_STREAM   = "dstream.binancefuture.com"

	TEST_BINANCE_API_SPOT_WS_API   = "testnet.binance.vision"
	TEST_BINANCE_API_FUTURE_WS_API = "testnet.binancefuture.com"
)

const (
	SUBSCRIBE          = "SUBSCRIBE"
	UNSUBSCRIBE        = "UNSUBSCRIBE"
	LIST_SUBSCRIPTIONS = "LIST_SUBSCRIPTIONS"
)

var (
	// WebsocketTimeout is an interval for sending ping/pong messages if WebsocketKeepalive is enabled
	WebsocketTimeout = time.Second * 10
	// WebsocketKeepalive enables sending ping/pong messages to check the connection stability
	WebsocketKeepalive = true
	//SUBSCRIBE_INTERVAL_TIME = 500 * time.Millisecond //订阅间隔

	ListenKeyRefreshInterval = 30 * time.Minute
)

type WsStreamPath int

const (
	WS_STREAM_PATH WsStreamPath = iota
	WS_ACCOUNT_PATH
	WS_SPOT_API_PATH
	WS_FUTURE_API_PATH
)

// 数据流订阅基础客户端
type WsStreamClient struct {
	apiType ApiType
	isGzip  bool
	conn    *websocket.Conn

	////行情订阅相关
	//commonSubMap   map[int64]*Subscription[SubscribeResult] //订阅的返回结果
	//klineSubMap    map[string]*Subscription[WsKline]
	//depthSubMap    map[string]*Subscription[WsDepth]
	//aggTradeSubMap map[string]*Subscription[WsAggTrade]
	//
	////账户推送相关
	//wsSpotPayloadMap   map[int64]*WsSpotPayload
	//wsFuturePayloadMap map[int64]*WsFuturePayload
	//wsSwapPayloadMap   map[int64]*WsSwapPayload

	//行情订阅相关
	commonSubMap   MySyncMap[int64, *Subscription[SubscribeResult]] //订阅的返回结果
	klineSubMap    MySyncMap[string, *Subscription[WsKline]]
	depthSubMap    MySyncMap[string, *Subscription[WsDepth]]
	aggTradeSubMap MySyncMap[string, *Subscription[WsAggTrade]]

	//账户推送相关
	wsSpotPayloadMap   MySyncMap[int64, *WsSpotPayload]
	wsFuturePayloadMap MySyncMap[int64, *WsFuturePayload]
	wsSwapPayloadMap   MySyncMap[int64, *WsSwapPayload]

	//wsApi交易相关
	waitWsApiResultMap MySyncMap[string, WsApiResultChan]

	resultChan               chan []byte
	errChan                  chan error
	isClose                  bool
	reSubscribeMu            *sync.Mutex
	apiKey                   string
	apiSecret                string
	wsStreamPath             WsStreamPath
	isListenWs               bool
	listenKey                string
	listenKeyRefreshStopChan *chan struct{}

	AutoReConnectTimes int //自动重连次数
}

// 订阅请求结构体
type SubscribeReq struct {
	Method string   `json:"method"`
	Params []string `json:"params,omitempty"`
	Id     int64    `json:"id"`
}

// 订阅结果结构体
type SubscribeResult struct {
	Result []string `json:"result"`
	Id     int64    `json:"id"`
	Error  string   `json:"error"`
}

// 数据流订阅参数
type SubscribeParams []string

// 数据流订阅返回标准结构体
type Subscription[T any] struct {
	ws         *WsStreamClient //订阅的连接
	ID         int64           //此次订阅ID
	Method     string          //订阅方法
	Params     SubscribeParams //订阅参数
	resultChan chan T          //接收订阅结果的通道
	errChan    chan error      //接收订阅错误的通道
	closeChan  chan struct{}   //接收订阅关闭的通道
}

func (ws *WsStreamClient) GetConn() *websocket.Conn {
	return ws.conn
}

// 获取订阅结果
func (sub *Subscription[T]) ResultChan() chan T {
	return sub.resultChan
}

// 获取错误订阅
func (sub *Subscription[T]) ErrChan() chan error {
	return sub.errChan
}

// 获取关闭订阅信号
func (sub *Subscription[T]) CloseChan() chan struct{} {
	return sub.closeChan
}

func sendMsg[T any](ws *WsStreamClient, id int64, method string, params []string) (*Subscription[T], error) {
	if ws == nil || ws.conn == nil || ws.isClose {
		return nil, fmt.Errorf("websocket is close")
	}
	if id == 0 {
		node, err := snowflake.NewNode(1)
		if err != nil {
			return nil, err
		}
		id = node.Generate().Int64()
	}
	log.Debugf("send msg: {%d:%s:%v}", id, method, params)
	subscribeReq := SubscribeReq{
		Method: method,
		Params: params,
		Id:     id,
	}
	data, err := json.Marshal(subscribeReq)
	if err != nil {
		return nil, err
	}
	err = ws.conn.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		return nil, err
	}
	result := &Subscription[T]{
		ID:         subscribeReq.Id,
		Method:     method,
		Params:     params,
		resultChan: make(chan T),
		errChan:    make(chan error),
		closeChan:  make(chan struct{}),
		ws:         ws,
	}

	return result, nil
}

func (ws *WsStreamClient) Close() error {
	ws.isClose = true

	err := ws.conn.Close()
	if err != nil {
		return err
	}
	//手动关闭成功，给所有订阅发送关闭信号
	go ws.sendWsCloseToAllSub()

	//等待5秒，如果还有订阅未关闭，则强制关闭
	time.Sleep(5 * time.Second)
	//初始化连接状态
	ws.conn = nil
	//close(ws.resultChan)
	//close(ws.errChan)
	ws.resultChan = nil
	ws.errChan = nil
	ws.initStructs()
	return nil
}

func (ws *WsStreamClient) initStructs() {
	//ws.commonSubMap = make(map[int64]*Subscription[SubscribeResult])
	//ws.klineSubMap = make(map[string]*Subscription[WsKline])
	//ws.depthSubMap = make(map[string]*Subscription[WsDepth])
	//ws.aggTradeSubMap = make(map[string]*Subscription[WsAggTrade])
	//
	//ws.wsSpotPayloadMap = make(map[int64]*WsSpotPayload)
	//ws.wsFuturePayloadMap = make(map[int64]*WsFuturePayload)
	//ws.wsSwapPayloadMap = make(map[int64]*WsSwapPayload)
	ws.commonSubMap = NewMySyncMap[int64, *Subscription[SubscribeResult]]()
	ws.klineSubMap = NewMySyncMap[string, *Subscription[WsKline]]()
	ws.depthSubMap = NewMySyncMap[string, *Subscription[WsDepth]]()
	ws.aggTradeSubMap = NewMySyncMap[string, *Subscription[WsAggTrade]]()

	ws.wsSpotPayloadMap = NewMySyncMap[int64, *WsSpotPayload]()
	ws.wsFuturePayloadMap = NewMySyncMap[int64, *WsFuturePayload]()
	ws.wsSwapPayloadMap = NewMySyncMap[int64, *WsSwapPayload]()

	ws.waitWsApiResultMap = NewMySyncMap[string, WsApiResultChan]()
}

func (ws *WsStreamClient) OpenConn() error {
	if ws.resultChan == nil {
		ws.resultChan = make(chan []byte)
	}
	if ws.errChan == nil {
		ws.errChan = make(chan error)
	}
	apiUrl := handlerWsStreamRequestApi(ws.wsStreamPath, ws.listenKey, ws.apiType, ws.isGzip)
	if ws.conn == nil {
		conn, err := wsStreamServe(apiUrl, ws.isGzip, ws.resultChan, ws.errChan)
		if err != nil {
			return err
		}
		ws.conn = conn
		ws.isClose = false
		log.Info("OpenConn success to ", apiUrl)
		ws.handleResult(ws.resultChan, ws.errChan)

	} else {
		conn, err := wsStreamServe(apiUrl, ws.isGzip, ws.resultChan, ws.errChan)
		if err != nil {
			return err
		}
		ws.conn = conn
		log.Info("Auto ReOpenConn success to ", apiUrl)
	}
	return nil
}

func (ws *WsStreamClient) sendSubscribeResultToChan(result SubscribeResult) {
	if result.Error != "" {
		if sub, ok := ws.commonSubMap.Load(result.Id); ok {
			sub.errChan <- fmt.Errorf("errHandler: %s===%v", result.Error, sub.Params)
		}
	} else {
		if sub, ok := ws.commonSubMap.Load(result.Id); ok {
			sub.resultChan <- result
		}
	}
}

func (ws *WsStreamClient) sendUnSubscribeSuccessToCloseChan(params []string) {
	isCloseMap := map[int64]bool{}
	for _, param := range params {
		if sub, ok := ws.klineSubMap.Load(param); ok {
			ws.klineSubMap.Delete(param)
			if _, ok2 := isCloseMap[sub.ID]; ok2 {
				continue
			}
			sub.closeChan <- struct{}{}
			isCloseMap[sub.ID] = true
		} else if sub, ok := ws.depthSubMap.Load(param); ok {
			ws.depthSubMap.Delete(param)
			if _, ok2 := isCloseMap[sub.ID]; ok2 {
				continue
			}
			sub.closeChan <- struct{}{}
			isCloseMap[sub.ID] = true
		} else if sub, ok := ws.aggTradeSubMap.Load(param); ok {
			ws.aggTradeSubMap.Delete(param)
			if _, ok2 := isCloseMap[sub.ID]; ok2 {
				continue
			}
			sub.closeChan <- struct{}{}
			isCloseMap[sub.ID] = true
		}
	}
}

func (ws *WsStreamClient) sendWsCloseToAllSub() {
	params := []string{}
	ws.commonSubMap.Range(func(_ int64, sub *Subscription[SubscribeResult]) bool {
		params = append(params, sub.Params...)
		return true
	})
	ws.sendUnSubscribeSuccessToCloseChan(params)
	ws.waitWsApiResultMap.Range(func(_ string, ch WsApiResultChan) bool {
		ch.CloseChan() <- struct{}{}
		return true
	})
}

func (ws *WsStreamClient) reSubscribeForReconnect() error {
	if ws.reSubscribeMu.TryLock() {
		defer ws.reSubscribeMu.Unlock()
	} else {
		return nil
	}

	var wErr error
	ws.commonSubMap.Range(func(_ int64, sub *Subscription[SubscribeResult]) bool {
		log.Infof("ReSubscribe:{%d,%s,%v}", sub.ID, sub.Method, sub.Params)
		doSub, err := sendMsg[SubscribeResult](ws, sub.ID, sub.Method, sub.Params)
		if err != nil {
			log.Error(err)
			wErr = err
			return false
		}

		ws.commonSubMap.Store(doSub.ID, doSub)

		select {
		case err := <-doSub.ErrChan():
			log.Error("Subscribe Error: ", err)
			wErr = err
			return false
		case subResult := <-doSub.ResultChan():
			if subResult.Error != "" {
				log.Error(subResult.Error)
				wErr = fmt.Errorf(subResult.Error)
			}
			log.Debugf("Subscribe Success: params:%v result:%v", doSub.Params, subResult)
		}
		log.Infof("reSubscribe Success: {%d,%s,%v}", sub.ID, sub.Method, sub.Params)
		time.Sleep(1000 * time.Millisecond)
		return true
	})

	return wErr
}

func (ws *WsStreamClient) handleResult(resultChan chan []byte, errChan chan error) {
	go func() {
		for {
			select {
			case err, ok := <-errChan:
				if !ok {
					return
				}
				log.Error(err)

				//错误处理 重连等
				//ws标记为非关闭 且返回错误包含EOF、close、reset时自动重连
				if !ws.isClose && (strings.Contains(err.Error(), "EOF") ||
					strings.Contains(err.Error(), "close") ||
					strings.Contains(err.Error(), "reset")) {
					log.Error("意外断连,5秒后自动重连: ", err.Error())
					time.Sleep(5 * time.Second)
					err := ws.OpenConn()
					for err != nil {
						time.Sleep(5000 * time.Millisecond)
						err = ws.OpenConn()
					}
					ws.AutoReConnectTimes += 1
					go func() {
						err = ws.reSubscribeForReconnect()
						if err != nil {
							log.Error(err)
						}
					}()

				} else {
					continue
				}
			case data, ok := <-resultChan:
				if !ok {
					return
				}

				//wsApi返回结果
				if strings.Contains(string(data), "id") &&
					strings.Contains(string(data), "rateLimits") &&
					(strings.Contains(string(data), "error") || strings.Contains(string(data), "result")) {
					wsApiResult := WsApiResult[struct{}]{}
					err := json.Unmarshal(data, &wsApiResult)
					if err != nil {
						log.Error(err)
						continue
					}
					if ch, ok := ws.waitWsApiResultMap.Load(wsApiResult.Id); ok {
						err = wsApiResult.Error.handlerError()
						if err != nil {
							ch.ErrChan() <- err
							continue
						} else {
							ch.ResultChan() <- data
						}
						ws.waitWsApiResultMap.Delete(wsApiResult.Id)
					}
					continue
				}

				// log.Debugf("[%s] receive result: %s", ws.apiType.String(), string(data))
				//处理订阅或查询订阅列表请求返回结果
				if strings.Contains(string(data), "error") || strings.Contains(string(data), "result") {
					result := SubscribeResult{}
					err := json.Unmarshal(data, &result)
					if err != nil {
						log.Error(err)
						continue
					}
					ws.sendSubscribeResultToChan(result)
					continue
				}

				//处理正常数据的返回结果
				//K线订阅
				if strings.Contains(string(data), "@kline") {
					var k *WsKline
					var err error
					//K线处理
					// if !ws.isGzip {
					k, err = HandleWsCombinedKline(ws.apiType, data)
					// } else {
					// 	k, err = HandleWsCombinedKlineGzip(ws.apiType, data)
					// }
					param := getKlineSubscribeParam(k.Symbol, k.Interval)
					if sub, ok := ws.klineSubMap.Load(param); ok {
						if err != nil {
							sub.errChan <- err
							continue
						}
						sub.resultChan <- *k
					}
					continue
				}

				//深度订阅
				if strings.Contains(string(data), "@depth") {
					var d *WsDepth
					var err error
					//深度处理
					// if !ws.isGzip {
					d, err = HandleWsCombinedDepth(ws.apiType, data)
					// } else {
					// 	d, err = HandleWsCombinedDepthGzip(ws.apiType, data, false, true)
					// }
					param := getLevelDepthSubscribeParam(d.Symbol, d.Level, d.USpeed)
					// log.Warn(param)
					if sub, ok := ws.depthSubMap.Load(param); ok {
						if err != nil {
							sub.errChan <- err
							continue
						}
						sub.resultChan <- *d
					}
					continue
				}

				//归集交易流订阅
				if strings.Contains(string(data), "@aggTrade") {
					var a *WsAggTrade
					var err error
					//交易流处理
					// if !ws.isGzip {
					a, err = HandleWsCombinedAggTrade(ws.apiType, data)
					// } else {
					// 	a, err = HandleWsCombinedAggTradeGzip(ws.apiType, data)
					// }
					param := getAggTradeParam(a.Symbol)
					if sub, ok := ws.aggTradeSubMap.Load(param); ok {
						if err != nil {
							sub.errChan <- err
							continue
						}
						sub.resultChan <- *a
					}
					continue
				}

				//现货账户更新推送
				if strings.Contains(string(data), "outboundAccountPosition") {
					res, err := HandleWsPayloadResult[WsSpotPayloadOutboundAccountPosition](data)
					if err != nil {
						log.Error(err)
						continue
					}
					ws.wsSpotPayloadMap.Range(func(_ int64, payload *WsSpotPayload) bool {
						if payload.OutboundAccountPositionPayload != nil {
							payload.OutboundAccountPositionPayload.resultChan <- *res
						}
						return true
					})
				}

				//现货余额更新推送
				if strings.Contains(string(data), "balanceUpdate") {
					res, err := HandleWsPayloadResult[WsSpotPayloadBalanceUpdate](data)
					if err != nil {
						log.Error(err)
						continue
					}
					ws.wsSpotPayloadMap.Range(func(_ int64, payload *WsSpotPayload) bool {
						if payload.BalanceUpdatePayload != nil {
							payload.BalanceUpdatePayload.resultChan <- *res
						}
						return true
					})
				}

				//现货订单推送
				if strings.Contains(string(data), "executionReport") {
					res, err := HandleWsPayloadResult[WsSpotPayloadExecutionReport](data)
					if err != nil {
						log.Error(err)
						continue
					}
					ws.wsSpotPayloadMap.Range(func(_ int64, payload *WsSpotPayload) bool {
						if payload.ExecutionReportPayload != nil {
							payload.ExecutionReportPayload.resultChan <- *res
						}
						return true
					})
				}

				//U本位合约及币本位合约 余额/仓位 更新推送
				if strings.Contains(string(data), "ACCOUNT_UPDATE") {
					switch ws.apiType {
					case FUTURE:
						res, err := HandleWsPayloadResult[WsFuturePayloadAccountUpdate](data)
						if err != nil {
							log.Error(err)
							continue
						}
						ws.wsFuturePayloadMap.Range(func(_ int64, payload *WsFuturePayload) bool {
							if payload.AccountUpdatePayload != nil {
								payload.AccountUpdatePayload.resultChan <- *res
							}
							return true
						})

					case SWAP:
						res, err := HandleWsPayloadResult[WsSwapPayloadAccountUpdate](data)
						if err != nil {
							log.Error(err)
							continue
						}
						ws.wsSwapPayloadMap.Range(func(_ int64, payload *WsSwapPayload) bool {
							if payload.AccountUpdatePayload != nil {
								payload.AccountUpdatePayload.resultChan <- *res
							}
							return true
						})
					default:
					}

				}

				//U本位合约及币本位合约 订单/成交 更新推送
				if strings.Contains(string(data), "ORDER_TRADE_UPDATE") {
					switch ws.apiType {
					case FUTURE:
						res, err := HandleWsPayloadResult[WsFuturePayloadOrderTradeUpdate](data)
						if err != nil {
							log.Error(err)
							continue
						}
						ws.wsFuturePayloadMap.Range(func(_ int64, payload *WsFuturePayload) bool {
							if payload.OrderTradeUpdatePayload != nil {
								payload.OrderTradeUpdatePayload.resultChan <- *res
							}
							return true
						})
					case SWAP:
						res, err := HandleWsPayloadResult[WsSwapPayloadOrderTradeUpdate](data)
						if err != nil {
							log.Error(err)
							continue
						}
						ws.wsSwapPayloadMap.Range(func(_ int64, payload *WsSwapPayload) bool {
							if payload.OrderTradeUpdatePayload != nil {
								payload.OrderTradeUpdatePayload.resultChan <- *res
							}
							return true
						})

					default:
					}
				}

			}
		}
	}()
}

// 获取当前所有订阅
func (ws *WsStreamClient) CurrentSubscribeList() ([]string, error) {
	listSub, err := sendMsg[SubscribeResult](ws, 0, LIST_SUBSCRIPTIONS, []string{})
	if err != nil {
		return nil, err
	}
	ws.commonSubMap.Store(listSub.ID, listSub)
	select {
	case err := <-listSub.ErrChan():
		log.Error("CurrentSubscribeList Error: ", err)
		return nil, err
	case subResult := <-listSub.ResultChan():
		if subResult.Error != "" {
			log.Error(subResult.Error)
			return nil, fmt.Errorf(subResult.Error)
		}
		log.Debug("CurrentSubscribeList Success: ", subResult)
		ws.commonSubMap.Delete(listSub.ID)
		return subResult.Result, nil
	}
}

// 取消订阅
func (sub *Subscription[T]) Unsubscribe() error {
	unSub, err := sendMsg[SubscribeResult](sub.ws, 0, UNSUBSCRIBE, sub.Params)
	if err != nil {
		return err
	}
	sub.ws.commonSubMap.Store(unSub.ID, unSub)

	select {
	case err := <-unSub.ErrChan():
		return fmt.Errorf("Unsubscribe Error: %v", err)
	case subResult := <-unSub.ResultChan():
		if subResult.Error != "" {
			log.Error(subResult.Error)
			return fmt.Errorf(subResult.Error)
		}
		log.Infof("Unsubscribe Success Params:%v Result:%v", unSub.Params, subResult)

		//取消订阅成功，给所有订阅消息的通道发送关闭信号
		sub.ws.sendUnSubscribeSuccessToCloseChan(unSub.Params)
		//删除当前订阅列表中已存在的记录
		sub.ws.commonSubMap.Delete(unSub.ID)
		sub.ws.commonSubMap.Delete(sub.ID)
		return nil
	}

}

// 标准订阅方法
func wsStreamServe(api string, isGzip bool, resultChan chan []byte, errChan chan error) (*websocket.Conn, error) {
	c, _, err := websocket.DefaultDialer.Dial(api, nil)
	if err != nil {
		return nil, err
	}
	c.SetReadLimit(655350)
	go func() {
		if WebsocketKeepalive {
			keepAlive(c, WebsocketTimeout)
		}
		if isGzip {
			for {
				_, message, err := c.ReadMessage()
				if err != nil {
					errChan <- err
					return
				}
				log.Warn(string(message))
				reader, err := gzip.NewReader(bytes.NewReader(message))
				if err != nil {
					errChan <- err
					return
				}
				unzipMessage, err := io.ReadAll(reader)
				if err != nil {
					errChan <- err
					return
				}
				err = reader.Close()
				if err != nil {
					errChan <- err
					return
				}
				resultChan <- unzipMessage
			}
		} else {
			for {
				_, message, err := c.ReadMessage()
				if err != nil {
					errChan <- err
					return
				}
				resultChan <- message
			}
		}
	}()
	return c, err
}

// 获取数据流请求URL
func handlerWsStreamRequestApi(wsStreamPath WsStreamPath, listenKey string, apiType ApiType, isGzip bool) string {
	host := ""
	switch wsStreamPath {
	case WS_SPOT_API_PATH, WS_FUTURE_API_PATH:
		host = getWsApiWsApi(apiType)
	case WS_STREAM_PATH, WS_ACCOUNT_PATH:
		host = getWsStreamWsApi(apiType, isGzip)
	}

	u := url.URL{
		Scheme:   "wss",
		Host:     host,
		Path:     getWsPath(wsStreamPath, listenKey),
		RawQuery: "",
	}
	return u.String()
}

// 获取数据流请求Path
func getWsStreamWsApi(apiType ApiType, isGzip bool) string {
	switch apiType {
	case SPOT:
		switch NowNetType {
		case MAIN_NET:
			return BINANCE_API_SPOT_WS_STREAM
		case TEST_NET:
			return TEST_BINANCE_API_SPOT_WS_STREAM
		}
	case SWAP:
		switch NowNetType {
		case MAIN_NET:
			if isGzip {
				return BINANCE_API_SWAP_WS_STREAM_GZIP
			} else {
				return BINANCE_API_SWAP_WS_STREAM
			}
		case TEST_NET:
			return TEST_BINANCE_API_SWAP_WS_STREAM
		}
	case FUTURE:
		switch NowNetType {
		case MAIN_NET:
			if isGzip {
				return BINANCE_API_FUTURE_WS_STREAM_GZIP
			} else {
				return BINANCE_API_FUTURE_WS_STREAM
			}
		case TEST_NET:
			return TEST_BINANCE_API_FUTURE_WS_STREAM
		}
	}
	log.Error("AccountType Error is ", apiType)
	return ""
}

// 获取数据流请求Path
func getWsApiWsApi(apiType ApiType) string {
	switch apiType {
	case SPOT:
		switch NowNetType {
		case MAIN_NET:
			return BINANCE_API_SPOT_WS_API
		case TEST_NET:
			return TEST_BINANCE_API_SPOT_WS_API
		}
	case SWAP:
		return ""
	case FUTURE:
		switch NowNetType {
		case MAIN_NET:
			return BINANCE_API_FUTURE_WS_API
		case TEST_NET:
			return TEST_BINANCE_API_FUTURE_WS_API
		}
	}
	log.Error("AccountType Error is ", apiType)
	return ""
}

func getWsPath(wsStreamPath WsStreamPath, listenKey string) string {
	switch wsStreamPath {
	case WS_STREAM_PATH:
		return "/stream"
	case WS_ACCOUNT_PATH:
		return fmt.Sprintf("/ws/%s", listenKey)
	case WS_SPOT_API_PATH:
		return "/ws-api/v3"
	case WS_FUTURE_API_PATH:
		return "/ws-fapi/v1"
	}
	log.Error("WsStreamPath Error is ", wsStreamPath)
	return ""
}

// 发送ping/pong消息以检查连接稳定性
func keepAlive(c *websocket.Conn, timeout time.Duration) {
	ticker := time.NewTicker(timeout)

	lastResponse := time.Now()
	c.SetPongHandler(func(msg string) error {
		lastResponse = time.Now()
		return nil
	})

	go func() {
		defer ticker.Stop()
		for {
			deadline := time.Now().Add(10 * time.Second)
			err := c.WriteControl(websocket.PingMessage, []byte{}, deadline)
			if err != nil {
				return
			}
			<-ticker.C
			if time.Since(lastResponse) > timeout {
				err := c.Close()
				if err != nil {
					return
				}
				return
			}
		}
	}()
}
