package mybinanceapi

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"net/url"
	"strings"
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
)

const (
	SUBSCRIBE          = "SUBSCRIBE"
	UNSUBSCRIBE        = "UNSUBSCRIBE"
	LIST_SUBSCRIPTIONS = "LIST_SUBSCRIPTIONS"
)

var (
	// WebsocketTimeout is an interval for sending ping/pong messages if WebsocketKeepalive is enabled
	WebsocketTimeout = time.Second * 60
	// WebsocketKeepalive enables sending ping/pong messages to check the connection stability
	WebsocketKeepalive      = true
	SUBSCRIBE_INTERVAL_TIME = 500 * time.Millisecond //订阅间隔
)

// 数据流订阅基础客户端
type WsStreamClient struct {
	apiType              ApiType
	isGzip               bool
	conn                 *websocket.Conn
	UnSubscribeMap       map[int64]*Subscription[SubscribeResult] //取消订阅的返回结果
	ListSubscriptionsMap map[int64]*Subscription[SubscribeParams] //查询当前所有订阅的返回结果
	CommonSubMap         map[int64]*Subscription[SubscribeResult] //订阅的返回结果
	KlineSubMap          map[string]*Subscription[WsKline]
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
	conn       *WsStreamClient //订阅的连接
	ID         int64           //此次订阅ID
	Method     string          //订阅方法
	Params     SubscribeParams //订阅参数
	resultChan chan T          //接收订阅结果的通道
	errChan    chan error      //接收订阅错误的通道
}

// 获取订阅结果
func (sub *Subscription[T]) ResultChan() <-chan T {
	return sub.resultChan
}

// 获取错误订阅
func (sub *Subscription[T]) ErrChan() <-chan error {
	return sub.errChan
}

func Subscribe[T any](ws *WsStreamClient, method string, params []string) (*Subscription[T], error) {
	node, err := snowflake.NewNode(1)
	id := node.Generate().Int64()
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
		resultChan: make(chan T, 1),
		errChan:    make(chan error, 1),
	}

	return result, nil
}

type SpotWsStreamClient struct {
	WsStreamClient
}
type FutureWsStreamClient struct {
	WsStreamClient
}
type SwapWsStreamClient struct {
	WsStreamClient
}

func (*MyBinance) NewSpotWsStreamClient() *SpotWsStreamClient {
	return &SpotWsStreamClient{
		WsStreamClient: WsStreamClient{
			apiType:              SPOT,
			isGzip:               false,
			UnSubscribeMap:       make(map[int64]*Subscription[SubscribeResult]),
			ListSubscriptionsMap: make(map[int64]*Subscription[SubscribeParams]),
			CommonSubMap:         make(map[int64]*Subscription[SubscribeResult]),
			KlineSubMap:          make(map[string]*Subscription[WsKline]),
		},
	}
}
func (*MyBinance) NewFutureWsStreamClient() *FutureWsStreamClient {
	return &FutureWsStreamClient{
		WsStreamClient: WsStreamClient{
			apiType:              FUTURE,
			isGzip:               IS_GZIP,
			UnSubscribeMap:       make(map[int64]*Subscription[SubscribeResult]),
			ListSubscriptionsMap: make(map[int64]*Subscription[SubscribeParams]),
			CommonSubMap:         make(map[int64]*Subscription[SubscribeResult]),
			KlineSubMap:          make(map[string]*Subscription[WsKline]),
		},
	}
}
func (*MyBinance) NewSwapWsStreamClient() *SwapWsStreamClient {
	return &SwapWsStreamClient{
		WsStreamClient: WsStreamClient{
			apiType:              SWAP,
			isGzip:               IS_GZIP,
			UnSubscribeMap:       make(map[int64]*Subscription[SubscribeResult]),
			ListSubscriptionsMap: make(map[int64]*Subscription[SubscribeParams]),
			CommonSubMap:         make(map[int64]*Subscription[SubscribeResult]),
			KlineSubMap:          make(map[string]*Subscription[WsKline]),
		},
	}
}

func (ws *WsStreamClient) SendSubscribeResultToChan(result SubscribeResult) {
	if result.Error != "" {
		if sub, ok := ws.UnSubscribeMap[result.Id]; ok {
			sub.errChan <- fmt.Errorf("errHandler: %s===%v", result.Error, sub.Params)
		}
		if sub, ok := ws.ListSubscriptionsMap[result.Id]; ok {
			sub.errChan <- fmt.Errorf("errHandler: %s===%v", result.Error, sub.Params)
		}
		if sub, ok := ws.CommonSubMap[result.Id]; ok {
			sub.errChan <- fmt.Errorf("errHandler: %s===%v", result.Error, sub.Params)
		}
	} else {
		if sub, ok := ws.UnSubscribeMap[result.Id]; ok {
			sub.resultChan <- result
		}
		if sub, ok := ws.ListSubscriptionsMap[result.Id]; ok {
			sub.resultChan <- result.Result
		}
		if sub, ok := ws.CommonSubMap[result.Id]; ok {
			sub.resultChan <- result
		}
	}
}

