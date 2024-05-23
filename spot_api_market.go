package mybinanceapi

// 行情接口
// binance SPOT spotTickerPrice rest价格 (NONE)
func (client *SpotRestClient) NewSpotTickerPrice() *SpotTickerPriceApi {
	return &SpotTickerPriceApi{
		client: client,
		req:    &SpotTickerPriceReq{},
	}
}
func (api *SpotTickerPriceApi) Do() (*SpotTickerPriceRes, error) {
	url := binanceHandlerRequestApi(SPOT, api.req, SpotApiMap[SpotTickerPrice])
	if api.req.Symbol != nil && *api.req.Symbol != "" {
		res, err := binanceCallApiWithSecret[SpotTickerPriceResRow](api.client.c, url, GET)
		if err != nil {
			return nil, err
		}
		return &SpotTickerPriceRes{*res}, nil
	} else {
		return binanceCallApiWithSecret[SpotTickerPriceRes](api.client.c, url, GET)
	}

}

// binance SPOT spotKlines restK线数据 (NONE)
func (client *SpotRestClient) NewSpotKlines() *SpotKlinesApi {
	return &SpotKlinesApi{
		client: client,
		req:    &SpotKlinesReq{},
	}
}
func (api *SpotKlinesApi) Do() (*KlinesRes, error) {
	url := binanceHandlerRequestApi(SPOT, api.req, SpotApiMap[SpotKlines])
	res, err := binanceCallApiWithSecret[KlinesMiddle](api.client.c, url, GET)
	if err != nil {
		return nil, err
	}
	res2 := res.ConvertToRes()
	return &res2, nil
}

// binance SPOT spotDepth rest深度信息 (NONE)
func (client *SpotRestClient) NewSpotDepth() *SpotDepthApi {
	return &SpotDepthApi{
		client: client,
		req:    &SpotDepthReq{},
	}
}
func (api *SpotDepthApi) Do() (*SpotDepthRes, error) {
	url := binanceHandlerRequestApi(SPOT, api.req, SpotApiMap[SpotDepth])
	res, err := binanceCallApiWithSecret[SpotDepthResMiddle](api.client.c, url, GET)
	if err != nil {
		return nil, err
	}
	return res.ConvertToRes(), nil
}

// binance SPOT SpotTrades rest最近成交 (NONE)
func (client *SpotRestClient) NewSpotTrades() *SpotTradesApi {
	return &SpotTradesApi{
		client: client,
		req:    &SpotTradesReq{},
	}
}
func (api *SpotTradesApi) Do() (*SpotTradesRes, error) {
	url := binanceHandlerRequestApi(SPOT, api.req, SpotApiMap[SpotTrades])
	return binanceCallApiWithSecret[SpotTradesRes](api.client.c, url, GET)
}

// binance SPOT spotHistoricalTrades rest历史成交 (NONE)
func (client *SpotRestClient) NewSpotHistoricalTrades() *SpotHistoricalTradesApi {
	return &SpotHistoricalTradesApi{
		client: client,
		req:    &SpotHistoricalTradesReq{},
	}
}
func (api *SpotHistoricalTradesApi) Do() (*SpotHistoricalTradesRes, error) {
	url := binanceHandlerRequestApi(SPOT, api.req, SpotApiMap[SpotHistoricalTrades])
	return binanceCallApiWithSecret[SpotHistoricalTradesRes](api.client.c, url, GET)
}

// binance SPOT spotAggTrades rest近期成交(归集)(NONE)
func (client *SpotRestClient) NewSpotAggTrades() *SpotAggTradesApi {
	return &SpotAggTradesApi{
		client: client,
		req:    &SpotAggTradesReq{},
	}
}
func (api *SpotAggTradesApi) Do() (*SpotAggTradesRes, error) {
	url := binanceHandlerRequestApi(SPOT, api.req, SpotApiMap[SpotAggTrades])
	return binanceCallApiWithSecret[SpotAggTradesRes](api.client.c, url, GET)
}

