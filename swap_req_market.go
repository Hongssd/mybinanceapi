package mybinanceapi

type SwapKlinesReq struct {
	Symbol    *string `json:"symbol"`    //YES
	Interval  *string `json:"interval"`  //YES	详见枚举定义：K线间隔
	StartTime *int64  `json:"startTime"` //NO
	EndTime   *int64  `json:"endTime"`   //NO
	Limit     *int    `json:"limit"`     //NO	默认 500; 最大 1000.
}
type SwapKlinesApi struct {
	client *SwapRestClient
	req    *SwapKlinesReq
}

func (api *SwapKlinesApi) Symbol(Symbol string) *SwapKlinesApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *SwapKlinesApi) Interval(Interval string) *SwapKlinesApi {
	api.req.Interval = GetPointer(Interval)
	return api
}
func (api *SwapKlinesApi) StartTime(StartTime int64) *SwapKlinesApi {
	api.req.StartTime = GetPointer(StartTime)
	return api
}
func (api *SwapKlinesApi) EndTime(EndTime int64) *SwapKlinesApi {
	api.req.EndTime = GetPointer(EndTime)
	return api
}
func (api *SwapKlinesApi) Limit(Limit int) *SwapKlinesApi {
	api.req.Limit = GetPointer(Limit)
	return api
}

type SwapTickerPriceReq struct {
	Symbol *string `json:"symbol"` //NO	交易对
	Pair   *string `json:"pair"`   //NO	标的交易对
}
type SwapTickerPriceApi struct {
	client *SwapRestClient
	req    *SwapTickerPriceReq
}

func (api *SwapTickerPriceApi) Symbol(Symbol string) *SwapTickerPriceApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *SwapTickerPriceApi) Pair(Pair string) *SwapTickerPriceApi {
	api.req.Pair = GetPointer(Pair)
	return api
}

type SwapDepthReq struct {
	Symbol *string `json:"symbol"` //YES	交易对
	Limit  *int    `json:"limit"`  //NO	默认 500; 可选值:[5, 10, 20, 50, 100, 500, 1000]
}
type SwapDepthApi struct {
	client *SwapRestClient
	req    *SwapDepthReq
}

func (api *SwapDepthApi) Symbol(Symbol string) *SwapDepthApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *SwapDepthApi) Limit(Limit int) *SwapDepthApi {
	api.req.Limit = GetPointer(Limit)
	return api
}
