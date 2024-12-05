package mybinanceapi

type PortfolioMarginUmOrderPostRes struct {
	ClientOrderId           string `json:"clientOrderId"`           // 用户自定义的订单号，不可以重复出现在挂单中。如空缺系统会自动赋值。必须满足正则规则: ^[\.A-Z\:/a-z0-9_-]{1,32}$
	CumQty                  string `json:"cumQty"`                  // 成交量
	CumQuote                string `json:"cumQuote"`                // 成交金额
	ExecutedQty             string `json:"executedQty"`             // 成交数量
	OrderId                 int64  `json:"orderId"`                 // 订单ID
	AvgPrice                string `json:"avgPrice"`                // 平均成交价
	OrigQty                 string `json:"origQty"`                 // 原始数量
	Price                   string `json:"price"`                   // 价格
	ReduceOnly              bool   `json:"reduceOnly"`              // true或false; 非双开模式下默认false；双开模式下不接受此参数
	Side                    string `json:"side"`                    // 方向
	PositionSide            string `json:"positionSide"`            // 持仓方向
	Status                  string `json:"status"`                  // 订单状态
	Symbol                  string `json:"symbol"`                  // 交易对
	TimeInForce             string `json:"timeInForce"`             // TIF
	Type                    string `json:"type"`                    // 订单类型
	SelfTradePreventionMode string `json:"selfTradePreventionMode"` // 订单自成交保护模式
	GoodTillDate            int64  `json:"goodTillDate"`            // TIF为GTD时的自动取消时间
	UpdateTime              int64  `json:"updateTime"`              // 更新时间
	PriceMatch              string `json:"priceMatch"`              // OPPONENT/ OPPONENT_5/ OPPONENT_10/ OPPONENT_20/QUEUE/ QUEUE_5/ QUEUE_10/ QUEUE_20；不能与price同时传
}

type PortfolioMarginUmOrderPutRes struct {
	OrderId                 int64  `json:"orderId"`                 // 订单ID
	Symbol                  string `json:"symbol"`                  // 交易对
	Status                  string `json:"status"`                  // 订单状态
	ClientOrderId           string `json:"clientOrderId"`           // 用户自定义的订单号，不可以重复出现在挂单中。如空缺系统会自动赋值。必须满足正则规则: ^[\.A-Z\:/a-z0-9_-]{1,32}$
	Price                   string `json:"price"`                   // 价格
	AvgPrice                string `json:"avgPrice"`                // 平均成交价
	OrigQty                 string `json:"origQty"`                 // 原始数量
	ExecutedQty             string `json:"executedQty"`             // 成交数量
	CumQty                  string `json:"cumQty"`                  // 成交量
	CumQuote                string `json:"cumQuote"`                // 成交金额
	TimeInForce             string `json:"timeInForce"`             // TIF
	Type                    string `json:"type"`                    // 订单类型
	ReduceOnly              bool   `json:"reduceOnly"`              // true或false; 非双开模式下默认false；双开模式下不接受此参数
	Side                    string `json:"side"`                    // 方向
	PositionSide            string `json:"positionSide"`            // 持仓方向
	OrigType                string `json:"origType"`                // 订单类型
	SelfTradePreventionMode string `json:"selfTradePreventionMode"` // 订单自成交保护模式
	GoodTillDate            int64  `json:"goodTillDate"`            // TIF为GTD时的自动取消时间
	UpdateTime              int64  `json:"updateTime"`              // 更新时间
	PriceMatch              string `json:"priceMatch"`              // OPPONENT/ OPPONENT_5/ OPPONENT_10/ OPPONENT_20/QUEUE/ QUEUE_5/ QUEUE_10/ QUEUE_20；不能与price同时传
}

type PortfolioMarginUmOrderDeleteRes struct {
	AvgPrice                string `json:"avgPrice"`                // 平均成交价
	ClientOrderId           string `json:"clientOrderId"`           // 用户自定义的订单号，不可以重复出现在挂单中。如空缺系统会自动赋值。必须满足正则规则: ^[\.A-Z\:/a-z0-9_-]{1,32}$
	CumQty                  string `json:"cumQty"`                  // 成交量
	CumQuote                string `json:"cumQuote"`                // 成交金额
	ExecutedQty             string `json:"executedQty"`             // 成交数量
	OrderId                 int64  `json:"orderId"`                 // 订单ID
	OrigQty                 string `json:"origQty"`                 // 原始数量
	Price                   string `json:"price"`                   // 价格
	ReduceOnly              bool   `json:"reduceOnly"`              // true或false; 非双开模式下默认false；双开模式下不接受此参数
	Side                    string `json:"side"`                    // 方向
	PositionSide            string `json:"positionSide"`            // 持仓方向
	Status                  string `json:"status"`                  // 订单状态
	Symbol                  string `json:"symbol"`                  // 交易对
	TimeInForce             string `json:"timeInForce"`             // TIF
	Type                    string `json:"type"`                    // 订单类型
	UpdateTime              int64  `json:"updateTime"`              // 更新时间
	SelfTradePreventionMode string `json:"selfTradePreventionMode"` // 订单自成交保护模式
	GoodTillDate            int64  `json:"goodTillDate"`            // TIF为GTD时的自动取消时间
	PriceMatch              string `json:"priceMatch"`              // OPPONENT/ OPPONENT_5/ OPPONENT_10/ OPPONENT_20/QUEUE/ QUEUE_5/ QUEUE_10/ QUEUE_20；不能与price同时传
}

type PortfolioMarginUmAllOpenOrdersDeleteRes struct{}

type PortfolioMarginUmOrderGetRes struct {
	AvgPrice                string `json:"avgPrice"`                // 平均成交价
	ClientOrderId           string `json:"clientOrderId"`           // 用户自定义的订单号，不可以重复出现在挂单中。如空缺系统会自动赋值。必须满足正则规则: ^[\.A-Z\:/a-z0-9_-]{1,32}$
	CumQuote                string `json:"cumQuote"`                // 成交金额
	ExecutedQty             string `json:"executedQty"`             // 成交数量
	OrderId                 int64  `json:"orderId"`                 // 订单ID
	OrigQty                 string `json:"origQty"`                 // 原始数量
	OrigType                string `json:"origType"`                // 订单类型
	Price                   string `json:"price"`                   // 价格
	ReduceOnly              bool   `json:"reduceOnly"`              // true或false; 非双开模式下默认false；双开模式下不接受此参数
	Side                    string `json:"side"`                    // 方向
	PositionSide            string `json:"positionSide"`            // 持仓方向
	Status                  string `json:"status"`                  // 订单状态
	Symbol                  string `json:"symbol"`                  // 交易对
	Time                    int64  `json:"time"`                    // 时间
	TimeInForce             string `json:"timeInForce"`             // TIF
	Type                    string `json:"type"`                    // 订单类型
	UpdateTime              int64  `json:"updateTime"`              // 更新时间
	SelfTradePreventionMode string `json:"selfTradePreventionMode"` // 订单自成交保护模式
	GoodTillDate            int64  `json:"goodTillDate"`            // TIF为GTD时的自动取消时间
	PriceMatch              string `json:"priceMatch"`              // OPPONENT/ OPPONENT_5/ OPPONENT_10/ OPPONENT_20/QUEUE/ QUEUE_5/ QUEUE_10/ QUEUE_20；不能与price同时传
}

