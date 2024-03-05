package mybinanceapi

import "github.com/shopspring/decimal"

//FutureAccount
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

//FuturePositionSideDualGet
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

//FutureMultiAssetsMarginGet
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

//FuturePositionSideDualPost
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

//FutureMultiAssetsMarginPost
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

//FutureLeverage
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

//FutureMarginType
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

//FutureLeverageBracket
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

type FuturePingReq struct {
}
type FuturePingApi struct {
	client *FutureRestClient
	req    *FuturePingReq
}

type FutureServerTimeReq struct {
}
type FutureServerTimeApi struct {
	client *FutureRestClient
	req    *FutureServerTimeReq
}

type FutureExchangeInfoReq struct {
}
type FutureExchangeInfoApi struct {
	client *FutureRestClient
	req    *FutureExchangeInfoReq
}

type FutureOpenOrdersReq struct {
	Symbol     *string `json:"symbol"` //No	交易对
	Recvwindow *int64  `json:"recvWindow"`
	Timestamp  *int64  `json:"timestamp"`
}

type FutureOpenOrdersApi struct {
	client *FutureRestClient
	req    *FutureOpenOrdersReq
}

func (api *FutureOpenOrdersApi) Symbol(Symbol string) *FutureOpenOrdersApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *FutureOpenOrdersApi) Recvwindow(Recvwindow int64) *FutureOpenOrdersApi {
	api.req.Recvwindow = GetPointer(Recvwindow)
	return api
}
func (api *FutureOpenOrdersApi) Timestamp(Timestamp int64) *FutureOpenOrdersApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

// symbol	STRING	YES	交易对
// orderId	LONG	NO	只返回此orderID及之后的订单，缺省返回最近的订单
// startTime	LONG	NO	起始时间
// endTime	LONG	NO	结束时间
// limit	INT	NO	返回的结果集数量 默认值:500 最大值:1000
// recvWindow	LONG	NO
// timestamp	LONG	YES
type FutureAllOrdersReq struct {
	Symbol     *string `json:"symbol"`    //No	交易对
	OrderId    *int64  `json:"orderId"`   //No	只返回此orderID及之后的订单，缺省返回最近的订单
	StartTime  *int64  `json:"startTime"` //No	起始时间
	EndTime    *int64  `json:"endTime"`   //No	结束时间
	Limit      *int64  `json:"limit"`     //No	返回的结果集数量 默认值:500 最大值:1000
	Recvwindow *int64  `json:"recvWindow"`
	Timestamp  *int64  `json:"timestamp"`
}

type FutureAllOrdersApi struct {
	client *FutureRestClient
	req    *FutureAllOrdersReq
}

func (api *FutureAllOrdersApi) Symbol(Symbol string) *FutureAllOrdersApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *FutureAllOrdersApi) OrderId(OrderId int64) *FutureAllOrdersApi {
	api.req.OrderId = GetPointer(OrderId)
	return api
}
func (api *FutureAllOrdersApi) StartTime(StartTime int64) *FutureAllOrdersApi {
	api.req.StartTime = GetPointer(StartTime)
	return api
}
func (api *FutureAllOrdersApi) EndTime(EndTime int64) *FutureAllOrdersApi {
	api.req.EndTime = GetPointer(EndTime)
	return api
}
func (api *FutureAllOrdersApi) Limit(Limit int64) *FutureAllOrdersApi {
	api.req.Limit = GetPointer(Limit)
	return api
}
func (api *FutureAllOrdersApi) Recvwindow(Recvwindow int64) *FutureAllOrdersApi {
	api.req.Recvwindow = GetPointer(Recvwindow)
	return api
}
func (api *FutureAllOrdersApi) Timestamp(Timestamp int64) *FutureAllOrdersApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

