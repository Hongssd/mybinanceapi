package mybinanceapi

import "github.com/shopspring/decimal"

type SpotSubAccountListReq struct {
	Email      *string `json:"email"`
	Isfreeze   *string `json:"isFreeze"`
	Page       *int    `json:"page"`
	Limit      *int    `json:"limit"`
	Recvwindow *int64  `json:"recvWindow"`
	Timestamp  *int64  `json:"timestamp"`
}

type SpotSubAccountListApi struct {
	SpotRestClient
	req *SpotSubAccountListReq
}

func (api *SpotSubAccountListApi) Email(Email string) *SpotSubAccountListApi {
	api.req.Email = GetPointer(Email)
	return api
}

func (api *SpotSubAccountListApi) Isfreeze(Isfreeze bool) *SpotSubAccountListApi {
	if Isfreeze {
		api.req.Isfreeze = GetPointer("true")
	} else {
		api.req.Isfreeze = GetPointer("false")
	}
	return api
}

func (api *SpotSubAccountListApi) Page(Page int) *SpotSubAccountListApi {
	api.req.Page = GetPointer(Page)
	return api
}

func (api *SpotSubAccountListApi) Limit(Limit int) *SpotSubAccountListApi {
	api.req.Limit = GetPointer(Limit)
	return api
}

func (api *SpotSubAccountListApi) Recvwindow(Recvwindow int64) *SpotSubAccountListApi {
	api.req.Recvwindow = GetPointer(Recvwindow)
	return api
}

func (api *SpotSubAccountListApi) Timestamp(Timestamp int64) *SpotSubAccountListApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SpotSubAccountVirtualSubAccountReq struct {
	SubAccountString *string `json:"subAccountString"`
	Recvwindow       *int64  `json:"recvWindow"`
	Timestamp        *int64  `json:"timestamp"`
}
type SpotSubAccountVirtualSubAccountApi struct {
	SpotRestClient
	req *SpotSubAccountVirtualSubAccountReq
}

func (api *SpotSubAccountVirtualSubAccountApi) SubAccountString(SubAccountString string) *SpotSubAccountVirtualSubAccountApi {
	api.req.SubAccountString = GetPointer(SubAccountString)
	return api
}
func (api *SpotSubAccountVirtualSubAccountApi) Recvwindow(Recvwindow int64) *SpotSubAccountVirtualSubAccountApi {
	api.req.Recvwindow = GetPointer(Recvwindow)
	return api
}
func (api *SpotSubAccountVirtualSubAccountApi) Timestamp(Timestamp int64) *SpotSubAccountVirtualSubAccountApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SpotSubAccountUniversalTransferReq struct {
	FromEmail       *string          `json:"fromEmail"`       //NO
	ToEmail         *string          `json:"toEmail"`         //NO
	FromAccountType *string          `json:"fromAccountType"` //YES	"SPOT","USDT_FUTURE","COIN_FUTURE","MARGIN"(Cross),"ISOLATED_MARGIN"
	ToAccountType   *string          `json:"toAccountType"`   //YES	"SPOT","USDT_FUTURE","COIN_FUTURE","MARGIN"(Cross),"ISOLATED_MARGIN"
	ClientTranId    *string          `json:"clientTranId"`    //NO	不可重复
	Symbol          *string          `json:"symbol"`          //NO	仅在ISOLATED_MARGIN类型下使用
	Asset           *string          `json:"asset"`           //YES
	Amount          *decimal.Decimal `json:"amount"`          //YES
	Recvwindow      *int64           `json:"recvWindow"`
	Timestamp       *int64           `json:"timestamp"`
}
type SpotSubAccountUniversalTransferApi struct {
	SpotRestClient
	req *SpotSubAccountUniversalTransferReq
}

func (api *SpotSubAccountUniversalTransferApi) FromEmail(FromEmail string) *SpotSubAccountUniversalTransferApi {
	api.req.FromEmail = GetPointer(FromEmail)
	return api
}
func (api *SpotSubAccountUniversalTransferApi) ToEmail(ToEmail string) *SpotSubAccountUniversalTransferApi {
	api.req.ToEmail = GetPointer(ToEmail)
	return api
}
func (api *SpotSubAccountUniversalTransferApi) FromAccountType(FromAccountType string) *SpotSubAccountUniversalTransferApi {
	api.req.FromAccountType = GetPointer(FromAccountType)
	return api
}
func (api *SpotSubAccountUniversalTransferApi) ToAccountType(ToAccountType string) *SpotSubAccountUniversalTransferApi {
	api.req.ToAccountType = GetPointer(ToAccountType)
	return api
}
func (api *SpotSubAccountUniversalTransferApi) ClientTranId(ClientTranId string) *SpotSubAccountUniversalTransferApi {
	api.req.ClientTranId = GetPointer(ClientTranId)
	return api
}
func (api *SpotSubAccountUniversalTransferApi) Symbol(Symbol string) *SpotSubAccountUniversalTransferApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *SpotSubAccountUniversalTransferApi) Asset(Asset string) *SpotSubAccountUniversalTransferApi {
	api.req.Asset = GetPointer(Asset)
	return api
}
func (api *SpotSubAccountUniversalTransferApi) Amount(Amount decimal.Decimal) *SpotSubAccountUniversalTransferApi {
	api.req.Amount = GetPointer(Amount)
	return api
}
func (api *SpotSubAccountUniversalTransferApi) Recvwindow(Recvwindow int64) *SpotSubAccountUniversalTransferApi {
	api.req.Recvwindow = GetPointer(Recvwindow)
	return api
}
func (api *SpotSubAccountUniversalTransferApi) Timestamp(Timestamp int64) *SpotSubAccountUniversalTransferApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SpotSubAccountUniversalTransferHistoryReq struct {
	FromEmail    *string `json:"fromEmail"`    //NO
	ToEmail      *string `json:"toEmail"`      //NO
	ClientTranId *string `json:"clientTranId"` //NO
	StartTime    *int64  `json:"startTime"`    //NO
	EndTime      *int64  `json:"endTime"`      //NO
	Page         *int    `json:"page"`         //NO	默认 1
	Limit        *int    `json:"limit"`        //NO	默认 500, 最大 500
	Recvwindow   *int64  `json:"recvWindow"`
	Timestamp    *int64  `json:"timestamp"`
}
type SpotSubAccountUniversalTransferHistoryApi struct {
	SpotRestClient
	req *SpotSubAccountUniversalTransferHistoryReq
}

func (api *SpotSubAccountUniversalTransferHistoryApi) FromEmail(FromEmail string) *SpotSubAccountUniversalTransferHistoryApi {
	api.req.FromEmail = GetPointer(FromEmail)
	return api
}
func (api *SpotSubAccountUniversalTransferHistoryApi) ToEmail(ToEmail string) *SpotSubAccountUniversalTransferHistoryApi {
	api.req.ToEmail = GetPointer(ToEmail)
	return api
}
func (api *SpotSubAccountUniversalTransferHistoryApi) ClientTranId(ClientTranId string) *SpotSubAccountUniversalTransferHistoryApi {
	api.req.ClientTranId = GetPointer(ClientTranId)
	return api
}
func (api *SpotSubAccountUniversalTransferHistoryApi) StartTime(StartTime int64) *SpotSubAccountUniversalTransferHistoryApi {
	api.req.StartTime = GetPointer(StartTime)
	return api
}
func (api *SpotSubAccountUniversalTransferHistoryApi) EndTime(EndTime int64) *SpotSubAccountUniversalTransferHistoryApi {
	api.req.EndTime = GetPointer(EndTime)
	return api
}
func (api *SpotSubAccountUniversalTransferHistoryApi) Page(Page int) *SpotSubAccountUniversalTransferHistoryApi {
	api.req.Page = GetPointer(Page)
	return api
}
func (api *SpotSubAccountUniversalTransferHistoryApi) Limit(Limit int) *SpotSubAccountUniversalTransferHistoryApi {
	api.req.Limit = GetPointer(Limit)
	return api
}
func (api *SpotSubAccountUniversalTransferHistoryApi) Recvwindow(Recvwindow int64) *SpotSubAccountUniversalTransferHistoryApi {
	api.req.Recvwindow = GetPointer(Recvwindow)
	return api
}
func (api *SpotSubAccountUniversalTransferHistoryApi) Timestamp(Timestamp int64) *SpotSubAccountUniversalTransferHistoryApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SpotSubAccountAssetsReq struct {
	Email      *string `json:"email"` //YES	子账户邮箱 备注
	Recvwindow *int64  `json:"recvWindow"`
	Timestamp  *int64  `json:"timestamp"`
}
type SpotSubAccountAssetsApi struct {
	SpotRestClient
	req *SpotSubAccountAssetsReq
}

func (api *SpotSubAccountAssetsApi) Email(Email string) *SpotSubAccountAssetsApi {
	api.req.Email = GetPointer(Email)
	return api
}
func (api *SpotSubAccountAssetsApi) Recvwindow(Recvwindow int64) *SpotSubAccountAssetsApi {
	api.req.Recvwindow = GetPointer(Recvwindow)
	return api
}
func (api *SpotSubAccountAssetsApi) Timestamp(Timestamp int64) *SpotSubAccountAssetsApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SpotSubAccountFuturesAccountReq struct {
	Email       *string `json:"email"`       //YES	子账户邮箱 备注
	FuturesType *int    `json:"futuresType"` //YES	1:USDT Margined Futures, 2:COIN Margined Futures
	RecvWindow  *int64  `json:"recvWindow"`  //NO
	Timestamp   *int64  `json:"timestamp"`   //YES
}

