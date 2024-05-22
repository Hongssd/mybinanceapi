package mybinanceapi

import "time"

// 交易接口
// binance FUTURE FutureOpenOrders rest查询当前挂单 (USER_DATA)
func (client *FutureRestClient) NewOpenOrders() *FutureOpenOrdersApi {
	return &FutureOpenOrdersApi{
		client: client,
		req:    &FutureOpenOrdersReq{},
	}
}
func (api *FutureOpenOrdersApi) Do() (*FutureOpenOrdersRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(FUTURE, api.req, FutureApiMap[FutureOpenOrders], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[FutureOpenOrdersRes](api.client.c, url, GET)
}

// binance FUTURE FutureAllOrders rest查询所有订单 (USER_DATA)
func (client *FutureRestClient) NewAllOrders() *FutureAllOrdersApi {
	return &FutureAllOrdersApi{
		client: client,
		req:    &FutureAllOrdersReq{},
	}
}
func (api *FutureAllOrdersApi) Do() (*FutureAllOrdersRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(FUTURE, api.req, FutureApiMap[FutureAllOrders], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[FutureAllOrdersRes](api.client.c, url, GET)
}

// binance FUTURE FutureOrderPost rest下单 (TRADE)
func (client *FutureRestClient) NewFutureOrderPost() *FutureOrderPostApi {
	return &FutureOrderPostApi{
		client: client,
		req:    &FutureOrderPostReq{},
	}
}
func (api *FutureOrderPostApi) Do() (*FutureOrderPostRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(FUTURE, api.req, FutureApiMap[FutureOrderPost], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[FutureOrderPostRes](api.client.c, url, POST)
}

// binance FUTURE FutureOrderPut rest修改订单 (TRADE)
func (client *FutureRestClient) NewFutureOrderPut() *FutureOrderPutApi {
	return &FutureOrderPutApi{
		client: client,
		req:    &FutureOrderPutReq{},
	}
}
func (api *FutureOrderPutApi) Do() (*FutureOrderPutRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(FUTURE, api.req, FutureApiMap[FutureOrderPut], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[FutureOrderPutRes](api.client.c, url, PUT)
}

// binance FUTURE FutureOrderGet  rest查询订单 (USER_DATA)
func (client *FutureRestClient) NewFutureOrderGet() *FutureOrderGetApi {
	return &FutureOrderGetApi{
		client: client,
		req:    &FutureOrderGetReq{},
	}
}
func (api *FutureOrderGetApi) Do() (*FutureOrderGetRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(FUTURE, api.req, FutureApiMap[FutureOrderGet], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[FutureOrderGetRes](api.client.c, url, GET)
}

// binance FUTURE FutureOrderDelete rest撤销订单 (TRADE)
func (client *FutureRestClient) NewFutureOrderDelete() *FutureOrderDeleteApi {
	return &FutureOrderDeleteApi{
		client: client,
		req:    &FutureOrderDeleteReq{},
	}
}
func (api *FutureOrderDeleteApi) Do() (*FutureOrderDeleteRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(FUTURE, api.req, FutureApiMap[FutureOrderDelete], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[FutureOrderDeleteRes](api.client.c, url, DELETE)
}

// binance FUTURE FutureBatchOrdersPost rest批量下单 (TRADE)
func (client *FutureRestClient) NewFutureBatchOrdersPost() *FutureBatchOrdersPostApi {
	return &FutureBatchOrdersPostApi{
		client: client,
		req:    &FutureBatchOrdersPostReq{},
	}
}
func (api *FutureBatchOrdersPostApi) Do() (*FutureBatchOrdersPostRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(FUTURE, api.req, FutureApiMap[FutureBatchOrdersPost], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[FutureBatchOrdersPostRes](api.client.c, url, POST)
}

// binance FUTURE FutureBatchOrdersPut rest批量修改订单 (TRADE)
func (client *FutureRestClient) NewFutureBatchOrdersPut() *FutureBatchOrdersPutApi {
	return &FutureBatchOrdersPutApi{
		client: client,
		req:    &FutureBatchOrdersPutReq{},
	}
}
func (api *FutureBatchOrdersPutApi) Do() (*FutureBatchOrdersPutRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(FUTURE, api.req, FutureApiMap[FutureBatchOrdersPut], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[FutureBatchOrdersPutRes](api.client.c, url, PUT)
}

// binance FUTURE FutureBatchOrdersDelete rest批量撤销订单 (TRADE)
func (client *FutureRestClient) NewFutureBatchOrdersDelete() *FutureBatchOrdersDeleteApi {
	return &FutureBatchOrdersDeleteApi{
		client: client,
		req:    &FutureBatchOrdersDeleteReq{},
	}
}
func (api *FutureBatchOrdersDeleteApi) Do() (*FutureBatchOrdersDeleteRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(FUTURE, api.req, FutureApiMap[FutureBatchOrdersDelete], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[FutureBatchOrdersDeleteRes](api.client.c, url, DELETE)
}

// binance FUTURE FutureUserTrades rest账户成交历史 (USER_DATA)
func (client *FutureRestClient) NewFutureUserTrades() *FutureUserTradesApi {
	return &FutureUserTradesApi{
		client: client,
		req:    &FutureUserTradesReq{},
	}
}
func (api *FutureUserTradesApi) Do() (*FutureUserTradesRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(FUTURE, api.req, FutureApiMap[FutureUserTrades], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[FutureUserTradesRes](api.client.c, url, GET)
}

// binance FUTURE  FutureCommissionRate rest查询用户当前的手续费率 (USER_DATA)
func (client *FutureRestClient) NewFutureCommissionRate() *FutureCommissionRateApi {
	return &FutureCommissionRateApi{
		client: client,
		req:    &FutureCommissionRateReq{},
	}
}
func (api *FutureCommissionRateApi) Do() (*FutureCommissionRateRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(FUTURE, api.req, FutureApiMap[FutureCommissionRate], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[FutureCommissionRateRes](api.client.c, url, GET)
}
