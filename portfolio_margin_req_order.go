package mybinanceapi

import "github.com/shopspring/decimal"

type PortfolioMarginUmOrderPostReq struct {
	Symbol                  *string          `json:"symbol"`                  // YES 交易对
	Side                    *string          `json:"side"`                    // YES 方向
	PositionSide            *string          `json:"positionSide"`            // NO 持仓方向，单向持仓模式下非必填，默认且仅可填BOTH;在双向持仓模式下必填,且仅可选择 LONG 或 SHORT
	Type                    *string          `json:"type"`                    // YES LIMIT, MARKET
	TimeInForce             *string          `json:"timeInForce"`             // NO 有效方法
	Quantity                *decimal.Decimal `json:"quantity"`                // NO 下单数量
	ReduceOnly              *string          `json:"reduceOnly"`              // NO true或false; 非双开模式下默认false；双开模式下不接受此参数
	Price                   *decimal.Decimal `json:"price"`                   // NO 委托价格
	NewClientOrderId        *string          `json:"newClientOrderId"`        // NO 用户自定义的订单号，不可以重复出现在挂单中。如空缺系统会自动赋值。必须满足正则规则: ^[\.A-Z\:/a-z0-9_-]{1,32}$
	NewOrderRespType        *string          `json:"newOrderRespType"`        // NO ACK， RESULT，默认 ACK
	PriceMatch              *string          `json:"priceMatch"`              // NO OPPONENT/ OPPONENT_5/ OPPONENT_10/ OPPONENT_20/QUEUE/ QUEUE_5/ QUEUE_10/ QUEUE_20；不能与price同时传
	SelfTradePreventionMode *string          `json:"selfTradePreventionMode"` // NO NONE / EXPIRE_TAKER/ EXPIRE_MAKER/ EXPIRE_BOTH； 默认NONE
	GoodTillDate            *int64           `json:"goodTillDate"`            // NO TIF为GTD时订单的自动取消时间， 当timeInforce为GTD时必传；传入的时间戳仅保留秒级精度，毫秒级部分会被自动忽略，时间戳需大于当前时间+600s且小于253402300799000
	RecvWindow              *int64           `json:"recvWindow"`              // NO
	Timestamp               *int64           `json:"timestamp"`               // YES
}

// YES 交易对
func (api *PortfolioMarginUmOrderPostApi) Symbol(symbol string) *PortfolioMarginUmOrderPostApi {
	api.req.Symbol = GetPointer(symbol)
	return api
}

// YES 方向
func (api *PortfolioMarginUmOrderPostApi) Side(side string) *PortfolioMarginUmOrderPostApi {
	api.req.Side = GetPointer(side)
	return api
}

// NO 持仓方向，单向持仓模式下非必填，默认且仅可填BOTH;在双向持仓模式下必填,且仅可选择 LONG 或 SHORT
func (api *PortfolioMarginUmOrderPostApi) PositionSide(positionSide string) *PortfolioMarginUmOrderPostApi {
	api.req.PositionSide = GetPointer(positionSide)
	return api
}

// YES LIMIT, MARKET
func (api *PortfolioMarginUmOrderPostApi) Type(type_ string) *PortfolioMarginUmOrderPostApi {
	api.req.Type = GetPointer(type_)
	return api
}

// NO 有效方法
func (api *PortfolioMarginUmOrderPostApi) TimeInForce(timeInForce string) *PortfolioMarginUmOrderPostApi {
	api.req.TimeInForce = GetPointer(timeInForce)
	return api
}

// NO 下单数量
func (api *PortfolioMarginUmOrderPostApi) Quantity(quantity decimal.Decimal) *PortfolioMarginUmOrderPostApi {
	api.req.Quantity = &quantity
	return api
}

// NO true或false; 非双开模式下默认false；双开模式下不接受此参数
func (api *PortfolioMarginUmOrderPostApi) ReduceOnly(reduceOnly string) *PortfolioMarginUmOrderPostApi {
	api.req.ReduceOnly = GetPointer(reduceOnly)
	return api
}

// NO 委托价格
func (api *PortfolioMarginUmOrderPostApi) Price(price decimal.Decimal) *PortfolioMarginUmOrderPostApi {
	api.req.Price = &price
	return api
}

// NO 用户自定义的订单号，不可以重复出现在挂单中。如空缺系统会自动赋值。必须满足正则规则: ^[\.A-Z\:/a-z0-9_-]{1,32}$
func (api *PortfolioMarginUmOrderPostApi) NewClientOrderId(newClientOrderId string) *PortfolioMarginUmOrderPostApi {
	api.req.NewClientOrderId = GetPointer(newClientOrderId)
	return api
}

// NO ACK， RESULT，默认 ACK
func (api *PortfolioMarginUmOrderPostApi) NewOrderRespType(newOrderRespType string) *PortfolioMarginUmOrderPostApi {
	api.req.NewOrderRespType = GetPointer(newOrderRespType)
	return api
}

// NO OPPONENT/ OPPONENT_5/ OPPONENT_10/ OPPONENT_20/QUEUE/ QUEUE_5/ QUEUE_10/ QUEUE_20；不能与price同时传
func (api *PortfolioMarginUmOrderPostApi) PriceMatch(priceMatch string) *PortfolioMarginUmOrderPostApi {
	api.req.PriceMatch = GetPointer(priceMatch)
	return api
}

// NO NONE / EXPIRE_TAKER/ EXPIRE_MAKER/ EXPIRE_BOTH； 默认NONE
func (api *PortfolioMarginUmOrderPostApi) SelfTradePreventionMode(selfTradePreventionMode string) *PortfolioMarginUmOrderPostApi {
	api.req.SelfTradePreventionMode = GetPointer(selfTradePreventionMode)
	return api
}

// NO TIF为GTD时订单的自动取消时间， 当timeInforce为GTD时必传；传入的时间戳仅保留秒级精度，毫秒级部分会被自动忽略，时间戳需大于当前时间+600s且小于253402300799000
func (api *PortfolioMarginUmOrderPostApi) GoodTillDate(goodTillDate int64) *PortfolioMarginUmOrderPostApi {
	api.req.GoodTillDate = &goodTillDate
	return api
}

// NO
func (api *PortfolioMarginUmOrderPostApi) RecvWindow(recvWindow int64) *PortfolioMarginUmOrderPostApi {
	api.req.RecvWindow = &recvWindow
	return api
}

// YES
func (api *PortfolioMarginUmOrderPostApi) Timestamp(timestamp int64) *PortfolioMarginUmOrderPostApi {
	api.req.Timestamp = GetPointer(timestamp)
	return api
}

type PortfolioMarginUmOrderPostApi struct {
	client *PortfolioMarginRestClient
	req    *PortfolioMarginUmOrderPostReq
}

type PortfolioMarginUmOrderDeleteReq struct {
	Symbol            *string `json:"symbol"`            // YES 交易对
	OrderId           *int64  `json:"orderId"`           // NO
	OrigClientOrderId *string `json:"origClientOrderId"` // NO
	RecvWindow        *int64  `json:"recvWindow"`        // NO
	Timestamp         *int64  `json:"timestamp"`         // YES
}

// YES 交易对
func (api *PortfolioMarginUmOrderDeleteApi) Symbol(symbol string) *PortfolioMarginUmOrderDeleteApi {
	api.req.Symbol = GetPointer(symbol)
	return api
}

// NO
func (api *PortfolioMarginUmOrderDeleteApi) OrderId(orderId int64) *PortfolioMarginUmOrderDeleteApi {
	api.req.OrderId = GetPointer(orderId)
	return api
}

// NO
func (api *PortfolioMarginUmOrderDeleteApi) OrigClientOrderId(origClientOrderId string) *PortfolioMarginUmOrderDeleteApi {
	api.req.OrigClientOrderId = GetPointer(origClientOrderId)
	return api
}

// NO
func (api *PortfolioMarginUmOrderDeleteApi) RecvWindow(recvWindow int64) *PortfolioMarginUmOrderDeleteApi {
	api.req.RecvWindow = GetPointer(recvWindow)
	return api
}

// YES
func (api *PortfolioMarginUmOrderDeleteApi) Timestamp(timestamp int64) *PortfolioMarginUmOrderDeleteApi {
	api.req.Timestamp = GetPointer(timestamp)
	return api
}

type PortfolioMarginUmOrderDeleteApi struct {
	client *PortfolioMarginRestClient
	req    *PortfolioMarginUmOrderDeleteReq
}

type PortfolioMarginUmAllOpenOrdersDeleteReq struct {
	Symbol     *string `json:"symbol"`     // YES 交易对
	RecvWindow *int64  `json:"recvWindow"` // NO
	Timestamp  *int64  `json:"timestamp"`  // YES
}

// YES 交易对
func (api *PortfolioMarginUmAllOpenOrdersDeleteApi) Symbol(symbol string) *PortfolioMarginUmAllOpenOrdersDeleteApi {
	api.req.Symbol = GetPointer(symbol)
	return api
}

// NO
func (api *PortfolioMarginUmAllOpenOrdersDeleteApi) RecvWindow(recvWindow int64) *PortfolioMarginUmAllOpenOrdersDeleteApi {
	api.req.RecvWindow = GetPointer(recvWindow)
	return api
}

// YES
func (api *PortfolioMarginUmAllOpenOrdersDeleteApi) Timestamp(timestamp int64) *PortfolioMarginUmAllOpenOrdersDeleteApi {
	api.req.Timestamp = GetPointer(timestamp)
	return api
}

