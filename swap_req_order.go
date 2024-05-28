package mybinanceapi

import "github.com/shopspring/decimal"

type SwapOpenOrdersReq struct {
	Symbol     *string `json:"symbol"` //No	交易对
	RecvWindow *int64  `json:"recvWindow"`
	Timestamp  *int64  `json:"timestamp"`
}
type SwapOpenOrdersApi struct {
	client *SwapRestClient
	req    *SwapOpenOrdersReq
}

func (api *SwapOpenOrdersApi) Symbol(Symbol string) *SwapOpenOrdersApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *SwapOpenOrdersApi) RecvWindow(RecvWindow int64) *SwapOpenOrdersApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *SwapOpenOrdersApi) Timestamp(Timestamp int64) *SwapOpenOrdersApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SwapAllOrdersReq struct {
	Symbol     *string `json:"symbol"`    //YES	交易对
	OrderId    *int64  `json:"orderId"`   //NO	只返回此orderID及之后的订单，缺省返回最近的订单
	StartTime  *int64  `json:"startTime"` //NO	起始时间
	EndTime    *int64  `json:"endTime"`   //NO	结束时间
	Limit      *int64  `json:"limit"`     //NO	返回的结果集数量 默认值:500 最大值:1000
	RecvWindow *int64  `json:"recvWindow"`
	Timestamp  *int64  `json:"timestamp"`
}
type SwapAllOrdersApi struct {
	client *SwapRestClient
	req    *SwapAllOrdersReq
}

func (api *SwapAllOrdersApi) Symbol(Symbol string) *SwapAllOrdersApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *SwapAllOrdersApi) OrderId(OrderId int64) *SwapAllOrdersApi {
	api.req.OrderId = GetPointer(OrderId)
	return api
}
func (api *SwapAllOrdersApi) StartTime(StartTime int64) *SwapAllOrdersApi {
	api.req.StartTime = GetPointer(StartTime)
	return api
}
func (api *SwapAllOrdersApi) EndTime(EndTime int64) *SwapAllOrdersApi {
	api.req.EndTime = GetPointer(EndTime)
	return api
}
func (api *SwapAllOrdersApi) Limit(Limit int64) *SwapAllOrdersApi {
	api.req.Limit = GetPointer(Limit)
	return api
}
func (api *SwapAllOrdersApi) RecvWindow(RecvWindow int64) *SwapAllOrdersApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *SwapAllOrdersApi) Timestamp(Timestamp int64) *SwapAllOrdersApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SwapOrderPostReq struct {
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
	RecvWindow       *int64           `json:"recvWindow"`       //No
	Timestamp        *int64           `json:"timestamp"`        //Yes
}
type SwapOrderPostApi struct {
	client *SwapRestClient
	req    *SwapOrderPostReq
}

func (api *SwapOrderPostApi) Symbol(Symbol string) *SwapOrderPostApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *SwapOrderPostApi) Side(Side string) *SwapOrderPostApi {
	api.req.Side = GetPointer(Side)
	return api
}
func (api *SwapOrderPostApi) PositionSide(PositionSide string) *SwapOrderPostApi {
	api.req.PositionSide = GetPointer(PositionSide)
	return api
}
func (api *SwapOrderPostApi) Type(Type string) *SwapOrderPostApi {
	api.req.Type = GetPointer(Type)
	return api
}
func (api *SwapOrderPostApi) ReduceOnly(ReduceOnly string) *SwapOrderPostApi {
	api.req.ReduceOnly = GetPointer(ReduceOnly)
	return api
}
func (api *SwapOrderPostApi) Quantity(Quantity decimal.Decimal) *SwapOrderPostApi {
	api.req.Quantity = GetPointer(Quantity)
	return api
}
func (api *SwapOrderPostApi) Price(Price decimal.Decimal) *SwapOrderPostApi {
	api.req.Price = GetPointer(Price)
	return api
}
func (api *SwapOrderPostApi) NewClientOrderId(NewClientOrderId string) *SwapOrderPostApi {
	api.req.NewClientOrderId = GetPointer(NewClientOrderId)
	return api
}
func (api *SwapOrderPostApi) StopPrice(StopPrice decimal.Decimal) *SwapOrderPostApi {
	api.req.StopPrice = GetPointer(StopPrice)
	return api
}
func (api *SwapOrderPostApi) ClosePosition(ClosePosition string) *SwapOrderPostApi {
	api.req.ClosePosition = GetPointer(ClosePosition)
	return api
}
func (api *SwapOrderPostApi) ActivationPrice(ActivationPrice decimal.Decimal) *SwapOrderPostApi {
	api.req.ActivationPrice = GetPointer(ActivationPrice)
	return api
}
func (api *SwapOrderPostApi) CallbackRate(CallbackRate decimal.Decimal) *SwapOrderPostApi {
	api.req.CallbackRate = GetPointer(CallbackRate)
	return api
}
func (api *SwapOrderPostApi) TimeInForce(TimeInForce string) *SwapOrderPostApi {
	api.req.TimeInForce = GetPointer(TimeInForce)
	return api
}
func (api *SwapOrderPostApi) WorkingType(WorkingType string) *SwapOrderPostApi {
	api.req.WorkingType = GetPointer(WorkingType)
	return api
}
func (api *SwapOrderPostApi) PriceProtect(PriceProtect string) *SwapOrderPostApi {
	api.req.PriceProtect = GetPointer(PriceProtect)
	return api
}
func (api *SwapOrderPostApi) NewOrderRespType(NewOrderRespType string) *SwapOrderPostApi {
	api.req.NewOrderRespType = GetPointer(NewOrderRespType)
	return api
}
func (api *SwapOrderPostApi) RecvWindow(RecvWindow int64) *SwapOrderPostApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *SwapOrderPostApi) Timestamp(Timestamp int64) *SwapOrderPostApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SwapOrderGetReq struct {
	Symbol            *string `json:"symbol"`            //YES 交易对
	OrderId           *int64  `json:"orderId"`           //NO 系统订单号
	OrigClientOrderId *string `json:"origClientOrderId"` //NO 用户自定义的订单号
	RecvWindow        *int64  `json:"recvWindow"`        //NO
	Timestamp         *int64  `json:"timestamp"`         //YES
}
type SwapOrderGetApi struct {
	client *SwapRestClient
	req    *SwapOrderGetReq
}