// symbol	STRING	YES	交易对
// side	ENUM	YES	买卖方向 SELL, BUY
// positionSide	ENUM	NO	持仓方向，单向持仓模式下非必填，默认且仅可填BOTH;在双向持仓模式下必填,且仅可选择 LONG 或 SHORT
// type	ENUM	YES	订单类型 LIMIT, MARKET, STOP, TAKE_PROFIT, STOP_MARKET, TAKE_PROFIT_MARKET, TRAILING_STOP_MARKET
// reduceOnly	STRING	NO	true, false; 非双开模式下默认false；双开模式下不接受此参数； 使用closePosition不支持此参数。
// quantity	DECIMAL	NO	下单数量,使用closePosition不支持此参数。
// price	DECIMAL	NO	委托价格
// newClientOrderId	STRING	NO	用户自定义的订单号，不可以重复出现在挂单中。如空缺系统会自动赋值。必须满足正则规则 ^[\.A-Z\:/a-z0-9_-]{1,36}$
// stopPrice	DECIMAL	NO	触发价, 仅 STOP, STOP_MARKET, TAKE_PROFIT, TAKE_PROFIT_MARKET 需要此参数
// closePosition	STRING	NO	true, false；触发后全部平仓，仅支持STOP_MARKET和TAKE_PROFIT_MARKET；不与quantity合用；自带只平仓效果，不与reduceOnly 合用
// activationPrice	DECIMAL	NO	追踪止损激活价格，仅TRAILING_STOP_MARKET 需要此参数, 默认为下单当前市场价格(支持不同workingType)
// callbackRate	DECIMAL	NO	追踪止损回调比例，可取值范围[0.1, 5],其中 1代表1% ,仅TRAILING_STOP_MARKET 需要此参数
// timeInForce	ENUM	NO	有效方法
// workingType	ENUM	NO	stopPrice 触发类型: MARK_PRICE(标记价格), CONTRACT_PRICE(合约最新价). 默认 CONTRACT_PRICE
// priceProtect	STRING	NO	条件单触发保护："TRUE","FALSE", 默认"FALSE". 仅 STOP, STOP_MARKET, TAKE_PROFIT, TAKE_PROFIT_MARKET 需要此参数
// newOrderRespType	ENUM	NO	"ACK", "RESULT", 默认 "ACK"
// recvWindow	LONG	NO
// timestamp	LONG	YES

type FutureOrderPostReq struct {
	Symbol           *string          `json:"symbol"`           //Yes	交易对
	Side             *string          `json:"side"`             //Yes	买卖方向 SELL, BUY
	PositionSide     *string          `json:"positionSide"`     //No	持仓方向，单向持仓模式下非必填，默认且仅可填BOTH;在双向持仓模式下必填,且仅可选择 LONG 或 SHORT
	Type             *string          `json:"type"`             //Yes	订单类型 LIMIT, MARKET, STOP, TAKE_PROFIT, STOP_MARKET, TAKE_PROFIT_MARKET, TRAILING_STOP_MARKET
	ReduceOnly       *string          `json:"reduceOnly"`       //No	true, false; 非双开模式下默认false；双开模式下不接受此参数； 使用closePosition不支持此参数。
	Quantity         *decimal.Decimal `json:"quantity"`         //No	下单数量,使用closePosition不支持此参数。
	Price            *decimal.Decimal `json:"price"`            //No	委托价格
	NewClientOrderId *string          `json:"newClientOrderId"` //No	用户自定义的订单号，不可以重复出现在挂单中。如空缺系统会自动赋值。必须满足正则规则 ^[\.A-Z\:/a-z0-9_-]{1,36}$
	StopPrice        *decimal.Decimal `json:"stopPrice"`        //No	触发价, 仅 STOP, STOP_MARKET, TAKE_PROFIT, TAKE_PROFIT_MARKET 需要此参数
	ClosePosition    *string          `json:"closePosition"`    //No	true, false；触发后全部平仓，仅支持STOP_MARKET和TAKE_PROFIT_MARKET；不与quantity合用；自带只平仓效果，不与reduceOnly 合用
	ActivationPrice  *decimal.Decimal `json:"activationPrice"`  //No	追踪止损激活价格，仅TRAILING_STOP_MARKET 需要此参数, 默认为下单当前市场价格(支持不同workingType)
	CallbackRate     *decimal.Decimal `json:"callbackRate"`     //No	追踪止损回调比例，可取值范围[0.1, 5],其中 1代表1% ,仅TRAILING_STOP_MARKET 需要此参数
	TimeInForce      *string          `json:"timeInForce"`      //No	有效方法
	WorkingType      *string          `json:"workingType"`      //No	stopPrice 触发类型: MARK_PRICE(标记价格), CONTRACT_PRICE(合约最新价). 默认 CONTRACT_PRICE
	PriceProtect     *string          `json:"priceProtect"`     //No	条件单触发保护："TRUE","FALSE", 默认"FALSE". 仅 STOP, STOP_MARKET, TAKE_PROFIT, TAKE_PROFIT_MARKET 需要此参数
	NewOrderRespType *string          `json:"newOrderRespType"` //No	"ACK", "RESULT", 默认 "ACK"
	Recvwindow       *int64           `json:"recvWindow"`       //No
	Timestamp        *int64           `json:"timestamp"`        //Yes
}

