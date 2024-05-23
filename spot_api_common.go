package mybinanceapi

// 通用接口
// binance SPOT ping rest测试连通性 (NONE)
func (client *SpotRestClient) NewPing() *SpotPingApi {
	return &SpotPingApi{
		client: client,
		req:    &SpotPingReq{},
	}
}
func (api *SpotPingApi) Do() (*SpotPingRes, error) {
	url := binanceHandlerRequestApi(SPOT, api.req, SpotApiMap[SpotPing])
	return binanceCallApiWithSecret[SpotPingRes](api.client.c, url, GET)
}

// binance SPOT serverTime rest服务器时间 (NONE)
func (client *SpotRestClient) NewServerTime() *SpotServerTimeApi {
	return &SpotServerTimeApi{
		client: client,
		req:    &SpotServerTimeReq{},
	}
}
func (api *SpotServerTimeApi) Do() (*SpotServerTimeRes, error) {
	url := binanceHandlerRequestApi(SPOT, api.req, SpotApiMap[SpotServerTime])
	return binanceCallApiWithSecret[SpotServerTimeRes](api.client.c, url, GET)
}

// binance SPOT exchangeInfo rest交易规范
func (client *SpotRestClient) NewExchangeInfo() *SpotExchangeInfoApi {
	return &SpotExchangeInfoApi{
		client: client,
		req:    &SpotExchangeInfoReq{},
	}
}
func (api *SpotExchangeInfoApi) Do() (*SpotExchangeInfoRes, error) {
	url := binanceHandlerRequestApi(SPOT, api.req, SpotApiMap[SpotExchangeInfo])
	return binanceCallApiWithSecret[SpotExchangeInfoRes](api.client.c, url, GET)
}
