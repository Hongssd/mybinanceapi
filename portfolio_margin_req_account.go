package mybinanceapi

// 账户接口
type PortfolioMarginBalanceReq struct {
	Asset      *string `json:"asset,omitempty"`      // NO 资产
	RevcWindow *int64  `json:"recvWindow,omitempty"` // NO 接收窗口
	Timestamp  *int64  `json:"timestamp,omitempty"`  // YES 时间戳
}

// 资产
func (api *PortfolioMarginBalanceApi) Asset(asset string) *PortfolioMarginBalanceApi {
	api.req.Asset = GetPointer(asset)
	return api
}

// 接收窗口
func (api *PortfolioMarginBalanceApi) RevcWindow(recvWindow int64) *PortfolioMarginBalanceApi {
	api.req.RevcWindow = GetPointer(recvWindow)
	return api
}

// 时间戳
func (api *PortfolioMarginBalanceApi) Timestamp(timestamp int64) *PortfolioMarginBalanceApi {
	api.req.Timestamp = GetPointer(timestamp)
	return api
}

type PortfolioMarginBalanceApi struct {
	client *PortfolioMarginRestClient
	req    *PortfolioMarginBalanceReq
}

type PortfolioMarginAccountReq struct {
	RevcWindow *int64 `json:"recvWindow,omitempty"` // NO 接收窗口
	Timestamp  *int64 `json:"timestamp,omitempty"`  // YES 时间戳
}

// 接收窗口
func (api *PortfolioMarginAccountApi) RevcWindow(recvWindow int64) *PortfolioMarginAccountApi {
	api.req.RevcWindow = GetPointer(recvWindow)
	return api
}

// 时间戳
func (api *PortfolioMarginAccountApi) Timestamp(timestamp int64) *PortfolioMarginAccountApi {
	api.req.Timestamp = GetPointer(timestamp)
	return api
}

type PortfolioMarginAccountApi struct {
	client *PortfolioMarginRestClient
	req    *PortfolioMarginAccountReq
}

type PortfolioMarginMaxBorrowableReq struct {
	Asset      *string `json:"asset,omitempty"`      // YES 资产
	RevcWindow *int64  `json:"recvWindow,omitempty"` // NO 接收窗口
	Timestamp  *int64  `json:"timestamp,omitempty"`  // YES 时间戳
}

// 资产
func (api *PortfolioMarginMaxBorrowableApi) Asset(asset string) *PortfolioMarginMaxBorrowableApi {
	api.req.Asset = GetPointer(asset)
	return api
}

// 接收窗口
func (api *PortfolioMarginMaxBorrowableApi) RevcWindow(recvWindow int64) *PortfolioMarginMaxBorrowableApi {
	api.req.RevcWindow = GetPointer(recvWindow)
	return api
}

// 时间戳
func (api *PortfolioMarginMaxBorrowableApi) Timestamp(timestamp int64) *PortfolioMarginMaxBorrowableApi {
	api.req.Timestamp = GetPointer(timestamp)
	return api
}

type PortfolioMarginMaxBorrowableApi struct {
	client *PortfolioMarginRestClient
	req    *PortfolioMarginMaxBorrowableReq
}

type PortfolioMarginMaxWithdrawReq struct {
	Asset      *string `json:"asset,omitempty"`      // YES 资产
	RevcWindow *int64  `json:"recvWindow,omitempty"` // NO 接收窗口
	Timestamp  *int64  `json:"timestamp,omitempty"`  // YES 时间戳
}

// 资产
func (api *PortfolioMarginMaxWithdrawApi) Asset(asset string) *PortfolioMarginMaxWithdrawApi {
	api.req.Asset = GetPointer(asset)
	return api
}

// 接收窗口
func (api *PortfolioMarginMaxWithdrawApi) RevcWindow(recvWindow int64) *PortfolioMarginMaxWithdrawApi {
	api.req.RevcWindow = GetPointer(recvWindow)
	return api
}