type PortfolioMarginUmAllOpenOrdersDeleteApi struct {
	client *PortfolioMarginRestClient
	req    *PortfolioMarginUmAllOpenOrdersDeleteReq
}

type PortfolioMarginUmOrderPutReq struct {
	OrderId           *int64           `json:"orderId"`           // NO 系统订单号
	OrigClientOrderId *string          `json:"origClientOrderId"` // NO 用户自定义的订单号
	Symbol            *string          `json:"symbol"`            // YES 交易对
	Side              *string          `json:"side"`              // YES 买卖方向 SELL, BUY; side需要和原订单相同
	Quantity          *decimal.Decimal `json:"quantity"`          // YES 下单数量
	Price             *decimal.Decimal `json:"price"`             // YES 委托价格
	PriceMatch        *string          `json:"priceMatch"`        // NO OPPONENT/ OPPONENT_5/ OPPONENT_10/ OPPONENT_20/QUEUE/ QUEUE_5/ QUEUE_10/ QUEUE_20；不能与price同时传
	RecvWindow        *int64           `json:"recvWindow"`        // NO
	Timestamp         *int64           `json:"timestamp"`         // YES
}

// NO 系统订单号
func (api *PortfolioMarginUmOrderPutApi) OrderId(orderId int64) *PortfolioMarginUmOrderPutApi {
	api.req.OrderId = GetPointer(orderId)
	return api
}

// NO 用户自定义的订单号
func (api *PortfolioMarginUmOrderPutApi) OrigClientOrderId(origClientOrderId string) *PortfolioMarginUmOrderPutApi {
	api.req.OrigClientOrderId = GetPointer(origClientOrderId)
	return api
}

// YES 交易对
func (api *PortfolioMarginUmOrderPutApi) Symbol(symbol string) *PortfolioMarginUmOrderPutApi {
	api.req.Symbol = GetPointer(symbol)
	return api
}

// YES 买卖方向 SELL, BUY; side需要和原订单相同
func (api *PortfolioMarginUmOrderPutApi) Side(side string) *PortfolioMarginUmOrderPutApi {
	api.req.Side = GetPointer(side)
	return api
}

// YES 下单数量
func (api *PortfolioMarginUmOrderPutApi) Quantity(quantity decimal.Decimal) *PortfolioMarginUmOrderPutApi {
	api.req.Quantity = &quantity
	return api
}

// YES 委托价格
func (api *PortfolioMarginUmOrderPutApi) Price(price decimal.Decimal) *PortfolioMarginUmOrderPutApi {
	api.req.Price = &price
	return api
}

// NO OPPONENT/ OPPONENT_5/ OPPONENT_10/ OPPONENT_20/QUEUE/ QUEUE_5/ QUEUE_10/ QUEUE_20；不能与price同时传
func (api *PortfolioMarginUmOrderPutApi) PriceMatch(priceMatch string) *PortfolioMarginUmOrderPutApi {
	api.req.PriceMatch = GetPointer(priceMatch)
	return api
}

// NO
func (api *PortfolioMarginUmOrderPutApi) RecvWindow(recvWindow int64) *PortfolioMarginUmOrderPutApi {
	api.req.RecvWindow = GetPointer(recvWindow)
	return api
}

// YES
func (api *PortfolioMarginUmOrderPutApi) Timestamp(timestamp int64) *PortfolioMarginUmOrderPutApi {
	api.req.Timestamp = GetPointer(timestamp)
	return api
}

type PortfolioMarginUmOrderPutApi struct {
	client *PortfolioMarginRestClient
	req    *PortfolioMarginUmOrderPutReq
}

type PortfolioMarginUmOrderGetReq struct {
	Symbol            *string `json:"symbol"`            // YES 交易对
	OrderId           *int64  `json:"orderId"`           // NO
	OrigClientOrderId *string `json:"origClientOrderId"` // NO
	RecvWindow        *int64  `json:"recvWindow"`        // NO
	Timestamp         *int64  `json:"timestamp"`         // YES
}

// YES 交易对
func (api *PortfolioMarginUmOrderGetApi) Symbol(symbol string) *PortfolioMarginUmOrderGetApi {
	api.req.Symbol = GetPointer(symbol)
	return api
}

// NO
func (api *PortfolioMarginUmOrderGetApi) OrderId(orderId int64) *PortfolioMarginUmOrderGetApi {
	api.req.OrderId = GetPointer(orderId)
	return api
}

// NO
func (api *PortfolioMarginUmOrderGetApi) OrigClientOrderId(origClientOrderId string) *PortfolioMarginUmOrderGetApi {
	api.req.OrigClientOrderId = GetPointer(origClientOrderId)
	return api
}

// NO
func (api *PortfolioMarginUmOrderGetApi) RecvWindow(recvWindow int64) *PortfolioMarginUmOrderGetApi {
	api.req.RecvWindow = GetPointer(recvWindow)
	return api
}

// YES
func (api *PortfolioMarginUmOrderGetApi) Timestamp(timestamp int64) *PortfolioMarginUmOrderGetApi {
	api.req.Timestamp = GetPointer(timestamp)
	return api
}

type PortfolioMarginUmOrderGetApi struct {
	client *PortfolioMarginRestClient
	req    *PortfolioMarginUmOrderGetReq
}

type PortfolioMarginUmOpenOrderGetReq struct {
	Symbol            *string `json:"symbol"`            // YES 交易对
	OrderId           *int64  `json:"orderId"`           // NO
	OrigClientOrderId *string `json:"origClientOrderId"` // NO
	RecvWindow        *int64  `json:"recvWindow"`        // NO
	Timestamp         *int64  `json:"timestamp"`         // YES
}

// YES 交易对
func (api *PortfolioMarginUmOpenOrderGetApi) Symbol(symbol string) *PortfolioMarginUmOpenOrderGetApi {
	api.req.Symbol = GetPointer(symbol)
	return api
}

// NO
func (api *PortfolioMarginUmOpenOrderGetApi) OrderId(orderId int64) *PortfolioMarginUmOpenOrderGetApi {
	api.req.OrderId = GetPointer(orderId)
	return api
}

// NO
func (api *PortfolioMarginUmOpenOrderGetApi) OrigClientOrderId(origClientOrderId string) *PortfolioMarginUmOpenOrderGetApi {
	api.req.OrigClientOrderId = GetPointer(origClientOrderId)
	return api
}

// NO
func (api *PortfolioMarginUmOpenOrderGetApi) RecvWindow(recvWindow int64) *PortfolioMarginUmOpenOrderGetApi {
	api.req.RecvWindow = GetPointer(recvWindow)
	return api
}

// YES
func (api *PortfolioMarginUmOpenOrderGetApi) Timestamp(timestamp int64) *PortfolioMarginUmOpenOrderGetApi {
	api.req.Timestamp = GetPointer(timestamp)
	return api
}

type PortfolioMarginUmOpenOrderGetApi struct {
	client *PortfolioMarginRestClient
	req    *PortfolioMarginUmOpenOrderGetReq
}

type PortfolioMarginUmOpenOrdersGetReq struct {
	Symbol     *string `json:"symbol"`     // YES 交易对
	RecvWindow *int64  `json:"recvWindow"` // NO
	Timestamp  *int64  `json:"timestamp"`  // YES
}

// YES 交易对
func (api *PortfolioMarginUmOpenOrdersGetApi) Symbol(symbol string) *PortfolioMarginUmOpenOrdersGetApi {
	api.req.Symbol = GetPointer(symbol)
	return api
}

// NO
func (api *PortfolioMarginUmOpenOrdersGetApi) RecvWindow(recvWindow int64) *PortfolioMarginUmOpenOrdersGetApi {
	api.req.RecvWindow = GetPointer(recvWindow)
	return api
}

// YES
func (api *PortfolioMarginUmOpenOrdersGetApi) Timestamp(timestamp int64) *PortfolioMarginUmOpenOrdersGetApi {
	api.req.Timestamp = GetPointer(timestamp)
	return api
}

type PortfolioMarginUmOpenOrdersGetApi struct {
	client *PortfolioMarginRestClient
	req    *PortfolioMarginUmOpenOrdersGetReq
}

type PortfolioMarginUmConditionalOrderPostApi struct {
	client *PortfolioMarginRestClient
	req    *PortfolioMarginUmConditionalOrderReq
}

