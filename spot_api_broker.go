package mybinanceapi

import "time"

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