func (api *SwapOrderGetApi) Symbol(Symbol string) *SwapOrderGetApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *SwapOrderGetApi) OrderId(OrderId int64) *SwapOrderGetApi {
	api.req.OrderId = GetPointer(OrderId)
	return api
}
func (api *SwapOrderGetApi) OrigClientOrderId(OrigClientOrderId string) *SwapOrderGetApi {
	api.req.OrigClientOrderId = GetPointer(OrigClientOrderId)
	return api
}
func (api *SwapOrderGetApi) RecvWindow(RecvWindow int64) *SwapOrderGetApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *SwapOrderGetApi) Timestamp(Timestamp int64) *SwapOrderGetApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SwapOrderDeleteReq struct {
	Symbol            *string `json:"symbol"`            //YES 交易对
	OrderId           *int64  `json:"orderId"`           //NO 系统订单号
	OrigClientOrderId *string `json:"origClientOrderId"` //NO 用户自定义的订单号
	RecvWindow        *int64  `json:"recvWindow"`        //NO
	Timestamp         *int64  `json:"timestamp"`         //YES
}
type SwapOrderDeleteApi struct {
	client *SwapRestClient
	req    *SwapOrderDeleteReq
}

func (api *SwapOrderDeleteApi) Symbol(Symbol string) *SwapOrderDeleteApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *SwapOrderDeleteApi) OrderId(OrderId int64) *SwapOrderDeleteApi {
	api.req.OrderId = GetPointer(OrderId)
	return api
}
func (api *SwapOrderDeleteApi) OrigClientOrderId(OrigClientOrderId string) *SwapOrderDeleteApi {
	api.req.OrigClientOrderId = GetPointer(OrigClientOrderId)
	return api
}
func (api *SwapOrderDeleteApi) RecvWindow(RecvWindow int64) *SwapOrderDeleteApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *SwapOrderDeleteApi) Timestamp(Timestamp int64) *SwapOrderDeleteApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SwapUserTradesReq struct {
	Symbol     *string `json:"symbol"`     //NO	交易对
	Pair       *string `json:"pair"`       //NO	标的交易对
	OrderId    *int64  `json:"orderId"`    //NO	订单号
	StartTime  *int64  `json:"startTime"`  //NO	起始时间
	EndTime    *int64  `json:"endTime"`    //NO	结束时间
	FromId     *int64  `json:"fromId"`     //NO	返回该fromId及之后的成交,缺省返回最近的成交,仅支持配合symbol使用
	Limit      *int64  `json:"limit"`      //NO	返回的结果集数量 默认值:50 最大值:1000
	RecvWindow *int64  `json:"recvWindow"` //NO
	Timestamp  *int64  `json:"timestamp"`  //YES
}

type SwapUserTradesApi struct {
	client *SwapRestClient
	req    *SwapUserTradesReq
}

func (api *SwapUserTradesApi) Symbol(Symbol string) *SwapUserTradesApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *SwapUserTradesApi) Pair(Pair string) *SwapUserTradesApi {
	api.req.Pair = GetPointer(Pair)
	return api
}
func (api *SwapUserTradesApi) OrderId(OrderId int64) *SwapUserTradesApi {
	api.req.OrderId = GetPointer(OrderId)
	return api
}
func (api *SwapUserTradesApi) StartTime(StartTime int64) *SwapUserTradesApi {
	api.req.StartTime = GetPointer(StartTime)
	return api
}
func (api *SwapUserTradesApi) EndTime(EndTime int64) *SwapUserTradesApi {
	api.req.EndTime = GetPointer(EndTime)
	return api
}
func (api *SwapUserTradesApi) FromId(FromId int64) *SwapUserTradesApi {
	api.req.FromId = GetPointer(FromId)
	return api
}
func (api *SwapUserTradesApi) Limit(Limit int64) *SwapUserTradesApi {
	api.req.Limit = GetPointer(Limit)
	return api
}
func (api *SwapUserTradesApi) RecvWindow(RecvWindow int64) *SwapUserTradesApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *SwapUserTradesApi) Timestamp(Timestamp int64) *SwapUserTradesApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}
