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
	SpotSubAccountTransferSubUserHistory
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
	SpotMarginOrderDelete
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
	SpotOrderDelete       //撤销订单 (TRADE)
	SpotAssetTransferPost //用户万向划转 (USER_DATA)
	SpotAssetTransferGet  //查询用户万向划转历史 (USER_DATA)
	SpotAssetTradeFee     //查询用户交易手续费率 (USER_DATA)

	//现货账户接口
	SpotMyTrades //账户成交历史 (USER_DATA)

	//通用接口
	SpotPing         //测试服务器连通性
	SpotServerTime   //获取服务器时间
	SpotExchangeInfo //获取交易规则和交易对信息。

	//行情接口
	SpotKlines      //K线数据
	SpotTickerPrice //获取交易对最新价格
	SpotDepth       //获取深度信息

	SpotTrades           //近期成交列表
	SpotHistoricalTrades //历史成交记录
	SpotAggTrades        //近期成交(归集)
	SpotAvgPrice         //当前平均价格
	SpotUiKlines         //UIK线数据
	SpotTicker24hr       //24hr 价格变动情况
	SpotTickerBookTicker //当前最优挂单
	SpotTicker           //滚动窗口价格变动统计

	//Ws账户推送相关
	SpotUserDataStreamPost   //(现货账户)生成 Listen Key (USER_STREAM)
	SpotUserDataStreamPut    //(现货账户)延长 Listen Key 有效期 (USER_STREAM)
	SpotUserDataStreamDelete //(现货账户)关闭 Listen Key (USER_STREAM)

	SpotMarginUserDataStreamPost   //(杠杆账户)生成 Listen Key (USER_STREAM)
	SpotMarginUserDataStreamPut    //(杠杆账户)延长 Listen Key 有效期 (USER_STREAM)
	SpotMarginUserDataStreamDelete //(杠杆账户)关闭 Listen Key (USER_STREAM)

	SpotMarginIsolatedUserDataStreamPost   //(逐仓杠杆账户)生成 Listen Key (USER_STREAM)
	SpotMarginIsolatedUserDataStreamPut    //(逐仓杠杆账户)延长 Listen Key 有效期 (USER_STREAM)
	SpotMarginIsolatedUserDataStreamDelete //(逐仓杠杆账户)关闭 Listen Key (USER_STREAM)
)

