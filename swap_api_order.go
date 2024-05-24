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
