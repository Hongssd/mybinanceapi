package mybinanceapi

import "time"

// 账户接口
// binance PortfolioMarginGetBalance rest查询账户余额 (NONE)
func (client *PortfolioMarginRestClient) NewGetBalance() *PortfolioMarginBalanceApi {
	return &PortfolioMarginBalanceApi{
		client: client,
		req:    &PortfolioMarginBalanceReq{},
	}
}
func (api *PortfolioMarginBalanceApi) Do() (*PortfolioMarginBalanceRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PORTFOLIO_MARGIN, api.req, PortfolioMarginApiMap[PortfolioMarginGetBalance], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PortfolioMarginBalanceRes](api.client.c, url, GET)
}

// binance PortfolioMarginGetAccount rest查询账户信息 (NONE)
func (client *PortfolioMarginRestClient) NewGetAccount() *PortfolioMarginAccountApi {
	return &PortfolioMarginAccountApi{
		client: client,
		req:    &PortfolioMarginAccountReq{},
	}
}
func (api *PortfolioMarginAccountApi) Do() (*PortfolioMarginAccountRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PORTFOLIO_MARGIN, api.req, PortfolioMarginApiMap[PortfolioMarginGetAccount], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PortfolioMarginAccountRes](api.client.c, url, GET)
}

// binance PortfolioMarginGetMaxBorrowable rest查询最大可借贷额度
func (client *PortfolioMarginRestClient) NewGetMaxBorrowable() *PortfolioMarginMaxBorrowableApi {
	return &PortfolioMarginMaxBorrowableApi{
		client: client,
		req:    &PortfolioMarginMaxBorrowableReq{},
	}
}
func (api *PortfolioMarginMaxBorrowableApi) Do() (*PortfolioMarginMaxBorrowableRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PORTFOLIO_MARGIN, api.req, PortfolioMarginApiMap[PortfolioMarginGetMaxBorrowable], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PortfolioMarginMaxBorrowableRes](api.client.c, url, GET)
}

// binance PortfolioMarginGetMaxWithdraw rest查询最大可转出额度
func (client *PortfolioMarginRestClient) NewGetMaxWithdraw() *PortfolioMarginMaxWithdrawApi {
	return &PortfolioMarginMaxWithdrawApi{
		client: client,
		req:    &PortfolioMarginMaxWithdrawReq{},
	}
}
func (api *PortfolioMarginMaxWithdrawApi) Do() (*PortfolioMarginMaxWithdrawRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PORTFOLIO_MARGIN, api.req, PortfolioMarginApiMap[PortfolioMarginGetMaxWithdraw], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PortfolioMarginMaxWithdrawRes](api.client.c, url, GET)
}

// binance PortfolioMarginSetUmLeverage rest设置UM开仓杠杆
func (client *PortfolioMarginRestClient) NewSetUmLeverage() *PortfolioMarginSetUmLeverageApi {
	return &PortfolioMarginSetUmLeverageApi{
		client: client,
		req:    &PortfolioMarginSetUmLeverageReq{},
	}
}
func (api *PortfolioMarginSetUmLeverageApi) Do() (*PortfolioMarginSetUmLeverageRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PORTFOLIO_MARGIN, api.req, PortfolioMarginApiMap[PortfolioMarginSetUmLeverage], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PortfolioMarginSetUmLeverageRes](api.client.c, url, POST)
}

// binance PortfolioMarginSetCmLeverage rest设置CM开仓杠杆
func (client *PortfolioMarginRestClient) NewSetCmLeverage() *PortfolioMarginSetCmLeverageApi {
	return &PortfolioMarginSetCmLeverageApi{
		client: client,
		req:    &PortfolioMarginSetCmLeverageReq{},
	}
}
func (api *PortfolioMarginSetCmLeverageApi) Do() (*PortfolioMarginSetCmLeverageRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PORTFOLIO_MARGIN, api.req, PortfolioMarginApiMap[PortfolioMarginSetCmLeverage], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PortfolioMarginSetCmLeverageRes](api.client.c, url, POST)
}

// binance PortfolioMarginUmPositionSideDualPost rest 更改UM持仓模式
func (client *PortfolioMarginRestClient) NewUmPositionSideDualPost() *PortfolioMarginUmPositionSideDualPostApi {
	return &PortfolioMarginUmPositionSideDualPostApi{
		client: client,
		req:    &PortfolioMarginUmPositionSideDualPostReq{},
	}
}

