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

	TEST_BINANCE_API_SPOT_WS_STREAM   = "testnet.binance.vision"
	TEST_BINANCE_API_FUTURE_WS_STREAM = "fstream.binancefuture.com"
	TEST_BINANCE_API_SWAP_WS_STREAM   = "dstream.binancefuture.com"
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
	WebsocketKeepalive      = true
	SUBSCRIBE_INTERVAL_TIME = 500 * time.Millisecond //订阅间隔

	ListenKeyRefreshInterval = 30 * time.Minute
)

type WsStreamPath int

const (
	WS_STREAM_PATH WsStreamPath = iota
	WS_ACCOUNT_PATH
)

// 数据流订阅基础客户端
type WsStreamClient struct {
	apiType ApiType
	isGzip  bool
	conn    *websocket.Conn

	//行情订阅相关
	commonSubMap   map[int64]*Subscription[SubscribeResult] //订阅的返回结果
	klineSubMap    map[string]*Subscription[WsKline]
	depthSubMap    map[string]*Subscription[WsDepth]
	aggTradeSubMap map[string]*Subscription[WsAggTrade]

	//账户推送相关
	wsSpotPayloadMap   map[int64]*WsSpotPayload
	wsFuturePayloadMap map[int64]*WsFuturePayload
	wsSwapPayloadMap   map[int64]*WsSwapPayload

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

type Payload[T any] struct {
	resultChan chan T        //接收订阅结果的通道
	errChan    chan error    //接收订阅错误的通道
	closeChan  chan struct{} //接收订阅关闭的通道
}

func (p *Payload[T]) ResultChan() chan T {
	return p.resultChan
}

func (p *Payload[T]) ErrChan() chan error {
	return p.errChan
}

func (p *Payload[T]) CloseChan() chan struct{} {
	return p.closeChan
}

type WsSpotPayload struct {
	Ws                             *SpotWsStreamClient
	Id                             int64
	OutboundAccountPositionPayload *Payload[WsSpotPayloadOutboundAccountPosition]
	BalanceUpdatePayload           *Payload[WsSpotPayloadBalanceUpdate]
	ExecutionReportPayload         *Payload[WsSpotPayloadExecutionReport]
}

type WsFuturePayload struct {
	Ws                      *FutureWsStreamClient
	Id                      int64
	AccountUpdatePayload    *Payload[WsFuturePayloadAccountUpdate]
	OrderTradeUpdatePayload *Payload[WsFuturePayloadOrderTradeUpdate]
}

type WsSwapPayload struct {
	Ws                      *SwapWsStreamClient
	Id                      int64
	AccountUpdatePayload    *Payload[WsSwapPayloadAccountUpdate]
	OrderTradeUpdatePayload *Payload[WsSwapPayloadOrderTradeUpdate]
}

func (payload *WsSpotPayload) Close() {
	if _, ok := payload.Ws.wsSpotPayloadMap[payload.Id]; ok {
		delete(payload.Ws.wsSpotPayloadMap, payload.Id)
		payload.OutboundAccountPositionPayload.closeChan <- struct{}{}
		payload.BalanceUpdatePayload.closeChan <- struct{}{}
		payload.ExecutionReportPayload.closeChan <- struct{}{}
	}
}

func (payload *WsFuturePayload) Close() {
	if _, ok := payload.Ws.wsFuturePayloadMap[payload.Id]; ok {
		delete(payload.Ws.wsFuturePayloadMap, payload.Id)
		payload.AccountUpdatePayload.closeChan <- struct{}{}
		payload.OrderTradeUpdatePayload.closeChan <- struct{}{}
	}
}

func (payload *WsSwapPayload) Close() {
	if _, ok := payload.Ws.wsSwapPayloadMap[payload.Id]; ok {
		delete(payload.Ws.wsSwapPayloadMap, payload.Id)
		payload.AccountUpdatePayload.closeChan <- struct{}{}
		payload.OrderTradeUpdatePayload.closeChan <- struct{}{}
	}
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
	ws.sendWsCloseToAllSub()

	//初始化连接状态
	ws.conn = nil
	close(ws.resultChan)
	close(ws.errChan)
	ws.resultChan = nil
	ws.errChan = nil
	ws.initStructs()
	return nil
}

func (ws *WsStreamClient) initStructs() {
	ws.commonSubMap = make(map[int64]*Subscription[SubscribeResult])
	ws.klineSubMap = make(map[string]*Subscription[WsKline])
	ws.depthSubMap = make(map[string]*Subscription[WsDepth])
	ws.aggTradeSubMap = make(map[string]*Subscription[WsAggTrade])

	ws.wsSpotPayloadMap = make(map[int64]*WsSpotPayload)
	ws.wsFuturePayloadMap = make(map[int64]*WsFuturePayload)
	ws.wsSwapPayloadMap = make(map[int64]*WsSwapPayload)
}

type SpotWsType int

const (
	SPOT_WS_TYPE SpotWsType = iota
	SPOT_MARGIN_WS_TYPE
	SPOT_ISOLATED_MARGIN_WS_TYPE
)

func (ws *SpotWsStreamClient) ConvertToAccountWs(apiKey string, apiSecret string, spotWsType SpotWsType, isolatedSymbol ...string) (*SpotWsStreamClient, error) {
	ws.wsStreamPath = WS_ACCOUNT_PATH
	ws.apiKey = apiKey
	ws.apiSecret = apiSecret
	ws.spotWsType = spotWsType
	ws.isListenWs = true
	b := MyBinance{}
	ws.client = b.NewSpotRestClient(apiKey, apiSecret)

	err := ws.listenKeyPost()
	if err != nil {
		return nil, err
	}
	//创建一个协程定时刷新listenKey，如果已存在旧的刷新协程则不再创建
	if ws.listenKeyRefreshStopChan == nil {
		stopChan := make(chan struct{})
		ws.listenKeyRefreshStopChan = &stopChan
		go func() {
			for {
				select {
				case <-time.After(ListenKeyRefreshInterval):
					err := ws.listenKeyPut()
					for err != nil {
						log.Error(err)
						time.Sleep(5 * time.Second)
						if strings.Contains(err.Error(), "-1125") {
							//如果是-1125错误，则Post更新
							err = ws.listenKeyPost()
						} else {
							err = ws.listenKeyPut()
						}
					}
				case <-*ws.listenKeyRefreshStopChan:
					return
				}
			}
		}()
	}

	return ws, nil
}

func (ws *SpotWsStreamClient) listenKeyPost() error {
	//申请新的listenKey
	switch ws.spotWsType {
	case SPOT_WS_TYPE:
		res, err := ws.client.NewSpotUserDataStreamPost().Do()
		if err != nil {
			log.Error(err)
			return err
		}
		ws.listenKey = res.ListenKey
		log.Debug("listenKey:", ws.listenKey)
	case SPOT_MARGIN_WS_TYPE:
		res, err := ws.client.NewSpotMarginUserDataStreamPost().Do()
		if err != nil {
			log.Error(err)
			return err
		}
		ws.listenKey = res.ListenKey
		log.Debug("listenKey:", ws.listenKey)
	case SPOT_ISOLATED_MARGIN_WS_TYPE:
		res, err := ws.client.NewSpotMarginIsolatedUserDataStreamPost().Do()
		if err != nil {
			log.Error(err)
			return err
		}
		ws.listenKey = res.ListenKey
		log.Debug("listenKey:", ws.listenKey)
	default:
		return fmt.Errorf("spotWsType is not support")
	}
	return nil
}

func (ws *SpotWsStreamClient) listenKeyPut() error {
	//申请新的listenKey
	switch ws.spotWsType {
	case SPOT_WS_TYPE:
		_, err := ws.client.NewSpotUserDataStreamPut().ListenKey(ws.listenKey).Do()
		if err != nil {
			log.Error(err)
			return err
		}
	case SPOT_MARGIN_WS_TYPE:
		_, err := ws.client.NewSpotMarginUserDataStreamPut().ListenKey(ws.listenKey).Do()
		if err != nil {
			log.Error(err)
			return err
		}

	case SPOT_ISOLATED_MARGIN_WS_TYPE:
		_, err := ws.client.NewSpotMarginIsolatedUserDataStreamPut().Symbol(ws.isolatedSymbol).ListenKey(ws.listenKey).Do()
		if err != nil {
			log.Error(err)
			return err
		}
	default:
		return fmt.Errorf("spotWsType is not support")
	}
	return nil
}

func (ws *SpotWsStreamClient) listenKeyDelete() error {
	switch ws.spotWsType {
	case SPOT_WS_TYPE:
		_, err := ws.client.NewSpotUserDataStreamDelete().ListenKey(ws.listenKey).Do()
		if err != nil {
			log.Error(err)
			return err
		}
	case SPOT_MARGIN_WS_TYPE:
		_, err := ws.client.NewSpotMarginUserDataStreamDelete().ListenKey(ws.listenKey).Do()
		if err != nil {
			log.Error(err)
			return err
		}

	case SPOT_ISOLATED_MARGIN_WS_TYPE:
		_, err := ws.client.NewSpotMarginIsolatedUserDataStreamDelete().Symbol(ws.isolatedSymbol).ListenKey(ws.listenKey).Do()
		if err != nil {
			log.Error(err)
			return err
		}
	default:
		return fmt.Errorf("spotWsType is not support")
	}
	return nil
}

func (ws *FutureWsStreamClient) ConvertToAccountWs(apiKey string, apiSecret string) (*FutureWsStreamClient, error) {
	ws.wsStreamPath = WS_ACCOUNT_PATH
	ws.apiKey = apiKey
	ws.apiSecret = apiSecret
	ws.isListenWs = true
	b := MyBinance{}
	ws.client = b.NewFutureRestClient(apiKey, apiSecret)

	err := ws.listenKeyPost()
	if err != nil {
		return nil, err
	}
	//创建一个协程定时刷新listenKey，如果已存在旧的刷新协程则不再创建
	if ws.listenKeyRefreshStopChan == nil {
		stopChan := make(chan struct{})
		ws.listenKeyRefreshStopChan = &stopChan
		go func() {
			for {
				select {
				case <-time.After(ListenKeyRefreshInterval):
					err := ws.listenKeyPut()
					for err != nil {
						log.Error(err)
						time.Sleep(5 * time.Second)
						if strings.Contains(err.Error(), "-1125") {
							//如果是-1125错误，则Post更新
							err = ws.listenKeyPost()
						} else {
							err = ws.listenKeyPut()
						}
					}
				case <-*ws.listenKeyRefreshStopChan:
					return
				}
			}
		}()
	}

	return ws, nil
}

func (ws *FutureWsStreamClient) listenKeyPost() error {
	res, err := ws.client.NewFutureListenKeyPost().Do()
	if err != nil {
		log.Error(err)
		return err
	}
	ws.listenKey = res.ListenKey
	log.Debug("listenKey:", ws.listenKey)
	return nil
}

func (ws *FutureWsStreamClient) listenKeyPut() error {
	_, err := ws.client.NewFutureListenKeyPut().Do()
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (ws *FutureWsStreamClient) listenKeyDelete() error {
	_, err := ws.client.NewFutureListenKeyDelete().Do()
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (ws *SwapWsStreamClient) ConvertToAccountWs(apiKey string, apiSecret string) (*SwapWsStreamClient, error) {
	ws.wsStreamPath = WS_ACCOUNT_PATH
	ws.apiKey = apiKey
	ws.apiSecret = apiSecret
	ws.isListenWs = true
	b := MyBinance{}
	ws.client = b.NewSwapRestClient(apiKey, apiSecret)

	err := ws.listenKeyPost()
	if err != nil {
		return nil, err
	}

	//创建一个协程定时刷新listenKey，如果已存在旧的刷新协程则不再创建
	if ws.listenKeyRefreshStopChan == nil {
		stopChan := make(chan struct{})
		ws.listenKeyRefreshStopChan = &stopChan
		go func() {
			for {
				select {
				case <-time.After(ListenKeyRefreshInterval):
					err := ws.listenKeyPut()
					for err != nil {
						log.Error(err)
						time.Sleep(5 * time.Second)
						if strings.Contains(err.Error(), "-1125") {
							//如果是-1125错误，则Post更新
							err = ws.listenKeyPost()
						} else {
							err = ws.listenKeyPut()
						}
					}
				case <-*ws.listenKeyRefreshStopChan:
					return
				}
			}
		}()
	}

	return ws, nil
}

func (ws *SwapWsStreamClient) listenKeyPost() error {
	res, err := ws.client.NewSwapListenKeyPost().Do()
	if err != nil {
		log.Error(err)
		return err
	}
	ws.listenKey = res.ListenKey
	log.Debug("listenKey:", ws.listenKey)
	return nil
}

func (ws *SwapWsStreamClient) listenKeyPut() error {
	_, err := ws.client.NewSwapListenKeyPut().Do()
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (ws *SwapWsStreamClient) listenKeyDelete() error {
	_, err := ws.client.NewSwapListenKeyDelete().Do()
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (ws *SpotWsStreamClient) Close() error {
	if ws.isListenWs {
		err := ws.listenKeyDelete()
		if err != nil {
			return err
		}
		*ws.listenKeyRefreshStopChan <- struct{}{}
	}
	return ws.WsStreamClient.Close()
}

func (ws *FutureWsStreamClient) Close() error {
	if ws.isListenWs {
		err := ws.listenKeyDelete()
		if err != nil {
			return err
		}
		*ws.listenKeyRefreshStopChan <- struct{}{}
	}
	return ws.WsStreamClient.Close()
}

func (ws *SwapWsStreamClient) Close() error {
	if ws.isListenWs {
		err := ws.listenKeyDelete()
		if err != nil {
			return err
		}
		*ws.listenKeyRefreshStopChan <- struct{}{}
	}
	return ws.WsStreamClient.Close()
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
		ws.conn = conn
		ws.isClose = false
		log.Info("OpenConn success to ", apiUrl)
		ws.handleResult(ws.resultChan, ws.errChan)
		return err
	} else {
		conn, err := wsStreamServe(apiUrl, ws.isGzip, ws.resultChan, ws.errChan)
		ws.conn = conn
		log.Info("Auto ReOpenConn success to ", apiUrl)
		return err
	}
}

type SpotWsStreamClient struct {
	WsStreamClient
	spotWsType     SpotWsType
	isolatedSymbol string
	client         *SpotRestClient
}
type FutureWsStreamClient struct {
	WsStreamClient
	client *FutureRestClient
}
type SwapWsStreamClient struct {
	WsStreamClient
	client *SwapRestClient
}

func (*MyBinance) NewSpotWsStreamClient() *SpotWsStreamClient {
	ws := &SpotWsStreamClient{
		WsStreamClient: WsStreamClient{
			apiType:       SPOT,
			isGzip:        false,
			reSubscribeMu: &sync.Mutex{},
			wsStreamPath:  WS_STREAM_PATH,
		},
	}
	ws.initStructs()
	return ws
}
func (*MyBinance) NewFutureWsStreamClient() *FutureWsStreamClient {
	ws := &FutureWsStreamClient{
		WsStreamClient: WsStreamClient{
			apiType:       FUTURE,
			isGzip:        false,
			reSubscribeMu: &sync.Mutex{},
			wsStreamPath:  WS_STREAM_PATH,
		},
	}
	ws.initStructs()
	return ws
}
func (*MyBinance) NewSwapWsStreamClient() *SwapWsStreamClient {
	ws := &SwapWsStreamClient{
		WsStreamClient: WsStreamClient{
			apiType:       SWAP,
			isGzip:        false,
			reSubscribeMu: &sync.Mutex{},
			wsStreamPath:  WS_STREAM_PATH,
		},
	}
	ws.initStructs()
	return ws
}

func (ws *WsStreamClient) sendSubscribeResultToChan(result SubscribeResult) {
	if result.Error != "" {
		if sub, ok := ws.commonSubMap[result.Id]; ok {
			sub.errChan <- fmt.Errorf("errHandler: %s===%v", result.Error, sub.Params)
		}
	} else {
		if sub, ok := ws.commonSubMap[result.Id]; ok {
			sub.resultChan <- result
		}
	}
}

func (ws *WsStreamClient) sendUnSubscribeSuccessToCloseChan(params []string) {
	isCloseMap := map[int64]bool{}
	for _, param := range params {
		if sub, ok := ws.klineSubMap[param]; ok {
			delete(ws.klineSubMap, param)
			if _, ok2 := isCloseMap[sub.ID]; ok2 {
				continue
			}
			sub.closeChan <- struct{}{}
			isCloseMap[sub.ID] = true
		} else if sub, ok := ws.depthSubMap[param]; ok {
			delete(ws.depthSubMap, param)
			if _, ok2 := isCloseMap[sub.ID]; ok2 {
				continue
			}
			sub.closeChan <- struct{}{}
			isCloseMap[sub.ID] = true
		} else if sub, ok := ws.aggTradeSubMap[param]; ok {
			delete(ws.aggTradeSubMap, param)
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
	for _, sub := range ws.commonSubMap {
		params = append(params, sub.Params...)
	}
	ws.sendUnSubscribeSuccessToCloseChan(params)
}

func (ws *WsStreamClient) reSubscribeForReconnect() error {
	ws.reSubscribeMu.Lock()
	defer ws.reSubscribeMu.Unlock()
	for _, sub := range ws.commonSubMap {
		log.Infof("ReSubscribe:{%d,%s,%v}", sub.ID, sub.Method, sub.Params)
		doSub, err := sendMsg[SubscribeResult](ws, 0, sub.Method, sub.Params)
		if err != nil {
			log.Error(err)
			return err
		}
		ws.commonSubMap[doSub.ID] = doSub

		select {
		case err := <-doSub.ErrChan():
			log.Error("SubscribeKline Error: ", err)
			return err
		case subResult := <-doSub.ResultChan():
			if subResult.Error != "" {
				log.Error(subResult.Error)
				return fmt.Errorf(subResult.Error)
			}
			log.Debugf("SubscribeKline Success: params:%v result:%v", doSub.Params, subResult)
		}
		log.Infof("reSubscribe Success: {%d,%s,%v}", sub.ID, sub.Method, sub.Params)
		sub.ID = doSub.ID
		delete(ws.commonSubMap, sub.ID)
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
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
					err := ws.OpenConn()
					for err != nil {
						time.Sleep(1000 * time.Millisecond)
						err = ws.OpenConn()
					}
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
					if sub, ok := ws.klineSubMap[param]; ok {
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
					if sub, ok := ws.depthSubMap[param]; ok {
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
					if sub, ok := ws.aggTradeSubMap[param]; ok {
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
					for _, payload := range ws.wsSpotPayloadMap {
						if payload.OutboundAccountPositionPayload != nil {
							payload.OutboundAccountPositionPayload.resultChan <- *res
						}
					}
				}

				//现货余额更新推送
				if strings.Contains(string(data), "balanceUpdate") {
					res, err := HandleWsPayloadResult[WsSpotPayloadBalanceUpdate](data)
					if err != nil {
						log.Error(err)
						continue
					}
					for _, payload := range ws.wsSpotPayloadMap {
						if payload.BalanceUpdatePayload != nil {
							payload.BalanceUpdatePayload.resultChan <- *res
						}
					}
				}

				//现货订单推送
				if strings.Contains(string(data), "executionReport") {
					res, err := HandleWsPayloadResult[WsSpotPayloadExecutionReport](data)
					if err != nil {
						log.Error(err)
						continue
					}
					for _, payload := range ws.wsSpotPayloadMap {
						if payload.ExecutionReportPayload != nil {
							payload.ExecutionReportPayload.resultChan <- *res
						}
					}
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
						for _, payload := range ws.wsFuturePayloadMap {
							if payload.AccountUpdatePayload != nil {
								payload.AccountUpdatePayload.resultChan <- *res
							}
						}
					case SWAP:
						res, err := HandleWsPayloadResult[WsSwapPayloadAccountUpdate](data)
						if err != nil {
							log.Error(err)
							continue
						}
						for _, payload := range ws.wsSwapPayloadMap {
							if payload.AccountUpdatePayload != nil {
								payload.AccountUpdatePayload.resultChan <- *res
							}
						}
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
						for _, payload := range ws.wsFuturePayloadMap {
							if payload.OrderTradeUpdatePayload != nil {
								payload.OrderTradeUpdatePayload.resultChan <- *res
							}
						}
					case SWAP:
						res, err := HandleWsPayloadResult[WsSwapPayloadOrderTradeUpdate](data)
						if err != nil {
							log.Error(err)
							continue
						}
						for _, payload := range ws.wsSwapPayloadMap {
							if payload.OrderTradeUpdatePayload != nil {
								payload.OrderTradeUpdatePayload.resultChan <- *res
							}
						}
					}
				}
			}
		}
	}()
}

// 订阅K线 如: "BTCUSDT","1m"
func (ws *WsStreamClient) SubscribeKline(symbol string, interval string) (*Subscription[WsKline], error) {
	param := getKlineSubscribeParam(symbol, interval)
	params := []string{param}
	doSub, err := sendMsg[SubscribeResult](ws, 0, SUBSCRIBE, params)
	if err != nil {
		return nil, err
	}
	ws.commonSubMap[doSub.ID] = doSub

	select {
	case err := <-doSub.ErrChan():
		log.Error("SubscribeKline Error: ", err)
		return nil, err
	case subResult := <-doSub.ResultChan():
		if subResult.Error != "" {
			log.Error(subResult.Error)
			return nil, fmt.Errorf(subResult.Error)
		}
		log.Infof("SubscribeKline Success: params:%v result:%v", doSub.Params, subResult)
		sub := &Subscription[WsKline]{
			ID:         doSub.ID,
			Method:     SUBSCRIBE,
			Params:     params,
			resultChan: make(chan WsKline),
			errChan:    make(chan error),
			closeChan:  make(chan struct{}),
			ws:         ws,
		}
		ws.klineSubMap[param] = sub
		return sub, nil
	}
}

// 批量订阅K线 如: []string{"BTCUSDT","ETHUSDT"},[]string{"1m","5m"}
func (ws *WsStreamClient) SubscribeKlineMultiple(symbols []string, intervals []string) (*Subscription[WsKline], error) {
	params := []string{}
	for _, symbol := range symbols {
		for _, interval := range intervals {
			param := getKlineSubscribeParam(symbol, interval)
			params = append(params, param)
		}
	}

	doSub, err := sendMsg[SubscribeResult](ws, 0, SUBSCRIBE, params)
	if err != nil {
		return nil, err
	}
	ws.commonSubMap[doSub.ID] = doSub

	select {
	case err := <-doSub.ErrChan():
		log.Error("SubscribeKline Error: ", err)
		return nil, err
	case subResult := <-doSub.ResultChan():
		if subResult.Error != "" {
			log.Error(subResult.Error)
			return nil, fmt.Errorf(subResult.Error)
		}
		log.Infof("SubscribeKline Success Params:%v Result:%v", doSub.Params, subResult)
		resultChan := make(chan WsKline)
		errChan := make(chan error)
		closeChan := make(chan struct{})
		sub := &Subscription[WsKline]{
			ID:         doSub.ID,
			Method:     SUBSCRIBE,
			Params:     params,
			resultChan: resultChan,
			errChan:    errChan,
			closeChan:  closeChan,
			ws:         ws,
		}
		for _, param := range params {
			ws.klineSubMap[param] = sub
		}
		return sub, nil
	}
}

// 订阅有限档深度 如: "BTCUSDT","20","100ms"
func (ws *WsStreamClient) SubscribeLevelDepth(symbol string, level string, USpeed string) (*Subscription[WsDepth], error) {
	param := getLevelDepthSubscribeParam(symbol, level, USpeed)
	params := []string{param}
	doSub, err := sendMsg[SubscribeResult](ws, 0, SUBSCRIBE, params)
	if err != nil {
		return nil, err
	}
	ws.commonSubMap[doSub.ID] = doSub

	select {
	case err := <-doSub.ErrChan():
		log.Error("SubscribeLevelDepth Error: ", err)
		return nil, err
	case subResult := <-doSub.ResultChan():
		if subResult.Error != "" {
			log.Error(subResult.Error)
			return nil, fmt.Errorf(subResult.Error)
		}
		log.Infof("SubscribeLevelDepth Success: params:%v result:%v", doSub.Params, subResult)
		sub := &Subscription[WsDepth]{
			ID:         doSub.ID,
			Method:     SUBSCRIBE,
			Params:     params,
			resultChan: make(chan WsDepth),
			errChan:    make(chan error),
			closeChan:  make(chan struct{}),
			ws:         ws,
		}

		ws.depthSubMap[param] = sub
		return sub, nil
	}
}

// 批量订阅有限档深度 如: []string{"BTCUSDT","ETHUSDT"},"20","100ms"
func (ws *WsStreamClient) SubscribeLevelDepthMultiple(symbols []string, level string, USpeed string) (*Subscription[WsDepth], error) {
	params := []string{}
	for _, symbol := range symbols {
		param := getLevelDepthSubscribeParam(symbol, level, USpeed)
		params = append(params, param)
	}
	doSub, err := sendMsg[SubscribeResult](ws, 0, SUBSCRIBE, params)
	if err != nil {
		return nil, err
	}
	ws.commonSubMap[doSub.ID] = doSub

	select {
	case err := <-doSub.ErrChan():
		log.Error("SubscribeLevelDepth Error: ", err)
		return nil, err
	case subResult := <-doSub.ResultChan():
		if subResult.Error != "" {
			log.Error(subResult.Error)
			return nil, fmt.Errorf(subResult.Error)
		}
		log.Infof("SubscribeLevelDepth Success: params:%v result:%v", doSub.Params, subResult)
		sub := &Subscription[WsDepth]{
			ID:         doSub.ID,
			Method:     SUBSCRIBE,
			Params:     params,
			resultChan: make(chan WsDepth),
			errChan:    make(chan error),
			closeChan:  make(chan struct{}),
			ws:         ws,
		}
		for _, param := range params {
			ws.depthSubMap[param] = sub
		}
		return sub, nil
	}
}

// 订阅增量深度 如: "BTCUSDT","100ms"
func (ws *WsStreamClient) SubscribeIncrementDepth(symbol string, USpeed string) (*Subscription[WsDepth], error) {
	param := getIncrementDepthSubscribeParam(symbol, USpeed)
	params := []string{param}
	doSub, err := sendMsg[SubscribeResult](ws, 0, SUBSCRIBE, params)
	if err != nil {
		return nil, err
	}
	ws.commonSubMap[doSub.ID] = doSub

	select {
	case err := <-doSub.ErrChan():
		log.Error("SubscribeIncrementDepth Error: ", err)
		return nil, err
	case subResult := <-doSub.ResultChan():
		if subResult.Error != "" {
			log.Error(subResult.Error)
			return nil, fmt.Errorf(subResult.Error)
		}
		log.Infof("SubscribeIncrementDepth Success: params:%v result:%v", doSub.Params, subResult)
		sub := &Subscription[WsDepth]{
			ID:         doSub.ID,
			Method:     SUBSCRIBE,
			Params:     params,
			resultChan: make(chan WsDepth),
			errChan:    make(chan error),
			closeChan:  make(chan struct{}),
			ws:         ws,
		}

		ws.depthSubMap[param] = sub
		return sub, nil
	}
}

// 批量订阅增量深度 如: []string{"BTCUSDT","ETHUSDT"},"100ms"
func (ws *WsStreamClient) SubscribeIncrementDepthMultiple(symbols []string, USpeed string) (*Subscription[WsDepth], error) {
	params := []string{}
	for _, symbol := range symbols {
		param := getIncrementDepthSubscribeParam(symbol, USpeed)
		params = append(params, param)
	}

	doSub, err := sendMsg[SubscribeResult](ws, 0, SUBSCRIBE, params)
	if err != nil {
		return nil, err
	}
	ws.commonSubMap[doSub.ID] = doSub

	select {
	case err := <-doSub.ErrChan():
		log.Error("SubscribeIncrementDepth Error: ", err)
		return nil, err
	case subResult := <-doSub.ResultChan():
		if subResult.Error != "" {
			log.Error(subResult.Error)
			return nil, fmt.Errorf(subResult.Error)
		}
		log.Infof("SubscribeIncrementDepth Success: params:%v result:%v", doSub.Params, subResult)
		sub := &Subscription[WsDepth]{
			ID:         doSub.ID,
			Method:     SUBSCRIBE,
			Params:     params,
			resultChan: make(chan WsDepth),
			errChan:    make(chan error),
			closeChan:  make(chan struct{}),
			ws:         ws,
		}
		for _, param := range params {
			ws.depthSubMap[param] = sub
		}
		return sub, nil
	}
}

// 订阅归集交易流 如: "BTCUSDT"
func (ws *WsStreamClient) SubscribeAggTrade(symbol string) (*Subscription[WsAggTrade], error) {
	param := getAggTradeParam(symbol)
	params := []string{param}
	doSub, err := sendMsg[SubscribeResult](ws, 0, SUBSCRIBE, params)
	if err != nil {
		return nil, err
	}
	ws.commonSubMap[doSub.ID] = doSub

	select {
	case err := <-doSub.ErrChan():
		log.Error("SubscribeAggTrade Error: ", err)
		return nil, err
	case subResult := <-doSub.ResultChan():
		if subResult.Error != "" {
			log.Error(subResult.Error)
			return nil, fmt.Errorf(subResult.Error)
		}
		log.Infof("SubscribeAggTrade Success: params:%v result:%v", doSub.Params, subResult)
		sub := &Subscription[WsAggTrade]{
			ID:         doSub.ID,
			Method:     SUBSCRIBE,
			Params:     params,
			resultChan: make(chan WsAggTrade),
			errChan:    make(chan error),
			closeChan:  make(chan struct{}),
			ws:         ws,
		}

		ws.aggTradeSubMap[param] = sub
		return sub, nil
	}
}

// 批量订阅归集交易流 如: []string{"BTCUSDT","ETHUSDT"}
func (ws *WsStreamClient) SubscribeAggTradeMultiple(symbols []string) (*Subscription[WsAggTrade], error) {
	params := []string{}
	for _, symbol := range symbols {
		param := getAggTradeParam(symbol)
		params = append(params, param)
	}
	doSub, err := sendMsg[SubscribeResult](ws, 0, SUBSCRIBE, params)
	if err != nil {
		return nil, err
	}
	ws.commonSubMap[doSub.ID] = doSub

	select {
	case err := <-doSub.ErrChan():
		log.Error("SubscribeAggTrade Error: ", err)
		return nil, err
	case subResult := <-doSub.ResultChan():
		if subResult.Error != "" {
			log.Error(subResult.Error)
			return nil, fmt.Errorf(subResult.Error)
		}
		log.Infof("SubscribeAggTrade Success: params:%v result:%v", doSub.Params, subResult)
		sub := &Subscription[WsAggTrade]{
			ID:         doSub.ID,
			Method:     SUBSCRIBE,
			Params:     params,
			resultChan: make(chan WsAggTrade),
			errChan:    make(chan error),
			closeChan:  make(chan struct{}),
			ws:         ws,
		}
		for _, param := range params {
			ws.aggTradeSubMap[param] = sub
		}
		return sub, nil
	}
}

func (ws *SpotWsStreamClient) CreatePayload() (*WsSpotPayload, error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return nil, err
	}
	id := node.Generate().Int64()
	payload := &WsSpotPayload{
		Ws: ws,
		Id: id,
		OutboundAccountPositionPayload: &Payload[WsSpotPayloadOutboundAccountPosition]{
			resultChan: make(chan WsSpotPayloadOutboundAccountPosition),
			errChan:    make(chan error),
			closeChan:  make(chan struct{}),
		},
		BalanceUpdatePayload: &Payload[WsSpotPayloadBalanceUpdate]{
			resultChan: make(chan WsSpotPayloadBalanceUpdate),
			errChan:    make(chan error),
			closeChan:  make(chan struct{}),
		},
		ExecutionReportPayload: &Payload[WsSpotPayloadExecutionReport]{
			resultChan: make(chan WsSpotPayloadExecutionReport),
			errChan:    make(chan error),
			closeChan:  make(chan struct{}),
		},
	}
	ws.wsSpotPayloadMap[id] = payload
	return payload, nil
}

func (ws *FutureWsStreamClient) CreatePayload() (*WsFuturePayload, error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return nil, err
	}
	id := node.Generate().Int64()
	payload := &WsFuturePayload{
		Ws: ws,
		Id: id,
		AccountUpdatePayload: &Payload[WsFuturePayloadAccountUpdate]{
			resultChan: make(chan WsFuturePayloadAccountUpdate),
			errChan:    make(chan error),
			closeChan:  make(chan struct{}),
		},
		OrderTradeUpdatePayload: &Payload[WsFuturePayloadOrderTradeUpdate]{
			resultChan: make(chan WsFuturePayloadOrderTradeUpdate),
			errChan:    make(chan error),
			closeChan:  make(chan struct{}),
		},
	}
	ws.wsFuturePayloadMap[id] = payload
	return payload, nil
}

func (ws *SwapWsStreamClient) CreatePayload() (*WsSwapPayload, error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return nil, err
	}
	id := node.Generate().Int64()
	payload := &WsSwapPayload{
		Ws: ws,
		Id: id,
		AccountUpdatePayload: &Payload[WsSwapPayloadAccountUpdate]{
			resultChan: make(chan WsSwapPayloadAccountUpdate),
			errChan:    make(chan error),
			closeChan:  make(chan struct{}),
		},
		OrderTradeUpdatePayload: &Payload[WsSwapPayloadOrderTradeUpdate]{
			resultChan: make(chan WsSwapPayloadOrderTradeUpdate),
			errChan:    make(chan error),
			closeChan:  make(chan struct{}),
		},
	}
	ws.wsSwapPayloadMap[id] = payload
	return payload, nil
}

// 获取当前所有订阅
func (ws *WsStreamClient) CurrentSubscribeList() ([]string, error) {
	listSub, err := sendMsg[SubscribeResult](ws, 0, LIST_SUBSCRIPTIONS, []string{})
	if err != nil {
		return nil, err
	}
	ws.commonSubMap[listSub.ID] = listSub
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
		delete(ws.commonSubMap, listSub.ID)
		return subResult.Result, nil
	}
}

// 取消订阅
func (sub *Subscription[T]) Unsubscribe() error {
	unSub, err := sendMsg[SubscribeResult](sub.ws, 0, UNSUBSCRIBE, sub.Params)
	if err != nil {
		return err
	}
	sub.ws.commonSubMap[unSub.ID] = unSub

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
		delete(sub.ws.commonSubMap, unSub.ID)
		delete(sub.ws.commonSubMap, sub.ID)
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
				reader.Close()
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
	u := url.URL{
		Scheme:   "wss",
		Host:     getWsApi(apiType, isGzip),
		Path:     getWsPath(wsStreamPath, listenKey),
		RawQuery: "",
	}
	return u.String()
}

// 获取数据流请求Path
func getWsApi(apiType ApiType, isGzip bool) string {
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

func getWsPath(wsStreamPath WsStreamPath, listenKey string) string {
	switch wsStreamPath {
	case WS_STREAM_PATH:
		return "/stream"
	case WS_ACCOUNT_PATH:
		return fmt.Sprintf("/ws/%s", listenKey)
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
				c.Close()
				return
			}
		}
	}()
}
func getKlineSubscribeParam(symbol string, interval string) string {
	return fmt.Sprintf("%s@kline_%s", strings.ToLower(symbol), interval)
}
func getLevelDepthSubscribeParam(symbol string, level string, USpeed string) string {
	if USpeed != "" {
		return fmt.Sprintf("%s@depth%s@%s", strings.ToLower(symbol), level, USpeed)
	} else {
		return fmt.Sprintf("%s@depth%s", strings.ToLower(symbol), level)
	}
}
func getIncrementDepthSubscribeParam(symbol string, USpeed string) string {
	if USpeed != "" {
		return fmt.Sprintf("%s@depth@%s", strings.ToLower(symbol), USpeed)
	} else {
		return fmt.Sprintf("%s@depth", strings.ToLower(symbol))
	}

}
func getAggTradeParam(symbol string) string {
	return fmt.Sprintf("%s@aggTrade", strings.ToLower(symbol))
}
