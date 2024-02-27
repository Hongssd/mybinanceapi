package mybinanceapi

import "github.com/shopspring/decimal"

type SwapAccountRes struct {
	Assets []struct {
		Asset                  string `json:"asset"`                  // 资产名
		WalletBalance          string `json:"walletBalance"`          // 账户余额
		UnrealizedProfit       string `json:"unrealizedProfit"`       // 全部持仓未实现盈亏
		MarginBalance          string `json:"marginBalance"`          // 保证金余额
		MaintMargin            string `json:"maintMargin"`            // 维持保证金
		InitialMargin          string `json:"initialMargin"`          // 当前所需起始保证金(按最新标标记价格)
		PositionInitialMargin  string `json:"positionInitialMargin"`  // 当前所需持仓起始保证金(按最新标标记价格)
		OpenOrderInitialMargin string `json:"openOrderInitialMargin"` // 当前所需挂单起始保证金(按最新标标记价格)
		MaxWithdrawAmount      string `json:"maxWithdrawAmount"`      // 最大可提款金额
		CrossWalletBalance     string `json:"crossWalletBalance"`     // 可用于全仓的账户余额
		CrossUnPnl             string `json:"crossUnPnl"`             // 所有全仓持仓的未实现盈亏
		AvailableBalance       string `json:"availableBalance"`       // 可用下单余额
	} `json:"assets"` // 资产内容

	Positions []struct {
		Symbol                 string `json:"symbol"`                 // 交易对
		PositionAmt            string `json:"positionAmt"`            // 持仓数量
		InitialMargin          string `json:"initialMargin"`          // 当前所需起始保证金(按最新标标记价格)
		MaintMargin            string `json:"maintMargin"`            // 持仓维持保证金
		UnrealizedProfit       string `json:"unrealizedProfit"`       // 持仓未实现盈亏
		PositionInitialMargin  string `json:"positionInitialMargin"`  // 当前所需持仓起始保证金(按最新标标记价格)
		OpenOrderInitialMargin string `json:"openOrderInitialMargin"` // 当前所需挂单起始保证金(按最新标标记价格)
		Leverage               string `json:"leverage"`               // 杠杆倍率
		Isolated               bool   `json:"isolated"`               // 是否是逐仓模式
		PositionSide           string `json:"positionSide"`           // 持仓方向
		EntryPrice             string `json:"entryPrice"`             // 平均持仓成本
		UpdateTime             int64  `json:"updateTime"`             // 最新更新时间
		MaxQty                 string `json:"maxQty"`                 // 当前杠杆下最大可开仓数(标的数量)
	} `json:"positions"` // 头寸
	CanDeposit  bool  `json:"canDeposit"`  // 是否可以入金
	CanTrade    bool  `json:"canTrade"`    // 是否可以交易
	CanWithdraw bool  `json:"canWithdraw"` // 是否可以出金
	FeeTier     int64 `json:"feeTier"`     // 手续费等级
	UpdateTime  int64 `json:"updateTime"`
}

type SwapCommonPostRes struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

type SwapPositionSideDualGetRes struct {
	DualSidePosition bool `json:"dualSidePosition"` // "true": 双向持仓模式；"false": 单向持仓模式
}

type SwapPositionSideDualPostRes SwapCommonPostRes

type SwapLeverageRes struct {
	Leverage         int64  `json:"leverage"`         // 杠杆倍数
	MaxNotionalValue string `json:"maxNotionalValue"` // 当前杠杆倍数下允许的最大名义价值
	Symbol           string `json:"symbol"`           // 交易对
}

type SwapMarginTypeRes SwapCommonPostRes

// [
//     {
//         "symbol": "BTCUSD_PERP",
//         "brackets": [
//             {
//                 "bracket": 1,   // 层级
//                 "initialLeverage": 125,  // 该层允许的最高初始杠杆倍数
//                 "qtyCap": 50,  // 该层对应的数量上限
//                 "qtylFloor": 0,  // 该层对应的数量下限
//                 "maintMarginRatio": 0.004 // 该层对应的维持保证金率
//                 "cum": 0.0  //速算数
//             },
//         ]
//     }
// ]

type SwapLeverageBracketRes []SwapLeverageBracketResResult

type SwapLeverageBracketResResult struct {
	Symbol   string                                 `json:"symbol"`   // 交易对
	Brackets []SwapLeverageBracketResResultBrackets `json:"brackets"` // 杠杆档位
}

type SwapLeverageBracketResResultBrackets struct {
	Bracket          int64           `json:"bracket"`          // 层级
	InitialLeverage  int64           `json:"initialLeverage"`  // 该层允许的最高初始杠杆倍数
	QtyCap           int64           `json:"qtyCap"`           // 该层对应的数量上限
	QtyFloor         int64           `json:"qtyFloor"`         // 该层对应的数量下限
	MaintMarginRatio decimal.Decimal `json:"maintMarginRatio"` // 该层对应的维持保证金率
	Cum              decimal.Decimal `json:"cum"`              // 速算数
}

type SwapPingRes struct {
}

type SwapServerTimeRes struct {
	ServerTime int64 `json:"serverTime"`
}