type PortfolioMarginUmConditionalOrderReq struct {
	Symbol                  *string          `json:"symbol"`                  // YES 交易对
	Side                    *string          `json:"side"`                    // YES 方向
	PositionSide            *string          `json:"positionSide"`            // NO 持仓方向，单向持仓模式下非必填，默认且仅可填BOTH;在双向持仓模式下必填,且仅可选择 LONG 或 SHORT
	StrategyType            *string          `json:"strategyType"`            // YES 条件单类型"STOP", "STOP_MARKET", "TAKE_PROFIT", "TAKE_PROFIT_MARKET"或"TRAILING_STOP_MARKET"
	TimeInForce             *string          `json:"timeInForce"`             // NO
	Quantity                *decimal.Decimal `json:"quantity"`                // NO
	ReduceOnly              *string          `json:"reduceOnly"`              // NO true或false; 非双开模式下默认false；双开模式下不接受此参数
	Price                   *decimal.Decimal `json:"price"`                   // NO
	WorkingType             *string          `json:"workingType"`             // NO stopPrice 触发类型: MARK_PRICE(标记价格), CONTRACT_PRICE(合约最新价). 默认 CONTRACT_PRICE
	PriceProtect            *string          `json:"priceProtect"`            // NO 条件单触发保护："TRUE","FALSE", 默认"FALSE". 仅 STOP, STOP_MARKET, TAKE_PROFIT, TAKE_PROFIT_MARKET 需要此参数
	NewClientStrategyId     *string          `json:"newClientStrategyId"`     // NO 用户自定义的订单号，不可以重复出现在挂单中。如空缺系统会自动赋值。必须满足正则规则: ^[\.A-Z\:/a-z0-9_-]{1,32}$
	StopPrice               *decimal.Decimal `json:"stopPrice"`               // NO Used with STOP/STOP_MARKET or TAKE_PROFIT/TAKE_PROFIT_MARKET orders.
	ActivationPrice         *decimal.Decimal `json:"activationPrice"`         // NO TRAILING_STOP_MARKET 单使用，默认标记价格
	CallbackRate            *decimal.Decimal `json:"callbackRate"`            // NO TRAILING_STOP_MARKET 单使用, 最小0.1, 最大5，1代表1%
	PriceMatch              *string          `json:"priceMatch"`              // NO OPPONENT/ OPPONENT_5/ OPPONENT_10/ OPPONENT_20/QUEUE/ QUEUE_5/ QUEUE_10/ QUEUE_20；不能与price同时传
	SelfTradePreventionMode *string          `json:"selfTradePreventionMode"` // NO NONE / EXPIRE_TAKER/ EXPIRE_MAKER/ EXPIRE_BOTH； 默认NONE
	GoodTillDate            *int64           `json:"goodTillDate"`            // NO TIF为GTD时订单的自动取消时间， 当timeInforce为GTD时必传；传入的时间戳仅保留秒级精度，毫秒级部分会被自动忽略，时间戳需大于当前时间+600s且小于253402300799000
	RecvWindow              *int64           `json:"recvWindow"`              // NO
	Timestamp               *int64           `json:"timestamp"`               // YES
}

// YES 交易对
func (api *PortfolioMarginUmConditionalOrderPostApi) Symbol(symbol string) *PortfolioMarginUmConditionalOrderPostApi {
	api.req.Symbol = GetPointer(symbol)
	return api
}

// YES 方向
func (api *PortfolioMarginUmConditionalOrderPostApi) Side(side string) *PortfolioMarginUmConditionalOrderPostApi {
	api.req.Side = GetPointer(side)
	return api
}

// NO 持仓方向，单向持仓模式下非必填，默认且仅可填BOTH;在双向持仓模式下必填,且仅可选择 LONG 或 SHORT
func (api *PortfolioMarginUmConditionalOrderPostApi) PositionSide(positionSide string) *PortfolioMarginUmConditionalOrderPostApi {
	api.req.PositionSide = GetPointer(positionSide)
	return api
}

// YES 条件单类型"STOP", "STOP_MARKET", "TAKE_PROFIT", "TAKE_PROFIT_MARKET"或"TRAILING_STOP_MARKET"
func (api *PortfolioMarginUmConditionalOrderPostApi) StrategyType(strategyType string) *PortfolioMarginUmConditionalOrderPostApi {
	api.req.StrategyType = GetPointer(strategyType)
	return api
}

// NO
func (api *PortfolioMarginUmConditionalOrderPostApi) TimeInForce(timeInForce string) *PortfolioMarginUmConditionalOrderPostApi {
	api.req.TimeInForce = GetPointer(timeInForce)
	return api
}

// NO
func (api *PortfolioMarginUmConditionalOrderPostApi) Quantity(quantity decimal.Decimal) *PortfolioMarginUmConditionalOrderPostApi {
	api.req.Quantity = &quantity
	return api
}

// NO true或false; 非双开模式下默认false；双开模式下不接受此参数
func (api *PortfolioMarginUmConditionalOrderPostApi) ReduceOnly(reduceOnly string) *PortfolioMarginUmConditionalOrderPostApi {
	api.req.ReduceOnly = GetPointer(reduceOnly)
	return api
}

// NO
func (api *PortfolioMarginUmConditionalOrderPostApi) Price(price decimal.Decimal) *PortfolioMarginUmConditionalOrderPostApi {
	api.req.Price = &price
	return api
}

// NO stopPrice 触发类型: MARK_PRICE(标记价格), CONTRACT_PRICE(合约最新价). 默认 CONTRACT_PRICE
func (api *PortfolioMarginUmConditionalOrderPostApi) WorkingType(workingType string) *PortfolioMarginUmConditionalOrderPostApi {
	api.req.WorkingType = GetPointer(workingType)
	return api
}

// NO 条件单触发保护："TRUE","FALSE", 默认"FALSE". 仅 STOP, STOP_MARKET, TAKE_PROFIT, TAKE_PROFIT_MARKET 需要此参数
func (api *PortfolioMarginUmConditionalOrderPostApi) PriceProtect(priceProtect string) *PortfolioMarginUmConditionalOrderPostApi {
	api.req.PriceProtect = GetPointer(priceProtect)
	return api
}

// NO 用户自定义的订单号，不可以重复出现在挂单中。如空缺系统会自动赋值。必须满足正则规则: ^[\.A-Z\:/a-z0-9_-]{1,32}$
func (api *PortfolioMarginUmConditionalOrderPostApi) NewClientStrategyId(newClientStrategyId string) *PortfolioMarginUmConditionalOrderPostApi {
	api.req.NewClientStrategyId = GetPointer(newClientStrategyId)
	return api
}

// NO Used with STOP/STOP_MARKET or TAKE_PROFIT/TAKE_PROFIT_MARKET orders.
func (api *PortfolioMarginUmConditionalOrderPostApi) StopPrice(stopPrice decimal.Decimal) *PortfolioMarginUmConditionalOrderPostApi {
	api.req.StopPrice = &stopPrice
	return api
}

// NO TRAILING_STOP_MARKET 单使用，默认标记价格
func (api *PortfolioMarginUmConditionalOrderPostApi) ActivationPrice(activationPrice decimal.Decimal) *PortfolioMarginUmConditionalOrderPostApi {
	api.req.ActivationPrice = &activationPrice
	return api
}

// NO TRAILING_STOP_MARKET 单使用, 最小0.1, 最大5，1代表1%
func (api *PortfolioMarginUmConditionalOrderPostApi) CallbackRate(callbackRate decimal.Decimal) *PortfolioMarginUmConditionalOrderPostApi {
	api.req.CallbackRate = &callbackRate
	return api
}

// NO OPPONENT/ OPPONENT_5/ OPPONENT_10/ OPPONENT_20/QUEUE/ QUEUE_5/ QUEUE_10/ QUEUE_20；不能与price同时传
func (api *PortfolioMarginUmConditionalOrderPostApi) PriceMatch(priceMatch string) *PortfolioMarginUmConditionalOrderPostApi {
	api.req.PriceMatch = GetPointer(priceMatch)
	return api
}

// NO NONE / EXPIRE_TAKER/ EXPIRE_MAKER/ EXPIRE_BOTH； 默认NONE
func (api *PortfolioMarginUmConditionalOrderPostApi) SelfTradePreventionMode(selfTradePreventionMode string) *PortfolioMarginUmConditionalOrderPostApi {
	api.req.SelfTradePreventionMode = GetPointer(selfTradePreventionMode)
	return api
}

// NO TIF为GTD时订单的自动取消时间， 当timeInforce为GTD时必传；传入的时间戳仅保留秒级精度，毫秒级部分会被自动忽略，时间戳需大于当前时间+600s且小于253402300799000
func (api *PortfolioMarginUmConditionalOrderPostApi) GoodTillDate(goodTillDate int64) *PortfolioMarginUmConditionalOrderPostApi {
	api.req.GoodTillDate = &goodTillDate
	return api
}

// NO
func (api *PortfolioMarginUmConditionalOrderPostApi) RecvWindow(recvWindow int64) *PortfolioMarginUmConditionalOrderPostApi {
	api.req.RecvWindow = &recvWindow
	return api
}

// YES
func (api *PortfolioMarginUmConditionalOrderPostApi) Timestamp(timestamp int64) *PortfolioMarginUmConditionalOrderPostApi {
	api.req.Timestamp = GetPointer(timestamp)
	return api
}

type PortfolioMarginCmOrderPostReq struct {
	Symbol           *string          `json:"symbol"`           // YES 交易对
	Side             *string          `json:"side"`             // YES 方向
	PositionSide     *string          `json:"positionSide"`     // NO 持仓方向，单向持仓模式下非必填，默认且仅可填BOTH;在双向持仓模式下必填,且仅可选择 LONG 或 SHORT
	Type             *string          `json:"type"`             // YES LIMIT, MARKET
	TimeInForce      *string          `json:"timeInForce"`      // NO 有效方法
	Quantity         *decimal.Decimal `json:"quantity"`         // NO 下单数量
	ReduceOnly       *string          `json:"reduceOnly"`       // NO true或false; 非双开模式下默认false；双开模式下不接受此参数
	Price            *decimal.Decimal `json:"price"`            // NO 委托价格
	NewClientOrderId *string          `json:"newClientOrderId"` // NO 用户自定义的订单号，不可以重复出现在挂单中。如空缺系统会自动赋值。必须满足正则规则: ^[\.A-Z\:/a-z0-9_-]{1,32}$
	NewOrderRespType *string          `json:"newOrderRespType"` // NO ACK， RESULT，默认 ACK
	RecvWindow       *int64           `json:"recvWindow"`       // NO
	Timestamp        *int64           `json:"timestamp"`        // YES
}

