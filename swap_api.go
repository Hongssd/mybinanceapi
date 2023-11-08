package mybinanceapi

import "time"

type SwapApi int

const (
	SwapAccount SwapApi = iota
	SwapPositionSideDualGet
	SwapLeverageBracket
	SwapPositionSideDualPost
	SwapLeverage
	SwapMarginType
	SwapOpenOrders
	SwapAllOrders
	SwapOrderPost
	SwapOrderGet

	//通用接口
	SwapPing
	SwapServerTime
	SwapExchangeInfo

	//行情接口
	SwapKlines
)

var SwapApiMap = map[SwapApi]string{
	SwapAccount: "/dapi/v1/account", //GET接口 账户信息 (USER_DATA)

	//GET接口
	SwapPositionSideDualGet: "/dapi/v1/positionSide/dual", //GET接口 (HMAC SHA256) 查询持仓模式(USER_DATA)
	SwapLeverageBracket:     "/dapi/v2/leverageBracket",   //GET接口 杠杆分层标准 (USER_DATA)
	//POST接口
	SwapPositionSideDualPost: "/dapi/v1/positionSide/dual", //POST接口 (HMAC SHA256) 更改持仓模式(TRADE)
	SwapLeverage:             "/dapi/v1/leverage",          //POST接口 (HMAC SHA256) 调整开仓杠杆 (TRADE)
	SwapMarginType:           "/dapi/v1/marginType",        //POST接口 (HMAC SHA256) 变换逐全仓模式 (TRADE)

	SwapOpenOrders: "/dapi/v1/openOrders", //GET接口 (HMAC SHA256) 查询当前挂单 (USER_DATA)
	SwapAllOrders:  "/dapi/v1/allOrders",  //GET接口 (HMAC SHA256) 查询所有订单 (USER_DATA)

	SwapOrderPost: "/dapi/v1/order", //POST接口 (HMAC SHA256) 下单 (TRADE)
	SwapOrderGet:  "/dapi/v1/order", //GET接口 (HMAC SHA256) 查询订单 (USER_DATA)

	//通用接口
	SwapPing:         "/dapi/v1/ping",         //GET接口 测试连接
	SwapServerTime:   "/dapi/v1/time",         //GET接口 获取服务器时间
	SwapExchangeInfo: "/dapi/v1/exchangeInfo", //GET接口 交易规则和交易对信息

	//行情接口
	SwapKlines: "/dapi/v1/klines", //GET接口 获取K线数据
}

//============GET接口
// binance SWAP  SwapAccount rest账户信息 (USER_DATA)
func (client *SwapRestClient) NewSwapAccount() *SwapAccountApi {
	return &SwapAccountApi{
		SwapRestClient: *client,
		req:            &SwapAccountReq{},
	}
}
func (api *SwapAccountApi) Do() (*SwapAccountRes, error) {
	api.Timestamp(time.Now().UnixMilli())
	url := binanceHandlerRequestApiWithSecretGet(SWAP, api.req, SwapApiMap[SwapAccount], api.c.ApiSecret)
	return binanceCallApiWithSecretGet[SwapAccountRes](api.SwapRestClient.c, url)
}

// binance SWAP  SwapPositionSideDualGet rest查询持仓模式 (USER_DATA)
func (client *SwapRestClient) NewSwapPositionSideDualGet() *SwapPositionSideDualGetApi {
	return &SwapPositionSideDualGetApi{
		SwapRestClient: *client,
		req:            &SwapPositionSideDualGetReq{},
	}
}
func (api *SwapPositionSideDualGetApi) Do() (*SwapPositionSideDualGetRes, error) {
	api.Timestamp(time.Now().UnixMilli())
	url := binanceHandlerRequestApiWithSecretGet(SWAP, api.req, SwapApiMap[SwapPositionSideDualGet], api.c.ApiSecret)
	return binanceCallApiWithSecretGet[SwapPositionSideDualGetRes](api.SwapRestClient.c, url)
}

