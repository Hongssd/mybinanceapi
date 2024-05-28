package mybinanceapi

import (
	"github.com/shopspring/decimal"
	"strconv"
)

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

type FutureOrderPostReq struct {
	Symbol           *string          `json:"symbol"`                     //Yes	交易对
	Side             *string          `json:"side"`                       //Yes	买卖方向 SELL, BUY
	PositionSide     *string          `json:"positionSide,omitempty"`     //No	持仓方向，单向持仓模式下非必填，默认且仅可填BOTH;在双向持仓模式下必填,且仅可选择 LONG 或 SHORT
	Type             *string          `json:"type"`                       //Yes	订单类型 LIMIT, MARKET, STOP, TAKE_PROFIT, STOP_MARKET, TAKE_PROFIT_MARKET, TRAILING_STOP_MARKET
	ReduceOnly       *string          `json:"reduceOnly,omitempty"`       //No	true, false; 非双开模式下默认false；双开模式下不接受此参数； 使用closePosition不支持此参数。
	Quantity         *decimal.Decimal `json:"quantity,omitempty"`         //No	下单数量,使用closePosition不支持此参数。
	Price            *decimal.Decimal `json:"price,omitempty"`            //No	委托价格
	NewClientOrderId *string          `json:"newClientOrderId,omitempty"` //No	用户自定义的订单号，不可以重复出现在挂单中。如空缺系统会自动赋值。必须满足正则规则 ^[\.A-Z\:/a-z0-9_-]{1,36}$
	StopPrice        *decimal.Decimal `json:"stopPrice,omitempty"`        //No	触发价, 仅 STOP, STOP_MARKET, TAKE_PROFIT, TAKE_PROFIT_MARKET 需要此参数
	ClosePosition    *string          `json:"closePosition,omitempty"`    //No	true, false；触发后全部平仓，仅支持STOP_MARKET和TAKE_PROFIT_MARKET；不与quantity合用；自带只平仓效果，不与reduceOnly 合用
	ActivationPrice  *decimal.Decimal `json:"activationPrice,omitempty"`  //No	追踪止损激活价格，仅TRAILING_STOP_MARKET 需要此参数, 默认为下单当前市场价格(支持不同workingType)
	CallbackRate     *decimal.Decimal `json:"callbackRate,omitempty"`     //No	追踪止损回调比例，可取值范围[0.1, 5],其中 1代表1% ,仅TRAILING_STOP_MARKET 需要此参数
	TimeInForce      *string          `json:"timeInForce,omitempty"`      //No	有效方法
	WorkingType      *string          `json:"workingType,omitempty"`      //No	stopPrice 触发类型: MARK_PRICE(标记价格), CONTRACT_PRICE(合约最新价). 默认 CONTRACT_PRICE
	PriceProtect     *string          `json:"priceProtect,omitempty"`     //No	条件单触发保护："TRUE","FALSE", 默认"FALSE". 仅 STOP, STOP_MARKET, TAKE_PROFIT, TAKE_PROFIT_MARKET 需要此参数
	NewOrderRespType *string          `json:"newOrderRespType,omitempty"` //No	"ACK", "RESULT", 默认 "ACK"
	Recvwindow       *int64           `json:"recvWindow,omitempty"`       //No
	Timestamp        *int64           `json:"timestamp,omitempty"`        //Yes
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

type FutureOrderPutReq struct {
	Symbol            *string          `json:"symbol"`                      //YES	交易对
	OrderId           *string          `json:"orderId,omitempty"`           //NO	系统订单号
	OrigClientOrderId *string          `json:"origClientOrderId,omitempty"` //NO	用户自定义的订单号
	Side              *string          `json:"side"`                        //YES	买卖方向 SELL, BUY; side需要和原订单相同
	Quantity          *decimal.Decimal `json:"quantity"`                    //YES	下单数量,使用closePosition不支持此参数。
	Price             *decimal.Decimal `json:"price"`                       //YES	委托价格
	PriceMatch        *string          `json:"priceMatch,omitempty"`        //NO	OPPONENT/ OPPONENT_5/ OPPONENT_10/ OPPONENT_20/QUEUE/ QUEUE_5/ QUEUE_10/ QUEUE_20；不能与price同时传
	RecvWindow        *int64           `json:"recvWindow,omitempty"`        //NO
	Timestamp         *int64           `json:"timestamp,omitempty"`         //YES
}
type FutureOrderPutApi struct {
	client *FutureRestClient
	req    *FutureOrderPutReq
}

func (api *FutureOrderPutApi) Symbol(Symbol string) *FutureOrderPutApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *FutureOrderPutApi) OrderId(OrderId int64) *FutureOrderPutApi {
	orderIdStr := strconv.FormatInt(OrderId, BIT_BASE_10)
	api.req.OrderId = GetPointer(orderIdStr)
	return api
}
func (api *FutureOrderPutApi) OrigClientOrderId(OrigClientOrderId string) *FutureOrderPutApi {
	api.req.OrigClientOrderId = GetPointer(OrigClientOrderId)
	return api
}
func (api *FutureOrderPutApi) Side(Side string) *FutureOrderPutApi {
	api.req.Side = GetPointer(Side)
	return api
}
func (api *FutureOrderPutApi) Quantity(Quantity decimal.Decimal) *FutureOrderPutApi {
	api.req.Quantity = GetPointer(Quantity)
	return api
}
func (api *FutureOrderPutApi) Price(Price decimal.Decimal) *FutureOrderPutApi {
	api.req.Price = GetPointer(Price)
	return api
}
func (api *FutureOrderPutApi) PriceMatch(PriceMatch string) *FutureOrderPutApi {
	api.req.PriceMatch = GetPointer(PriceMatch)
	return api
}
func (api *FutureOrderPutApi) Recvwindow(Recvwindow int64) *FutureOrderPutApi {
	api.req.RecvWindow = GetPointer(Recvwindow)
	return api
}
func (api *FutureOrderPutApi) Timestamp(Timestamp int64) *FutureOrderPutApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type FutureOrderGetReq struct {
	Symbol            *string `json:"symbol"`                      //YES 交易对
	OrderId           *int64  `json:"orderId,omitempty"`           //NO 系统订单号
	OrigClientOrderId *string `json:"origClientOrderId,omitempty"` //NO 用户自定义的订单号
	RecvWindow        *int64  `json:"recvWindow,omitempty"`        //NO
	Timestamp         *int64  `json:"timestamp"`                   //YES
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

type FutureOrderDeleteReq struct {
	Symbol            *string `json:"symbol"`                      //YES 交易对
	OrderId           *int64  `json:"orderId,omitempty"`           //NO 系统订单号
	OrigClientOrderId *string `json:"origClientOrderId,omitempty"` //NO 用户自定义的订单号
	RecvWindow        *int64  `json:"recvWindow,omitempty"`        //NO
	Timestamp         *int64  `json:"timestamp"`                   //YES
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

type FutureBatchOrdersPostReq struct {
	BatchOrders *[]FutureOrderPostReq `json:"batchOrders"` //YES	批量下单 最多支持5个
	RecvWindow  *int64                `json:"recvWindow"`  //NO
	Timestamp   *int64                `json:"timestamp"`   //YES
}
type FutureBatchOrdersPostApi struct {
	client *FutureRestClient
	req    *FutureBatchOrdersPostReq
}

func (api *FutureBatchOrdersPostApi) AddOrders(orderApis ...*FutureOrderPostApi) *FutureBatchOrdersPostApi {
	if api.req.BatchOrders == nil {
		api.req.BatchOrders = &[]FutureOrderPostReq{}
	}
	for _, order := range orderApis {
		*api.req.BatchOrders = append(*api.req.BatchOrders, *order.req)
	}
	return api
}
func (api *FutureBatchOrdersPostApi) SetOrders(orderApi []*FutureOrderPostApi) *FutureBatchOrdersPostApi {
	api.req.BatchOrders = &[]FutureOrderPostReq{}
	for _, order := range orderApi {
		*api.req.BatchOrders = append(*api.req.BatchOrders, *order.req)
	}
	return api
}
func (api *FutureBatchOrdersPostApi) Recvwindow(Recvwindow int64) *FutureBatchOrdersPostApi {
	api.req.RecvWindow = GetPointer(Recvwindow)
	return api
}
func (api *FutureBatchOrdersPostApi) Timestamp(Timestamp int64) *FutureBatchOrdersPostApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type FutureBatchOrdersPutReq struct {
	BatchOrders *[]FutureOrderPutReq `json:"batchOrders"` //YES	批量下单 最多支持5个
	RecvWindow  *int64               `json:"recvWindow"`  //NO
	Timestamp   *int64               `json:"timestamp"`   //YES
}
type FutureBatchOrdersPutApi struct {
	client *FutureRestClient
	req    *FutureBatchOrdersPutReq
}

func (api *FutureBatchOrdersPutApi) AddOrders(orderApi ...*FutureOrderPutApi) *FutureBatchOrdersPutApi {
	if api.req.BatchOrders == nil {
		api.req.BatchOrders = &[]FutureOrderPutReq{}
	}
	for _, order := range orderApi {
		*api.req.BatchOrders = append(*api.req.BatchOrders, *order.req)
	}
	return api
}
func (api *FutureBatchOrdersPutApi) SetOrders(orderApi []*FutureOrderPutApi) *FutureBatchOrdersPutApi {
	api.req.BatchOrders = &[]FutureOrderPutReq{}
	for _, order := range orderApi {
		*api.req.BatchOrders = append(*api.req.BatchOrders, *order.req)
	}
	return api
}
func (api *FutureBatchOrdersPutApi) Recvwindow(Recvwindow int64) *FutureBatchOrdersPutApi {
	api.req.RecvWindow = GetPointer(Recvwindow)
	return api
}
func (api *FutureBatchOrdersPutApi) Timestamp(Timestamp int64) *FutureBatchOrdersPutApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

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
