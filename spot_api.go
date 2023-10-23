package mybinanceapi

import (
	"time"
)

type SpotApi int

const (

	//子母账户接口
	//GET
	SpotSubAccountList SpotApi = iota
	SpotSubAccountUniversalTransferHistory
	SpotSubAccountAssets
	SpotSubAccountFuturesAccount
	SpotSubAccountApiIpRestriction
	//POST
	SpotSubAccountVirtualSubAccount
	SpotSubAccountUniversalTransfer
	SpotSubAccountFuturesEnable

	//杠杆账户接口
	//GET
	SpotMarginAllPairs
	SpotMarginIsolatedAllPairs
	SpotMarginAccount
	SpotMarginIsolatedAccount
	SpotMarginMaxBorrowable
	SpotMarginMaxTransferable
	SpotMarginInterestHistory
	SpotMarginOrderGet
	SpotMarginAllOrders
	SpotMarginOpenOrders

	//POST
	SpotMarginTransfer
	SpotMarginIsolatedTransfer
	SpotMarginLoan
	SpotMarginRepay

	//下单接口
	SpotOrderPost
	SpotMarginOrderPost

	//钱包接口
	SpotAccountApiTradingStatus //账户API交易状态(USER_DATA)
	SpotAccount
	SpotOpenOrders
	SpotAllOrders
	SpotOrderGet

	//通用接口
	SpotPing
	SpotServerTime
	SpotExchangeInfo

	//行情接口
	SpotTickerPrice
)

var SpotApiMap = map[SpotApi]string{
	//子母账户接口
	//GET
	SpotSubAccountList:                     "/sapi/v1/sub-account/list",                        //GET接口 查询子账户列表(适用主账户)
	SpotSubAccountUniversalTransferHistory: "/sapi/v1/sub-account/universalTransfer",           //GET接口 子母账户万能划转历史查询 (适用主账户)
	SpotSubAccountAssets:                   "/sapi/v4/sub-account/assets",                      //GET接口 查询子账户资产(适用主账户)(USER_DATA)
	SpotSubAccountFuturesAccount:           "/sapi/v2/sub-account/futures/account",             //GET接口 查询子账户Futures账户详情V2 (适用主账户)
	SpotSubAccountApiIpRestriction:         "/sapi/v1/sub-account/subAccountApi/ipRestriction", //GET接口 查询子账户API Key IP白名单 (适用母账户)

	//POST
	SpotSubAccountVirtualSubAccount: "/sapi/v1/sub-account/virtualSubAccount", //POST接口 创建虚拟子账户(适用主账户)
	SpotSubAccountUniversalTransfer: "/sapi/v1/sub-account/universalTransfer", //POST接口 子母账户万能划转 (适用主账户)
	SpotSubAccountFuturesEnable:     "/sapi/v1/sub-account/futures/enable",    //POST接口 为子账户开通Futures (适用主账户)

	//杠杆账户接口
	//GET
	SpotMarginAllPairs:         "/sapi/v1/margin/allPairs",          //GET 获取所有全仓杠杆交易对(MARKET_DATA)
	SpotMarginIsolatedAllPairs: "/sapi/v1/margin/isolated/allPairs", //GET 获取所有逐仓杠杆交易对(USER_DATA)
	SpotMarginAccount:          "/sapi/v1/margin/account",           //GET 查询全仓杠杆账户详情 (USER_DATA)
	SpotMarginIsolatedAccount:  "/sapi/v1/margin/isolated/account",  //GET 查询逐仓杠杆账户详情 (USER_DATA)
	SpotMarginMaxBorrowable:    "/sapi/v1/margin/maxBorrowable",     //GET 查询账户最大可借贷额度(USER_DATA)
	SpotMarginMaxTransferable:  "/sapi/v1/margin/maxTransferable",   //GET 查询最大可转出额 (USER_DATA)
	SpotMarginInterestHistory:  "/sapi/v1/margin/interestHistory",   //GET 获取利息历史 (USER_DATA)
	SpotMarginOrderGet:         "/sapi/v1/margin/order",             //GET 查询杠杆账户订单 (USER_DATA)
	SpotMarginAllOrders:        "/sapi/v1/margin/allOrders",         //GET 查询杠杆账户所有订单 (USER_DATA)
	SpotMarginOpenOrders:       "/sapi/v1/margin/openOrders",        //GET 查询杠杆账户挂单记录 (USER_DATA)

	//POST
	SpotMarginTransfer:         "/sapi/v1/margin/transfer",          //POST  全仓杠杆账户划转 (MARGIN)
	SpotMarginIsolatedTransfer: "/sapi/v1/margin/isolated/transfer", //POST  逐仓杠杆账户划转 (MARGIN)
	SpotMarginLoan:             "/sapi/v1/margin/loan",              //POST  杠杆账户借贷 (MARGIN) 支持逐仓和全仓
	SpotMarginRepay:            "/sapi/v1/margin/repay",             //POST  杠杆账户归还借贷 (MARGIN) 支持逐仓和全仓

	//现货下单
	SpotOrderPost:       "/api/v3/order",         // POST /api/v3/order (HMAC SHA256) 下单 (TRADE)
	SpotMarginOrderPost: "/sapi/v1/margin/order", // POST /sapi/v1/margin/order (HMAC SHA256) 杠杆账户下单 (TRADE)
	//现货钱包接口
	SpotAccount:                 "/api/v3/account",                   //GET接口 账户信息 (USER_DATA)
	SpotAccountApiTradingStatus: "/sapi/v1/account/apiTradingStatus", //GET接口 账户API交易状态(USER_DATA)
	SpotOpenOrders:              "/api/v3/openOrders",                //GET接口 查询当前挂单 (USER_DATA)
	SpotAllOrders:               "/api/v3/allOrders",                 //GET接口 查询所有订单 (USER_DATA)
	SpotOrderGet:                "/api/v3/order",                     //GET接口 查询订单 (USER_DATA)

	//通用接口
	SpotPing:         "/api/v3/ping",         //GET接口 测试连通性
	SpotServerTime:   "/api/v3/time",         //GET接口 获取服务器时间
	SpotExchangeInfo: "/api/v3/exchangeInfo", //GET接口 获取交易规范

	//行情接口
	SpotTickerPrice: "/api/v3/ticker/price", //GET接口 获取交易对最新价格
}