type PortfolioMarginUmConditionalOrderPostRes struct {
	NewClientStrategyId     string `json:"newClientStrategyId"`     // 用户自定义的订单号，不可以重复出现在挂单中。如空缺系统会自动赋值。必须满足正则规则: ^[\.A-Z\:/a-z0-9_-]{1,32}$
	StrategyId              int64  `json:"strategyId"`              // 策略ID
	StrategyStatus          string `json:"strategyStatus"`          // 策略状态
	StrategyType            string `json:"strategyType"`            // 条件单类型"STOP", "STOP_MARKET", "TAKE_PROFIT", "TAKE_PROFIT_MARKET"或"TRAILING_STOP_MARKET"
	OrigQty                 string `json:"origQty"`                 // 原始数量
	Price                   string `json:"price"`                   // 价格
	ReduceOnly              bool   `json:"reduceOnly"`              // true或false; 非双开模式下默认false；双开模式下不接受此参数
	Side                    string `json:"side"`                    // 方向
	PositionSide            string `json:"positionSide"`            // 持仓方向
	StopPrice               string `json:"stopPrice"`               // 触发价格
	Symbol                  string `json:"symbol"`                  // 交易对
	TimeInForce             string `json:"timeInForce"`             // TIF
	ActivatePrice           string `json:"activatePrice"`           // 触发价格
	PriceRate               string `json:"priceRate"`               // 触发百分比
	BookTime                int64  `json:"bookTime"`                // 条件单下单时间
	UpdateTime              int64  `json:"updateTime"`              // 更新时间
	WorkingType             string `json:"workingType"`             // 触发类型
	PriceProtect            bool   `json:"priceProtect"`            // 条件单触发保护
	SelfTradePreventionMode string `json:"selfTradePreventionMode"` // 订单自成交保护模式
	GoodTillDate            int64  `json:"goodTillDate"`            // TIF为GTD时的自动取消时间
	PriceMatch              string `json:"priceMatch"`              // OPPONENT/ OPPONENT_5/ OPPONENT_10/ OPPONENT_20/QUEUE/ QUEUE_5/ QUEUE_10/ QUEUE_20；不能与price同时传
}

type PortfolioMarginUmConditionalOrderDeleteRes struct {
	NewClientStrategyId     string `json:"newClientStrategyId"`     // 用户自定义的订单号，不可以重复出现在挂单中。如空缺系统会自动赋值。必须满足正则规则: ^[\.A-Z\:/a-z0-9_-]{1,32}$
	StrategyId              int64  `json:"strategyId"`              // 策略ID
	StrategyStatus          string `json:"strategyStatus"`          // 策略状态
	StrategyType            string `json:"strategyType"`            // 条件单类型"STOP", "STOP_MARKET", "TAKE_PROFIT", "TAKE_PROFIT_MARKET"或"TRAILING_STOP_MARKET"
	OrigQty                 string `json:"origQty"`                 // 原始数量
	Price                   string `json:"price"`                   // 价格
	ReduceOnly              bool   `json:"reduceOnly"`              // true或false; 非双开模式下默认false；双开模式下不接受此参数
	Side                    string `json:"side"`                    // 方向
	PositionSide            string `json:"positionSide"`            // 持仓方向
	StopPrice               string `json:"stopPrice"`               // 触发价格
	Symbol                  string `json:"symbol"`                  // 交易对
	TimeInForce             string `json:"timeInForce"`             // TIF
	ActivatePrice           string `json:"activatePrice"`           // 触发价格
	PriceRate               string `json:"priceRate"`               // 触发百分比
	BookTime                int64  `json:"bookTime"`                // 条件单下单时间
	UpdateTime              int64  `json:"updateTime"`              // 更新时间
	WorkingType             string `json:"workingType"`             // 触发类型
	PriceProtect            bool   `json:"priceProtect"`            // 条件单触发保护
	SelfTradePreventionMode string `json:"selfTradePreventionMode"` // 订单自成交保护模式
	GoodTillDate            int64  `json:"goodTillDate"`            // TIF为GTD时的自动取消时间
	PriceMatch              string `json:"priceMatch"`              // OPPONENT/ OPPONENT_5/ OPPONENT_10/ OPPONENT_20/QUEUE/ QUEUE_5/ QUEUE_10/ QUEUE_20；不能与price同时传
}

type PortfolioMarginUmOpenOrderGetRes struct {
	AvgPrice                string `json:"avgPrice"`                // 平均成交价
	ClientOrderId           string `json:"clientOrderId"`           // 用户自定义的订单号，不可以重复出现在挂单中。如空缺系统会自动赋值。必须满足正则规则: ^[\.A-Z\:/a-z0-9_-]{1,32}$
	CumQuote                string `json:"cumQuote"`                // 成交金额
	ExecutedQty             string `json:"executedQty"`             // 成交数量
	OrderId                 int64  `json:"orderId"`                 // 订单ID
	OrigQty                 string `json:"origQty"`                 // 原始数量
	OrigType                string `json:"origType"`                // 订单类型
	Price                   string `json:"price"`                   // 价格
	ReduceOnly              bool   `json:"reduceOnly"`              // true或false; 非双开模式下默认false；双开模式下不接受此参数
	Side                    string `json:"side"`                    // 方向
	PositionSide            string `json:"positionSide"`            // 持仓方向
	Status                  string `json:"status"`                  // 订单状态
	Symbol                  string `json:"symbol"`                  // 交易对
	Time                    int64  `json:"time"`                    // 时间
	TimeInForce             string `json:"timeInForce"`             // TIF
	Type                    string `json:"type"`                    // 订单类型
	UpdateTime              int64  `json:"updateTime"`              // 更新时间
	SelfTradePreventionMode string `json:"selfTradePreventionMode"` // 订单自成交保护模式
	GoodTillDate            int64  `json:"goodTillDate"`            // TIF为GTD时的自动取消时间
	PriceMatch              string `json:"priceMatch"`              // OPPONENT/ OPPONENT_5/ OPPONENT_10/ OPPONENT_20/QUEUE/ QUEUE_5/ QUEUE_10/ QUEUE_20；不能与price同时传
}

