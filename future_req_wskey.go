package mybinanceapi

// listenKey相关
type FutureListenKeyPostReq struct{}

type FutureListenKeyPostApi struct {
	client *FutureRestClient
	req    *FutureListenKeyPostReq
}

type FutureListenKeyPutReq struct{}

type FutureListenKeyPutApi struct {
	client *FutureRestClient
	req    *FutureListenKeyPutReq
}

type FutureListenKeyDeleteReq struct{}

type FutureListenKeyDeleteApi struct {
	client *FutureRestClient
	req    *FutureListenKeyDeleteReq
}
