package mybinanceapi

import "time"

type FutureApi int

const (
	//钱包接口
	FutureAccount FutureApi = iota
	FuturePositionSideDualGet
	FutureMultiAssetsMarginGet
	FuturePositionSideDualPost
	FutureMultiAssetsMarginPost
	FutureLeverage
	FutureMarginType
	FutureLeverageBracket
	FutureOpenOrders
	FutureAllOrders
	FutureOrderPost
	FutureOrderGet
	FutureOrderDelete
	FutureBatchOrdersDelete

	FutureUserTrades

	FutureCommissionRate

	//通用接口
	FuturePing
	FutureServerTime
	FutureExchangeInfo

	//行情接口
	FutureKlines
)

var FutureApiMap = map[FutureApi]string{

	//U合约 GET接口
	FutureAccount:              "/fapi/v2/account",           //GET接口 账户信息V2 (USER_DATA)
	FuturePositionSideDualGet:  "/fapi/v1/positionSide/dual", //GET接口 (HMAC SHA256)查询持仓模式(USER_DATA)
	FutureMultiAssetsMarginGet: "/fapi/v1/multiAssetsMargin", //GET接口 (HMAC SHA256)查询联合保证金模式(USER_DATA)
	//U合约 POST接口
	FuturePositionSideDualPost:  "/fapi/v1/positionSide/dual", //POST接口 (HMAC SHA256)更改持仓模式(TRADE)
	FutureMultiAssetsMarginPost: "/fapi/v1/multiAssetsMargin", //POST接口 (HMAC SHA256)更改联合保证金模式(TRADE)

	FutureLeverage:   "/fapi/v1/leverage",   //POST接口 (HMAC SHA256)调整开仓杠杆 (TRADE)
	FutureMarginType: "/fapi/v1/marginType", //POST接口 (HMAC SHA256)变换逐全仓模式 (TRADE)

	FutureLeverageBracket: "/fapi/v1/leverageBracket", //GET接口 杠杆分层标准 (USER_DATA)

	FutureOpenOrders: "/fapi/v1/openOrders", //GET接口 (HMAC SHA256)查询当前挂单 (USER_DATA)
	FutureAllOrders:  "/fapi/v1/allOrders",  //GET接口 (HMAC SHA256)查询所有订单 (USER_DATA)

	FutureOrderPost:         "/fapi/v1/order",       //POST接口 (HMAC SHA256)下单 (TRADE)
	FutureOrderGet:          "/fapi/v1/order",       //GET接口 (HMAC SHA256)查询订单 (USER_DATA)
	FutureOrderDelete:       "/fapi/v1/order",       //DELETE接口 (HMAC SHA256)撤销订单 (TRADE)
	FutureBatchOrdersDelete: "/fapi/v1/batchOrders", //DELETE接口 (HMAC SHA256)批量撤销订单 (TRADE)

	FutureUserTrades: "/fapi/v1/userTrades", //GET接口 (HMAC SHA256)账户成交历史 (USER_DATA)

	FutureCommissionRate: "/fapi/v1/commissionRate", //GET接口 (HMAC SHA256)查询用户当前的手续费率

	//通用接口
	FuturePing:         "/fapi/v1/ping",         //GET接口 测试服务器连通性
	FutureServerTime:   "/fapi/v1/time",         //GET接口 获取服务器时间
	FutureExchangeInfo: "/fapi/v1/exchangeInfo", //GET接口 交易规则和交易对信息

	FutureKlines: "/fapi/v1/klines", //GET接口 K线数据
}

//=======GET接口
// binance FUTURE  FutureAccount rest账户信息 (USER_DATA)
func (client *FutureRestClient) NewFutureAccount() *FutureAccountApi {
	return &FutureAccountApi{
		FutureRestClient: *client,
		req:              &FutureAccountReq{},
	}
}
func (api *FutureAccountApi) Do() (*FutureAccountRes, error) {
	api.Timestamp(time.Now().UnixMilli())
	url := binanceHandlerRequestApiWithSecret(FUTURE, api.req, FutureApiMap[FutureAccount], api.c.ApiSecret)
	return binanceCallApiWithSecret[FutureAccountRes](api.FutureRestClient.c, url, GET)
}

