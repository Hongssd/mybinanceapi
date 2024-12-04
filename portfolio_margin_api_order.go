package mybinanceapi

import "time"

// binance PortfolioMarginUmOrderPost rest UM下单（TRADE)
func (client *PortfolioMarginRestClient) NewUmOrderPost() *PortfolioMarginUmOrderPostApi {
	return &PortfolioMarginUmOrderPostApi{
		client: client,
		req:    &PortfolioMarginUmOrderPostReq{},
	}
}

func (api *PortfolioMarginUmOrderPostApi) Do() (*PortfolioMarginUmOrderPostRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PORTFOLIO_MARGIN, api.req, PortfolioMarginApiMap[PortfolioMarginUmOrderPost], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PortfolioMarginUmOrderPostRes](api.client.c, url, POST)
}

// binance PortfolioMarginUmOrderDelete rest UM撤单（TRADE)
func (client *PortfolioMarginRestClient) NewUmOrderDelete() *PortfolioMarginUmOrderDeleteApi {
	return &PortfolioMarginUmOrderDeleteApi{
		client: client,
		req:    &PortfolioMarginUmOrderDeleteReq{},
	}
}

func (api *PortfolioMarginUmOrderDeleteApi) Do() (*PortfolioMarginUmOrderDeleteRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PORTFOLIO_MARGIN, api.req, PortfolioMarginApiMap[PortfolioMarginUmOrderDelete], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PortfolioMarginUmOrderDeleteRes](api.client.c, url, DELETE)
}

// binance PortfolioMarginUmAllOpenOrdersDelete rest 撤销UM当前全部挂单（TRADE)
func (client *PortfolioMarginRestClient) NewUmAllOpenOrdersDelete() *PortfolioMarginUmAllOpenOrdersDeleteApi {
	return &PortfolioMarginUmAllOpenOrdersDeleteApi{
		client: client,
		req:    &PortfolioMarginUmAllOpenOrdersDeleteReq{},
	}
}

func (api *PortfolioMarginUmAllOpenOrdersDeleteApi) Do() (*PortfolioMarginUmAllOpenOrdersDeleteRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PORTFOLIO_MARGIN, api.req, PortfolioMarginApiMap[PortfolioMarginUmAllOpenOrdersDelete], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PortfolioMarginUmAllOpenOrdersDeleteRes](api.client.c, url, DELETE)
}

// binance PortfolioMarginUmOrderPut rest UM修改订单（TRADE)
func (client *PortfolioMarginRestClient) NewUmOrderPut() *PortfolioMarginUmOrderPutApi {
	return &PortfolioMarginUmOrderPutApi{
		client: client,
		req:    &PortfolioMarginUmOrderPutReq{},
	}
}

func (api *PortfolioMarginUmOrderPutApi) Do() (*PortfolioMarginUmOrderPutRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PORTFOLIO_MARGIN, api.req, PortfolioMarginApiMap[PortfolioMarginUmOrderPut], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PortfolioMarginUmOrderPutRes](api.client.c, url, PUT)
}

// binance PortfolioMarginUmConditionalOrderPost rest UM条件下单（TRADE)
func (client *PortfolioMarginRestClient) NewUmConditionalOrderPost() *PortfolioMarginUmConditionalOrderPostApi {
	return &PortfolioMarginUmConditionalOrderPostApi{
		client: client,
		req:    &PortfolioMarginUmConditionalOrderReq{},
	}
}

func (api *PortfolioMarginUmConditionalOrderPostApi) Do() (*PortfolioMarginUmConditionalOrderPostRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PORTFOLIO_MARGIN, api.req, PortfolioMarginApiMap[PortfolioMarginUmConditionalOrderPost], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PortfolioMarginUmConditionalOrderPostRes](api.client.c, url, POST)
}

// binance PortfolioMarginUmConditionalOpenOrder rest 查询UM当前条件挂单（TRADE)
func (client *PortfolioMarginRestClient) NewUmConditionalOpenOrder() *PortfolioMarginUmConditionalOpenOrderApi {
	return &PortfolioMarginUmConditionalOpenOrderApi{
		client: client,
		req:    &PortfolioMarginUmConditionalOpenOrderReq{},
	}
}

