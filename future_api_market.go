package mybinanceapi

// 行情接口
// binance FUTURE FutureKlines restK线数据 (MARKET_DATA)
func (client *FutureRestClient) NewFutureKlines() *FutureKlinesApi {
	return &FutureKlinesApi{
		client: client,
		req:    &FutureKlinesReq{},
	}
}
func (api *FutureKlinesApi) Do() (*KlinesRes, error) {
	url := binanceHandlerRequestApi(FUTURE, api.req, FutureApiMap[FutureKlines])
	res, err := binanceCallApiWithSecret[KlinesMiddle](api.client.c, url, GET)
	if err != nil {
		return nil, err
	}
	res2 := res.ConvertToRes()
	return &res2, nil
}

// binance FUTURE FutureDepth rest深度信息 (MARKET_DATA)
func (client *FutureRestClient) NewFutureDepth() *FutureDepthApi {
	return &FutureDepthApi{
		client: client,
		req:    &FutureDepthReq{},
	}
}
func (api *FutureDepthApi) Do() (*FutureDepthRes, error) {
	url := binanceHandlerRequestApi(FUTURE, api.req, FutureApiMap[FutureDepth])
	res, err := binanceCallApiWithSecret[FutureDepthResMiddle](api.client.c, url, GET)
	if err != nil {
		return nil, err
	}
	return res.ConvertToRes(), nil
}

// binance FUTURE FutureTrades rest最新成交 (MARKET_DATA)
func (client *FutureRestClient) NewFutureTrades() *FutureTradesApi {
	return &FutureTradesApi{
		client: client,
		req:    &FutureTradesReq{},
	}
}
func (api *FutureTradesApi) Do() (*FutureTradesRes, error) {
	url := binanceHandlerRequestApi(FUTURE, api.req, FutureApiMap[FutureTrades])
	return binanceCallApiWithSecret[FutureTradesRes](api.client.c, url, GET)
}

// binance FUTURE FutureHistoricalTrades rest历史成交 (MARKET_DATA)
func (client *FutureRestClient) NewFutureHistoricalTrades() *FutureHistoricalTradesApi {
	return &FutureHistoricalTradesApi{
		client: client,
		req:    &FutureHistoricalTradesReq{},
	}
}
func (api *FutureHistoricalTradesApi) Do() (*FutureHistoricalTradesRes, error) {
	url := binanceHandlerRequestApi(FUTURE, api.req, FutureApiMap[FutureHistoricalTrades])
	return binanceCallApiWithSecret[FutureHistoricalTradesRes](api.client.c, url, GET)
}

// binance FUTURE FutureAggTrades rest近期成交(归集) (MARKET_DATA)
func (client *FutureRestClient) NewFutureAggTrades() *FutureAggTradesApi {
	return &FutureAggTradesApi{
		client: client,
		req:    &FutureAggTradesReq{},
	}
}
func (api *FutureAggTradesApi) Do() (*FutureAggTradesRes, error) {
	url := binanceHandlerRequestApi(FUTURE, api.req, FutureApiMap[FutureAggTrades])
	return binanceCallApiWithSecret[FutureAggTradesRes](api.client.c, url, GET)
}

// binance FUTURE FuturePremiumIndex rest最新标记价格和资金费率 (MARKET_DATA)
func (client *FutureRestClient) NewFuturePremiumIndex() *FuturePremiumIndexApi {
	return &FuturePremiumIndexApi{
		client: client,
		req:    &FuturePremiumIndexReq{},
	}
}
func (api *FuturePremiumIndexApi) Do() (*FuturePremiumIndexRes, error) {
	url := binanceHandlerRequestApi(FUTURE, api.req, FutureApiMap[FuturePremiumIndex])
	if api.req.Symbol != nil && *api.req.Symbol != "" {
		res, err := binanceCallApiWithSecret[FuturePremiumIndexResRow](api.client.c, url, GET)
		if err != nil {
			return nil, err
		}
		return &FuturePremiumIndexRes{*res}, nil
	} else {
		return binanceCallApiWithSecret[FuturePremiumIndexRes](api.client.c, url, GET)
	}

}