var SpotApiMap = map[SpotApi]string{
	//子母账户接口
	//GET
	SpotSubAccountList:                     "/sapi/v1/sub-account/list",                        //GET接口 查询子账户列表(适用主账户)
	SpotSubAccountUniversalTransferHistory: "/sapi/v1/sub-account/universalTransfer",           //GET接口 子母账户万能划转历史查询 (适用主账户)
	SpotSubAccountAssets:                   "/sapi/v4/sub-account/assets",                      //GET接口 查询子账户资产(适用主账户)(USER_DATA)
	SpotSubAccountFuturesAccount:           "/sapi/v2/sub-account/futures/account",             //GET接口 查询子账户Futures账户详情V2 (适用主账户)
	SpotSubAccountApiIpRestriction:         "/sapi/v1/sub-account/subAccountApi/ipRestriction", //GET接口 查询子账户API Key IP白名单 (适用母账户)
	SpotSubAccountTransferSubUserHistory:   "/sapi/v1/sub-account/transfer/subUserHistory",     //GET 查询子账户划转历史 (仅适用子账户)

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
	SpotMarginOrderDelete:      "/sapi/v1/margin/order",             //DELETE 撤销杠杆账户订单 (TRADE)
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
	SpotOrderDelete:             "/api/v3/order",                     //DELETE接口 撤销订单 (TRADE)
	SpotAssetTransferPost:       "/sapi/v1/asset/transfer",           //用户万向划转 (USER_DATA)
	SpotAssetTransferGet:        "/sapi/v1/asset/transfer",           //查询用户万向划转历史 (USER_DATA)
	SpotAssetTradeFee:           "/sapi/v1/asset/tradeFee",           //查询用户交易手续费率 (USER_DATA)

	//现货账户接口
	SpotMyTrades: "/api/v3/myTrades", //GET接口 账户成交历史 (USER_DATA)

	//通用接口
	SpotPing:         "/api/v3/ping",         //GET接口 测试连通性
	SpotServerTime:   "/api/v3/time",         //GET接口 获取服务器时间
	SpotExchangeInfo: "/api/v3/exchangeInfo", //GET接口 获取交易规范

	//行情接口
	SpotTickerPrice:      "/api/v3/ticker/price",      //GET接口 获取交易对最新价格
	SpotKlines:           "/api/v3/klines",            //GET接口 获取K线数据
	SpotDepth:            "/api/v3/depth",             //GET接口 获取深度信息
	SpotTrades:           "/api/v3/trades",            //GET接口 近期成交列表
	SpotHistoricalTrades: "/api/v3/historicalTrades",  //GET接口 历史成交记录
	SpotAggTrades:        "/api/v3/aggTrades",         //GET接口 近期成交(归集)
	SpotAvgPrice:         "/api/v3/avgPrice",          //GET接口 当前平均价格
	SpotUiKlines:         "/api/v3/uiKlines",          //GET接口 UIK线数据
	SpotTicker24hr:       "/api/v3/ticker/24hr",       //GET接口 24hr 价格变动情况
	SpotTickerBookTicker: "/api/v3/ticker/bookTicker", //GET接口 当前最优挂单
	SpotTicker:           "/api/v3/ticker",            //GET接口 滚动窗口价格变动统计

	//Ws账户推送相关
	SpotUserDataStreamPost:   "/api/v3/userDataStream", //POST接口 (现货账户)生成 Listen Key (USER_STREAM)
	SpotUserDataStreamPut:    "/api/v3/userDataStream", //PUT接口 (现货账户)延长 Listen Key 有效期 (USER_STREAM)
	SpotUserDataStreamDelete: "/api/v3/userDataStream", //DELETE接口 (现货账户)关闭 Listen Key (USER_STREAM)

	SpotMarginUserDataStreamPost:   "/sapi/v1/userDataStream", //POST接口 (杠杆账户)生成 Listen Key (USER_STREAM)
	SpotMarginUserDataStreamPut:    "/sapi/v1/userDataStream", //PUT接口 (杠杆账户)延长 Listen Key 有效期 (USER_STREAM)
	SpotMarginUserDataStreamDelete: "/sapi/v1/userDataStream", //DELETE接口 (杠杆账户)关闭 Listen Key (USER_STREAM)

	SpotMarginIsolatedUserDataStreamPost:   "/sapi/v1/userDataStream/isolated", //POST接口 (逐仓杠杆账户)生成 Listen Key (USER_STREAM)
	SpotMarginIsolatedUserDataStreamPut:    "/sapi/v1/userDataStream/isolated", //PUT接口 (逐仓杠杆账户)延长 Listen Key 有效期 (USER_STREAM)
	SpotMarginIsolatedUserDataStreamDelete: "/sapi/v1/userDataStream/isolated", //DELETE接口 (逐仓杠杆账户)关闭 Listen Key (USER_STREAM)
}

// ================以下为子母账户GET接口
// binance SPOT子母账户接口  SubAccountList rest查询子账户列表
func (client *SpotRestClient) NewSubAccountList() *SpotSubAccountListApi {
	return &SpotSubAccountListApi{
		client: client,
		req:    &SpotSubAccountListReq{},
	}
}
func (api *SpotSubAccountListApi) Do() (*SubAccountListRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotSubAccountList], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SubAccountListRes](api.client.c, url, GET)
}

// binance SPOT子母账户接口  SubAccountUniversalTransferHistory rest查询子母账户万能划转历史查询
func (client *SpotRestClient) NewSubAccountUniversalTransferHistory() *SpotSubAccountUniversalTransferHistoryApi {
	return &SpotSubAccountUniversalTransferHistoryApi{
		client: client,
		req:    &SpotSubAccountUniversalTransferHistoryReq{},
	}
}
func (api *SpotSubAccountUniversalTransferHistoryApi) Do() (*SubAccountUniversalTransferHistoryRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotSubAccountUniversalTransferHistory], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SubAccountUniversalTransferHistoryRes](api.client.c, url, GET)
}

// binance SPOT子母账户接口  SubAccountAssets rest查询子账户资产 (适用主账户)
func (client *SpotRestClient) NewSubAccountAssets() *SpotSubAccountAssetsApi {
	return &SpotSubAccountAssetsApi{
		client: client,
		req:    &SpotSubAccountAssetsReq{},
	}
}
func (api *SpotSubAccountAssetsApi) Do() (*SubAccountAssetsRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotSubAccountAssets], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SubAccountAssetsRes](api.client.c, url, GET)
}

