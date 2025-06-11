package mybinanceapi

import "time"

// 账户接口
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

// binance FUTURE  FuturePositionRisk rest用户持仓风险V2 (USER_DATA)
func (client *FutureRestClient) NewFuturePositionRisk() *FuturePositionRiskApi {
	return &FuturePositionRiskApi{
		client: client,
		req:    &FuturePositionRiskReq{},
	}
}
func (api *FuturePositionRiskApi) Do() (*FuturePositionRiskRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(FUTURE, api.req, FutureApiMap[FuturePositionRisk], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[FuturePositionRiskRes](api.client.c, url, GET)
}

// binance FUTURE  FutureIncomeAsyn rest获取合约资金流水下载Id (USER_DATA)
func (client *FutureRestClient) NewFutureIncomeAsyn() *FutureIncomeAsynApi {
	return &FutureIncomeAsynApi{
		client: client,
		req:    &FutureIncomeAsynReq{},
	}
}
func (api *FutureIncomeAsynApi) Do() (*FutureIncomeAsynRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(FUTURE, api.req, FutureApiMap[FutureIncomeAsyn], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[FutureIncomeAsynRes](api.client.c, url, GET)
}

// binance FUTURE  FutureIncomeAsynId rest通过下载Id获取合约资金流水下载链接 (USER_DATA)
func (client *FutureRestClient) NewFutureIncomeAsynId() *FutureIncomeAsynIdApi {
	return &FutureIncomeAsynIdApi{
		client: client,
		req:    &FutureIncomeAsynIdReq{},
	}
}
func (api *FutureIncomeAsynIdApi) Do() (*FutureIncomeAsynIdRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(FUTURE, api.req, FutureApiMap[FutureIncomeAsynId], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[FutureIncomeAsynIdRes](api.client.c, url, GET)
}