type SpotSubAccountFuturesAccountApi struct {
	SpotRestClient
	req *SpotSubAccountFuturesAccountReq
}

func (api *SpotSubAccountFuturesAccountApi) Email(Email string) *SpotSubAccountFuturesAccountApi {
	api.req.Email = GetPointer(Email)
	return api
}
func (api *SpotSubAccountFuturesAccountApi) FuturesType(FuturesType int) *SpotSubAccountFuturesAccountApi {
	api.req.FuturesType = GetPointer(FuturesType)
	return api
}
func (api *SpotSubAccountFuturesAccountApi) RecvWindow(RecvWindow int64) *SpotSubAccountFuturesAccountApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *SpotSubAccountFuturesAccountApi) Timestamp(Timestamp int64) *SpotSubAccountFuturesAccountApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SpotSubAccountFuturesEnableReq struct {
	Email      *string `json:"email"` //YES	子账户邮箱 备注
	Recvwindow *int64  `json:"recvWindow"`
	Timestamp  *int64  `json:"timestamp"`
}

type SpotSubAccountFuturesEnableApi struct {
	SpotRestClient
	req *SpotSubAccountFuturesEnableReq
}

func (api *SpotSubAccountFuturesEnableApi) Email(Email string) *SpotSubAccountFuturesEnableApi {
	api.req.Email = GetPointer(Email)
	return api
}
func (api *SpotSubAccountFuturesEnableApi) Recvwindow(Recvwindow int64) *SpotSubAccountFuturesEnableApi {
	api.req.Recvwindow = GetPointer(Recvwindow)
	return api
}
func (api *SpotSubAccountFuturesEnableApi) Timestamp(Timestamp int64) *SpotSubAccountFuturesEnableApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SpotAccountReq struct {
	Recvwindow *int64 `json:"recvWindow"`
	Timestamp  *int64 `json:"timestamp"`
}
type SpotAccountApi struct {
	SpotRestClient
	req *SpotAccountReq
}

func (api *SpotAccountApi) Recvwindow(Recvwindow int64) *SpotAccountApi {
	api.req.Recvwindow = GetPointer(Recvwindow)
	return api
}
func (api *SpotAccountApi) Timestamp(Timestamp int64) *SpotAccountApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SpotSubAccountApiIpRestrictionReq struct {
	Email            *string `json:"email"`            //YES	Sub-account email
	SubAccountApiKey *string `json:"subAccountApiKey"` //YES
	Recvwindow       *int64  `json:"recvWindow"`
	Timestamp        *int64  `json:"timestamp"`
}

type SpotSubAccountApiIpRestrictionApi struct {
	SpotRestClient
	req *SpotSubAccountApiIpRestrictionReq
}

func (api *SpotSubAccountApiIpRestrictionApi) Email(Email string) *SpotSubAccountApiIpRestrictionApi {
	api.req.Email = GetPointer(Email)
	return api
}
func (api *SpotSubAccountApiIpRestrictionApi) SubAccountApiKey(SubAccountApiKey string) *SpotSubAccountApiIpRestrictionApi {
	api.req.SubAccountApiKey = GetPointer(SubAccountApiKey)
	return api
}
func (api *SpotSubAccountApiIpRestrictionApi) Recvwindow(Recvwindow int64) *SpotSubAccountApiIpRestrictionApi {
	api.req.Recvwindow = GetPointer(Recvwindow)
	return api
}
func (api *SpotSubAccountApiIpRestrictionApi) Timestamp(Timestamp int64) *SpotSubAccountApiIpRestrictionApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SpotMarginAllPairsReq struct {
	Recvwindow *int64 `json:"recvWindow"`
	Timestamp  *int64 `json:"timestamp"`
}

type SpotMarginAllPairsApi struct {
	SpotRestClient
	req *SpotMarginAllPairsReq
}

func (api *SpotMarginAllPairsApi) Recvwindow(Recvwindow int64) *SpotMarginAllPairsApi {
	api.req.Recvwindow = GetPointer(Recvwindow)
	return api
}
func (api *SpotMarginAllPairsApi) Timestamp(Timestamp int64) *SpotMarginAllPairsApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SpotMarginIsolatedAllPairsReq struct {
	Recvwindow *int64 `json:"recvWindow"`
	Timestamp  *int64 `json:"timestamp"`
}

type SpotMarginIsolatedAllPairsApi struct {
	SpotRestClient
	req *SpotMarginIsolatedAllPairsReq
}

func (api *SpotMarginIsolatedAllPairsApi) Recvwindow(Recvwindow int64) *SpotMarginIsolatedAllPairsApi {
	api.req.Recvwindow = GetPointer(Recvwindow)
	return api
}
func (api *SpotMarginIsolatedAllPairsApi) Timestamp(Timestamp int64) *SpotMarginIsolatedAllPairsApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SpotMarginAccountReq struct {
	Recvwindow *int64 `json:"recvWindow"`
	Timestamp  *int64 `json:"timestamp"`
}

type SpotMarginAccountApi struct {
	SpotRestClient
	req *SpotMarginAccountReq
}

func (api *SpotMarginAccountApi) Recvwindow(Recvwindow int64) *SpotMarginAccountApi {
	api.req.Recvwindow = GetPointer(Recvwindow)
	return api
}
func (api *SpotMarginAccountApi) Timestamp(Timestamp int64) *SpotMarginAccountApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SpotMarginIsolatedAccountReq struct {
	Symbols    *string `json:"symbols"` //NO	最多可以传5个symbol; 由","分隔的字符串表示. e.g. "BTCUSDT,BNBUSDT,ADAUSDT"
	Recvwindow *int64  `json:"recvWindow"`
	Timestamp  *int64  `json:"timestamp"`
}

type SpotMarginIsolatedAccountApi struct {
	SpotRestClient
	req *SpotMarginIsolatedAccountReq
}

func (api *SpotMarginIsolatedAccountApi) Symbols(Symbols string) *SpotMarginIsolatedAccountApi {
	api.req.Symbols = GetPointer(Symbols)
	return api
}

func (api *SpotMarginIsolatedAccountApi) Recvwindow(Recvwindow int64) *SpotMarginIsolatedAccountApi {
	api.req.Recvwindow = GetPointer(Recvwindow)
	return api
}

func (api *SpotMarginIsolatedAccountApi) Timestamp(Timestamp int64) *SpotMarginIsolatedAccountApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SpotMarginTransferReq struct {
	Asset      *string          `json:"asset"`      //YES 被划转的资产, 比如, BTC
	Amount     *decimal.Decimal `json:"amount"`     //YES 划转数量
	Type       *int             `json:"type"`       //YES 1: 主账户向全仓杠杆账户划转 2: 全仓杠杆账户向主账户划转
	Recvwindow *int64           `json:"recvWindow"` //NO 赋值不能大于 60000
	Timestamp  *int64           `json:"timestamp"`  //YES
}

type SpotMarginTransferApi struct {
	SpotRestClient
	req *SpotMarginTransferReq
}

func (api *SpotMarginTransferApi) Asset(Asset string) *SpotMarginTransferApi {
	api.req.Asset = GetPointer(Asset)
	return api
}
func (api *SpotMarginTransferApi) Amount(Amount decimal.Decimal) *SpotMarginTransferApi {
	api.req.Amount = GetPointer(Amount)
	return api
}
func (api *SpotMarginTransferApi) Type(Type int) *SpotMarginTransferApi {
	api.req.Type = GetPointer(Type)
	return api
}
func (api *SpotMarginTransferApi) Recvwindow(Recvwindow int64) *SpotMarginTransferApi {
	api.req.Recvwindow = GetPointer(Recvwindow)
	return api
}
func (api *SpotMarginTransferApi) Timestamp(Timestamp int64) *SpotMarginTransferApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SpotMarginIsolatedTransferReq struct {
	Asset      *string          `json:"asset"`      //YES 被划转的资产, 比如, BTC
	Symbol     *string          `json:"symbol"`     //YES 逐仓 symbol
	TransFrom  *string          `json:"transFrom"`  //YES "SPOT", "ISOLATED_MARGIN"
	TransTo    *string          `json:"transTo"`    //YES "SPOT", "ISOLATED_MARGIN"
	Amount     *decimal.Decimal `json:"amount"`     //YES 划转数量
	Recvwindow *int64           `json:"recvWindow"` //NO 赋值不能大于 60000
	Timestamp  *int64           `json:"timestamp"`  //YES
}

type SpotMarginIsolatedTransferApi struct {
	SpotRestClient
	req *SpotMarginIsolatedTransferReq
}

