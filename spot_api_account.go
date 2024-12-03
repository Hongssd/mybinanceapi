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
