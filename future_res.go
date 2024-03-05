package mybinanceapi

import "github.com/shopspring/decimal"

type FutureAccountResAsset struct {
	Asset                  string `json:"asset"`                  //资产
	WalletBalance          string `json:"walletBalance"`          //余额
	UnrealizedProfit       string `json:"unrealizedProfit"`       // 未实现盈亏
	MarginBalance          string `json:"marginBalance"`          // 保证金余额
	MaintMargin            string `json:"maintMargin"`            // 维持保证金
	InitialMargin          string `json:"initialMargin"`          // 当前所需起始保证金
	PositionInitialMargin  string `json:"positionInitialMargin"`  // 持仓所需起始保证金(基于最新标记价格)
	OpenOrderInitialMargin string `json:"openOrderInitialMargin"` // 当前挂单所需起始保证金(基于最新标记价格)
	CrossWalletBalance     string `json:"crossWalletBalance"`     //全仓账户余额
	CrossUnPnl             string `json:"crossUnPnl"`             // 全仓持仓未实现盈亏
	AvailableBalance       string `json:"availableBalance"`       // 可用余额
	MaxWithdrawAmount      string `json:"maxWithdrawAmount"`      // 最大可转出余额
	MarginAvailable        bool   `json:"marginAvailable"`        // 是否可用作联合保证金
	UpdateTime             int64  `json:"updateTime"`             //更新时间
}

type FutureAccountResPosition struct {
	//根据用户持仓模式展示持仓方向，即单向模式下只返回BOTH持仓情况，双向模式下只返回 LONG 和 SHORT 持仓情况
	Symbol                 string `json:"symbol"`                 // 交易对
	InitialMargin          string `json:"initialMargin"`          // 当前所需起始保证金(基于最新标记价格)
	MaintMargin            string `json:"maintMargin"`            //维持保证金
	UnrealizedProfit       string `json:"unrealizedProfit"`       // 持仓未实现盈亏
	PositionInitialMargin  string `json:"positionInitialMargin"`  // 持仓所需起始保证金(基于最新标记价格)
	OpenOrderInitialMargin string `json:"openOrderInitialMargin"` // 当前挂单所需起始保证金(基于最新标记价格)
	Leverage               string `json:"leverage"`               // 杠杆倍率
	Isolated               bool   `json:"isolated"`               // 是否是逐仓模式
	EntryPrice             string `json:"entryPrice"`             // 持仓成本价
	MaxNotional            string `json:"maxNotional"`            // 当前杠杆下用户可用的最大名义价值
	BidNotional            string `json:"bidNotional"`            // 买单净值，忽略
	AskNotional            string `json:"askNotional"`            // 卖单净值，忽略
	PositionSide           string `json:"positionSide"`           // 持仓方向
	PositionAmt            string `json:"positionAmt"`            // 持仓数量
	UpdateTime             int64  `json:"updateTime"`             // 更新时间
}

type FutureAccountRes struct {
	FeeTier                     int64                      `json:"feeTier"`     // 手续费等级
	CanTrade                    bool                       `json:"canTrade"`    // 是否可以交易
	CanDeposit                  bool                       `json:"canDeposit"`  // 是否可以入金
	CanWithdraw                 bool                       `json:"canWithdraw"` // 是否可以出金
	UpdateTime                  int64                      `json:"updateTime"`  // 保留字段，请忽略
	MultiAssetsMargin           bool                       `json:"multiAssetsMargin"`
	TotalInitialMargin          string                     `json:"totalInitialMargin"`          // 当前所需起始保证金总额(存在逐仓请忽略), 仅计算u
	TotalMaintMargin            string                     `json:"totalMaintMargin"`            // 维持保证金总额, 仅计算usdt资产
	TotalWalletBalance          string                     `json:"totalWalletBalance"`          // 账户总余额, 仅计算usdt资产
	TotalUnrealizedProfit       string                     `json:"totalUnrealizedProfit"`       // 持仓未实现盈亏总额, 仅计算usdt资产
	TotalMarginBalance          string                     `json:"totalMarginBalance"`          // 保证金总余额, 仅计算usdt资产
	TotalPositionInitialMargin  string                     `json:"totalPositionInitialMargin"`  // 持仓所需起始保证金(基于最新标记价格), 仅计算usdt
	TotalOpenOrderInitialMargin string                     `json:"totalOpenOrderInitialMargin"` // 当前挂单所需起始保证金(基于最新标记价格), 仅计算
	TotalCrossWalletBalance     string                     `json:"totalCrossWalletBalance"`     // 全仓账户余额, 仅计算usdt资产
	TotalCrossUnPnl             string                     `json:"totalCrossUnPnl"`             // 全仓持仓未实现盈亏总额, 仅计算usdt资产
	AvailableBalance            string                     `json:"availableBalance"`            // 可用余额, 仅计算usdt资产
	MaxWithdrawAmount           string                     `json:"maxWithdrawAmount"`           // 最大可转出余额, 仅计算usdt资产
	Assets                      []FutureAccountResAsset    `json:"assets"`
	Positions                   []FutureAccountResPosition `json:"positions"` // 头寸，将返回所有市场symbol。
}