// 时间戳
func (api *PortfolioMarginMaxWithdrawApi) Timestamp(timestamp int64) *PortfolioMarginMaxWithdrawApi {
	api.req.Timestamp = GetPointer(timestamp)
	return api
}

type PortfolioMarginMaxWithdrawApi struct {
	client *PortfolioMarginRestClient
	req    *PortfolioMarginMaxWithdrawReq
}

type PortfolioMarginSetUmLeverageReq struct {
	Symbol     *string `json:"symbol,omitempty"`     // YES
	Leverage   *int    `json:"leverage,omitempty"`   // YES target initial leverage: int from 1 to 125
	RevcWindow *int64  `json:"recvWindow,omitempty"` // NO
	Timestamp  *int64  `json:"timestamp,omitempty"`  // YES
}

// 交易对
func (api *PortfolioMarginSetUmLeverageApi) Symbol(symbol string) *PortfolioMarginSetUmLeverageApi {
	api.req.Symbol = GetPointer(symbol)
	return api
}

// 杠杆 target initial leverage: int from 1 to 125
func (api *PortfolioMarginSetUmLeverageApi) Leverage(leverage int) *PortfolioMarginSetUmLeverageApi {
	api.req.Leverage = GetPointer(leverage)
	return api
}

// 接收窗口
func (api *PortfolioMarginSetUmLeverageApi) RevcWindow(recvWindow int64) *PortfolioMarginSetUmLeverageApi {
	api.req.RevcWindow = GetPointer(recvWindow)
	return api
}

// 时间戳
func (api *PortfolioMarginSetUmLeverageApi) Timestamp(timestamp int64) *PortfolioMarginSetUmLeverageApi {
	api.req.Timestamp = GetPointer(timestamp)
	return api
}

type PortfolioMarginSetUmLeverageApi struct {
	client *PortfolioMarginRestClient
	req    *PortfolioMarginSetUmLeverageReq
}

type PortfolioMarginSetCmLeverageReq struct {
	Symbol     *string `json:"symbol,omitempty"`     // YES
	Leverage   *int    `json:"leverage,omitempty"`   // YES target initial leverage: int from 1 to 125
	RevcWindow *int64  `json:"recvWindow,omitempty"` // NO
	Timestamp  *int64  `json:"timestamp,omitempty"`  // YES
}

// 交易对
func (api *PortfolioMarginSetCmLeverageApi) Symbol(symbol string) *PortfolioMarginSetCmLeverageApi {
	api.req.Symbol = GetPointer(symbol)
	return api
}

// 杠杆 target initial leverage: int from 1 to 125
func (api *PortfolioMarginSetCmLeverageApi) Leverage(leverage int) *PortfolioMarginSetCmLeverageApi {
	api.req.Leverage = GetPointer(leverage)
	return api
}

// 接收窗口
func (api *PortfolioMarginSetCmLeverageApi) RevcWindow(recvWindow int64) *PortfolioMarginSetCmLeverageApi {
	api.req.RevcWindow = GetPointer(recvWindow)
	return api
}

// 时间戳
func (api *PortfolioMarginSetCmLeverageApi) Timestamp(timestamp int64) *PortfolioMarginSetCmLeverageApi {
	api.req.Timestamp = GetPointer(timestamp)
	return api
}

type PortfolioMarginSetCmLeverageApi struct {
	client *PortfolioMarginRestClient
	req    *PortfolioMarginSetCmLeverageReq
}

type PortfolioMarginUmPositionSideDualPostReq struct {
	DualSidePosition *string `json:"dualSidePosition,omitempty"` // YES true: 双向持仓模式 false: 单向持仓模式
	RevcWindow       *int64  `json:"recvWindow,omitempty"`       // NO
	Timestamp        *int64  `json:"timestamp,omitempty"`        // YES
}