// binance FUTURE  FuturePositionSideDualGet rest查询持仓模式 (USER_DATA)
func (client *FutureRestClient) NewFuturePositionSideDualGet() *FuturePositionSideDualGetApi {
	return &FuturePositionSideDualGetApi{
		FutureRestClient: *client,
		req:              &FuturePositionSideDualGetReq{},
	}
}
func (api *FuturePositionSideDualGetApi) Do() (*FuturePositionSideDualGetRes, error) {
	api.Timestamp(time.Now().UnixMilli())
	url := binanceHandlerRequestApiWithSecret(FUTURE, api.req, FutureApiMap[FuturePositionSideDualGet], api.c.ApiSecret)
	return binanceCallApiWithSecret[FuturePositionSideDualGetRes](api.FutureRestClient.c, url, GET)
}

// binance FUTURE  FutureMultiAssetsMarginGet rest查询联合保证金模式 (USER_DATA)
func (client *FutureRestClient) NewFutureMultiAssetsMarginGet() *FutureMultiAssetsMarginGetApi {
	return &FutureMultiAssetsMarginGetApi{
		FutureRestClient: *client,
		req:              &FutureMultiAssetsMarginGetReq{},
	}
}
func (api *FutureMultiAssetsMarginGetApi) Do() (*FutureMultiAssetsMarginGetRes, error) {
	api.Timestamp(time.Now().UnixMilli())
	url := binanceHandlerRequestApiWithSecret(FUTURE, api.req, FutureApiMap[FutureMultiAssetsMarginGet], api.c.ApiSecret)
	return binanceCallApiWithSecret[FutureMultiAssetsMarginGetRes](api.FutureRestClient.c, url, GET)
}

// binance FUTURE  FutureLeverageBracket rest杠杆分层标准 (USER_DATA)
func (client *FutureRestClient) NewFutureLeverageBracket() *FutureLeverageBracketApi {
	return &FutureLeverageBracketApi{
		FutureRestClient: *client,
		req:              &FutureLeverageBracketReq{},
	}
}
func (api *FutureLeverageBracketApi) Do() (*FutureLeverageBracketRes, error) {
	api.Timestamp(time.Now().UnixMilli())
	url := binanceHandlerRequestApiWithSecret(FUTURE, api.req, FutureApiMap[FutureLeverageBracket], api.c.ApiSecret)
	return binanceCallApiWithSecret[FutureLeverageBracketRes](api.FutureRestClient.c, url, GET)
}

//========POST接口
// binance FUTURE  FuturePositionSideDualPost rest更改持仓模式 (TRADE)
func (client *FutureRestClient) NewFuturePositionSideDualPost() *FuturePositionSideDualPostApi {
	return &FuturePositionSideDualPostApi{
		FutureRestClient: *client,
		req:              &FuturePositionSideDualPostReq{},
	}
}
func (api *FuturePositionSideDualPostApi) Do() (*FuturePositionSideDualPostRes, error) {
	api.Timestamp(time.Now().UnixMilli())
	url := binanceHandlerRequestApiWithSecret(FUTURE, api.req, FutureApiMap[FuturePositionSideDualPost], api.c.ApiSecret)
	return binanceCallApiWithSecret[FuturePositionSideDualPostRes](api.FutureRestClient.c, url, POST)
}

// binance FUTURE  FutureMultiAssetsMarginPost rest更改联合保证金模式 (TRADE)
func (client *FutureRestClient) NewFutureMultiAssetsMarginPost() *FutureMultiAssetsMarginPostApi {
	return &FutureMultiAssetsMarginPostApi{
		FutureRestClient: *client,
		req:              &FutureMultiAssetsMarginPostReq{},
	}
}
func (api *FutureMultiAssetsMarginPostApi) Do() (*FutureMultiAssetsMarginPostRes, error) {
	api.Timestamp(time.Now().UnixMilli())
	url := binanceHandlerRequestApiWithSecret(FUTURE, api.req, FutureApiMap[FutureMultiAssetsMarginPost], api.c.ApiSecret)
	return binanceCallApiWithSecret[FutureMultiAssetsMarginPostRes](api.FutureRestClient.c, url, POST)
}

// binance FUTURE  FutureLeverage rest调整开仓杠杆 (TRADE)
func (client *FutureRestClient) NewFutureLeverage() *FutureLeverageApi {
	return &FutureLeverageApi{
		FutureRestClient: *client,
		req:              &FutureLeverageReq{},
	}
}
func (api *FutureLeverageApi) Do() (*FutureLeverageRes, error) {
	api.Timestamp(time.Now().UnixMilli())
	url := binanceHandlerRequestApiWithSecret(FUTURE, api.req, FutureApiMap[FutureLeverage], api.c.ApiSecret)
	return binanceCallApiWithSecret[FutureLeverageRes](api.FutureRestClient.c, url, POST)
}