// binance SWAP  SwapLeverageBracket rest杠杆分层标准 (USER_DATA)
func (client *SwapRestClient) NewSwapLeverageBracket() *SwapLeverageBracketApi {
	return &SwapLeverageBracketApi{
		SwapRestClient: *client,
		req:            &SwapLeverageBracketReq{},
	}
}
func (api *SwapLeverageBracketApi) Do() (*SwapLeverageBracketRes, error) {
	api.Timestamp(time.Now().UnixMilli())
	url := binanceHandlerRequestApiWithSecretGet(SWAP, api.req, SwapApiMap[SwapLeverageBracket], api.c.ApiSecret)
	return binanceCallApiWithSecretGet[SwapLeverageBracketRes](api.SwapRestClient.c, url)
}

//============POST接口
// binance SWAP  SwapPositionSideDualPost rest更改持仓模式 (TRADE)
func (client *SwapRestClient) NewSwapPositionSideDualPost() *SwapPositionSideDualPostApi {
	return &SwapPositionSideDualPostApi{
		SwapRestClient: *client,
		req:            &SwapPositionSideDualPostReq{},
	}
}
func (api *SwapPositionSideDualPostApi) Do() (*SwapPositionSideDualPostRes, error) {
	api.Timestamp(time.Now().UnixMilli())
	url, reqBody := binanceHandlerRequestApiWithSecretPost(SWAP, api.req, SwapApiMap[SwapPositionSideDualPost], api.c.ApiSecret)
	return binanceCallApiWithSecretPost[SwapPositionSideDualPostRes](api.SwapRestClient.c, url, reqBody)
}

// binance SWAP  SwapLeverage rest调整开仓杠杆 (TRADE)
func (client *SwapRestClient) NewSwapLeverage() *SwapLeverageApi {
	return &SwapLeverageApi{
		SwapRestClient: *client,
		req:            &SwapLeverageReq{},
	}
}
func (api *SwapLeverageApi) Do() (*SwapLeverageRes, error) {
	api.Timestamp(time.Now().UnixMilli())
	url, reqBody := binanceHandlerRequestApiWithSecretPost(SWAP, api.req, SwapApiMap[SwapLeverage], api.c.ApiSecret)
	return binanceCallApiWithSecretPost[SwapLeverageRes](api.SwapRestClient.c, url, reqBody)
}

// binance SWAP  SwapMarginType rest变换逐全仓模式 (TRADE)
func (client *SwapRestClient) NewSwapMarginType() *SwapMarginTypeApi {
	return &SwapMarginTypeApi{
		SwapRestClient: *client,
		req:            &SwapMarginTypeReq{},
	}
}
func (api *SwapMarginTypeApi) Do() (*SwapMarginTypeRes, error) {
	api.Timestamp(time.Now().UnixMilli())
	url, reqBody := binanceHandlerRequestApiWithSecretPost(SWAP, api.req, SwapApiMap[SwapMarginType], api.c.ApiSecret)
	return binanceCallApiWithSecretPost[SwapMarginTypeRes](api.SwapRestClient.c, url, reqBody)
}

//通用接口
// binance SWAP  SwapPing rest测试连接
func (client *SwapRestClient) NewPing() *SwapPingApi {
	return &SwapPingApi{
		SwapRestClient: *client,
		req:            &SwapPingReq{},
	}
}
func (api *SwapPingApi) Do() (*SwapPingRes, error) {
	url := binanceHandlerRequestApiGet(SWAP, api.req, SwapApiMap[SwapPing])
	return binanceCallApiWithSecretGet[SwapPingRes](api.SwapRestClient.c, url)
}

// binance SWAP  SwapServerTime rest获取服务器时间
func (client *SwapRestClient) NewServerTime() *SwapServerTimeApi {
	return &SwapServerTimeApi{
		SwapRestClient: *client,
		req:            &SwapServerTimeReq{},
	}
}
func (api *SwapServerTimeApi) Do() (*SwapServerTimeRes, error) {
	url := binanceHandlerRequestApiGet(SWAP, api.req, SwapApiMap[SwapServerTime])
	return binanceCallApiWithSecretGet[SwapServerTimeRes](api.SwapRestClient.c, url)
}