func (api *SpotMarginIsolatedTransferApi) Asset(Asset string) *SpotMarginIsolatedTransferApi {
	api.req.Asset = GetPointer(Asset)
	return api
}
func (api *SpotMarginIsolatedTransferApi) Symbol(Symbol string) *SpotMarginIsolatedTransferApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *SpotMarginIsolatedTransferApi) TransFrom(TransFrom string) *SpotMarginIsolatedTransferApi {
	api.req.TransFrom = GetPointer(TransFrom)
	return api
}
func (api *SpotMarginIsolatedTransferApi) TransTo(TransTo string) *SpotMarginIsolatedTransferApi {
	api.req.TransTo = GetPointer(TransTo)
	return api
}
func (api *SpotMarginIsolatedTransferApi) Amount(Amount decimal.Decimal) *SpotMarginIsolatedTransferApi {
	api.req.Amount = GetPointer(Amount)
	return api
}
func (api *SpotMarginIsolatedTransferApi) Recvwindow(Recvwindow int64) *SpotMarginIsolatedTransferApi {
	api.req.Recvwindow = GetPointer(Recvwindow)
	return api
}
func (api *SpotMarginIsolatedTransferApi) Timestamp(Timestamp int64) *SpotMarginIsolatedTransferApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SpotMarginLoanReq struct {
	Asset      *string          `json:"asset"`      //YES 被划转的资产, 比如, BTC
	IsIsolated *string          `json:"isIsolated"` //NO 是否逐仓杠杆，"TRUE", "FALSE", 默认 "FALSE"
	Symbol     *string          `json:"symbol"`     //NO 逐仓交易对，配合逐仓使用
	Amount     *decimal.Decimal `json:"amount"`     //YES 划转数量
	Recvwindow *int64           `json:"recvWindow"` //NO 赋值不能大于 60000
	Timestamp  *int64           `json:"timestamp"`  //YES
}

type SpotMarginLoanApi struct {
	SpotRestClient
	req *SpotMarginLoanReq
}

func (api *SpotMarginLoanApi) Asset(Asset string) *SpotMarginLoanApi {
	api.req.Asset = GetPointer(Asset)
	return api
}
func (api *SpotMarginLoanApi) IsIsolated(IsIsolated string) *SpotMarginLoanApi {
	api.req.IsIsolated = GetPointer(IsIsolated)
	return api
}
func (api *SpotMarginLoanApi) Symbol(Symbol string) *SpotMarginLoanApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *SpotMarginLoanApi) Amount(Amount decimal.Decimal) *SpotMarginLoanApi {
	api.req.Amount = GetPointer(Amount)
	return api
}
func (api *SpotMarginLoanApi) Recvwindow(Recvwindow int64) *SpotMarginLoanApi {
	api.req.Recvwindow = GetPointer(Recvwindow)
	return api
}
func (api *SpotMarginLoanApi) Timestamp(Timestamp int64) *SpotMarginLoanApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SpotMarginRepayReq struct {
	Asset      *string          `json:"asset"`      //YES 被划转的资产, 比如, BTC
	IsIsolated *string          `json:"isIsolated"` //NO 是否逐仓杠杆，"TRUE", "FALSE", 默认 "FALSE"
	Symbol     *string          `json:"symbol"`     //NO 逐仓交易对，配合逐仓使用
	Amount     *decimal.Decimal `json:"amount"`     //YES 划转数量
	Recvwindow *int64           `json:"recvWindow"` //NO 赋值不能大于 60000
	Timestamp  *int64           `json:"timestamp"`  //YES
}

type SpotMarginRepayApi struct {
	SpotRestClient
	req *SpotMarginRepayReq
}

func (api *SpotMarginRepayApi) Asset(Asset string) *SpotMarginRepayApi {
	api.req.Asset = GetPointer(Asset)
	return api
}
func (api *SpotMarginRepayApi) IsIsolated(IsIsolated string) *SpotMarginRepayApi {
	api.req.IsIsolated = GetPointer(IsIsolated)
	return api
}
func (api *SpotMarginRepayApi) Symbol(Symbol string) *SpotMarginRepayApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *SpotMarginRepayApi) Amount(Amount decimal.Decimal) *SpotMarginRepayApi {
	api.req.Amount = GetPointer(Amount)
	return api
}
func (api *SpotMarginRepayApi) Recvwindow(Recvwindow int64) *SpotMarginRepayApi {
	api.req.Recvwindow = GetPointer(Recvwindow)
	return api
}
func (api *SpotMarginRepayApi) Timestamp(Timestamp int64) *SpotMarginRepayApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SpotMarginMaxBorrowableReq struct {
	Asset          *string `json:"asset"`          //YES 币种, 比如, BTC
	IsolatedSymbol *string `json:"isolatedSymbol"` //NO 逐仓交易对，适用于逐仓查询
	Recvwindow     *int64  `json:"recvWindow"`     //NO 赋值不能大于 60000
	Timestamp      *int64  `json:"timestamp"`      //YES
}

type SpotMarginMaxBorrowableApi struct {
	SpotRestClient
	req *SpotMarginMaxBorrowableReq
}

func (api *SpotMarginMaxBorrowableApi) Asset(Asset string) *SpotMarginMaxBorrowableApi {
	api.req.Asset = GetPointer(Asset)
	return api
}
func (api *SpotMarginMaxBorrowableApi) IsolatedSymbol(IsolatedSymbol string) *SpotMarginMaxBorrowableApi {
	api.req.IsolatedSymbol = GetPointer(IsolatedSymbol)
	return api
}
func (api *SpotMarginMaxBorrowableApi) Recvwindow(Recvwindow int64) *SpotMarginMaxBorrowableApi {
	api.req.Recvwindow = GetPointer(Recvwindow)
	return api
}
func (api *SpotMarginMaxBorrowableApi) Timestamp(Timestamp int64) *SpotMarginMaxBorrowableApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SpotMarginMaxTransferableReq struct {
	Asset          *string `json:"asset"`          //YES 币种, 比如, BTC
	IsolatedSymbol *string `json:"isolatedSymbol"` //NO 逐仓交易对，适用于逐仓查询
	Recvwindow     *int64  `json:"recvWindow"`     //NO 赋值不能大于 60000
	Timestamp      *int64  `json:"timestamp"`      //YES
}

type SpotMarginMaxTransferableApi struct {
	SpotRestClient
	req *SpotMarginMaxTransferableReq
}

func (api *SpotMarginMaxTransferableApi) Asset(Asset string) *SpotMarginMaxTransferableApi {
	api.req.Asset = GetPointer(Asset)
	return api
}
func (api *SpotMarginMaxTransferableApi) IsolatedSymbol(IsolatedSymbol string) *SpotMarginMaxTransferableApi {
	api.req.IsolatedSymbol = GetPointer(IsolatedSymbol)
	return api
}
func (api *SpotMarginMaxTransferableApi) Recvwindow(Recvwindow int64) *SpotMarginMaxTransferableApi {
	api.req.Recvwindow = GetPointer(Recvwindow)
	return api
}
func (api *SpotMarginMaxTransferableApi) Timestamp(Timestamp int64) *SpotMarginMaxTransferableApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SpotMarginInterestHistoryReq struct {
	Asset          *string `json:"asset"`          //NO 币种, 比如, BTC
	IsolatedSymbol *string `json:"isolatedSymbol"` //NO 逐仓交易对，适用于逐仓查询
	StartTime      *int64  `json:"startTime"`      //NO
	EndTime        *int64  `json:"endTime"`        //NO
	Current        *int64  `json:"current"`        //NO	当前查询页。 开始值 1. 默认:1
	Size           *int64  `json:"size"`           //NO	默认:10 最大:100
	Archived       *string `json:"archived"`       //NO	默认: false. 查询6个月以前的数据，需要设为 true
	Recvwindow     *int64  `json:"recvWindow"`     //NO 赋值不能大于 60000
	Timestamp      *int64  `json:"timestamp"`      //YES
}

type SpotMarginInterestHistoryApi struct {
	SpotRestClient
	req *SpotMarginInterestHistoryReq
}

func (api *SpotMarginInterestHistoryApi) Asset(Asset string) *SpotMarginInterestHistoryApi {
	api.req.Asset = GetPointer(Asset)
	return api
}
func (api *SpotMarginInterestHistoryApi) IsolatedSymbol(IsolatedSymbol string) *SpotMarginInterestHistoryApi {
	api.req.IsolatedSymbol = GetPointer(IsolatedSymbol)
	return api
}
func (api *SpotMarginInterestHistoryApi) StartTime(StartTime int64) *SpotMarginInterestHistoryApi {
	api.req.StartTime = GetPointer(StartTime)
	return api
}
func (api *SpotMarginInterestHistoryApi) EndTime(EndTime int64) *SpotMarginInterestHistoryApi {
	api.req.EndTime = GetPointer(EndTime)
	return api
}
func (api *SpotMarginInterestHistoryApi) Current(Current int64) *SpotMarginInterestHistoryApi {
	api.req.Current = GetPointer(Current)
	return api
}
func (api *SpotMarginInterestHistoryApi) Size(Size int64) *SpotMarginInterestHistoryApi {
	api.req.Size = GetPointer(Size)
	return api
}
func (api *SpotMarginInterestHistoryApi) Archived(Archived string) *SpotMarginInterestHistoryApi {
	api.req.Archived = GetPointer(Archived)
	return api
}
func (api *SpotMarginInterestHistoryApi) Recvwindow(Recvwindow int64) *SpotMarginInterestHistoryApi {
	api.req.Recvwindow = GetPointer(Recvwindow)
	return api
}
func (api *SpotMarginInterestHistoryApi) Timestamp(Timestamp int64) *SpotMarginInterestHistoryApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SpotMarginOrderGetReq struct {
	Symbol            *string `json:"symbol"`            //NO
	IsIsolated        *string `json:"isIsolated"`        //NO 	是否逐仓杠杆，"TRUE", "FALSE", 默认 "FALSE"
	OrderId           *int64  `json:"orderId"`           //NO
	OrigClientOrderId *string `json:"origClientOrderId"` //NO
	Recvwindow        *int64  `json:"recvWindow"`        //NO 赋值不能大于 60000
	Timestamp         *int64  `json:"timestamp"`         //YES
}