type FuturePositionSideDualGetRes struct {
	DualSidePosition bool `json:"dualSidePosition"` // "true": 双向持仓模式；"false": 单向持仓模式
}

type FutureMultiAssetsMarginGetRes struct {
	MultiAssetsMargin bool `json:"multiAssetsMargin"` // "true": 联合保证金模式开启；"false": 联合保证金模式关闭
}
type FutureCommonPostRes struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

type FuturePositionSideDualPostRes FutureCommonPostRes

type FutureMultiAssetsMarginPostRes FutureCommonPostRes

type FutureLeverageRes struct {
	Leverage         int64  `json:"leverage"`         // 杠杆倍数
	MaxNotionalValue string `json:"maxNotionalValue"` // 当前杠杆倍数下允许的最大名义价值
	Symbol           string `json:"symbol"`           // 交易对
}

type FutureMarginTypeRes FutureCommonPostRes

// [
//
//	{
//	    "symbol": "ETHUSDT",
//	    "brackets": [
//	        {
//	            "bracket": 1,   // 层级
//	            "initialLeverage": 75,  // 该层允许的最高初始杠杆倍数
//	            "notionalCap": 10000,  // 该层对应的名义价值上限
//	            "notionalFloor": 0,  // 该层对应的名义价值下限
//	            "maintMarginRatio": 0.0065, // 该层对应的维持保证金率
//	            "cum":0 // 速算数
//	        },
//	    ]
//	}
//
// ]

type FutureLeverageBracketRes []FutureLeverageBracketResResult
type FutureLeverageBracketResResult struct {
	Symbol   string                                   `json:"symbol"`   // 交易对
	Brackets []FutureLeverageBracketResResultBrackets `json:"brackets"` // 交易对
}

type FutureLeverageBracketResResultBrackets struct {
	Bracket          int64           `json:"bracket"`          // 层级
	InitialLeverage  int64           `json:"initialLeverage"`  // 该层允许的最高初始杠杆倍数
	NotionalCap      int64           `json:"notionalCap"`      // 该层对应的名义价值上限
	NotionalFloor    int64           `json:"notionalFloor"`    // 该层对应的名义价值下限
	MaintMarginRatio decimal.Decimal `json:"maintMarginRatio"` // 该层对应的维持保证金率
	Cum              decimal.Decimal `json:"cum"`              // 速算数
}

type FuturePingRes struct {
}

type FutureTimeRes struct {
	ServerTime int64 `json:"serverTime"`
}

type FutureExchangeInfoRes struct {
	ExchangeFilters []map[string]interface{}         `json:"exchangeFilters"`
	RateLimits      []FutureExchangeInfoResRateLimit `json:"rateLimits"` // API访问的限制
	Assets          []FutureExchangeInfoResAsset     `json:"assets"`     // 资产信息
	Symbols         []FutureExchangeInfoResSymbol    `json:"symbols"`    // 交易对信息
	Timezone        string                           `json:"timezone"`   // 服务器所用的时间区域
}

type FutureExchangeInfoResRateLimit struct {
	Interval      string `json:"interval"`      // 限制的时间间隔
	IntervalNum   int64  `json:"intervalNum"`   // 限制的时间间隔数量
	Limit         int64  `json:"limit"`         // 限制的请求数量
	RateLimitType string `json:"rateLimitType"` // 限制的请求数量的类型
}

type FutureExchangeInfoResAsset struct {
	Asset             string `json:"asset"`             // 资产
	MarginAvailable   bool   `json:"marginAvailable"`   // 是否可用作保证金
	AutoAssetExchange string `json:"autoAssetExchange"` // 保证金资产自动兑换阈值
}