// {
// 	"exchangeFilters": [],
// 	"rateLimits": [ // API访问的限制
// 		{
// 			"interval": "MINUTE", // 按照分钟计算
// 			"intervalNum": 1, // 按照1分钟计算
// 			"limit": 2400, // 上限次数
// 			"rateLimitType": "REQUEST_WEIGHT" // 按照访问权重来计算
// 		},
// 		{
// 			"interval": "MINUTE",
// 			"intervalNum": 1,
// 			"limit": 1200,
// 			"rateLimitType": "ORDERS" // 按照订单数量来计算
// 		}
// 	],
// 	"serverTime": 1565613908500, // 请忽略。如果需要获取系统时间，请查询接口 “GET /dapi/v1/time”
// 	"symbols": [ // 交易对信息
// 		{
// 			"filters": [
// 				{
// 					"filterType": "PRICE_FILTER", // 价格限制
// 					"maxPrice": "100000", // 价格上限, 最大价格
// 					"minPrice": "0.1", // 价格下限, 最小价格
// 					"tickSize": "0.1" // 下单最小价格间隔
// 				},
// 				{
// 					"filterType": "LOT_SIZE", // 数量限制
// 					"maxQty": "100000", // 数量上限, 最大数量
// 					"minQty": "1", // 数量下限, 最小数量
// 					"stepSize": "1" // 下单最小数量间隔
// 				},
// 				{
// 					"filterType": "MARKET_LOT_SIZE", // 市价订单数量限制
// 					"maxQty": "100000", // 数量上限, 最大数量
// 					"minQty": "1", // 数量下限, 最小数量
// 					"stepSize": "1" // 允许的步进值
// 				},
// 				{
// 					"filterType": "MAX_NUM_ORDERS", // 最多挂单数限制
// 					"limit": 200
// 				},
// 				{
// 					"filterType": "PERCENT_PRICE", // 价格比限制
// 					"multiplierUp": "1.0500", // 价格上限百分比
// 					"multiplierDown": "0.9500", // 价格下限百分比
// 					"multiplierDecimal": 4
// 				}
// 			],
// 			"OrderType": [ // 订单类型
// 				"LIMIT",  // 限价单
// 				"MARKET",  // 市价单
// 				"STOP", // 止损单
// 				"TAKE_PROFIT", // 止盈单
// 				"TRAILING_STOP_MARKET" // 跟踪止损单
// 			],
// 			"timeInForce": [ // 有效方式
// 				"GTC", // 成交为止, 一直有效
// 				"IOC", // 无法立即成交(吃单)的部分就撤销
// 				"FOK", // 无法全部立即成交就撤销
// 				"GTX" // 无法成为挂单方就撤销
// 			],
// 			"liquidationFee": "0.010000",   // 强平费率
// 			"marketTakeBound": "0.30",  // 市价吃单(相对于标记价格)允许可造成的最大价格偏离比例
// 			"symbol": "BTCUSD_200925", // 交易对
// 			"pair": "BTCUSD",   // 标的交易对
// 			"contractType": "CURRENT_QUARTER",   // 合约类型
// 			"deliveryDate": 1601020800000,
// 			"onboardDate": 1590739200000,
// 			"contractStatus": "TRADING", // 交易对状态
// 			"contractSize": 100,     //
// 			"quoteAsset": "USD", // 报价币种
// 			"baseAsset": "BTC",  // 标的物
// 			"marginAsset": "BTC",   // 保证金币种
// 			"pricePrecision": 1,   // 价格小数点位数(仅作为系统精度使用，注意同tickSize 区分)
// 			"quantityPrecision": 0, // 数量小数点位数(仅作为系统精度使用，注意同stepSize 区分)
// 			"baseAssetPrecision": 8,
// 			"quotePrecision": 8,
// 			"equalQtyPrecision": 4,     // 请忽略
// 			"triggerProtect": "0.0500", // 开启"priceProtect"的条件订单的触发阈值
// 			"maintMarginPercent": "2.5000", // 请忽略
// 			"requiredMarginPercent": "5.0000", // 请忽略
// 			"underlyingType": "COIN",  // 标的类型
// 			"underlyingSubType": []     // 标的物子类型
// 		}
// 	],
// 	"timezone": "UTC" // 服务器所用的时间区域
// }
type SwapExchangeInfoRes struct {
	ExchangeFilters []map[string]interface{}       `json:"exchangeFilters"`
	RateLimits      []SwapExchangeInfoResRateLimit `json:"rateLimits"`
	Symbols         []SwapExchangeInfoResSymbol    `json:"symbols"`
	Timezone        string                         `json:"timezone"` // 服务器所用的时间区域
}

type SwapExchangeInfoResRateLimit struct { // API访问的限制
	Interval      string `json:"interval"`      // 按照分钟计算
	IntervalNum   int64  `json:"intervalNum"`   // 按照1分钟计算
	Limit         int64  `json:"limit"`         // 上限次数
	RateLimitType string `json:"rateLimitType"` // 按照访问权重来计算
}