// YES 交易对
func (api *PortfolioMarginCmOrderPostApi) Symbol(symbol string) *PortfolioMarginCmOrderPostApi {
	api.req.Symbol = GetPointer(symbol)
	return api
}

// YES 方向
func (api *PortfolioMarginCmOrderPostApi) Side(side string) *PortfolioMarginCmOrderPostApi {
	api.req.Side = GetPointer(side)
	return api
}

// NO 持仓方向，单向持仓模式下非必填，默认且仅可填BOTH;在双向持仓模式下必填,且仅可选择 LONG 或 SHORT
func (api *PortfolioMarginCmOrderPostApi) PositionSide(positionSide string) *PortfolioMarginCmOrderPostApi {
	api.req.PositionSide = GetPointer(positionSide)
	return api
}

// YES LIMIT, MARKET
func (api *PortfolioMarginCmOrderPostApi) Type(type_ string) *PortfolioMarginCmOrderPostApi {
	api.req.Type = GetPointer(type_)
	return api
}

// NO 有效方法
func (api *PortfolioMarginCmOrderPostApi) TimeInForce(timeInForce string) *PortfolioMarginCmOrderPostApi {
	api.req.TimeInForce = GetPointer(timeInForce)
	return api
}

// NO 下单数量
func (api *PortfolioMarginCmOrderPostApi) Quantity(quantity decimal.Decimal) *PortfolioMarginCmOrderPostApi {
	api.req.Quantity = &quantity
	return api
}

// NO true或false; 非双开模式下默认false；双开模式下不接受此参数
func (api *PortfolioMarginCmOrderPostApi) ReduceOnly(reduceOnly string) *PortfolioMarginCmOrderPostApi {
	api.req.ReduceOnly = GetPointer(reduceOnly)
	return api
}

// NO 委托价格
func (api *PortfolioMarginCmOrderPostApi) Price(price decimal.Decimal) *PortfolioMarginCmOrderPostApi {
	api.req.Price = &price
	return api
}

// NO 用户自定义的订单号，不可以重复出现在挂单中。如空缺系统会自动赋值。必须满足正则规则: ^[\.A-Z\:/a-z0-9_-]{1,32}$
func (api *PortfolioMarginCmOrderPostApi) NewClientOrderId(newClientOrderId string) *PortfolioMarginCmOrderPostApi {
	api.req.NewClientOrderId = GetPointer(newClientOrderId)
	return api
}

// NO ACK， RESULT，默认 ACK
func (api *PortfolioMarginCmOrderPostApi) NewOrderRespType(newOrderRespType string) *PortfolioMarginCmOrderPostApi {
	api.req.NewOrderRespType = GetPointer(newOrderRespType)
	return api
}

// NO
func (api *PortfolioMarginCmOrderPostApi) RecvWindow(recvWindow int64) *PortfolioMarginCmOrderPostApi {
	api.req.RecvWindow = GetPointer(recvWindow)
	return api
}

// YES
func (api *PortfolioMarginCmOrderPostApi) Timestamp(timestamp int64) *PortfolioMarginCmOrderPostApi {
	api.req.Timestamp = GetPointer(timestamp)
	return api
}

type PortfolioMarginCmOrderPostApi struct {
	client *PortfolioMarginRestClient
	req    *PortfolioMarginCmOrderPostReq
}

type PortfolioMarginCmConditionalOrderPostReq struct {
	Symbol              *string          `json:"symbol"`              // YES 交易对
	Side                *string          `json:"side"`                // YES 方向
	PositionSide        *string          `json:"positionSide"`        // NO 持仓方向，单向持仓模式下非必填，默认且仅可填BOTH;在双向持仓模式下必填,且仅可选择 LONG 或 SHORT
	StrategyType        *string          `json:"strategyType"`        // YES 条件单类型"STOP", "STOP_MARKET", "TAKE_PROFIT", "TAKE_PROFIT_MARKET"或"TRAILING_STOP_MARKET"
	TimeInForce         *string          `json:"timeInForce"`         // NO
	Quantity            *decimal.Decimal `json:"quantity"`            // NO
	ReduceOnly          *string          `json:"reduceOnly"`          // NO true或false; 非双开模式下默认false；双开模式下不接受此参数
	Price               *decimal.Decimal `json:"price"`               // NO
	WorkingType         *string          `json:"workingType"`         // NO stopPrice 触发类型: MARK_PRICE(标记价格), CONTRACT_PRICE(合约最新价). 默认 CONTRACT_PRICE
	PriceProtect        *string          `json:"priceProtect"`        // NO 条件单触发保护："TRUE","FALSE", 默认"FALSE". 仅 STOP, STOP_MARKET, TAKE_PROFIT, TAKE_PROFIT_MARKET 需要此参数
	NewClientStrategyId *string          `json:"newClientStrategyId"` // NO 用户自定义的订单号，不可以重复出现在挂单中。如空缺系统会自动赋值。必须满足正则规则: ^[\.A-Z\:/a-z0-9_-]{1,32}$
	StopPrice           *decimal.Decimal `json:"stopPrice"`           // NO Used with STOP/STOP_MARKET or TAKE_PROFIT/TAKE_PROFIT_MARKET orders.
	ActivationPrice     *decimal.Decimal `json:"activationPrice"`     // NO TRAILING_STOP_MARKET 单使用，默认标记价格
	CallbackRate        *decimal.Decimal `json:"callbackRate"`        // NO TRAILING_STOP_MARKET 单使用, 最小0.1, 最大5，1代表1%
	RecvWindow          *int64           `json:"recvWindow"`          // NO
	Timestamp           *int64           `json:"timestamp"`           // YES
}

// YES 交易对
func (api *PortfolioMarginCmConditionalOrderPostApi) Symbol(symbol string) *PortfolioMarginCmConditionalOrderPostApi {
	api.req.Symbol = GetPointer(symbol)
	return api
}

// YES 方向
func (api *PortfolioMarginCmConditionalOrderPostApi) Side(side string) *PortfolioMarginCmConditionalOrderPostApi {
	api.req.Side = GetPointer(side)
	return api
}

// NO 持仓方向，单向持仓模式下非必填，默认且仅可填BOTH;在双向持仓模式下必填,且仅可选择 LONG 或 SHORT
func (api *PortfolioMarginCmConditionalOrderPostApi) PositionSide(positionSide string) *PortfolioMarginCmConditionalOrderPostApi {
	api.req.PositionSide = GetPointer(positionSide)
	return api
}

// YES 条件单类型"STOP", "STOP_MARKET", "TAKE_PROFIT", "TAKE_PROFIT_MARKET"或"TRAILING_STOP_MARKET"
func (api *PortfolioMarginCmConditionalOrderPostApi) StrategyType(strategyType string) *PortfolioMarginCmConditionalOrderPostApi {
	api.req.StrategyType = GetPointer(strategyType)
	return api
}

// NO
func (api *PortfolioMarginCmConditionalOrderPostApi) TimeInForce(timeInForce string) *PortfolioMarginCmConditionalOrderPostApi {
	api.req.TimeInForce = GetPointer(timeInForce)
	return api
}

// NO
func (api *PortfolioMarginCmConditionalOrderPostApi) Quantity(quantity decimal.Decimal) *PortfolioMarginCmConditionalOrderPostApi {
	api.req.Quantity = &quantity
	return api
}

// NO true或false; 非双开模式下默认false；双开模式下不接受此参数
func (api *PortfolioMarginCmConditionalOrderPostApi) ReduceOnly(reduceOnly string) *PortfolioMarginCmConditionalOrderPostApi {
	api.req.ReduceOnly = GetPointer(reduceOnly)
	return api
}

// NO
func (api *PortfolioMarginCmConditionalOrderPostApi) Price(price decimal.Decimal) *PortfolioMarginCmConditionalOrderPostApi {
	api.req.Price = &price
	return api
}

// NO stopPrice 触发类型: MARK_PRICE(标记价格), CONTRACT_PRICE(合约最新价). 默认 CONTRACT_PRICE
func (api *PortfolioMarginCmConditionalOrderPostApi) WorkingType(workingType string) *PortfolioMarginCmConditionalOrderPostApi {
	api.req.WorkingType = GetPointer(workingType)
	return api
}

// NO 条件单触发保护："TRUE","FALSE", 默认"FALSE". 仅 STOP, STOP_MARKET, TAKE_PROFIT, TAKE_PROFIT_MARKET 需要此参数
func (api *PortfolioMarginCmConditionalOrderPostApi) PriceProtect(priceProtect string) *PortfolioMarginCmConditionalOrderPostApi {
	api.req.PriceProtect = GetPointer(priceProtect)
	return api
}

// NO 用户自定义的订单号，不可以重复出现在挂单中。如空缺系统会自动赋值。必须满足正则规则: ^[\.A-Z\:/a-z0-9_-]{1,32}$
func (api *PortfolioMarginCmConditionalOrderPostApi) NewClientStrategyId(newClientStrategyId string) *PortfolioMarginCmConditionalOrderPostApi {
	api.req.NewClientStrategyId = GetPointer(newClientStrategyId)
	return api
}

