package mybinanceapi

import (
	"fmt"
	"strings"
)

func getAggTradeParam(symbol string) string {
	return fmt.Sprintf("%s@aggTrade", strings.ToLower(symbol))
}

// 订阅归集交易流 如: "BTCUSDT"
func (ws *WsStreamClient) SubscribeAggTrade(symbol string) (*Subscription[WsAggTrade], error) {
	return ws.SubscribeAggTradeMultiple([]string{symbol})
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
	ws.commonSubMap.Store(doSub.ID, doSub)

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
			ws.aggTradeSubMap.Store(param, sub)
		}
		return sub, nil
	}
}