func (api *PortfolioMarginUmConditionalOpenOrderApi) Do() (*PortfolioMarginUmConditionalOpenOrderRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PORTFOLIO_MARGIN, api.req, PortfolioMarginApiMap[PortfolioMarginUmConditionalOpenOrder], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PortfolioMarginUmConditionalOpenOrderRes](api.client.c, url, GET)
}

// binance PortfolioMarginUmOrderGet rest 查询UM订单（TRADE)
func (client *PortfolioMarginRestClient) NewUmOrderGet() *PortfolioMarginUmOrderGetApi {
	return &PortfolioMarginUmOrderGetApi{
		client: client,
		req:    &PortfolioMarginUmOrderGetReq{},
	}
}

func (api *PortfolioMarginUmOrderGetApi) Do() (*PortfolioMarginUmOrderGetRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PORTFOLIO_MARGIN, api.req, PortfolioMarginApiMap[PortfolioMarginUmOrderGet], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PortfolioMarginUmOrderGetRes](api.client.c, url, GET)
}

// binance PortfolioMarginUmOpenOrderGet rest 查询UM当前挂单（TRADE)
func (client *PortfolioMarginRestClient) NewUmOpenOrderGet() *PortfolioMarginUmOpenOrderGetApi {
	return &PortfolioMarginUmOpenOrderGetApi{
		client: client,
		req:    &PortfolioMarginUmOpenOrderGetReq{},
	}
}

func (api *PortfolioMarginUmOpenOrderGetApi) Do() (*PortfolioMarginUmOpenOrderGetRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PORTFOLIO_MARGIN, api.req, PortfolioMarginApiMap[PortfolioMarginUmOpenOrderGet], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PortfolioMarginUmOpenOrderGetRes](api.client.c, url, GET)
}

// binance PortfolioMarginUmOpenOrdersGet rest 查询UM当前全部挂单（TRADE)
func (client *PortfolioMarginRestClient) NewUmOpenOrdersGet() *PortfolioMarginUmOpenOrdersGetApi {
	return &PortfolioMarginUmOpenOrdersGetApi{
		client: client,
		req:    &PortfolioMarginUmOpenOrdersGetReq{},
	}
}

func (api *PortfolioMarginUmOpenOrdersGetApi) Do() (*PortfolioMarginUmOpenOrdersGetRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PORTFOLIO_MARGIN, api.req, PortfolioMarginApiMap[PortfolioMarginUmOpenOrdersGet], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PortfolioMarginUmOpenOrdersGetRes](api.client.c, url, GET)
}

// binance PortfolioMarginCmOrderPost rest CM下单（TRADE)
func (client *PortfolioMarginRestClient) NewCmOrderPost() *PortfolioMarginCmOrderPostApi {
	return &PortfolioMarginCmOrderPostApi{
		client: client,
		req:    &PortfolioMarginCmOrderPostReq{},
	}
}

func (api *PortfolioMarginCmOrderPostApi) Do() (*PortfolioMarginCmOrderPostRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PORTFOLIO_MARGIN, api.req, PortfolioMarginApiMap[PortfolioMarginCmOrderPost], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PortfolioMarginCmOrderPostRes](api.client.c, url, POST)
}

// binance PortfolioMarginCmConditionalOrderPost rest CM条件下单（TRADE)
func (client *PortfolioMarginRestClient) NewCmConditionalOrderPost() *PortfolioMarginCmConditionalOrderPostApi {
	return &PortfolioMarginCmConditionalOrderPostApi{
		client: client,
		req:    &PortfolioMarginCmConditionalOrderPostReq{},
	}
}

func (api *PortfolioMarginCmConditionalOrderPostApi) Do() (*PortfolioMarginCmConditionalOrderPostRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PORTFOLIO_MARGIN, api.req, PortfolioMarginApiMap[PortfolioMarginCmConditionalOrderPost], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PortfolioMarginCmConditionalOrderPostRes](api.client.c, url, POST)
}

// binance PortfolioMarginMarginOrderPost rest 杠杆下单（TRADE)
func (client *PortfolioMarginRestClient) NewMarginOrderPost() *PortfolioMarginMarginOrderPostApi {
	return &PortfolioMarginMarginOrderPostApi{
		client: client,
		req:    &PortfolioMarginMarginOrderPostReq{},
	}
}

