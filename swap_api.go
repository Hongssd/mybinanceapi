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
	SwapOrderDelete

	//通用接口
	SwapPing
	SwapServerTime
	SwapExchangeInfo

	//行情接口
	SwapKlines
	SwapTickerPrice //获取交易对最新价格
	SwapDepth       //获取深度信息

	//Ws账户推送相关
	SwapListenKeyPost //生成listenKey (USER_STREAM)
	SwapListenKeyPut  //延长listenKey有效期 (USER_STREAM)
	SwapListenKeyDel  //关闭listenKey (USER_STREAM)
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

	SwapOrderPost:   "/dapi/v1/order", //POST接口 (HMAC SHA256) 下单 (TRADE)
	SwapOrderGet:    "/dapi/v1/order", //GET接口 (HMAC SHA256) 查询订单 (USER_DATA)
	SwapOrderDelete: "/dapi/v1/order", //DELETE接口 (HMAC SHA256) 撤销订单 (TRADE)
	//通用接口
	SwapPing:         "/dapi/v1/ping",         //GET接口 测试连接
	SwapServerTime:   "/dapi/v1/time",         //GET接口 获取服务器时间
	SwapExchangeInfo: "/dapi/v1/exchangeInfo", //GET接口 交易规则和交易对信息

	//行情接口
	SwapKlines:      "/dapi/v1/klines",       //GET接口 获取K线数据
	SwapTickerPrice: "/dapi/v1/ticker/price", //GET接口 获取交易对最新价格
	SwapDepth:       "/dapi/v1/depth",        //GET接口 获取深度信息

	//Ws账户推送相关
	SwapListenKeyPost: "/dapi/v1/listenKey", //POST接口 (HMAC SHA256) 创建listenKey (USER_STREAM)\
	SwapListenKeyPut:  "/dapi/v1/listenKey", //PUT接口 (HMAC SHA256) 延长listenKey有效期 (USER_STREAM)
	SwapListenKeyDel:  "/dapi/v1/listenKey", //DELETE接口 (HMAC SHA256) 关闭listenKey (USER_STREAM)
}

//============GET接口
// binance SWAP  SwapAccount rest账户信息 (USER_DATA)
func (client *SwapRestClient) NewSwapAccount() *SwapAccountApi {
	return &SwapAccountApi{
		client: client,
		req:    &SwapAccountReq{},
	}
}
func (api *SwapAccountApi) Do() (*SwapAccountRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SWAP, api.req, SwapApiMap[SwapAccount], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SwapAccountRes](api.client.c, url, GET)
}

// binance SWAP  SwapPositionSideDualGet rest查询持仓模式 (USER_DATA)
func (client *SwapRestClient) NewSwapPositionSideDualGet() *SwapPositionSideDualGetApi {
	return &SwapPositionSideDualGetApi{
		client: client,
		req:    &SwapPositionSideDualGetReq{},
	}
}
func (api *SwapPositionSideDualGetApi) Do() (*SwapPositionSideDualGetRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SWAP, api.req, SwapApiMap[SwapPositionSideDualGet], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SwapPositionSideDualGetRes](api.client.c, url, GET)
}

// binance SWAP  SwapLeverageBracket rest杠杆分层标准 (USER_DATA)
func (client *SwapRestClient) NewSwapLeverageBracket() *SwapLeverageBracketApi {
	return &SwapLeverageBracketApi{
		client: client,
		req:    &SwapLeverageBracketReq{},
	}
}
func (api *SwapLeverageBracketApi) Do() (*SwapLeverageBracketRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SWAP, api.req, SwapApiMap[SwapLeverageBracket], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SwapLeverageBracketRes](api.client.c, url, GET)
}

//============POST接口
// binance SWAP  SwapPositionSideDualPost rest更改持仓模式 (TRADE)
func (client *SwapRestClient) NewSwapPositionSideDualPost() *SwapPositionSideDualPostApi {
	return &SwapPositionSideDualPostApi{
		client: client,
		req:    &SwapPositionSideDualPostReq{},
	}
}
func (api *SwapPositionSideDualPostApi) Do() (*SwapPositionSideDualPostRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SWAP, api.req, SwapApiMap[SwapPositionSideDualPost], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SwapPositionSideDualPostRes](api.client.c, url, POST)
}

// binance SWAP  SwapLeverage rest调整开仓杠杆 (TRADE)
func (client *SwapRestClient) NewSwapLeverage() *SwapLeverageApi {
	return &SwapLeverageApi{
		client: client,
		req:    &SwapLeverageReq{},
	}
}
func (api *SwapLeverageApi) Do() (*SwapLeverageRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SWAP, api.req, SwapApiMap[SwapLeverage], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SwapLeverageRes](api.client.c, url, POST)
}

// binance SWAP  SwapMarginType rest变换逐全仓模式 (TRADE)
func (client *SwapRestClient) NewSwapMarginType() *SwapMarginTypeApi {
	return &SwapMarginTypeApi{
		client: client,
		req:    &SwapMarginTypeReq{},
	}
}
func (api *SwapMarginTypeApi) Do() (*SwapMarginTypeRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SWAP, api.req, SwapApiMap[SwapMarginType], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SwapMarginTypeRes](api.client.c, url, POST)
}