// ================以下为子母账户GET接口
// binance SPOT子母账户接口  SubAccountList rest查询子账户列表
func (client *SpotRestClient) NewSubAccountList() *SpotSubAccountListApi {
	return &SpotSubAccountListApi{
		SpotRestClient: *client,
		req:            &SpotSubAccountListReq{},
	}
}
func (api *SpotSubAccountListApi) Do() (*SubAccountListRes, error) {
	api.Timestamp(time.Now().UnixMilli())
	url := binanceHandlerRequestApiWithSecretGet(SPOT, api.req, SpotApiMap[SpotSubAccountList], api.c.ApiSecret)
	return binanceCallApiWithSecretGet[SubAccountListRes](api.SpotRestClient.c, url)
}

// binance SPOT子母账户接口  SubAccountUniversalTransferHistory rest查询子母账户万能划转历史查询
func (client *SpotRestClient) NewSubAccountUniversalTransferHistory() *SpotSubAccountUniversalTransferHistoryApi {
	return &SpotSubAccountUniversalTransferHistoryApi{
		SpotRestClient: *client,
		req:            &SpotSubAccountUniversalTransferHistoryReq{},
	}
}
func (api *SpotSubAccountUniversalTransferHistoryApi) Do() (*SubAccountUniversalTransferHistoryRes, error) {
	api.Timestamp(time.Now().UnixMilli())
	url := binanceHandlerRequestApiWithSecretGet(SPOT, api.req, SpotApiMap[SpotSubAccountUniversalTransferHistory], api.c.ApiSecret)
	return binanceCallApiWithSecretGet[SubAccountUniversalTransferHistoryRes](api.SpotRestClient.c, url)
}

