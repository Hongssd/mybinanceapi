package mybinanceapi

// 通用接口
// binance SWAP  SwapPing rest测试连接
func (client *SwapRestClient) NewPing() *SwapPingApi {
	return &SwapPingApi{
		client: client,
		req:    &SwapPingReq{},
	}
}
func (api *SwapPingApi) Do() (*SwapPingRes, error) {
	url := binanceHandlerRequestApi(SWAP, api.req, SwapApiMap[SwapPing])
	return binanceCallApiWithSecret[SwapPingRes](api.client.c, url, GET)
}

// binance SWAP  SwapServerTime rest获取服务器时间
func (client *SwapRestClient) NewServerTime() *SwapServerTimeApi {
	return &SwapServerTimeApi{
		client: client,
		req:    &SwapServerTimeReq{},
	}
}
func (api *SwapServerTimeApi) Do() (*SwapServerTimeRes, error) {
	url := binanceHandlerRequestApi(SWAP, api.req, SwapApiMap[SwapServerTime])
	return binanceCallApiWithSecret[SwapServerTimeRes](api.client.c, url, GET)
}

// binance SWAP  SwapExchangeInfo rest交易规则和交易对信息
func (client *SwapRestClient) NewExchangeInfo() *SwapExchangeInfoApi {
	return &SwapExchangeInfoApi{
		client: client,
		req:    &SwapExchangeInfoReq{},
	}
}
func (api *SwapExchangeInfoApi) Do() (*SwapExchangeInfoRes, error) {
	url := binanceHandlerRequestApi(SWAP, api.req, SwapApiMap[SwapExchangeInfo])
	return binanceCallApiWithSecret[SwapExchangeInfoRes](api.client.c, url, GET)
}