// binance FUTURE  FutureMarginType rest变换逐全仓模式 (TRADE)
func (client *FutureRestClient) NewFutureMarginType() *FutureMarginTypeApi {
	return &FutureMarginTypeApi{
		FutureRestClient: *client,
		req:              &FutureMarginTypeReq{},
	}
}
func (api *FutureMarginTypeApi) Do() (*FutureMarginTypeRes, error) {
	api.Timestamp(time.Now().UnixMilli())
	url := binanceHandlerRequestApiWithSecret(FUTURE, api.req, FutureApiMap[FutureMarginType], api.c.ApiSecret)
	return binanceCallApiWithSecret[FutureMarginTypeRes](api.FutureRestClient.c, url, POST)
}

// binance FUTURE  FutureCommissionRate rest查询用户当前的手续费率 (USER_DATA)
func (client *FutureRestClient) NewFutureCommissionRate() *FutureCommissionRateApi {
	return &FutureCommissionRateApi{
		FutureRestClient: *client,
		req:              &FutureCommissionRateReq{},
	}
}
func (api *FutureCommissionRateApi) Do() (*FutureCommissionRateRes, error) {
	api.Timestamp(time.Now().UnixMilli())
	url := binanceHandlerRequestApiWithSecret(FUTURE, api.req, FutureApiMap[FutureCommissionRate], api.c.ApiSecret)
	return binanceCallApiWithSecret[FutureCommissionRateRes](api.FutureRestClient.c, url, GET)
}

//通用接口
// binance FUTURE FuturePing rest测试服务器连通性 (NONE)
func (client *FutureRestClient) NewPing() *FuturePingApi {
	return &FuturePingApi{
		FutureRestClient: *client,
		req:              &FuturePingReq{},
	}
}

func (api *FuturePingApi) Do() (*FuturePingRes, error) {
	url := binanceHandlerRequestApi(FUTURE, api.req, FutureApiMap[FuturePing])
	return binanceCallApiWithSecret[FuturePingRes](api.FutureRestClient.c, url, GET)
}

// binance FUTURE FutureTime rest获取服务器时间 (NONE)
func (client *FutureRestClient) NewServerTime() *FutureServerTimeApi {
	return &FutureServerTimeApi{
		FutureRestClient: *client,
		req:              &FutureServerTimeReq{},
	}
}

func (api *FutureServerTimeApi) Do() (*FutureTimeRes, error) {
	url := binanceHandlerRequestApi(FUTURE, api.req, FutureApiMap[FutureServerTime])
	return binanceCallApiWithSecret[FutureTimeRes](api.FutureRestClient.c, url, GET)
}

// binance FUTURE FutureExchangeInfo rest交易规范信息和交易对信息 (NONE)
func (client *FutureRestClient) NewExchangeInfo() *FutureExchangeInfoApi {
	return &FutureExchangeInfoApi{
		FutureRestClient: *client,
		req:              &FutureExchangeInfoReq{},
	}
}

func (api *FutureExchangeInfoApi) Do() (*FutureExchangeInfoRes, error) {
	url := binanceHandlerRequestApi(FUTURE, api.req, FutureApiMap[FutureExchangeInfo])
	return binanceCallApiWithSecret[FutureExchangeInfoRes](api.FutureRestClient.c, url, GET)
}

// binance FUTURE FutureOpenOrders rest查询当前挂单 (USER_DATA)
func (client *FutureRestClient) NewOpenOrders() *FutureOpenOrdersApi {
	return &FutureOpenOrdersApi{
		FutureRestClient: *client,
		req:              &FutureOpenOrdersReq{},
	}
}

func (api *FutureOpenOrdersApi) Do() (*FutureOpenOrdersRes, error) {
	api.Timestamp(time.Now().UnixMilli())
	url := binanceHandlerRequestApiWithSecret(FUTURE, api.req, FutureApiMap[FutureOpenOrders], api.c.ApiSecret)
	return binanceCallApiWithSecret[FutureOpenOrdersRes](api.FutureRestClient.c, url, GET)
}

// binance FUTURE FutureAllOrders rest查询所有订单 (USER_DATA)
func (client *FutureRestClient) NewAllOrders() *FutureAllOrdersApi {
	return &FutureAllOrdersApi{
		FutureRestClient: *client,
		req:              &FutureAllOrdersReq{},
	}
}

func (api *FutureAllOrdersApi) Do() (*FutureAllOrdersRes, error) {
	api.Timestamp(time.Now().UnixMilli())
	url := binanceHandlerRequestApiWithSecret(FUTURE, api.req, FutureApiMap[FutureAllOrders], api.c.ApiSecret)
	return binanceCallApiWithSecret[FutureAllOrdersRes](api.FutureRestClient.c, url, GET)
}