type SpotMarginOrderGetApi struct {
	SpotRestClient
	req *SpotMarginOrderGetReq
}

func (api *SpotMarginOrderGetApi) Symbol(Symbol string) *SpotMarginOrderGetApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *SpotMarginOrderGetApi) IsIsolated(IsIsolated string) *SpotMarginOrderGetApi {
	api.req.IsIsolated = GetPointer(IsIsolated)
	return api
}
func (api *SpotMarginOrderGetApi) OrderId(OrderId int64) *SpotMarginOrderGetApi {
	api.req.OrderId = GetPointer(OrderId)
	return api
}
func (api *SpotMarginOrderGetApi) OrigClientOrderId(OrigClientOrderId string) *SpotMarginOrderGetApi {
	api.req.OrigClientOrderId = GetPointer(OrigClientOrderId)
	return api
}
func (api *SpotMarginOrderGetApi) Recvwindow(Recvwindow int64) *SpotMarginOrderGetApi {
	api.req.Recvwindow = GetPointer(Recvwindow)
	return api
}
func (api *SpotMarginOrderGetApi) Timestamp(Timestamp int64) *SpotMarginOrderGetApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SpotMarginAllOrdersReq struct {
	Symbol     *string `json:"symbol"`     //YES
	IsIsolated *string `json:"isIsolated"` //NO 	是否逐仓杠杆，"TRUE", "FALSE", 默认 "FALSE"
	OrderId    *int64  `json:"orderId"`    //NO
	StartTime  *int64  `json:"startTime"`  //NO
	EndTime    *int64  `json:"endTime"`    //NO
	Limit      *int64  `json:"limit"`      //NO	默认 500;最大500.
	Recvwindow *int64  `json:"recvWindow"` //NO 赋值不能大于 60000
	Timestamp  *int64  `json:"timestamp"`  //YES
}

type SpotMarginAllOrdersApi struct {
	SpotRestClient
	req *SpotMarginAllOrdersReq
}

func (api *SpotMarginAllOrdersApi) Symbol(Symbol string) *SpotMarginAllOrdersApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *SpotMarginAllOrdersApi) IsIsolated(IsIsolated string) *SpotMarginAllOrdersApi {
	api.req.IsIsolated = GetPointer(IsIsolated)
	return api
}
func (api *SpotMarginAllOrdersApi) OrderId(OrderId int64) *SpotMarginAllOrdersApi {
	api.req.OrderId = GetPointer(OrderId)
	return api
}
func (api *SpotMarginAllOrdersApi) StartTime(StartTime int64) *SpotMarginAllOrdersApi {
	api.req.StartTime = GetPointer(StartTime)
	return api
}
func (api *SpotMarginAllOrdersApi) EndTime(EndTime int64) *SpotMarginAllOrdersApi {
	api.req.EndTime = GetPointer(EndTime)
	return api
}
func (api *SpotMarginAllOrdersApi) Limit(Limit int64) *SpotMarginAllOrdersApi {
	api.req.Limit = GetPointer(Limit)
	return api
}
func (api *SpotMarginAllOrdersApi) Recvwindow(Recvwindow int64) *SpotMarginAllOrdersApi {
	api.req.Recvwindow = GetPointer(Recvwindow)
	return api
}
func (api *SpotMarginAllOrdersApi) Timestamp(Timestamp int64) *SpotMarginAllOrdersApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SpotMarginOpenOrdersReq struct {
	Symbol     *string `json:"symbol"`     //NO
	IsIsolated *string `json:"isIsolated"` //NO 	是否逐仓杠杆，"TRUE", "FALSE", 默认 "FALSE"
	Recvwindow *int64  `json:"recvWindow"` //NO 赋值不能大于 60000
	Timestamp  *int64  `json:"timestamp"`  //YES
}

type SpotMarginOpenOrdersApi struct {
	SpotRestClient
	req *SpotMarginOpenOrdersReq
}

func (api *SpotMarginOpenOrdersApi) Symbol(Symbol string) *SpotMarginOpenOrdersApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *SpotMarginOpenOrdersApi) IsIsolated(IsIsolated string) *SpotMarginOpenOrdersApi {
	api.req.IsIsolated = GetPointer(IsIsolated)
	return api
}
func (api *SpotMarginOpenOrdersApi) Recvwindow(Recvwindow int64) *SpotMarginOpenOrdersApi {
	api.req.Recvwindow = GetPointer(Recvwindow)
	return api
}
func (api *SpotMarginOpenOrdersApi) Timestamp(Timestamp int64) *SpotMarginOpenOrdersApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SpotOrderPostReq struct {
	Symbol                  *string          `json:"symbol"`                  //YES
	Side                    *string          `json:"side"`                    //YES
	Type                    *string          `json:"type"`                    //YES
	TimeInForce             *string          `json:"timeInForce"`             //NO 	详见枚举定义：有效方式
	Quantity                *decimal.Decimal `json:"quantity"`                //NO
	QuoteOrderQty           *decimal.Decimal `json:"quoteOrderQty"`           //NO
	Price                   *decimal.Decimal `json:"price"`                   //NO
	NewClientOrderId        *string          `json:"newClientOrderId"`        //NO 	客户自定义的唯一订单ID。 如果未发送，则自动生成。
	StopPrice               *decimal.Decimal `json:"stopPrice"`               //NO 	仅 STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT 和 TAKE_PROFIT_LIMIT 需要此参数。
	TrailingDelta           *int64           `json:"trailingDelta"`           //NO 	用于 STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT 和 TAKE_PROFIT_LIMIT 类型的订单。更多追踪止盈止损订单细节, 请参考 追踪止盈止损(Trailing Stop)订单常见问题。
	IcebergQty              *decimal.Decimal `json:"icebergQty"`              //NO 	仅使用 LIMIT, STOP_LOSS_LIMIT, 和 TAKE_PROFIT_LIMIT 创建新的 iceberg 订单时需要此参数。
	NewOrderRespType        *string          `json:"newOrderRespType"`        //NO 	设置响应JSON。ACK，RESULT 或 FULL；MARKET 和 LIMIT 订单类型默认为 FULL，所有其他订单默认为 ACK。
	SelfTradePreventionMode *string          `json:"selfTradePreventionMode"` //NO 	允许的 ENUM 取决于交易对的配置。支持的值有 EXPIRE_TAKER，EXPIRE_MAKER，EXPIRE_BOTH，NONE。
	StrategyId              *int             `json:"strategyId"`              //NO
	StrategyType            *int             `json:"strategyType"`            //NO 	不能低于 1000000
	RecvWindow              *int64           `json:"recvWindow"`              //NO 	赋值不能大于 60000
	Timestamp               *int64           `json:"timestamp"`               //YES
}

type SpotOrderPostApi struct {
	SpotRestClient
	req *SpotOrderPostReq
}

func (api *SpotOrderPostApi) Symbol(Symbol string) *SpotOrderPostApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *SpotOrderPostApi) Side(Side string) *SpotOrderPostApi {
	api.req.Side = GetPointer(Side)
	return api
}
func (api *SpotOrderPostApi) Type(Type string) *SpotOrderPostApi {
	api.req.Type = GetPointer(Type)
	return api
}
func (api *SpotOrderPostApi) TimeInForce(TimeInForce string) *SpotOrderPostApi {
	api.req.TimeInForce = GetPointer(TimeInForce)
	return api
}
func (api *SpotOrderPostApi) Quantity(Quantity decimal.Decimal) *SpotOrderPostApi {
	api.req.Quantity = GetPointer(Quantity)
	return api
}
func (api *SpotOrderPostApi) QuoteOrderQty(QuoteOrderQty decimal.Decimal) *SpotOrderPostApi {
	api.req.QuoteOrderQty = GetPointer(QuoteOrderQty)
	return api
}
func (api *SpotOrderPostApi) Price(Price decimal.Decimal) *SpotOrderPostApi {
	api.req.Price = GetPointer(Price)
	return api
}
func (api *SpotOrderPostApi) NewClientOrderId(NewClientOrderId string) *SpotOrderPostApi {
	api.req.NewClientOrderId = GetPointer(NewClientOrderId)
	return api
}
func (api *SpotOrderPostApi) StopPrice(StopPrice decimal.Decimal) *SpotOrderPostApi {
	api.req.StopPrice = GetPointer(StopPrice)
	return api
}
func (api *SpotOrderPostApi) TrailingDelta(TrailingDelta int64) *SpotOrderPostApi {
	api.req.TrailingDelta = GetPointer(TrailingDelta)
	return api
}
func (api *SpotOrderPostApi) IcebergQty(IcebergQty decimal.Decimal) *SpotOrderPostApi {
	api.req.IcebergQty = GetPointer(IcebergQty)
	return api
}
func (api *SpotOrderPostApi) NewOrderRespType(NewOrderRespType string) *SpotOrderPostApi {
	api.req.NewOrderRespType = GetPointer(NewOrderRespType)
	return api
}
func (api *SpotOrderPostApi) SelfTradePreventionMode(SelfTradePreventionMode string) *SpotOrderPostApi {
	api.req.SelfTradePreventionMode = GetPointer(SelfTradePreventionMode)
	return api
}
func (api *SpotOrderPostApi) StrategyId(StrategyId int) *SpotOrderPostApi {
	api.req.StrategyId = GetPointer(StrategyId)
	return api
}
func (api *SpotOrderPostApi) StrategyType(StrategyType int) *SpotOrderPostApi {
	api.req.StrategyType = GetPointer(StrategyType)
	return api
}
func (api *SpotOrderPostApi) RecvWindow(RecvWindow int64) *SpotOrderPostApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *SpotOrderPostApi) Timestamp(Timestamp int64) *SpotOrderPostApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