// 双向持仓模式
func (api *PortfolioMarginUmPositionSideDualPostApi) DualSidePosition(dualSidePosition string) *PortfolioMarginUmPositionSideDualPostApi {
	api.req.DualSidePosition = GetPointer(dualSidePosition)
	return api
}

// 接收窗口
func (api *PortfolioMarginUmPositionSideDualPostApi) RevcWindow(recvWindow int64) *PortfolioMarginUmPositionSideDualPostApi {
	api.req.RevcWindow = GetPointer(recvWindow)
	return api
}

// 时间戳
func (api *PortfolioMarginUmPositionSideDualPostApi) Timestamp(timestamp int64) *PortfolioMarginUmPositionSideDualPostApi {
	api.req.Timestamp = GetPointer(timestamp)
	return api
}

type PortfolioMarginUmPositionSideDualPostApi struct {
	client *PortfolioMarginRestClient
	req    *PortfolioMarginUmPositionSideDualPostReq
}

type PortfolioMarginUmPositionSideDualGetReq struct {
	RevcWindow *int64 `json:"recvWindow,omitempty"` // NO
	Timestamp  *int64 `json:"timestamp,omitempty"`  // YES
}

// 接收窗口
func (api *PortfolioMarginUmPositionSideDualGetApi) RevcWindow(recvWindow int64) *PortfolioMarginUmPositionSideDualGetApi {
	api.req.RevcWindow = GetPointer(recvWindow)
	return api
}

// 时间戳
func (api *PortfolioMarginUmPositionSideDualGetApi) Timestamp(timestamp int64) *PortfolioMarginUmPositionSideDualGetApi {
	api.req.Timestamp = GetPointer(timestamp)
	return api
}

type PortfolioMarginUmPositionSideDualGetApi struct {
	client *PortfolioMarginRestClient
	req    *PortfolioMarginUmPositionSideDualGetReq
}

type PortfolioMarginCmPositionSideDualPostReq struct {
	DualSidePosition *string `json:"dualSidePosition,omitempty"` // YES true: 双向持仓模式 false: 单向持仓模式
	RevcWindow       *int64  `json:"recvWindow,omitempty"`       // NO
	Timestamp        *int64  `json:"timestamp,omitempty"`        // YES
}

// 双向持仓模式
func (api *PortfolioMarginCmPositionSideDualPostApi) DualSidePosition(dualSidePosition string) *PortfolioMarginCmPositionSideDualPostApi {
	api.req.DualSidePosition = GetPointer(dualSidePosition)
	return api
}

// 接收窗口
func (api *PortfolioMarginCmPositionSideDualPostApi) RevcWindow(recvWindow int64) *PortfolioMarginCmPositionSideDualPostApi {
	api.req.RevcWindow = GetPointer(recvWindow)
	return api
}

// 时间戳
func (api *PortfolioMarginCmPositionSideDualPostApi) Timestamp(timestamp int64) *PortfolioMarginCmPositionSideDualPostApi {
	api.req.Timestamp = GetPointer(timestamp)
	return api
}

type PortfolioMarginCmPositionSideDualPostApi struct {
	client *PortfolioMarginRestClient
	req    *PortfolioMarginCmPositionSideDualPostReq
}

type PortfolioMarginCmPositionSideDualGetReq struct {
	RevcWindow *int64 `json:"recvWindow,omitempty"` // NO
	Timestamp  *int64 `json:"timestamp,omitempty"`  // YES
}

// 接收窗口
func (api *PortfolioMarginCmPositionSideDualGetApi) RevcWindow(recvWindow int64) *PortfolioMarginCmPositionSideDualGetApi {
	api.req.RevcWindow = GetPointer(recvWindow)
	return api
}

// 时间戳
func (api *PortfolioMarginCmPositionSideDualGetApi) Timestamp(timestamp int64) *PortfolioMarginCmPositionSideDualGetApi {
	api.req.Timestamp = GetPointer(timestamp)
	return api
}

