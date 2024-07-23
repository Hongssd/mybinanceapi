package mybinanceapi

import "github.com/shopspring/decimal"

type SpotAccountApiTradingStatusReq struct {
	RecvWindow *int64 `json:"recvWindow"` //NO 	赋值不能大于 60000
	Timestamp  *int64 `json:"timestamp"`  //YES
}
type SpotAccountApiTradingStatusApi struct {
	client *SpotRestClient
	req    *SpotAccountApiTradingStatusReq
}

func (api *SpotAccountApiTradingStatusApi) RecvWindow(RecvWindow int64) *SpotAccountApiTradingStatusApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *SpotAccountApiTradingStatusApi) Timestamp(Timestamp int64) *SpotAccountApiTradingStatusApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SpotAssetGetFundingAssetReq struct {
	Asset            *string `json:"asset"`            //NO
	NeedBtcValuation *string `json:"needBtcValuation"` //NO	true or false
	RecvWindow       *int64  `json:"recvWindow"`       //NO
	Timestamp        *int64  `json:"timestamp"`        //YES
}
type SpotAssetGetFundingAssetApi struct {
	client *SpotRestClient
	req    *SpotAssetGetFundingAssetReq
}

func (api *SpotAssetGetFundingAssetApi) Asset(Asset string) *SpotAssetGetFundingAssetApi {
	api.req.Asset = GetPointer(Asset)
	return api
}
func (api *SpotAssetGetFundingAssetApi) NeedBtcValuation(NeedBtcValuation string) *SpotAssetGetFundingAssetApi {
	api.req.NeedBtcValuation = GetPointer(NeedBtcValuation)
	return api
}
func (api *SpotAssetGetFundingAssetApi) RecvWindow(RecvWindow int64) *SpotAssetGetFundingAssetApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *SpotAssetGetFundingAssetApi) Timestamp(Timestamp int64) *SpotAssetGetFundingAssetApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SpotAccountReq struct {
	OmitZeroBalances *bool  `json:"omitZeroBalances"` //NO	默认 false, 设为 true 时，返回结果不包含为 0 的资产
	RecvWindow       *int64 `json:"recvWindow"`
	Timestamp        *int64 `json:"timestamp"`
}
type SpotAccountApi struct {
	client *SpotRestClient
	req    *SpotAccountReq
}

func (api *SpotAccountApi) OmitZeroBalances(OmitZeroBalances bool) *SpotAccountApi {
	api.req.OmitZeroBalances = GetPointer(OmitZeroBalances)
	return api
}
func (api *SpotAccountApi) RecvWindow(RecvWindow int64) *SpotAccountApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *SpotAccountApi) Timestamp(Timestamp int64) *SpotAccountApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

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
	client *SpotRestClient
	req    *SpotAssetTransferPostReq
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
	client *SpotRestClient
	req    *SpotAssetTransferGetReq
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

type SpotAssetTradeFeeReq struct {
	Symbol     *string `json:"symbol"`     //NO
	RecvWindow *int64  `json:"recvWindow"` //NO
	Timestamp  *int64  `json:"timestamp"`  //YES
}
type SpotAssetTradeFeeApi struct {
	client *SpotRestClient
	req    *SpotAssetTradeFeeReq
}

func (api *SpotAssetTradeFeeApi) Symbol(Symbol string) *SpotAssetTradeFeeApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *SpotAssetTradeFeeApi) RecvWindow(RecvWindow int64) *SpotAssetTradeFeeApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *SpotAssetTradeFeeApi) Timestamp(Timestamp int64) *SpotAssetTradeFeeApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}