// NO Used with STOP/STOP_MARKET or TAKE_PROFIT/TAKE_PROFIT_MARKET orders.
func (api *PortfolioMarginCmConditionalOrderPostApi) StopPrice(stopPrice decimal.Decimal) *PortfolioMarginCmConditionalOrderPostApi {
	api.req.StopPrice = &stopPrice
	return api
}

// NO TRAILING_STOP_MARKET 单使用，默认标记价格
func (api *PortfolioMarginCmConditionalOrderPostApi) ActivationPrice(activationPrice decimal.Decimal) *PortfolioMarginCmConditionalOrderPostApi {
	api.req.ActivationPrice = &activationPrice
	return api
}

// NO TRAILING_STOP_MARKET 单使用, 最小0.1, 最大5，1代表1%
func (api *PortfolioMarginCmConditionalOrderPostApi) CallbackRate(callbackRate decimal.Decimal) *PortfolioMarginCmConditionalOrderPostApi {
	api.req.CallbackRate = &callbackRate
	return api
}

// NO
func (api *PortfolioMarginCmConditionalOrderPostApi) RecvWindow(recvWindow int64) *PortfolioMarginCmConditionalOrderPostApi {
	api.req.RecvWindow = &recvWindow
	return api
}

// YES
func (api *PortfolioMarginCmConditionalOrderPostApi) Timestamp(timestamp int64) *PortfolioMarginCmConditionalOrderPostApi {
	api.req.Timestamp = GetPointer(timestamp)
	return api
}

type PortfolioMarginCmConditionalOrderPostApi struct {
	client *PortfolioMarginRestClient
	req    *PortfolioMarginCmConditionalOrderPostReq
}

type PortfolioMarginMarginOrderPostReq struct {
	Symbol                  *string          `json:"symbol"`                  // YES 交易对
	Side                    *string          `json:"side"`                    // YES 方向
	Type                    *string          `json:"type"`                    // YES 详见枚举定义：订单类型
	Quantity                *decimal.Decimal `json:"quantity"`                // NO
	QuoteOrderQty           *decimal.Decimal `json:"quoteOrderQty"`           // NO
	Price                   *decimal.Decimal `json:"price"`                   // NO
	StopPrice               *decimal.Decimal `json:"stopPrice"`               // NO 与 STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT，和 TAKE_PROFIT_LIMIT 订单一起使用
	NewClientOrderId        *string          `json:"newClientOrderId"`        // NO 客户自定义的唯一订单ID。若未发送自动生成。
	NewOrderRespType        *string          `json:"newOrderRespType"`        // NO 设置响应: JSON。 ACK， RESULT， 或 FULL; MARKET 和 LIMIT 订单类型默认为 FULL， 所有其他订单默认为 ACK
	IcebergQty              *decimal.Decimal `json:"icebergQty"`              // NO 与 LIMIT， STOP_LOSS_LIMIT，和 TAKE_PROFIT_LIMIT 一起使用创建 iceberg 订单
	SideEffectType          *string          `json:"sideEffectType"`          // NO NO_SIDE_EFFECT， MARGIN_BUY，AUTO_REPAY,AUTO_BORROW_REPAY；默认为 NO_SIDE_EFFECT
	TimeInForce             *string          `json:"timeInForce"`             // NO GTC，IOC，FOK
	SelfTradePreventionMode *string          `json:"selfTradePreventionMode"` // NO 允许的 ENUM 取决于交易对的配置。支持的值有 EXPIRE_TAKER，EXPIRE_MAKER，EXPIRE_BOTH，NONE
	AutoRepayAtCancel       *string          `json:"autoRepayAtCancel"`       // NO Only when MARGIN_BUY or AUTO_BORROW_REPAY order takes effect, true means that the debt generated by the order needs to be repay after the order is cancelled. The default is true
	RecvWindow              *int64           `json:"recvWindow"`              // NO 赋值不能超过 60000
	Timestamp               *int64           `json:"timestamp"`               // YES
}

// YES 交易对
func (api *PortfolioMarginMarginOrderPostApi) Symbol(symbol string) *PortfolioMarginMarginOrderPostApi {
	api.req.Symbol = GetPointer(symbol)
	return api
}

// YES 方向
func (api *PortfolioMarginMarginOrderPostApi) Side(side string) *PortfolioMarginMarginOrderPostApi {
	api.req.Side = GetPointer(side)
	return api
}

// YES 详见枚举定义：订单类型
func (api *PortfolioMarginMarginOrderPostApi) Type(type_ string) *PortfolioMarginMarginOrderPostApi {
	api.req.Type = GetPointer(type_)
	return api
}

// NO
func (api *PortfolioMarginMarginOrderPostApi) Quantity(quantity decimal.Decimal) *PortfolioMarginMarginOrderPostApi {
	api.req.Quantity = &quantity
	return api
}

// NO
func (api *PortfolioMarginMarginOrderPostApi) QuoteOrderQty(quoteOrderQty decimal.Decimal) *PortfolioMarginMarginOrderPostApi {
	api.req.QuoteOrderQty = &quoteOrderQty
	return api
}

// NO
func (api *PortfolioMarginMarginOrderPostApi) Price(price decimal.Decimal) *PortfolioMarginMarginOrderPostApi {
	api.req.Price = &price
	return api
}

// NO 与 STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT，和 TAKE_PROFIT_LIMIT 订单一起使用
func (api *PortfolioMarginMarginOrderPostApi) StopPrice(stopPrice decimal.Decimal) *PortfolioMarginMarginOrderPostApi {
	api.req.StopPrice = &stopPrice
	return api
}

// NO 客户自定义的唯一订单ID。若未发送自动生成。
func (api *PortfolioMarginMarginOrderPostApi) NewClientOrderId(newClientOrderId string) *PortfolioMarginMarginOrderPostApi {
	api.req.NewClientOrderId = GetPointer(newClientOrderId)
	return api
}

// NO 设置响应: JSON。 ACK， RESULT， 或 FULL; MARKET 和 LIMIT 订单类型默认为 FULL， 所有其他订单默认为 ACK
func (api *PortfolioMarginMarginOrderPostApi) NewOrderRespType(newOrderRespType string) *PortfolioMarginMarginOrderPostApi {
	api.req.NewOrderRespType = GetPointer(newOrderRespType)
	return api
}

// NO 与 LIMIT， STOP_LOSS_LIMIT，和 TAKE_PROFIT_LIMIT 一起使用创建 iceberg 订单
func (api *PortfolioMarginMarginOrderPostApi) IcebergQty(icebergQty decimal.Decimal) *PortfolioMarginMarginOrderPostApi {
	api.req.IcebergQty = &icebergQty
	return api
}

// NO NO_SIDE_EFFECT， MARGIN_BUY，AUTO_REPAY,AUTO_BORROW_REPAY；默认为 NO_SIDE_EFFECT
func (api *PortfolioMarginMarginOrderPostApi) SideEffectType(sideEffectType string) *PortfolioMarginMarginOrderPostApi {
	api.req.SideEffectType = GetPointer(sideEffectType)
	return api
}

// NO GTC，IOC，FOK
func (api *PortfolioMarginMarginOrderPostApi) TimeInForce(timeInForce string) *PortfolioMarginMarginOrderPostApi {
	api.req.TimeInForce = GetPointer(timeInForce)
	return api
}

// NO 允许的 ENUM 取决于交易对的配置。支持的值有 EXPIRE_TAKER，EXPIRE_MAKER，EXPIRE_BOTH，NONE
func (api *PortfolioMarginMarginOrderPostApi) SelfTradePreventionMode(selfTradePreventionMode string) *PortfolioMarginMarginOrderPostApi {
	api.req.SelfTradePreventionMode = GetPointer(selfTradePreventionMode)
	return api
}

// NO Only when MARGIN_BUY or AUTO_BORROW_REPAY order takes effect, true means that the debt generated by the order needs to be repay after the order is cancelled. The default is true
func (api *PortfolioMarginMarginOrderPostApi) AutoRepayAtCancel(autoRepayAtCancel string) *PortfolioMarginMarginOrderPostApi {
	api.req.AutoRepayAtCancel = GetPointer(autoRepayAtCancel)
	return api
}

// NO 赋值不能超过 60000
func (api *PortfolioMarginMarginOrderPostApi) RecvWindow(recvWindow int64) *PortfolioMarginMarginOrderPostApi {
	api.req.RecvWindow = &recvWindow
	return api
}

// YES
func (api *PortfolioMarginMarginOrderPostApi) Timestamp(timestamp int64) *PortfolioMarginMarginOrderPostApi {
	api.req.Timestamp = GetPointer(timestamp)
	return api
}

type PortfolioMarginMarginOrderPostApi struct {
	client *PortfolioMarginRestClient
	req    *PortfolioMarginMarginOrderPostReq
}

