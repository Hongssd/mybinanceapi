package mybinanceapi

// Ws账户推送相关
// binance SWAP  SwapListenKeyPost rest创建listenKey (USER_STREAM)
func (client *SwapRestClient) NewSwapListenKeyPost() *SwapListenKeyPostApi {
	return &SwapListenKeyPostApi{
		client: client,
		req:    &SwapListenKeyPostReq{},
	}
}
func (api *SwapListenKeyPostApi) Do() (*SwapListenKeyPostRes, error) {
	url := binanceHandlerRequestApiWithSecret(SWAP, api.req, SwapApiMap[SwapListenKeyPost], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SwapListenKeyPostRes](api.client.c, url, POST)
}

// binance SWAP  SwapListenKeyPut rest延长listenKey有效期 (USER_STREAM)
func (client *SwapRestClient) NewSwapListenKeyPut() *SwapListenKeyPutApi {
	return &SwapListenKeyPutApi{
		client: client,
		req:    &SwapListenKeyPutReq{},
	}
}
func (api *SwapListenKeyPutApi) Do() (*SwapListenKeyPutRes, error) {
	url := binanceHandlerRequestApiWithSecret(SWAP, api.req, SwapApiMap[SwapListenKeyPut], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SwapListenKeyPutRes](api.client.c, url, PUT)
}

// binance SWAP  SwapListenKeyDelete rest关闭listenKey (USER_STREAM)
func (client *SwapRestClient) NewSwapListenKeyDelete() *SwapListenKeyDeleteApi {
	return &SwapListenKeyDeleteApi{
		client: client,
		req:    &SwapListenKeyDeleteReq{},
	}
}
func (api *SwapListenKeyDeleteApi) Do() (*SwapListenKeyDeleteRes, error) {
	url := binanceHandlerRequestApiWithSecret(SWAP, api.req, SwapApiMap[SwapListenKeyDel], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SwapListenKeyDeleteRes](api.client.c, url, DELETE)
}