func (api *PortfolioMarginUmPositionSideDualPostApi) Do() (*PortfolioMarginUmPositionSideDualPostRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PORTFOLIO_MARGIN, api.req, PortfolioMarginApiMap[PortfolioMarginUmPositionSideDualPost], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PortfolioMarginUmPositionSideDualPostRes](api.client.c, url, POST)
}

// binance PortfolioMarginUmPositionSideDualGet rest 查询UM持仓模式
func (client *PortfolioMarginRestClient) NewUmPositionSideDualGet() *PortfolioMarginUmPositionSideDualGetApi {
	return &PortfolioMarginUmPositionSideDualGetApi{
		client: client,
		req:    &PortfolioMarginUmPositionSideDualGetReq{},
	}
}

func (api *PortfolioMarginUmPositionSideDualGetApi) Do() (*PortfolioMarginUmPositionSideDualGetRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PORTFOLIO_MARGIN, api.req, PortfolioMarginApiMap[PortfolioMarginUmPositionSideDualGet], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PortfolioMarginUmPositionSideDualGetRes](api.client.c, url, GET)
}

// binance PortfolioMarginCmPositionSideDualPost rest 更改CM持仓模式
func (client *PortfolioMarginRestClient) NewCmPositionSideDualPost() *PortfolioMarginCmPositionSideDualPostApi {
	return &PortfolioMarginCmPositionSideDualPostApi{
		client: client,
		req:    &PortfolioMarginCmPositionSideDualPostReq{},
	}
}

func (api *PortfolioMarginCmPositionSideDualPostApi) Do() (*PortfolioMarginCmPositionSideDualPostRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PORTFOLIO_MARGIN, api.req, PortfolioMarginApiMap[PortfolioMarginCmPositionSideDualPost], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PortfolioMarginCmPositionSideDualPostRes](api.client.c, url, POST)
}

// binance PortfolioMarginCmPositionSideDualGet rest 查询CM持仓模式
func (client *PortfolioMarginRestClient) NewCmPositionSideDualGet() *PortfolioMarginCmPositionSideDualGetApi {
	return &PortfolioMarginCmPositionSideDualGetApi{
		client: client,
		req:    &PortfolioMarginCmPositionSideDualGetReq{},
	}
}

func (api *PortfolioMarginCmPositionSideDualGetApi) Do() (*PortfolioMarginCmPositionSideDualGetRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PORTFOLIO_MARGIN, api.req, PortfolioMarginApiMap[PortfolioMarginCmPositionSideDualGet], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PortfolioMarginCmPositionSideDualGetRes](api.client.c, url, GET)
}

// binance PortfolioMarginUmAccountV1 rest 获取UM账户信息
func (client *PortfolioMarginRestClient) NewUmAccountV1() *PortfolioMarginUmAccountV1Api {
	return &PortfolioMarginUmAccountV1Api{
		client: client,
		req:    &PortfolioMarginUmAccountV1Req{},
	}
}

func (api *PortfolioMarginUmAccountV1Api) Do() (*PortfolioMarginUmAccountV1Res, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PORTFOLIO_MARGIN, api.req, PortfolioMarginApiMap[PortfolioMarginUmAccountV1], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PortfolioMarginUmAccountV1Res](api.client.c, url, GET)
}

// binance PortfolioMarginUmAccountV2 rest 获取UM账户信息
func (client *PortfolioMarginRestClient) NewUmAccountV2() *PortfolioMarginUmAccountV2Api {
	return &PortfolioMarginUmAccountV2Api{
		client: client,
		req:    &PortfolioMarginUmAccountV2Req{},
	}
}

func (api *PortfolioMarginUmAccountV2Api) Do() (*PortfolioMarginUmAccountV2Res, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PORTFOLIO_MARGIN, api.req, PortfolioMarginApiMap[PortfolioMarginUmAccountV2], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PortfolioMarginUmAccountV2Res](api.client.c, url, GET)
}

// binance PortfolioMarginCmAccount rest 获取CM账户信息
func (client *PortfolioMarginRestClient) NewCmAccount() *PortfolioMarginCmAccountApi {
	return &PortfolioMarginCmAccountApi{
		client: client,
		req:    &PortfolioMarginCmAccountReq{},
	}
}