type PortfolioMarginUmOpenOrdersGetResRow struct {
	AvgPrice                string `json:"avgPrice"`                // 平均成交价
	ClientOrderId           string `json:"clientOrderId"`           // 用户自定义的订单号，不可以重复出现在挂单中。如空缺系统会自动赋值。必须满足正则规则: ^[\.A-Z\:/a-z0-9_-]{1,32}$
	CumQuote                string `json:"cumQuote"`                // 成交金额
	ExecutedQty             string `json:"executedQty"`             // 成交数量
	OrderId                 int64  `json:"orderId"`                 // 订单ID
	OrigQty                 string `json:"origQty"`                 // 原始数量
	OrigType                string `json:"origType"`                // 订单类型
	Price                   string `json:"price"`                   // 价格
	ReduceOnly              bool   `json:"reduceOnly"`              // true或false; 非双开模式下默认false；双开模式下不接受此参数
	Side                    string `json:"side"`                    // 方向
	PositionSide            string `json:"positionSide"`            // 持仓方向
	Status                  string `json:"status"`                  // 订单状态
	Symbol                  string `json:"symbol"`                  // 交易对
	Time                    int64  `json:"time"`                    // 时间
	TimeInForce             string `json:"timeInForce"`             // TIF
	Type                    string `json:"type"`                    // 订单类型
	UpdateTime              int64  `json:"updateTime"`              // 更新时间
	SelfTradePreventionMode string `json:"selfTradePreventionMode"` // 订单自成交保护模式
	GoodTillDate            int64  `json:"goodTillDate"`            // TIF为GTD时的自动取消时间
	PriceMatch              string `json:"priceMatch"`              // OPPONENT/ OPPONENT_5/ OPPONENT_10/ OPPONENT_20/QUEUE/ QUEUE_5/ QUEUE_10/ QUEUE_20；不能与price同时传
}
type PortfolioMarginUmOpenOrdersGetRes []PortfolioMarginUmOpenOrdersGetResRow

type PortfolioMarginCmOrderPostRes struct {
	ClientOrderId string `json:"clientOrderId"` // 用户自定义的订单号，不可以重复出现在挂单中。如空缺系统会自动赋值。必须满足正则规则: ^[\.A-Z\:/a-z0-9_-]{1,32}$
	CumQty        string `json:"cumQty"`        // 成交量
	CumBase       string `json:"cumBase"`       // 成交基础币量
	ExecutedQty   string `json:"executedQty"`   // 成交数量
	OrderId       int64  `json:"orderId"`       // 订单ID
	AvgPrice      string `json:"avgPrice"`      // 平均成交价
	OrigQty       string `json:"origQty"`       // 原始数量
	Price         string `json:"price"`         // 价格
	ReduceOnly    bool   `json:"reduceOnly"`    // true或false; 非双开模式下默认false；双开模式下不接受此参数
	Side          string `json:"side"`          // 方向
	PositionSide  string `json:"positionSide"`  // 持仓方向
	Status        string `json:"status"`        // 订单状态
	Symbol        string `json:"symbol"`        // 交易对
	Pair          string `json:"pair"`          // 交易对
	TimeInForce   string `json:"timeInForce"`   // TIF
	Type          string `json:"type"`          // 订单类型
	UpdateTime    int64  `json:"updateTime"`    // 更新时间
}

type PortfolioMarginCmConditionalOrderPostRes struct {
	NewClientStrategyId string `json:"newClientStrategyId"` // 用户自定义的订单号，不可以重复出现在挂单中。如空缺系统会自动赋值。必须满足正则规则: ^[\.A-Z\:/a-z0-9_-]{1,32}$
	StrategyId          int64  `json:"strategyId"`          // 策略ID
	StrategyStatus      string `json:"strategyStatus"`      // 策略状态
	StrategyType        string `json:"strategyType"`        // 条件单类型"STOP", "STOP_MARKET", "TAKE_PROFIT", "TAKE_PROFIT_MARKET"或"TRAILING_STOP_MARKET"
	OrigQty             string `json:"origQty"`             // 原始数量
	Price               string `json:"price"`               // 价格
	ReduceOnly          bool   `json:"reduceOnly"`          // true或false; 非双开模式下默认false；双开模式下不接受此参数
	Side                string `json:"side"`                // 方向
	PositionSide        string `json:"positionSide"`        // 持仓方向
	StopPrice           string `json:"stopPrice"`           // 触发价格
	Symbol              string `json:"symbol"`              // 交易对
	Pair                string `json:"pair"`                // 交易对
	TimeInForce         string `json:"timeInForce"`         // TIF
	ActivatePrice       string `json:"activatePrice"`       // 触发价格
	PriceRate           string `json:"priceRate"`           // 触发百分比
	BookTime            int64  `json:"bookTime"`            // 条件单下单时间
	UpdateTime          int64  `json:"updateTime"`          // 更新时间
	WorkingType         string `json:"workingType"`         // 触发类型
	PriceProtect        bool   `json:"priceProtect"`        // 条件单触发保护
}

type PortfolioMarginMarginOrderPostRes struct {
	Symbol                string `json:"symbol"`                // 交易对
	OrderId               int64  `json:"orderId"`               // 订单ID
	ClientOrderId         string `json:"clientOrderId"`         // 用户自定义的订单号
	TransactTime          int64  `json:"transactTime"`          // 成交时间
	Price                 string `json:"price"`                 // 价格
	OrigQty               string `json:"origQty"`               // 原始数量
	ExecutedQty           string `json:"executedQty"`           // 成交数量
	CummulativeQuoteQty   string `json:"cummulativeQuoteQty"`   // 成交金额
	Status                string `json:"status"`                // 订单状态
	TimeInForce           string `json:"timeInForce"`           // TIF
	Type                  string `json:"type"`                  // 订单类型
	Side                  string `json:"side"`                  // 方向
	MarginBuyBorrowAmount int    `json:"marginBuyBorrowAmount"` // 下单后没有发生借款则不返回该字段
	MarginBuyBorrowAsset  string `json:"marginBuyBorrowAsset"`  // 下单后没有发生借款则不返回该字段
	Fills                 []struct {
		Price           string `json:"price"`           // 成交价格
		Qty             string `json:"qty"`             // 成交数量
		Commission      string `json:"commission"`      // 佣金
		CommissionAsset string `json:"commissionAsset"` // 佣金资产
	} `json:"fills"` // 成交明细
}

