package mybinanceapi

import "time"

// 子母账户接口
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

// binance SPOT子母账户接口  ManagedSubAccountQueryTransLog rest查询托管子账户的划转记录(适用交易团队子账户)(USER_DATA)
func (client *SpotRestClient) NewManagedSubAccountQueryTransLog() *SpotManagedSubAccountQueryTransLogApi {
	return &SpotManagedSubAccountQueryTransLogApi{
		client: client,
		req:    &SpotManagedSubAccountQueryTransLogReq{},
	}
}
func (api *SpotManagedSubAccountQueryTransLogApi) Do() (*ManagedSubAccountQueryTransLogRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotManagedSubAccountQueryTransLog], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[ManagedSubAccountQueryTransLogRes](api.client.c, url, GET)
}

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

// binance SPOT子母账户接口  SpotSubAccountMarginEnable rest为子账户开通杠杆 (适用主账户)
func (client *SpotRestClient) NewSpotSubAccountMarginEnable() *SpotSubAccountMarginEnableApi {
	return &SpotSubAccountMarginEnableApi{
		client: client,
		req:    &SpotSubAccountMarginEnableReq{},
	}
}
func (api *SpotSubAccountMarginEnableApi) Do() (*SpotSubAccountMarginEnableRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotSubAccountMarginEnable], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SpotSubAccountMarginEnableRes](api.client.c, url, POST)
}

// binance SPOT子母账户接口  SpotSubAccountSubAccountApiIpRestrictionPost 为子账户API Key增加IP白名单(适用母账户)
func (client *SpotRestClient) NewSpotSubAccountSubAccountApiIpRestrictionPost() *SpotSubAccountSubAccountApiIpRestrictionPostApi {
	return &SpotSubAccountSubAccountApiIpRestrictionPostApi{
		client: client,
		req:    &SpotSubAccountSubAccountApiIpRestrictionPostReq{},
	}
}
func (api *SpotSubAccountSubAccountApiIpRestrictionPostApi) Do() (*SpotSubAccountSubAccountApiIpRestrictionPostRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotSubAccountSubAccountApiIpRestrictionPost], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SpotSubAccountSubAccountApiIpRestrictionPostRes](api.client.c, url, POST)
}

// binance SPOT子母账户接口  SpotSubAccountSubAccountApiIpRestrictionGet 为子账户API Key查询IP白名单(适用母账户)
func (client *SpotRestClient) NewSpotSubAccountSubAccountApiIpRestrictionGet() *SpotSubAccountSubAccountApiIpRestrictionGetApi {
	return &SpotSubAccountSubAccountApiIpRestrictionGetApi{
		client: client,
		req:    &SpotSubAccountSubAccountApiIpRestrictionGetReq{},
	}
}
func (api *SpotSubAccountSubAccountApiIpRestrictionGetApi) Do() (*SpotSubAccountSubAccountApiIpRestrictionGetRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotSubAccountSubAccountApiIpRestrictionGet], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SpotSubAccountSubAccountApiIpRestrictionGetRes](api.client.c, url, GET)
}

// binance SPOT子母账户接口  SpotSubAccountSubAccountApiIpRestrictionDelete 为子账户API Key删除IP白名单(适用母账户)
func (client *SpotRestClient) NewSpotSubAccountSubAccountApiIpRestrictionDelete() *SpotSubAccountSubAccountApiIpRestrictionDeleteApi {
	return &SpotSubAccountSubAccountApiIpRestrictionDeleteApi{
		client: client,
		req:    &SpotSubAccountSubAccountApiIpRestrictionDeleteReq{},
	}
}
func (api *SpotSubAccountSubAccountApiIpRestrictionDeleteApi) Do() (*SpotSubAccountSubAccountApiIpRestrictionDeleteRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotSubAccountSubAccountApiIpRestrictionDelete], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SpotSubAccountSubAccountApiIpRestrictionDeleteRes](api.client.c, url, DELETE)
}

// binance SPOT子母账户接口 SpotSubAccountCapitalDepositSubAddress 查询子账户充值地址 (适用主账户)
func (client *SpotRestClient) NewSpotSubAccountCapitalDepositSubAddress() *SpotSubAccountCapitalDepositSubAddressApi {
	return &SpotSubAccountCapitalDepositSubAddressApi{
		client: client,
		req:    &SpotSubAccountCapitalDepositSubAddressReq{},
	}
}
func (api *SpotSubAccountCapitalDepositSubAddressApi) Do() (*SpotSubAccountCapitalDepositSubAddressRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotSubAccountCapitalDepositSubAddress], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SpotSubAccountCapitalDepositSubAddressRes](api.client.c, url, GET)
}

// binance SPOT子母账户接口 SpotSubAccountCapitalDepositSubHisrec 查询子账户充值历史 (适用主账户)
func (client *SpotRestClient) NewSpotSubAccountCapitalDepositSubHisrec() *SpotSubAccountCapitalDepositSubHisrecApi {
	return &SpotSubAccountCapitalDepositSubHisrecApi{
		client: client,
		req:    &SpotSubAccountCapitalDepositSubHisrecReq{},
	}
}
func (api *SpotSubAccountCapitalDepositSubHisrecApi) Do() (*SpotSubAccountCapitalDepositSubHisrecRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotSubAccountCapitalDepositSubHisrec], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SpotSubAccountCapitalDepositSubHisrecRes](api.client.c, url, GET)
}

func (client *SpotRestClient) NewSpotSubAccountFuturesPositionRisk() *SpotSubAccountFuturesPositionRiskApi {
	return &SpotSubAccountFuturesPositionRiskApi{
		client: client,
		req:    &SpotSubAccountFuturesPositionRiskReq{},
	}
}
func (api *SpotSubAccountFuturesPositionRiskApi) Do() (*SpotSubAccountFuturesPositionRiskRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotSubAccountFuturesPositionRisk], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SpotSubAccountFuturesPositionRiskRes](api.client.c, url, GET)
}

// binance SPOT子母账户接口 SpotSubAccountSpotSummary 查询子账户汇总账户信息 (适用主账户)
func (api *SpotRestClient) NewSpotSubAccountSpotSummary() *SpotSubAccountSpotSummaryApi {
	return &SpotSubAccountSpotSummaryApi{
		client: api,
		req:    &SpotSubAccountSpotSummaryReq{},
	}
}
func (api *SpotSubAccountSpotSummaryApi) Do() (*SpotSubAccountSpotSummaryRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotSubAccountSpotSummary], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SpotSubAccountSpotSummaryRes](api.client.c, url, GET)
}