// binance SPOT子母账户接口  SubAccountAssets rest查询子账户资产 (适用主账户)
func (client *SpotRestClient) NewSubAccountAssets() *SpotSubAccountAssetsApi {
	return &SpotSubAccountAssetsApi{
		SpotRestClient: *client,
		req:            &SpotSubAccountAssetsReq{},
	}
}
func (api *SpotSubAccountAssetsApi) Do() (*SubAccountAssetsRes, error) {
	api.Timestamp(time.Now().UnixMilli())
	url := binanceHandlerRequestApiWithSecretGet(SPOT, api.req, SpotApiMap[SpotSubAccountAssets], api.c.ApiSecret)
	return binanceCallApiWithSecretGet[SubAccountAssetsRes](api.SpotRestClient.c, url)
}

// binance SPOT子母账户接口  SubAccountAssets rest查询子账户Futures账户详情V2 (适用主账户)
func (client *SpotRestClient) NewSubAccountFuturesAccount() *SpotSubAccountFuturesAccountApi {
	return &SpotSubAccountFuturesAccountApi{
		SpotRestClient: *client,
		req:            &SpotSubAccountFuturesAccountReq{},
	}
}
func (api *SpotSubAccountFuturesAccountApi) Do() (*SubAccountFuturesAccountRes, error) {
	api.Timestamp(time.Now().UnixMilli())
	url := binanceHandlerRequestApiWithSecretGet(SPOT, api.req, SpotApiMap[SpotSubAccountFuturesAccount], api.c.ApiSecret)
	return binanceCallApiWithSecretGet[SubAccountFuturesAccountRes](api.SpotRestClient.c, url)
}

// binance SPOT现货账户和交易接口  SpotAccount rest账户信息 (USER_DATA)
func (client *SpotRestClient) NewSpotAccount() *SpotAccountApi {
	return &SpotAccountApi{
		SpotRestClient: *client,
		req:            &SpotAccountReq{},
	}
}
func (api *SpotAccountApi) Do() (*SpotAccountRes, error) {
	api.Timestamp(time.Now().UnixMilli())
	url := binanceHandlerRequestApiWithSecretGet(SPOT, api.req, SpotApiMap[SpotAccount], api.c.ApiSecret)
	return binanceCallApiWithSecretGet[SpotAccountRes](api.SpotRestClient.c, url)
}

// binance 查询子账户API Key IP白名单 (适用母账户)  SpotSubAccountApiIpRestriction rest查询子账户API Key IP白名单 (适用母账户)
func (client *SpotRestClient) NewSpotSubAccountApiIpRestriction() *SpotSubAccountApiIpRestrictionApi {
	return &SpotSubAccountApiIpRestrictionApi{
		SpotRestClient: *client,
		req:            &SpotSubAccountApiIpRestrictionReq{},
	}
}
func (api *SpotSubAccountApiIpRestrictionApi) Do() (*SpotSubAccountApiIpRestrictionRes, error) {
	api.Timestamp(time.Now().UnixMilli())
	url := binanceHandlerRequestApiWithSecretGet(SPOT, api.req, SpotApiMap[SpotSubAccountApiIpRestriction], api.c.ApiSecret)
	return binanceCallApiWithSecretGet[SpotSubAccountApiIpRestrictionRes](api.SpotRestClient.c, url)
}

// ================以下为子母账户POST接口
// binance SPOT子母账户接口  SubAccountVirtualSubAccount rest创建虚拟子账户
func (client *SpotRestClient) NewSubAccountVirtualSubAccount() *SpotSubAccountVirtualSubAccountApi {
	return &SpotSubAccountVirtualSubAccountApi{
		SpotRestClient: *client,
		req:            &SpotSubAccountVirtualSubAccountReq{},
	}
}
func (api *SpotSubAccountVirtualSubAccountApi) Do() (*SubAccountVirtualSubAccountRes, error) {
	api.Timestamp(time.Now().UnixMilli())
	url, reqBody := binanceHandlerRequestApiWithSecretPost(SPOT, api.req, SpotApiMap[SpotSubAccountVirtualSubAccount], api.c.ApiSecret)
	return binanceCallApiWithSecretPost[SubAccountVirtualSubAccountRes](api.SpotRestClient.c, url, reqBody)
}

