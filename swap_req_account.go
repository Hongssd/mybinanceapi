package mybinanceapi

type SwapAccountReq struct {
	RecvWindow *int64 `json:"recvWindow"`
	Timestamp  *int64 `json:"timestamp"`
}
type SwapAccountApi struct {
	client *SwapRestClient
	req    *SwapAccountReq
}

func (api *SwapAccountApi) RecvWindow(RecvWindow int64) *SwapAccountApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *SwapAccountApi) Timestamp(Timestamp int64) *SwapAccountApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SwapPositionSideDualGetReq struct {
	RecvWindow *int64 `json:"recvWindow"`
	Timestamp  *int64 `json:"timestamp"`
}
type SwapPositionSideDualGetApi struct {
	client *SwapRestClient
	req    *SwapPositionSideDualGetReq
}

func (api *SwapPositionSideDualGetApi) RecvWindow(RecvWindow int64) *SwapPositionSideDualGetApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *SwapPositionSideDualGetApi) Timestamp(Timestamp int64) *SwapPositionSideDualGetApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SwapLeverageBracketReq struct {
	Symbol     *string `json:"symbol"` //No	交易对
	RecvWindow *int64  `json:"recvWindow"`
	Timestamp  *int64  `json:"timestamp"`
}
type SwapLeverageBracketApi struct {
	client *SwapRestClient
	req    *SwapLeverageBracketReq
}

func (api *SwapLeverageBracketApi) Symbol(Symbol string) *SwapLeverageBracketApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *SwapLeverageBracketApi) RecvWindow(RecvWindow int64) *SwapLeverageBracketApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *SwapLeverageBracketApi) Timestamp(Timestamp int64) *SwapLeverageBracketApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SwapPositionSideDualPostReq struct {
	DualSidePosition *string `json:"dualSidePosition"` //YES "true": 双向持仓模式；"false": 单向持仓模式
	RecvWindow       *int64  `json:"recvWindow"`
	Timestamp        *int64  `json:"timestamp"`
}
type SwapPositionSideDualPostApi struct {
	client *SwapRestClient
	req    *SwapPositionSideDualPostReq
}

func (api *SwapPositionSideDualPostApi) DualSidePosition(DualSidePosition string) *SwapPositionSideDualPostApi {
	api.req.DualSidePosition = GetPointer(DualSidePosition)
	return api
}
func (api *SwapPositionSideDualPostApi) RecvWindow(RecvWindow int64) *SwapPositionSideDualPostApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *SwapPositionSideDualPostApi) Timestamp(Timestamp int64) *SwapPositionSideDualPostApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SwapLeverageReq struct {
	Symbol     *string `json:"symbol"`   //YES	交易对
	Leverage   *int64  `json:"leverage"` //YES	目标杠杆倍数：1 到 125 整数
	RecvWindow *int64  `json:"recvWindow"`
	Timestamp  *int64  `json:"timestamp"`
}
type SwapLeverageApi struct {
	client *SwapRestClient
	req    *SwapLeverageReq
}

func (api *SwapLeverageApi) Symbol(Symbol string) *SwapLeverageApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *SwapLeverageApi) Leverage(Leverage int64) *SwapLeverageApi {
	api.req.Leverage = GetPointer(Leverage)
	return api
}
func (api *SwapLeverageApi) RecvWindow(RecvWindow int64) *SwapLeverageApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *SwapLeverageApi) Timestamp(Timestamp int64) *SwapLeverageApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SwapMarginTypeReq struct {
	Symbol     *string `json:"symbol"`     //YES	交易对
	MarginType *string `json:"marginType"` //YES	保证金模式 ISOLATED(逐仓), CROSSED(全仓)
	RecvWindow *int64  `json:"recvWindow"`
	Timestamp  *int64  `json:"timestamp"`
}
type SwapMarginTypeApi struct {
	client *SwapRestClient
	req    *SwapMarginTypeReq
}

func (api *SwapMarginTypeApi) Symbol(Symbol string) *SwapMarginTypeApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *SwapMarginTypeApi) MarginType(MarginType string) *SwapMarginTypeApi {
	api.req.MarginType = GetPointer(MarginType)
	return api
}
func (api *SwapMarginTypeApi) RecvWindow(RecvWindow int64) *SwapMarginTypeApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *SwapMarginTypeApi) Timestamp(Timestamp int64) *SwapMarginTypeApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}
