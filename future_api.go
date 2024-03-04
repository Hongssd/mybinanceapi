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

	FutureDepth            //深度信息
	FutureTrades           //最新成交
	FutureHistoricalTrades //历史成交
	FutureAggTrades        //近期成交(归集)
	FuturePremiumIndex     //最新标记价格和资金费率
	FutureFundingRate      //查询资金费率历史
	FutureFundingInfo      //查询资金费率信息
	FutureTicker24hr       //24hr价格变动情况
	FutureTickerPrice      //最新价格
	FutureTickerBookTicker //当前最优挂单
	FutureDataBasis        //基差数据
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

	FutureKlines:           "/fapi/v1/klines",            //GET接口 K线数据
	FutureDepth:            "/fapi/v1/depth",             //GET接口 深度信息
	FutureTrades:           "/fapi/v1/trades",            //GET接口 最新成交
	FutureHistoricalTrades: "/fapi/v1/historicalTrades",  //GET接口 查询历史成交(MARKET_DATA)
	FutureAggTrades:        "/fapi/v1/aggTrades",         //GET接口 近期成交(归集)
	FuturePremiumIndex:     "/fapi/v1/premiumIndex",      //GET接口 最新标记价格和资金费率
	FutureFundingRate:      "/fapi/v1/fundingRate",       //GET接口 查询资金费率历史
	FutureFundingInfo:      "/fapi/v1/fundingInfo",       //GET接口 查询资金费率信息
	FutureTicker24hr:       "/fapi/v1/ticker/24hr",       //GET接口 24hr价格变动情况
	FutureTickerPrice:      "/fapi/v1/ticker/price",      //GET接口 最新价格
	FutureTickerBookTicker: "/fapi/v1/ticker/bookTicker", //GET接口 当前最优挂单
	FutureDataBasis:        "/futures/data/basis",        //GET接口 基差数据

}