type PortfolioMarginCmPositionSideDualGetApi struct {
	client *PortfolioMarginRestClient
	req    *PortfolioMarginCmPositionSideDualGetReq
}

type PortfolioMarginUmAccountV1Req struct {
	RevcWindow *int64 `json:"recvWindow,omitempty"` // NO
	Timestamp  *int64 `json:"timestamp,omitempty"`  // YES
}

// 接收窗口
func (api *PortfolioMarginUmAccountV1Api) RevcWindow(recvWindow int64) *PortfolioMarginUmAccountV1Api {
	api.req.RevcWindow = GetPointer(recvWindow)
	return api
}

// 时间戳
func (api *PortfolioMarginUmAccountV1Api) Timestamp(timestamp int64) *PortfolioMarginUmAccountV1Api {
	api.req.Timestamp = GetPointer(timestamp)
	return api
}

type PortfolioMarginUmAccountV1Api struct {
	client *PortfolioMarginRestClient
	req    *PortfolioMarginUmAccountV1Req
}

type PortfolioMarginUmAccountV2Req struct {
	RevcWindow *int64 `json:"recvWindow,omitempty"` // NO
	Timestamp  *int64 `json:"timestamp,omitempty"`  // YES
}

// 接收窗口
func (api *PortfolioMarginUmAccountV2Api) RevcWindow(recvWindow int64) *PortfolioMarginUmAccountV2Api {
	api.req.RevcWindow = GetPointer(recvWindow)
	return api
}

// 时间戳
func (api *PortfolioMarginUmAccountV2Api) Timestamp(timestamp int64) *PortfolioMarginUmAccountV2Api {
	api.req.Timestamp = GetPointer(timestamp)
	return api
}

type PortfolioMarginUmAccountV2Api struct {
	client *PortfolioMarginRestClient
	req    *PortfolioMarginUmAccountV2Req
}

type PortfolioMarginCmAccountReq struct {
	RevcWindow *int64 `json:"recvWindow,omitempty"` // NO
	Timestamp  *int64 `json:"timestamp,omitempty"`  // YES
}

// 接收窗口
func (api *PortfolioMarginCmAccountApi) RevcWindow(recvWindow int64) *PortfolioMarginCmAccountApi {
	api.req.RevcWindow = GetPointer(recvWindow)
	return api
}

// 时间戳
func (api *PortfolioMarginCmAccountApi) Timestamp(timestamp int64) *PortfolioMarginCmAccountApi {
	api.req.Timestamp = GetPointer(timestamp)
	return api
}

type PortfolioMarginCmAccountApi struct {
	client *PortfolioMarginRestClient
	req    *PortfolioMarginCmAccountReq
}

type PortfolioMarginUmPositionRiskReq struct {
	Symbol     *string `json:"symbol,omitempty"`     // NO
	RevcWindow *int64  `json:"recvWindow,omitempty"` // NO
	Timestamp  *int64  `json:"timestamp,omitempty"`  // YES
}

// 交易对
func (api *PortfolioMarginUmPositionRiskApi) Symbol(symbol string) *PortfolioMarginUmPositionRiskApi {
	api.req.Symbol = GetPointer(symbol)
	return api
}

// 接收窗口
func (api *PortfolioMarginUmPositionRiskApi) RevcWindow(recvWindow int64) *PortfolioMarginUmPositionRiskApi {
	api.req.RevcWindow = GetPointer(recvWindow)
	return api
}

// 时间戳
func (api *PortfolioMarginUmPositionRiskApi) Timestamp(timestamp int64) *PortfolioMarginUmPositionRiskApi {
	api.req.Timestamp = GetPointer(timestamp)
	return api
}

type PortfolioMarginUmPositionRiskApi struct {
	client *PortfolioMarginRestClient
	req    *PortfolioMarginUmPositionRiskReq
}

type PortfolioMarginCmPositionRiskReq struct {
	MarginAsset *string `json:"marginAsset,omitempty"` // NO
	Pair        *string `json:"pair,omitempty"`        // NO
	RevcWindow  *int64  `json:"recvWindow,omitempty"`  // NO
	Timestamp   *int64  `json:"timestamp,omitempty"`   // YES
}

