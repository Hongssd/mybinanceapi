package mybinanceapi

// Ws账户推送相关
// binance FUTURE FutureListenKeyPost rest生成listenKey (USER_STREAM)
func (client *FutureRestClient) NewFutureListenKeyPost() *FutureListenKeyPostApi {
	return &FutureListenKeyPostApi{
		client: client,
		req:    &FutureListenKeyPostReq{},
	}
}
func (api *FutureListenKeyPostApi) Do() (*FutureListenKeyPostRes, error) {
	url := binanceHandlerRequestApiWithSecret(FUTURE, api.req, FutureApiMap[FutureListenKeyPost], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[FutureListenKeyPostRes](api.client.c, url, POST)
}

// binance FUTURE FutureListenKeyPut rest延长listenKey有效期 (USER_STREAM)
func (client *FutureRestClient) NewFutureListenKeyPut() *FutureListenKeyPutApi {
	return &FutureListenKeyPutApi{
		client: client,
		req:    &FutureListenKeyPutReq{},
	}
}
func (api *FutureListenKeyPutApi) Do() (*FutureListenKeyPutRes, error) {
	url := binanceHandlerRequestApiWithSecret(FUTURE, api.req, FutureApiMap[FutureListenKeyPut], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[FutureListenKeyPutRes](api.client.c, url, PUT)
}

// binance FUTURE FutureListenKeyDelete rest关闭listenKey (USER_STREAM)
func (client *FutureRestClient) NewFutureListenKeyDelete() *FutureListenKeyDeleteApi {
	return &FutureListenKeyDeleteApi{
		client: client,
		req:    &FutureListenKeyDeleteReq{},
	}
}
func (api *FutureListenKeyDeleteApi) Do() (*FutureListenKeyDeleteRes, error) {
	url := binanceHandlerRequestApiWithSecret(FUTURE, api.req, FutureApiMap[FutureListenKeyDelete], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[FutureListenKeyDeleteRes](api.client.c, url, DELETE)
}