type PortfolioMarginMarginOrderOcoPostRes struct {
	OrderListId           int    `json:"orderListId"`           // 订单列表ID
	ContingencyType       string `json:"contingencyType"`       // 条件类型
	ListStatusType        string `json:"listStatusType"`        // 订单列表状态
	ListOrderStatus       string `json:"listOrderStatus"`       // 订单列表状态
	ListClientOrderId     string `json:"listClientOrderId"`     // 用户自定义的订单号
	TransactionTime       int64  `json:"transactionTime"`       // 交易时间
	Symbol                string `json:"symbol"`                // 交易对
	MarginBuyBorrowAmount int    `json:"marginBuyBorrowAmount"` // 下单后没有发生借款则不返回该字段
	MarginBuyBorrowAsset  string `json:"marginBuyBorrowAsset"`  // 下单后没有发生借款则不返回该字段
	Orders                []struct {
		Symbol        string `json:"symbol"`        // 交易对
		OrderId       int64  `json:"orderId"`       // 订单ID
		ClientOrderId string `json:"clientOrderId"` // 用户自定义的订单号
	} `json:"orders"` // 订单列表
	OrderReports []struct {
		Symbol              string `json:"symbol"`              // 交易对
		OrderId             int64  `json:"orderId"`             // 订单ID
		OrderListId         int    `json:"orderListId"`         // 订单列表ID
		ClientOrderId       string `json:"clientOrderId"`       // 用户自定义的订单号
		TransactTime        int64  `json:"transactTime"`        // 交易时间
		Price               string `json:"price"`               // 价格
		OrigQty             string `json:"origQty"`             // 原始数量
		EcecutedQty         string `json:"executedQty"`         // 成交数量
		CummulativeQuoteQty string `json:"cummulativeQuoteQty"` // 成交金额
		Status              string `json:"status"`              // 订单状态
		TimeInForce         string `json:"timeInForce"`         // TIF
		Type                string `json:"type"`                // 订单类型
		Side                string `json:"side"`                // 方向
		StopPrice           string `json:"stopPrice"`           // 触发价格
	} `json:"orderReports"` // 订单列表
}

type PortfolioMarginUmConditionalOpenOrderRes struct {
	NewClientStrategyId     string `json:"newClientStrategyId"`     // 用户自定义的订单号，不可以重复出现在挂单中。如空缺系统会自动赋值。必须满足正则规则: ^[\.A-Z\:/a-z0-9_-]{1,32}$
	StrategyId              int64  `json:"strategyId"`              // 策略ID
	StrategyStatus          string `json:"strategyStatus"`          // 策略状态
	StrategyType            string `json:"strategyType"`            // 条件单类型"STOP", "STOP_MARKET", "TAKE_PROFIT", "TAKE_PROFIT_MARKET"或"TRAILING_STOP_MARKET"
	OrigQty                 string `json:"origQty"`                 // 原始数量
	Price                   string `json:"price"`                   // 价格
	ReduceOnly              bool   `json:"reduceOnly"`              // true或false; 非双开模式下默认false；双开模式下不接受此参数
	Side                    string `json:"side"`                    // 方向
	PositionSide            string `json:"positionSide"`            // 持仓方向
	StopPrice               string `json:"stopPrice"`               // 触发价格
	Symbol                  string `json:"symbol"`                  // 交易对
	BookTime                int64  `json:"bookTime"`                // 条件单下单时间
	UpdateTime              int64  `json:"updateTime"`              // 更新时间
	TimeInForce             string `json:"timeInForce"`             // TIF
	ActivatePrice           string `json:"activatePrice"`           // 触发价格
	PriceRate               string `json:"priceRate"`               // 触发百分比
	SelfTradePreventionMode string `json:"selfTradePreventionMode"` // 订单自成交保护模式
	GoodTillDate            int64  `json:"goodTillDate"`            // TIF为GTD时的自动取消时间
	PriceMatch              string `json:"priceMatch"`              // OPPONENT/ OPPONENT_5/ OPPONENT_10/ OPPONENT_20/QUEUE/ QUEUE_5/ QUEUE_10/ QUEUE_20；不能与price同时传
}

type PortfolioMarginUmConditionalOpenOrdersResRow struct {
	NewClientStrategyId     string `json:"newClientStrategyId"`     // 用户自定义的订单号，不可以重复出现在挂单中。如空缺系统会自动赋值。必须满足正则规则: ^[\.A-Z\:/a-z0-9_-]{1,32}$
	StrategyId              int64  `json:"strategyId"`              // 策略ID
	StrategyStatus          string `json:"strategyStatus"`          // 策略状态
	StrategyType            string `json:"strategyType"`            // 条件单类型"STOP", "STOP_MARKET", "TAKE_PROFIT", "TAKE_PROFIT_MARKET"或"TRAILING_STOP_MARKET"
	OrigQty                 string `json:"origQty"`                 // 原始数量
	Price                   string `json:"price"`                   // 价格
	ReduceOnly              bool   `json:"reduceOnly"`              // true或false; 非双开模式下默认false；双开模式下不接受此参数
	Side                    string `json:"side"`                    // 方向
	PositionSide            string `json:"positionSide"`            // 持仓方向
	StopPrice               string `json:"stopPrice"`               // 触发价格
	Symbol                  string `json:"symbol"`                  // 交易对
	BookTime                int64  `json:"bookTime"`                // 条件单下单时间
	UpdateTime              int64  `json:"updateTime"`              // 更新时间
	TimeInForce             string `json:"timeInForce"`             // TIF
	ActivatePrice           string `json:"activatePrice"`           // 触发价格
	PriceRate               string `json:"priceRate"`               // 触发百分比
	SelfTradePreventionMode string `json:"selfTradePreventionMode"` // 订单自成交保护模式
	GoodTillDate            int64  `json:"goodTillDate"`            // TIF为GTD时的自动取消时间
	PriceMatch              string `json:"priceMatch"`              // OPPONENT/ OPPONENT_5/ OPPONENT_10/ OPPONENT_20/QUEUE/ QUEUE_5/ QUEUE_10/ QUEUE_20；不能与price同时传
}

