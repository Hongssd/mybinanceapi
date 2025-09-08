package mybinanceapi

import "time"

// 现货账户接口
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

// binance SPOT spotAssetGetFundingAsset rest资金账户 (USER_DATA)
func (client *SpotRestClient) NewSpotAssetGetFundingAsset() *SpotAssetGetFundingAssetApi {
	return &SpotAssetGetFundingAssetApi{
		client: client,
		req:    &SpotAssetGetFundingAssetReq{},
	}
}
func (api *SpotAssetGetFundingAssetApi) Do() (*SpotAssetGetFundingAssetRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotAssetGetFundingAsset], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SpotAssetGetFundingAssetRes](api.client.c, url, POST)
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

// binance SPOT SpotCapitalDepositAddress rest获取充值地址 (USER_DATA)
func (client *SpotRestClient) NewSpotCapitalDepositAddress() *SpotCapitalDepositAddressApi {
	return &SpotCapitalDepositAddressApi{
		client: client,
		req:    &SpotCapitalDepositAddressReq{},
	}
}
func (api *SpotCapitalDepositAddressApi) Do() (*SpotCapitalDepositAddressRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotCapitalDepositAddress], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SpotCapitalDepositAddressRes](api.client.c, url, GET)
}

// binance SPOT SpotCapitalDepositHisrec rest获取充值历史 (USER_DATA)
func (client *SpotRestClient) NewSpotCapitalDepositHisrec() *SpotCapitalDepositHisrecApi {
	return &SpotCapitalDepositHisrecApi{
		client: client,
		req:    &SpotCapitalDepositHisrecReq{},
	}
}
func (api *SpotCapitalDepositHisrecApi) Do() (*SpotCapitalDepositHisrecRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotCapitalDepositHisrec], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SpotCapitalDepositHisrecRes](api.client.c, url, GET)
}

// binance SPOT SpotAssetDustBtc rest获取可以转换成BNB的小额资产 (USER_DATA)
func (client *SpotRestClient) NewSpotAssetDustBtc() *SpotAssetDustBtcApi {
	return &SpotAssetDustBtcApi{
		client: client,
		req:    &SpotAssetDustBtcReq{},
	}
}
func (api *SpotAssetDustBtcApi) Do() (*SpotAssetDustBtcRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotAssetDustBtc], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SpotAssetDustBtcRes](api.client.c, url, POST)
}

// binance SPOT SpotAssetDust rest小额资产转换 (USER_DATA)
func (client *SpotRestClient) NewSpotAssetDust() *SpotAssetDustApi {
	return &SpotAssetDustApi{
		client: client,
		req:    &SpotAssetDustReq{},
	}
}
func (api *SpotAssetDustApi) Do() (*SpotAssetDustRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotAssetDust], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SpotAssetDustRes](api.client.c, url, POST)
}

// binance SPOT SpotMarginExchangeSmallLiabilityGet rest查询可小额负债转换的资产 (USER_DATA)
func (client *SpotRestClient) NewSpotMarginExchangeSmallLiabilityGet() *SpotMarginExchangeSmallLiabilityGetApi {
	return &SpotMarginExchangeSmallLiabilityGetApi{
		client: client,
		req:    &SpotMarginExchangeSmallLiabilityGetReq{},
	}
}
func (api *SpotMarginExchangeSmallLiabilityGetApi) Do() (*SpotMarginExchangeSmallLiabilityGetRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotMarginExchangeSmallLiabilityGet], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SpotMarginExchangeSmallLiabilityGetRes](api.client.c, url, GET)
}

//binance SPOT SpotMarginExchangeSmallLiabilityPost rest全仓杠杆小额负债转换 (MARGIN)
func (client *SpotRestClient) NewSpotMarginExchangeSmallLiabilityPost() *SpotMarginExchangeSmallLiabilityPostApi {
	return &SpotMarginExchangeSmallLiabilityPostApi{
		client: client,
		req:    &SpotMarginExchangeSmallLiabilityPostReq{},
	}
}
func (api *SpotMarginExchangeSmallLiabilityPostApi) Do() (*SpotMarginExchangeSmallLiabilityPostRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotMarginExchangeSmallLiabilityPost], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SpotMarginExchangeSmallLiabilityPostRes](api.client.c, url, POST)
}

// binance SPOT SpotBrokerRebateRecentRecord rest查询经纪商返佣近期记录（现货）
func (client *SpotRestClient) NewSpotBrokerRebateRecentRecord() *SpotBrokerRebateRecentRecordApi {
	return &SpotBrokerRebateRecentRecordApi{
		client: client,
		req:    &SpotBrokerRebateRecentRecordReq{},
	}
}
func (api *SpotBrokerRebateRecentRecordApi) Do() (*SpotBrokerRebateRecentRecordRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotBrokerRebateRecentRecord], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SpotBrokerRebateRecentRecordRes](api.client.c, url, GET)
}

// binance SPOT SpotBrokerRebateFuturesRecentRecord rest查询经纪商返佣近期记录（合约）
func (client *SpotRestClient) NewSpotBrokerRebateFuturesRecentRecord() *SpotBrokerRebateFuturesRecentRecordApi {
	return &SpotBrokerRebateFuturesRecentRecordApi{
		client: client,
		req:    &SpotBrokerRebateFuturesRecentRecordReq{},
	}
}
func (api *SpotBrokerRebateFuturesRecentRecordApi) Do() (*SpotBrokerRebateFuturesRecentRecordRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotBrokerRebateFuturesRecentRecord], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SpotBrokerRebateFuturesRecentRecordRes](api.client.c, url, GET)
}

// binance SPOT SpotApiReferralIfNewUser rest查询返佣资格
func (client *SpotRestClient) NewSpotApiReferralIfNewUser() *SpotApiReferralIfNewUserApi {
	return &SpotApiReferralIfNewUserApi{
		client: client,
		req:    &SpotApiReferralIfNewUserReq{},
	}
}
func (api *SpotApiReferralIfNewUserApi) Do() (*SpotApiReferralIfNewUserRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotApiReferralIfNewUser], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SpotApiReferralIfNewUserRes](api.client.c, url, GET)
}