// MarginAsset NO 保证金资产 marginasset 和 pair 不要同时提供
func (api *PortfolioMarginCmPositionRiskApi) MarginAsset(marginAsset string) *PortfolioMarginCmPositionRiskApi {
	api.req.MarginAsset = GetPointer(marginAsset)
	return api
}

// Pair NO 交易对 marginasset 和 pair 不要同时提供
func (api *PortfolioMarginCmPositionRiskApi) Pair(pair string) *PortfolioMarginCmPositionRiskApi {
	api.req.Pair = GetPointer(pair)
	return api
}

// 接收窗口
func (api *PortfolioMarginCmPositionRiskApi) RevcWindow(recvWindow int64) *PortfolioMarginCmPositionRiskApi {
	api.req.RevcWindow = GetPointer(recvWindow)
	return api
}

// 时间戳
func (api *PortfolioMarginCmPositionRiskApi) Timestamp(timestamp int64) *PortfolioMarginCmPositionRiskApi {
	api.req.Timestamp = GetPointer(timestamp)
	return api
}

type PortfolioMarginCmPositionRiskApi struct {
	client *PortfolioMarginRestClient
	req    *PortfolioMarginCmPositionRiskReq
}

type PortfolioMarginUmCommissionRateReq struct {
	Symbol     *string `json:"symbol,omitempty"`     // YES
	RevcWindow *int64  `json:"recvWindow,omitempty"` // NO
	Timestamp  *int64  `json:"timestamp,omitempty"`  // YES
}

// 交易对
func (api *PortfolioMarginUmCommissionRateApi) Symbol(symbol string) *PortfolioMarginUmCommissionRateApi {
	api.req.Symbol = GetPointer(symbol)
	return api
}

// 接收窗口
func (api *PortfolioMarginUmCommissionRateApi) RevcWindow(recvWindow int64) *PortfolioMarginUmCommissionRateApi {
	api.req.RevcWindow = GetPointer(recvWindow)
	return api
}

// 时间戳
func (api *PortfolioMarginUmCommissionRateApi) Timestamp(timestamp int64) *PortfolioMarginUmCommissionRateApi {
	api.req.Timestamp = GetPointer(timestamp)
	return api
}

type PortfolioMarginUmCommissionRateApi struct {
	client *PortfolioMarginRestClient
	req    *PortfolioMarginUmCommissionRateReq
}

type PortfolioMarginCmCommissionRateReq struct {
	Symbol     *string `json:"symbol,omitempty"`     // YES
	RevcWindow *int64  `json:"recvWindow,omitempty"` // NO
	Timestamp  *int64  `json:"timestamp,omitempty"`  // YES
}

// 交易对
func (api *PortfolioMarginCmCommissionRateApi) Symbol(symbol string) *PortfolioMarginCmCommissionRateApi {
	api.req.Symbol = GetPointer(symbol)
	return api
}

// 接收窗口
func (api *PortfolioMarginCmCommissionRateApi) RevcWindow(recvWindow int64) *PortfolioMarginCmCommissionRateApi {
	api.req.RevcWindow = GetPointer(recvWindow)
	return api
}

// 时间戳
func (api *PortfolioMarginCmCommissionRateApi) Timestamp(timestamp int64) *PortfolioMarginCmCommissionRateApi {
	api.req.Timestamp = GetPointer(timestamp)
	return api
}

type PortfolioMarginCmCommissionRateApi struct {
	client *PortfolioMarginRestClient
	req    *PortfolioMarginCmCommissionRateReq
}

type PortfolioMarginAutoCollectionReq struct {
	RecvWindow *int64 `json:"recvWindow,omitempty"` // NO
	Timestamp  *int64 `json:"timestamp,omitempty"`  // YES
}

