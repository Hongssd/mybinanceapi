package mybinanceapi

import "time"

// 杠杆账户接口
// binance SPOT杠杆接口  MarginAllPairs rest获取所有全仓杠杆交易对(MARKET_DATA)
func (client *SpotRestClient) NewSpotMarginAllPairs() *SpotMarginAllPairsApi {
	return &SpotMarginAllPairsApi{
		client: client,
		req:    &SpotMarginAllPairsReq{},
	}
}
func (api *SpotMarginAllPairsApi) Do() (*MarginAllPairsRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotMarginAllPairs], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[MarginAllPairsRes](api.client.c, url, GET)
}

// binance SPOT杠杆接口  MarginIsolatedAllPairs rest获取所有逐仓杠杆交易对(MARKET_DATA)
func (client *SpotRestClient) NewSpotMarginIsolatedAllPairs() *SpotMarginIsolatedAllPairsApi {
	return &SpotMarginIsolatedAllPairsApi{
		client: client,
		req:    &SpotMarginIsolatedAllPairsReq{},
	}
}
func (api *SpotMarginIsolatedAllPairsApi) Do() (*MarginIsolatedAllPairsRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotMarginIsolatedAllPairs], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[MarginIsolatedAllPairsRes](api.client.c, url, GET)
}

// binance SPOT杠杆接口  MarginAccount rest查询全仓杠杆账户详情 (USER_DATA)
func (client *SpotRestClient) NewSpotMarginAccount() *SpotMarginAccountApi {
	return &SpotMarginAccountApi{
		client: client,
		req:    &SpotMarginAccountReq{},
	}
}
func (api *SpotMarginAccountApi) Do() (*MarginAccountRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotMarginAccount], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[MarginAccountRes](api.client.c, url, GET)
}

// binance SPOT杠杆接口  MarginIsolatedAccount rest查询逐仓杠杆账户详情 (USER_DATA)
func (client *SpotRestClient) NewSpotMarginIsolatedAccount() *SpotMarginIsolatedAccountApi {
	return &SpotMarginIsolatedAccountApi{
		client: client,
		req:    &SpotMarginIsolatedAccountReq{},
	}
}
func (api *SpotMarginIsolatedAccountApi) Do() (*MarginIsolatedAccountRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotMarginIsolatedAccount], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[MarginIsolatedAccountRes](api.client.c, url, GET)
}

// binance SPOT杠杆接口 MarginMaxBorrowable rest查询最大可借 (MARKET_DATA)
func (client *SpotRestClient) NewSpotMarginMaxBorrowable() *SpotMarginMaxBorrowableApi {
	return &SpotMarginMaxBorrowableApi{
		client: client,
		req:    &SpotMarginMaxBorrowableReq{},
	}
}
func (api *SpotMarginMaxBorrowableApi) Do() (*MarginMaxBorrowableRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotMarginMaxBorrowable], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[MarginMaxBorrowableRes](api.client.c, url, GET)
}

// binance SPOT杠杆接口  MarginMaxTransferable rest查询最大可转 (MARKET_DATA)
func (client *SpotRestClient) NewSpotMarginMaxTransferable() *SpotMarginMaxTransferableApi {
	return &SpotMarginMaxTransferableApi{
		client: client,
		req:    &SpotMarginMaxTransferableReq{},
	}
}
func (api *SpotMarginMaxTransferableApi) Do() (*MarginMaxTransferableRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotMarginMaxTransferable], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[MarginMaxTransferableRes](api.client.c, url, GET)
}

// binance SPOT杠杆接口 MarginInterestHistory rest获取杠杆账户借息历史 (USER_DATA)
func (client *SpotRestClient) NewSpotMarginInterestHistory() *SpotMarginInterestHistoryApi {
	return &SpotMarginInterestHistoryApi{
		client: client,
		req:    &SpotMarginInterestHistoryReq{},
	}
}
func (api *SpotMarginInterestHistoryApi) Do() (*MarginInterestHistoryRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotMarginInterestHistory], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[MarginInterestHistoryRes](api.client.c, url, GET)
}

// binance SPOT杠杆接口  MarginOrderGet rest查询杠杆账户订单 (USER_DATA)
func (client *SpotRestClient) NewSpotMarginOrderGet() *SpotMarginOrderGetApi {
	return &SpotMarginOrderGetApi{
		client: client,
		req:    &SpotMarginOrderGetReq{},
	}
}
func (api *SpotMarginOrderGetApi) Do() (*MarginOrderGetRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotMarginOrderGet], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[MarginOrderGetRes](api.client.c, url, GET)
}