// binance SPOT spotAvgPrice rest当前平均价格 (NONE)
func (client *SpotRestClient) NewSpotAvgPrice() *SpotAvgPriceApi {
	return &SpotAvgPriceApi{
		client: client,
		req:    &SpotAvgPriceReq{},
	}
}
func (api *SpotAvgPriceApi) Do() (*SpotAvgPriceRes, error) {
	url := binanceHandlerRequestApi(SPOT, api.req, SpotApiMap[SpotAvgPrice])
	return binanceCallApiWithSecret[SpotAvgPriceRes](api.client.c, url, GET)
}

// binance SPOT spotUiKlines restUI K线数据 (NONE)
func (client *SpotRestClient) NewSpotUiKlines() *SpotUiKlinesApi {
	return &SpotUiKlinesApi{
		client: client,
		req:    &SpotKlinesReq{},
	}
}
func (api *SpotUiKlinesApi) Do() (*KlinesRes, error) {
	url := binanceHandlerRequestApi(SPOT, api.req, SpotApiMap[SpotKlines])
	res, err := binanceCallApiWithSecret[KlinesMiddle](api.client.c, url, GET)
	if err != nil {
		return nil, err
	}
	res2 := res.ConvertToRes()
	return &res2, nil
}

// binance SPOT spotTicker24hr rest24hr价格变动情况 (NONE)
func (client *SpotRestClient) NewSpotTicker24hr() *SpotTicker24hrApi {
	return &SpotTicker24hrApi{
		client: client,
		req:    &SpotTicker24hrReq{},
	}
}
func (api *SpotTicker24hrApi) Do() (*SpotTicker24hrRes, error) {
	url := binanceHandlerRequestApi(SPOT, api.req, SpotApiMap[SpotTicker24hr])
	if api.req.Symbol != nil && *api.req.Symbol != "" {
		res, err := binanceCallApiWithSecret[SpotTicker24hrResRow](api.client.c, url, GET)
		if err != nil {
			return nil, err
		}
		return &SpotTicker24hrRes{*res}, nil
	} else {
		return binanceCallApiWithSecret[SpotTicker24hrRes](api.client.c, url, GET)
	}
}

// binance SPOT spotTickerBookTicker rest当前最优挂单 (NONE)
func (client *SpotRestClient) NewSpotTickerBookTicker() *SpotTickerBookTickerApi {
	return &SpotTickerBookTickerApi{
		client: client,
		req:    &SpotTickerBookTickerReq{},
	}
}
func (api *SpotTickerBookTickerApi) Do() (*SpotTickerBookTickerRes, error) {
	url := binanceHandlerRequestApi(SPOT, api.req, SpotApiMap[SpotTickerBookTicker])
	if api.req.Symbol != nil && *api.req.Symbol != "" {
		res, err := binanceCallApiWithSecret[SpotTickerBookTickerResRow](api.client.c, url, GET)
		if err != nil {
			return nil, err
		}
		return &SpotTickerBookTickerRes{*res}, nil
	} else {
		return binanceCallApiWithSecret[SpotTickerBookTickerRes](api.client.c, url, GET)
	}
}

// binance SPOT spotTicker rest滚动窗口价格变动统计
func (client *SpotRestClient) NewSpotTicker() *SpotTickerApi {
	return &SpotTickerApi{
		client: client,
		req:    &SpotTickerReq{},
	}
}
func (api *SpotTickerApi) Do() (*SpotTickerRes, error) {
	url := binanceHandlerRequestApi(SPOT, api.req, SpotApiMap[SpotTicker])
	if api.req.Symbol != nil && *api.req.Symbol != "" {
		res, err := binanceCallApiWithSecret[SpotTickerResRow](api.client.c, url, GET)
		if err != nil {
			return nil, err
		}
		return &SpotTickerRes{*res}, nil
	} else {
		return binanceCallApiWithSecret[SpotTickerRes](api.client.c, url, GET)
	}
}