type PortfolioMarginUmConditionalOpenOrdersRes []PortfolioMarginUmConditionalOpenOrdersResRow

type PortfolioMarginUmConditionalOrderHistoryRes struct {
	NewClientStrategyId     string `json:"newClientStrategyId"`     // 用户自定义的订单号，不可以重复出现在挂单中。如空缺系统会自动赋值。必须满足正则规则: ^[\.A-Z\:/a-z0-9_-]{1,32}$
	StrategyId              int64  `json:"strategyId"`              // 策略ID
	StrategyStatus          string `json:"strategyStatus"`          // 策略状态
	StrategyType            string `json:"strategyType"`            // 条件单类型"STOP", "STOP_MARKET", "TAKE_PROFIT", "TAKE_PROFIT_MARKET"或"TRAILING_STOP_MARKET"
	OrigQty                 string `json:"origQty"`                 // 原始数量
	Price                   string `json:"price"`                   // 价格
	ReduceOnly              bool   `json:"reduceOnly"`              // true或false; 非双开模式下默认false；双开模式下不接受此参数
	Side                    string `json:"side"`                    // 方向
	PositionSide            string `json:"positionSide"`            // 持仓方向
	StopPrice               string `json:"stopPrice"`               // 触发价格
	Symbol                  string `json:"symbol"`                  // 交易对
	OrderId                 int64  `json:"orderId"`                 // 条件单触发后普通单id，当条件单被触发后才有
	Status                  string `json:"status"`                  // 条件单触发后普通单状态, 当条件单被触发后才有
	BookTime                int64  `json:"bookTime"`                // 条件单下单时间
	UpdateTime              int64  `json:"updateTime"`              // 更新时间
	TriggerTime             int64  `json:"triggerTime"`             // 触发时间
	TimeInForce             string `json:"timeInForce"`             // TIF
	Type                    string `json:"type"`                    // 订单类型
	ActivatePrice           string `json:"activatePrice"`           // 触发价格
	PriceRate               string `json:"priceRate"`               // 触发百分比
	WorkingType             string `json:"workingType"`             // 触发类型
	PriceProtect            bool   `json:"priceProtect"`            // 条件单触发保护
	SelfTradePreventionMode string `json:"selfTradePreventionMode"` // 订单自成交保护模式
	GoodTillDate            int64  `json:"goodTillDate"`            // TIF为GTD时的自动取消时间
}

type PortfolioMarginCmOrderDeleteRes struct {
	AvgPrice      string `json:"avgPrice"`      // 平均成交价
	ClientOrderId string `json:"clientOrderId"` // 用户自定义的订单号，不可以重复出现在挂单中。如空缺系统会自动赋值。必须满足正则规则: ^[\.A-Z\:/a-z0-9_-]{1,32}$
	CumQty        string `json:"cumQty"`        // 成交量
	CumBase       string `json:"cumBase"`       // 成交基础币量
	ExecutedQty   string `json:"executedQty"`   // 成交数量
	OrderId       int64  `json:"orderId"`       // 订单ID
	OrigQty       string `json:"origQty"`       // 原始数量
	Price         string `json:"price"`         // 价格
	ReduceOnly    bool   `json:"reduceOnly"`    // true或false; 非双开模式下默认false；双开模式下不接受此参数
	Side          string `json:"side"`          // 方向
	PositionSide  string `json:"positionSide"`  // 持仓方向
	Status        string `json:"status"`        // 订单状态
	Symbol        string `json:"symbol"`        // 交易对
	Pair          string `json:"pair"`          // 交易对
	TimeInForce   string `json:"timeInForce"`   // TIF
	Type          string `json:"type"`          // 订单类型
	UpdateTime    int64  `json:"updateTime"`    // 更新时间
}

type PortfolioMarginCmAllOpenOrdersDeleteRes struct{}

type PortfolioMarginCmOrderPutRes struct {
	OrderId       int64  `json:"orderId"`       // 订单ID
	Symbol        string `json:"symbol"`        // 交易对
	Pair          string `json:"pair"`          // 交易对
	Status        string `json:"status"`        // 订单状态
	ClientOrderId string `json:"clientOrderId"` // 用户自定义的订单号，不可以重复出现在挂单中。如空缺系统会自动赋值。必须满足正则规则: ^[\.A-Z\:/a-z0-9_-]{1,32}$
	Price         string `json:"price"`         // 价格
	AvgPrice      string `json:"avgPrice"`      // 平均成交价
	OrigQty       string `json:"origQty"`       // 原始数量
	ExecutedQty   string `json:"executedQty"`   // 成交数量
	CumQty        string `json:"cumQty"`        // 成交量
	CumBase       string `json:"cumBase"`       // 成交基础币量
	TimeInForce   string `json:"timeInForce"`   // TIF
	Type          string `json:"type"`          // 订单类型
	ReduceOnly    bool   `json:"reduceOnly"`    // true或false; 非双开模式下默认false；双开模式下不接受此参数
	Side          string `json:"side"`          // 方向
	PositionSide  string `json:"positionSide"`  // 持仓方向
	OrigType      string `json:"origType"`      // 订单类型
	UpdateTime    int64  `json:"updateTime"`    // 更新时间
}

type PortfolioMarginCmConditionalOrderDeleteRes struct {
	NewClientStrategyId string `json:"newClientStrategyId"` // 用户自定义的订单号，不可以重复出现在挂单中。如空缺系统会自动赋值。必须满足正则规则: ^[\.A-Z\:/a-z0-9_-]{1,32}$
	StrategyId          int64  `json:"strategyId"`          // 策略ID
	StrategyStatus      string `json:"strategyStatus"`      // 策略状态
	StrategyType        string `json:"strategyType"`        // 条件单类型"STOP", "STOP_MARKET", "TAKE_PROFIT", "TAKE_PROFIT_MARKET"或"TRAILING_STOP_MARKET"
	OrigQty             string `json:"origQty"`             // 原始数量
	Price               string `json:"price"`               // 价格
	ReduceOnly          bool   `json:"reduceOnly"`          // true或false; 非双开模式下默认false；双开模式下不接受此参数
	Side                string `json:"side"`                // 方向
	PositionSide        string `json:"positionSide"`        // 持仓方向
	StopPrice           string `json:"stopPrice"`           // 触发价格
	Symbol              string `json:"symbol"`              // 交易对
	TimeInForce         string `json:"timeInForce"`         // TIF
	ActivatePrice       string `json:"activatePrice"`       // 触发价格
	PriceRate           string `json:"priceRate"`           // 触发百分比
	UpdateTime          int64  `json:"updateTime"`          // 更新时间
	WorkingType         string `json:"workingType"`         // 触发类型
	PriceProtect        bool   `json:"priceProtect"`        // 条件单触发保护
}