//=======GET接口
// binance FUTURE  FutureAccount rest账户信息 (USER_DATA)
func (client *FutureRestClient) NewFutureAccount() *FutureAccountApi {
	return &FutureAccountApi{
		client: client,
		req:    &FutureAccountReq{},
	}
}
func (api *FutureAccountApi) Do() (*FutureAccountRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(FUTURE, api.req, FutureApiMap[FutureAccount], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[FutureAccountRes](api.client.c, url, GET)
}

// binance FUTURE  FuturePositionSideDualGet rest查询持仓模式 (USER_DATA)
func (client *FutureRestClient) NewFuturePositionSideDualGet() *FuturePositionSideDualGetApi {
	return &FuturePositionSideDualGetApi{
		client: client,
		req:    &FuturePositionSideDualGetReq{},
	}
}
func (api *FuturePositionSideDualGetApi) Do() (*FuturePositionSideDualGetRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(FUTURE, api.req, FutureApiMap[FuturePositionSideDualGet], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[FuturePositionSideDualGetRes](api.client.c, url, GET)
}

// binance FUTURE  FutureMultiAssetsMarginGet rest查询联合保证金模式 (USER_DATA)
func (client *FutureRestClient) NewFutureMultiAssetsMarginGet() *FutureMultiAssetsMarginGetApi {
	return &FutureMultiAssetsMarginGetApi{
		client: client,
		req:    &FutureMultiAssetsMarginGetReq{},
	}
}
func (api *FutureMultiAssetsMarginGetApi) Do() (*FutureMultiAssetsMarginGetRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(FUTURE, api.req, FutureApiMap[FutureMultiAssetsMarginGet], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[FutureMultiAssetsMarginGetRes](api.client.c, url, GET)
}

// binance FUTURE  FutureLeverageBracket rest杠杆分层标准 (USER_DATA)
func (client *FutureRestClient) NewFutureLeverageBracket() *FutureLeverageBracketApi {
	return &FutureLeverageBracketApi{
		client: client,
		req:    &FutureLeverageBracketReq{},
	}
}
func (api *FutureLeverageBracketApi) Do() (*FutureLeverageBracketRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(FUTURE, api.req, FutureApiMap[FutureLeverageBracket], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[FutureLeverageBracketRes](api.client.c, url, GET)
}

//========POST接口
// binance FUTURE  FuturePositionSideDualPost rest更改持仓模式 (TRADE)
func (client *FutureRestClient) NewFuturePositionSideDualPost() *FuturePositionSideDualPostApi {
	return &FuturePositionSideDualPostApi{
		client: client,
		req:    &FuturePositionSideDualPostReq{},
	}
}
func (api *FuturePositionSideDualPostApi) Do() (*FuturePositionSideDualPostRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(FUTURE, api.req, FutureApiMap[FuturePositionSideDualPost], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[FuturePositionSideDualPostRes](api.client.c, url, POST)
}

// binance FUTURE  FutureMultiAssetsMarginPost rest更改联合保证金模式 (TRADE)
func (client *FutureRestClient) NewFutureMultiAssetsMarginPost() *FutureMultiAssetsMarginPostApi {
	return &FutureMultiAssetsMarginPostApi{
		client: client,
		req:    &FutureMultiAssetsMarginPostReq{},
	}
}
func (api *FutureMultiAssetsMarginPostApi) Do() (*FutureMultiAssetsMarginPostRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(FUTURE, api.req, FutureApiMap[FutureMultiAssetsMarginPost], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[FutureMultiAssetsMarginPostRes](api.client.c, url, POST)
}

// binance FUTURE  FutureLeverage rest调整开仓杠杆 (TRADE)
func (client *FutureRestClient) NewFutureLeverage() *FutureLeverageApi {
	return &FutureLeverageApi{
		client: client,
		req:    &FutureLeverageReq{},
	}
}
func (api *FutureLeverageApi) Do() (*FutureLeverageRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(FUTURE, api.req, FutureApiMap[FutureLeverage], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[FutureLeverageRes](api.client.c, url, POST)
}

// binance FUTURE  FutureMarginType rest变换逐全仓模式 (TRADE)
func (client *FutureRestClient) NewFutureMarginType() *FutureMarginTypeApi {
	return &FutureMarginTypeApi{
		client: client,
		req:    &FutureMarginTypeReq{},
	}
}
func (api *FutureMarginTypeApi) Do() (*FutureMarginTypeRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(FUTURE, api.req, FutureApiMap[FutureMarginType], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[FutureMarginTypeRes](api.client.c, url, POST)
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

//通用接口
// binance FUTURE FuturePing rest测试服务器连通性 (NONE)
func (client *FutureRestClient) NewPing() *FuturePingApi {
	return &FuturePingApi{
		client: client,
		req:    &FuturePingReq{},
	}
}

func (api *FuturePingApi) Do() (*FuturePingRes, error) {
	url := binanceHandlerRequestApi(FUTURE, api.req, FutureApiMap[FuturePing])
	return binanceCallApiWithSecret[FuturePingRes](api.client.c, url, GET)
}

// binance FUTURE FutureTime rest获取服务器时间 (NONE)
func (client *FutureRestClient) NewServerTime() *FutureServerTimeApi {
	return &FutureServerTimeApi{
		client: client,
		req:    &FutureServerTimeReq{},
	}
}

func (api *FutureServerTimeApi) Do() (*FutureTimeRes, error) {
	url := binanceHandlerRequestApi(FUTURE, api.req, FutureApiMap[FutureServerTime])
	return binanceCallApiWithSecret[FutureTimeRes](api.client.c, url, GET)
}

// binance FUTURE FutureExchangeInfo rest交易规范信息和交易对信息 (NONE)
func (client *FutureRestClient) NewExchangeInfo() *FutureExchangeInfoApi {
	return &FutureExchangeInfoApi{
		client: client,
		req:    &FutureExchangeInfoReq{},
	}
}

func (api *FutureExchangeInfoApi) Do() (*FutureExchangeInfoRes, error) {
	url := binanceHandlerRequestApi(FUTURE, api.req, FutureApiMap[FutureExchangeInfo])
	return binanceCallApiWithSecret[FutureExchangeInfoRes](api.client.c, url, GET)
}

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

// binance FUTURE FutureKlines restK线数据 (MARKET_DATA)
func (client *FutureRestClient) NewFutureKlines() *FutureKlinesApi {
	return &FutureKlinesApi{
		client: client,
		req:    &FutureKlinesReq{},
	}
}
func (api *FutureKlinesApi) Do() (*KlinesRes, error) {
	url := binanceHandlerRequestApi(FUTURE, api.req, FutureApiMap[FutureKlines])
	res, err := binanceCallApiWithSecret[KlinesMiddle](api.client.c, url, GET)
	if err != nil {
		return nil, err
	}
	res2 := res.ConvertToRes()
	return &res2, nil
}

// binance FUTURE FutureDepth rest深度信息 (MARKET_DATA)
func (client *FutureRestClient) NewFutureDepth() *FutureDepthApi {
	return &FutureDepthApi{
		client: client,
		req:    &FutureDepthReq{},
	}
}
func (api *FutureDepthApi) Do() (*FutureDepthRes, error) {
	url := binanceHandlerRequestApi(FUTURE, api.req, FutureApiMap[FutureDepth])
	res, err := binanceCallApiWithSecret[FutureDepthResMiddle](api.client.c, url, GET)
	if err != nil {
		return nil, err
	}
	return res.ConvertToRes(), nil
}

// binance FUTURE FutureTrades rest最新成交 (MARKET_DATA)
func (client *FutureRestClient) NewFutureTrades() *FutureTradesApi {
	return &FutureTradesApi{
		client: client,
		req:    &FutureTradesReq{},
	}
}
func (api *FutureTradesApi) Do() (*FutureTradesRes, error) {
	url := binanceHandlerRequestApi(FUTURE, api.req, FutureApiMap[FutureTrades])
	return binanceCallApiWithSecret[FutureTradesRes](api.client.c, url, GET)
}

// binance FUTURE FutureHistoricalTrades rest历史成交 (MARKET_DATA)
func (client *FutureRestClient) NewFutureHistoricalTrades() *FutureHistoricalTradesApi {
	return &FutureHistoricalTradesApi{
		client: client,
		req:    &FutureHistoricalTradesReq{},
	}
}
func (api *FutureHistoricalTradesApi) Do() (*FutureHistoricalTradesRes, error) {
	url := binanceHandlerRequestApi(FUTURE, api.req, FutureApiMap[FutureHistoricalTrades])
	return binanceCallApiWithSecret[FutureHistoricalTradesRes](api.client.c, url, GET)
}

// binance FUTURE FutureAggTrades rest近期成交(归集) (MARKET_DATA)
func (client *FutureRestClient) NewFutureAggTrades() *FutureAggTradesApi {
	return &FutureAggTradesApi{
		client: client,
		req:    &FutureAggTradesReq{},
	}
}
func (api *FutureAggTradesApi) Do() (*FutureAggTradesRes, error) {
	url := binanceHandlerRequestApi(FUTURE, api.req, FutureApiMap[FutureAggTrades])
	return binanceCallApiWithSecret[FutureAggTradesRes](api.client.c, url, GET)
}

// binance FUTURE FuturePremiumIndex rest最新标记价格和资金费率 (MARKET_DATA)
func (client *FutureRestClient) NewFuturePremiumIndex() *FuturePremiumIndexApi {
	return &FuturePremiumIndexApi{
		client: client,
		req:    &FuturePremiumIndexReq{},
	}
}
func (api *FuturePremiumIndexApi) Do() (*FuturePremiumIndexRes, error) {
	url := binanceHandlerRequestApi(FUTURE, api.req, FutureApiMap[FuturePremiumIndex])
	if api.req.Symbol != nil && *api.req.Symbol != "" {
		res, err := binanceCallApiWithSecret[FuturePremiumIndexResRow](api.client.c, url, GET)
		if err != nil {
			return nil, err
		}
		return &FuturePremiumIndexRes{*res}, nil
	} else {
		return binanceCallApiWithSecret[FuturePremiumIndexRes](api.client.c, url, GET)
	}

}

// binance FUTURE FutureFundingRate rest查询资金费率历史 (MARKET_DATA)
func (client *FutureRestClient) NewFutureFundingRate() *FutureFundingRateApi {
	return &FutureFundingRateApi{
		client: client,
		req:    &FutureFundingRateReq{},
	}
}
func (api *FutureFundingRateApi) Do() (*FutureFundingRateRes, error) {
	url := binanceHandlerRequestApi(FUTURE, api.req, FutureApiMap[FutureFundingRate])
	return binanceCallApiWithSecret[FutureFundingRateRes](api.client.c, url, GET)
}

// binance FUTURE FutureFundingInfo rest查询资金费率信息 (MARKET_DATA)
func (client *FutureRestClient) NewFutureFundingInfo() *FutureFundingInfoApi {
	return &FutureFundingInfoApi{
		client: client,
		req:    &FutureFundingInfoReq{},
	}
}
func (api *FutureFundingInfoApi) Do() (*FutureFundingInfoRes, error) {
	url := binanceHandlerRequestApi(FUTURE, api.req, FutureApiMap[FutureFundingInfo])
	return binanceCallApiWithSecret[FutureFundingInfoRes](api.client.c, url, GET)
}

// binance FUTURE FutureTicker24hr rest24hr价格变动情况 (MARKET_DATA)
func (client *FutureRestClient) NewFutureTicker24hr() *FutureTicker24hrApi {
	return &FutureTicker24hrApi{
		client: client,
		req:    &FutureTicker24hrReq{},
	}
}
func (api *FutureTicker24hrApi) Do() (*FutureTicker24hrRes, error) {
	url := binanceHandlerRequestApi(FUTURE, api.req, FutureApiMap[FutureTicker24hr])
	if api.req.Symbol != nil && *api.req.Symbol != "" {
		res, err := binanceCallApiWithSecret[FutureTicker24hrResRow](api.client.c, url, GET)
		if err != nil {
			return nil, err
		}
		return &FutureTicker24hrRes{*res}, nil
	} else {
		return binanceCallApiWithSecret[FutureTicker24hrRes](api.client.c, url, GET)
	}
}

// binance FUTURE FutureTickerPrice rest最新价格 (MARKET_DATA)
func (client *FutureRestClient) NewFutureTickerPrice() *FutureTickerPriceApi {
	return &FutureTickerPriceApi{
		client: client,
		req:    &FutureTickerPriceReq{},
	}
}
func (api *FutureTickerPriceApi) Do() (*FutureTickerPriceRes, error) {
	url := binanceHandlerRequestApi(FUTURE, api.req, FutureApiMap[FutureTickerPrice])
	if api.req.Symbol != nil && *api.req.Symbol != "" {
		res, err := binanceCallApiWithSecret[FutureTickerPriceResRow](api.client.c, url, GET)
		if err != nil {
			return nil, err
		}
		return &FutureTickerPriceRes{*res}, nil
	} else {
		return binanceCallApiWithSecret[FutureTickerPriceRes](api.client.c, url, GET)
	}
}

// binance FUTURE FutureTickerBookTicker rest当前最优挂单 (MARKET_DATA)
func (client *FutureRestClient) NewFutureTickerBookTicker() *FutureTickerBookTickerApi {
	return &FutureTickerBookTickerApi{
		client: client,
		req:    &FutureTickerBookTickerReq{},
	}
}
func (api *FutureTickerBookTickerApi) Do() (*FutureTickerBookTickerRes, error) {
	url := binanceHandlerRequestApi(FUTURE, api.req, FutureApiMap[FutureTickerBookTicker])
	if api.req.Symbol != nil && *api.req.Symbol != "" {
		res, err := binanceCallApiWithSecret[FutureTickerBookTickerResRow](api.client.c, url, GET)
		if err != nil {
			return nil, err
		}
		return &FutureTickerBookTickerRes{*res}, nil
	} else {
		return binanceCallApiWithSecret[FutureTickerBookTickerRes](api.client.c, url, GET)
	}
}

// binance FUTURE FutureDataBasis rest基差数据 (MARKET_DATA)
func (client *FutureRestClient) NewFutureDataBasis() *FutureDataBasisApi {
	return &FutureDataBasisApi{
		client: client,
		req:    &FutureDataBasisReq{},
	}
}
func (api *FutureDataBasisApi) Do() (*FutureDataBasisRes, error) {
	url := binanceHandlerRequestApi(FUTURE, api.req, FutureApiMap[FutureDataBasis])
	return binanceCallApiWithSecret[FutureDataBasisRes](api.client.c, url, GET)
}
