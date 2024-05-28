package mybinanceapi

import (
	"github.com/shopspring/decimal"
)

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
	Symbol            *string `json:"symbol"`                      //YES
	OrderId           *int64  `json:"orderId,omitempty"`           //NO
	OrigClientOrderId *string `json:"origClientOrderId,omitempty"` //NO
	RecvWindow        *int64  `json:"recvWindow,omitempty"`        //NO 	赋值不能大于 60000
	Timestamp         *int64  `json:"timestamp"`                   //YES
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
	Symbol                  *string          `json:"symbol"`                            //YES
	Side                    *string          `json:"side"`                              //YES
	Type                    *string          `json:"type"`                              //YES
	TimeInForce             *string          `json:"timeInForce,omitempty"`             //NO 	详见枚举定义：有效方式
	Quantity                *decimal.Decimal `json:"quantity,omitempty"`                //NO
	QuoteOrderQty           *decimal.Decimal `json:"quoteOrderQty,omitempty"`           //NO
	Price                   *decimal.Decimal `json:"price,omitempty"`                   //NO
	NewClientOrderId        *string          `json:"newClientOrderId,omitempty"`        //NO 	客户自定义的唯一订单ID。 如果未发送，则自动生成。
	StopPrice               *decimal.Decimal `json:"stopPrice,omitempty"`               //NO 	仅 STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT 和 TAKE_PROFIT_LIMIT 需要此参数。
	TrailingDelta           *int64           `json:"trailingDelta,omitempty"`           //NO 	用于 STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT 和 TAKE_PROFIT_LIMIT 类型的订单。更多追踪止盈止损订单细节, 请参考 追踪止盈止损(Trailing Stop)订单常见问题。
	IcebergQty              *decimal.Decimal `json:"icebergQty,omitempty"`              //NO 	仅使用 LIMIT, STOP_LOSS_LIMIT, 和 TAKE_PROFIT_LIMIT 创建新的 iceberg 订单时需要此参数。
	NewOrderRespType        *string          `json:"newOrderRespType,omitempty"`        //NO 	设置响应JSON。ACK，RESULT 或 FULL；MARKET 和 LIMIT 订单类型默认为 FULL，所有其他订单默认为 ACK。
	SelfTradePreventionMode *string          `json:"selfTradePreventionMode,omitempty"` //NO 	允许的 ENUM 取决于交易对的配置。支持的值有 EXPIRE_TAKER，EXPIRE_MAKER，EXPIRE_BOTH，NONE。
	StrategyId              *int             `json:"strategyId,omitempty"`              //NO
	StrategyType            *int             `json:"strategyType,omitempty"`            //NO 	不能低于 1000000
	RecvWindow              *int64           `json:"recvWindow,omitempty"`              //NO 	赋值不能大于 60000
	Timestamp               *int64           `json:"timestamp"`                         //YES
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
	Symbol             *string `json:"symbol"`                       //YES
	OrderId            *int64  `json:"orderId,omitempty"`            //NO
	OrigClientOrderId  *string `json:"origClientOrderId,omitempty"`  //NO
	NewClientOrderId   *string `json:"newClientOrderId,omitempty"`   //NO	用户自定义的本次撤销操作的ID(注意不是被撤销的订单的自定义ID)。如无指定会自动赋值。
	CancelRestrictions *string `json:"cancelRestrictions,omitempty"` //NO	支持的值: ONLY_NEW - 如果订单状态为 NEW，订单取消将成功。 ONLY_PARTIALLY_FILLED - 如果订单状态为 PARTIALLY_FILLED，订单取消将成功。
	RecvWindow         *int64  `json:"recvWindow,omitempty"`         //NO	赋值不得大于 60000
	Timestamp          *int64  `json:"timestamp"`                    //YES
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

