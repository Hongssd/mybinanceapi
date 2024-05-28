package mybinanceapi

import (
	"fmt"
	"strings"
)

func getKlineSubscribeParam(symbol string, interval string) string {
	return fmt.Sprintf("%s@kline_%s", strings.ToLower(symbol), interval)
}

// 订阅K线 如: "BTCUSDT","1m"
func (ws *WsStreamClient) SubscribeKline(symbol string, interval string) (*Subscription[WsKline], error) {
	return ws.SubscribeKlineMultiple([]string{symbol}, []string{interval})
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
	ws.commonSubMap.Store(doSub.ID, doSub)

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
			ws.klineSubMap.Store(param, sub)
		}
		return sub, nil
	}
}