func (api *PortfolioMarginCmAccountApi) Do() (*PortfolioMarginCmAccountRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PORTFOLIO_MARGIN, api.req, PortfolioMarginApiMap[PortfolioMarginCmAccount], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PortfolioMarginCmAccountRes](api.client.c, url, GET)
}

// binance PortfolioMarginUmPositionRisk rest 查询UM持仓风险
func (client *PortfolioMarginRestClient) NewUmPositionRisk() *PortfolioMarginUmPositionRiskApi {
	return &PortfolioMarginUmPositionRiskApi{
		client: client,
		req:    &PortfolioMarginUmPositionRiskReq{},
	}
}

func (api *PortfolioMarginUmPositionRiskApi) Do() (*PortfolioMarginUmPositionRiskRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PORTFOLIO_MARGIN, api.req, PortfolioMarginApiMap[PortfolioMarginUmPositionRisk], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PortfolioMarginUmPositionRiskRes](api.client.c, url, GET)
}

// binance PortfolioMarginCmPositionRisk rest 查询CM持仓风险
func (client *PortfolioMarginRestClient) NewCmPositionRisk() *PortfolioMarginCmPositionRiskApi {
	return &PortfolioMarginCmPositionRiskApi{
		client: client,
		req:    &PortfolioMarginCmPositionRiskReq{},
	}
}

func (api *PortfolioMarginCmPositionRiskApi) Do() (*PortfolioMarginCmPositionRiskRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PORTFOLIO_MARGIN, api.req, PortfolioMarginApiMap[PortfolioMarginCmPositionRisk], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PortfolioMarginCmPositionRiskRes](api.client.c, url, GET)
}

// binance PortfolioMarginUmCommissionRate rest 查询UM手续费率
func (client *PortfolioMarginRestClient) NewUmCommissionRate() *PortfolioMarginUmCommissionRateApi {
	return &PortfolioMarginUmCommissionRateApi{
		client: client,
		req:    &PortfolioMarginUmCommissionRateReq{},
	}
}

func (api *PortfolioMarginUmCommissionRateApi) Do() (*PortfolioMarginUmCommissionRateRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PORTFOLIO_MARGIN, api.req, PortfolioMarginApiMap[PortfolioMarginUmCommissionRate], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PortfolioMarginUmCommissionRateRes](api.client.c, url, GET)
}

// binance PortfolioMarginCmCommissionRate rest 查询CM手续费率
func (client *PortfolioMarginRestClient) NewCmCommissionRate() *PortfolioMarginCmCommissionRateApi {
	return &PortfolioMarginCmCommissionRateApi{
		client: client,
		req:    &PortfolioMarginCmCommissionRateReq{},
	}
}

func (api *PortfolioMarginCmCommissionRateApi) Do() (*PortfolioMarginCmCommissionRateRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PORTFOLIO_MARGIN, api.req, PortfolioMarginApiMap[PortfolioMarginCmCommissionRate], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PortfolioMarginCmCommissionRateRes](api.client.c, url, GET)
}

// binance PortfolioMarginAutoCollection rest 统一账户资金归集
func (client *PortfolioMarginRestClient) NewAutoCollection() *PortfolioMarginAutoCollectionApi {
	return &PortfolioMarginAutoCollectionApi{
		client: client,
		req:    &PortfolioMarginAutoCollectionReq{},
	}
}

func (api *PortfolioMarginAutoCollectionApi) Do() (*PortfolioMarginAutoCollectionRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PORTFOLIO_MARGIN, api.req, PortfolioMarginApiMap[PortfolioMarginAutoCollection], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PortfolioMarginAutoCollectionRes](api.client.c, url, POST)
}

// binance PortfolioMarginAssetCollection rest 特定资产资金归集
func (client *PortfolioMarginRestClient) NewAssetCollection() *PortfolioMarginAssetCollectionApi {
	return &PortfolioMarginAssetCollectionApi{
		client: client,
		req:    &PortfolioMarginAssetCollectionReq{},
	}
}

func (api *PortfolioMarginAssetCollectionApi) Do() (*PortfolioMarginAssetCollectionRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PORTFOLIO_MARGIN, api.req, PortfolioMarginApiMap[PortfolioMarginAssetCollection], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PortfolioMarginAssetCollectionRes](api.client.c, url, POST)
}