// symbol	STRING	YES
// listClientOrderId	STRING	NO	整个 orderList 的唯一 ID
// side	ENUM	YES	详见枚举定义：订单方向
// quantity	DECIMAL	YES
// limitClientOrderId	STRING	NO	限价单的唯一 ID
// price	DECIMAL	YES
// limitIcebergQty	DECIMAL	NO
// stopClientOrderId	STRING	NO	止损/止损限价单的唯一 ID
// stopPrice	DECIMAL	YES
// stopLimitPrice	DECIMAL	NO	如果提供，须配合提交stopLimitTimeInForce
// stopIcebergQty	DECIMAL	NO
// stopLimitTimeInForce	ENUM	NO	有效值 GTC/FOK/IOC
// newOrderRespType	ENUM	NO	详见枚举定义：订单返回类型
// sideEffectType	ENUM	NO	NO_SIDE_EFFECT, MARGIN_BUY, AUTO_REPAY; 默认为 NO_SIDE_EFFECT
// recvWindow	LONG	NO	不能大于 60000
// timestamp	LONG	YES
type PortfolioMarginMarginOrderOcoPostReq struct {
	Symbol               *string          `json:"symbol"`               // YES
	ListClientOrderId    *string          `json:"listClientOrderId"`    // NO	整个 orderList 的唯一 ID
	Side                 *string          `json:"side"`                 // YES	详见枚举定义：订��方向
	Quantity             *decimal.Decimal `json:"quantity"`             // YES
	LimitClientOrderId   *string          `json:"limitClientOrderId"`   // NO	限价单的唯一 ID
	Price                *decimal.Decimal `json:"price"`                // YES
	LimitIcebergQty      *decimal.Decimal `json:"limitIcebergQty"`      // NO
	StopClientOrderId    *string          `json:"stopClientOrderId"`    // NO	止损/止损限价单的唯一 ID
	StopPrice            *decimal.Decimal `json:"stopPrice"`            // YES
	StopLimitPrice       *decimal.Decimal `json:"stopLimitPrice"`       // NO	如果提供，须配合提交stopLimitTimeInForce
	StopIcebergQty       *decimal.Decimal `json:"stopIcebergQty"`       // NO
	StopLimitTimeInForce *string          `json:"stopLimitTimeInForce"` // NO	有效值 GTC/FOK/IOC
	NewOrderRespType     *string          `json:"newOrderRespType"`     // NO	详见枚举定义：订单返回类型
	SideEffectType       *string          `json:"sideEffectType"`       // NO	NO_SIDE_EFFECT, MARGIN_BUY, AUTO_REPAY; 默认为 NO_SIDE_EFFECT
	RecvWindow           *int64           `json:"recvWindow"`           // NO	不能大于 60000
	Timestamp            *int64           `json:"timestamp"`            // YES
}

// YES
func (api *PortfolioMarginMarginOrderOcoPostApi) Symbol(symbol string) *PortfolioMarginMarginOrderOcoPostApi {
	api.req.Symbol = GetPointer(symbol)
	return api
}

// NO	整个 orderList 的唯一 ID
func (api *PortfolioMarginMarginOrderOcoPostApi) ListClientOrderId(listClientOrderId string) *PortfolioMarginMarginOrderOcoPostApi {
	api.req.ListClientOrderId = GetPointer(listClientOrderId)
	return api
}

// YES	详见枚举定义：订单方向
func (api *PortfolioMarginMarginOrderOcoPostApi) Side(side string) *PortfolioMarginMarginOrderOcoPostApi {
	api.req.Side = GetPointer(side)
	return api
}

// YES
func (api *PortfolioMarginMarginOrderOcoPostApi) Quantity(quantity decimal.Decimal) *PortfolioMarginMarginOrderOcoPostApi {
	api.req.Quantity = &quantity
	return api
}

// NO	限价单的唯一 ID
func (api *PortfolioMarginMarginOrderOcoPostApi) LimitClientOrderId(limitClientOrderId string) *PortfolioMarginMarginOrderOcoPostApi {
	api.req.LimitClientOrderId = GetPointer(limitClientOrderId)
	return api
}

// YES
func (api *PortfolioMarginMarginOrderOcoPostApi) Price(price decimal.Decimal) *PortfolioMarginMarginOrderOcoPostApi {
	api.req.Price = &price
	return api
}

// NO
func (api *PortfolioMarginMarginOrderOcoPostApi) LimitIcebergQty(limitIcebergQty decimal.Decimal) *PortfolioMarginMarginOrderOcoPostApi {
	api.req.LimitIcebergQty = &limitIcebergQty
	return api
}

// NO	止损/止损限价单的唯一 ID
func (api *PortfolioMarginMarginOrderOcoPostApi) StopClientOrderId(stopClientOrderId string) *PortfolioMarginMarginOrderOcoPostApi {
	api.req.StopClientOrderId = GetPointer(stopClientOrderId)
	return api
}

// YES
func (api *PortfolioMarginMarginOrderOcoPostApi) StopPrice(stopPrice decimal.Decimal) *PortfolioMarginMarginOrderOcoPostApi {
	api.req.StopPrice = &stopPrice
	return api
}

// NO	如果提供，须配合提交stopLimitTimeInForce
func (api *PortfolioMarginMarginOrderOcoPostApi) StopLimitPrice(stopLimitPrice decimal.Decimal) *PortfolioMarginMarginOrderOcoPostApi {
	api.req.StopLimitPrice = &stopLimitPrice
	return api
}

// NO
func (api *PortfolioMarginMarginOrderOcoPostApi) StopIcebergQty(stopIcebergQty decimal.Decimal) *PortfolioMarginMarginOrderOcoPostApi {
	api.req.StopIcebergQty = &stopIcebergQty
	return api
}

// NO	有效值 GTC/FOK/IOC
func (api *PortfolioMarginMarginOrderOcoPostApi) StopLimitTimeInForce(stopLimitTimeInForce string) *PortfolioMarginMarginOrderOcoPostApi {
	api.req.StopLimitTimeInForce = GetPointer(stopLimitTimeInForce)
	return api
}

// NO	详见枚举定义：订单返回类型
func (api *PortfolioMarginMarginOrderOcoPostApi) NewOrderRespType(newOrderRespType string) *PortfolioMarginMarginOrderOcoPostApi {
	api.req.NewOrderRespType = GetPointer(newOrderRespType)
	return api
}

// NO	NO_SIDE_EFFECT, MARGIN_BUY, AUTO_REPAY; 默认为 NO_SIDE_EFFECT
func (api *PortfolioMarginMarginOrderOcoPostApi) SideEffectType(sideEffectType string) *PortfolioMarginMarginOrderOcoPostApi {
	api.req.SideEffectType = GetPointer(sideEffectType)
	return api
}

// NO	不能大于 60000
func (api *PortfolioMarginMarginOrderOcoPostApi) RecvWindow(recvWindow int64) *PortfolioMarginMarginOrderOcoPostApi {
	api.req.RecvWindow = GetPointer(recvWindow)
	return api
}

// YES
func (api *PortfolioMarginMarginOrderOcoPostApi) Timestamp(timestamp int64) *PortfolioMarginMarginOrderOcoPostApi {
	api.req.Timestamp = GetPointer(timestamp)
	return api
}

type PortfolioMarginMarginOrderOcoPostApi struct {
	client *PortfolioMarginRestClient
	req    *PortfolioMarginMarginOrderOcoPostReq
}

type PortfolioMarginUmConditionalOpenOrderReq struct {
	Symbol              *string `json:"symbol"`              // YES 交易对
	StrategyId          *int64  `json:"strategyId"`          // YES 用户
	NewClientStrategyId *string `json:"newClientStrategyId"` // NO 用户自定义的订单号，不可以重复出现在挂单中。如空缺系统会自动赋值。必须满足正则规则: ^[\.A-Z\:/a-z0-9_-]{1,32}$
	RecvWindow          *int64  `json:"recvWindow"`          // NO
	Timestamp           *int64  `json:"timestamp"`           // YES
}

// YES 交易对
func (api *PortfolioMarginUmConditionalOpenOrderApi) Symbol(symbol string) *PortfolioMarginUmConditionalOpenOrderApi {
	api.req.Symbol = GetPointer(symbol)
	return api
}

// YES 用户
func (api *PortfolioMarginUmConditionalOpenOrderApi) StrategyId(strategyId int64) *PortfolioMarginUmConditionalOpenOrderApi {
	api.req.StrategyId = GetPointer(strategyId)
	return api
}

// NO 用户自定义的订单号，不可以重复出现在挂单中。如空缺系统会自动赋值。必须满足正则规则: ^[\.A-Z\:/a-z0-9_-]{1,32}$
func (api *PortfolioMarginUmConditionalOpenOrderApi) NewClientStrategyId(newClientStrategyId string) *PortfolioMarginUmConditionalOpenOrderApi {
	api.req.NewClientStrategyId = GetPointer(newClientStrategyId)
	return api
}

// NO
func (api *PortfolioMarginUmConditionalOpenOrderApi) RecvWindow(recvWindow int64) *PortfolioMarginUmConditionalOpenOrderApi {
	api.req.RecvWindow = &recvWindow
	return api
}

// YES
func (api *PortfolioMarginUmConditionalOpenOrderApi) Timestamp(timestamp int64) *PortfolioMarginUmConditionalOpenOrderApi {
	api.req.Timestamp = GetPointer(timestamp)
	return api
}

type PortfolioMarginUmConditionalOpenOrderApi struct {
	client *PortfolioMarginRestClient
	req    *PortfolioMarginUmConditionalOpenOrderReq
}

type PortfolioMarginUmConditionalOpenOrdersReq struct {
	Symbol     *string `json:"symbol"`     // YES 交易对
	RecvWindow *int64  `json:"recvWindow"` // NO
	Timestamp  *int64  `json:"timestamp"`  // YES
}