// binance SPOT子母账户接口  SubAccountUniversalTransfer rest子母账户万能划转
func (client *SpotRestClient) NewSubAccountUniversalTransfer() *SpotSubAccountUniversalTransferApi {
	return &SpotSubAccountUniversalTransferApi{
		SpotRestClient: *client,
		req:            &SpotSubAccountUniversalTransferReq{},
	}
}
func (api *SpotSubAccountUniversalTransferApi) Do() (*SubAccountUniversalTransferRes, error) {
	api.Timestamp(time.Now().UnixMilli())
	url, reqBody := binanceHandlerRequestApiWithSecretPost(SPOT, api.req, SpotApiMap[SpotSubAccountUniversalTransfer], api.c.ApiSecret)
	return binanceCallApiWithSecretPost[SubAccountUniversalTransferRes](api.SpotRestClient.c, url, reqBody)
}

// binance SPOT子母账户接口  SubAccountFuturesEnable rest为子账户开通Futures (适用主账户)
func (client *SpotRestClient) NewSubAccountFuturesEnable() *SpotSubAccountFuturesEnableApi {
	return &SpotSubAccountFuturesEnableApi{
		SpotRestClient: *client,
		req:            &SpotSubAccountFuturesEnableReq{},
	}
}
func (api *SpotSubAccountFuturesEnableApi) Do() (*SubAccountFuturesEnableRes, error) {
	api.Timestamp(time.Now().UnixMilli())
	url, reqBody := binanceHandlerRequestApiWithSecretPost(SPOT, api.req, SpotApiMap[SpotSubAccountFuturesEnable], api.c.ApiSecret)
	return binanceCallApiWithSecretPost[SubAccountFuturesEnableRes](api.SpotRestClient.c, url, reqBody)
}

// ================以下为杠杆账户GET接口
// binance SPOT杠杆接口  MarginAllPairs rest获取所有全仓杠杆交易对(MARKET_DATA)
func (client *SpotRestClient) NewMarginAllPairs() *SpotMarginAllPairsApi {
	return &SpotMarginAllPairsApi{
		SpotRestClient: *client,
		req:            &SpotMarginAllPairsReq{},
	}
}
func (api *SpotMarginAllPairsApi) Do() (*MarginAllPairsRes, error) {
	api.Timestamp(time.Now().UnixMilli())
	url := binanceHandlerRequestApiWithSecretGet(SPOT, api.req, SpotApiMap[SpotMarginAllPairs], api.c.ApiSecret)
	return binanceCallApiWithSecretGet[MarginAllPairsRes](api.SpotRestClient.c, url)
}

// binance SPOT杠杆接口  MarginIsolatedAllPairs rest获取所有逐仓杠杆交易对(MARKET_DATA)
func (client *SpotRestClient) NewMarginIsolatedAllPairs() *SpotMarginIsolatedAllPairsApi {
	return &SpotMarginIsolatedAllPairsApi{
		SpotRestClient: *client,
		req:            &SpotMarginIsolatedAllPairsReq{},
	}
}
func (api *SpotMarginIsolatedAllPairsApi) Do() (*MarginIsolatedAllPairsRes, error) {
	api.Timestamp(time.Now().UnixMilli())
	url := binanceHandlerRequestApiWithSecretGet(SPOT, api.req, SpotApiMap[SpotMarginIsolatedAllPairs], api.c.ApiSecret)
	return binanceCallApiWithSecretGet[MarginIsolatedAllPairsRes](api.SpotRestClient.c, url)
}

// binance SPOT杠杆接口  MarginAccount rest查询全仓杠杆账户详情 (USER_DATA)
func (client *SpotRestClient) NewMarginAccount() *SpotMarginAccountApi {
	return &SpotMarginAccountApi{
		SpotRestClient: *client,
		req:            &SpotMarginAccountReq{},
	}
}
func (api *SpotMarginAccountApi) Do() (*MarginAccountRes, error) {
	api.Timestamp(time.Now().UnixMilli())
	url := binanceHandlerRequestApiWithSecretGet(SPOT, api.req, SpotApiMap[SpotMarginAccount], api.c.ApiSecret)
	return binanceCallApiWithSecretGet[MarginAccountRes](api.SpotRestClient.c, url)
}

