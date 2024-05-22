package mybinanceapi

type FuturePingReq struct {
}
type FuturePingApi struct {
	client *FutureRestClient
	req    *FuturePingReq
}

type FutureServerTimeReq struct {
}
type FutureServerTimeApi struct {
	client *FutureRestClient
	req    *FutureServerTimeReq
}

type FutureExchangeInfoReq struct {
}
type FutureExchangeInfoApi struct {
	client *FutureRestClient
	req    *FutureExchangeInfoReq
}
