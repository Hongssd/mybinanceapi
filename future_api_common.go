package mybinanceapi

// 通用接口
// binance FUTURE FuturePing rest测试服务器连通性 (NONE)
func (client *FutureRestClient) NewPing() *FuturePingApi {
	return &FuturePingApi{
		client: client,
		req:    &FuturePingReq{},
	}
}
func (api *FuturePingApi) Do() (*FuturePingRes, error) {
	url := binanceHandlerRequestApi(FUTURE, api.req, FutureApiMap[FuturePing])
	return binanceCallApiWithSecret[FuturePingRes](api.client.c, url, GET)
}

// binance FUTURE FutureTime rest获取服务器时间 (NONE)
func (client *FutureRestClient) NewServerTime() *FutureServerTimeApi {
	return &FutureServerTimeApi{
		client: client,
		req:    &FutureServerTimeReq{},
	}
}
func (api *FutureServerTimeApi) Do() (*FutureTimeRes, error) {
	url := binanceHandlerRequestApi(FUTURE, api.req, FutureApiMap[FutureServerTime])
	return binanceCallApiWithSecret[FutureTimeRes](api.client.c, url, GET)
}

// binance FUTURE FutureExchangeInfo rest交易规范信息和交易对信息 (NONE)
func (client *FutureRestClient) NewExchangeInfo() *FutureExchangeInfoApi {
	return &FutureExchangeInfoApi{
		client: client,
		req:    &FutureExchangeInfoReq{},
	}
}
func (api *FutureExchangeInfoApi) Do() (*FutureExchangeInfoRes, error) {
	url := binanceHandlerRequestApi(FUTURE, api.req, FutureApiMap[FutureExchangeInfo])
	return binanceCallApiWithSecret[FutureExchangeInfoRes](api.client.c, url, GET)
}