// binance SPOT杠杆接口  MarginIsolatedAccount rest查询逐仓杠杆账户详情 (USER_DATA)
func (client *SpotRestClient) NewMarginIsolatedAccount() *SpotMarginIsolatedAccountApi {
	return &SpotMarginIsolatedAccountApi{
		SpotRestClient: *client,
		req:            &SpotMarginIsolatedAccountReq{},
	}
}
func (api *SpotMarginIsolatedAccountApi) Do() (*MarginIsolatedAccountRes, error) {
	api.Timestamp(time.Now().UnixMilli())
	url := binanceHandlerRequestApiWithSecretGet(SPOT, api.req, SpotApiMap[SpotMarginIsolatedAccount], api.c.ApiSecret)
	return binanceCallApiWithSecretGet[MarginIsolatedAccountRes](api.SpotRestClient.c, url)
}

// binance SPOT杠杆接口 MarginMaxBorrowable rest查询最大可借 (MARKET_DATA)
func (client *SpotRestClient) NewMarginMaxBorrowable() *SpotMarginMaxBorrowableApi {
	return &SpotMarginMaxBorrowableApi{
		SpotRestClient: *client,
		req:            &SpotMarginMaxBorrowableReq{},
	}
}
func (api *SpotMarginMaxBorrowableApi) Do() (*MarginMaxBorrowableRes, error) {
	api.Timestamp(time.Now().UnixMilli())
	url := binanceHandlerRequestApiWithSecretGet(SPOT, api.req, SpotApiMap[SpotMarginMaxBorrowable], api.c.ApiSecret)
	return binanceCallApiWithSecretGet[MarginMaxBorrowableRes](api.SpotRestClient.c, url)
}

// binance SPOT杠杆接口  MarginMaxTransferable rest查询最大可转 (MARKET_DATA)
func (client *SpotRestClient) NewMarginMaxTransferable() *SpotMarginMaxTransferableApi {
	return &SpotMarginMaxTransferableApi{
		SpotRestClient: *client,
		req:            &SpotMarginMaxTransferableReq{},
	}
}
func (api *SpotMarginMaxTransferableApi) Do() (*MarginMaxTransferableRes, error) {
	api.Timestamp(time.Now().UnixMilli())
	url := binanceHandlerRequestApiWithSecretGet(SPOT, api.req, SpotApiMap[SpotMarginMaxTransferable], api.c.ApiSecret)
	return binanceCallApiWithSecretGet[MarginMaxTransferableRes](api.SpotRestClient.c, url)
}

// binance SPOT杠杆接口 MarginInterestHistory rest获取杠杆账户借息历史 (USER_DATA)
func (client *SpotRestClient) NewMarginInterestHistory() *SpotMarginInterestHistoryApi {
	return &SpotMarginInterestHistoryApi{
		SpotRestClient: *client,
		req:            &SpotMarginInterestHistoryReq{},
	}
}
func (api *SpotMarginInterestHistoryApi) Do() (*MarginInterestHistoryRes, error) {
	api.Timestamp(time.Now().UnixMilli())
	url := binanceHandlerRequestApiWithSecretGet(SPOT, api.req, SpotApiMap[SpotMarginInterestHistory], api.c.ApiSecret)
	return binanceCallApiWithSecretGet[MarginInterestHistoryRes](api.SpotRestClient.c, url)
}

// binance SPOT杠杆接口  MarginOrderGet rest查询杠杆账户订单 (USER_DATA)
func (client *SpotRestClient) NewMarginOrderGet() *SpotMarginOrderGetApi {
	return &SpotMarginOrderGetApi{
		SpotRestClient: *client,
		req:            &SpotMarginOrderGetReq{},
	}
}
func (api *SpotMarginOrderGetApi) Do() (*MarginOrderGetRes, error) {
	api.Timestamp(time.Now().UnixMilli())
	url := binanceHandlerRequestApiWithSecretGet(SPOT, api.req, SpotApiMap[SpotMarginOrderGet], api.c.ApiSecret)
	return binanceCallApiWithSecretGet[MarginOrderGetRes](api.SpotRestClient.c, url)
}