// symbol	STRING	YES
// side	ENUM	YES
// type	ENUM	YES
// cancelReplaceMode	ENUM	YES	指定类型：STOP_ON_FAILURE - 如果撤消订单失败将不会继续重新下单。
// ALLOW_FAILURE - 不管撤消订单是否成功都会继续重新下单。
// timeInForce	ENUM	NO
// quantity	DECIMAL	NO
// quoteOrderQty	DECIMAL	NO
// price	DECIMAL	NO
// cancelNewClientOrderId	STRING	NO	用户自定义的id，如空缺系统会自动赋值
// cancelOrigClientOrderId	STRING	NO	必须提供cancelOrigClientOrderId 或者 cancelOrderId。 如果两个参数都提供, cancelOrderId 会占优先。
// cancelOrderId	LONG	NO	必须提供cancelOrigClientOrderId 或者 cancelOrderId。 如果两个参数都提供, cancelOrderId 会占优先。
// newClientOrderId	STRING	NO	用于辨识新订单。
// strategyId	INT	NO
// strategyType	INT	NO	不能低于 1000000
// stopPrice	DECIMAL	NO
// trailingDelta	LONG	NO
// icebergQty	DECIMAL	NO
// newOrderRespType	ENUM	NO	指定响应类型:
// 指定响应类型 ACK, RESULT, 或者 FULL; MARKET 与 LIMIT 订单默认为FULL, 其他默认为ACK.
// selfTradePreventionMode	ENUM	NO	允许的 ENUM 取决于交易对的配置。支持的值有 EXPIRE_TAKER，EXPIRE_MAKER，EXPIRE_BOTH，NONE。
// cancelRestrictions	ENUM	NO	支持的值:
// ONLY_NEW - 如果订单状态为 NEW，撤销将成功。
// ONLY_PARTIALLY_FILLED - 如果订单状态为 PARTIALLY_FILLED，撤销将成功。
// recvWindow	LONG	NO	不能大于 60000
// timestamp	LONG	YES
type SpotOrderCancelReplaceReq struct {
	Symbol                  *string          `json:"symbol"`                            //YES
	Side                    *string          `json:"side"`                              //YES
	Type                    *string          `json:"type"`                              //YES
	CancelReplaceMode       *string          `json:"cancelReplaceMode"`                 //YES	指定类型：STOP_ON_FAILURE - 如果撤消订单失败将不会继续重新下单。 ALLOW_FAILURE - 不管撤消订单是否成功都会继续重新下单。
	TimeInForce             *string          `json:"timeInForce,omitempty"`             //NO
	Quantity                *decimal.Decimal `json:"quantity,omitempty"`                //NO
	QuoteOrderQty           *decimal.Decimal `json:"quoteOrderQty,omitempty"`           //NO
	Price                   *decimal.Decimal `json:"price,omitempty"`                   //NO
	CancelNewClientOrderId  *string          `json:"cancelNewClientOrderId,omitempty"`  //NO	用户自定义的id，如空缺系统会自动赋值
	CancelOrigClientOrderId *string          `json:"cancelOrigClientOrderId,omitempty"` //NO	必须提供cancelOrigClientOrderId 或者 cancelOrderId。 如果两个参数都提供, cancelOrderId 会占优先。
	CancelOrderId           *int64           `json:"cancelOrderId,omitempty"`           //NO	必须提供cancelOrigClientOrderId 或者 cancelOrderId。 如果两个参数都提供, cancelOrderId 会占优先。
	NewClientOrderId        *string          `json:"newClientOrderId,omitempty"`        //NO	用于辨识新订单。
	StrategyId              *int             `json:"strategyId,omitempty"`              //NO
	StrategyType            *int             `json:"strategyType,omitempty"`            //NO	不能低于 1000000
	StopPrice               *decimal.Decimal `json:"stopPrice,omitempty"`               //NO
	TrailingDelta           *int64           `json:"trailingDelta,omitempty"`           //NO
	IcebergQty              *decimal.Decimal `json:"icebergQty,omitempty"`              //NO
	NewOrderRespType        *string          `json:"newOrderRespType,omitempty"`        //NO	指定响应类型: 指定响应类型 ACK, RESULT, 或者 FULL; MARKET 与 LIMIT 订单默认为FULL, 其他默认为ACK.
	SelfTradePreventionMode *string          `json:"selfTradePreventionMode,omitempty"` //NO	允许的 ENUM 取决于交易对的配置。支持的值有 EXPIRE_TAKER，EXPIRE_MAKER，EXPIRE_BOTH，NONE。
	CancelRestrictions      *string          `json:"cancelRestrictions,omitempty"`      //NO	支持的值: ONLY_NEW - 如果订单状态为 NEW，撤销将成功。 ONLY_PARTIALLY_FILLED - 如果订单状态为 PARTIALLY_FILLED，撤销将成功。
	RecvWindow              *int64           `json:"recvWindow,omitempty"`              //NO	不能大于 60000
	Timestamp               *int64           `json:"timestamp"`                         //YES
}
type SpotOrderCancelReplaceApi struct {
	client *SpotRestClient
	req    *SpotOrderCancelReplaceReq
}