type FutureExchangeInfoResSymbol struct {
	Symbol             string                   `json:"symbol"`             // 交易对
	Pair               string                   `json:"pair"`               // 标的交易对
	ContractType       string                   `json:"contractType"`       // 合约类型
	DeliveryDate       int64                    `json:"deliveryDate"`       // 交割日期
	OnboardDate        int64                    `json:"onboardDate"`        // 上线日期
	Status             string                   `json:"status"`             // 交易对状态
	BaseAsset          string                   `json:"baseAsset"`          // 标的资产
	QuoteAsset         string                   `json:"quoteAsset"`         // 报价资产
	MarginAsset        string                   `json:"marginAsset"`        // 保证金资产
	PricePrecision     int                      `json:"pricePrecision"`     // 价格小数点位数(仅作为系统精度使用，注意同tickSize 区分）
	QuantityPrecision  int                      `json:"quantityPrecision"`  // 数量小数点位数(仅作为系统精度使用，注意同stepSize 区分）
	BaseAssetPrecision int                      `json:"baseAssetPrecision"` // 标的资产精度
	QuotePrecision     int                      `json:"quotePrecision"`     // 报价资产精度
	UnderlyingType     string                   `json:"underlyingType"`
	UnderlyingSubType  []string                 `json:"underlyingSubType"`
	SettlePlan         int                      `json:"settlePlan"`
	TriggerProtect     string                   `json:"triggerProtect"` // 开启"priceProtect"的条件订单的触发阈值
	Filters            []map[string]interface{} `json:"filters"`
	OrderType          []string                 `json:"OrderType"`       // 订单类型
	TimeInForce        []string                 `json:"timeInForce"`     // 有效方式
	LiquidationFee     string                   `json:"liquidationFee"`  // 强平费率
	MarketTakeBound    string                   `json:"marketTakeBound"` // 市价吃单(相对于标记价格)允许可造成的最大价格偏离比例
}

// {
//     "avgPrice": "0.00000",              // 平均成交价
//     "clientOrderId": "abc",             // 用户自定义的订单号
//     "cumQuote": "0",                        // 成交金额
//     "executedQty": "0",                 // 成交量
//     "orderId": 1917641,                 // 系统订单号
//     "origQty": "0.40",                  // 原始委托数量
//     "origType": "TRAILING_STOP_MARKET", // 触发前订单类型
//     "price": "0",                   // 委托价格
//     "reduceOnly": false,                // 是否仅减仓
//     "side": "BUY",                      // 买卖方向
//     "status": "NEW",                    // 订单状态
//     "positionSide": "SHORT", // 持仓方向
//     "stopPrice": "9300",                    // 触发价，对`TRAILING_STOP_MARKET`无效
//     "closePosition": false,   // 是否条件全平仓
//     "symbol": "BTCUSDT",                // 交易对
//     "time": 1579276756075,              // 订单时间
//     "timeInForce": "GTC",               // 有效方法
//     "type": "TRAILING_STOP_MARKET",     // 订单类型
//     "activatePrice": "9020", // 跟踪止损激活价格, 仅`TRAILING_STOP_MARKET` 订单返回此字段
//     "priceRate": "0.3", // 跟踪止损回调比例, 仅`TRAILING_STOP_MARKET` 订单返回此字段
//     "updateTime": 1579276756075,        // 更新时间
//     "workingType": "CONTRACT_PRICE", // 条件价格触发类型
//     "priceProtect": false            // 是否开启条件单触发保护
// }

type FutureOrderOrder struct {
	AvgPrice      string `json:"avgPrice"`                  // 平均成交价
	ClientOrderId string `json:"clientOrderId"`             // 用户自定义的订单号
	CumQuote      string `json:"cumQuote"`                  // 成交金额
	ExecutedQty   string `json:"executedQty"`               // 成交量
	OrderId       int64  `gorm:"primaryKey" json:"orderId"` // 系统订单号
	OrigQty       string `json:"origQty"`                   // 原始委托数量
	OrigType      string `json:"origType"`                  // 触发前订单类型
	Price         string `json:"price"`                     // 委托价格
	ReduceOnly    bool   `json:"reduceOnly"`                // 是否仅减仓
	Side          string `json:"side"`                      // 买卖方向
	Status        string `json:"status"`                    // 订单状态
	PositionSide  string `json:"positionSide"`              // 持仓方向
	StopPrice     string `json:"stopPrice"`                 // 触发价，对`TRAILING_STOP_MARKET`无效
	ClosePosition bool   `json:"closePosition"`             // 是否条件全平仓
	Symbol        string `json:"symbol"`                    // 交易对
	Time          int64  `json:"time"`                      // 订单时间
	TimeInForce   string `json:"timeInForce"`               // 有效方法
	Type          string `json:"type"`                      // 订单类型
	ActivatePrice string `json:"activatePrice"`             // 跟踪止损激活价格, 仅`TRAILING_STOP_MARKET` 订单返回此字段
	PriceRate     string `json:"priceRate"`                 // 跟踪止损回调比例, 仅`TRAILING_STOP_MARKET` 订单返回此字段
	UpdateTime    int64  `json:"updateTime"`                // 更新时间
	WorkingType   string `json:"workingType"`               // 条件价格触发类型
	PriceProtect  bool   `json:"priceProtect"`              // 是否开启条件单触发保护
}