type FutureOrderPostApi struct {
	client *FutureRestClient
	req    *FutureOrderPostReq
}

func (api *FutureOrderPostApi) Symbol(Symbol string) *FutureOrderPostApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *FutureOrderPostApi) Side(Side string) *FutureOrderPostApi {
	api.req.Side = GetPointer(Side)
	return api
}
func (api *FutureOrderPostApi) PositionSide(PositionSide string) *FutureOrderPostApi {
	api.req.PositionSide = GetPointer(PositionSide)
	return api
}
func (api *FutureOrderPostApi) Type(Type string) *FutureOrderPostApi {
	api.req.Type = GetPointer(Type)
	return api
}
func (api *FutureOrderPostApi) ReduceOnly(ReduceOnly string) *FutureOrderPostApi {
	api.req.ReduceOnly = GetPointer(ReduceOnly)
	return api
}
func (api *FutureOrderPostApi) Quantity(Quantity decimal.Decimal) *FutureOrderPostApi {
	api.req.Quantity = GetPointer(Quantity)
	return api
}
func (api *FutureOrderPostApi) Price(Price decimal.Decimal) *FutureOrderPostApi {
	api.req.Price = GetPointer(Price)
	return api
}
func (api *FutureOrderPostApi) NewClientOrderId(NewClientOrderId string) *FutureOrderPostApi {
	api.req.NewClientOrderId = GetPointer(NewClientOrderId)
	return api
}
func (api *FutureOrderPostApi) StopPrice(StopPrice decimal.Decimal) *FutureOrderPostApi {
	api.req.StopPrice = GetPointer(StopPrice)
	return api
}

func (api *FutureOrderPostApi) ClosePosition(ClosePosition string) *FutureOrderPostApi {
	api.req.ClosePosition = GetPointer(ClosePosition)
	return api
}
func (api *FutureOrderPostApi) ActivationPrice(ActivationPrice decimal.Decimal) *FutureOrderPostApi {
	api.req.ActivationPrice = GetPointer(ActivationPrice)
	return api
}
func (api *FutureOrderPostApi) CallbackRate(CallbackRate decimal.Decimal) *FutureOrderPostApi {
	api.req.CallbackRate = GetPointer(CallbackRate)
	return api
}
func (api *FutureOrderPostApi) TimeInForce(TimeInForce string) *FutureOrderPostApi {
	api.req.TimeInForce = GetPointer(TimeInForce)
	return api
}
func (api *FutureOrderPostApi) WorkingType(WorkingType string) *FutureOrderPostApi {
	api.req.WorkingType = GetPointer(WorkingType)
	return api
}
func (api *FutureOrderPostApi) PriceProtect(PriceProtect string) *FutureOrderPostApi {
	api.req.PriceProtect = GetPointer(PriceProtect)
	return api
}
func (api *FutureOrderPostApi) NewOrderRespType(NewOrderRespType string) *FutureOrderPostApi {
	api.req.NewOrderRespType = GetPointer(NewOrderRespType)
	return api
}
func (api *FutureOrderPostApi) Recvwindow(Recvwindow int64) *FutureOrderPostApi {
	api.req.Recvwindow = GetPointer(Recvwindow)
	return api
}
func (api *FutureOrderPostApi) Timestamp(Timestamp int64) *FutureOrderPostApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

