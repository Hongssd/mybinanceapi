package mybinanceapi

import "time"

// binance SPOT SpotApiVipVipPortalFutureIpWhitelistsGet rest查询VIP专属合约IP白名单 MM低延迟接口
func (client *SpotRestClient) NewSpotApiVipVipPortalFutureIpWhitelistsGet() *SpotApiVipVipPortalFutureIpWhitelistsGetApi {
	return &SpotApiVipVipPortalFutureIpWhitelistsGetApi{
		client: client,
		req:    &SpotApiVipVipPortalFutureIpWhitelistsGetReq{},
	}
}
func (api *SpotApiVipVipPortalFutureIpWhitelistsGetApi) Do() (*SpotApiVipVipPortalFutureIpWhitelistsGetRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(SPOT, api.req, SpotApiMap[SpotApiVipVipPortalFutureIpWhitelistsGet], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[SpotApiVipVipPortalFutureIpWhitelistsGetRes](api.client.c, url, GET)
}

// binance SPOT SpotApiVipVipPortalFutureIpWhitelistsPost rest更新VIP专属合约IP白名单 MM低延迟接口
func (client *SpotRestClient) NewSpotApiVipVipPortalFutureIpWhitelistsPost() *SpotApiVipVipPortalFutureIpWhitelistsPostApi {
	return &SpotApiVipVipPortalFutureIpWhitelistsPostApi{
		client: client,
		req:    &SpotApiVipVipPortalFutureIpWhitelistsPostReq{},
	}
}
func (api *SpotApiVipVipPortalFutureIpWhitelistsPostApi) Do() (*SpotApiVipVipPortalFutureIpWhitelistsPostRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	targetJson := struct {
		IpAddresses *[]string `json:"ipAddresses,omitempty"`
	}{
		IpAddresses: api.req.IpAddresses,
	}
	jsonBody, _ := json.Marshal(targetJson)
	_ = jsonBody
	queryReq := *api.req
	queryReq.IpAddresses = nil

	body, url := binanceHandlerRequestApiWithSecretForUrlRequestAndJsonBody(SPOT, &queryReq, jsonBody, SpotApiMap[SpotApiVipVipPortalFutureIpWhitelistsPost], api.client.c.ApiSecret)
	return binanceCallApiWithSecretForJsonBody[SpotApiVipVipPortalFutureIpWhitelistsPostRes](api.client.c, url, POST, body)
}

// binance SPOT SpotApiVipVipPortalFutureIpWhitelistsDelete rest删除VIP专属合约IP白名单 MM低延迟接口
func (client *SpotRestClient) NewSpotApiVipVipPortalFutureIpWhitelistsDelete() *SpotApiVipVipPortalFutureIpWhitelistsDeleteApi {
	return &SpotApiVipVipPortalFutureIpWhitelistsDeleteApi{
		client: client,
		req:    &SpotApiVipVipPortalFutureIpWhitelistsDeleteReq{},
	}
}
func (api *SpotApiVipVipPortalFutureIpWhitelistsDeleteApi) Do() (*SpotApiVipVipPortalFutureIpWhitelistsDeleteRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	targetJson := struct {
		IpAddresses *[]string `json:"ipAddresses,omitempty"`
	}{
		IpAddresses: api.req.IpAddresses,
	}
	jsonBody, _ := json.Marshal(targetJson)
	_ = jsonBody
	queryReq := *api.req
	queryReq.IpAddresses = nil

	body, url := binanceHandlerRequestApiWithSecretForUrlRequestAndJsonBody(SPOT, &queryReq, jsonBody, SpotApiMap[SpotApiVipVipPortalFutureIpWhitelistsDelete], api.client.c.ApiSecret)
	return binanceCallApiWithSecretForJsonBody[SpotApiVipVipPortalFutureIpWhitelistsDeleteRes](api.client.c, url, DELETE, body)
}
