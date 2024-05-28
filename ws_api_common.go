package mybinanceapi

import (
	"context"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/gorilla/websocket"
	"strconv"
	"time"
)

type WsApiResultChan interface {
	ResultChan() chan []byte
	ErrChan() chan error
	CloseChan() chan struct{}
}

type WsApiReq[T any] struct {
	Id     string `json:"id"`     //YES	任意的 ID 用于匹配对请求的响应
	Method string `json:"method"` //YES	请求函数名称
	Params T      `json:"params"` //YES	请求参数
}

type WsApiResult[T any] struct {
	Id         string          `json:"id"`     // YES	与原来请求的ID一样
	Status     int             `json:"status"` // YES	响应状态。请看 状态代码
	Result     T               `json:"result"` // YES	响应内容。请求成功则显示
	Error      BinanceErrorRes `json:"error"`  // 请求失败则显示
	RateLimits []RateLimit     `json:"rateLimits"`
}

type RateLimit struct {
	RateLimitType string `json:"rateLimitType"`
	Interval      string `json:"interval"`
	IntervalNum   int    `json:"intervalNum"`
	Limit         int    `json:"limit"`
	Count         int    `json:"count"`
}

type WsApiResultChanImpl struct {
	resultChan chan []byte
	errChan    chan error
	closeChan  chan struct{}
}

func (w *WsApiResultChanImpl) ResultChan() chan []byte {
	return w.resultChan
}
func (w *WsApiResultChanImpl) ErrChan() chan error {
	return w.errChan
}
func (w *WsApiResultChanImpl) CloseChan() chan struct{} {
	return w.closeChan
}

// 转换为Ws API 可以进行ws下单等操作
func (ws *SpotWsStreamClient) ConvertToWsApi(apiKey string, apiSecret string) (*SpotWsStreamClient, error) {
	ws.wsStreamPath = WS_SPOT_API_PATH
	ws.apiKey = apiKey
	ws.apiSecret = apiSecret
	ws.isListenWs = true
	b := MyBinance{}
	ws.client = b.NewSpotRestClient(apiKey, apiSecret)
	return ws, nil
}

func (ws *FutureWsStreamClient) ConvertToWsApi(apiKey string, apiSecret string) (*FutureWsStreamClient, error) {
	ws.wsStreamPath = WS_FUTURE_API_PATH
	ws.apiKey = apiKey
	ws.apiSecret = apiSecret
	ws.isListenWs = true
	b := MyBinance{}
	ws.client = b.NewFutureRestClient(apiKey, apiSecret)
	return ws, nil
}

func sendApiMsg[T any, R any](ws *WsStreamClient, id int64, method string, params T) (*WsApiResult[R], error) {
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

	wsApiReq := WsApiReq[T]{
		Id:     strconv.FormatInt(id, 10),
		Method: method,
		Params: params,
	}
	d, _ := json.Marshal(wsApiReq)
	log.Debugf("send api msg: %s", string(d))
	result := &WsApiResultChanImpl{
		resultChan: make(chan []byte, 1),
		errChan:    make(chan error, 1),
		closeChan:  make(chan struct{}, 1),
	}
	ws.waitWsApiResultMap.Store(wsApiReq.Id, result)
	data, err := json.Marshal(wsApiReq)
	if err != nil {
		return nil, err
	}
	err = ws.conn.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		return nil, err
	}

	//阻塞等待，且5秒后超时
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resultChan := make(chan []byte, 1)
	errChan := make(chan error, 1)
	closeChan := make(chan struct{}, 1)
	go func() {
		select {
		case <-ctx.Done():
			log.Info("api msg timeout")
		case nErr := <-result.ErrChan():
			errChan <- nErr
		case r := <-result.ResultChan():
			resultChan <- r
		case <-result.CloseChan():
			closeChan <- struct{}{}
		}
	}()

	select {
	case result := <-resultChan:
		res := &WsApiResult[R]{}
		err = json.Unmarshal(result, res)
		if err != nil {
			return nil, err
		}
		return res, nil
	case nErr := <-errChan:
		return nil, nErr
	case <-closeChan:
		return nil, fmt.Errorf("ws is closed")
	case <-ctx.Done():
		return nil, fmt.Errorf("api msg timeout")
	}
}