// binance SPOT子母账户接口  SubAccountAssets rest查询子账户Futures账户详情V2 (适用主账户)
func (client *SpotRestClient) NewSubAccountFuturesAccount() *SpotSubAccountFuturesAccountApi {
	return &SpotSubAccountFuturesAccountApi{
		client: client,
		req:    &SpotSubAccountFuturesAccountReq{},
	}
}
func (api *SpotSubAccountFuturesAccountApi) Do() (*SubAccountFuturesAccountRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotSubAccountFuturesAccount], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SubAccountFuturesAccountRes](api.client.c, url, GET)
}

// binance SPOT子母账户接口  SubAccountTransferSubUserHistory rest查询子账户划转历史 (仅适用子账户)
func (client *SpotRestClient) NewSubAccountTransferSubUserHistory() *SpotSubAccountTransferSubUserHistoryApi {
	return &SpotSubAccountTransferSubUserHistoryApi{
		client: client,
		req:    &SpotSubAccountTransferSubUserHistoryReq{},
	}
}
func (api *SpotSubAccountTransferSubUserHistoryApi) Do() (*SpotSubAccountTransferSubUserHistoryRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotSubAccountTransferSubUserHistory], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SpotSubAccountTransferSubUserHistoryRes](api.client.c, url, GET)
}

// binance SPOT现货账户和交易接口  SpotAccount rest账户信息 (USER_DATA)
func (client *SpotRestClient) NewSpotAccount() *SpotAccountApi {
	return &SpotAccountApi{
		client: client,
		req:    &SpotAccountReq{},
	}
}
func (api *SpotAccountApi) Do() (*SpotAccountRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotAccount], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SpotAccountRes](api.client.c, url, GET)
}

// binance 查询子账户API Key IP白名单 (适用母账户)  SpotSubAccountApiIpRestriction rest查询子账户API Key IP白名单 (适用母账户)
func (client *SpotRestClient) NewSpotSubAccountApiIpRestriction() *SpotSubAccountApiIpRestrictionApi {
	return &SpotSubAccountApiIpRestrictionApi{
		client: client,
		req:    &SpotSubAccountApiIpRestrictionReq{},
	}
}
func (api *SpotSubAccountApiIpRestrictionApi) Do() (*SpotSubAccountApiIpRestrictionRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotSubAccountApiIpRestriction], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SpotSubAccountApiIpRestrictionRes](api.client.c, url, GET)
}

// ================以下为子母账户POST接口
// binance SPOT子母账户接口  SubAccountVirtualSubAccount rest创建虚拟子账户
func (client *SpotRestClient) NewSubAccountVirtualSubAccount() *SpotSubAccountVirtualSubAccountApi {
	return &SpotSubAccountVirtualSubAccountApi{
		client: client,
		req:    &SpotSubAccountVirtualSubAccountReq{},
	}
}
func (api *SpotSubAccountVirtualSubAccountApi) Do() (*SubAccountVirtualSubAccountRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotSubAccountVirtualSubAccount], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SubAccountVirtualSubAccountRes](api.client.c, url, POST)
}

// binance SPOT子母账户接口  SubAccountUniversalTransfer rest子母账户万能划转
func (client *SpotRestClient) NewSubAccountUniversalTransfer() *SpotSubAccountUniversalTransferApi {
	return &SpotSubAccountUniversalTransferApi{
		client: client,
		req:    &SpotSubAccountUniversalTransferReq{},
	}
}
func (api *SpotSubAccountUniversalTransferApi) Do() (*SubAccountUniversalTransferRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotSubAccountUniversalTransfer], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SubAccountUniversalTransferRes](api.client.c, url, POST)
}

// binance SPOT子母账户接口  SubAccountFuturesEnable rest为子账户开通Futures (适用主账户)
func (client *SpotRestClient) NewSubAccountFuturesEnable() *SpotSubAccountFuturesEnableApi {
	return &SpotSubAccountFuturesEnableApi{
		client: client,
		req:    &SpotSubAccountFuturesEnableReq{},
	}
}
func (api *SpotSubAccountFuturesEnableApi) Do() (*SubAccountFuturesEnableRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotSubAccountFuturesEnable], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SubAccountFuturesEnableRes](api.client.c, url, POST)
}