// binance FUTURE FutureFundingRate rest查询资金费率历史 (MARKET_DATA)
func (client *FutureRestClient) NewFutureFundingRate() *FutureFundingRateApi {
	return &FutureFundingRateApi{
		client: client,
		req:    &FutureFundingRateReq{},
	}
}
func (api *FutureFundingRateApi) Do() (*FutureFundingRateRes, error) {
	url := binanceHandlerRequestApi(FUTURE, api.req, FutureApiMap[FutureFundingRate])
	return binanceCallApiWithSecret[FutureFundingRateRes](api.client.c, url, GET)
}

// binance FUTURE FutureFundingInfo rest查询资金费率信息 (MARKET_DATA)
func (client *FutureRestClient) NewFutureFundingInfo() *FutureFundingInfoApi {
	return &FutureFundingInfoApi{
		client: client,
		req:    &FutureFundingInfoReq{},
	}
}
func (api *FutureFundingInfoApi) Do() (*FutureFundingInfoRes, error) {
	url := binanceHandlerRequestApi(FUTURE, api.req, FutureApiMap[FutureFundingInfo])
	return binanceCallApiWithSecret[FutureFundingInfoRes](api.client.c, url, GET)
}

// binance FUTURE FutureTicker24hr rest24hr价格变动情况 (MARKET_DATA)
func (client *FutureRestClient) NewFutureTicker24hr() *FutureTicker24hrApi {
	return &FutureTicker24hrApi{
		client: client,
		req:    &FutureTicker24hrReq{},
	}
}
func (api *FutureTicker24hrApi) Do() (*FutureTicker24hrRes, error) {
	url := binanceHandlerRequestApi(FUTURE, api.req, FutureApiMap[FutureTicker24hr])
	if api.req.Symbol != nil && *api.req.Symbol != "" {
		res, err := binanceCallApiWithSecret[FutureTicker24hrResRow](api.client.c, url, GET)
		if err != nil {
			return nil, err
		}
		return &FutureTicker24hrRes{*res}, nil
	} else {
		return binanceCallApiWithSecret[FutureTicker24hrRes](api.client.c, url, GET)
	}
}

// binance FUTURE FutureTickerPrice rest最新价格 (MARKET_DATA)
func (client *FutureRestClient) NewFutureTickerPrice() *FutureTickerPriceApi {
	return &FutureTickerPriceApi{
		client: client,
		req:    &FutureTickerPriceReq{},
	}
}
func (api *FutureTickerPriceApi) Do() (*FutureTickerPriceRes, error) {
	url := binanceHandlerRequestApi(FUTURE, api.req, FutureApiMap[FutureTickerPrice])
	if api.req.Symbol != nil && *api.req.Symbol != "" {
		res, err := binanceCallApiWithSecret[FutureTickerPriceResRow](api.client.c, url, GET)
		if err != nil {
			return nil, err
		}
		return &FutureTickerPriceRes{*res}, nil
	} else {
		return binanceCallApiWithSecret[FutureTickerPriceRes](api.client.c, url, GET)
	}
}

// binance FUTURE FutureTickerBookTicker rest当前最优挂单 (MARKET_DATA)
func (client *FutureRestClient) NewFutureTickerBookTicker() *FutureTickerBookTickerApi {
	return &FutureTickerBookTickerApi{
		client: client,
		req:    &FutureTickerBookTickerReq{},
	}
}
func (api *FutureTickerBookTickerApi) Do() (*FutureTickerBookTickerRes, error) {
	url := binanceHandlerRequestApi(FUTURE, api.req, FutureApiMap[FutureTickerBookTicker])
	if api.req.Symbol != nil && *api.req.Symbol != "" {
		res, err := binanceCallApiWithSecret[FutureTickerBookTickerResRow](api.client.c, url, GET)
		if err != nil {
			return nil, err
		}
		return &FutureTickerBookTickerRes{*res}, nil
	} else {
		return binanceCallApiWithSecret[FutureTickerBookTickerRes](api.client.c, url, GET)
	}
}

// binance FUTURE FutureDataBasis rest基差数据 (MARKET_DATA)
func (client *FutureRestClient) NewFutureDataBasis() *FutureDataBasisApi {
	return &FutureDataBasisApi{
		client: client,
		req:    &FutureDataBasisReq{},
	}
}
func (api *FutureDataBasisApi) Do() (*FutureDataBasisRes, error) {
	url := binanceHandlerRequestApi(FUTURE, api.req, FutureApiMap[FutureDataBasis])
	return binanceCallApiWithSecret[FutureDataBasisRes](api.client.c, url, GET)
}