type PortfolioMarginCmConditionalAllOpenOrdersDeleteRes struct{}

type PortfolioMarginCmOrderGetRes struct {
	AvgPrice      string `json:"avgPrice"`      // 平均成交价
	ClientOrderId string `json:"clientOrderId"` // 用户自定义的订单号
	CumBase       string `json:"cumBase"`       // 成交基础币量
	ExecuteQty    string `json:"executeQty"`    // 成交数量
	OrderId       int64  `json:"orderId"`       // 订单ID
	OrigQty       string `json:"origQty"`       // 原始数量
	OrigType      string `json:"origType"`      // 订单类型
	Price         string `json:"price"`         // 价格
	ReduceOnly    bool   `json:"reduceOnly"`    // true或false; 非双开模式下默认false；双开模式下不接受此参数
	Side          string `json:"side"`          // 方向
	Status        string `json:"status"`        // 订单状态
	Symbol        string `json:"symbol"`        // 交易对
	Pair          string `json:"pair"`          // 交易对
	PositionSide  string `json:"positionSide"`  // 持仓方向
	Time          int64  `json:"time"`          // 时间
	TimeInForce   string `json:"timeInForce"`   // TIF
	Type          string `json:"type"`          // 订单类型
	UpdateTime    int64  `json:"updateTime"`    // 更新时间
}

type PortfolioMarginCmAllOrdersResRow struct {
	AvgPrice      string `json:"avgPrice"`      // 平均成交价
	ClientOrderId string `json:"clientOrderId"` // 用户自定义的订单号
	CumBase       string `json:"cumBase"`       // 成交基础币量
	ExecuteQty    string `json:"executeQty"`    // 成交数量
	OrderId       int64  `json:"orderId"`       // 订单ID
	OrigQty       string `json:"origQty"`       // 原始数量
	OrigType      string `json:"origType"`      // 订单类型
	Price         string `json:"price"`         // 价格
	ReduceOnly    bool   `json:"reduceOnly"`    // true或false; 非双开模式下默认false；双开模式下不接受此参数
	Side          string `json:"side"`          // 方向
	PositionSide  string `json:"positionSide"`  // 持仓方向
	Status        string `json:"status"`        // 订单状态
	Symbol        string `json:"symbol"`        // 交易对
	Pair          string `json:"pair"`          // 交易对
	Time          int64  `json:"time"`          // 时间
	TimeInForce   string `json:"timeInForce"`   // TIF
	Type          string `json:"type"`          // 订单类型
	UpdateTime    int64  `json:"updateTime"`    // 更新时间
}

type PortfolioMarginCmAllOrdersRes []PortfolioMarginCmAllOrdersResRow

type PortfolioMarginCmConditionalOpenOrderRes struct {
	NewClientStrategyId string `json:"newClientStrategyId"` // 用户自定义的订单号
	StrategyId          int64  `json:"strategyId"`          // 策略ID
	StrategyStatus      string `json:"strategyStatus"`      // 策略状态
	StrategyType        string `json:"strategyType"`        // 条件单类型
	OrigQty             string `json:"origQty"`             // 原始数量
	Price               string `json:"price"`               // 价格
	ReduceOnly          bool   `json:"reduceOnly"`          // true或false
	Side                string `json:"side"`                // 方向
	PositionSide        string `json:"positionSide"`        // 持仓方向
	StopPrice           string `json:"stopPrice"`           // 触发价格
	Symbol              string `json:"symbol"`              // 交易对
	BookTime            int64  `json:"bookTime"`            // 条件单下单时间
	UpdateTime          int64  `json:"updateTime"`          // 更新时间
	TimeInForce         string `json:"timeInForce"`         // TIF
	ActivatePrice       string `json:"activatePrice"`       // 触发价格
	PriceRate           string `json:"priceRate"`           // 触发百分比
}

type PortfolioMarginCmConditionalAllOrdersResRow struct {
	NewClientStrategyId string `json:"newClientStrategyId"` // 用户自定义的订单号
	StrategyId          int64  `json:"strategyId"`          // 策略ID
	StrategyStatus      string `json:"strategyStatus"`      // 策略状态
	StrategyType        string `json:"strategyType"`        // 条件单类型
	OrigQty             string `json:"origQty"`             // 原始数量
	Price               string `json:"price"`               // 价格
	ReduceOnly          bool   `json:"reduceOnly"`          // true或false
	Side                string `json:"side"`                // 方向
	PositionSide        string `json:"positionSide"`        // 持仓方向
	StopPrice           string `json:"stopPrice"`           // 触发价格
	Symbol              string `json:"symbol"`              // 交易对
	OrderId             int64  `json:"orderId"`             // 条件单触发后普通单id，当条件单被触发后才有
	Status              string `json:"status"`              // 条件单触发后普通单状态, 当条件单被触发后才有
	BookTime            int64  `json:"bookTime"`            // 条件单下单时间
	UpdateTime          int64  `json:"updateTime"`          // 更新时间
	TriggerTime         int64  `json:"triggerTime"`         // 触发时间
	TimeInForce         string `json:"timeInForce"`         // TIF
	Type                string `json:"type"`                // 订单类型
	ActivatePrice       string `json:"activatePrice"`       // 触发价格
	PriceRate           string `json:"priceRate"`           // 触发百分比
}

type PortfolioMarginCmConditionalAllOrdersRes []PortfolioMarginCmConditionalAllOrdersResRow

