package mybinanceapi

type SpotTickerPriceReq struct {
	Symbol *string `json:"symbol"` //YES
}
type SpotTickerPriceApi struct {
	client *SpotRestClient
	req    *SpotTickerPriceReq
}

func (api *SpotTickerPriceApi) Symbol(Symbol string) *SpotTickerPriceApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}

type SpotKlinesReq struct {
	Symbol    *string `json:"symbol"`    //YES
	Interval  *string `json:"interval"`  //YES	详见枚举定义：K线间隔
	StartTime *int64  `json:"startTime"` //NO
	EndTime   *int64  `json:"endTime"`   //NO
	Limit     *int    `json:"limit"`     //NO	默认 500; 最大 1000.
}
type SpotKlinesApi struct {
	client *SpotRestClient
	req    *SpotKlinesReq
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

type SpotDepthReq struct {
	Symbol *string `json:"symbol"` //YES
	Limit  *int    `json:"limit"`  //NO	默认 100; 最大 5000. 可选值:[5, 10, 20, 50, 100, 500, 1000, 5000]
}
type SpotDepthApi struct {
	client *SpotRestClient
	req    *SpotDepthReq
}

func (api *SpotDepthApi) Symbol(Symbol string) *SpotDepthApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *SpotDepthApi) Limit(Limit int) *SpotDepthApi {
	api.req.Limit = GetPointer(Limit)
	return api
}

type SpotTradesReq struct {
	Symbol *string `json:"symbol"` //YES
	Limit  *int    `json:"limit"`  //NO	默认 500; 最大值 1000.
}
type SpotTradesApi struct {
	client *SpotRestClient
	req    *SpotTradesReq
}

func (api *SpotTradesApi) Symbol(Symbol string) *SpotTradesApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *SpotTradesApi) Limit(Limit int) *SpotTradesApi {
	api.req.Limit = GetPointer(Limit)
	return api
}

type SpotHistoricalTradesReq struct {
	Symbol *string `json:"symbol"` //YES
	Limit  *int    `json:"limit"`  //NO	默认 500; 最大值 1000.
	FromId *int64  `json:"fromId"` //NO	从哪一条成交id开始返回. 缺省返回最近的成交记录。
}
type SpotHistoricalTradesApi struct {
	client *SpotRestClient
	req    *SpotHistoricalTradesReq
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

type SpotAggTradesReq struct {
	Symbol    *string `json:"symbol"`    //YES
	FromId    *int64  `json:"fromId"`    //NO	从包含fromId的成交id开始返回结果
	StartTime *int64  `json:"startTime"` //NO	从该时刻之后的成交记录开始返回结果
	EndTime   *int64  `json:"endTime"`   //NO	返回该时刻为止的成交记录
	Limit     *int    `json:"limit"`     //NO	默认 500; 最大 1000.
}
type SpotAggTradesApi struct {
	client *SpotRestClient
	req    *SpotAggTradesReq
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

type SpotAvgPriceReq struct {
	Symbol *string `json:"symbol"` //YES
}
type SpotAvgPriceApi struct {
	client *SpotRestClient
	req    *SpotAvgPriceReq
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

type SpotTicker24hrReq struct {
	Symbol  *string   `json:"symbol"`  //NO
	Symbols *[]string `json:"symbols"` //NO
	Type    *string   `json:"type"`    //NO	可接受的参数: FULL or MINI. 如果不提供, 默认值为 FULL
}
type SpotTicker24hrApi struct {
	client *SpotRestClient
	req    *SpotTicker24hrReq
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
	client *SpotRestClient
	req    *SpotTickerBookTickerReq
}

func (api *SpotTickerBookTickerApi) Symbol(Symbol string) *SpotTickerBookTickerApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *SpotTickerBookTickerApi) Symbols(Symbols []string) *SpotTickerBookTickerApi {
	api.req.Symbols = GetPointer(Symbols)
	return api
}

type SpotTickerReq struct {
	Symbol     *string   `json:"symbol"`     //NO	提供 symbol或者symbols 其中之一
	Symbols    *[]string `json:"symbols"`    //NO
	WindowSize *string   `json:"windowSize"` //NO	默认为 1d
	Type       *string   `json:"type"`       //NO	可接受的参数: FULL or MINI. 如果不提供, 默认值为 FULL
}
type SpotTickerApi struct {
	client *SpotRestClient
	req    *SpotTickerReq
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