// binance  SPOT杠杆接口 MarginAllOrders rest查询杠杆账户全部订单 (USER_DATA)
func (client *SpotRestClient) NewMarginAllOrders() *SpotMarginAllOrdersApi {
	return &SpotMarginAllOrdersApi{
		SpotRestClient: *client,
		req:            &SpotMarginAllOrdersReq{},
	}
}
func (api *SpotMarginAllOrdersApi) Do() (*MarginAllOrdersRes, error) {
	api.Timestamp(time.Now().UnixMilli())
	url := binanceHandlerRequestApiWithSecretGet(SPOT, api.req, SpotApiMap[SpotMarginAllOrders], api.c.ApiSecret)
	return binanceCallApiWithSecretGet[MarginAllOrdersRes](api.SpotRestClient.c, url)
}

// binance SPOT杠杆接口 MarginOpenOrders  rest查询杠杆账户挂单记录 (USER_DATA)
func (client *SpotRestClient) NewMarginOpenOrders() *SpotMarginOpenOrdersApi {
	return &SpotMarginOpenOrdersApi{
		SpotRestClient: *client,
		req:            &SpotMarginOpenOrdersReq{},
	}
}
func (api *SpotMarginOpenOrdersApi) Do() (*MarginOpenOrdersRes, error) {
	api.Timestamp(time.Now().UnixMilli())
	url := binanceHandlerRequestApiWithSecretGet(SPOT, api.req, SpotApiMap[SpotMarginOpenOrders], api.c.ApiSecret)
	return binanceCallApiWithSecretGet[MarginOpenOrdersRes](api.SpotRestClient.c, url)
}

// ================以下为杠杆账户POST接口
// binance SPOT杠杆接口  MarginTransfer rest全仓杠杆账户划转 (MARGIN)
func (client *SpotRestClient) NewMarginTransfer() *SpotMarginTransferApi {
	return &SpotMarginTransferApi{
		SpotRestClient: *client,
		req:            &SpotMarginTransferReq{},
	}
}
func (api *SpotMarginTransferApi) Do() (*MarginTransferRes, error) {
	api.Timestamp(time.Now().UnixMilli())
	url, reqBody := binanceHandlerRequestApiWithSecretPost(SPOT, api.req, SpotApiMap[SpotMarginTransfer], api.c.ApiSecret)
	return binanceCallApiWithSecretPost[MarginTransferRes](api.SpotRestClient.c, url, reqBody)
}

// binance SPOT杠杆接口  MarginIsolatedTransfer rest逐仓杠杆账户划转 (MARGIN)
func (client *SpotRestClient) NewMarginIsolatedTransfer() *SpotMarginIsolatedTransferApi {
	return &SpotMarginIsolatedTransferApi{
		SpotRestClient: *client,
		req:            &SpotMarginIsolatedTransferReq{},
	}
}
func (api *SpotMarginIsolatedTransferApi) Do() (*MarginIsolatedTransferRes, error) {
	api.Timestamp(time.Now().UnixMilli())
	url, reqBody := binanceHandlerRequestApiWithSecretPost(SPOT, api.req, SpotApiMap[SpotMarginIsolatedTransfer], api.c.ApiSecret)
	return binanceCallApiWithSecretPost[MarginIsolatedTransferRes](api.SpotRestClient.c, url, reqBody)
}

// binance SPOT杠杆接口 MarginLoan rest 杠杆账户借贷 (MARGIN) 支持逐仓和全仓
func (client *SpotRestClient) NewMarginLoan() *SpotMarginLoanApi {
	return &SpotMarginLoanApi{
		SpotRestClient: *client,
		req:            &SpotMarginLoanReq{},
	}
}
func (api *SpotMarginLoanApi) Do() (*MarginLoanRes, error) {
	api.Timestamp(time.Now().UnixMilli())
	url, reqBody := binanceHandlerRequestApiWithSecretPost(SPOT, api.req, SpotApiMap[SpotMarginLoan], api.c.ApiSecret)
	return binanceCallApiWithSecretPost[MarginLoanRes](api.SpotRestClient.c, url, reqBody)
}