// symbol	STRING	YES
// isIsolated	STRING	NO	是否逐仓杠杆，"TRUE", "FALSE", 默认 "FALSE"
// side	ENUM	YES	BUY

// SELL
// type	ENUM	YES	详见枚举定义：订单类型
// quantity	DECIMAL	NO
// quoteOrderQty	DECIMAL	NO
// price	DECIMAL	NO
// stopPrice	DECIMAL	NO	与STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT, 和 TAKE_PROFIT_LIMIT 订单一起使用.
// newClientOrderId	STRING	NO	客户自定义的唯一订单ID。若未发送自动生成。
// icebergQty	DECIMAL	NO	与 LIMIT, STOP_LOSS_LIMIT, 和 TAKE_PROFIT_LIMIT 一起使用创建 iceberg 订单.
// newOrderRespType	ENUM	NO	设置响应: JSON. ACK, RESULT, 或 FULL; MARKET 和 LIMIT 订单类型默认为 FULL, 所有其他订单默认为 ACK.
// sideEffectType	ENUM	NO	NO_SIDE_EFFECT, MARGIN_BUY, AUTO_REPAY;默认为 NO_SIDE_EFFECT.
// timeInForce	ENUM	NO	GTC,IOC,FOK
// selfTradePreventionMode	ENUM	NO	允许的 ENUM 取决于交易对的配置。支持的值有 EXPIRE_TAKER，EXPIRE_MAKER，EXPIRE_BOTH，NONE
// autoRepayAtCancel	BOOLEAN	NO	只有在自动借款单或者自动借还单生效，true表示的是撤单后需要把订单产生的借款归还，默认为true
// recvWindow	LONG	NO	赋值不能大于 60000
// timestamp	LONG	YES
type SpotMarginOrderPostReq struct {
	Symbol                  *string          `json:"symbol"`                  //YES
	IsIsolated              *string          `json:"isIsolated"`              //NO	是否逐仓杠杆，"TRUE", "FALSE", 默认 "FALSE"
	Side                    *string          `json:"side"`                    //YES	BUY SELL
	Type                    *string          `json:"type"`                    //YES	详见枚举定义：订单类型
	Quantity                *decimal.Decimal `json:"quantity"`                //NO
	QuoteOrderQty           *decimal.Decimal `json:"quoteOrderQty"`           //NO
	Price                   *decimal.Decimal `json:"price"`                   //NO
	StopPrice               *decimal.Decimal `json:"stopPrice"`               //NO	与STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT, 和 TAKE_PROFIT_LIMIT 订单一起使用.
	NewClientOrderId        *string          `json:"newClientOrderId"`        //NO	客户自定义的唯一订单ID。若未发送自动生成。
	IcebergQty              *decimal.Decimal `json:"icebergQty"`              //NO	与 LIMIT, STOP_LOSS_LIMIT, 和 TAKE_PROFIT_LIMIT 一起使用创建 iceberg 订单.
	NewOrderRespType        *string          `json:"newOrderRespType"`        //NO	设置响应: JSON. ACK, RESULT, 或 FULL; MARKET 和 LIMIT 订单类型默认为 FULL, 所有其他订单默认为 ACK.
	SideEffectType          *string          `json:"sideEffectType"`          //NO	NO_SIDE_EFFECT, MARGIN_BUY, AUTO_REPAY;默认为 NO_SIDE_EFFECT.
	TimeInForce             *string          `json:"timeInForce"`             //NO	GTC,IOC,FOK
	SelfTradePreventionMode *string          `json:"selfTradePreventionMode"` //NO	允许的 ENUM 取决于交易对的配置。支持的值有 EXPIRE_TAKER，EXPIRE_MAKER，EXPIRE_BOTH，NONE
	AutoRepayAtCancel       *bool            `json:"autoRepayAtCancel"`       //NO	只有在自动借款单或者自动借还单生效，true表示的是撤单后需要把订单产生的借款归还，默认为true
	RecvWindow              *int64           `json:"recvWindow"`              //NO	赋值不能大于 60000
	Timestamp               *int64           `json:"timestamp"`               //YES
}

type SpotMarginOrderPostApi struct {
	SpotRestClient
	req *SpotMarginOrderPostReq
}

func (api *SpotMarginOrderPostApi) Symbol(Symbol string) *SpotMarginOrderPostApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *SpotMarginOrderPostApi) IsIsolated(IsIsolated string) *SpotMarginOrderPostApi {
	api.req.IsIsolated = GetPointer(IsIsolated)
	return api
}
func (api *SpotMarginOrderPostApi) Side(Side string) *SpotMarginOrderPostApi {
	api.req.Side = GetPointer(Side)
	return api
}
func (api *SpotMarginOrderPostApi) Type(Type string) *SpotMarginOrderPostApi {
	api.req.Type = GetPointer(Type)
	return api
}
func (api *SpotMarginOrderPostApi) Quantity(Quantity decimal.Decimal) *SpotMarginOrderPostApi {
	api.req.Quantity = GetPointer(Quantity)
	return api
}
func (api *SpotMarginOrderPostApi) QuoteOrderQty(QuoteOrderQty decimal.Decimal) *SpotMarginOrderPostApi {
	api.req.QuoteOrderQty = GetPointer(QuoteOrderQty)
	return api
}
func (api *SpotMarginOrderPostApi) Price(Price decimal.Decimal) *SpotMarginOrderPostApi {
	api.req.Price = GetPointer(Price)
	return api
}
func (api *SpotMarginOrderPostApi) StopPrice(StopPrice decimal.Decimal) *SpotMarginOrderPostApi {
	api.req.StopPrice = GetPointer(StopPrice)
	return api
}
func (api *SpotMarginOrderPostApi) NewClientOrderId(NewClientOrderId string) *SpotMarginOrderPostApi {
	api.req.NewClientOrderId = GetPointer(NewClientOrderId)
	return api
}
func (api *SpotMarginOrderPostApi) IcebergQty(IcebergQty decimal.Decimal) *SpotMarginOrderPostApi {
	api.req.IcebergQty = GetPointer(IcebergQty)
	return api
}
func (api *SpotMarginOrderPostApi) NewOrderRespType(NewOrderRespType string) *SpotMarginOrderPostApi {
	api.req.NewOrderRespType = GetPointer(NewOrderRespType)
	return api
}
func (api *SpotMarginOrderPostApi) SideEffectType(SideEffectType string) *SpotMarginOrderPostApi {
	api.req.SideEffectType = GetPointer(SideEffectType)
	return api
}
func (api *SpotMarginOrderPostApi) TimeInForce(TimeInForce string) *SpotMarginOrderPostApi {
	api.req.TimeInForce = GetPointer(TimeInForce)
	return api
}
func (api *SpotMarginOrderPostApi) SelfTradePreventionMode(SelfTradePreventionMode string) *SpotMarginOrderPostApi {
	api.req.SelfTradePreventionMode = GetPointer(SelfTradePreventionMode)
	return api
}
func (api *SpotMarginOrderPostApi) AutoRepayAtCancel(AutoRepayAtCancel bool) *SpotMarginOrderPostApi {
	api.req.AutoRepayAtCancel = GetPointer(AutoRepayAtCancel)
	return api
}
func (api *SpotMarginOrderPostApi) RecvWindow(RecvWindow int64) *SpotMarginOrderPostApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *SpotMarginOrderPostApi) Timestamp(Timestamp int64) *SpotMarginOrderPostApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SpotPingReq struct {
}
type SpotPingApi struct {
	SpotRestClient
	req *SpotPingReq
}

type SpotServerTimeReq struct {
}
type SpotServerTimeApi struct {
	SpotRestClient
	req *SpotServerTimeReq
}

type SpotExchangeInfoReq struct {
}
type SpotExchangeInfoApi struct {
	SpotRestClient
	req *SpotExchangeInfoReq
}

