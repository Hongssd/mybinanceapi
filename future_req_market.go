package mybinanceapi

type FutureKlinesReq struct {
	Symbol    *string `json:"symbol"`    //YES
	Interval  *string `json:"interval"`  //YES	详见枚举定义：K线间隔
	StartTime *int64  `json:"startTime"` //NO
	EndTime   *int64  `json:"endTime"`   //NO
	Limit     *int    `json:"limit"`     //NO	默认 500; 最大 1000.
}

type FutureKlinesApi struct {
	client *FutureRestClient
	req    *FutureKlinesReq
}

func (api *FutureKlinesApi) Symbol(Symbol string) *FutureKlinesApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *FutureKlinesApi) Interval(Interval string) *FutureKlinesApi {
	api.req.Interval = GetPointer(Interval)
	return api
}
func (api *FutureKlinesApi) StartTime(StartTime int64) *FutureKlinesApi {
	api.req.StartTime = GetPointer(StartTime)
	return api
}
func (api *FutureKlinesApi) EndTime(EndTime int64) *FutureKlinesApi {
	api.req.EndTime = GetPointer(EndTime)
	return api
}
func (api *FutureKlinesApi) Limit(Limit int) *FutureKlinesApi {
	api.req.Limit = GetPointer(Limit)
	return api
}

type FutureDepthReq struct {
	Symbol *string `json:"symbol"` //YES	交易对
	Limit  *int    `json:"limit"`  //NO	默认 500; 可选值:[5, 10, 20, 50, 100, 500, 1000]
}

type FutureDepthApi struct {
	client *FutureRestClient
	req    *FutureDepthReq
}

// YES	交易对
func (api *FutureDepthApi) Symbol(Symbol string) *FutureDepthApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}

// NO	默认 500; 可选值:[5, 10, 20, 50, 100, 500, 1000]
func (api *FutureDepthApi) Limit(Limit int) *FutureDepthApi {
	api.req.Limit = GetPointer(Limit)
	return api
}

type FutureTradesReq struct {
	Symbol *string `json:"symbol"` //YES	交易对
	Limit  *int    `json:"limit"`  //NO	默认:500，最大1000
}

type FutureTradesApi struct {
	client *FutureRestClient
	req    *FutureTradesReq
}

// YES	交易对
func (api *FutureTradesApi) Symbol(Symbol string) *FutureTradesApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}

// NO	默认:500，最大1000
func (api *FutureTradesApi) Limit(Limit int) *FutureTradesApi {
	api.req.Limit = GetPointer(Limit)
	return api
}

type FutureHistoricalTradesReq struct {
	Symbol *string `json:"symbol"` //YES	交易对
	Limit  *int    `json:"limit"`  //NO	默认值:500 最大值:1000.
	FromId *int64  `json:"fromId"` //NO	从哪一条成交id开始返回. 缺省返回最近的成交记录
}

type FutureHistoricalTradesApi struct {
	client *FutureRestClient
	req    *FutureHistoricalTradesReq
}

// YES	交易对
func (api *FutureHistoricalTradesApi) Symbol(Symbol string) *FutureHistoricalTradesApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}

// NO	默认值:500 最大值:1000.
func (api *FutureHistoricalTradesApi) Limit(Limit int) *FutureHistoricalTradesApi {
	api.req.Limit = GetPointer(Limit)
	return api
}

// NO	从哪一条成交id开始返回. 缺省返回最近的成交记录
func (api *FutureHistoricalTradesApi) FromId(FromId int64) *FutureHistoricalTradesApi {
	api.req.FromId = GetPointer(FromId)
	return api
}

type FutureAggTradesReq struct {
	Symbol    *string `json:"symbol"`    //YES	交易对
	FromId    *int64  `json:"fromId"`    //NO	从包含fromID的成交开始返回结果
	StartTime *int64  `json:"startTime"` //NO	从该时刻之后的成交记录开始返回结果
	EndTime   *int64  `json:"endTime"`   //NO	返回该时刻为止的成交记录
	Limit     *int    `json:"limit"`     //NO	默认 500; 最大 1000.
}

type FutureAggTradesApi struct {
	client *FutureRestClient
	req    *FutureAggTradesReq
}

// YES	交易对
func (api *FutureAggTradesApi) Symbol(Symbol string) *FutureAggTradesApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}

// NO	从包含fromID的成交开始返回结果
func (api *FutureAggTradesApi) FromId(FromId int64) *FutureAggTradesApi {
	api.req.FromId = GetPointer(FromId)
	return api
}

// NO	从该时刻之后的成交记录开始返回结果
func (api *FutureAggTradesApi) StartTime(StartTime int64) *FutureAggTradesApi {
	api.req.StartTime = GetPointer(StartTime)
	return api
}

// NO	返回该时刻为止的成交记录
func (api *FutureAggTradesApi) EndTime(EndTime int64) *FutureAggTradesApi {
	api.req.EndTime = GetPointer(EndTime)
	return api
}

