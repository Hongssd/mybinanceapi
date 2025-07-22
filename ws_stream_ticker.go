package mybinanceapi

import (
	"fmt"
	"strings"
)

func getTickerSubscribeParam(symbol string) string {
	return fmt.Sprintf("%s@ticker", strings.ToLower(symbol))
}

func getTickerAllSubscribeParam() string {
	return "!ticker@arr"
}

// 订阅单个Symbol的Ticker 如: "BTCUSDT"
func (ws *WsStreamClient) SubscribeTicker(symbol string) (*Subscription[WsTicker], error) {
	return ws.SubscribeTickerMultiple([]string{symbol})
}

// 批量订阅多个Symbol的Ticker 如: []string{"BTCUSDT","ETHUSDT"}
func (ws *WsStreamClient) SubscribeTickerMultiple(symbols []string) (*Subscription[WsTicker], error) {
	params := []string{}
	for _, symbol := range symbols {
		param := getTickerSubscribeParam(symbol)
		params = append(params, param)
	}

	doSub, err := sendMsg[SubscribeResult](ws, 0, SUBSCRIBE, params)
	if err != nil {
		return nil, err
	}
	ws.commonSubMap.Store(doSub.ID, doSub)

	select {
	case err := <-doSub.ErrChan():
		log.Error("SubscribeTicker Error: ", err)
		return nil, err
	case subResult := <-doSub.ResultChan():
		if subResult.Error != "" {
			log.Error(subResult.Error)
			return nil, fmt.Errorf(subResult.Error)
		}
		log.Infof("SubscribeTicker Success Params:%v Result:%v", doSub.Params, subResult)
		resultChan := make(chan WsTicker)
		errChan := make(chan error)
		closeChan := make(chan struct{})
		sub := &Subscription[WsTicker]{
			ID:         doSub.ID,
			Method:     SUBSCRIBE,
			Params:     params,
			resultChan: resultChan,
			errChan:    errChan,
			closeChan:  closeChan,
			ws:         ws,
		}
		for _, param := range params {
			ws.tickerSubMap.Store(param, sub)
		}
		return sub, nil
	}
}

// 订阅全市场所有交易对的Ticker（数组形式推送）
func (ws *WsStreamClient) SubscribeTickerAll() (*Subscription[[]*WsTicker], error) {
	param := getTickerAllSubscribeParam()
	params := []string{param}

	doSub, err := sendMsg[SubscribeResult](ws, 0, SUBSCRIBE, params)
	if err != nil {
		return nil, err
	}
	ws.commonSubMap.Store(doSub.ID, doSub)

	select {
	case err := <-doSub.ErrChan():
		log.Error("SubscribeTickerAll Error: ", err)
		return nil, err
	case subResult := <-doSub.ResultChan():
		if subResult.Error != "" {
			log.Error(subResult.Error)
			return nil, fmt.Errorf(subResult.Error)
		}
		log.Infof("SubscribeTickerAll Success Params:%v Result:%v", doSub.Params, subResult)
		resultChan := make(chan []*WsTicker)
		errChan := make(chan error)
		closeChan := make(chan struct{})
		sub := &Subscription[[]*WsTicker]{
			ID:         doSub.ID,
			Method:     SUBSCRIBE,
			Params:     params,
			resultChan: resultChan,
			errChan:    errChan,
			closeChan:  closeChan,
			ws:         ws,
		}
		ws.tickerAllSubMap.Store(param, sub)
		return sub, nil
	}
}