// YES 交易对
func (api *PortfolioMarginUmConditionalOpenOrdersApi) Symbol(symbol string) *PortfolioMarginUmConditionalOpenOrdersApi {
	api.req.Symbol = GetPointer(symbol)
	return api
}

// NO
func (api *PortfolioMarginUmConditionalOpenOrdersApi) RecvWindow(recvWindow int64) *PortfolioMarginUmConditionalOpenOrdersApi {
	api.req.RecvWindow = &recvWindow
	return api
}

// YES
func (api *PortfolioMarginUmConditionalOpenOrdersApi) Timestamp(timestamp int64) *PortfolioMarginUmConditionalOpenOrdersApi {
	api.req.Timestamp = GetPointer(timestamp)
	return api
}

type PortfolioMarginUmConditionalOpenOrdersApi struct {
	client *PortfolioMarginRestClient
	req    *PortfolioMarginUmConditionalOpenOrdersReq
}

type PortfolioMarginUmConditionalOrderHistoryReq struct {
	Symbol              *string `json:"symbol"`              // YES 交易对
	StrategyId          *int64  `json:"strategyId"`          // NO
	NewClientStrategyId *string `json:"newClientStrategyId"` // NO
	RecvWindow          *int64  `json:"recvWindow"`          // NO
	Timestamp           *int64  `json:"timestamp"`           // YES
}

// YES 交易对
func (api *PortfolioMarginUmConditionalOrderHistoryApi) Symbol(symbol string) *PortfolioMarginUmConditionalOrderHistoryApi {
	api.req.Symbol = GetPointer(symbol)
	return api
}

// NO
func (api *PortfolioMarginUmConditionalOrderHistoryApi) StrategyId(strategyId int64) *PortfolioMarginUmConditionalOrderHistoryApi {
	api.req.StrategyId = GetPointer(strategyId)
	return api
}

// NO 用户自定义的订单号，不可以重复出现在挂单中。如空缺系统会自动赋值。必须满足正则规则: ^[\.A-Z\:/a-z0-9_-]{1,32}$
func (api *PortfolioMarginUmConditionalOrderHistoryApi) NewClientStrategyId(newClientStrategyId string) *PortfolioMarginUmConditionalOrderHistoryApi {
	api.req.NewClientStrategyId = GetPointer(newClientStrategyId)
	return api
}

// NO
func (api *PortfolioMarginUmConditionalOrderHistoryApi) RecvWindow(recvWindow int64) *PortfolioMarginUmConditionalOrderHistoryApi {
	api.req.RecvWindow = GetPointer(recvWindow)
	return api
}

// YES
func (api *PortfolioMarginUmConditionalOrderHistoryApi) Timestamp(timestamp int64) *PortfolioMarginUmConditionalOrderHistoryApi {
	api.req.Timestamp = GetPointer(timestamp)
	return api
}

type PortfolioMarginUmConditionalOrderHistoryApi struct {
	client *PortfolioMarginRestClient
	req    *PortfolioMarginUmConditionalOrderHistoryReq
}

type PortfolioMarginCmOrderDeleteReq struct {
	Symbol            *string `json:"symbol"`            // YES 交易对
	OrderId           *int64  `json:"orderId"`           // NO
	OrigClientOrderId *string `json:"origClientOrderId"` // NO
	NewClientOrderId  *string `json:"newClientOrderId"`  // NO
	RecvWindow        *int64  `json:"recvWindow"`        // NO
	Timestamp         *int64  `json:"timestamp"`         // YES
}

// YES 交易对
func (api *PortfolioMarginCmOrderDeleteApi) Symbol(symbol string) *PortfolioMarginCmOrderDeleteApi {
	api.req.Symbol = GetPointer(symbol)
	return api
}

// NO
func (api *PortfolioMarginCmOrderDeleteApi) OrderId(orderId int64) *PortfolioMarginCmOrderDeleteApi {
	api.req.OrderId = GetPointer(orderId)
	return api
}

// NO
func (api *PortfolioMarginCmOrderDeleteApi) OrigClientOrderId(origClientOrderId string) *PortfolioMarginCmOrderDeleteApi {
	api.req.OrigClientOrderId = GetPointer(origClientOrderId)
	return api
}

// NO
func (api *PortfolioMarginCmOrderDeleteApi) NewClientOrderId(newClientOrderId string) *PortfolioMarginCmOrderDeleteApi {
	api.req.NewClientOrderId = GetPointer(newClientOrderId)
	return api
}

// NO
func (api *PortfolioMarginCmOrderDeleteApi) RecvWindow(recvWindow int64) *PortfolioMarginCmOrderDeleteApi {
	api.req.RecvWindow = GetPointer(recvWindow)
	return api
}

// YES
func (api *PortfolioMarginCmOrderDeleteApi) Timestamp(timestamp int64) *PortfolioMarginCmOrderDeleteApi {
	api.req.Timestamp = GetPointer(timestamp)
	return api
}

type PortfolioMarginCmOrderDeleteApi struct {
	client *PortfolioMarginRestClient
	req    *PortfolioMarginCmOrderDeleteReq
}

type PortfolioMarginCmAllOpenOrdersDeleteReq struct {
	Symbol     *string `json:"symbol"`     // YES 交易对
	RecvWindow *int64  `json:"recvWindow"` // NO
	Timestamp  *int64  `json:"timestamp"`  // YES
}

// YES 交易对
func (api *PortfolioMarginCmAllOpenOrdersDeleteApi) Symbol(symbol string) *PortfolioMarginCmAllOpenOrdersDeleteApi {
	api.req.Symbol = GetPointer(symbol)
	return api
}

// NO
func (api *PortfolioMarginCmAllOpenOrdersDeleteApi) RecvWindow(recvWindow int64) *PortfolioMarginCmAllOpenOrdersDeleteApi {
	api.req.RecvWindow = GetPointer(recvWindow)
	return api
}

// YES
func (api *PortfolioMarginCmAllOpenOrdersDeleteApi) Timestamp(timestamp int64) *PortfolioMarginCmAllOpenOrdersDeleteApi {
	api.req.Timestamp = GetPointer(timestamp)
	return api
}

type PortfolioMarginCmAllOpenOrdersDeleteApi struct {
	client *PortfolioMarginRestClient
	req    *PortfolioMarginCmAllOpenOrdersDeleteReq
}

type PortfolioMarginCmOrderPutReq struct {
	OrderId           *int64           `json:"orderId"`           // YES
	OrigClientOrderId *string          `json:"origClientOrderId"` // NO
	Symbol            *string          `json:"symbol"`            // YES 交易对
	Side              *string          `json:"side"`              // YES 交易方向
	Quantity          *decimal.Decimal `json:"quantity"`          // NO 下单数量
	Price             *decimal.Decimal `json:"price"`             // NO 委托价格
	RecvWindow        *int64           `json:"recvWindow"`        // NO
	Timestamp         *int64           `json:"timestamp"`         // YES
}

// YES
func (api *PortfolioMarginCmOrderPutApi) OrderId(orderId int64) *PortfolioMarginCmOrderPutApi {
	api.req.OrderId = GetPointer(orderId)
	return api
}

// NO
func (api *PortfolioMarginCmOrderPutApi) OrigClientOrderId(origClientOrderId string) *PortfolioMarginCmOrderPutApi {
	api.req.OrigClientOrderId = GetPointer(origClientOrderId)
	return api
}

// YES 交易对
func (api *PortfolioMarginCmOrderPutApi) Symbol(symbol string) *PortfolioMarginCmOrderPutApi {
	api.req.Symbol = GetPointer(symbol)
	return api
}

// YES 交易方向
func (api *PortfolioMarginCmOrderPutApi) Side(side string) *PortfolioMarginCmOrderPutApi {
	api.req.Side = GetPointer(side)
	return api
}

// NO 下单数量
func (api *PortfolioMarginCmOrderPutApi) Quantity(quantity decimal.Decimal) *PortfolioMarginCmOrderPutApi {
	api.req.Quantity = &quantity
	return api
}

// NO 委托价格
func (api *PortfolioMarginCmOrderPutApi) Price(price decimal.Decimal) *PortfolioMarginCmOrderPutApi {
	api.req.Price = &price
	return api
}

// NO
func (api *PortfolioMarginCmOrderPutApi) RecvWindow(recvWindow int64) *PortfolioMarginCmOrderPutApi {
	api.req.RecvWindow = GetPointer(recvWindow)
	return api
}

// YES
func (api *PortfolioMarginCmOrderPutApi) Timestamp(timestamp int64) *PortfolioMarginCmOrderPutApi {
	api.req.Timestamp = GetPointer(timestamp)
	return api
}

type PortfolioMarginCmOrderPutApi struct {
	client *PortfolioMarginRestClient
	req    *PortfolioMarginCmOrderPutReq
}

type PortfolioMarginCmConditionalOrderDeleteReq struct {
	Symbol              *string `json:"symbol"`              // YES 交易对
	StrategyId          *int64  `json:"strategyId"`          // YES 策略ID
	NewClientStrategyId *string `json:"newClientStrategyId"` // NO 用户自定义的订单号，不可以重复出现在挂单中。如空缺系统会自动赋值。必须满足正则规则: ^[\.A-Z\:/a-z0-9_-]{1,32}$
	RecvWindow          *int64  `json:"recvWindow"`          // NO
	Timestamp           *int64  `json:"timestamp"`           // YES
}