func (api *SpotOrderCancelReplaceApi) Symbol(Symbol string) *SpotOrderCancelReplaceApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *SpotOrderCancelReplaceApi) Side(Side string) *SpotOrderCancelReplaceApi {
	api.req.Side = GetPointer(Side)
	return api
}
func (api *SpotOrderCancelReplaceApi) Type(Type string) *SpotOrderCancelReplaceApi {
	api.req.Type = GetPointer(Type)
	return api
}
func (api *SpotOrderCancelReplaceApi) CancelReplaceMode(CancelReplaceMode string) *SpotOrderCancelReplaceApi {
	api.req.CancelReplaceMode = GetPointer(CancelReplaceMode)
	return api
}
func (api *SpotOrderCancelReplaceApi) TimeInForce(TimeInForce string) *SpotOrderCancelReplaceApi {
	api.req.TimeInForce = GetPointer(TimeInForce)
	return api
}
func (api *SpotOrderCancelReplaceApi) Quantity(Quantity decimal.Decimal) *SpotOrderCancelReplaceApi {
	api.req.Quantity = GetPointer(Quantity)
	return api
}
func (api *SpotOrderCancelReplaceApi) QuoteOrderQty(QuoteOrderQty decimal.Decimal) *SpotOrderCancelReplaceApi {
	api.req.QuoteOrderQty = GetPointer(QuoteOrderQty)
	return api
}
func (api *SpotOrderCancelReplaceApi) Price(Price decimal.Decimal) *SpotOrderCancelReplaceApi {
	api.req.Price = GetPointer(Price)
	return api
}
func (api *SpotOrderCancelReplaceApi) CancelNewClientOrderId(CancelNewClientOrderId string) *SpotOrderCancelReplaceApi {
	api.req.CancelNewClientOrderId = GetPointer(CancelNewClientOrderId)
	return api
}
func (api *SpotOrderCancelReplaceApi) CancelOrigClientOrderId(CancelOrigClientOrderId string) *SpotOrderCancelReplaceApi {
	api.req.CancelOrigClientOrderId = GetPointer(CancelOrigClientOrderId)
	return api
}
func (api *SpotOrderCancelReplaceApi) CancelOrderId(CancelOrderId int64) *SpotOrderCancelReplaceApi {
	api.req.CancelOrderId = GetPointer(CancelOrderId)
	return api
}
func (api *SpotOrderCancelReplaceApi) NewClientOrderId(NewClientOrderId string) *SpotOrderCancelReplaceApi {
	api.req.NewClientOrderId = GetPointer(NewClientOrderId)
	return api
}
func (api *SpotOrderCancelReplaceApi) StrategyId(StrategyId int) *SpotOrderCancelReplaceApi {
	api.req.StrategyId = GetPointer(StrategyId)
	return api
}
func (api *SpotOrderCancelReplaceApi) StrategyType(StrategyType int) *SpotOrderCancelReplaceApi {
	api.req.StrategyType = GetPointer(StrategyType)
	return api
}
func (api *SpotOrderCancelReplaceApi) StopPrice(StopPrice decimal.Decimal) *SpotOrderCancelReplaceApi {
	api.req.StopPrice = GetPointer(StopPrice)
	return api
}
func (api *SpotOrderCancelReplaceApi) TrailingDelta(TrailingDelta int64) *SpotOrderCancelReplaceApi {
	api.req.TrailingDelta = GetPointer(TrailingDelta)
	return api
}
func (api *SpotOrderCancelReplaceApi) IcebergQty(IcebergQty decimal.Decimal) *SpotOrderCancelReplaceApi {
	api.req.IcebergQty = GetPointer(IcebergQty)
	return api
}
func (api *SpotOrderCancelReplaceApi) NewOrderRespType(NewOrderRespType string) *SpotOrderCancelReplaceApi {
	api.req.NewOrderRespType = GetPointer(NewOrderRespType)
	return api
}
func (api *SpotOrderCancelReplaceApi) SelfTradePreventionMode(SelfTradePreventionMode string) *SpotOrderCancelReplaceApi {
	api.req.SelfTradePreventionMode = GetPointer(SelfTradePreventionMode)
	return api
}
func (api *SpotOrderCancelReplaceApi) CancelRestrictions(CancelRestrictions string) *SpotOrderCancelReplaceApi {
	api.req.CancelRestrictions = GetPointer(CancelRestrictions)
	return api
}
func (api *SpotOrderCancelReplaceApi) RecvWindow(RecvWindow int64) *SpotOrderCancelReplaceApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *SpotOrderCancelReplaceApi) Timestamp(Timestamp int64) *SpotOrderCancelReplaceApi {
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