// symbol	STRING	YES	交易对
// orderId	LONG	NO	系统订单号
// origClientOrderId	STRING	NO	用户自定义的订单号
// recvWindow	LONG	NO
// timestamp	LONG	YES
type FutureOrderGetReq struct {
	Symbol            *string `json:"symbol"`            //YES 交易对
	OrderId           *int64  `json:"orderId"`           //NO 系统订单号
	OrigClientOrderId *string `json:"origClientOrderId"` //NO 用户自定义的订单号
	RecvWindow        *int64  `json:"recvWindow"`        //NO
	Timestamp         *int64  `json:"timestamp"`         //YES
}

type FutureOrderGetApi struct {
	client *FutureRestClient
	req    *FutureOrderGetReq
}

func (api *FutureOrderGetApi) Symbol(Symbol string) *FutureOrderGetApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *FutureOrderGetApi) OrderId(OrderId int64) *FutureOrderGetApi {
	api.req.OrderId = GetPointer(OrderId)
	return api
}
func (api *FutureOrderGetApi) OrigClientOrderId(OrigClientOrderId string) *FutureOrderGetApi {
	api.req.OrigClientOrderId = GetPointer(OrigClientOrderId)
	return api
}
func (api *FutureOrderGetApi) Recvwindow(Recvwindow int64) *FutureOrderGetApi {
	api.req.RecvWindow = GetPointer(Recvwindow)
	return api
}
func (api *FutureOrderGetApi) Timestamp(Timestamp int64) *FutureOrderGetApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type FutureCommissionRateReq struct {
	Symbol     *string `json:"symbol"`     //NO	交易对
	RecvWindow *int64  `json:"recvWindow"` //NO
	Timestamp  *int64  `json:"timestamp"`  //YES
}
type FutureCommissionRateApi struct {
	client *FutureRestClient
	req    *FutureCommissionRateReq
}

func (api *FutureCommissionRateApi) Symbol(Symbol string) *FutureCommissionRateApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *FutureCommissionRateApi) Recvwindow(Recvwindow int64) *FutureCommissionRateApi {
	api.req.RecvWindow = GetPointer(Recvwindow)
	return api
}
func (api *FutureCommissionRateApi) Timestamp(Timestamp int64) *FutureCommissionRateApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

// symbol	STRING	YES	交易对
// orderId	LONG	NO	必须要和参数symbol一起使用
// startTime	LONG	NO	起始时间
// endTime	LONG	NO	结束时间
// fromId	LONG	NO	返回该fromId及之后的成交，缺省返回最近的成交
// limit	INT	NO	返回的结果集数量 默认值:500 最大值:1000.
// recvWindow	LONG	NO
// timestamp	LONG	YES
type FutureUserTradesReq struct {
	Symbol     *string `json:"symbol"`     //YES	交易对
	OrderId    *int64  `json:"orderId"`    //NO	必须要和参数symbol一起使用
	StartTime  *int64  `json:"startTime"`  //NO	起始时间
	EndTime    *int64  `json:"endTime"`    //NO	结束时间
	FromId     *int64  `json:"fromId"`     //NO	返回该fromId及之后的成交，缺省返回最近的成交
	Limit      *int64  `json:"limit"`      //NO	返回的结果集数量 默认值:500 最大值:1000.
	RecvWindow *int64  `json:"recvWindow"` //NO
	Timestamp  *int64  `json:"timestamp"`  //YES
}

type FutureUserTradesApi struct {
	client *FutureRestClient
	req    *FutureUserTradesReq
}