type SpotAccountApiTradingStatusReq struct {
	RecvWindow *int64 `json:"recvWindow"` //NO 	赋值不能大于 60000
	Timestamp  *int64 `json:"timestamp"`  //YES
}

type SpotAccountApiTradingStatusApi struct {
	SpotRestClient
	req *SpotAccountApiTradingStatusReq
}

func (api *SpotAccountApiTradingStatusApi) RecvWindow(RecvWindow int64) *SpotAccountApiTradingStatusApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *SpotAccountApiTradingStatusApi) Timestamp(Timestamp int64) *SpotAccountApiTradingStatusApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

// symbol	STRING	NO
// recvWindow	LONG	NO	赋值不得大于 60000
// timestamp	LONG	YES
type SpotOpenOrdersReq struct {
	Symbol     *string `json:"symbol"`     //NO
	RecvWindow *int64  `json:"recvWindow"` //NO 	赋值不能大于 60000
	Timestamp  *int64  `json:"timestamp"`  //YES
}

type SpotOpenOrdersApi struct {
	SpotRestClient
	req *SpotOpenOrdersReq
}

func (api *SpotOpenOrdersApi) Symbol(Symbol string) *SpotOpenOrdersApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *SpotOpenOrdersApi) RecvWindow(RecvWindow int64) *SpotOpenOrdersApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *SpotOpenOrdersApi) Timestamp(Timestamp int64) *SpotOpenOrdersApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

// symbol	STRING	YES
// orderId	LONG	NO
// startTime	LONG	NO
// endTime	LONG	NO
// limit	INT	NO	默认 500; 最大 1000.
// recvWindow	LONG	NO	赋值不得大于 60000
// timestamp	LONG	YES
type SpotAllOrdersReq struct {
	Symbol     *string `json:"symbol"`     //YES
	OrderId    *int64  `json:"orderId"`    //NO
	StartTime  *int64  `json:"startTime"`  //NO
	EndTime    *int64  `json:"endTime"`    //NO
	Limit      *int    `json:"limit"`      //NO 默认 500; 最大 1000.
	RecvWindow *int64  `json:"recvWindow"` //NO 	赋值不能大于 60000
	Timestamp  *int64  `json:"timestamp"`  //YES
}

type SpotAllOrdersApi struct {
	SpotRestClient
	req *SpotAllOrdersReq
}

func (api *SpotAllOrdersApi) Symbol(Symbol string) *SpotAllOrdersApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *SpotAllOrdersApi) OrderId(OrderId int64) *SpotAllOrdersApi {
	api.req.OrderId = GetPointer(OrderId)
	return api
}
func (api *SpotAllOrdersApi) StartTime(StartTime int64) *SpotAllOrdersApi {
	api.req.StartTime = GetPointer(StartTime)
	return api
}
func (api *SpotAllOrdersApi) EndTime(EndTime int64) *SpotAllOrdersApi {
	api.req.EndTime = GetPointer(EndTime)
	return api
}
func (api *SpotAllOrdersApi) Limit(Limit int) *SpotAllOrdersApi {
	api.req.Limit = GetPointer(Limit)
	return api
}
func (api *SpotAllOrdersApi) RecvWindow(RecvWindow int64) *SpotAllOrdersApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *SpotAllOrdersApi) Timestamp(Timestamp int64) *SpotAllOrdersApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

// symbol	STRING	YES
// orderId	LONG	NO
// origClientOrderId	STRING	NO
// recvWindow	LONG	NO	赋值不得大于 60000
// timestamp	LONG	YES

type SpotOrderGetReq struct {
	Symbol            *string `json:"symbol"`            //YES
	OrderId           *int64  `json:"orderId"`           //NO
	OrigClientOrderId *string `json:"origClientOrderId"` //NO
	RecvWindow        *int64  `json:"recvWindow"`        //NO 	赋值不能大于 60000
	Timestamp         *int64  `json:"timestamp"`         //YES
}

type SpotOrderGetApi struct {
	SpotRestClient
	req *SpotOrderGetReq
}

func (api *SpotOrderGetApi) Symbol(Symbol string) *SpotOrderGetApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *SpotOrderGetApi) OrderId(OrderId int64) *SpotOrderGetApi {
	api.req.OrderId = GetPointer(OrderId)
	return api
}
func (api *SpotOrderGetApi) OrigClientOrderId(OrigClientOrderId string) *SpotOrderGetApi {
	api.req.OrigClientOrderId = GetPointer(OrigClientOrderId)
	return api
}
func (api *SpotOrderGetApi) RecvWindow(RecvWindow int64) *SpotOrderGetApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *SpotOrderGetApi) Timestamp(Timestamp int64) *SpotOrderGetApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SpotTickerPriceReq struct {
	Symbol *string `json:"symbol"` //YES
}

type SpotTickerPriceApi struct {
	SpotRestClient
	req *SpotTickerPriceReq
}

func (api *SpotTickerPriceApi) Symbol(Symbol string) *SpotTickerPriceApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}

// type	ENUM	YES
// asset	STRING	YES
// amount	DECIMAL	YES
// fromSymbol	STRING	NO
// toSymbol	STRING	NO
// recvWindow	LONG	NO
// timestamp	LONG	YES
type SpotAssetTransferPostReq struct {
	Type       *AssertTransferType `json:"type"`       //YES
	Asset      *string             `json:"asset"`      //YES
	Amount     *decimal.Decimal    `json:"amount"`     //YES
	FromSymbol *string             `json:"fromSymbol"` //NO
	ToSymbol   *string             `json:"toSymbol"`   //NO
	RecvWindow *int64              `json:"recvWindow"` //NO
	Timestamp  *int64              `json:"timestamp"`  //YES
}

type SpotAssetTransferPostApi struct {
	SpotRestClient
	req *SpotAssetTransferPostReq
}

func (api *SpotAssetTransferPostApi) Type(Type AssertTransferType) *SpotAssetTransferPostApi {
	api.req.Type = GetPointer(Type)
	return api
}
func (api *SpotAssetTransferPostApi) Asset(Asset string) *SpotAssetTransferPostApi {
	api.req.Asset = GetPointer(Asset)
	return api
}
func (api *SpotAssetTransferPostApi) Amount(Amount decimal.Decimal) *SpotAssetTransferPostApi {
	api.req.Amount = GetPointer(Amount)
	return api
}
func (api *SpotAssetTransferPostApi) FromSymbol(FromSymbol string) *SpotAssetTransferPostApi {
	api.req.FromSymbol = GetPointer(FromSymbol)
	return api
}
func (api *SpotAssetTransferPostApi) ToSymbol(ToSymbol string) *SpotAssetTransferPostApi {
	api.req.ToSymbol = GetPointer(ToSymbol)
	return api
}
func (api *SpotAssetTransferPostApi) RecvWindow(RecvWindow int64) *SpotAssetTransferPostApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *SpotAssetTransferPostApi) Timestamp(Timestamp int64) *SpotAssetTransferPostApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

// type	ENUM	YES
// startTime	LONG	NO
// endTime	LONG	NO
// current	INT	NO	默认 1
// size	INT	NO	默认 10, 最大 100
// fromSymbol	STRING	NO
// toSymbol	STRING	NO
// recvWindow	LONG	NO
// timestamp	LONG	YES

type SpotAssetTransferGetReq struct {
	Type       *AssertTransferType `json:"type"`       //YES
	StartTime  *int64              `json:"startTime"`  //NO
	EndTime    *int64              `json:"endTime"`    //NO
	Current    *int64              `json:"current"`    //NO	默认 1
	Size       *int64              `json:"size"`       //NO	默认 10, 最大 100
	FromSymbol *string             `json:"fromSymbol"` //NO
	ToSymbol   *string             `json:"toSymbol"`   //NO
	RecvWindow *int64              `json:"recvWindow"` //NO
	Timestamp  *int64              `json:"timestamp"`  //YES
}

type SpotAssetTransferGetApi struct {
	SpotRestClient
	req *SpotAssetTransferGetReq
}

func (api *SpotAssetTransferGetApi) Type(Type AssertTransferType) *SpotAssetTransferGetApi {
	api.req.Type = GetPointer(Type)
	return api
}
func (api *SpotAssetTransferGetApi) StartTime(StartTime int64) *SpotAssetTransferGetApi {
	api.req.StartTime = GetPointer(StartTime)
	return api
}
func (api *SpotAssetTransferGetApi) EndTime(EndTime int64) *SpotAssetTransferGetApi {
	api.req.EndTime = GetPointer(EndTime)
	return api
}
func (api *SpotAssetTransferGetApi) Current(Current int64) *SpotAssetTransferGetApi {
	api.req.Current = GetPointer(Current)
	return api
}
func (api *SpotAssetTransferGetApi) Size(Size int64) *SpotAssetTransferGetApi {
	api.req.Size = GetPointer(Size)
	return api
}
func (api *SpotAssetTransferGetApi) FromSymbol(FromSymbol string) *SpotAssetTransferGetApi {
	api.req.FromSymbol = GetPointer(FromSymbol)
	return api
}
func (api *SpotAssetTransferGetApi) ToSymbol(ToSymbol string) *SpotAssetTransferGetApi {
	api.req.ToSymbol = GetPointer(ToSymbol)
	return api
}
func (api *SpotAssetTransferGetApi) RecvWindow(RecvWindow int64) *SpotAssetTransferGetApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *SpotAssetTransferGetApi) Timestamp(Timestamp int64) *SpotAssetTransferGetApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