func (api *PortfolioMarginMarginOrderPostApi) Do() (*PortfolioMarginMarginOrderPostRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PORTFOLIO_MARGIN, api.req, PortfolioMarginApiMap[PortfolioMarginMarginOrderPost], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PortfolioMarginMarginOrderPostRes](api.client.c, url, POST)
}

// binance PortfolioMarginMarginOrderOcoPost rest 杠杆OCO下单（TRADE)
func (client *PortfolioMarginRestClient) NewMarginOrderOcoPost() *PortfolioMarginMarginOrderOcoPostApi {
	return &PortfolioMarginMarginOrderOcoPostApi{
		client: client,
		req:    &PortfolioMarginMarginOrderOcoPostReq{},
	}
}

func (api *PortfolioMarginMarginOrderOcoPostApi) Do() (*PortfolioMarginMarginOrderOcoPostRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PORTFOLIO_MARGIN, api.req, PortfolioMarginApiMap[PortfolioMarginMarginOrderOcoPost], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PortfolioMarginMarginOrderOcoPostRes](api.client.c, url, POST)
}

// binance PortfolioMarginUmConditionalOpenOrders rest 查询UM当前条件全部挂单（TRADE)
func (client *PortfolioMarginRestClient) NewUmConditionalOpenOrders() *PortfolioMarginUmConditionalOpenOrdersApi {
	return &PortfolioMarginUmConditionalOpenOrdersApi{
		client: client,
		req:    &PortfolioMarginUmConditionalOpenOrdersReq{},
	}
}

func (api *PortfolioMarginUmConditionalOpenOrdersApi) Do() (*PortfolioMarginUmConditionalOpenOrdersRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PORTFOLIO_MARGIN, api.req, PortfolioMarginApiMap[PortfolioMarginUmConditionalOpenOrders], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PortfolioMarginUmConditionalOpenOrdersRes](api.client.c, url, GET)
}

// binance PortfolioMarginUmConditionalOrderHistory rest 查询UM条件订单历史（TRADE)
func (client *PortfolioMarginRestClient) NewUmConditionalOrderHistory() *PortfolioMarginUmConditionalOrderHistoryApi {
	return &PortfolioMarginUmConditionalOrderHistoryApi{
		client: client,
		req:    &PortfolioMarginUmConditionalOrderHistoryReq{},
	}
}

func (api *PortfolioMarginUmConditionalOrderHistoryApi) Do() (*PortfolioMarginUmConditionalOrderHistoryRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PORTFOLIO_MARGIN, api.req, PortfolioMarginApiMap[PortfolioMarginUmConditionalOrderHistory], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PortfolioMarginUmConditionalOrderHistoryRes](api.client.c, url, GET)
}

// binance PortfolioMarginCmOrderDelete rest CM撤单（TRADE)
func (client *PortfolioMarginRestClient) NewCmOrderDelete() *PortfolioMarginCmOrderDeleteApi {
	return &PortfolioMarginCmOrderDeleteApi{
		client: client,
		req:    &PortfolioMarginCmOrderDeleteReq{},
	}
}

func (api *PortfolioMarginCmOrderDeleteApi) Do() (*PortfolioMarginCmOrderDeleteRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PORTFOLIO_MARGIN, api.req, PortfolioMarginApiMap[PortfolioMarginCmOrderDelete], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PortfolioMarginCmOrderDeleteRes](api.client.c, url, DELETE)
}

// binance PortfolioMarginCmAllOpenOrdersDelete rest 撤销CM当前全部挂单（TRADE)
func (client *PortfolioMarginRestClient) NewCmAllOpenOrdersDelete() *PortfolioMarginCmAllOpenOrdersDeleteApi {
	return &PortfolioMarginCmAllOpenOrdersDeleteApi{
		client: client,
		req:    &PortfolioMarginCmAllOpenOrdersDeleteReq{},
	}
}

func (api *PortfolioMarginCmAllOpenOrdersDeleteApi) Do() (*PortfolioMarginCmAllOpenOrdersDeleteRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PORTFOLIO_MARGIN, api.req, PortfolioMarginApiMap[PortfolioMarginCmAllOpenOrdersDelete], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PortfolioMarginCmAllOpenOrdersDeleteRes](api.client.c, url, DELETE)
}

// binance PortfolioMarginCmOrderPut rest CM修改订单（TRADE)
func (client *PortfolioMarginRestClient) NewCmOrderPut() *PortfolioMarginCmOrderPutApi {
	return &PortfolioMarginCmOrderPutApi{
		client: client,
		req:    &PortfolioMarginCmOrderPutReq{},
	}
}