func (api *FutureUserTradesApi) Symbol(Symbol string) *FutureUserTradesApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *FutureUserTradesApi) OrderId(OrderId int64) *FutureUserTradesApi {
	api.req.OrderId = GetPointer(OrderId)
	return api
}
func (api *FutureUserTradesApi) StartTime(StartTime int64) *FutureUserTradesApi {
	api.req.StartTime = GetPointer(StartTime)
	return api
}
func (api *FutureUserTradesApi) EndTime(EndTime int64) *FutureUserTradesApi {
	api.req.EndTime = GetPointer(EndTime)
	return api
}
func (api *FutureUserTradesApi) FromId(FromId int64) *FutureUserTradesApi {
	api.req.FromId = GetPointer(FromId)
	return api
}
func (api *FutureUserTradesApi) Limit(Limit int64) *FutureUserTradesApi {
	api.req.Limit = GetPointer(Limit)
	return api
}
func (api *FutureUserTradesApi) Recvwindow(Recvwindow int64) *FutureUserTradesApi {
	api.req.RecvWindow = GetPointer(Recvwindow)
	return api
}
func (api *FutureUserTradesApi) Timestamp(Timestamp int64) *FutureUserTradesApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

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

// symbol	STRING	YES	交易对
// orderId	LONG	NO	系统订单号
// origClientOrderId	STRING	NO	用户自定义的订单号
// recvWindow	LONG	NO
// timestamp	LONG	YES

type FutureOrderDeleteReq struct {
	Symbol            *string `json:"symbol"`            //YES 交易对
	OrderId           *int64  `json:"orderId"`           //NO 系统订单号
	OrigClientOrderId *string `json:"origClientOrderId"` //NO 用户自定义的订单号
	RecvWindow        *int64  `json:"recvWindow"`        //NO
	Timestamp         *int64  `json:"timestamp"`         //YES
}

type FutureOrderDeleteApi struct {
	client *FutureRestClient
	req    *FutureOrderDeleteReq
}

func (api *FutureOrderDeleteApi) Symbol(Symbol string) *FutureOrderDeleteApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *FutureOrderDeleteApi) OrderId(OrderId int64) *FutureOrderDeleteApi {
	api.req.OrderId = GetPointer(OrderId)
	return api
}
func (api *FutureOrderDeleteApi) OrigClientOrderId(OrigClientOrderId string) *FutureOrderDeleteApi {
	api.req.OrigClientOrderId = GetPointer(OrigClientOrderId)
	return api
}
func (api *FutureOrderDeleteApi) Recvwindow(Recvwindow int64) *FutureOrderDeleteApi {
	api.req.RecvWindow = GetPointer(Recvwindow)
	return api
}
func (api *FutureOrderDeleteApi) Timestamp(Timestamp int64) *FutureOrderDeleteApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

// symbol	STRING	YES	交易对
// orderIdList	LIST<LONG>	NO	系统订单号, 最多支持10个订单

// 比如[1234567,2345678]
// origClientOrderIdList	LIST<STRING>	NO	用户自定义的订单号, 最多支持10个订单

// 比如["my_id_1","my_id_2"] 需要encode双引号。逗号后面没有空格。
// recvWindow	LONG	NO
// timestamp	LONG	YES

type FutureBatchOrdersDeleteReq struct {
	Symbol                *string   `json:"symbol"`                //YES 交易对
	OrderIdList           *[]int64  `json:"orderIdList"`           //NO 系统订单号, 最多支持10个订单
	OrigClientOrderIdList *[]string `json:"origClientOrderIdList"` //NO 用户自定义的订单号, 最多支持10个订单
	RecvWindow            *int64    `json:"recvWindow"`            //NO
	Timestamp             *int64    `json:"timestamp"`             //YES
}

type FutureBatchOrdersDeleteApi struct {
	client *FutureRestClient
	req    *FutureBatchOrdersDeleteReq
}