// asset	STRING	NO	如不提供，返回所有asset 划转记录
// type	INT	NO	1: transfer in, 2: transfer out; 如不提供，返回transfer out方向划转记录
// startTime	LONG	NO
// endTime	LONG	NO
// limit	INT	NO	默认值: 500
// returnFailHistory	BOOLEAN	NO	默认False，返回PROCESS和SUCCESS状态的数据；如果传True返回PROCESS、SUCCESS、FAILURE状态的数据
// recvWindow	LONG	NO
// timestamp	LONG	YES
type SpotSubAccountTransferSubUserHistoryReq struct {
	Asset             *string `json:"asset"`             //NO	如不提供，返回所有asset 划转记录
	Type              *int64  `json:"type"`              //NO	1: transfer in, 2: transfer out; 如不提供，返回transfer out方向划转记录
	StartTime         *int64  `json:"startTime"`         //NO
	EndTime           *int64  `json:"endTime"`           //NO
	Limit             *int64  `json:"limit"`             //NO	默认值: 500
	ReturnFailHistory *bool   `json:"returnFailHistory"` //NO	默认False，返回PROCESS和SUCCESS状态的数据；如果传True返回PROCESS、SUCCESS、FAILURE状态的数据
	RecvWindow        *int64  `json:"recvWindow"`        //NO
	Timestamp         *int64  `json:"timestamp"`         //YES
}

type SpotSubAccountTransferSubUserHistoryApi struct {
	SpotRestClient
	req *SpotSubAccountTransferSubUserHistoryReq
}

func (api *SpotSubAccountTransferSubUserHistoryApi) Asset(Asset string) *SpotSubAccountTransferSubUserHistoryApi {
	api.req.Asset = GetPointer(Asset)
	return api
}
func (api *SpotSubAccountTransferSubUserHistoryApi) Type(Type int64) *SpotSubAccountTransferSubUserHistoryApi {
	api.req.Type = GetPointer(Type)
	return api
}
func (api *SpotSubAccountTransferSubUserHistoryApi) StartTime(StartTime int64) *SpotSubAccountTransferSubUserHistoryApi {
	api.req.StartTime = GetPointer(StartTime)
	return api
}
func (api *SpotSubAccountTransferSubUserHistoryApi) EndTime(EndTime int64) *SpotSubAccountTransferSubUserHistoryApi {
	api.req.EndTime = GetPointer(EndTime)
	return api
}
func (api *SpotSubAccountTransferSubUserHistoryApi) Limit(Limit int64) *SpotSubAccountTransferSubUserHistoryApi {
	api.req.Limit = GetPointer(Limit)
	return api
}
func (api *SpotSubAccountTransferSubUserHistoryApi) ReturnFailHistory(ReturnFailHistory bool) *SpotSubAccountTransferSubUserHistoryApi {
	api.req.ReturnFailHistory = GetPointer(ReturnFailHistory)
	return api
}
func (api *SpotSubAccountTransferSubUserHistoryApi) RecvWindow(RecvWindow int64) *SpotSubAccountTransferSubUserHistoryApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *SpotSubAccountTransferSubUserHistoryApi) Timestamp(Timestamp int64) *SpotSubAccountTransferSubUserHistoryApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

// symbol	STRING	YES
// interval	ENUM	YES	详见枚举定义：K线间隔
// startTime	LONG	NO
// endTime	LONG	NO
// limit	INT	NO	默认 500; 最大 1000.
type SpotKlinesReq struct {
	Symbol    *string `json:"symbol"`    //YES
	Interval  *string `json:"interval"`  //YES	详见枚举定义：K线间隔
	StartTime *int64  `json:"startTime"` //NO
	EndTime   *int64  `json:"endTime"`   //NO
	Limit     *int    `json:"limit"`     //NO	默认 500; 最大 1000.
}

type SpotKlinesApi struct {
	SpotRestClient
	req *SpotKlinesReq
}

func (api *SpotKlinesApi) Symbol(Symbol string) *SpotKlinesApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *SpotKlinesApi) Interval(Interval string) *SpotKlinesApi {
	api.req.Interval = GetPointer(Interval)
	return api
}
func (api *SpotKlinesApi) StartTime(StartTime int64) *SpotKlinesApi {
	api.req.StartTime = GetPointer(StartTime)
	return api
}
func (api *SpotKlinesApi) EndTime(EndTime int64) *SpotKlinesApi {
	api.req.EndTime = GetPointer(EndTime)
	return api
}
func (api *SpotKlinesApi) Limit(Limit int) *SpotKlinesApi {
	api.req.Limit = GetPointer(Limit)
	return api
}

// symbol	STRING	YES
// limit	INT	NO	默认 100; 最大 5000. 可选值:[5, 10, 20, 50, 100, 500, 1000, 5000]
// 如果 limit > 5000, 最多返回5000条数据.
type SpotDepthReq struct {
	Symbol *string `json:"symbol"` //YES
	Limit  *int    `json:"limit"`  //NO	默认 100; 最大 5000. 可选值:[5, 10, 20, 50, 100, 500, 1000, 5000]
}

type SpotDepthApi struct {
	SpotRestClient
	req *SpotDepthReq
}

func (api *SpotDepthApi) Symbol(Symbol string) *SpotDepthApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *SpotDepthApi) Limit(Limit int) *SpotDepthApi {
	api.req.Limit = GetPointer(Limit)
	return api
}

// symbol	STRING	YES
// orderId	LONG	NO
// origClientOrderId	STRING	NO
// newClientOrderId	STRING	NO	用户自定义的本次撤销操作的ID(注意不是被撤销的订单的自定义ID)。如无指定会自动赋值。
// cancelRestrictions	ENUM	NO	支持的值:
// 			ONLY_NEW - 如果订单状态为 NEW，订单取消将成功。
// 			ONLY_PARTIALLY_FILLED - 如果订单状态为 PARTIALLY_FILLED，订单取消将成功。
// recvWindow	LONG	NO	赋值不得大于 60000
// timestamp	LONG	YES

type SpotOrderDeleteReq struct {
	Symbol             *string `json:"symbol"`             //YES
	OrderId            *int64  `json:"orderId"`            //NO
	OrigClientOrderId  *string `json:"origClientOrderId"`  //NO
	NewClientOrderId   *string `json:"newClientOrderId"`   //NO	用户自定义的本次撤销操作的ID(注意不是被撤销的订单的自定义ID)。如无指定会自动赋值。
	CancelRestrictions *string `json:"cancelRestrictions"` //NO	支持的值: ONLY_NEW - 如果订单状态为 NEW，订单取消将成功。 ONLY_PARTIALLY_FILLED - 如果订单状态为 PARTIALLY_FILLED，订单取消将成功。
	RecvWindow         *int64  `json:"recvWindow"`         //NO	赋值不得大于 60000
	Timestamp          *int64  `json:"timestamp"`          //YES
}

type SpotOrderDeleteApi struct {
	SpotRestClient
	req *SpotOrderDeleteReq
}

func (api *SpotOrderDeleteApi) Symbol(Symbol string) *SpotOrderDeleteApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *SpotOrderDeleteApi) OrderId(OrderId int64) *SpotOrderDeleteApi {
	api.req.OrderId = GetPointer(OrderId)
	return api
}
func (api *SpotOrderDeleteApi) OrigClientOrderId(OrigClientOrderId string) *SpotOrderDeleteApi {
	api.req.OrigClientOrderId = GetPointer(OrigClientOrderId)
	return api
}
func (api *SpotOrderDeleteApi) NewClientOrderId(NewClientOrderId string) *SpotOrderDeleteApi {
	api.req.NewClientOrderId = GetPointer(NewClientOrderId)
	return api
}
func (api *SpotOrderDeleteApi) CancelRestrictions(CancelRestrictions string) *SpotOrderDeleteApi {
	api.req.CancelRestrictions = GetPointer(CancelRestrictions)
	return api
}
func (api *SpotOrderDeleteApi) RecvWindow(RecvWindow int64) *SpotOrderDeleteApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *SpotOrderDeleteApi) Timestamp(Timestamp int64) *SpotOrderDeleteApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

// symbol	STRING	YES
// isIsolated	STRING	NO	是否逐仓杠杆，"TRUE", "FALSE", 默认 "FALSE"
// orderId	LONG	NO
// origClientOrderId	STRING	NO
// newClientOrderId	STRING	NO	用于唯一识别此撤销订单，默认自动生成。
// recvWindow	LONG	NO	T赋值不能大于 60000
// timestamp	LONG	YES
type SpotMarginOrderDeleteReq struct {
	Symbol            *string `json:"symbol"`            //YES
	IsIsolated        *string `json:"isIsolated"`        //NO	是否逐仓杠杆，"TRUE", "FALSE", 默认 "FALSE"
	OrderId           *int64  `json:"orderId"`           //NO
	OrigClientOrderId *string `json:"origClientOrderId"` //NO
	NewClientOrderId  *string `json:"newClientOrderId"`  //NO	用于唯一识别此撤销订单，默认自动生成。
	RecvWindow        *int64  `json:"recvWindow"`        //NO	T赋值不能大于 60000
	Timestamp         *int64  `json:"timestamp"`         //YES
}