// binance SPOT杠杆接口 MarginRepay rest 杠杆账户还贷 (MARGIN) 支持逐仓和全仓
func (client *SpotRestClient) NewMarginRepay() *SpotMarginRepayApi {
	return &SpotMarginRepayApi{
		SpotRestClient: *client,
		req:            &SpotMarginRepayReq{},
	}
}
func (api *SpotMarginRepayApi) Do() (*MarginRepayRes, error) {
	api.Timestamp(time.Now().UnixMilli())
	url, reqBody := binanceHandlerRequestApiWithSecretPost(SPOT, api.req, SpotApiMap[SpotMarginRepay], api.c.ApiSecret)
	return binanceCallApiWithSecretPost[MarginRepayRes](api.SpotRestClient.c, url, reqBody)
}

// 下单接口
// binance SPOT下单接口 SpotOrderPost rest 下单 (TRADE)
func (client *SpotRestClient) NewSpotOrderPost() *SpotOrderPostApi {
	return &SpotOrderPostApi{
		SpotRestClient: *client,
		req:            &SpotOrderPostReq{},
	}
}
func (api *SpotOrderPostApi) Do() (*SpotOrderPostRes, error) {
	api.Timestamp(time.Now().UnixMilli())
	url, reqBody := binanceHandlerRequestApiWithSecretPost(SPOT, api.req, SpotApiMap[SpotOrderPost], api.c.ApiSecret)
	log.Warn(url)
	return binanceCallApiWithSecretPost[SpotOrderPostRes](api.SpotRestClient.c, url, reqBody)
}

// binance SPOT杠杆下单接口 SpotMarginOrderPost rest 杠杆下单 (TRADE)
func (client *SpotRestClient) NewSpotMarginOrderPost() *SpotMarginOrderPostApi {
	return &SpotMarginOrderPostApi{
		SpotRestClient: *client,
		req:            &SpotMarginOrderPostReq{},
	}
}
func (api *SpotMarginOrderPostApi) Do() (*SpotMarginOrderPostRes, error) {
	api.Timestamp(time.Now().UnixMilli())
	url, reqBody := binanceHandlerRequestApiWithSecretPost(SPOT, api.req, SpotApiMap[SpotMarginOrderPost], api.c.ApiSecret)
	log.Debug(url)
	return binanceCallApiWithSecretPost[SpotMarginOrderPostRes](api.SpotRestClient.c, url, reqBody)
}

// 钱包接口
// binance SPOT钱包接口 账户API交易状态(USER_DATA)
func (client *SpotRestClient) NewAccountApiTradingStatus() *SpotAccountApiTradingStatusApi {
	return &SpotAccountApiTradingStatusApi{
		SpotRestClient: *client,
		req:            &SpotAccountApiTradingStatusReq{},
	}
}
func (api *SpotAccountApiTradingStatusApi) Do() (*SpotAccountApiTradingStatusRes, error) {
	api.Timestamp(time.Now().UnixMilli())
	url := binanceHandlerRequestApiWithSecretGet(SPOT, api.req, SpotApiMap[SpotAccountApiTradingStatus], api.c.ApiSecret)
	return binanceCallApiWithSecretGet[SpotAccountApiTradingStatusRes](api.SpotRestClient.c, url)
}

// 通用接口
// binance SPOT ping rest测试连通性 (NONE)
func (client *SpotRestClient) NewPing() *SpotPingApi {
	return &SpotPingApi{
		SpotRestClient: *client,
		req:            &SpotPingReq{},
	}
}
func (api *SpotPingApi) Do() (*SpotPingRes, error) {
	url := binanceHandlerRequestApiGet(SPOT, api.req, SpotApiMap[SpotPing])
	return binanceCallApiWithSecretGet[SpotPingRes](api.SpotRestClient.c, url)
}