// NO	默认 500; 最大 1000.
func (api *FutureAggTradesApi) Limit(Limit int) *FutureAggTradesApi {
	api.req.Limit = GetPointer(Limit)
	return api
}

// symbol	STRING	NO	交易对
type FuturePremiumIndexReq struct {
	Symbol *string `json:"symbol"` //NO	交易对
}

type FuturePremiumIndexApi struct {
	client *FutureRestClient
	req    *FuturePremiumIndexReq
}

// NO	交易对
func (api *FuturePremiumIndexApi) Symbol(Symbol string) *FuturePremiumIndexApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}

type FutureFundingRateReq struct {
	Symbol    *string `json:"symbol"`    //交易对
	StartTime *int64  `json:"startTime"` //起始时间
	EndTime   *int64  `json:"endTime"`   //结束时间
	Limit     *int    `json:"limit"`     //默认值:100 最大值:1000
}

type FutureFundingRateApi struct {
	client *FutureRestClient
	req    *FutureFundingRateReq
}

// NO	交易对
func (api *FutureFundingRateApi) Symbol(Symbol string) *FutureFundingRateApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}

// NO	起始时间
func (api *FutureFundingRateApi) StartTime(StartTime int64) *FutureFundingRateApi {
	api.req.StartTime = GetPointer(StartTime)
	return api
}

// NO	结束时间
func (api *FutureFundingRateApi) EndTime(EndTime int64) *FutureFundingRateApi {
	api.req.EndTime = GetPointer(EndTime)
	return api
}

// NO	默认值:100 最大值:1000
func (api *FutureFundingRateApi) Limit(Limit int) *FutureFundingRateApi {
	api.req.Limit = GetPointer(Limit)
	return api
}

type FutureFundingInfoReq struct {
}

type FutureFundingInfoApi struct {
	client *FutureRestClient
	req    *FutureFundingInfoReq
}

// symbol	STRING	NO	交易对
type FutureTicker24hrReq struct {
	Symbol *string `json:"symbol"` //NO	交易对
}

type FutureTicker24hrApi struct {
	client *FutureRestClient
	req    *FutureTicker24hrReq
}

// NO	交易对
func (api *FutureTicker24hrApi) Symbol(Symbol string) *FutureTicker24hrApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}

// symbol	STRING	NO	交易对
type FutureTickerPriceReq struct {
	Symbol *string `json:"symbol"` //NO	交易对
}

type FutureTickerPriceApi struct {
	client *FutureRestClient
	req    *FutureTickerPriceReq
}

// NO	交易对
func (api *FutureTickerPriceApi) Symbol(Symbol string) *FutureTickerPriceApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}

// symbol	STRING	NO	交易对

type FutureTickerBookTickerReq struct {
	Symbol *string `json:"symbol"` //NO	交易对
}

type FutureTickerBookTickerApi struct {
	client *FutureRestClient
	req    *FutureTickerBookTickerReq
}

// NO	交易对
func (api *FutureTickerBookTickerApi) Symbol(Symbol string) *FutureTickerBookTickerApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}

type FutureDataBasisReq struct {
	Pair         *string `json:"pair"`         //YES	BTCUSDT
	ContractType *string `json:"contractType"` //YES	CURRENT_QUARTER, NEXT_QUARTER, PERPETUAL
	Period       *string `json:"period"`       //YES	"5m","15m","30m","1h","2h","4h","6h","12h","1d"
	Limit        *int64  `json:"limit"`        //NO	Default 30,Max 500
	StartTime    *int64  `json:"startTime"`    //NO
	EndTime      *int64  `json:"endTime"`      //NO
}

type FutureDataBasisApi struct {
	client *FutureRestClient
	req    *FutureDataBasisReq
}

// YES	BTCUSDT
func (api *FutureDataBasisApi) Pair(Pair string) *FutureDataBasisApi {
	api.req.Pair = GetPointer(Pair)
	return api
}

// YES	CURRENT_QUARTER, NEXT_QUARTER, PERPETUAL
func (api *FutureDataBasisApi) ContractType(ContractType string) *FutureDataBasisApi {
	api.req.ContractType = GetPointer(ContractType)
	return api
}

// YES	"5m","15m","30m","1h","2h","4h","6h","12h","1d"
func (api *FutureDataBasisApi) Period(Period string) *FutureDataBasisApi {
	api.req.Period = GetPointer(Period)
	return api
}

// NO	Default 30,Max 500
func (api *FutureDataBasisApi) Limit(Limit int64) *FutureDataBasisApi {
	api.req.Limit = GetPointer(Limit)
	return api
}

// NO
func (api *FutureDataBasisApi) StartTime(StartTime int64) *FutureDataBasisApi {
	api.req.StartTime = GetPointer(StartTime)
	return api
}

// NO
func (api *FutureDataBasisApi) EndTime(EndTime int64) *FutureDataBasisApi {
	api.req.EndTime = GetPointer(EndTime)
	return api
}
