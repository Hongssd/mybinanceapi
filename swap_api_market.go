package mybinanceapi

// 行情接口
// binance SWAP  SwapKlines rest获取K线数据
func (client *SwapRestClient) NewSwapKlines() *SwapKlinesApi {
	return &SwapKlinesApi{
		client: client,
		req:    &SwapKlinesReq{},
	}
}
func (api *SwapKlinesApi) Do() (*KlinesRes, error) {
	url := binanceHandlerRequestApi(SWAP, api.req, SwapApiMap[SwapKlines])
	res, err := binanceCallApiWithSecret[KlinesMiddle](api.client.c, url, GET)
	if err != nil {
		return nil, err
	}
	res2 := res.ConvertToRes()
	return &res2, nil
}

// binance SWAP  SwapTickerPrice rest获取交易对最新价格
func (client *SwapRestClient) NewSwapTickerPrice() *SwapTickerPriceApi {
	return &SwapTickerPriceApi{
		client: client,
		req:    &SwapTickerPriceReq{},
	}
}
func (api *SwapTickerPriceApi) Do() (*SwapTickerPriceRes, error) {
	url := binanceHandlerRequestApi(SWAP, api.req, SwapApiMap[SwapTickerPrice])
	return binanceCallApiWithSecret[SwapTickerPriceRes](api.client.c, url, GET)
}

// binance SWAP  SwapDepth rest获取深度信息
func (client *SwapRestClient) NewSwapDepth() *SwapDepthApi {
	return &SwapDepthApi{
		client: client,
		req:    &SwapDepthReq{},
	}
}
func (api *SwapDepthApi) Do() (*SwapDepthRes, error) {
	url := binanceHandlerRequestApi(SWAP, api.req, SwapApiMap[SwapDepth])
	res, err := binanceCallApiWithSecret[SwapDepthResMiddle](api.client.c, url, GET)
	if err != nil {
		return nil, err
	}
	return res.ConvertToRes(), nil
}