// 接收窗口
func (api *PortfolioMarginAutoCollectionApi) RecvWindow(recvWindow int64) *PortfolioMarginAutoCollectionApi {
	api.req.RecvWindow = GetPointer(recvWindow)
	return api
}

// 时间戳
func (api *PortfolioMarginAutoCollectionApi) Timestamp(timestamp int64) *PortfolioMarginAutoCollectionApi {
	api.req.Timestamp = GetPointer(timestamp)
	return api
}

type PortfolioMarginAutoCollectionApi struct {
	client *PortfolioMarginRestClient
	req    *PortfolioMarginAutoCollectionReq
}

type PortfolioMarginAssetCollectionReq struct {
	Asset      *string `json:"asset,omitempty"`      // YES
	RevcWindow *int64  `json:"recvWindow,omitempty"` // NO
	Timestamp  *int64  `json:"timestamp,omitempty"`  // YES
}

// 资产
func (api *PortfolioMarginAssetCollectionApi) Asset(asset string) *PortfolioMarginAssetCollectionApi {
	api.req.Asset = GetPointer(asset)
	return api
}

// 接收窗口
func (api *PortfolioMarginAssetCollectionApi) RevcWindow(recvWindow int64) *PortfolioMarginAssetCollectionApi {
	api.req.RevcWindow = GetPointer(recvWindow)
	return api
}

// 时间戳
func (api *PortfolioMarginAssetCollectionApi) Timestamp(timestamp int64) *PortfolioMarginAssetCollectionApi {
	api.req.Timestamp = GetPointer(timestamp)
	return api
}

type PortfolioMarginAssetCollectionApi struct {
	client *PortfolioMarginRestClient
	req    *PortfolioMarginAssetCollectionReq
}

// symbol	STRING	NO
// incomeType	STRING	NO	TRANSFER, WELCOME_BONUS, REALIZED_PNL, FUNDING_FEE, COMMISSION, INSURANCE_CLEAR, REFERRAL_KICKBACK, COMMISSION_REBATE, API_REBATE, CONTEST_REWARD, CROSS_COLLATERAL_TRANSFER, OPTIONS_PREMIUM_FEE, OPTIONS_SETTLE_PROFIT, INTERNAL_TRANSFER, AUTO_EXCHANGE, DELIVERED_SETTELMENT, COIN_SWAP_DEPOSIT, COIN_SWAP_WITHDRAW, POSITION_LIMIT_INCREASE_FEE
// startTime	LONG	NO	起始时间
// endTime	LONG	NO	结束时间
// limit	INT	NO	默认 100; 最大 1000
// recvWindow	LONG	NO
// timestamp	LONG	YES
type PortfolioMarginUmIncomeReq struct {
	Symbol     *string `json:"symbol,omitempty"`     // NO
	IncomeType *string `json:"incomeType,omitempty"` // NO
	StartTime  *int64  `json:"startTime,omitempty"`  // NO
	EndTime    *int64  `json:"endTime,omitempty"`    // NO
	Limit      *int    `json:"limit,omitempty"`      // NO
	RecvWindow *int64  `json:"recvWindow,omitempty"` // NO
	Timestamp  *int64  `json:"timestamp,omitempty"`  // YES
}
type PortfolioMarginUmIncomeApi struct {
	client *PortfolioMarginRestClient
	req    *PortfolioMarginUmIncomeReq
}

func (api *PortfolioMarginUmIncomeApi) Symbol(symbol string) *PortfolioMarginUmIncomeApi {
	api.req.Symbol = GetPointer(symbol)
	return api
}

func (api *PortfolioMarginUmIncomeApi) IncomeType(incomeType string) *PortfolioMarginUmIncomeApi {
	api.req.IncomeType = GetPointer(incomeType)
	return api
}

func (api *PortfolioMarginUmIncomeApi) StartTime(StartTime int64) *PortfolioMarginUmIncomeApi {
	api.req.StartTime = GetPointer(StartTime)
	return api
}