type SpotMarginOrderDeleteApi struct {
	SpotRestClient
	req *SpotMarginOrderDeleteReq
}

func (api *SpotMarginOrderDeleteApi) Symbol(Symbol string) *SpotMarginOrderDeleteApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *SpotMarginOrderDeleteApi) IsIsolated(IsIsolated string) *SpotMarginOrderDeleteApi {
	api.req.IsIsolated = GetPointer(IsIsolated)
	return api
}
func (api *SpotMarginOrderDeleteApi) OrderId(OrderId int64) *SpotMarginOrderDeleteApi {
	api.req.OrderId = GetPointer(OrderId)
	return api
}
func (api *SpotMarginOrderDeleteApi) OrigClientOrderId(OrigClientOrderId string) *SpotMarginOrderDeleteApi {
	api.req.OrigClientOrderId = GetPointer(OrigClientOrderId)
	return api
}
func (api *SpotMarginOrderDeleteApi) NewClientOrderId(NewClientOrderId string) *SpotMarginOrderDeleteApi {
	api.req.NewClientOrderId = GetPointer(NewClientOrderId)
	return api
}
func (api *SpotMarginOrderDeleteApi) RecvWindow(RecvWindow int64) *SpotMarginOrderDeleteApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *SpotMarginOrderDeleteApi) Timestamp(Timestamp int64) *SpotMarginOrderDeleteApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

// symbol	STRING	YES
// limit	INT	NO	默认 500; 最大值 1000.
type SpotTradesReq struct {
	Symbol *string `json:"symbol"` //YES
	Limit  *int    `json:"limit"`  //NO	默认 500; 最大值 1000.
}

type SpotTradesApi struct {
	SpotRestClient
	req *SpotTradesReq
}

func (api *SpotTradesApi) Symbol(Symbol string) *SpotTradesApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *SpotTradesApi) Limit(Limit int) *SpotTradesApi {
	api.req.Limit = GetPointer(Limit)
	return api
}

// symbol	STRING	YES
// limit	INT	NO	默认 500; 最大值 1000.
// fromId	LONG	NO	从哪一条成交id开始返回. 缺省返回最近的成交记录。

type SpotHistoricalTradesReq struct {
	Symbol *string `json:"symbol"` //YES
	Limit  *int    `json:"limit"`  //NO	默认 500; 最大值 1000.
	FromId *int64  `json:"fromId"` //NO	从哪一条成交id开始返回. 缺省返回最近的成交记录。
}

type SpotHistoricalTradesApi struct {
	SpotRestClient
	req *SpotHistoricalTradesReq
}

func (api *SpotHistoricalTradesApi) Symbol(Symbol string) *SpotHistoricalTradesApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *SpotHistoricalTradesApi) Limit(Limit int) *SpotHistoricalTradesApi {
	api.req.Limit = GetPointer(Limit)
	return api
}
func (api *SpotHistoricalTradesApi) FromId(FromId int64) *SpotHistoricalTradesApi {
	api.req.FromId = GetPointer(FromId)
	return api
}

// symbol	STRING	YES
// fromId	LONG	NO	从包含fromId的成交id开始返回结果
// startTime	LONG	NO	从该时刻之后的成交记录开始返回结果
// endTime	LONG	NO	返回该时刻为止的成交记录
// limit	INT	NO	默认 500; 最大 1000.
type SpotAggTradesReq struct {
	Symbol    *string `json:"symbol"`    //YES
	FromId    *int64  `json:"fromId"`    //NO	从包含fromId的成交id开始返回结果
	StartTime *int64  `json:"startTime"` //NO	从该时刻之后的成交记录开始返回结果
	EndTime   *int64  `json:"endTime"`   //NO	返回该时刻为止的成交记录
	Limit     *int    `json:"limit"`     //NO	默认 500; 最大 1000.
}

type SpotAggTradesApi struct {
	SpotRestClient
	req *SpotAggTradesReq
}

func (api *SpotAggTradesApi) Symbol(Symbol string) *SpotAggTradesApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *SpotAggTradesApi) FromId(FromId int64) *SpotAggTradesApi {
	api.req.FromId = GetPointer(FromId)
	return api
}
func (api *SpotAggTradesApi) StartTime(StartTime int64) *SpotAggTradesApi {
	api.req.StartTime = GetPointer(StartTime)
	return api
}
func (api *SpotAggTradesApi) EndTime(EndTime int64) *SpotAggTradesApi {
	api.req.EndTime = GetPointer(EndTime)
	return api
}
func (api *SpotAggTradesApi) Limit(Limit int) *SpotAggTradesApi {
	api.req.Limit = GetPointer(Limit)
	return api
}

// symbol	STRING	YES
type SpotAvgPriceReq struct {
	Symbol *string `json:"symbol"` //YES
}

type SpotAvgPriceApi struct {
	SpotRestClient
	req *SpotAvgPriceReq
}

func (api *SpotAvgPriceApi) Symbol(Symbol string) *SpotAvgPriceApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}

type SpotUiKlinesApi SpotKlinesApi

func (api *SpotUiKlinesApi) Symbol(Symbol string) *SpotUiKlinesApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *SpotUiKlinesApi) Interval(Interval string) *SpotUiKlinesApi {
	api.req.Interval = GetPointer(Interval)
	return api
}
func (api *SpotUiKlinesApi) StartTime(StartTime int64) *SpotUiKlinesApi {
	api.req.StartTime = GetPointer(StartTime)
	return api
}
func (api *SpotUiKlinesApi) EndTime(EndTime int64) *SpotUiKlinesApi {
	api.req.EndTime = GetPointer(EndTime)
	return api
}
func (api *SpotUiKlinesApi) Limit(Limit int) *SpotUiKlinesApi {
	api.req.Limit = GetPointer(Limit)
	return api
}

// symbol	1	2
// 		不提供symbol	80
// symbols	1-20	2
// 		21-100	40
// 		>= 101	80
// 		不提供symbol	80
// type	ENUM	NO	可接受的参数: FULL or MINI.
// 		如果不提供, 默认值为 FULL
type SpotTicker24hrReq struct {
	Symbol  *string   `json:"symbol"`  //NO
	Symbols *[]string `json:"symbols"` //NO
	Type    *string   `json:"type"`    //NO	可接受的参数: FULL or MINI. 如果不提供, 默认值为 FULL
}

type SpotTicker24hrApi struct {
	SpotRestClient
	req *SpotTicker24hrReq
}

func (api *SpotTicker24hrApi) Symbol(Symbol string) *SpotTicker24hrApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *SpotTicker24hrApi) Symbols(Symbols []string) *SpotTicker24hrApi {
	api.req.Symbols = GetPointer(Symbols)
	return api
}
func (api *SpotTicker24hrApi) Type(Type string) *SpotTicker24hrApi {
	api.req.Type = GetPointer(Type)
	return api
}

type SpotTickerBookTickerReq struct {
	Symbol  *string   `json:"symbol"`  //NO
	Symbols *[]string `json:"symbols"` //NO
}

type SpotTickerBookTickerApi struct {
	SpotRestClient
	req *SpotTickerBookTickerReq
}

func (api *SpotTickerBookTickerApi) Symbol(Symbol string) *SpotTickerBookTickerApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *SpotTickerBookTickerApi) Symbols(Symbols []string) *SpotTickerBookTickerApi {
	api.req.Symbols = GetPointer(Symbols)
	return api
}

// symbol	STRING	YES	提供 symbol或者symbols 其中之一
// symbols
// windowSize	ENUM	NO	默认为 1d
// 	windowSize 支持的值:
// 	如果是分钟: 1m,2m....59m
// 	如果是小时: 1h, 2h....23h
// 	如果是天: 1d...7d
//  不可以组合使用, 比如1d2h
// type	ENUM	NO	可接受的参数: FULL or MINI.
//  如果不提供, 默认值为 FULL

type SpotTickerReq struct {
	Symbol     *string   `json:"symbol"`     //NO	提供 symbol或者symbols 其中之一
	Symbols    *[]string `json:"symbols"`    //NO
	WindowSize *string   `json:"windowSize"` //NO	默认为 1d
	Type       *string   `json:"type"`       //NO	可接受的参数: FULL or MINI. 如果不提供, 默认值为 FULL
}

type SpotTickerApi struct {
	SpotRestClient
	req *SpotTickerReq
}

func (api *SpotTickerApi) Symbol(Symbol string) *SpotTickerApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *SpotTickerApi) Symbols(Symbols []string) *SpotTickerApi {
	api.req.Symbols = GetPointer(Symbols)
	return api
}
func (api *SpotTickerApi) WindowSize(WindowSize string) *SpotTickerApi {
	api.req.WindowSize = GetPointer(WindowSize)
	return api
}
func (api *SpotTickerApi) Type(Type string) *SpotTickerApi {
	api.req.Type = GetPointer(Type)
	return api
}


