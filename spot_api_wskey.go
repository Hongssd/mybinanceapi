package mybinanceapi

// Ws账户推送相关

// binance SPOT ws账户推送  SpotUserDataStreamPost rest现货创建一个listenKey (USER_STREAM)
func (client *SpotRestClient) NewSpotUserDataStreamPost() *SpotUserDataStreamPostApi {
	return &SpotUserDataStreamPostApi{
		client: client,
		req:    &SpotUserDataStreamPostReq{},
	}
}
func (api *SpotUserDataStreamPostApi) Do() (*SpotUserDataStreamPostRes, error) {
	url := binanceHandlerRequestApi(SPOT, api.req, SpotApiMap[SpotUserDataStreamPost])
	return binanceCallApiWithSecret[SpotUserDataStreamPostRes](api.client.c, url, POST)
}

// binance SPOT ws账户推送  SpotUserDataStreamPut rest现货延长listenKey有效期 (USER_STREAM)
func (client *SpotRestClient) NewSpotUserDataStreamPut() *SpotUserDataStreamPutApi {
	return &SpotUserDataStreamPutApi{
		client: client,
		req:    &SpotUserDataStreamPutReq{},
	}
}
func (api *SpotUserDataStreamPutApi) Do() (*SpotUserDataStreamPutRes, error) {
	url := binanceHandlerRequestApi(SPOT, api.req, SpotApiMap[SpotUserDataStreamPut])
	return binanceCallApiWithSecret[SpotUserDataStreamPutRes](api.client.c, url, PUT)
}

// binance SPOT ws账户推送  SpotUserDataStreamDelete rest现货关闭listenKey (USER_STREAM)
func (client *SpotRestClient) NewSpotUserDataStreamDelete() *SpotUserDataStreamDeleteApi {
	return &SpotUserDataStreamDeleteApi{
		client: client,
		req:    &SpotUserDataStreamDeleteReq{},
	}
}
func (api *SpotUserDataStreamDeleteApi) Do() (*SpotUserDataStreamDeleteRes, error) {
	url := binanceHandlerRequestApi(SPOT, api.req, SpotApiMap[SpotUserDataStreamDelete])
	return binanceCallApiWithSecret[SpotUserDataStreamDeleteRes](api.client.c, url, DELETE)
}

// binance SPOT ws账户推送 SpotMarginUserDataStreamPost rest现货杠杆创建一个listenKey (USER_STREAM)
func (client *SpotRestClient) NewSpotMarginUserDataStreamPost() *SpotMarginUserDataStreamPostApi {
	return &SpotMarginUserDataStreamPostApi{
		client: client,
		req:    &SpotMarginUserDataStreamPostReq{},
	}
}
func (api *SpotMarginUserDataStreamPostApi) Do() (*SpotMarginUserDataStreamPostRes, error) {
	url := binanceHandlerRequestApi(SPOT, api.req, SpotApiMap[SpotMarginUserDataStreamPost])
	return binanceCallApiWithSecret[SpotMarginUserDataStreamPostRes](api.client.c, url, POST)
}

// binance SPOT ws账户推送  SpotMarginUserDataStreamPut rest现货杠杆延长listenKey有效期 (USER_STREAM)
func (client *SpotRestClient) NewSpotMarginUserDataStreamPut() *SpotMarginUserDataStreamPutApi {
	return &SpotMarginUserDataStreamPutApi{
		client: client,
		req:    &SpotMarginUserDataStreamPutReq{},
	}
}
func (api *SpotMarginUserDataStreamPutApi) Do() (*SpotMarginUserDataStreamPutRes, error) {
	url := binanceHandlerRequestApi(SPOT, api.req, SpotApiMap[SpotMarginUserDataStreamPut])
	return binanceCallApiWithSecret[SpotMarginUserDataStreamPutRes](api.client.c, url, PUT)
}

// binance SPOT ws账户推送  SpotMarginUserDataStreamDelete rest现货杠杆关闭listenKey (USER_STREAM)
func (client *SpotRestClient) NewSpotMarginUserDataStreamDelete() *SpotMarginUserDataStreamDeleteApi {
	return &SpotMarginUserDataStreamDeleteApi{
		client: client,
		req:    &SpotMarginUserDataStreamDeleteReq{},
	}
}
func (api *SpotMarginUserDataStreamDeleteApi) Do() (*SpotMarginUserDataStreamDeleteRes, error) {
	url := binanceHandlerRequestApi(SPOT, api.req, SpotApiMap[SpotMarginUserDataStreamDelete])
	return binanceCallApiWithSecret[SpotMarginUserDataStreamDeleteRes](api.client.c, url, DELETE)
}

// binance SPOT ws账户推送  SpotMarginIsolatedUserDataStreamPost rest现货杠杆逐仓创建一个listenKey (USER_STREAM)
func (client *SpotRestClient) NewSpotMarginIsolatedUserDataStreamPost() *SpotMarginIsolatedUserDataStreamPostApi {
	return &SpotMarginIsolatedUserDataStreamPostApi{
		client: client,
		req:    &SpotMarginIsolatedUserDataStreamPostReq{},
	}
}
func (api *SpotMarginIsolatedUserDataStreamPostApi) Do() (*SpotMarginIsolatedUserDataStreamPostRes, error) {
	url := binanceHandlerRequestApi(SPOT, api.req, SpotApiMap[SpotMarginIsolatedUserDataStreamPost])
	return binanceCallApiWithSecret[SpotMarginIsolatedUserDataStreamPostRes](api.client.c, url, POST)
}

// binance SPOT ws账户推送  SpotMarginIsolatedUserDataStreamPut rest现货杠杆逐仓延长listenKey有效期 (USER_STREAM)
func (client *SpotRestClient) NewSpotMarginIsolatedUserDataStreamPut() *SpotMarginIsolatedUserDataStreamPutApi {
	return &SpotMarginIsolatedUserDataStreamPutApi{
		client: client,
		req:    &SpotMarginIsolatedUserDataStreamPutReq{},
	}
}
func (api *SpotMarginIsolatedUserDataStreamPutApi) Do() (*SpotMarginIsolatedUserDataStreamPutRes, error) {
	url := binanceHandlerRequestApi(SPOT, api.req, SpotApiMap[SpotMarginIsolatedUserDataStreamPut])
	return binanceCallApiWithSecret[SpotMarginIsolatedUserDataStreamPutRes](api.client.c, url, PUT)
}

// binance SPOT ws账户推送  SpotMarginIsolatedUserDataStreamDelete rest现货杠杆逐仓关闭listenKey (USER_STREAM)
func (client *SpotRestClient) NewSpotMarginIsolatedUserDataStreamDelete() *SpotMarginIsolatedUserDataStreamDeleteApi {
	return &SpotMarginIsolatedUserDataStreamDeleteApi{
		client: client,
		req:    &SpotMarginIsolatedUserDataStreamDeleteReq{},
	}
}
func (api *SpotMarginIsolatedUserDataStreamDeleteApi) Do() (*SpotMarginIsolatedUserDataStreamDeleteRes, error) {
	url := binanceHandlerRequestApi(SPOT, api.req, SpotApiMap[SpotMarginIsolatedUserDataStreamDelete])
	return binanceCallApiWithSecret[SpotMarginIsolatedUserDataStreamDeleteRes](api.client.c, url, DELETE)
}