func (api *FutureBatchOrdersDeleteApi) Symbol(Symbol string) *FutureBatchOrdersDeleteApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *FutureBatchOrdersDeleteApi) OrderIdList(OrderIdList []int64) *FutureBatchOrdersDeleteApi {
	api.req.OrderIdList = GetPointer(OrderIdList)
	return api
}
func (api *FutureBatchOrdersDeleteApi) OrigClientOrderIdList(OrigClientOrderIdList []string) *FutureBatchOrdersDeleteApi {
	api.req.OrigClientOrderIdList = GetPointer(OrigClientOrderIdList)
	return api
}
func (api *FutureBatchOrdersDeleteApi) Recvwindow(Recvwindow int64) *FutureBatchOrdersDeleteApi {
	api.req.RecvWindow = GetPointer(Recvwindow)
	return api
}
func (api *FutureBatchOrdersDeleteApi) Timestamp(Timestamp int64) *FutureBatchOrdersDeleteApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

// symbol	STRING	YES	交易对
// limit	INT	NO	默认 500; 可选值:[5, 10, 20, 50, 100, 500, 1000]

type FutureDepthReq struct {
	Symbol *string `json:"symbol"` //YES	交易对
	Limit  *int    `json:"limit"`  //NO	默认 500; 可选值:[5, 10, 20, 50, 100, 500, 1000]
}

type FutureDepthApi struct {
	client *FutureRestClient
	req    *FutureDepthReq
}

//YES	交易对
func (api *FutureDepthApi) Symbol(Symbol string) *FutureDepthApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}

//NO	默认 500; 可选值:[5, 10, 20, 50, 100, 500, 1000]
func (api *FutureDepthApi) Limit(Limit int) *FutureDepthApi {
	api.req.Limit = GetPointer(Limit)
	return api
}

// symbol	STRING	YES	交易对
// limit	INT	NO	默认:500，最大1000
type FutureTradesReq struct {
	Symbol *string `json:"symbol"` //YES	交易对
	Limit  *int    `json:"limit"`  //NO	默认:500，最大1000
}

type FutureTradesApi struct {
	client *FutureRestClient
	req    *FutureTradesReq
}

//YES	交易对
func (api *FutureTradesApi) Symbol(Symbol string) *FutureTradesApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}

//NO	默认:500，最大1000
func (api *FutureTradesApi) Limit(Limit int) *FutureTradesApi {
	api.req.Limit = GetPointer(Limit)
	return api
}

// symbol	STRING	YES	交易对
// limit	INT	NO	默认值:500 最大值:1000.
// fromId	LONG	NO	从哪一条成交id开始返回. 缺省返回最近的成交记录
type FutureHistoricalTradesReq struct {
	Symbol *string `json:"symbol"` //YES	交易对
	Limit  *int    `json:"limit"`  //NO	默认值:500 最大值:1000.
	FromId *int64  `json:"fromId"` //NO	从哪一条成交id开始返回. 缺省返回最近的成交记录
}

type FutureHistoricalTradesApi struct {
	client *FutureRestClient
	req    *FutureHistoricalTradesReq
}

//YES	交易对
func (api *FutureHistoricalTradesApi) Symbol(Symbol string) *FutureHistoricalTradesApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}

//NO	默认值:500 最大值:1000.
func (api *FutureHistoricalTradesApi) Limit(Limit int) *FutureHistoricalTradesApi {
	api.req.Limit = GetPointer(Limit)
	return api
}

//NO	从哪一条成交id开始返回. 缺省返回最近的成交记录
func (api *FutureHistoricalTradesApi) FromId(FromId int64) *FutureHistoricalTradesApi {
	api.req.FromId = GetPointer(FromId)
	return api
}

// symbol	STRING	YES	交易对
// fromId	LONG	NO	从包含fromID的成交开始返回结果
// startTime	LONG	NO	从该时刻之后的成交记录开始返回结果
// endTime	LONG	NO	返回该时刻为止的成交记录
// limit	INT	NO	默认 500; 最大 1000.
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

//YES	交易对
func (api *FutureAggTradesApi) Symbol(Symbol string) *FutureAggTradesApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}

//NO	从包含fromID的成交开始返回结果
func (api *FutureAggTradesApi) FromId(FromId int64) *FutureAggTradesApi {
	api.req.FromId = GetPointer(FromId)
	return api
}