func (api *PortfolioMarginUmIncomeApi) EndTime(EndTime int64) *PortfolioMarginUmIncomeApi {
	api.req.EndTime = GetPointer(EndTime)
	return api
}

func (api *PortfolioMarginUmIncomeApi) Limit(Limit int) *PortfolioMarginUmIncomeApi {
	api.req.Limit = GetPointer(Limit)
	return api
}

func (api *PortfolioMarginUmIncomeApi) RecvWindow(RecvWindow int64) *PortfolioMarginUmIncomeApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *PortfolioMarginUmIncomeApi) Timestamp(Timestamp int64) *PortfolioMarginUmIncomeApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

// startTime	LONG	YES	起始时间，ms格式时间戳
// endTime	LONG	YES	结束时间，ms格式时间戳
// recvWindow	LONG	NO
// timestamp	LONG	YES
type PortfolioMarginUmIncomeAsynReq struct {
	StartTime  *int64 `json:"startTime"`
	EndTime    *int64 `json:"endTime"`
	RecvWindow *int64 `json:"recvWindow"`
	Timestamp  *int64 `json:"timestamp"`
}

type PortfolioMarginUmIncomeAsynApi struct {
	client *PortfolioMarginRestClient
	req    *PortfolioMarginUmIncomeAsynReq
}

func (api *PortfolioMarginUmIncomeAsynApi) StartTime(StartTime int64) *PortfolioMarginUmIncomeAsynApi {
	api.req.StartTime = GetPointer(StartTime)
	return api
}

func (api *PortfolioMarginUmIncomeAsynApi) EndTime(EndTime int64) *PortfolioMarginUmIncomeAsynApi {
	api.req.EndTime = GetPointer(EndTime)
	return api
}

func (api *PortfolioMarginUmIncomeAsynApi) RecvWindow(RecvWindow int64) *PortfolioMarginUmIncomeAsynApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}

func (api *PortfolioMarginUmIncomeAsynApi) Timestamp(Timestamp int64) *PortfolioMarginUmIncomeAsynApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

// downloadId	STRING	YES	通过下载id 接口获取
// recvWindow	LONG	NO
// timestamp	LONG	YES
type PortfolioMarginUmIncomeAsynIdReq struct {
	DownloadId *string `json:"downloadId"`
	RecvWindow *int64  `json:"recvWindow"`
	Timestamp  *int64  `json:"timestamp"`
}

type PortfolioMarginUmIncomeAsynIdApi struct {
	client *PortfolioMarginRestClient
	req    *PortfolioMarginUmIncomeAsynIdReq
}

func (api *PortfolioMarginUmIncomeAsynIdApi) DownloadId(DownloadId string) *PortfolioMarginUmIncomeAsynIdApi {
	api.req.DownloadId = GetPointer(DownloadId)
	return api
}

func (api *PortfolioMarginUmIncomeAsynIdApi) RecvWindow(RecvWindow int64) *PortfolioMarginUmIncomeAsynIdApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}

func (api *PortfolioMarginUmIncomeAsynIdApi) Timestamp(Timestamp int64) *PortfolioMarginUmIncomeAsynIdApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type PortfolioMarginRepayFuturesNegativeBalanceReq struct {
	RecvWindow *int64 `json:"recvWindow"`
	Timestamp  *int64 `json:"timestamp"`
}

type PortfolioMarginRepayFuturesNegativeBalanceApi struct {
	client *PortfolioMarginRestClient
	req    *PortfolioMarginRepayFuturesNegativeBalanceReq
}

func (api *PortfolioMarginRepayFuturesNegativeBalanceApi) RecvWindow(RecvWindow int64) *PortfolioMarginRepayFuturesNegativeBalanceApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *PortfolioMarginRepayFuturesNegativeBalanceApi) Timestamp(Timestamp int64) *PortfolioMarginRepayFuturesNegativeBalanceApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}
