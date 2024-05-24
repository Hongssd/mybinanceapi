package mybinanceapi

type SpotPingReq struct {
}
type SpotPingApi struct {
	client *SpotRestClient
	req    *SpotPingReq
}

type SpotServerTimeReq struct {
}
type SpotServerTimeApi struct {
	client *SpotRestClient
	req    *SpotServerTimeReq
}

type SpotExchangeInfoReq struct {
}
type SpotExchangeInfoApi struct {
	client *SpotRestClient
	req    *SpotExchangeInfoReq
}
