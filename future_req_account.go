package mybinanceapi

// FutureAccount
type FutureAccountReq struct {
	Recvwindow *int64 `json:"recvWindow"`
	Timestamp  *int64 `json:"timestamp"`
}
type FutureAccountApi struct {
	client *FutureRestClient
	req    *FutureAccountReq
}

func (api *FutureAccountApi) Recvwindow(Recvwindow int64) *FutureAccountApi {
	api.req.Recvwindow = GetPointer(Recvwindow)
	return api
}
func (api *FutureAccountApi) Timestamp(Timestamp int64) *FutureAccountApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

// FuturePositionSideDualGet
type FuturePositionSideDualGetReq struct {
	Recvwindow *int64 `json:"recvWindow"`
	Timestamp  *int64 `json:"timestamp"`
}
type FuturePositionSideDualGetApi struct {
	client *FutureRestClient
	req    *FuturePositionSideDualGetReq
}

func (api *FuturePositionSideDualGetApi) Recvwindow(Recvwindow int64) *FuturePositionSideDualGetApi {
	api.req.Recvwindow = GetPointer(Recvwindow)
	return api
}
func (api *FuturePositionSideDualGetApi) Timestamp(Timestamp int64) *FuturePositionSideDualGetApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

// FutureMultiAssetsMarginGet
type FutureMultiAssetsMarginGetReq struct {
	Recvwindow *int64 `json:"recvWindow"`
	Timestamp  *int64 `json:"timestamp"`
}
type FutureMultiAssetsMarginGetApi struct {
	client *FutureRestClient
	req    *FutureMultiAssetsMarginGetReq
}

func (api *FutureMultiAssetsMarginGetApi) Recvwindow(Recvwindow int64) *FutureMultiAssetsMarginGetApi {
	api.req.Recvwindow = GetPointer(Recvwindow)
	return api
}
func (api *FutureMultiAssetsMarginGetApi) Timestamp(Timestamp int64) *FutureMultiAssetsMarginGetApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

// FuturePositionSideDualPost
type FuturePositionSideDualPostReq struct {
	DualSidePosition *string `json:"dualSidePosition"` //YES "true": 双向持仓模式；"false": 单向持仓模式
	Recvwindow       *int64  `json:"recvWindow"`
	Timestamp        *int64  `json:"timestamp"`
}

type FuturePositionSideDualPostApi struct {
	client *FutureRestClient
	req    *FuturePositionSideDualPostReq
}

func (api *FuturePositionSideDualPostApi) DualSidePosition(DualSidePosition string) *FuturePositionSideDualPostApi {
	api.req.DualSidePosition = GetPointer(DualSidePosition)
	return api
}
func (api *FuturePositionSideDualPostApi) Recvwindow(Recvwindow int64) *FuturePositionSideDualPostApi {
	api.req.Recvwindow = GetPointer(Recvwindow)
	return api
}
func (api *FuturePositionSideDualPostApi) Timestamp(Timestamp int64) *FuturePositionSideDualPostApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

// FutureMultiAssetsMarginPost
type FutureMultiAssetsMarginPostReq struct {
	MultiAssetsMargin *string `json:"multiAssetsMargin"` //YES "true": 联合保证金模式开启；"false": 联合保证金模式关闭
	Recvwindow        *int64  `json:"recvWindow"`
	Timestamp         *int64  `json:"timestamp"`
}

type FutureMultiAssetsMarginPostApi struct {
	client *FutureRestClient
	req    *FutureMultiAssetsMarginPostReq
}

func (api *FutureMultiAssetsMarginPostApi) MultiAssetsMargin(MultiAssetsMargin string) *FutureMultiAssetsMarginPostApi {
	api.req.MultiAssetsMargin = GetPointer(MultiAssetsMargin)
	return api
}
func (api *FutureMultiAssetsMarginPostApi) Recvwindow(Recvwindow int64) *FutureMultiAssetsMarginPostApi {
	api.req.Recvwindow = GetPointer(Recvwindow)
	return api
}
func (api *FutureMultiAssetsMarginPostApi) Timestamp(Timestamp int64) *FutureMultiAssetsMarginPostApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

// FutureLeverage
type FutureLeverageReq struct {
	Symbol     *string `json:"symbol"`   //YES	交易对
	Leverage   *int64  `json:"leverage"` //YES	目标杠杆倍数：1 到 125 整数
	Recvwindow *int64  `json:"recvWindow"`
	Timestamp  *int64  `json:"timestamp"`
}

type FutureLeverageApi struct {
	client *FutureRestClient
	req    *FutureLeverageReq
}

func (api *FutureLeverageApi) Symbol(Symbol string) *FutureLeverageApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *FutureLeverageApi) Leverage(Leverage int64) *FutureLeverageApi {
	api.req.Leverage = GetPointer(Leverage)
	return api
}
func (api *FutureLeverageApi) Recvwindow(Recvwindow int64) *FutureLeverageApi {
	api.req.Recvwindow = GetPointer(Recvwindow)
	return api
}
func (api *FutureLeverageApi) Timestamp(Timestamp int64) *FutureLeverageApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

// FutureMarginType
type FutureMarginTypeReq struct {
	Symbol     *string `json:"symbol"`     //YES	交易对
	MarginType *string `json:"marginType"` //YES	保证金模式 ISOLATED(逐仓), CROSSED(全仓)
	Recvwindow *int64  `json:"recvWindow"`
	Timestamp  *int64  `json:"timestamp"`
}

type FutureMarginTypeApi struct {
	client *FutureRestClient
	req    *FutureMarginTypeReq
}

func (api *FutureMarginTypeApi) Symbol(Symbol string) *FutureMarginTypeApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *FutureMarginTypeApi) MarginType(MarginType string) *FutureMarginTypeApi {
	api.req.MarginType = GetPointer(MarginType)
	return api
}
func (api *FutureMarginTypeApi) Recvwindow(Recvwindow int64) *FutureMarginTypeApi {
	api.req.Recvwindow = GetPointer(Recvwindow)
	return api
}
func (api *FutureMarginTypeApi) Timestamp(Timestamp int64) *FutureMarginTypeApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

// FutureLeverageBracket
type FutureLeverageBracketReq struct {
	Symbol     *string `json:"symbol"` //No	交易对
	Recvwindow *int64  `json:"recvWindow"`
	Timestamp  *int64  `json:"timestamp"`
}

type FutureLeverageBracketApi struct {
	client *FutureRestClient
	req    *FutureLeverageBracketReq
}

func (api *FutureLeverageBracketApi) Symbol(Symbol string) *FutureLeverageBracketApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *FutureLeverageBracketApi) Recvwindow(Recvwindow int64) *FutureLeverageBracketApi {
	api.req.Recvwindow = GetPointer(Recvwindow)
	return api
}
func (api *FutureLeverageBracketApi) Timestamp(Timestamp int64) *FutureLeverageBracketApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}