type PortfolioMarginCmConditionalOrderHistoryRes struct {
	NewClientStrategyId string `json:"newClientStrategyId"` // 用户自定义的订单号
	StrategyId          int64  `json:"strategyId"`          // 策略ID
	StrategyStatus      string `json:"strategyStatus"`      // 策略状态
	StrategyType        string `json:"strategyType"`        // 条件单类型
	OrigQty             string `json:"origQty"`             // 原始数量
	Price               string `json:"price"`               // 价格
	ReduceOnly          bool   `json:"reduceOnly"`          // true或false
	Side                string `json:"side"`                // 方向
	PositionSide        string `json:"positionSide"`        // 持仓方向
	StopPrice           string `json:"stopPrice"`           // 触发价格
	Symbol              string `json:"symbol"`              // 交易对
	OrderId             int64  `json:"orderId"`             // 条件单触发后普通单id，当条件单被触发后才有
	Status              string `json:"status"`              // 条件单触发后普通单状态, 当条件单被触发后才有
	BookTime            int64  `json:"bookTime"`            // 条件单下单时间
	UpdateTime          int64  `json:"updateTime"`          // 更新时间
	TriggerTime         int64  `json:"triggerTime"`         // 触发时间
	TimeInForce         string `json:"timeInForce"`         // TIF
	Type                string `json:"type"`                // 订单类型
	ActivatePrice       string `json:"activatePrice"`       // 触发价格
	PriceRate           string `json:"priceRate"`           // 触发百分比
	WorkingType         string `json:"workingType"`         // 触发类型
	PriceProtect        bool   `json:"priceProtect"`        // 条件单触发保护
	PriceMatch          string `json:"priceMatch"`          // OPPONENT/ OPPONENT_5/ OPPONENT_10/ OPPONENT_20/QUEUE/ QUEUE_5/ QUEUE_10/ QUEUE_20；不能与price同时传
}

type PortfolioMarginCmOpenOrdersResRow struct {
	AvgPrice      string `json:"avgPrice"`      // 平均成交价
	ClientOrderId string `json:"clientOrderId"` // 用户自定义的订单号
	CumBase       string `json:"cumBase"`       // 成交基础币量
	ExecutedQty   string `json:"executedQty"`   // 成交数量
	OrderId       int64  `json:"orderId"`       // 订单ID
	OrigQty       string `json:"origQty"`       // 原始数量
	OrigType      string `json:"origType"`      // 订单类型
	Price         string `json:"price"`         // 价格
	ReduceOnly    bool   `json:"reduceOnly"`    // true或false
	Side          string `json:"side"`          // 方向
	PositionSide  string `json:"positionSide"`  // 持仓方向
	Status        string `json:"status"`        // 订单状态
	Symbol        string `json:"symbol"`        // 交易对
	Pair          string `json:"pair"`          // 交易对
	Time          int64  `json:"time"`          // 时间
	TimeInForce   string `json:"timeInForce"`   // TIF
	Type          string `json:"type"`          // 订单类型
	UpdateTime    int64  `json:"updateTime"`    // 更新时间
}

type PortfolioMarginCmOpenOrdersRes []PortfolioMarginCmOpenOrdersResRow

type PortfolioMarginMarginOrderDeleteRes struct {
	Symbol              string `json:"symbol"`              // 交易对
	OrderId             int64  `json:"orderId"`             // 订单ID
	OrigClientOrderId   string `json:"origClientOrderId"`   // 原始订单号
	ClientOrderId       string `json:"clientOrderId"`       // 用户自定义的订单号
	Price               string `json:"price"`               // 价格
	OrigQty             string `json:"origQty"`             // 原始数量
	ExecutedQty         string `json:"executedQty"`         // 成交数量
	CummulativeQuoteQty string `json:"cummulativeQuoteQty"` // 成交金额
	Status              string `json:"status"`              // 订单状态
	TimeInForce         string `json:"timeInForce"`         // TIF
	Type                string `json:"type"`                // 订单类型
	Side                string `json:"side"`                // 方向
}

type PortfolioMarginMarginOrderOcoDeleteRes struct {
	OrderListId       int    `json:"orderListId"`       // 订单列表ID
	ContingencyType   string `json:"contingencyType"`   // 条件类型
	ListStatusType    string `json:"listStatusType"`    // 订单列表状态
	ListOrderStatus   string `json:"listOrderStatus"`   // 订单列表状态
	ListClientOrderId string `json:"listClientOrderId"` // 用户自定义的订单号
	TransactionTime   int64  `json:"transactionTime"`   // 交易时间
	Symbol            string `json:"symbol"`            // 交易对
	Orders            []struct {
		Symbol        string `json:"symbol"`        // 交易对
		OrderId       int64  `json:"orderId"`       // 订单ID
		ClientOrderId string `json:"clientOrderId"` // 用户自定义的订单号
	} `json:"orders"` // 订单列表
	OrderReports []struct {
		Symbol              string `json:"symbol"`              // 交易对
		OrigClientOrderId   string `json:"origClientOrderId"`   // 原始订单号
		OrderId             int64  `json:"orderId"`             // 订单ID
		OrderListId         int64  `json:"orderListId"`         // 订单列表ID
		ClientOrderId       string `json:"clientOrderId"`       // 用户自定义的订单号
		Price               string `json:"price"`               // 价格
		OrigQty             string `json:"origQty"`             // 原始数量
		CummulativeQuoteQty string `json:"cummulativeQuoteQty"` // 成交金额
		Status              string `json:"status"`              // 订单状态
		TimeInForce         string `json:"timeInForce"`         // TIF
		Type                string `json:"type"`                // 订单类型
		Side                string `json:"side"`                // 方向
		StopPrice           string `json:"stopPrice"`           // 触发价格
	} `json:"orderReports"` // 订单列表
}

type PortfolioMarginMarginOrderGetRes struct {
	ClientOrderId           string `json:"clientOrderId"`           // 用户自定义的订单号
	CummulativeQuoteQty     string `json:"cummulativeQuoteQty"`     // 成交金额
	ExecutedQty             string `json:"executedQty"`             // 成交数量
	IcebergQty              string `json:"icebergQty"`              // 冰山订单数量
	IsWorking               bool   `json:"isWorking"`               // 是否在交易
	OrderId                 int64  `json:"orderId"`                 // 订单ID
	OrigQty                 string `json:"origQty"`                 // 原始数量
	Price                   string `json:"price"`                   // 价格
	Side                    string `json:"side"`                    // 买卖方向
	Status                  string `json:"status"`                  // 订单状态
	StopPrice               string `json:"stopPrice"`               // 触发价格
	Symbol                  string `json:"symbol"`                  // 交易对
	Time                    int64  `json:"time"`                    // 时间
	TimeInForce             string `json:"timeInForce"`             // 有效方法
	Type                    string `json:"type"`                    // 订单类型
	UpdateTime              int64  `json:"updateTime"`              // 更新时间
	AccountId               int64  `json:"accountId"`               // 账户ID
	SelfTradePreventionMode string `json:"selfTradePreventionMode"` // 自成交保护模式
	PreventedMatchId        string `json:"preventedMatchId"`        // 防止成交ID
	PreventedQuantity       string `json:"preventedQuantity"`       // 防止成交数量
}