// ================以下为杠杆账户GET接口
// binance SPOT杠杆接口  MarginAllPairs rest获取所有全仓杠杆交易对(MARKET_DATA)
func (client *SpotRestClient) NewMarginAllPairs() *SpotMarginAllPairsApi {
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
func (client *SpotRestClient) NewMarginIsolatedAllPairs() *SpotMarginIsolatedAllPairsApi {
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
func (client *SpotRestClient) NewMarginAccount() *SpotMarginAccountApi {
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
func (client *SpotRestClient) NewMarginIsolatedAccount() *SpotMarginIsolatedAccountApi {
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
func (client *SpotRestClient) NewMarginMaxBorrowable() *SpotMarginMaxBorrowableApi {
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
func (client *SpotRestClient) NewMarginMaxTransferable() *SpotMarginMaxTransferableApi {
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
func (client *SpotRestClient) NewMarginInterestHistory() *SpotMarginInterestHistoryApi {
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
func (client *SpotRestClient) NewMarginOrderGet() *SpotMarginOrderGetApi {
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

// binance SPOT杠杆接口  MarginOrderDelete rest撤销杠杆账户订单 (TRADE)
func (client *SpotRestClient) NewMarginOrderDelete() *SpotMarginOrderDeleteApi {
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
func (client *SpotRestClient) NewMarginAllOrders() *SpotMarginAllOrdersApi {
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
func (client *SpotRestClient) NewMarginOpenOrders() *SpotMarginOpenOrdersApi {
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

// ================以下为杠杆账户POST接口
// binance SPOT杠杆接口  MarginTransfer rest全仓杠杆账户划转 (MARGIN)
func (client *SpotRestClient) NewMarginTransfer() *SpotMarginTransferApi {
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
func (client *SpotRestClient) NewMarginIsolatedTransfer() *SpotMarginIsolatedTransferApi {
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
func (client *SpotRestClient) NewMarginLoan() *SpotMarginLoanApi {
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
func (client *SpotRestClient) NewMarginRepay() *SpotMarginRepayApi {
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

// 下单接口
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
	log.Warn(url)
	return binanceCallApiWithSecret[SpotOrderPostRes](api.client.c, url, POST)
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

// 钱包接口
// binance SPOT钱包接口 账户API交易状态(USER_DATA)
func (client *SpotRestClient) NewAccountApiTradingStatus() *SpotAccountApiTradingStatusApi {
	return &SpotAccountApiTradingStatusApi{
		client: client,
		req:    &SpotAccountApiTradingStatusReq{},
	}
}
func (api *SpotAccountApiTradingStatusApi) Do() (*SpotAccountApiTradingStatusRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotAccountApiTradingStatus], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SpotAccountApiTradingStatusRes](api.client.c, url, GET)
}

// 通用接口
// binance SPOT ping rest测试连通性 (NONE)
func (client *SpotRestClient) NewPing() *SpotPingApi {
	return &SpotPingApi{
		client: client,
		req:    &SpotPingReq{},
	}
}
func (api *SpotPingApi) Do() (*SpotPingRes, error) {
	url := binanceHandlerRequestApi(SPOT, api.req, SpotApiMap[SpotPing])
	return binanceCallApiWithSecret[SpotPingRes](api.client.c, url, GET)
}

// binance SPOT serverTime rest服务器时间 (NONE)
func (client *SpotRestClient) NewServerTime() *SpotServerTimeApi {
	return &SpotServerTimeApi{
		client: client,
		req:    &SpotServerTimeReq{},
	}
}
func (api *SpotServerTimeApi) Do() (*SpotServerTimeRes, error) {
	url := binanceHandlerRequestApi(SPOT, api.req, SpotApiMap[SpotServerTime])
	return binanceCallApiWithSecret[SpotServerTimeRes](api.client.c, url, GET)
}

// binance SPOT exchangeInfo rest交易规范
func (client *SpotRestClient) NewExchangeInfo() *SpotExchangeInfoApi {
	return &SpotExchangeInfoApi{
		client: client,
		req:    &SpotExchangeInfoReq{},
	}
}
func (api *SpotExchangeInfoApi) Do() (*SpotExchangeInfoRes, error) {
	url := binanceHandlerRequestApi(SPOT, api.req, SpotApiMap[SpotExchangeInfo])
	return binanceCallApiWithSecret[SpotExchangeInfoRes](api.client.c, url, GET)
}

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

// binance SPOT spotAssetTransferPost rest用户万向划转 (USER_DATA)
func (client *SpotRestClient) NewSpotAssetTransferPost() *SpotAssetTransferPostApi {
	return &SpotAssetTransferPostApi{
		client: client,
		req:    &SpotAssetTransferPostReq{},
	}
}
func (api *SpotAssetTransferPostApi) Do() (*SpotAssetTransferPostRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotAssetTransferPost], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SpotAssetTransferPostRes](api.client.c, url, POST)
}

// binance SPOT spotAssetTransferGet rest查询用户万向划转历史 (USER_DATA)
func (client *SpotRestClient) NewSpotAssetTransferGet() *SpotAssetTransferGetApi {
	return &SpotAssetTransferGetApi{
		client: client,
		req:    &SpotAssetTransferGetReq{},
	}
}
func (api *SpotAssetTransferGetApi) Do() (*SpotAssetTransferGetRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotAssetTransferGet], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SpotAssetTransferGetRes](api.client.c, url, GET)
}

// binance SPOT spotAssetTradeFee rest查询用户交易手续费率 (USER_DATA)
func (client *SpotRestClient) NewSpotAssetTradeFee() *SpotAssetTradeFeeApi {
	return &SpotAssetTradeFeeApi{
		client: client,
		req:    &SpotAssetTradeFeeReq{},
	}
}
func (api *SpotAssetTradeFeeApi) Do() (*SpotAssetTradeFeeRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotAssetTradeFee], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SpotAssetTradeFeeRes](api.client.c, url, GET)
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

// binance SPOT spotTickerPrice rest价格 (NONE)
func (client *SpotRestClient) NewSpotTickerPrice() *SpotTickerPriceApi {
	return &SpotTickerPriceApi{
		client: client,
		req:    &SpotTickerPriceReq{},
	}
}
func (api *SpotTickerPriceApi) Do() (*SpotTickerPriceRes, error) {
	url := binanceHandlerRequestApi(SPOT, api.req, SpotApiMap[SpotTickerPrice])
	if api.req.Symbol != nil && *api.req.Symbol != "" {
		res, err := binanceCallApiWithSecret[SpotTickerPriceResRow](api.client.c, url, GET)
		if err != nil {
			return nil, err
		}
		return &SpotTickerPriceRes{*res}, nil
	} else {
		return binanceCallApiWithSecret[SpotTickerPriceRes](api.client.c, url, GET)
	}

}

// binance SPOT spotKlines restK线数据 (NONE)
func (client *SpotRestClient) NewSpotKlines() *SpotKlinesApi {
	return &SpotKlinesApi{
		client: client,
		req:    &SpotKlinesReq{},
	}
}
func (api *SpotKlinesApi) Do() (*KlinesRes, error) {
	url := binanceHandlerRequestApi(SPOT, api.req, SpotApiMap[SpotKlines])
	res, err := binanceCallApiWithSecret[KlinesMiddle](api.client.c, url, GET)
	if err != nil {
		return nil, err
	}
	res2 := res.ConvertToRes()
	return &res2, nil
}

// binance SPOT spotDepth rest深度信息 (NONE)
func (client *SpotRestClient) NewSpotDepth() *SpotDepthApi {
	return &SpotDepthApi{
		client: client,
		req:    &SpotDepthReq{},
	}
}
func (api *SpotDepthApi) Do() (*SpotDepthRes, error) {
	url := binanceHandlerRequestApi(SPOT, api.req, SpotApiMap[SpotDepth])
	res, err := binanceCallApiWithSecret[SpotDepthResMiddle](api.client.c, url, GET)
	if err != nil {
		return nil, err
	}
	return res.ConvertToRes(), nil
}

// binance SPOT SpotTrades rest最近成交 (NONE)
func (client *SpotRestClient) NewSpotTrades() *SpotTradesApi {
	return &SpotTradesApi{
		client: client,
		req:    &SpotTradesReq{},
	}
}
func (api *SpotTradesApi) Do() (*SpotTradesRes, error) {
	url := binanceHandlerRequestApi(SPOT, api.req, SpotApiMap[SpotTrades])
	return binanceCallApiWithSecret[SpotTradesRes](api.client.c, url, GET)
}

// binance SPOT spotHistoricalTrades rest历史成交 (NONE)
func (client *SpotRestClient) NewSpotHistoricalTrades() *SpotHistoricalTradesApi {
	return &SpotHistoricalTradesApi{
		client: client,
		req:    &SpotHistoricalTradesReq{},
	}
}
func (api *SpotHistoricalTradesApi) Do() (*SpotHistoricalTradesRes, error) {
	url := binanceHandlerRequestApi(SPOT, api.req, SpotApiMap[SpotHistoricalTrades])
	return binanceCallApiWithSecret[SpotHistoricalTradesRes](api.client.c, url, GET)
}

// binance SPOT spotAggTrades rest近期成交(归集)(NONE)
func (client *SpotRestClient) NewSpotAggTrades() *SpotAggTradesApi {
	return &SpotAggTradesApi{
		client: client,
		req:    &SpotAggTradesReq{},
	}
}
func (api *SpotAggTradesApi) Do() (*SpotAggTradesRes, error) {
	url := binanceHandlerRequestApi(SPOT, api.req, SpotApiMap[SpotAggTrades])
	return binanceCallApiWithSecret[SpotAggTradesRes](api.client.c, url, GET)
}

// binance SPOT spotAvgPrice rest当前平均价格 (NONE)
func (client *SpotRestClient) NewSpotAvgPrice() *SpotAvgPriceApi {
	return &SpotAvgPriceApi{
		client: client,
		req:    &SpotAvgPriceReq{},
	}
}
func (api *SpotAvgPriceApi) Do() (*SpotAvgPriceRes, error) {
	url := binanceHandlerRequestApi(SPOT, api.req, SpotApiMap[SpotAvgPrice])
	return binanceCallApiWithSecret[SpotAvgPriceRes](api.client.c, url, GET)
}

// binance SPOT spotUiKlines restUI K线数据 (NONE)
func (client *SpotRestClient) NewSpotUiKlines() *SpotUiKlinesApi {
	return &SpotUiKlinesApi{
		client: client,
		req:    &SpotKlinesReq{},
	}
}
func (api *SpotUiKlinesApi) Do() (*KlinesRes, error) {
	url := binanceHandlerRequestApi(SPOT, api.req, SpotApiMap[SpotKlines])
	res, err := binanceCallApiWithSecret[KlinesMiddle](api.client.c, url, GET)
	if err != nil {
		return nil, err
	}
	res2 := res.ConvertToRes()
	return &res2, nil
}

// binance SPOT spotTicker24hr rest24hr价格变动情况 (NONE)
func (client *SpotRestClient) NewSpotTicker24hr() *SpotTicker24hrApi {
	return &SpotTicker24hrApi{
		client: client,
		req:    &SpotTicker24hrReq{},
	}
}
func (api *SpotTicker24hrApi) Do() (*SpotTicker24hrRes, error) {
	url := binanceHandlerRequestApi(SPOT, api.req, SpotApiMap[SpotTicker24hr])
	if api.req.Symbol != nil && *api.req.Symbol != "" {
		res, err := binanceCallApiWithSecret[SpotTicker24hrResRow](api.client.c, url, GET)
		if err != nil {
			return nil, err
		}
		return &SpotTicker24hrRes{*res}, nil
	} else {
		return binanceCallApiWithSecret[SpotTicker24hrRes](api.client.c, url, GET)
	}
}

// binance SPOT spotTickerBookTicker rest当前最优挂单 (NONE)
func (client *SpotRestClient) NewSpotTickerBookTicker() *SpotTickerBookTickerApi {
	return &SpotTickerBookTickerApi{
		client: client,
		req:    &SpotTickerBookTickerReq{},
	}
}
func (api *SpotTickerBookTickerApi) Do() (*SpotTickerBookTickerRes, error) {
	url := binanceHandlerRequestApi(SPOT, api.req, SpotApiMap[SpotTickerBookTicker])
	if api.req.Symbol != nil && *api.req.Symbol != "" {
		res, err := binanceCallApiWithSecret[SpotTickerBookTickerResRow](api.client.c, url, GET)
		if err != nil {
			return nil, err
		}
		return &SpotTickerBookTickerRes{*res}, nil
	} else {
		return binanceCallApiWithSecret[SpotTickerBookTickerRes](api.client.c, url, GET)
	}
}

// binance SPOT spotTicker rest滚动窗口价格变动统计
func (client *SpotRestClient) NewSpotTicker() *SpotTickerApi {
	return &SpotTickerApi{
		client: client,
		req:    &SpotTickerReq{},
	}
}
func (api *SpotTickerApi) Do() (*SpotTickerRes, error) {
	url := binanceHandlerRequestApi(SPOT, api.req, SpotApiMap[SpotTicker])
	if api.req.Symbol != nil && *api.req.Symbol != "" {
		res, err := binanceCallApiWithSecret[SpotTickerResRow](api.client.c, url, GET)
		if err != nil {
			return nil, err
		}
		return &SpotTickerRes{*res}, nil
	} else {
		return binanceCallApiWithSecret[SpotTickerRes](api.client.c, url, GET)
	}
}

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