func (api *PortfolioMarginCmOrderPutApi) Do() (*PortfolioMarginCmOrderPutRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PORTFOLIO_MARGIN, api.req, PortfolioMarginApiMap[PortfolioMarginCmOrderPut], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PortfolioMarginCmOrderPutRes](api.client.c, url, PUT)
}

// binance PortfolioMarginCmConditionalOrderDelete rest 取消CM条件订单（TRADE)
func (client *PortfolioMarginRestClient) NewCmConditionalOrderDelete() *PortfolioMarginCmConditionalOrderDeleteApi {
	return &PortfolioMarginCmConditionalOrderDeleteApi{
		client: client,
		req:    &PortfolioMarginCmConditionalOrderDeleteReq{},
	}
}

func (api *PortfolioMarginCmConditionalOrderDeleteApi) Do() (*PortfolioMarginCmConditionalOrderDeleteRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PORTFOLIO_MARGIN, api.req, PortfolioMarginApiMap[PortfolioMarginCmConditionalOrderDelete], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PortfolioMarginCmConditionalOrderDeleteRes](api.client.c, url, DELETE)
}

// binance PortfolioMarginCmConditionalAllOpenOrdersDelete rest 取消全部CM条件单（TRADE)
func (client *PortfolioMarginRestClient) NewCmConditionalAllOpenOrdersDelete() *PortfolioMarginCmConditionalAllOpenOrdersDeleteApi {
	return &PortfolioMarginCmConditionalAllOpenOrdersDeleteApi{
		client: client,
		req:    &PortfolioMarginCmConditionalAllOpenOrdersDeleteReq{},
	}
}

func (api *PortfolioMarginCmConditionalAllOpenOrdersDeleteApi) Do() (*PortfolioMarginCmConditionalAllOpenOrdersDeleteRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PORTFOLIO_MARGIN, api.req, PortfolioMarginApiMap[PortfolioMarginCmConditionalAllOpenOrdersDelete], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PortfolioMarginCmConditionalAllOpenOrdersDeleteRes](api.client.c, url, DELETE)
}

// binance PortfolioMarginCmOrderGet rest 查询CM订单（TRADE)
func (client *PortfolioMarginRestClient) NewCmOrderGet() *PortfolioMarginCmOrderGetApi {
	return &PortfolioMarginCmOrderGetApi{
		client: client,
		req:    &PortfolioMarginCmOrderGetReq{},
	}
}

func (api *PortfolioMarginCmOrderGetApi) Do() (*PortfolioMarginCmOrderGetRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PORTFOLIO_MARGIN, api.req, PortfolioMarginApiMap[PortfolioMarginCmOrderGet], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PortfolioMarginCmOrderGetRes](api.client.c, url, GET)
}

// binance PortfolioMarginCmAllOrders rest 查询CM所有订单（TRADE)
func (client *PortfolioMarginRestClient) NewCmAllOrders() *PortfolioMarginCmAllOrdersApi {
	return &PortfolioMarginCmAllOrdersApi{
		client: client,
		req:    &PortfolioMarginCmAllOrdersReq{},
	}
}

func (api *PortfolioMarginCmAllOrdersApi) Do() (*PortfolioMarginCmAllOrdersRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PORTFOLIO_MARGIN, api.req, PortfolioMarginApiMap[PortfolioMarginCmAllOrders], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PortfolioMarginCmAllOrdersRes](api.client.c, url, GET)
}

// binance PortfolioMarginCmConditionalOpenOrder rest 查询CM当前条件挂单（TRADE)
func (client *PortfolioMarginRestClient) NewCmConditionalOpenOrder() *PortfolioMarginCmConditionalOpenOrderApi {
	return &PortfolioMarginCmConditionalOpenOrderApi{
		client: client,
		req:    &PortfolioMarginCmConditionalOpenOrderReq{},
	}
}

func (api *PortfolioMarginCmConditionalOpenOrderApi) Do() (*PortfolioMarginCmConditionalOpenOrderRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PORTFOLIO_MARGIN, api.req, PortfolioMarginApiMap[PortfolioMarginCmConditionalOpenOrder], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PortfolioMarginCmConditionalOpenOrderRes](api.client.c, url, GET)
}