//NO	从该时刻之后的成交记录开始返回结果
func (api *FutureAggTradesApi) StartTime(StartTime int64) *FutureAggTradesApi {
	api.req.StartTime = GetPointer(StartTime)
	return api
}

//NO	返回该时刻为止的成交记录
func (api *FutureAggTradesApi) EndTime(EndTime int64) *FutureAggTradesApi {
	api.req.EndTime = GetPointer(EndTime)
	return api
}

//NO	默认 500; 最大 1000.
func (api *FutureAggTradesApi) Limit(Limit int) *FutureAggTradesApi {
	api.req.Limit = GetPointer(Limit)
	return api
}

//symbol	STRING	NO	交易对
type FuturePremiumIndexReq struct {
	Symbol *string `json:"symbol"` //NO	交易对
}

type FuturePremiumIndexApi struct {
	client *FutureRestClient
	req    *FuturePremiumIndexReq
}

//NO	交易对
func (api *FuturePremiumIndexApi) Symbol(Symbol string) *FuturePremiumIndexApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}

// symbol	STRING	NO	交易对
// startTime	LONG	NO	起始时间
// endTime	LONG	NO	结束时间
// limit	INT	NO	默认值:100 最大值:1000
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

//NO	交易对
func (api *FutureFundingRateApi) Symbol(Symbol string) *FutureFundingRateApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}

//NO	起始时间
func (api *FutureFundingRateApi) StartTime(StartTime int64) *FutureFundingRateApi {
	api.req.StartTime = GetPointer(StartTime)
	return api
}

//NO	结束时间
func (api *FutureFundingRateApi) EndTime(EndTime int64) *FutureFundingRateApi {
	api.req.EndTime = GetPointer(EndTime)
	return api
}

//NO	默认值:100 最大值:1000
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

//NO	交易对
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

//NO	交易对
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

//NO	交易对
func (api *FutureTickerBookTickerApi) Symbol(Symbol string) *FutureTickerBookTickerApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}

// pair	STRING	YES	BTCUSDT
// contractType	ENUM	YES	CURRENT_QUARTER, NEXT_QUARTER, PERPETUAL
// period	ENUM	YES	"5m","15m","30m","1h","2h","4h","6h","12h","1d"
// limit	LONG	NO	Default 30,Max 500
// startTime	LONG	NO
// endTime	LONG	NO

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

//YES	BTCUSDT
func (api *FutureDataBasisApi) Pair(Pair string) *FutureDataBasisApi {
	api.req.Pair = GetPointer(Pair)
	return api
}

//YES	CURRENT_QUARTER, NEXT_QUARTER, PERPETUAL
func (api *FutureDataBasisApi) ContractType(ContractType string) *FutureDataBasisApi {
	api.req.ContractType = GetPointer(ContractType)
	return api
}

//YES	"5m","15m","30m","1h","2h","4h","6h","12h","1d"
func (api *FutureDataBasisApi) Period(Period string) *FutureDataBasisApi {
	api.req.Period = GetPointer(Period)
	return api
}

//NO	Default 30,Max 500
func (api *FutureDataBasisApi) Limit(Limit int64) *FutureDataBasisApi {
	api.req.Limit = GetPointer(Limit)
	return api
}

//NO
func (api *FutureDataBasisApi) StartTime(StartTime int64) *FutureDataBasisApi {
	api.req.StartTime = GetPointer(StartTime)
	return api
}

//NO
func (api *FutureDataBasisApi) EndTime(EndTime int64) *FutureDataBasisApi {
	api.req.EndTime = GetPointer(EndTime)
	return api
}

//listenKey相关
type FutureListenKeyPostReq struct{}

type FutureListenKeyPostApi struct {
	client *FutureRestClient
	req    *FutureListenKeyPostReq
}

type FutureListenKeyPutReq struct{}

type FutureListenKeyPutApi struct {
	client *FutureRestClient
	req    *FutureListenKeyPutReq
}

type FutureListenKeyDeleteReq struct{}

type FutureListenKeyDeleteApi struct {
	client *FutureRestClient
	req    *FutureListenKeyDeleteReq
}
