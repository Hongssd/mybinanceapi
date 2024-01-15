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
	WebsocketTimeout = time.Second * 10
	// WebsocketKeepalive enables sending ping/pong messages to check the connection stability
	WebsocketKeepalive      = true
	SUBSCRIBE_INTERVAL_TIME = 500 * time.Millisecond //订阅间隔
)

// 数据流订阅基础客户端
type WsStreamClient struct {
	apiType        ApiType
	isGzip         bool
	conn           *websocket.Conn
	commonSubMap   map[int64]*Subscription[SubscribeResult] //订阅的返回结果
	klineSubMap    map[string]*Subscription[WsKline]
	depthSubMap    map[string]*Subscription[WsDepth]
	aggTradeSubMap map[string]*Subscription[WsAggTrade]
	resultChan     chan []byte
	errChan        chan error
	isClose        bool
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
	ws.sendWsCloseToAllSub()

	//初始化连接状态
	ws.conn = nil
	close(ws.resultChan)
	close(ws.errChan)
	ws.resultChan = nil
	ws.errChan = nil
	ws.commonSubMap = make(map[int64]*Subscription[SubscribeResult])
	ws.klineSubMap = make(map[string]*Subscription[WsKline])
	ws.depthSubMap = make(map[string]*Subscription[WsDepth])
	ws.aggTradeSubMap = make(map[string]*Subscription[WsAggTrade])
	return nil
}

func (ws *WsStreamClient) OpenConn() error {
	if ws.resultChan == nil {
		ws.resultChan = make(chan []byte)
	}
	if ws.errChan == nil {
		ws.errChan = make(chan error)
	}
	apiUrl := handlerWsStreamRequestApi(ws.apiType, ws.isGzip)
	if ws.conn == nil {
		conn, err := wsStreamServe(apiUrl, ws.isGzip, ws.resultChan, ws.errChan)
		ws.conn = conn
		ws.isClose = false
		log.Debug("OpenConn success to ", apiUrl)
		ws.handleResult(ws.resultChan, ws.errChan)
		return err
	} else {
		conn, err := wsStreamServe(apiUrl, ws.isGzip, ws.resultChan, ws.errChan)
		ws.conn = conn
		log.Debug("Auto ReOpenConn success to ", apiUrl)
		return err
	}
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
			apiType:        SPOT,
			isGzip:         false,
			commonSubMap:   make(map[int64]*Subscription[SubscribeResult]),
			klineSubMap:    make(map[string]*Subscription[WsKline]),
			depthSubMap:    make(map[string]*Subscription[WsDepth]),
			aggTradeSubMap: make(map[string]*Subscription[WsAggTrade]),
		},
	}
}
func (*MyBinance) NewFutureWsStreamClient() *FutureWsStreamClient {
	return &FutureWsStreamClient{
		WsStreamClient: WsStreamClient{
			apiType:        FUTURE,
			isGzip:         false,
			commonSubMap:   make(map[int64]*Subscription[SubscribeResult]),
			klineSubMap:    make(map[string]*Subscription[WsKline]),
			depthSubMap:    make(map[string]*Subscription[WsDepth]),
			aggTradeSubMap: make(map[string]*Subscription[WsAggTrade]),
		},
	}
}
func (*MyBinance) NewSwapWsStreamClient() *SwapWsStreamClient {
	return &SwapWsStreamClient{
		WsStreamClient: WsStreamClient{
			apiType:        SWAP,
			isGzip:         false,
			commonSubMap:   make(map[int64]*Subscription[SubscribeResult]),
			klineSubMap:    make(map[string]*Subscription[WsKline]),
			depthSubMap:    make(map[string]*Subscription[WsDepth]),
			aggTradeSubMap: make(map[string]*Subscription[WsAggTrade]),
		},
	}
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
	for _, sub := range ws.commonSubMap {
		log.Infof("ReSubscribe:{%d,%s,%v}", sub.ID, sub.Method, sub.Params)
		resultSub, err := sendMsg[SubscribeResult](ws, 0, sub.Method, sub.Params)
		if err != nil {
			log.Error(err)
			continue
		}
		sub.ID = resultSub.ID
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
						time.Sleep(500 * time.Millisecond)
						err = ws.OpenConn()
					}
					ws.reSubscribeForReconnect()
				} else {
					continue
				}
			case data, ok := <-resultChan:
				if !ok {
					return
				}
				// log.Debug("receive result: ", string(data))
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
		log.Debugf("SubscribeKline Success: params:%v result:%v", doSub.Params, subResult)
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
		log.Debugf("SubscribeKline Success Params:%v Result:%v", doSub.Params, subResult)
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
		log.Debugf("SubscribeLevelDepth Success: params:%v result:%v", doSub.Params, subResult)
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
		log.Debugf("SubscribeLevelDepth Success: params:%v result:%v", doSub.Params, subResult)
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
		log.Debugf("SubscribeIncrementDepth Success: params:%v result:%v", doSub.Params, subResult)
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
		log.Debugf("SubscribeIncrementDepth Success: params:%v result:%v", doSub.Params, subResult)
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
		log.Debugf("SubscribeAggTrade Success: params:%v result:%v", doSub.Params, subResult)
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
		log.Debugf("SubscribeAggTrade Success: params:%v result:%v", doSub.Params, subResult)
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
		log.Debugf("Unsubscribe Success Params:%v Result:%v", unSub.Params, subResult)

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