// binance SPOT杠杆下单接口 SpotMarginOrderPost rest 杠杆下单 (TRADE)
func (client *SpotRestClient) NewSpotMarginOrderPost() *SpotMarginOrderPostApi {
	return &SpotMarginOrderPostApi{
		client: client,
		req:    &SpotMarginOrderPostReq{},
	}
}
func (api *SpotMarginOrderPostApi) Do() (*SpotMarginOrderPostRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotMarginOrderPost], api.client.c.ApiSecret)
	log.Debug(url)
	return binanceCallApiWithSecret[SpotMarginOrderPostRes](api.client.c, url, POST)
}

// binance SPOT杠杆接口  MarginOrderDelete rest撤销杠杆账户订单 (TRADE)
func (client *SpotRestClient) NewSpotMarginOrderDelete() *SpotMarginOrderDeleteApi {
	return &SpotMarginOrderDeleteApi{
		client: client,
		req:    &SpotMarginOrderDeleteReq{},
	}
}
func (api *SpotMarginOrderDeleteApi) Do() (*SpotMarginOrderDeleteRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotMarginOrderDelete], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SpotMarginOrderDeleteRes](api.client.c, url, DELETE)
}

// binance  SPOT杠杆接口 MarginAllOrders rest查询杠杆账户全部订单 (USER_DATA)
func (client *SpotRestClient) NewSpotMarginAllOrders() *SpotMarginAllOrdersApi {
	return &SpotMarginAllOrdersApi{
		client: client,
		req:    &SpotMarginAllOrdersReq{},
	}
}
func (api *SpotMarginAllOrdersApi) Do() (*MarginAllOrdersRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotMarginAllOrders], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[MarginAllOrdersRes](api.client.c, url, GET)
}

// binance SPOT杠杆接口 MarginOpenOrders  rest查询杠杆账户挂单记录 (USER_DATA)
func (client *SpotRestClient) NewSpotMarginOpenOrders() *SpotMarginOpenOrdersApi {
	return &SpotMarginOpenOrdersApi{
		client: client,
		req:    &SpotMarginOpenOrdersReq{},
	}
}
func (api *SpotMarginOpenOrdersApi) Do() (*MarginOpenOrdersRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotMarginOpenOrders], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[MarginOpenOrdersRes](api.client.c, url, GET)
}

// binance SPOT杠杆接口  MarginTransfer rest全仓杠杆账户划转 (MARGIN)
func (client *SpotRestClient) NewSpotMarginTransfer() *SpotMarginTransferApi {
	return &SpotMarginTransferApi{
		client: client,
		req:    &SpotMarginTransferReq{},
	}
}
func (api *SpotMarginTransferApi) Do() (*MarginTransferRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotMarginTransfer], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[MarginTransferRes](api.client.c, url, POST)
}

// binance SPOT杠杆接口  MarginIsolatedTransfer rest逐仓杠杆账户划转 (MARGIN)
func (client *SpotRestClient) NewSpotMarginIsolatedTransfer() *SpotMarginIsolatedTransferApi {
	return &SpotMarginIsolatedTransferApi{
		client: client,
		req:    &SpotMarginIsolatedTransferReq{},
	}
}
func (api *SpotMarginIsolatedTransferApi) Do() (*MarginIsolatedTransferRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotMarginIsolatedTransfer], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[MarginIsolatedTransferRes](api.client.c, url, POST)
}

// binance SPOT杠杆接口 MarginLoan rest 杠杆账户借贷 (MARGIN) 支持逐仓和全仓
func (client *SpotRestClient) NewSpotMarginLoan() *SpotMarginLoanApi {
	return &SpotMarginLoanApi{
		client: client,
		req:    &SpotMarginLoanReq{},
	}
}
func (api *SpotMarginLoanApi) Do() (*MarginLoanRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotMarginLoan], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[MarginLoanRes](api.client.c, url, POST)
}

// binance SPOT杠杆接口 MarginRepay rest 杠杆账户还贷 (MARGIN) 支持逐仓和全仓
func (client *SpotRestClient) NewSpotMarginRepay() *SpotMarginRepayApi {
	return &SpotMarginRepayApi{
		client: client,
		req:    &SpotMarginRepayReq{},
	}
}
func (api *SpotMarginRepayApi) Do() (*MarginRepayRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotMarginRepay], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[MarginRepayRes](api.client.c, url, POST)
}

func (client *SpotRestClient) NewSpotMarginMaxLeverage() *SpotMarginMaxLeverageApi {
	return &SpotMarginMaxLeverageApi{
		client: client,
		req:    &SpotMarginMaxLeverageReq{},
	}
}
func (api *SpotMarginMaxLeverageApi) Do() (*MarginMaxLeverageRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotMarginMaxLeverage], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[MarginMaxLeverageRes](api.client.c, url, POST)
}