func (ws *WsStreamClient) HandleResult(resultChan chan []byte, errChan chan error) {
	go func() {
		for {
			log.Info("handle")
			select {
			case err := <-errChan:
				log.Error(err)
				//TODO:错误处理 重连等
				return
			case data := <-resultChan:
				//处理订阅或查询订阅列表请求返回结果
				if strings.Contains(string(data), "error") || strings.Contains(string(data), "result") {
					result := SubscribeResult{}
					err := json.Unmarshal(data, &result)
					if err != nil {
						log.Error(err)
						continue
					}
					ws.SendSubscribeResultToChan(result)
				}

				//处理正常数据的返回结果
				if strings.Contains(string(data), "kline") {

					var k *WsKline
					var err error
					//K线处理
					if !ws.isGzip || ws.apiType == SPOT {
						k, err = HandleWsCombinedKline(ws.apiType, data)
					} else {
						k, err = HandleWsCombinedKlineGzip(ws.apiType, data)
					}
					param := fmt.Sprintf("%s@kline_%s", k.Symbol, k.Interval)
					if sub, ok := ws.KlineSubMap[param]; ok {
						if err != nil {
							sub.errChan <- err
							continue
						}
						sub.resultChan <- *k
					}

				}
			}
		}
	}()
}

func (ws *WsStreamClient) Close() error {
	return ws.conn.Close()
}

func (ws *WsStreamClient) OpenConn() error {
	conn, resultChan, errChan, err := wsStreamServe(handlerWsStreamRequestApi(ws.apiType, ws.isGzip), ws.isGzip)
	ws.conn = conn
	ws.HandleResult(resultChan, errChan)
	errChan <- fmt.Errorf("aaa%d", 1)
	return err
}

// 订阅K线
func (ws *WsStreamClient) SubscribeKline(symbol string, interval string) (*Subscription[WsKline], error) {
	param := fmt.Sprintf("%s@kline_%s", symbol, interval)
	params := []string{param}
	doSub, err := Subscribe[SubscribeResult](ws, SUBSCRIBE, params)
	if err != nil {
		return nil, err
	}
	ws.CommonSubMap[doSub.ID] = doSub
	log.Info(doSub)
	select {
	case <-doSub.ErrChan():
		return nil, err
	case <-doSub.ResultChan():
		sub := &Subscription[WsKline]{
			ID:         doSub.ID,
			Method:     SUBSCRIBE,
			Params:     params,
			resultChan: make(chan WsKline, 1),
			errChan:    make(chan error, 1),
		}
		ws.KlineSubMap[param] = sub
		return sub, nil
	}
}

// 获取当前所有订阅
func (ws *WsStreamClient) CurrentSubscribeList() ([]string, error) {
	listSub, err := Subscribe[SubscribeParams](ws, LIST_SUBSCRIPTIONS, []string{})
	if err != nil {
		return nil, err
	}
	ws.ListSubscriptionsMap[listSub.ID] = listSub
	select {
	case <-listSub.ErrChan():
		return nil, err
	case list := <-listSub.ResultChan():
		return list, nil
	}
}

// 取消订阅
func (sub *Subscription[T]) Unsubscribe() error {
	unSub, err := Subscribe[SubscribeResult](sub.conn, UNSUBSCRIBE, sub.Params)
	if err != nil {
		return err
	}
	select {
	case <-unSub.ErrChan():
		return err
	case <-unSub.ResultChan():
		return nil
	}
}

// 标准订阅方法
func wsStreamServe(api string, isGzip bool) (*websocket.Conn, chan []byte, chan error, error) {
	resultChan := make(chan []byte, 1)
	errChan := make(chan error, 1)
	c, _, err := websocket.DefaultDialer.Dial(api, nil)
	if err != nil {
		return nil, nil, nil, err
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
	return c, nil, nil, err
}

// 获取数据流请求URL
func handlerWsStreamRequestApi(apiType ApiType, isGzip bool) string {
	u := url.URL{
		Scheme:   "wss",
		Host:     getWsApi(apiType, isGzip),
		Path:     "/stream",
		RawQuery: "",
	}
	return u.String()
}

// 获取数据流请求Path
func getWsApi(apiType ApiType, isGzip bool) string {
	switch apiType {
	case SPOT:
		return BINANCE_API_SPOT_WS_STREAM
	case SWAP:
		if isGzip {
			return BINANCE_API_SWAP_WS_STREAM_GZIP
		} else {
			return BINANCE_API_SWAP_WS_STREAM
		}
	case FUTURE:
		if isGzip {
			return BINANCE_API_FUTURE_WS_STREAM_GZIP
		} else {
			return BINANCE_API_FUTURE_WS_STREAM
		}
	default:
		log.Error("AccountType Error is ", apiType)
		return ""
	}
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
