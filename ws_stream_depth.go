package mybinanceapi

import (
	"fmt"
	"strings"
)

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

// 订阅有限档深度 如: "BTCUSDT","20","100ms"
func (ws *WsStreamClient) SubscribeLevelDepth(symbol string, level string, USpeed string) (*Subscription[WsDepth], error) {
	return ws.SubscribeLevelDepthMultiple([]string{symbol}, level, USpeed)
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
	ws.commonSubMap.Store(doSub.ID, doSub)

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
			ws.depthSubMap.Store(param, sub)
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
	ws.commonSubMap.Store(doSub.ID, doSub)

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

		ws.depthSubMap.Store(param, sub)
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
	ws.commonSubMap.Store(doSub.ID, doSub)

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
			ws.depthSubMap.Store(param, sub)
		}
		return sub, nil
	}
}
