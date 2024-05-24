package mybinanceapi

import "github.com/shopspring/decimal"

type SpotOpenOrdersReq struct {
	Symbol     *string `json:"symbol"`     //NO
	RecvWindow *int64  `json:"recvWindow"` //NO 	赋值不能大于 60000
	Timestamp  *int64  `json:"timestamp"`  //YES
}
type SpotOpenOrdersApi struct {
	client *SpotRestClient
	req    *SpotOpenOrdersReq
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
	client *SpotRestClient
	req    *SpotAllOrdersReq
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

type SpotOrderGetReq struct {
	Symbol            *string `json:"symbol"`            //YES
	OrderId           *int64  `json:"orderId"`           //NO
	OrigClientOrderId *string `json:"origClientOrderId"` //NO
	RecvWindow        *int64  `json:"recvWindow"`        //NO 	赋值不能大于 60000
	Timestamp         *int64  `json:"timestamp"`         //YES
}
type SpotOrderGetApi struct {
	client *SpotRestClient
	req    *SpotOrderGetReq
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
	client *SpotRestClient
	req    *SpotOrderPostReq
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
	client *SpotRestClient
	req    *SpotOrderDeleteReq
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

type SpotMyTradesReq struct {
	Symbol     *string `json:"symbol"`     //YES
	OrderId    *int64  `json:"orderId"`    //NO
	StartTime  *int64  `json:"startTime"`  //NO
	EndTime    *int64  `json:"endTime"`    //NO
	FromId     *int64  `json:"fromId"`     //NO
	Limit      *int    `json:"limit"`      //NO	默认 500; 最大 1000.
	RecvWindow *int64  `json:"recvWindow"` //NO	赋值不能超过 60000
	Timestamp  *int64  `json:"timestamp"`  //YES
}
type SpotMyTradesApi struct {
	client *SpotRestClient
	req    *SpotMyTradesReq
}

func (api *SpotMyTradesApi) Symbol(Symbol string) *SpotMyTradesApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *SpotMyTradesApi) OrderId(OrderId int64) *SpotMyTradesApi {
	api.req.OrderId = GetPointer(OrderId)
	return api
}
func (api *SpotMyTradesApi) StartTime(StartTime int64) *SpotMyTradesApi {
	api.req.StartTime = GetPointer(StartTime)
	return api
}
func (api *SpotMyTradesApi) EndTime(EndTime int64) *SpotMyTradesApi {
	api.req.EndTime = GetPointer(EndTime)
	return api
}
func (api *SpotMyTradesApi) FromId(FromId int64) *SpotMyTradesApi {
	api.req.FromId = GetPointer(FromId)
	return api
}
func (api *SpotMyTradesApi) Limit(Limit int) *SpotMyTradesApi {
	api.req.Limit = GetPointer(Limit)
	return api
}
func (api *SpotMyTradesApi) RecvWindow(RecvWindow int64) *SpotMyTradesApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *SpotMyTradesApi) Timestamp(Timestamp int64) *SpotMyTradesApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}
