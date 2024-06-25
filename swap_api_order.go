package mybinanceapi

import "time"

// 币合约订单接口
// binance SWAP  SwapOpenOrders rest查询当前挂单 (USER_DATA)
func (client *SwapRestClient) NewOpenOrders() *SwapOpenOrdersApi {
	return &SwapOpenOrdersApi{
		client: client,
		req:    &SwapOpenOrdersReq{},
	}
}
func (api *SwapOpenOrdersApi) Do() (*SwapOpenOrdersRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SWAP, api.req, SwapApiMap[SwapOpenOrders], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SwapOpenOrdersRes](api.client.c, url, GET)
}

// binance SWAP  SwapAllOrders rest查询所有订单 (USER_DATA)
func (client *SwapRestClient) NewAllOrders() *SwapAllOrdersApi {
	return &SwapAllOrdersApi{
		client: client,
		req:    &SwapAllOrdersReq{},
	}
}
func (api *SwapAllOrdersApi) Do() (*SwapAllOrdersRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SWAP, api.req, SwapApiMap[SwapAllOrders], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SwapAllOrdersRes](api.client.c, url, GET)
}

// binance SWAP  SwapOrderPost rest下单 (TRADE)
func (client *SwapRestClient) NewSwapOrderPost() *SwapOrderPostApi {
	return &SwapOrderPostApi{
		client: client,
		req:    &SwapOrderPostReq{},
	}
}
func (api *SwapOrderPostApi) Do() (*SwapOrderPostRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SWAP, api.req, SwapApiMap[SwapOrderPost], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SwapOrderPostRes](api.client.c, url, POST)
}

// binance SWAP  SwapOrderPut rest修改订单 (TRADE)
func (client *SwapRestClient) NewSwapOrderPut() *SwapOrderPutApi {
	return &SwapOrderPutApi{
		client: client,
		req:    &SwapOrderPutReq{},
	}
}
func (api *SwapOrderPutApi) Do() (*SwapOrderPutRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SWAP, api.req, SwapApiMap[SwapOrderPut], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SwapOrderPutRes](api.client.c, url, PUT)
}

// binance SWAP  SwapOrderGet	rest查询订单 (USER_DATA)
func (client *SwapRestClient) NewSwapOrderGet() *SwapOrderGetApi {
	return &SwapOrderGetApi{
		client: client,
		req:    &SwapOrderGetReq{},
	}
}
func (api *SwapOrderGetApi) Do() (*SwapOrderGetRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SWAP, api.req, SwapApiMap[SwapOrderGet], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SwapOrderGetRes](api.client.c, url, GET)
}

// binance SWAP  SwapOrderDelete rest撤销订单 (TRADE)
func (client *SwapRestClient) NewSwapOrderDelete() *SwapOrderDeleteApi {
	return &SwapOrderDeleteApi{
		client: client,
		req:    &SwapOrderDeleteReq{},
	}
}
func (api *SwapOrderDeleteApi) Do() (*SwapOrderDeleteRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SWAP, api.req, SwapApiMap[SwapOrderDelete], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SwapOrderDeleteRes](api.client.c, url, DELETE)
}

// binance SWAP  SwapUserTrades rest查询账户成交记录 (USER_DATA)
func (client *SwapRestClient) NewSwapUserTrades() *SwapUserTradesApi {
	return &SwapUserTradesApi{
		client: client,
		req:    &SwapUserTradesReq{},
	}
}
func (api *SwapUserTradesApi) Do() (*SwapUserTradesRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SWAP, api.req, SwapApiMap[SwapUserTrades], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SwapUserTradesRes](api.client.c, url, GET)
}

// binance SWAP  SwapBatchOrdersPost rest批量下单 (TRADE)
func (client *SwapRestClient) NewSwapBatchOrdersPost() *SwapBatchOrdersPostApi {
	return &SwapBatchOrdersPostApi{
		client: client,
		req:    &SwapBatchOrdersPostReq{},
	}
}
func (api *SwapBatchOrdersPostApi) Do() (*SwapBatchOrdersPostRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SWAP, api.req, SwapApiMap[SwapBatchOrdersPost], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SwapBatchOrdersPostRes](api.client.c, url, POST)
}

// binance SWAP  SwapBatchOrdersPut rest批量修改订单 (TRADE)
func (client *SwapRestClient) NewSwapBatchOrdersPut() *SwapBatchOrdersPutApi {
	return &SwapBatchOrdersPutApi{
		client: client,
		req:    &SwapBatchOrdersPutReq{},
	}
}
func (api *SwapBatchOrdersPutApi) Do() (*SwapBatchOrdersPutRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SWAP, api.req, SwapApiMap[SwapBatchOrdersPut], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SwapBatchOrdersPutRes](api.client.c, url, PUT)
}

// binance SWAP  SwapBatchOrdersDelete rest批量撤销订单 (TRADE)
func (client *SwapRestClient) NewSwapBatchOrdersDelete() *SwapBatchOrdersDeleteApi {
	return &SwapBatchOrdersDeleteApi{
		client: client,
		req:    &SwapBatchOrdersDeleteReq{},
	}
}
func (api *SwapBatchOrdersDeleteApi) Do() (*SwapBatchOrdersDeleteRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SWAP, api.req, SwapApiMap[SwapBatchOrdersDelete], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SwapBatchOrdersDeleteRes](api.client.c, url, DELETE)
}

// binance SWAP  SwapCommissionRate rest查询用户当前的手续费率
func (client *SwapRestClient) NewSwapCommissionRate() *SwapCommissionRateApi {
	return &SwapCommissionRateApi{
		client: client,
		req:    &SwapCommissionRateReq{},
	}
}
func (api *SwapCommissionRateApi) Do() (*SwapCommissionRateRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SWAP, api.req, SwapApiMap[SwapCommissionRate], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SwapCommissionRateRes](api.client.c, url, GET)
}