// YES 交易对
func (api *PortfolioMarginCmConditionalOrderDeleteApi) Symbol(symbol string) *PortfolioMarginCmConditionalOrderDeleteApi {
	api.req.Symbol = GetPointer(symbol)
	return api
}

// YES 策略ID
func (api *PortfolioMarginCmConditionalOrderDeleteApi) StrategyId(strategyId int64) *PortfolioMarginCmConditionalOrderDeleteApi {
	api.req.StrategyId = GetPointer(strategyId)
	return api
}

// NO 用户自定义的订单号，不可以重复出现在挂单中。如空缺系统会自动赋值。必须满足正则规则: ^[\.A-Z\:/a-z0-9_-]{1,32}$
func (api *PortfolioMarginCmConditionalOrderDeleteApi) NewClientStrategyId(newClientStrategyId string) *PortfolioMarginCmConditionalOrderDeleteApi {
	api.req.NewClientStrategyId = GetPointer(newClientStrategyId)
	return api
}

// NO
func (api *PortfolioMarginCmConditionalOrderDeleteApi) RecvWindow(recvWindow int64) *PortfolioMarginCmConditionalOrderDeleteApi {
	api.req.RecvWindow = GetPointer(recvWindow)
	return api
}

// YES
func (api *PortfolioMarginCmConditionalOrderDeleteApi) Timestamp(timestamp int64) *PortfolioMarginCmConditionalOrderDeleteApi {
	api.req.Timestamp = GetPointer(timestamp)
	return api
}

type PortfolioMarginCmConditionalOrderDeleteApi struct {
	client *PortfolioMarginRestClient
	req    *PortfolioMarginCmConditionalOrderDeleteReq
}

type PortfolioMarginCmConditionalAllOpenOrdersDeleteReq struct {
	Symbol     *string `json:"symbol"`     // YES 交易对
	RecvWindow *int64  `json:"recvWindow"` // NO
	Timestamp  *int64  `json:"timestamp"`  // YES
}

// YES 交易对
func (api *PortfolioMarginCmConditionalAllOpenOrdersDeleteApi) Symbol(symbol string) *PortfolioMarginCmConditionalAllOpenOrdersDeleteApi {
	api.req.Symbol = GetPointer(symbol)
	return api
}

// NO
func (api *PortfolioMarginCmConditionalAllOpenOrdersDeleteApi) RecvWindow(recvWindow int64) *PortfolioMarginCmConditionalAllOpenOrdersDeleteApi {
	api.req.RecvWindow = GetPointer(recvWindow)
	return api
}

// YES
func (api *PortfolioMarginCmConditionalAllOpenOrdersDeleteApi) Timestamp(timestamp int64) *PortfolioMarginCmConditionalAllOpenOrdersDeleteApi {
	api.req.Timestamp = GetPointer(timestamp)
	return api
}

type PortfolioMarginCmConditionalAllOpenOrdersDeleteApi struct {
	client *PortfolioMarginRestClient
	req    *PortfolioMarginCmConditionalAllOpenOrdersDeleteReq
}

type PortfolioMarginCmOrderGetReq struct {
	Symbol            *string `json:"symbol"`            // YES 交易对
	OrderId           *int64  `json:"orderId"`           // NO
	OrigClientOrderId *string `json:"origClientOrderId"` // NO
	RecvWindow        *int64  `json:"recvWindow"`        // NO
	Timestamp         *int64  `json:"timestamp"`         // YES
}

// YES 交易对
func (api *PortfolioMarginCmOrderGetApi) Symbol(symbol string) *PortfolioMarginCmOrderGetApi {
	api.req.Symbol = GetPointer(symbol)
	return api
}

// NO
func (api *PortfolioMarginCmOrderGetApi) OrderId(orderId int64) *PortfolioMarginCmOrderGetApi {
	api.req.OrderId = GetPointer(orderId)
	return api
}

// NO
func (api *PortfolioMarginCmOrderGetApi) OrigClientOrderId(origClientOrderId string) *PortfolioMarginCmOrderGetApi {
	api.req.OrigClientOrderId = GetPointer(origClientOrderId)
	return api
}

// NO
func (api *PortfolioMarginCmOrderGetApi) RecvWindow(recvWindow int64) *PortfolioMarginCmOrderGetApi {
	api.req.RecvWindow = GetPointer(recvWindow)
	return api
}

// YES
func (api *PortfolioMarginCmOrderGetApi) Timestamp(timestamp int64) *PortfolioMarginCmOrderGetApi {
	api.req.Timestamp = GetPointer(timestamp)
	return api
}

type PortfolioMarginCmOrderGetApi struct {
	client *PortfolioMarginRestClient
	req    *PortfolioMarginCmOrderGetReq
}

type PortfolioMarginCmAllOrdersReq struct {
	Symbol     *string `json:"symbol"`     // YES 交易对
	Pair       *string `json:"pair"`       // NO
	OrderId    *int64  `json:"orderId"`    // NO
	StartTime  *int64  `json:"startTime"`  // NO
	EndTime    *int64  `json:"endTime"`    // NO
	Limit      *int64  `json:"limit"`      // NO 默认值: 50; 最大值: 100
	RecvWindow *int64  `json:"recvWindow"` // NO
	Timestamp  *int64  `json:"timestamp"`  // YES
}

// YES 交易对
func (api *PortfolioMarginCmAllOrdersApi) Symbol(symbol string) *PortfolioMarginCmAllOrdersApi {
	api.req.Symbol = GetPointer(symbol)
	return api
}

// NO
func (api *PortfolioMarginCmAllOrdersApi) Pair(pair string) *PortfolioMarginCmAllOrdersApi {
	api.req.Pair = GetPointer(pair)
	return api
}

// NO
func (api *PortfolioMarginCmAllOrdersApi) OrderId(orderId int64) *PortfolioMarginCmAllOrdersApi {
	api.req.OrderId = GetPointer(orderId)
	return api
}

// NO
func (api *PortfolioMarginCmAllOrdersApi) StartTime(startTime int64) *PortfolioMarginCmAllOrdersApi {
	api.req.StartTime = GetPointer(startTime)
	return api
}

// NO
func (api *PortfolioMarginCmAllOrdersApi) EndTime(endTime int64) *PortfolioMarginCmAllOrdersApi {
	api.req.EndTime = GetPointer(endTime)
	return api
}

// NO 默认值: 50; 最大值: 100
func (api *PortfolioMarginCmAllOrdersApi) Limit(limit int64) *PortfolioMarginCmAllOrdersApi {
	api.req.Limit = GetPointer(limit)
	return api
}

// NO
func (api *PortfolioMarginCmAllOrdersApi) RecvWindow(recvWindow int64) *PortfolioMarginCmAllOrdersApi {
	api.req.RecvWindow = GetPointer(recvWindow)
	return api
}

// YES
func (api *PortfolioMarginCmAllOrdersApi) Timestamp(timestamp int64) *PortfolioMarginCmAllOrdersApi {
	api.req.Timestamp = GetPointer(timestamp)
	return api
}

type PortfolioMarginCmAllOrdersApi struct {
	client *PortfolioMarginRestClient
	req    *PortfolioMarginCmAllOrdersReq
}

type PortfolioMarginCmConditionalOpenOrderReq struct {
	Symbol              *string `json:"symbol"`              // YES 交易对
	StrategyId          *int64  `json:"strategyId"`          // YES 用户
	NewClientStrategyId *string `json:"newClientStrategyId"` // NO 用户自定义的订单号，不可以重复出现在挂单中。如空缺系统会自动赋值。必须满足正则规则: ^[\.A-Z\:/a-z0-9_-]{1,32}$
	RecvWindow          *int64  `json:"recvWindow"`          // NO
	Timestamp           *int64  `json:"timestamp"`           // YES
}

// YES 交易对
func (api *PortfolioMarginCmConditionalOpenOrderApi) Symbol(symbol string) *PortfolioMarginCmConditionalOpenOrderApi {
	api.req.Symbol = GetPointer(symbol)
	return api
}

// YES 用户
func (api *PortfolioMarginCmConditionalOpenOrderApi) StrategyId(strategyId int64) *PortfolioMarginCmConditionalOpenOrderApi {
	api.req.StrategyId = GetPointer(strategyId)
	return api
}

// NO 用户自定义的订单号，不可以重复出现在挂单中。如空缺系统会自动赋值。必须满足正则规则: ^[\.A-Z\:/a-z0-9_-]{1,32}$
func (api *PortfolioMarginCmConditionalOpenOrderApi) NewClientStrategyId(newClientStrategyId string) *PortfolioMarginCmConditionalOpenOrderApi {
	api.req.NewClientStrategyId = GetPointer(newClientStrategyId)
	return api
}

// NO
func (api *PortfolioMarginCmConditionalOpenOrderApi) RecvWindow(recvWindow int64) *PortfolioMarginCmConditionalOpenOrderApi {
	api.req.RecvWindow = &recvWindow
	return api
}

// YES
func (api *PortfolioMarginCmConditionalOpenOrderApi) Timestamp(timestamp int64) *PortfolioMarginCmConditionalOpenOrderApi {
	api.req.Timestamp = GetPointer(timestamp)
	return api
}

type PortfolioMarginCmConditionalOpenOrderApi struct {
	client *PortfolioMarginRestClient
	req    *PortfolioMarginCmConditionalOpenOrderReq
}
