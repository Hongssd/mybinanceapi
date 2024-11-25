package mybinanceapi

import "time"

// 经纪商子母账号接口
func (client *SpotRestClient) SpotBrokerSubAccountPost() *SpotBrokerSubAccountPostApi {
	return &SpotBrokerSubAccountPostApi{
		client: client,
		req:    &SpotBrokerSubAccountPostReq{},
	}
}
func (api *SpotBrokerSubAccountPostApi) Do() (*SpotBrokerSubAccountPostRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotBrokerSubAccountPost], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SpotBrokerSubAccountPostRes](api.client.c, url, POST)
}

// 经纪商子母账号接口
func (client *SpotRestClient) NewSpotBrokerSubAccountGet() *SpotBrokerSubAccountGetApi {
	return &SpotBrokerSubAccountGetApi{
		client: client,
		req:    &SpotBrokerSubAccountGetReq{},
	}
}

func (api *SpotBrokerSubAccountGetApi) Do() (*SpotBrokerSubAccountGetRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotBrokerSubAccountGet], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SpotBrokerSubAccountGetRes](api.client.c, url, GET)
}

// POST 授权经纪商子账户Futures权限
func (client *SpotRestClient) NewSpotBrokerSubAccountFutures() *SpotBrokerSubAccountFuturesApi {
	return &SpotBrokerSubAccountFuturesApi{
		client: client,
		req:    &SpotBrokerSubAccountFuturesReq{},
	}
}
func (api *SpotBrokerSubAccountFuturesApi) Do() (*SpotBrokerSubAccountFuturesRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotBrokerSubAccountFutures], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SpotBrokerSubAccountFuturesRes](api.client.c, url, POST)
}

// POST 创建经纪商子账号ApiKey
func (client *SpotRestClient) NewSpotBrokerSubAccountApiPost() *SpotBrokerSubAccountApiPostApi {
	return &SpotBrokerSubAccountApiPostApi{
		client: client,
		req:    &SpotBrokerSubAccountApiPostReq{},
	}
}
func (api *SpotBrokerSubAccountApiPostApi) Do() (*SpotBrokerSubAccountApiPostRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotBrokerSubAccountApiPost], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SpotBrokerSubAccountApiPostRes](api.client.c, url, POST)
}

// GET 查询经纪商子账号ApiKey
func (client *SpotRestClient) NewSpotBrokerSubAccountApiGet() *SpotBrokerSubAccountApiGetApi {
	return &SpotBrokerSubAccountApiGetApi{
		client: client,
		req:    &SpotBrokerSubAccountApiGetReq{},
	}
}
func (api *SpotBrokerSubAccountApiGetApi) Do() (*SpotBrokerSubAccountApiGetRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotBrokerSubAccountApiGet], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SpotBrokerSubAccountApiGetRes](api.client.c, url, GET)
}

// DELETE 删除经纪商子账号ApiKey
func (client *SpotRestClient) NewSpotBrokerSubAccountApiDelete() *SpotBrokerSubAccountApiDeleteApi {
	return &SpotBrokerSubAccountApiDeleteApi{
		client: client,
		req:    &SpotBrokerSubAccountApiDeleteReq{},
	}
}
func (api *SpotBrokerSubAccountApiDeleteApi) Do() (*SpotBrokerSubAccountApiDeleteRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotBrokerSubAccountApiDelete], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SpotBrokerSubAccountApiDeleteRes](api.client.c, url, DELETE)
}

// POST 授权经纪商子账户万能划转权限
func (client *SpotRestClient) NewSpotBrokerSubAccountPermissionUniversalTransfer() *SpotBrokerSubAccountPermissionUniversalTransferApi {
	return &SpotBrokerSubAccountPermissionUniversalTransferApi{
		client: client,
		req:    &SpotBrokerSubAccountPermissionUniversalTransferReq{},
	}
}
func (api *SpotBrokerSubAccountPermissionUniversalTransferApi) Do() (*SpotBrokerSubAccountPermissionUniversalTransferRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotBrokerSubAccountPermissionUniversalTransfer], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SpotBrokerSubAccountPermissionUniversalTransferRes](api.client.c, url, POST)
}