// binance FUTURE FutureOrderPost rest下单 (TRADE)
func (client *FutureRestClient) NewFutureOrderPost() *FutureOrderPostApi {
	return &FutureOrderPostApi{
		FutureRestClient: *client,
		req:              &FutureOrderPostReq{},
	}
}

func (api *FutureOrderPostApi) Do() (*FutureOrderPostRes, error) {
	api.Timestamp(time.Now().UnixMilli())
	url := binanceHandlerRequestApiWithSecret(FUTURE, api.req, FutureApiMap[FutureOrderPost], api.c.ApiSecret)
	return binanceCallApiWithSecret[FutureOrderPostRes](api.FutureRestClient.c, url, POST)
}

// binance FUTURE FutureOrderGet  rest查询订单 (USER_DATA)
func (client *FutureRestClient) NewFutureOrderGet() *FutureOrderGetApi {
	return &FutureOrderGetApi{
		FutureRestClient: *client,
		req:              &FutureOrderGetReq{},
	}
}

func (api *FutureOrderGetApi) Do() (*FutureOrderGetRes, error) {
	api.Timestamp(time.Now().UnixMilli())
	url := binanceHandlerRequestApiWithSecret(FUTURE, api.req, FutureApiMap[FutureOrderGet], api.c.ApiSecret)
	return binanceCallApiWithSecret[FutureOrderGetRes](api.FutureRestClient.c, url, GET)
}

// binance FUTURE FutureOrderDelete rest撤销订单 (TRADE)
func (client *FutureRestClient) NewFutureOrderDelete() *FutureOrderDeleteApi {
	return &FutureOrderDeleteApi{
		FutureRestClient: *client,
		req:              &FutureOrderDeleteReq{},
	}
}

func (api *FutureOrderDeleteApi) Do() (*FutureOrderDeleteRes, error) {
	api.Timestamp(time.Now().UnixMilli())
	url := binanceHandlerRequestApiWithSecret(FUTURE, api.req, FutureApiMap[FutureOrderDelete], api.c.ApiSecret)
	return binanceCallApiWithSecret[FutureOrderDeleteRes](api.FutureRestClient.c, url, DELETE)
}

// binance FUTURE FutureBatchOrdersDelete rest批量撤销订单 (TRADE)
func (client *FutureRestClient) NewFutureBatchOrdersDelete() *FutureBatchOrdersDeleteApi {
	return &FutureBatchOrdersDeleteApi{
		FutureRestClient: *client,
		req:              &FutureBatchOrdersDeleteReq{},
	}
}

func (api *FutureBatchOrdersDeleteApi) Do() (*FutureBatchOrdersDeleteRes, error) {
	api.Timestamp(time.Now().UnixMilli())
	url := binanceHandlerRequestApiWithSecret(FUTURE, api.req, FutureApiMap[FutureBatchOrdersDelete], api.c.ApiSecret)
	return binanceCallApiWithSecret[FutureBatchOrdersDeleteRes](api.FutureRestClient.c, url, DELETE)
}

// binance FUTURE FutureUserTrades rest账户成交历史 (USER_DATA)
func (client *FutureRestClient) NewFutureUserTrades() *FutureUserTradesApi {
	return &FutureUserTradesApi{
		FutureRestClient: *client,
		req:              &FutureUserTradesReq{},
	}
}

func (api *FutureUserTradesApi) Do() (*FutureUserTradesRes, error) {
	api.Timestamp(time.Now().UnixMilli())
	url := binanceHandlerRequestApiWithSecret(FUTURE, api.req, FutureApiMap[FutureUserTrades], api.c.ApiSecret)
	return binanceCallApiWithSecret[FutureUserTradesRes](api.FutureRestClient.c, url, GET)
}

// binance FUTURE FutureKlines restK线数据 (MARKET_DATA)
func (client *FutureRestClient) NewFutureKlines() *FutureKlinesApi {
	return &FutureKlinesApi{
		FutureRestClient: *client,
		req:              &FutureKlinesReq{},
	}
}
func (api *FutureKlinesApi) Do() (*KlinesRes, error) {
	url := binanceHandlerRequestApi(FUTURE, api.req, FutureApiMap[FutureKlines])
	res, err := binanceCallApiWithSecret[KlinesMiddle](api.FutureRestClient.c, url, GET)
	if err != nil {
		return nil, err
	}
	res2 := res.ConvertToRes()
	return &res2, nil
}
