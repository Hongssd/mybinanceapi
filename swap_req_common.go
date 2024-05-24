package mybinanceapi

type SwapPingReq struct {
}
type SwapPingApi struct {
	client *SwapRestClient
	req    *SwapPingReq
}

type SwapServerTimeReq struct {
}
type SwapServerTimeApi struct {
	client *SwapRestClient
	req    *SwapServerTimeReq
}

type SwapExchangeInfoReq struct {
}
type SwapExchangeInfoApi struct {
	client *SwapRestClient
	req    *SwapExchangeInfoReq
}