type FutureOpenOrdersRes []FutureOrderOrder

type FutureAllOrdersRes []FutureOrderOrder

type FutureOrderGetRes FutureOrderOrder

// {
//     "clientOrderId": "testOrder", // 用户自定义的订单号
//     "cumQty": "0",
//     "cumQuote": "0", // 成交金额
//     "executedQty": "0", // 成交量
//     "orderId": 22542179, // 系统订单号
//     "avgPrice": "0.00000",  // 平均成交价
//     "origQty": "10", // 原始委托数量
//     "price": "0", // 委托价格
//     "reduceOnly": false, // 仅减仓
//     "side": "SELL", // 买卖方向
//     "positionSide": "SHORT", // 持仓方向
//     "status": "NEW", // 订单状态
//     "stopPrice": "0", // 触发价，对`TRAILING_STOP_MARKET`无效
//     "closePosition": false,   // 是否条件全平仓
//     "symbol": "BTCUSDT", // 交易对
//     "timeInForce": "GTC", // 有效方法
//     "type": "TRAILING_STOP_MARKET", // 订单类型
//     "origType": "TRAILING_STOP_MARKET",  // 触发前订单类型
//     "activatePrice": "9020", // 跟踪止损激活价格, 仅`TRAILING_STOP_MARKET` 订单返回此字段
//     "priceRate": "0.3", // 跟踪止损回调比例, 仅`TRAILING_STOP_MARKET` 订单返回此字段
//     "updateTime": 1566818724722, // 更新时间
//     "workingType": "CONTRACT_PRICE", // 条件价格触发类型
//     "priceProtect": false            // 是否开启条件单触发保护
// }
type FutureOrderPostRes struct {
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

// {
//     "symbol": "BTCUSDT",
//     "makerCommissionRate": "0.0002",  // 0.02%
//     "takerCommissionRate": "0.0004"   // 0.04%
// }
type FutureCommissionRateRes struct {
	Symbol              string `json:"symbol"`
	MakerCommissionRate string `json:"makerCommissionRate"`
	TakerCommissionRate string `json:"takerCommissionRate"`
}

// [
//   {
//     "buyer": false, // 是否是买方
//     "commission": "-0.07819010", // 手续费
//     "commissionAsset": "USDT", // 手续费计价单位
//     "id": 698759,   // 交易ID
//     "maker": false, // 是否是挂单方
//     "orderId": 25851813, // 订单编号
//     "price": "7819.01", // 成交价
//     "qty": "0.002", // 成交量
//     "quoteQty": "15.63802", // 成交额
//     "realizedPnl": "-0.91539999",   // 实现盈亏
//     "side": "SELL", // 买卖方向
//     "positionSide": "SHORT",  // 持仓方向
//     "symbol": "BTCUSDT", // 交易对
//     "time": 1569514978020 // 时间
//   }
// ]

type FutureUserTrade struct {
	Buyer           bool   `json:"buyer"`           // 是否是买方
	Commission      string `json:"commission"`      // 手续费
	CommissionAsset string `json:"commissionAsset"` // 手续费计价单位
	Id              int64  `json:"id"`              // 交易ID
	Maker           bool   `json:"maker"`           // 是否是挂单方
	OrderId         int64  `json:"orderId"`         // 订单编号
	Price           string `json:"price"`           // 成交价
	Qty             string `json:"qty"`             // 成交量
	QuoteQty        string `json:"quoteQty"`        // 成交额
	RealizedPnl     string `json:"realizedPnl"`     // 实现盈亏
	Side            string `json:"side"`            // 买卖方向
	PositionSide    string `json:"positionSide"`    // 持仓方向
	Symbol          string `json:"symbol"`          // 交易对
	Time            int64  `json:"time"`            // 时间
}
type FutureUserTradesRes []FutureUserTrade

// {
//     "clientOrderId": "myOrder1", // 用户自定义的订单号
//     "cumQty": "0",
//     "cumQuote": "0", // 成交金额
//     "executedQty": "0", // 成交量
//     "orderId": 283194212, // 系统订单号
//     "origQty": "11", // 原始委托数量
//     "price": "0", // 委托价格
//     "reduceOnly": false, // 仅减仓
//     "side": "BUY", // 买卖方向
//     "positionSide": "SHORT", // 持仓方向
//     "status": "CANCELED", // 订单状态
//     "stopPrice": "9300", // 触发价，对`TRAILING_STOP_MARKET`无效
//     "closePosition": false,   // 是否条件全平仓
//     "symbol": "BTCUSDT", // 交易对
//     "timeInForce": "GTC", // 有效方法
//     "origType": "TRAILING_STOP_MARKET", // 触发前订单类型
//     "type": "TRAILING_STOP_MARKET", // 订单类型
//     "activatePrice": "9020", // 跟踪止损激活价格, 仅`TRAILING_STOP_MARKET` 订单返回此字段
//     "priceRate": "0.3", // 跟踪止损回调比例, 仅`TRAILING_STOP_MARKET` 订单返回此字段
//     "updateTime": 1571110484038, // 更新时间
//     "workingType": "CONTRACT_PRICE", // 条件价格触发类型
//     "priceProtect": false,            // 是否开启条件单触发保护
//     "priceMatch": "NONE",              //盘口价格下单模式
//     "selfTradePreventionMode": "NONE", //订单自成交保护模式
//     "goodTillDate": 0                  //订单TIF为GTD时的自动取消时间
// }

type FutureOrderDeleteRes struct {
	BinanceErrorRes
	ClientOrderId           string `json:"clientOrderId"` // 用户自定义的订单号
	CumQty                  string `json:"cumQty"`
	CumQuote                string `json:"cumQuote"`                // 成交金额
	ExecutedQty             string `json:"executedQty"`             // 成交量
	OrderId                 int64  `json:"orderId"`                 // 系统订单号
	OrigQty                 string `json:"origQty"`                 // 原始委托数量
	Price                   string `json:"price"`                   // 委托价格
	ReduceOnly              bool   `json:"reduceOnly"`              // 仅减仓
	Side                    string `json:"side"`                    // 买卖方向
	PositionSide            string `json:"positionSide"`            // 持仓方向
	Status                  string `json:"status"`                  // 订单状态
	StopPrice               string `json:"stopPrice"`               // 触发价，对`TRAILING_STOP_MARKET`无效
	ClosePosition           bool   `json:"closePosition"`           // 是否条件全平仓
	Symbol                  string `json:"symbol"`                  // 交易对
	TimeInForce             string `json:"timeInForce"`             // 有效方法
	OrigType                string `json:"origType"`                // 触发前订单类型
	Type                    string `json:"type"`                    // 订单类型
	ActivatePrice           string `json:"activatePrice"`           // 跟踪止损激活价格, 仅`TRAILING_STOP_MARKET` 订单返回此字段
	PriceRate               string `json:"priceRate"`               // 跟踪止损回调比例, 仅`TRAILING_STOP_MARKET` 订单返回此字段
	UpdateTime              int64  `json:"updateTime"`              // 更新时间
	WorkingType             string `json:"workingType"`             // 条件价格触发类型
	PriceProtect            bool   `json:"priceProtect"`            // 是否开启条件单触发保护
	PriceMatch              string `json:"priceMatch"`              // 盘口价格下单模式
	SelfTradePreventionMode string `json:"selfTradePreventionMode"` // 订单自成交保护模式
	GoodTillDate            int64  `json:"goodTillDate"`            // 订单TIF为GTD时的自动取消时间
}

// [
//     {
//         "clientOrderId": "myOrder1", // 用户自定义的订单号
//         "cumQty": "0",
//         "cumQuote": "0", // 成交金额
//         "executedQty": "0", // 成交量
//         "orderId": 283194212, // 系统订单号
//         "origQty": "11", // 原始委托数量
//         "price": "0", // 委托价格
//         "reduceOnly": false, // 仅减仓
//         "side": "BUY", // 买卖方向
//         "positionSide": "SHORT", // 持仓方向
//         "status": "CANCELED", // 订单状态
//         "stopPrice": "9300", // 触发价，对`TRAILING_STOP_MARKET`无效
//         "closePosition": false,   // 是否条件全平仓
//         "symbol": "BTCUSDT", // 交易对
//         "timeInForce": "GTC", // 有效方法
//         "origType": "TRAILING_STOP_MARKET", // 触发前订单类型
//         "type": "TRAILING_STOP_MARKET", // 订单类型
//         "activatePrice": "9020", // 跟踪止损激活价格, 仅`TRAILING_STOP_MARKET` 订单返回此字段
//         "priceRate": "0.3", // 跟踪止损回调比例, 仅`TRAILING_STOP_MARKET` 订单返回此字段
//         "updateTime": 1571110484038, // 更新时间
//         "workingType": "CONTRACT_PRICE", // 条件价格触发类型
//         "priceProtect": false,            // 是否开启条件单触发保护
//         "priceMatch": "NONE",              //盘口价格下单模式
//         "selfTradePreventionMode": "NONE", //订单自成交保护模式
//         "goodTillDate": 0                  //订单TIF为GTD时的自动取消时间
//     },
//     {
//         "code": -2011,
//         "msg": "Unknown order sent."
//     }
// ]

type FutureBatchOrdersDeleteRes []FutureOrderDeleteRes

// {
// 	"lastUpdateId": 1027024,
// 	"E": 1589436922972,   // 消息时间
// 	"T": 1589436922959,   // 撮合引擎时间
// 	"bids": [             // 买单
// 	  [
// 		"4.00000000",     // 价格
// 		"431.00000000"    // 数量
// 	  ]
// 	],
// 	"asks": [             // 卖单
// 	  [
// 		"4.00000200",     // 价格
// 		"12.00000000"     // 数量
// 	  ]
// 	]
// }
type FutureDepthRes struct {
	LastUpdateId int64       `json:"lastUpdateId"` // 最后更新ID
	MessageTime  int64       `json:"E"`            // 消息时间
	EngineTime   int64       `json:"T"`            // 撮合引擎时间
	Bids         []DepthGear `json:"bids"`         // 买单
	Asks         []DepthGear `json:"asks"`         // 卖单
}

type FutureDepthResMiddle struct {
	LastUpdateId int64           `json:"lastUpdateId"` // 最后更新ID
	MessageTime  int64           `json:"E"`            // 消息时间
	EngineTime   int64           `json:"T"`            // 撮合引擎时间
	Bids         [][]interface{} `json:"bids"`         // 买单
	Asks         [][]interface{} `json:"asks"`         // 卖单
}

func (middle *FutureDepthResMiddle) ConvertToRes() *FutureDepthRes {
	res := FutureDepthRes{
		LastUpdateId: middle.LastUpdateId,
		MessageTime:  middle.MessageTime,
		EngineTime:   middle.EngineTime,
	}

	res.Bids = []DepthGear{}
	res.Asks = []DepthGear{}
	for _, gear := range middle.Bids {
		res.Bids = append(res.Bids, DepthGear{
			Price:    gear[0].(string),
			Quantity: gear[1].(string),
		})
	}
	for _, gear := range middle.Asks {
		res.Asks = append(res.Asks, DepthGear{
			Price:    gear[0].(string),
			Quantity: gear[1].(string),
		})
	}
	return &res
}

// [
//   {
//     "id": 28457,                // 成交ID
//     "price": "4.00000100",      // 成交价格
//     "qty": "12.00000000",       // 成交量
//     "quoteQty": "48.00",        // 成交额
//     "time": 1499865549590,      // 时间
//     "isBuyerMaker": true        // 买方是否为挂单方
//   }
// ]
type FutureTradesRes []FutureTradesResRow
type FutureTradesResRow struct {
	Id           int64  `json:"id"`           // 成交ID
	Price        string `json:"price"`        // 成交价格
	Qty          string `json:"qty"`          // 成交量
	QuoteQty     string `json:"quoteQty"`     // 成交额
	Time         int64  `json:"time"`         // 时间
	IsBuyerMaker bool   `json:"isBuyerMaker"` // 买方是否为挂单方
}

type FutureHistoricalTradesRes []FutureTradesResRow

// [
//   {
//     "a": 26129,         // 归集成交ID
//     "p": "0.01633102",  // 成交价
//     "q": "4.70443515",  // 成交量
//     "f": 27781,         // 被归集的首个成交ID
//     "l": 27781,         // 被归集的末个成交ID
//     "T": 1498793709153, // 成交时间
//     "m": true,          // 是否为主动卖出单
//   }
// ]
type FutureAggTradesRes []FutureAggTradesResRow

type FutureAggTradesResRow struct {
	Id           int64  `json:"a"` // 归集成交ID
	Price        string `json:"p"` // 成交价
	Qty          string `json:"q"` // 成交量
	FirstTradeId int64  `json:"f"` // 被归集的首个成交ID
	LastTradeId  int64  `json:"l"` // 被归集的末个成交ID
	Time         int64  `json:"T"` // 成交时间
	IsBuyerMaker bool   `json:"m"` // 是否为主动卖出单
}

// {
//     "symbol": "BTCUSDT",                // 交易对
//     "markPrice": "11793.63104562",      // 标记价格
//     "indexPrice": "11781.80495970",     // 指数价格
//     "estimatedSettlePrice": "11781.16138815",  // 预估结算价,仅在交割开始前最后一小时有意义
//     "lastFundingRate": "0.00038246",    // 最近更新的资金费率
//     "nextFundingTime": 1597392000000,   // 下次资金费时间
//     "interestRate": "0.00010000",       // 标的资产基础利率
//     "time": 1597370495002               // 更新时间
// }

type FuturePremiumIndexRes []FuturePremiumIndexResRow
type FuturePremiumIndexResRow struct {
	Symbol               string `json:"symbol"`               // 交易对
	MarkPrice            string `json:"markPrice"`            // 标记价格
	IndexPrice           string `json:"indexPrice"`           // 指数价格
	EstimatedSettlePrice string `json:"estimatedSettlePrice"` // 预估结算价,仅在交割开始前最后一小时有意义
	LastFundingRate      string `json:"lastFundingRate"`      // 最近更新的资金费率
	NextFundingTime      int64  `json:"nextFundingTime"`      // 下次资金费时间
	InterestRate         string `json:"interestRate"`         // 标的资产基础利率
	Time                 int64  `json:"time"`                 // 更新时间
}

// [
//     {
//         "symbol": "BTCUSDT",            // 交易对
//         "fundingTime": 1698768000000,   // 资金费率
//         "fundingRate": "0.00010000",    // 资金费时间
//         "markPrice": "34287.54619963"   // 资金费对应标记价格
//     },
//     {
//         "symbol": "BTCUSDT",
//         "fundingTime": 1698796800000,
//         "fundingRate": "0.00010000",
//         "markPrice": "34651.40000000"
//     }
// ]
type FutureFundingRateRes []FutureFundingRateResRow
type FutureFundingRateResRow struct {
	Symbol      string `json:"symbol"`      // 交易对
	FundingTime int64  `json:"fundingTime"` // 资金费率
	FundingRate string `json:"fundingRate"` // 资金费时间
	MarkPrice   string `json:"markPrice"`   // 资金费对应标记价格
}

// [
//     {
//         "symbol": "BLZUSDT",
//         "adjustedFundingRateCap": "0.02500000",
//         "adjustedFundingRateFloor": "-0.02500000",
//         "fundingIntervalHours": 8,
//         "disclaimer": false
//     }
// ]
type FutureFundingInfoRes []FutureFundingInfoResRow
type FutureFundingInfoResRow struct {
	Symbol                   string `json:"symbol"`                   // 交易对
	AdjustedFundingRateCap   string `json:"adjustedFundingRateCap"`   // 资金费率上限
	AdjustedFundingRateFloor string `json:"adjustedFundingRateFloor"` // 资金费率下限
	FundingIntervalHours     int64  `json:"fundingIntervalHours"`     // 资金费率间隔时间
	Disclaimer               bool   `json:"disclaimer"`
}

// {
// 	"symbol": "BTCUSDT",
// 	"priceChange": "-94.99999800",    //24小时价格变动
// 	"priceChangePercent": "-95.960",  //24小时价格变动百分比
// 	"weightedAvgPrice": "0.29628482", //加权平均价
// 	"lastPrice": "4.00000200",        //最近一次成交价
// 	"lastQty": "200.00000000",        //最近一次成交额
// 	"openPrice": "99.00000000",       //24小时内第一次成交的价格
// 	"highPrice": "100.00000000",      //24小时最高价
// 	"lowPrice": "0.10000000",         //24小时最低价
// 	"volume": "8913.30000000",        //24小时成交量
// 	"quoteVolume": "15.30000000",     //24小时成交额
// 	"openTime": 1499783499040,        //24小时内，第一笔交易的发生时间
// 	"closeTime": 1499869899040,       //24小时内，最后一笔交易的发生时间
// 	"firstId": 28385,   // 首笔成交id
// 	"lastId": 28460,    // 末笔成交id
// 	"count": 76         // 成交笔数
//   }
type FutureTicker24hrRes []FutureTicker24hrResRow
type FutureTicker24hrResRow struct {
	Symbol             string `json:"symbol"`             // 交易对
	PriceChange        string `json:"priceChange"`        // 24小时价格变动
	PriceChangePercent string `json:"priceChangePercent"` // 24小时价格变动百分比
	WeightedAvgPrice   string `json:"weightedAvgPrice"`   // 加权平均价
	LastPrice          string `json:"lastPrice"`          // 最近一次成交价
	LastQty            string `json:"lastQty"`            // 最近一次成交额
	OpenPrice          string `json:"openPrice"`          // 24小时内第一次成交的价格
	HighPrice          string `json:"highPrice"`          // 24小时最高价
	LowPrice           string `json:"lowPrice"`           // 24小时最低价
	Volume             string `json:"volume"`             // 24小时成交量
	QuoteVolume        string `json:"quoteVolume"`        // 24小时成交额
	OpenTime           int64  `json:"openTime"`           // 24小时内，第一笔交易的发生时间
	CloseTime          int64  `json:"closeTime"`          // 24小时内，最后一笔交易的发生时间
	FirstId            int64  `json:"firstId"`            // 首笔成交id
	LastId             int64  `json:"lastId"`             // 末笔成交id
	Count              int64  `json:"count"`              // 成交笔数
}

// {
// 	"symbol": "LTCBTC",       // 交易对
// 	"price": "4.00000200",        // 价格
// 	"time": 1589437530011   // 撮合引擎时间
//   }

type FutureTickerPriceRes []FutureTickerPriceResRow
type FutureTickerPriceResRow struct {
	Symbol string `json:"symbol"` // 交易对
	Price  string `json:"price"`  // 价格
	Time   int64  `json:"time"`   // 撮合引擎时间
}

// {
// 	"lastUpdateId": 1027024,
// 	"symbol": "BTCUSDT", // 交易对
// 	"bidPrice": "4.00000000", //最优买单价
// 	"bidQty": "431.00000000", //挂单量
// 	"askPrice": "4.00000200", //最优卖单价
// 	"askQty": "9.00000000", //挂单量
// 	"time": 1589437530011   // 撮合引擎时间
//   }

type FutureTickerBookTickerRes []FutureTickerBookTickerResRow
type FutureTickerBookTickerResRow struct {
	LastUpdateId int64  `json:"lastUpdateId"` // 最后更新ID
	Symbol       string `json:"symbol"`       // 交易对
	BidPrice     string `json:"bidPrice"`     // 最优买单价
	BidQty       string `json:"bidQty"`       // 挂单量
	AskPrice     string `json:"askPrice"`     // 最优卖单价
	AskQty       string `json:"askQty"`       // 挂单量
	Time         int64  `json:"time"`         // 撮合引擎时间
}

// [
//     {
//         "indexPrice": "34400.15945055",
//         "contractType": "PERPETUAL",
//         "basisRate": "0.0004",
//         "futuresPrice": "34414.10",
//         "annualizedBasisRate": "",
//         "basis": "13.94054945",
//         "pair": "BTCUSDT",
//         "timestamp": 1698742800000
//     }
// ]

type FutureDataBasisRes []FutureDataBasisResRow
type FutureDataBasisResRow struct {
	IndexPrice          string `json:"indexPrice"`
	ContractType        string `json:"contractType"`
	BasisRate           string `json:"basisRate"`
	FuturesPrice        string `json:"futuresPrice"`
	AnnualizedBasisRate string `json:"annualizedBasisRate"`
	Basis               string `json:"basis"`
	Pair                string `json:"pair"`
	Timestamp           int64  `json:"timestamp"`
}

type FutureListenKeyPostRes struct {
	ListenKey string `json:"listenKey"`
}

type FutureListenKeyPutRes struct{}

type FutureListenKeyDeleteRes struct{}