type PortfolioMarginMarginOpenOrdersResRow struct {
	ClientOrderId           string `json:"clientOrderId"`           // 用户自定义的订单号
	CummulativeQuoteQty     string `json:"cummulativeQuoteQty"`     // 成交金额
	ExecutedQty             string `json:"executedQty"`             // 成交数量
	IcebergQty              string `json:"icebergQty"`              // 冰山订单数量
	IsWorking               bool   `json:"isWorking"`               // 是否在交易
	OrderId                 int64  `json:"orderId"`                 // 订单ID
	OrigQty                 string `json:"origQty"`                 // 原始数量
	Price                   string `json:"price"`                   // 价格
	Side                    string `json:"side"`                    // 买卖方向
	Status                  string `json:"status"`                  // 订单状态
	StopPrice               string `json:"stopPrice"`               // 触发价格
	Symbol                  string `json:"symbol"`                  // 交易对
	Time                    int64  `json:"time"`                    // 时间
	TimeInForce             string `json:"timeInForce"`             // 有效方法
	Type                    string `json:"type"`                    // 订单类型
	UpdateTime              int64  `json:"updateTime"`              // 更新时间
	AccountId               int64  `json:"accountId"`               // 账户ID
	SelfTradePreventionMode string `json:"selfTradePreventionMode"` // 自成交保护模式
	PreventedMatchId        string `json:"preventedMatchId"`        // 防止成交ID
	PreventedQuantity       string `json:"preventedQuantity"`       // 防止成交数量
}

type PortfolioMarginMarginOpenOrdersRes []PortfolioMarginMarginOpenOrdersResRow

type PortfolioMarginMarginAllOrdersResRow struct {
	ClientOrderId           string `json:"clientOrderId"`           // 用户自定义的订单号
	CummulativeQuoteQty     string `json:"cummulativeQuoteQty"`     // 成交金额
	ExecutedQty             string `json:"executedQty"`             // 成交数量
	IcebergQty              string `json:"icebergQty"`              // 冰山订单数量
	IsWorking               bool   `json:"isWorking"`               // 是否在交易
	OrderId                 int64  `json:"orderId"`                 // 订单ID
	OrigQty                 string `json:"origQty"`                 // 原始数量
	Price                   string `json:"price"`                   // 价格
	Side                    string `json:"side"`                    // 买卖方向
	Status                  string `json:"status"`                  // 订单状态
	StopPrice               string `json:"stopPrice"`               // 触发价格
	Symbol                  string `json:"symbol"`                  // 交易对
	Time                    int64  `json:"time"`                    // 时间
	TimeInForce             string `json:"timeInForce"`             // 有效方法
	Type                    string `json:"type"`                    // 订单类型
	UpdateTime              int64  `json:"updateTime"`              // 更新时间
	AccountId               int64  `json:"accountId"`               // 账户ID
	SelfTradePreventionMode string `json:"selfTradePreventionMode"` // 自成交保护模式
	PreventedMatchId        string `json:"preventedMatchId"`        // 防止成交ID
	PreventedQuantity       string `json:"preventedQuantity"`       // 防止成交数量
}
type PortfolioMarginMarginAllOrdersRes []PortfolioMarginMarginAllOrdersResRow

type PortfolioMarginMarginOrderOcoGetRes struct {
	OrderListId       int    `json:"orderListId"`       // 订单列表ID
	ContingencyType   string `json:"contingencyType"`   // 条件类型
	ListStatusType    string `json:"listStatusType"`    // 订单列表状态
	ListOrderStatus   string `json:"listOrderStatus"`   // 订单列表状态
	ListClientOrderId string `json:"listClientOrderId"` // 用户自定义的订单号
	TransactionTime   int64  `json:"transactionTime"`   // 交易时间
	Symbol            string `json:"symbol"`            // 交易对
	Orders            []struct {
		Symbol        string `json:"symbol"`        // 交易对
		OrderId       int64  `json:"orderId"`       // 订单ID
		ClientOrderId string `json:"clientOrderId"` // 用户自定义的订单号
	} `json:"orders"` // 订单列表
}

type PortfolioMarginMarginAllOpenOrdersDeleteResRow struct {
	PortfolioMarginMarginOrderDeleteRes
	PortfolioMarginMarginOrderOcoDeleteRes
}

type PortfolioMarginMarginAllOpenOrdersDeleteRes []PortfolioMarginMarginAllOpenOrdersDeleteResRow

type PortfolioMarginMarginOrderOcoOpenOrderListResRow struct {
	OrderListId       int    `json:"orderListId"`       // 订单列表ID
	ContingencyType   string `json:"contingencyType"`   // 条件类型
	ListStatusType    string `json:"listStatusType"`    // 订单列表状态
	ListOrderStatus   string `json:"listOrderStatus"`   // 订单列表状态
	ListClientOrderId string `json:"listClientOrderId"` // 用户自定义的订单号
	TransactionTime   int64  `json:"transactionTime"`   // 交易时间
	Symbol            string `json:"symbol"`            // 交易对
	Orders            []struct {
		Symbol        string `json:"symbol"`        // 交易对
		OrderId       int64  `json:"orderId"`       // 订单ID
		ClientOrderId string `json:"clientOrderId"` // 用户自定义的订单号
	} `json:"orders"` // 订单列表
}
type PortfolioMarginMarginOrderOcoOpenOrderListRes []PortfolioMarginMarginOrderOcoOpenOrderListResRow

type PortfolioMarginMarginOcoAllOrderListResRow struct {
	OrderListId       int    `json:"orderListId"`       // 订单列表ID
	ContingencyType   string `json:"contingencyType"`   // 条件类型
	ListStatusType    string `json:"listStatusType"`    // 订单列表状态
	ListOrderStatus   string `json:"listOrderStatus"`   // 订单列表状态
	ListClientOrderId string `json:"listClientOrderId"` // 用户自定义的订单号
	TransactionTime   int64  `json:"transactionTime"`   // 交易时间
	Symbol            string `json:"symbol"`            // 交易对
	Orders            []struct {
		Symbol        string `json:"symbol"`        // 交易对
		OrderId       int64  `json:"orderId"`       // 订单ID
		ClientOrderId string `json:"clientOrderId"` // 用户自定义的订单号
	} `json:"orders"` // 订单列表
}
type PortfolioMarginMarginOcoAllOrderListRes []PortfolioMarginMarginOcoAllOrderListResRow
