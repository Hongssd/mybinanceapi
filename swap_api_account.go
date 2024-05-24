package mybinanceapi

import "time"

// 币合约账户接口
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
