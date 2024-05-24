package mybinanceapi

type SwapListenKeyPostReq struct{}
type SwapListenKeyPostApi struct {
	client *SwapRestClient
	req    *SwapListenKeyPostReq
}

type SwapListenKeyPutReq struct{}
type SwapListenKeyPutApi struct {
	client *SwapRestClient
	req    *SwapListenKeyPutReq
}

type SwapListenKeyDeleteReq struct{}
type SwapListenKeyDeleteApi struct {
	client *SwapRestClient
	req    *SwapListenKeyDeleteReq
}