// POST 更新经纪商子账户API Key IP白名单
func (client *SpotRestClient) NewSpotBrokerSubAccountApiIpRestriction() *SpotBrokerSubAccountApiIpRestrictionApi {
	return &SpotBrokerSubAccountApiIpRestrictionApi{
		client: client,
		req:    &SpotBrokerSubAccountApiIpRestrictionReq{},
	}
}
func (api *SpotBrokerSubAccountApiIpRestrictionApi) Do() (*SpotBrokerSubAccountApiIpRestrictionPostRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotBrokerSubAccountApiIpRestrictionPost], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SpotBrokerSubAccountApiIpRestrictionPostRes](api.client.c, url, POST)
}

// GET 查询经纪商子账户API Key IP白名单
func (client *SpotRestClient) NewSpotBrokerSubAccountApiIpRestrictionGet() *SpotBrokerSubAccountApiIpRestrictionGetApi {
	return &SpotBrokerSubAccountApiIpRestrictionGetApi{
		client: client,
		req:    &SpotBrokerSubAccountApiIpRestrictionGetReq{},
	}
}
func (api *SpotBrokerSubAccountApiIpRestrictionGetApi) Do() (*SpotBrokerSubAccountApiIpRestrictionGetRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotBrokerSubAccountApiIpRestrictionGet], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SpotBrokerSubAccountApiIpRestrictionGetRes](api.client.c, url, GET)
}

// DELETE 删除经纪商子账户API Key IP白名单
func (client *SpotRestClient) NewSpotBrokerSubAccountApiIpRestrictionDelete() *SpotBrokerSubAccountApiIpRestrictionDeleteApi {
	return &SpotBrokerSubAccountApiIpRestrictionDeleteApi{
		client: client,
		req:    &SpotBrokerSubAccountApiIpRestrictionDeleteReq{},
	}
}
func (api *SpotBrokerSubAccountApiIpRestrictionDeleteApi) Do() (*SpotBrokerSubAccountApiIpRestrictionDeleteRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotBrokerSubAccountApiIpRestrictionDelete], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SpotBrokerSubAccountApiIpRestrictionDeleteRes](api.client.c, url, DELETE)
}

// GET 查询经纪商子账户充值历史
func (client *SpotRestClient) NewSpotBrokerSubAccountDepositHist() *SpotBrokerSubAccountDepositHistApi {
	return &SpotBrokerSubAccountDepositHistApi{
		client: client,
		req:    &SpotBrokerSubAccountDepositHistReq{},
	}
}
func (api *SpotBrokerSubAccountDepositHistApi) Do() (*SpotBrokerSubAccountDepositHistRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotBrokerSubAccountDepositHist], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SpotBrokerSubAccountDepositHistRes](api.client.c, url, GET)
}

// POST 经纪商子账户万能划转
// CAUTION!!!
// You need to enable "internal transfer" option for the api key which requests this endpoint.
// Transfer from master account if fromId not sent.
// Transfer to master account if toId not sent.
// Transfer between futures acount is not supported.
func (client *SpotRestClient) NewSpotBrokerUniversalTransferPost() *SpotBrokerUniversalTransferPostApi {
	return &SpotBrokerUniversalTransferPostApi{
		client: client,
		req:    &SpotBrokerUniversalTransferPostReq{},
	}
}
func (api *SpotBrokerUniversalTransferPostApi) Do() (*SpotBrokerUniversalTransferPostRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotBrokerUniversalTransferPost], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SpotBrokerUniversalTransferPostRes](api.client.c, url, POST)
}

// GET 查询经纪商子账户万能划转历史
func (client *SpotRestClient) NewSpotBrokerUniversalTransferGet() *SpotBrokerUniversalTransferGetApi {
	return &SpotBrokerUniversalTransferGetApi{
		client: client,
		req:    &SpotBrokerUniversalTransferGetReq{},
	}
}
func (api *SpotBrokerUniversalTransferGetApi) Do() (*SpotBrokerUniversalTransferGetRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotBrokerUniversalTransferGet], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SpotBrokerUniversalTransferGetRes](api.client.c, url, GET)
}
