package mybinanceapi

type SwapOrderOrder struct {
	AvgPrice      string `json:"avgPrice"`      // 平均成交价
	ClientOrderId string `json:"clientOrderId"` // 用户自定义的订单号
	CumQuote      string `json:"cumQuote"`      // 成交金额
	ExecutedQty   string `json:"executedQty"`   // 成交量
	OrderId       int64  `json:"orderId"`       // 系统订单号
	OrigQty       string `json:"origQty"`       // 原始委托数量
	OrigType      string `json:"origType"`      // 触发前订单类型
	Price         string `json:"price"`         // 委托价格
	ReduceOnly    bool   `json:"reduceOnly"`    // 是否仅减仓
	Side          string `json:"side"`          // 买卖方向
	PositionSide  string `json:"positionSide"`  // 持仓方向
	Status        string `json:"status"`        // 订单状态
	StopPrice     string `json:"stopPrice"`     // 触发价，对`TRAILING_STOP_MARKET`无效
	ClosePosition bool   `json:"closePosition"` // 是否条件全平仓
	Symbol        string `json:"symbol"`        // 交易对
	Time          int64  `json:"time"`          // 订单时间
	TimeInForce   string `json:"timeInForce"`   // 有效方法
	Type          string `json:"type"`          // 订单类型
	ActivatePrice string `json:"activatePrice"` // 跟踪止损激活价格, 仅`TRAILING_STOP_MARKET` 订单返回此字段
	PriceRate     string `json:"priceRate"`     // 跟踪止损回调比例, 仅`TRAILING_STOP_MARKET` 订单返回此字段
	UpdateTime    int64  `json:"updateTime"`    // 更新时间
	WorkingType   string `json:"workingType"`   // 条件价格触发类型
	PriceProtect  bool   `json:"priceProtect"`  // 是否开启条件单触发保护
}
type SwapOpenOrdersRes []SwapOrderOrder

type SwapAllOrdersRes []SwapOrderOrder

type SwapOrderPostRes struct {
	ClientOrderId string `json:"clientOrderId"` // 用户自定义的订单号
	CumQty        string `json:"cumQty"`
	CumQuote      string `json:"cumQuote"` // 成交金额
	ExecutedQty   string `json:"executedQty"`
	OrderId       int64  `json:"orderId"` // 系统订单号
	AvgPrice      string `json:"avgPrice"`
	OrigQty       string `json:"origQty"` // 原始委托数量
	Price         string `json:"price"`   // 委托价格
	ReduceOnly    bool   `json:"reduceOnly"`
	Side          string `json:"side"`         // 买卖方向
	PositionSide  string `json:"positionSide"` // 持仓方向
	Status        string `json:"status"`       // 订单状态
	StopPrice     string `json:"stopPrice"`    // 触发价，对`TRAILING_STOP_MARKET`无效
	ClosePosition bool   `json:"closePosition"`
	Symbol        string `json:"symbol"`      // 交易对
	TimeInForce   string `json:"timeInForce"` // 有效方法
	Type          string `json:"type"`        // 订单类型
	OrigType      string `json:"origType"`    // 触发前订单类型
}

type SwapOrderPutRes SwapOrderPostRes

type SwapOrderGetRes SwapOrderOrder

type SwapOrderDeleteRes struct {
	BinanceErrorRes
	SwapOrderOrder
}

type SwapUserTradesRes []SwapUserTrade
type SwapUserTrade struct {
	Symbol          string `json:"symbol"`          // 交易对
	Id              int64  `json:"id"`              // 交易ID
	OrderId         int64  `json:"orderId"`         // 订单ID
	Pair            string `json:"pair"`            // 标的交易对
	Side            string `json:"side"`            // 买卖方向
	Price           string `json:"price"`           // 成交价
	Qty             string `json:"qty"`             // 成交量(张数)
	RealizedPnl     string `json:"realizedPnl"`     // 实现盈亏
	MarginAsset     string `json:"marginAsset"`     // 保证金币种
	BaseQty         string `json:"baseQty"`         // 成交额(标的数量)
	Commission      string `json:"commission"`      // 手续费
	CommissionAsset string `json:"commissionAsset"` // 手续费单位
	Time            int64  `json:"time"`            // 时间
	PositionSide    string `json:"positionSide"`    // 持仓方向
	Buyer           bool   `json:"buyer"`           // 是否是买方
	Maker           bool   `json:"maker"`           // 是否是挂单方
}

type SwapBatchOrdersPostRes []SwapOrderPostResRow
type SwapOrderPostResRow struct {
	BinanceErrorRes
	SwapOrderPostRes
}

type SwapBatchOrdersPutRes []SwapBatchOrdersPutResRow
type SwapBatchOrdersPutResRow struct {
	BinanceErrorRes
	SwapOrderPutRes
}

type SwapBatchOrdersDeleteRes []SwapOrderDeleteRes

type SwapCommissionRateRes struct {
	Symbol              string `json:"symbol"`
	MakerCommissionRate string `json:"makerCommissionRate"`
	TakerCommissionRate string `json:"takerCommissionRate"`
}
