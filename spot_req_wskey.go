package mybinanceapi

type SpotUserDataStreamPostReq struct{}
type SpotUserDataStreamPostApi struct {
	client *SpotRestClient
	req    *SpotUserDataStreamPostReq
}

type SpotUserDataStreamPutReq struct {
	ListenKey *string `json:"listenKey"` //YES
}
type SpotUserDataStreamPutApi struct {
	client *SpotRestClient
	req    *SpotUserDataStreamPutReq
}

func (api *SpotUserDataStreamPutApi) ListenKey(ListenKey string) *SpotUserDataStreamPutApi {
	api.req.ListenKey = GetPointer(ListenKey)
	return api
}

type SpotUserDataStreamDeleteReq struct {
	ListenKey *string `json:"listenKey"` //YES
}
type SpotUserDataStreamDeleteApi struct {
	client *SpotRestClient
	req    *SpotUserDataStreamDeleteReq
}

func (api *SpotUserDataStreamDeleteApi) ListenKey(ListenKey string) *SpotUserDataStreamDeleteApi {
	api.req.ListenKey = GetPointer(ListenKey)
	return api
}

type SpotMarginUserDataStreamPostReq struct{}
type SpotMarginUserDataStreamPostApi struct {
	client *SpotRestClient
	req    *SpotMarginUserDataStreamPostReq
}

type SpotMarginUserDataStreamPutReq struct {
	ListenKey *string `json:"listenKey"` //YES
}
type SpotMarginUserDataStreamPutApi struct {
	client *SpotRestClient
	req    *SpotMarginUserDataStreamPutReq
}

func (api *SpotMarginUserDataStreamPutApi) ListenKey(ListenKey string) *SpotMarginUserDataStreamPutApi {
	api.req.ListenKey = GetPointer(ListenKey)
	return api
}

type SpotMarginUserDataStreamDeleteReq struct {
	ListenKey *string `json:"listenKey"` //YES
}
type SpotMarginUserDataStreamDeleteApi struct {
	client *SpotRestClient
	req    *SpotMarginUserDataStreamDeleteReq
}

func (api *SpotMarginUserDataStreamDeleteApi) ListenKey(ListenKey string) *SpotMarginUserDataStreamDeleteApi {
	api.req.ListenKey = GetPointer(ListenKey)
	return api
}

type SpotMarginIsolatedUserDataStreamPostReq struct {
	Symbol *string `json:"symbol"` //YES
}
type SpotMarginIsolatedUserDataStreamPostApi struct {
	client *SpotRestClient
	req    *SpotMarginIsolatedUserDataStreamPostReq
}

func (api *SpotMarginIsolatedUserDataStreamPostApi) Symbol(Symbol string) *SpotMarginIsolatedUserDataStreamPostApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}

type SpotMarginIsolatedUserDataStreamPutReq struct {
	Symbol    *string `json:"symbol"`    //YES
	ListenKey *string `json:"listenKey"` //YES
}
type SpotMarginIsolatedUserDataStreamPutApi struct {
	client *SpotRestClient
	req    *SpotMarginIsolatedUserDataStreamPutReq
}

func (api *SpotMarginIsolatedUserDataStreamPutApi) Symbol(Symbol string) *SpotMarginIsolatedUserDataStreamPutApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *SpotMarginIsolatedUserDataStreamPutApi) ListenKey(ListenKey string) *SpotMarginIsolatedUserDataStreamPutApi {
	api.req.ListenKey = GetPointer(ListenKey)
	return api
}

type SpotMarginIsolatedUserDataStreamDeleteReq struct {
	Symbol    *string `json:"symbol"`    //YES
	ListenKey *string `json:"listenKey"` //YES
}
type SpotMarginIsolatedUserDataStreamDeleteApi struct {
	client *SpotRestClient
	req    *SpotMarginIsolatedUserDataStreamDeleteReq
}

func (api *SpotMarginIsolatedUserDataStreamDeleteApi) Symbol(Symbol string) *SpotMarginIsolatedUserDataStreamDeleteApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *SpotMarginIsolatedUserDataStreamDeleteApi) ListenKey(ListenKey string) *SpotMarginIsolatedUserDataStreamDeleteApi {
	api.req.ListenKey = GetPointer(ListenKey)
	return api
}
