package mybinanceapi

type SpotOrderOrder struct {
	Symbol                  string `json:"symbol"`                  // 交易对
	OrderId                 int64  `json:"orderId"`                 // 系统的订单ID
	OrderListId             int64  `json:"orderListId"`             // OCO订单ID，否则为 -1
	ClientOrderId           string `json:"clientOrderId"`           // 客户自己设置的ID
	Price                   string `json:"price"`                   // 订单价格
	OrigQty                 string `json:"origQty"`                 // 用户设置的原始订单数量
	ExecutedQty             string `json:"executedQty"`             // 交易的订单数量
	CummulativeQuoteQty     string `json:"cummulativeQuoteQty"`     // 累计交易的金额
	Status                  string `json:"status"`                  // 订单状态
	TimeInForce             string `json:"timeInForce"`             // 订单的时效方式
	Type                    string `json:"type"`                    // 订单类型， 比如市价单，现价单等
	Side                    string `json:"side"`                    // 订单方向，买还是卖
	StopPrice               string `json:"stopPrice"`               // 触发价，对追踪止损单无效
	IcebergQty              string `json:"icebergQty"`              // 冰山订单的数量
	Time                    int64  `json:"time"`                    // 订单添加到 order book 的时间
	UpdateTime              int64  `json:"updateTime"`              // 订单最后一次成交的时间
	IsWorking               bool   `json:"isWorking"`               // 订单是否在 order book 中被激活
	WorkingTime             int64  `json:"workingTime"`             // 订单添加到 order book 的时间
	OrigQuoteOrderQty       string `json:"origQuoteOrderQty"`       // 原始的报价订单数量
	SelfTradePreventionMode string `json:"selfTradePreventionMode"` // 自成交防范模式
}
type SpotOpenOrdersRes []SpotOrderOrder

type SpotAllOrdersRes []SpotOrderOrder

type SpotOrderGetRes SpotOrderOrder

type SpotOrderFill struct {
	Price           string `json:"price"`           // 交易的价格
	Qty             string `json:"qty"`             // 交易的数量
	Commission      string `json:"commission"`      // 手续费金额
	CommissionAsset string `json:"commissionAsset"` // 手续费的币种
	TradeId         int64  `json:"tradeId"`         // 交易ID
}
type SpotOrderPostRes struct {
	Symbol                  string          `json:"symbol"`                  // 交易对
	OrderId                 int64           `json:"orderId"`                 // 系统的订单ID
	OrderListId             int64           `json:"orderListId"`             // OCO订单ID，否则为 -1
	ClientOrderId           string          `json:"clientOrderId"`           // 客户自己设置的ID
	TransactTime            int64           `json:"transactTime"`            // 交易的时间戳
	Price                   string          `json:"price"`                   // 订单价格
	OrigQty                 string          `json:"origQty"`                 // 用户设置的原始订单数量
	ExecutedQty             string          `json:"executedQty"`             // 交易的订单数量
	CummulativeQuoteQty     string          `json:"cummulativeQuoteQty"`     // 累计交易的金额
	Status                  string          `json:"status"`                  // 订单状态
	TimeInForce             string          `json:"timeInForce"`             // 订单的时效方式
	Type                    string          `json:"type"`                    // 订单类型， 比如市价单，现价单等
	Side                    string          `json:"side"`                    // 订单方向，买还是卖
	WorkingTime             int64           `json:"workingTime"`             // 订单添加到 order book 的时间
	SelfTradePreventionMode string          `json:"selfTradePreventionMode"` // 自我交易预防模式
	Fills                   []SpotOrderFill `json:"fills"`                   // 订单中交易的信息
}

type SpotOrderDeleteRes struct {
	Symbol                  string `json:"symbol"`                  // 交易对
	OrigClientOrderId       string `json:"origClientOrderId"`       // 原始的客户端订单ID
	OrderId                 int64  `json:"orderId"`                 // 系统的订单ID
	OrderListId             int64  `json:"orderListId"`             // OCO订单ID，否则为 -1
	ClientOrderId           string `json:"clientOrderId"`           // 客户自己设置的ID
	TransactTime            int64  `json:"transactTime"`            // 交易的时间戳
	Price                   string `json:"price"`                   // 订单价格
	OrigQty                 string `json:"origQty"`                 // 用户设置的原始订单数量
	ExecutedQty             string `json:"executedQty"`             // 交易的订单数量
	CummulativeQuoteQty     string `json:"cummulativeQuoteQty"`     // 累计交易的金额
	Status                  string `json:"status"`                  // 订单状态
	TimeInForce             string `json:"timeInForce"`             // 订单的时效方式
	Type                    string `json:"type"`                    // 订单类型， 比如市价单，现价单等
	Side                    string `json:"side"`                    // 订单方向，买还是卖
	SelfTradePreventionMode string `json:"selfTradePreventionMode"` // 自我交易预防模式
}

type SpotOrderCancelReplaceRes struct {
	CancelResult     string             `json:"cancelResult"`
	NewOrderResult   string             `json:"newOrderResult"`
	CancelResponse   SpotOrderDeleteRes `json:"cancelResponse"`
	NewOrderResponse SpotOrderPostRes   `json:"newOrderResponse"`
}

type SpotMyTradesResRow struct {
	Symbol          string `json:"symbol"`          // 交易对
	Id              int64  `json:"id"`              // trade ID
	OrderId         int64  `json:"orderId"`         // 订单ID
	OrderListId     int64  `json:"orderListId"`     // OCO订单的ID，不然就是-1
	Price           string `json:"price"`           // 成交价格
	Qty             string `json:"qty"`             // 成交量
	QuoteQty        string `json:"quoteQty"`        // 成交金额
	Commission      string `json:"commission"`      // 交易费金额
	CommissionAsset string `json:"commissionAsset"` // 交易费资产类型
	Time            int64  `json:"time"`            // 交易时间
	IsBuyer         bool   `json:"isBuyer"`         // 是否是买家
	IsMaker         bool   `json:"isMaker"`         // 是否是挂单方
	IsBestMatch     bool   `json:"isBestMatch"`
}
type SpotMyTradesRes []SpotMyTradesResRow