//通用接口
// binance SWAP  SwapPing rest测试连接
func (client *SwapRestClient) NewPing() *SwapPingApi {
	return &SwapPingApi{
		client: client,
		req:    &SwapPingReq{},
	}
}
func (api *SwapPingApi) Do() (*SwapPingRes, error) {
	url := binanceHandlerRequestApi(SWAP, api.req, SwapApiMap[SwapPing])
	return binanceCallApiWithSecret[SwapPingRes](api.client.c, url, GET)
}

// binance SWAP  SwapServerTime rest获取服务器时间
func (client *SwapRestClient) NewServerTime() *SwapServerTimeApi {
	return &SwapServerTimeApi{
		client: client,
		req:    &SwapServerTimeReq{},
	}
}
func (api *SwapServerTimeApi) Do() (*SwapServerTimeRes, error) {
	url := binanceHandlerRequestApi(SWAP, api.req, SwapApiMap[SwapServerTime])
	return binanceCallApiWithSecret[SwapServerTimeRes](api.client.c, url, GET)
}

// binance SWAP  SwapExchangeInfo rest交易规则和交易对信息
func (client *SwapRestClient) NewExchangeInfo() *SwapExchangeInfoApi {
	return &SwapExchangeInfoApi{
		client: client,
		req:    &SwapExchangeInfoReq{},
	}
}
func (api *SwapExchangeInfoApi) Do() (*SwapExchangeInfoRes, error) {
	url := binanceHandlerRequestApi(SWAP, api.req, SwapApiMap[SwapExchangeInfo])
	return binanceCallApiWithSecret[SwapExchangeInfoRes](api.client.c, url, GET)
}

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

// binance SWAP  SwapKlines rest获取K线数据
func (client *SwapRestClient) NewSwapKlines() *SwapKlinesApi {
	return &SwapKlinesApi{
		client: client,
		req:    &SwapKlinesReq{},
	}
}
func (api *SwapKlinesApi) Do() (*KlinesRes, error) {
	url := binanceHandlerRequestApi(SWAP, api.req, SwapApiMap[SwapKlines])
	res, err := binanceCallApiWithSecret[KlinesMiddle](api.client.c, url, GET)
	if err != nil {
		return nil, err
	}
	res2 := res.ConvertToRes()
	return &res2, nil
}

// binance SWAP  SwapTickerPrice rest获取交易对最新价格
func (client *SwapRestClient) NewSwapTickerPrice() *SwapTickerPriceApi {
	return &SwapTickerPriceApi{
		client: client,
		req:    &SwapTickerPriceReq{},
	}
}
func (api *SwapTickerPriceApi) Do() (*SwapTickerPriceRes, error) {
	url := binanceHandlerRequestApi(SWAP, api.req, SwapApiMap[SwapTickerPrice])
	return binanceCallApiWithSecret[SwapTickerPriceRes](api.client.c, url, GET)
}

// binance SWAP  SwapDepth rest获取深度信息
func (client *SwapRestClient) NewSwapDepth() *SwapDepthApi {
	return &SwapDepthApi{
		client: client,
		req:    &SwapDepthReq{},
	}
}
func (api *SwapDepthApi) Do() (*SwapDepthRes, error) {
	url := binanceHandlerRequestApi(SWAP, api.req, SwapApiMap[SwapDepth])
	res, err := binanceCallApiWithSecret[SwapDepthResMiddle](api.client.c, url, GET)
	if err != nil {
		return nil, err
	}
	return res.ConvertToRes(), nil
}

//============Ws账户推送相关

// binance SWAP  SwapListenKeyPost rest创建listenKey (USER_STREAM)
func (client *SwapRestClient) NewSwapListenKeyPost() *SwapListenKeyPostApi {
	return &SwapListenKeyPostApi{
		client: client,
		req:    &SwapListenKeyPostReq{},
	}
}
func (api *SwapListenKeyPostApi) Do() (*SwapListenKeyPostRes, error) {
	url := binanceHandlerRequestApiWithSecret(SWAP, api.req, SwapApiMap[SwapListenKeyPost], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SwapListenKeyPostRes](api.client.c, url, POST)
}

// binance SWAP  SwapListenKeyPut rest延长listenKey有效期 (USER_STREAM)
func (client *SwapRestClient) NewSwapListenKeyPut() *SwapListenKeyPutApi {
	return &SwapListenKeyPutApi{
		client: client,
		req:    &SwapListenKeyPutReq{},
	}
}
func (api *SwapListenKeyPutApi) Do() (*SwapListenKeyPutRes, error) {
	url := binanceHandlerRequestApiWithSecret(SWAP, api.req, SwapApiMap[SwapListenKeyPut], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SwapListenKeyPutRes](api.client.c, url, PUT)
}

// binance SWAP  SwapListenKeyDelete rest关闭listenKey (USER_STREAM)
func (client *SwapRestClient) NewSwapListenKeyDelete() *SwapListenKeyDeleteApi {
	return &SwapListenKeyDeleteApi{
		client: client,
		req:    &SwapListenKeyDeleteReq{},
	}
}
func (api *SwapListenKeyDeleteApi) Do() (*SwapListenKeyDeleteRes, error) {
	url := binanceHandlerRequestApiWithSecret(SWAP, api.req, SwapApiMap[SwapListenKeyDel], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SwapListenKeyDeleteRes](api.client.c, url, DELETE)
}