// binance SPOT serverTime rest服务器时间 (NONE)
func (client *SpotRestClient) NewServerTime() *SpotServerTimeApi {
	return &SpotServerTimeApi{
		SpotRestClient: *client,
		req:            &SpotServerTimeReq{},
	}
}
func (api *SpotServerTimeApi) Do() (*SpotServerTimeRes, error) {
	url := binanceHandlerRequestApiGet(SPOT, api.req, SpotApiMap[SpotServerTime])
	return binanceCallApiWithSecretGet[SpotServerTimeRes](api.SpotRestClient.c, url)
}

// binance SPOT exchangeInfo rest交易规范
func (client *SpotRestClient) NewExchangeInfo() *SpotExchangeInfoApi {
	return &SpotExchangeInfoApi{
		SpotRestClient: *client,
		req:            &SpotExchangeInfoReq{},
	}
}
func (api *SpotExchangeInfoApi) Do() (*SpotExchangeInfoRes, error) {
	url := binanceHandlerRequestApiGet(SPOT, api.req, SpotApiMap[SpotExchangeInfo])
	return binanceCallApiWithSecretGet[SpotExchangeInfoRes](api.SpotRestClient.c, url)
}

// binance SPOT openOrders rest当前挂单 (USER_DATA)
func (client *SpotRestClient) NewOpenOrders() *SpotOpenOrdersApi {
	return &SpotOpenOrdersApi{
		SpotRestClient: *client,
		req:            &SpotOpenOrdersReq{},
	}
}
func (api *SpotOpenOrdersApi) Do() (*SpotOpenOrdersRes, error) {
	api.Timestamp(time.Now().UnixMilli())
	url := binanceHandlerRequestApiWithSecretGet(SPOT, api.req, SpotApiMap[SpotOpenOrders], api.c.ApiSecret)
	return binanceCallApiWithSecretGet[SpotOpenOrdersRes](api.SpotRestClient.c, url)
}

// binance SPOT allOrders rest所有订单 (USER_DATA)
func (client *SpotRestClient) NewAllOrders() *SpotAllOrdersApi {
	return &SpotAllOrdersApi{
		SpotRestClient: *client,
		req:            &SpotAllOrdersReq{},
	}
}
func (api *SpotAllOrdersApi) Do() (*SpotAllOrdersRes, error) {
	api.Timestamp(time.Now().UnixMilli())
	url := binanceHandlerRequestApiWithSecretGet(SPOT, api.req, SpotApiMap[SpotAllOrders], api.c.ApiSecret)
	return binanceCallApiWithSecretGet[SpotAllOrdersRes](api.SpotRestClient.c, url)
}

// binance SPOT orderGet rest订单查询 (USER_DATA)
func (client *SpotRestClient) NewSpotOrderGet() *SpotOrderGetApi {
	return &SpotOrderGetApi{
		SpotRestClient: *client,
		req:            &SpotOrderGetReq{},
	}
}
func (api *SpotOrderGetApi) Do() (*SpotOrderGetRes, error) {
	api.Timestamp(time.Now().UnixMilli())
	url := binanceHandlerRequestApiWithSecretGet(SPOT, api.req, SpotApiMap[SpotOrderGet], api.c.ApiSecret)
	return binanceCallApiWithSecretGet[SpotOrderGetRes](api.SpotRestClient.c, url)
}

// binance SPOT spotTickerPrice rest价格 (NONE)
func (client *SpotRestClient) NewSpotTickerPrice() *SpotTickerPriceApi {
	return &SpotTickerPriceApi{
		SpotRestClient: *client,
		req:            &SpotTickerPriceReq{},
	}
}
func (api *SpotTickerPriceApi) Do() (*SpotTickerPriceRes, error) {
	url := binanceHandlerRequestApiGet(SPOT, api.req, SpotApiMap[SpotTickerPrice])
	return binanceCallApiWithSecretGet[SpotTickerPriceRes](api.SpotRestClient.c, url)
}