// binance SWAP  SwapExchangeInfo rest交易规则和交易对信息
func (client *SwapRestClient) NewExchangeInfo() *SwapExchangeInfoApi {
	return &SwapExchangeInfoApi{
		SwapRestClient: *client,
		req:            &SwapExchangeInfoReq{},
	}
}
func (api *SwapExchangeInfoApi) Do() (*SwapExchangeInfoRes, error) {
	url := binanceHandlerRequestApiGet(SWAP, api.req, SwapApiMap[SwapExchangeInfo])
	return binanceCallApiWithSecretGet[SwapExchangeInfoRes](api.SwapRestClient.c, url)
}

// binance SWAP  SwapOpenOrders rest查询当前挂单 (USER_DATA)
func (client *SwapRestClient) NewOpenOrders() *SwapOpenOrdersApi {
	return &SwapOpenOrdersApi{
		SwapRestClient: *client,
		req:            &SwapOpenOrdersReq{},
	}
}
func (api *SwapOpenOrdersApi) Do() (*SwapOpenOrdersRes, error) {
	api.Timestamp(time.Now().UnixMilli())
	url := binanceHandlerRequestApiWithSecretGet(SWAP, api.req, SwapApiMap[SwapOpenOrders], api.c.ApiSecret)
	return binanceCallApiWithSecretGet[SwapOpenOrdersRes](api.SwapRestClient.c, url)
}

// binance SWAP  SwapAllOrders rest查询所有订单 (USER_DATA)
func (client *SwapRestClient) NewAllOrders() *SwapAllOrdersApi {
	return &SwapAllOrdersApi{
		SwapRestClient: *client,
		req:            &SwapAllOrdersReq{},
	}
}
func (api *SwapAllOrdersApi) Do() (*SwapAllOrdersRes, error) {
	api.Timestamp(time.Now().UnixMilli())
	url := binanceHandlerRequestApiWithSecretGet(SWAP, api.req, SwapApiMap[SwapAllOrders], api.c.ApiSecret)
	return binanceCallApiWithSecretGet[SwapAllOrdersRes](api.SwapRestClient.c, url)
}

// binance SWAP  SwapOrderPost rest下单 (TRADE)
func (client *SwapRestClient) NewSwapOrderPost() *SwapOrderPostApi {
	return &SwapOrderPostApi{
		SwapRestClient: *client,
		req:            &SwapOrderPostReq{},
	}
}
func (api *SwapOrderPostApi) Do() (*SwapOrderPostRes, error) {
	api.Timestamp(time.Now().UnixMilli())
	url, reqBody := binanceHandlerRequestApiWithSecretPost(SWAP, api.req, SwapApiMap[SwapOrderPost], api.c.ApiSecret)
	return binanceCallApiWithSecretPost[SwapOrderPostRes](api.SwapRestClient.c, url, reqBody)
}

// binance SWAP  SwapOrderGet	rest查询订单 (USER_DATA)
func (client *SwapRestClient) NewSwapOrderGet() *SwapOrderGetApi {
	return &SwapOrderGetApi{
		SwapRestClient: *client,
		req:            &SwapOrderGetReq{},
	}
}
func (api *SwapOrderGetApi) Do() (*SwapOrderGetRes, error) {
	api.Timestamp(time.Now().UnixMilli())
	url := binanceHandlerRequestApiWithSecretGet(SWAP, api.req, SwapApiMap[SwapOrderGet], api.c.ApiSecret)
	return binanceCallApiWithSecretGet[SwapOrderGetRes](api.SwapRestClient.c, url)
}

// binance SWAP  SwapKlines rest获取K线数据
func (client *SwapRestClient) NewSwapKlines() *SwapKlinesApi {
	return &SwapKlinesApi{
		SwapRestClient: *client,
		req:            &SwapKlinesReq{},
	}
}
func (api *SwapKlinesApi) Do() (*KlinesRes, error) {
	url := binanceHandlerRequestApiGet(SWAP, api.req, SwapApiMap[SwapKlines])
	res, err := binanceCallApiWithSecretGet[KlinesMiddle](api.SwapRestClient.c, url)
	if err != nil {
		return nil, err
	}
	res2 := res.ConvertToRes()
	return &res2, nil
}
