package mybinanceapi

// FutureAccount
type FutureAccountReq struct {
	RecvWindow *int64 `json:"recvWindow"`
	Timestamp  *int64 `json:"timestamp"`
}
type FutureAccountApi struct {
	client *FutureRestClient
	req    *FutureAccountReq
}

func (api *FutureAccountApi) RecvWindow(RecvWindow int64) *FutureAccountApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *FutureAccountApi) Timestamp(Timestamp int64) *FutureAccountApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

// FuturePositionSideDualGet
type FuturePositionSideDualGetReq struct {
	RecvWindow *int64 `json:"recvWindow"`
	Timestamp  *int64 `json:"timestamp"`
}
type FuturePositionSideDualGetApi struct {
	client *FutureRestClient
	req    *FuturePositionSideDualGetReq
}

func (api *FuturePositionSideDualGetApi) RecvWindow(RecvWindow int64) *FuturePositionSideDualGetApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *FuturePositionSideDualGetApi) Timestamp(Timestamp int64) *FuturePositionSideDualGetApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

// FutureMultiAssetsMarginGet
type FutureMultiAssetsMarginGetReq struct {
	RecvWindow *int64 `json:"recvWindow"`
	Timestamp  *int64 `json:"timestamp"`
}
type FutureMultiAssetsMarginGetApi struct {
	client *FutureRestClient
	req    *FutureMultiAssetsMarginGetReq
}

func (api *FutureMultiAssetsMarginGetApi) RecvWindow(RecvWindow int64) *FutureMultiAssetsMarginGetApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *FutureMultiAssetsMarginGetApi) Timestamp(Timestamp int64) *FutureMultiAssetsMarginGetApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

// FuturePositionSideDualPost
type FuturePositionSideDualPostReq struct {
	DualSidePosition *string `json:"dualSidePosition"` //YES "true": 双向持仓模式；"false": 单向持仓模式
	RecvWindow       *int64  `json:"recvWindow"`
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
func (api *FuturePositionSideDualPostApi) RecvWindow(RecvWindow int64) *FuturePositionSideDualPostApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *FuturePositionSideDualPostApi) Timestamp(Timestamp int64) *FuturePositionSideDualPostApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

// FutureMultiAssetsMarginPost
type FutureMultiAssetsMarginPostReq struct {
	MultiAssetsMargin *string `json:"multiAssetsMargin"` //YES "true": 联合保证金模式开启；"false": 联合保证金模式关闭
	RecvWindow        *int64  `json:"recvWindow"`
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
func (api *FutureMultiAssetsMarginPostApi) RecvWindow(RecvWindow int64) *FutureMultiAssetsMarginPostApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
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
	RecvWindow *int64  `json:"recvWindow"`
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
func (api *FutureLeverageApi) RecvWindow(RecvWindow int64) *FutureLeverageApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
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
	RecvWindow *int64  `json:"recvWindow"`
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
func (api *FutureMarginTypeApi) RecvWindow(RecvWindow int64) *FutureMarginTypeApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *FutureMarginTypeApi) Timestamp(Timestamp int64) *FutureMarginTypeApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

// FutureLeverageBracket
type FutureLeverageBracketReq struct {
	Symbol     *string `json:"symbol"` //No	交易对
	RecvWindow *int64  `json:"recvWindow"`
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
func (api *FutureLeverageBracketApi) RecvWindow(RecvWindow int64) *FutureLeverageBracketApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *FutureLeverageBracketApi) Timestamp(Timestamp int64) *FutureLeverageBracketApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type FuturePositionRiskReq struct {
	Symbol     *string `json:"symbol"` //No	交易对
	RecvWindow *int64  `json:"recvWindow"`
	Timestamp  *int64  `json:"timestamp"`
}

type FuturePositionRiskApi struct {
	client *FutureRestClient
	req    *FuturePositionRiskReq
}

func (api *FuturePositionRiskApi) Symbol(Symbol string) *FuturePositionRiskApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *FuturePositionRiskApi) RecvWindow(RecvWindow int64) *FuturePositionRiskApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *FuturePositionRiskApi) Timestamp(Timestamp int64) *FuturePositionRiskApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

// startTime	LONG	YES	起始时间，ms格式时间戳
// endTime	LONG	YES	结束时间，ms格式时间戳
// recvWindow	LONG	NO
// timestamp	LONG	YES
type FutureIncomeAsynReq struct {
	StartTime  *int64 `json:"startTime"`
	EndTime    *int64 `json:"endTime"`
	RecvWindow *int64 `json:"recvWindow"`
	Timestamp  *int64 `json:"timestamp"`
}

type FutureIncomeAsynApi struct {
	client *FutureRestClient
	req    *FutureIncomeAsynReq
}

func (api *FutureIncomeAsynApi) StartTime(StartTime int64) *FutureIncomeAsynApi {
	api.req.StartTime = GetPointer(StartTime)
	return api
}
func (api *FutureIncomeAsynApi) EndTime(EndTime int64) *FutureIncomeAsynApi {
	api.req.EndTime = GetPointer(EndTime)
	return api
}
func (api *FutureIncomeAsynApi) RecvWindow(RecvWindow int64) *FutureIncomeAsynApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *FutureIncomeAsynApi) Timestamp(Timestamp int64) *FutureIncomeAsynApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

// downloadId	STRING	YES	通过下载id 接口获取
// recvWindow	LONG	NO
// timestamp	LONG	YES
type FutureIncomeAsynIdReq struct {
	DownloadId *string `json:"downloadId"`
	RecvWindow *int64  `json:"recvWindow"`
	Timestamp  *int64  `json:"timestamp"`
}

type FutureIncomeAsynIdApi struct {
	client *FutureRestClient
	req    *FutureIncomeAsynIdReq
}

func (api *FutureIncomeAsynIdApi) DownloadId(DownloadId string) *FutureIncomeAsynIdApi {
	api.req.DownloadId = GetPointer(DownloadId)
	return api
}
func (api *FutureIncomeAsynIdApi) RecvWindow(RecvWindow int64) *FutureIncomeAsynIdApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *FutureIncomeAsynIdApi) Timestamp(Timestamp int64) *FutureIncomeAsynIdApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}
