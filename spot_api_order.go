package mybinanceapi

import "time"

// 现货订单接口
// binance SPOT openOrders rest当前挂单 (USER_DATA)
func (client *SpotRestClient) NewOpenOrders() *SpotOpenOrdersApi {
	return &SpotOpenOrdersApi{
		client: client,
		req:    &SpotOpenOrdersReq{},
	}
}
func (api *SpotOpenOrdersApi) Do() (*SpotOpenOrdersRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotOpenOrders], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SpotOpenOrdersRes](api.client.c, url, GET)
}

// binance SPOT allOrders rest所有订单 (USER_DATA)
func (client *SpotRestClient) NewAllOrders() *SpotAllOrdersApi {
	return &SpotAllOrdersApi{
		client: client,
		req:    &SpotAllOrdersReq{},
	}
}
func (api *SpotAllOrdersApi) Do() (*SpotAllOrdersRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotAllOrders], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SpotAllOrdersRes](api.client.c, url, GET)
}

// binance SPOT orderGet rest订单查询 (USER_DATA)
func (client *SpotRestClient) NewSpotOrderGet() *SpotOrderGetApi {
	return &SpotOrderGetApi{
		client: client,
		req:    &SpotOrderGetReq{},
	}
}
func (api *SpotOrderGetApi) Do() (*SpotOrderGetRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotOrderGet], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SpotOrderGetRes](api.client.c, url, GET)
}

// binance SPOT下单接口 SpotOrderPost rest 下单 (TRADE)
func (client *SpotRestClient) NewSpotOrderPost() *SpotOrderPostApi {
	return &SpotOrderPostApi{
		client: client,
		req:    &SpotOrderPostReq{},
	}
}
func (api *SpotOrderPostApi) Do() (*SpotOrderPostRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotOrderPost], api.client.c.ApiSecret)
	//log.Warn(url)
	return binanceCallApiWithSecret[SpotOrderPostRes](api.client.c, url, POST)
}

// binance SPOT orderDelete rest撤销订单 (TRADE)
func (client *SpotRestClient) NewSpotOrderDelete() *SpotOrderDeleteApi {
	return &SpotOrderDeleteApi{
		client: client,
		req:    &SpotOrderDeleteReq{},
	}
}
func (api *SpotOrderDeleteApi) Do() (*SpotOrderDeleteRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotOrderDelete], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SpotOrderDeleteRes](api.client.c, url, DELETE)
}

// binance SPOT orderCancelReplace rest撤消挂单再下单 (TRADE)
func (client *SpotRestClient) NewSpotOrderCancelReplace() *SpotOrderCancelReplaceApi {
	return &SpotOrderCancelReplaceApi{
		client: client,
		req:    &SpotOrderCancelReplaceReq{},
	}
}
func (api *SpotOrderCancelReplaceApi) Do() (*SpotOrderCancelReplaceRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotOrderCancelReplace], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SpotOrderCancelReplaceRes](api.client.c, url, POST)

}

// binance SPOT spotMyTrades rest账户成交历史 (USER_DATA)
func (client *SpotRestClient) NewSpotMyTrades() *SpotMyTradesApi {
	return &SpotMyTradesApi{
		client: client,
		req:    &SpotMyTradesReq{},
	}
}
func (api *SpotMyTradesApi) Do() (*SpotMyTradesRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotMyTrades], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SpotMyTradesRes](api.client.c, url, GET)
}