type SwapExchangeInfoResSymbol struct {
	Filters            []map[string]interface{} `json:"filters"`         // 交易对限制
	OrderType          []string                 `json:"orderType"`       // 订单类型
	TimeInForce        []string                 `json:"timeInForce"`     // 有效方式
	LiquidationFee     string                   `json:"liquidationFee"`  // 强平费率
	MarketTakeBound    string                   `json:"marketTakeBound"` // 市价吃单(相对于标记价格)允许可造成的最大价格偏离比例
	Symbol             string                   `json:"symbol"`          // 交易对
	Pair               string                   `json:"pair"`            // 标的交易对
	ContractType       string                   `json:"contractType"`    // 合约类型
	DeliveryDate       int64                    `json:"deliveryDate"`
	OnboardDate        int64                    `json:"onboardDate"`
	ContractStatus     string                   `json:"contractStatus"` // 交易对状态
	ContractSize       int64                    `json:"contractSize"`
	QuoteAsset         string                   `json:"quoteAsset"`        // 报价币种
	BaseAsset          string                   `json:"baseAsset"`         // 标的物
	MarginAsset        string                   `json:"marginAsset"`       // 保证金币种
	PricePrecision     int64                    `json:"pricePrecision"`    // 价格小数点位数(仅作为系统精度使用，注意同tickSize 区分)
	QuantityPrecision  int64                    `json:"quantityPrecision"` // 数量小数点位数(仅作为系统精度使用，注意同stepSize 区分)
	BaseAssetPrecision int64                    `json:"baseAssetPrecision"`
	QuotePrecision     int64                    `json:"quotePrecision"`
	TriggerProtect     string                   `json:"triggerProtect"`    // 开启"priceProtect"的条件订单的触发阈值
	UnderlyingType     string                   `json:"underlyingType"`    // 标的类型
	UnderlyingSubType  []string                 `json:"underlyingSubType"` // 标的物子类型
}

// [
//   {
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
//     "positionSide": "SHORT", // 持仓方向
//     "status": "NEW",                    // 订单状态
//     "stopPrice": "9300",                    // 触发价，对`TRAILING_STOP_MARKET`无效
//     "closePosition": false,             // 是否条件全平仓
//     "symbol": "BTCUSDT",                // 交易对
//     "time": 1579276756075,              // 订单时间
//     "timeInForce": "GTC",               // 有效方法
//     "type": "TRAILING_STOP_MARKET",     // 订单类型
//     "activatePrice": "9020", // 跟踪止损激活价格, 仅`TRAILING_STOP_MARKET` 订单返回此字段
//     "priceRate": "0.3", // 跟踪止损回调比例, 仅`TRAILING_STOP_MARKET` 订单返回此字段
//     "updateTime": 1579276756075,        // 更新时间
//     "workingType": "CONTRACT_PRICE", // 条件价格触发类型
//     "priceProtect": false            // 是否开启条件单触发保护
//   }
// ]

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

type SwapOrderGetRes SwapOrderOrder

type SwapOrderDeleteRes SwapOrderOrder

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

type SwapTickerPriceRes []SwapTickerPriceResRow
type SwapTickerPriceResRow struct {
	Symbol string `json:"symbol"` // 交易对
	Ps     string `json:"ps"`     // 标的交易对
	Price  string `json:"price"`  // 价格
	Time   int64  `json:"time"`   // 时间
}

// {
// 	"lastUpdateId": 16769853,
// 	"symbol": "BTCUSD_PERP", // 交易对
// 	"pair": "BTCUSD",      // 标的交易对
// 	"E": 1591250106370,   // 消息时间
// 	"T": 1591250106368,   // 撮合时间
// 	"bids": [              // 买单
// 	  [
// 		"9638.0",         // 价格
// 		"431"             // 数量
// 	  ]
// 	],
// 	"asks": [             // 卖单
// 	  [
// 		"9638.2",         // 价格
// 		"12"              // 数量
// 	  ]
// 	]
//   }
type SwapDepthRes struct {
	LastUpdateId int64       `json:"lastUpdateId"`
	Symbol       string      `json:"symbol"` // 交易对
	Pair         string      `json:"pair"`   // 标的交易对
	E            int64       `json:"E"`      // 消息时间
	T            int64       `json:"T"`      // 撮合时间
	Bids         []DepthGear `json:"bids"`
	Asks         []DepthGear `json:"asks"`
}

type SwapDepthResMiddle struct {
	LastUpdateId int64           `json:"lastUpdateId"`
	Symbol       string          `json:"symbol"` // 交易对
	Pair         string          `json:"pair"`   // 标的交易对
	E            int64           `json:"E"`      // 消息时间
	T            int64           `json:"T"`      // 撮合时间
	Bids         [][]interface{} `json:"bids"`
	Asks         [][]interface{} `json:"asks"`
}

func (middle *SwapDepthResMiddle) ConvertToRes() *SwapDepthRes {
	res := SwapDepthRes{}
	res.LastUpdateId = middle.LastUpdateId
	res.Symbol = middle.Symbol
	res.Pair = middle.Pair
	res.E = middle.E
	res.T = middle.T
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
