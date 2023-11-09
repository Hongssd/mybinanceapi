package mybinanceapi

type SubAccountResSubAccount struct {
	Email                       string `json:"email"`
	IsFreeze                    bool   `json:"isFreeze"`
	CreateTime                  int64  `json:"createTime"`
	IsManagedSubAccount         bool   `json:"isManagedSubAccount"`
	IsAssetManagementSubAccount bool   `json:"isAssetManagementSubAccount"`
}
type SubAccountListRes struct {
	SubAccounts []SubAccountResSubAccount `json:"subAccounts"`
}

type SubAccountVirtualSubAccountRes struct {
	Email string `json:"email"`
}

type SubAccountUniversalTransferRes struct {
	TranId       int64  `json:"tranId"`
	ClientTranId string `json:"clientTranId"`
}

type SubAccountUniversalTransferHistoryResResult struct {
	TranId          int64  `json:"tranId"`
	FromEmail       string `json:"fromEmail"`
	ToEmail         string `json:"toEmail"`
	Asset           string `json:"asset"`
	Amount          string `json:"amount"`
	CreateTimeStamp int64  `json:"createTimeStamp"`
	FromAccountType string `json:"fromAccountType"`
	ToAccountType   string `json:"toAccountType"`
	Status          string `json:"status"`
	ClientTranId    string `json:"clientTranId"`
}
type SubAccountUniversalTransferHistoryRes struct {
	Result     []SubAccountUniversalTransferHistoryResResult `json:"result"`
	TotalCount int                                           `json:"totalCount"`
}

type SubAccountAssetsRes struct {
	Balance []SpotBalance `json:"balances"`
}

type SubAccountFuturesAccountAssets struct {
	Asset                  string `json:"asset"`
	InitialMargin          string `json:"initialMargin"`
	MaintenanceMargin      string `json:"maintenanceMargin"`
	MarginBalance          string `json:"marginBalance"`
	MaxWithdrawAmount      string `json:"maxWithdrawAmount"`
	OpenOrderInitialMargin string `json:"openOrderInitialMargin"`
	PositionInitialMargin  string `json:"positionInitialMargin"`
	UnrealizedProfit       string `json:"unrealizedProfit"`
	WalletBalance          string `json:"walletBalance"`
}
type SubAccountFuturesAccountAccountResp struct {
	Email                       string                           `json:"email"`
	Assets                      []SubAccountFuturesAccountAssets `json:"assets"`
	CanDeposit                  bool                             `json:"canDeposit"`
	CanTrade                    bool                             `json:"canTrade"`
	CanWithdraw                 bool                             `json:"canWithdraw"`
	FeeTier                     int                              `json:"feeTier"`
	UpdateTime                  int64                            `json:"updateTime"`
	MaxWithdrawAmount           string                           `json:"maxWithdrawAmount"`
	TotalInitialMargin          string                           `json:"totalInitialMargin"`
	TotalMaintenanceMargin      string                           `json:"totalMaintenanceMargin"`
	TotalMarginBalance          string                           `json:"totalMarginBalance"`
	TotalOpenOrderInitialMargin string                           `json:"totalOpenOrderInitialMargin"`
	TotalPositionInitialMargin  string                           `json:"totalPositionInitialMargin"`
	TotalUnrealizedProfit       string                           `json:"totalUnrealizedProfit"`
	TotalWalletBalance          string                           `json:"totalWalletBalance"`
}

type SubAccountFuturesAccountRes struct {
	FutureAccountResp   SubAccountFuturesAccountAccountResp `json:"futureAccountResp"`
	DeliveryAccountResp SubAccountFuturesAccountAccountResp `json:"deliveryAccountResp"`
}

type SubAccountFuturesEnableRes struct {
	Email            string `json:"email"`
	IsFuturesEnabled bool   `json:"isFuturesEnabled"`
}

type SpotAccountCommissionRates struct {
	Maker  string `json:"maker"`
	Taker  string `json:"taker"`
	Buyer  string `json:"buyer"`
	Seller string `json:"seller"`
}

type SpotBalance struct {
	Asset  string `json:"asset"`
	Free   string `json:"free"`
	Locked string `json:"locked"`
}
type SpotAccountRes struct {
	MakerCommission           int                        `json:"makerCommission"`
	TakerCommission           int                        `json:"takerCommission"`
	BuyerCommission           int                        `json:"buyerCommission"`
	SellerCommission          int                        `json:"sellerCommission"`
	CommissionRates           SpotAccountCommissionRates `json:"commissionRates"`
	CanTrade                  bool                       `json:"canTrade"`
	CanWithdraw               bool                       `json:"canWithdraw"`
	CanDeposit                bool                       `json:"canDeposit"`
	Brokered                  bool                       `json:"brokered"`
	RequireSelfTradePreventio bool                       `json:"requireSelfTradePreventio"`
	UpdateTime                int64                      `json:"updateTime"`
	AccountType               string                     `json:"accountType"`
	Balance                   []SpotBalance              `json:"balances"`
	Permissions               []string                   `json:"permissions"`
}

type SpotSubAccountApiIpRestrictionRes struct {
	IpRestrict string   `json:"ipRestrict"`
	IpList     []string `json:"ipList"`
	UpdateTime int64    `json:"updateTime"`
	ApiKey     string   `json:"apiKey"`
}

type MarginPairs struct {
	Base          string `json:"base"`          //基础币种
	Id            uint64 `json:"id"`            //币种id
	IsBuyAllowed  bool   `json:"isBuyAllowed"`  //是否允许买入
	IsMarginTrade bool   `json:"isMarginTrade"` //是否允许杠杆交易
	IsSellAllowed bool   `json:"isSellAllowed"` //是否允许卖出
	Quote         string `json:"quote"`         //计价币种
	Symbol        string `json:"symbol"`        //交易对
}
type MarginAllPairsRes []MarginPairs
type MarginIsolatedAllPairsRes []MarginPairs

type MarginAsset struct {
	Asset    string `json:"asset"`
	Borrowed string `json:"borrowed"`
	Free     string `json:"free"`
	Interest string `json:"interest"`
	Locked   string `json:"locked"`
	NetAsset string `json:"netAsset"`
}
type MarginAccountRes struct {
	BorrowEnabled       bool          `json:"borrowEnabled"`
	MarginLevel         string        `json:"marginLevel"`
	TotalAssetOfBtc     string        `json:"totalAssetOfBtc"`
	TotalLiabilityOfBtc string        `json:"totalLiabilityOfBtc"`
	TotalNetAssetOfBtc  string        `json:"totalNetAssetOfBtc"`
	TradeEnabled        bool          `json:"tradeEnabled"`
	TransferEnabled     bool          `json:"transferEnabled"`
	UserAssets          []MarginAsset `json:"userAssets"`
}
type MarginIsolatedAsset struct {
	BaseAsset         MarginIsolatedInnerAsset `json:"baseAsset"`
	QuoteAsset        MarginIsolatedInnerAsset `json:"quoteAsset"`
	Symbol            string                   `json:"symbol"`
	IsolatedCreated   bool                     `json:"isolatedCreated"`
	Enabled           bool                     `json:"enabled"`
	MarginLevel       string                   `json:"marginLevel"`
	MarginLevelStatus string                   `json:"marginLevelStatus"`
	MarginRatio       string                   `json:"marginRatio"`
	IndexPrice        string                   `json:"indexPrice"`
	LiquidatePrice    string                   `json:"liquidatePrice"`
	LiquidateRate     string                   `json:"liquidateRate"`
	TradeEnabled      bool                     `json:"tradeEnabled"`
}
type MarginIsolatedInnerAsset struct {
	Asset         string `json:"asset"`
	BorrowEnabled bool   `json:"borrowEnabled"`
	Borrowed      string `json:"borrowed"`
	Free          string `json:"free"`
	Interest      string `json:"interest"`
	Locked        string `json:"locked"`
	NetAsset      string `json:"netAsset"`
	NetAssetOfBtc string `json:"netAssetOfBtc"`
	RepayEnabled  bool   `json:"repayEnabled"`
	TotalAsset    string `json:"totalAsset"`
}
type MarginIsolatedAccountRes struct {
	Assets              []MarginIsolatedAsset `json:"assets"`
	TotalAssetOfBtc     string                `json:"totalAssetOfBtc"`
	TotalLiabilityOfBtc string                `json:"totalLiabilityOfBtc"`
	TotalNetAssetOfBtc  string                `json:"totalNetAssetOfBtc"`
}

type TranRes struct {
	TranId int64 `json:"tranId"` //transaction id
}
type MarginTransferRes TranRes

type MarginIsolatedTransferRes TranRes

type MarginLoanRes TranRes

type MarginRepayRes TranRes

type MarginMaxBorrowableRes struct {
	Amount      string `json:"amount"`      //系统可借充足情况下用户账户当前最大可借额度
	BorrowLimit string `json:"borrowLimit"` //平台限制的用户当前等级可以借的额度
}

type MarginMaxTransferableRes struct {
	Amount string `json:"amount"` //用户账户当前最大可转出额度
}

type MarginInterestHistoryRes struct {
	Rows  []MarginInterestHistoryRow `json:"rows"`
	Total int                        `json:"total"`
}

type MarginInterestHistoryRow struct {
	TxId                int64  `json:"txId"`
	InterestAccuredTime int64  `json:"interestAccuredTime"`
	Asset               string `json:"asset"`
	RawAsset            string `json:"rawAsset"`
	Principal           string `json:"principal"`
	Interest            string `json:"interest"`
	InterestRate        string `json:"interestRate"`
	Type                string `json:"type"`
	IsolatedSymbol      string `json:"isolatedSymbol"`
}

type MarginOrderOrder struct {
	ClientOrderId           string `json:"clientOrderId"`
	CummulativeQuoteQty     string `json:"cummulativeQuoteQty"`
	ExecutedQty             string `json:"executedQty"`
	IcebergQty              string `json:"icebergQty"`
	IsWorking               bool   `json:"isWorking"`
	OrderId                 int64  `json:"orderId"`
	OrigQty                 string `json:"origQty"`
	Price                   string `json:"price"`
	Side                    string `json:"side"`
	Status                  string `json:"status"`
	StopPrice               string `json:"stopPrice"`
	Symbol                  string `json:"symbol"`
	IsIsolated              bool   `json:"isIsolated"`
	Time                    int64  `json:"time"`
	TimeInForce             string `json:"timeInForce"`
	Type                    string `json:"type"`
	SelfTradePreventionMode string `json:"selfTradePreventionMode"`
	UpdateTime              int64  `json:"updateTime"`
}

type MarginOrderGetRes MarginOrderOrder

type MarginAllOrdersRes []MarginOrderOrder

type MarginOpenOrdersRes []MarginOrderOrder

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

type SpotOrderFill struct {
	Price           string `json:"price"`           // 交易的价格
	Qty             string `json:"qty"`             // 交易的数量
	Commission      string `json:"commission"`      // 手续费金额
	CommissionAsset string `json:"commissionAsset"` // 手续费的币种
	TradeId         int64  `json:"tradeId"`         // 交易ID
}

//	{
//		"symbol": "BTCUSDT",
//		"orderId": 28,
//		"clientOrderId": "6gCrw2kRUAF9CvJDGP16IP",
//		"transactTime": 1507725176595,
//		"price": "0.00000000",
//		"origQty": "10.00000000",
//		"executedQty": "10.00000000",
//		"cummulativeQuoteQty": "10.00000000",
//		"status": "FILLED",
//		"timeInForce": "GTC",
//		"type": "MARKET",
//		"side": "SELL",
//		"marginBuyBorrowAmount": 5,       // 下单后没有发生借款则不返回该字段
//		"marginBuyBorrowAsset": "BTC",    // 下单后没有发生借款则不返回该字段
//		"isIsolated": true,       // 是否是逐仓symbol交易
//		"selfTradePreventionMode": "NONE",
//		"fills": [
//		  {
//			"price": "4000.00000000",
//			"qty": "1.00000000",
//			"commission": "4.00000000",
//			"commissionAsset": "USDT"
//		  },
//		  {
//			"price": "3999.00000000",
//			"qty": "5.00000000",
//			"commission": "19.99500000",
//			"commissionAsset": "USDT"
//		  },
//		  {
//			"price": "3998.00000000",
//			"qty": "2.00000000",
//			"commission": "7.99600000",
//			"commissionAsset": "USDT"
//		  },
//		  {
//			"price": "3997.00000000",
//			"qty": "1.00000000",
//			"commission": "3.99700000",
//			"commissionAsset": "USDT"
//		  },
//		  {
//			"price": "3995.00000000",
//			"qty": "1.00000000",
//			"commission": "3.99500000",
//			"commissionAsset": "USDT"
//		  }
//		]
//	  }
type SpotMarginOrderPostRes struct {
	Symbol                  string          `json:"symbol"`                  // 交易对
	OrderId                 int64           `json:"orderId"`                 // 系统的订单ID
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
	MarginBuyBorrowAmount   string          `json:"marginBuyBorrowAmount"`   // 下单后没有发生借款则不返回该字段
	MarginBuyBorrowAsset    string          `json:"marginBuyBorrowAsset"`    // 下单后没有发生借款则不返回该字段
	IsIsolated              bool            `json:"isIsolated"`              // 是否是逐仓symbol交易
	SelfTradePreventionMode string          `json:"selfTradePreventionMode"` // 自我交易预防模式
	Fills                   []SpotOrderFill `json:"fills"`                   // 订单中交易的信息
}

type SpotPingRes struct {
}

type SpotServerTimeRes struct {
	ServerTime int64 `json:"serverTime"`
}

type SpotExchangeInfoRes struct {
	Timezone        string                         `json:"timezone"`
	ServerTime      int64                          `json:"serverTime"`
	RateLimits      []SpotExchangeInfoResRateLimit `json:"rateLimits"`
	ExchangeFilters []map[string]interface{}       `json:"exchangeFilters"`
	Symbols         []SpotExchangeInfoResSymbol    `json:"symbols"`
}
type SpotExchangeInfoResRateLimit struct {
	Interval      string `json:"interval"`      // 限制的时间间隔
	IntervalNum   int64  `json:"intervalNum"`   // 限制的时间间隔数量
	Limit         int64  `json:"limit"`         // 限制的请求数量
	RateLimitType string `json:"rateLimitType"` // 限制的请求数量的类型
}
type SpotExchangeInfoResSymbol struct {
	Symbol                          string                   `json:"symbol"`
	Status                          string                   `json:"status"`
	BaseAsset                       string                   `json:"baseAsset"`
	BaseAssetPrecision              int64                    `json:"baseAssetPrecision"`
	QuoteAsset                      string                   `json:"quoteAsset"`
	QuotePrecision                  int64                    `json:"quotePrecision"`
	QuoteAssetPrecision             int64                    `json:"quoteAssetPrecision"`
	OrderTypes                      []string                 `json:"orderTypes"`
	IcebergAllowed                  bool                     `json:"icebergAllowed"`
	OcoAllowed                      bool                     `json:"ocoAllowed"`
	QuoteOrderQtyMarketAllowed      bool                     `json:"quoteOrderQtyMarketAllowed"`
	AllowTrailingStop               bool                     `json:"allowTrailingStop"`
	IsSpotTradingAllowed            bool                     `json:"isSpotTradingAllowed"`
	IsMarginTradingAllowed          bool                     `json:"isMarginTradingAllowed"`
	CancelReplaceAllowed            bool                     `json:"cancelReplaceAllowed"`
	Filters                         []map[string]interface{} `json:"filters"`
	Permissions                     []string                 `json:"permissions"`
	DefaultSelfTradePreventionMode  string                   `json:"defaultSelfTradePreventionMode"`
	AllowedSelfTradePreventionModes []string                 `json:"allowedSelfTradePreventionModes"`
}

type SpotAccountApiTradingStatusRes struct {
	Data SpotAccountApiTradingStatusData `json:"data"`
}

type SpotAccountApiTradingStatusData struct {
	IsLocked           bool                                        `json:"isLocked"`
	PlannedRecoverTime int64                                       `json:"plannedRecoverTime"`
	TriggerCondition   SpotAccountApiTradingStatusTriggerCondition `json:"triggerCondition"`
	UpdateTime         int64                                       `json:"updateTime"`
}

type SpotAccountApiTradingStatusTriggerCondition struct {
	GCR  int64 `json:"GCR"`
	IFER int64 `json:"IFER"`
	UFR  int64 `json:"UFR"`
}

//	{
//	    "symbol": "LTCBTC",
//	    "orderId": 1,
//	    "orderListId": -1, // OCO订单ID，否则为 -1
//	    "clientOrderId": "myOrder1",
//	    "price": "0.1",
//	    "origQty": "1.0",
//	    "executedQty": "0.0",
//	    "cummulativeQuoteQty": "0.0",
//	    "status": "NEW",
//	    "timeInForce": "GTC",
//	    "type": "LIMIT",
//	    "side": "BUY",
//	    "stopPrice": "0.0",
//	    "icebergQty": "0.0",
//	    "time": 1499827319559,
//	    "updateTime": 1499827319559,
//	    "isWorking": true,
//	    "workingTime": 1499827319559,
//	    "origQuoteOrderQty": "0.000000",
//	    "selfTradePreventionMode": "NONE"
//	  }
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

type SpotTickerPriceRes struct {
	Symbol string `json:"symbol"` // 交易对
	Price  string `json:"price"`  // 价格
}

type SpotAssetTransferPostRes struct {
	TranId int64 `json:"tranId"`
}

// {
//     "total":2,
//     "rows":[
//         {
//             "asset":"USDT",
//             "amount":"1",
//             "type":"MAIN_UMFUTURE"
//             "status": "CONFIRMED", // status: CONFIRMED / FAILED / PENDING
//             "tranId": 11415955596,
//             "timestamp":1544433328000
//         },
//         {
//             "asset":"USDT",
//             "amount":"2",
//             "type":"MAIN_UMFUTURE",
//             "status": "CONFIRMED",
//             "tranId": 11366865406,
//             "timestamp":1544433328000
//         }
//     ]
// }

type SpotAssetTransferGetRes struct {
	Total int                       `json:"total"`
	Rows  []SpotAssetTransferGetRow `json:"rows"`
}

type SpotAssetTransferGetRow struct {
	Asset     string `json:"asset"`
	Amount    string `json:"amount"`
	Type      string `json:"type"`
	Status    string `json:"status"`
	TranId    int64  `json:"tranId"`
	Timestamp int64  `json:"timestamp"`
}

type SpotSubAccountTransferSubUserHistoryRes []SpotSubAccountTransferSubUserHistoryResRow
type SpotSubAccountTransferSubUserHistoryResRow struct {
	CounterParty    string `json:"counterParty"`
	Email           string `json:"email"`
	Type            int64  `json:"type"` // 1 for transfer in , 2 for transfer out
	Asset           string `json:"asset"`
	Qty             string `json:"qty"`
	FromAccountType string `json:"fromAccountType"`
	ToAccountType   string `json:"toAccountType"`
	Status          string `json:"status"` // status: PROCESS / SUCCESS / FAILURE
	TranId          int64  `json:"tranId"`
	Time            int64  `json:"time"`
}

// {
// 	"lastUpdateId": 1027024,
// 	"bids": [
// 	  [
// 		"4.00000000",     // 价位
// 		"431.00000000"    // 挂单量
// 	  ]
// 	],
// 	"asks": [
// 	  [
// 		"4.00000200",
// 		"12.00000000"
// 	  ]
// 	]
//   }

type SpotDepthRes struct {
	LastUpdateId int64       `json:"lastUpdateId"`
	Bids         []DepthGear `json:"bids"`
	Asks         []DepthGear `json:"asks"`
}

type SpotDepthResMiddle struct {
	LastUpdateId int64           `json:"lastUpdateId"`
	Bids         [][]interface{} `json:"bids"`
	Asks         [][]interface{} `json:"asks"`
}

func (middle *SpotDepthResMiddle) ConvertToRes() *SpotDepthRes {
	res := SpotDepthRes{}
	res.LastUpdateId = middle.LastUpdateId
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

//	{
//		"symbol": "LTCBTC",
//		"origClientOrderId": "myOrder1",
//		"orderId": 4,
//		"orderListId": -1, // OCO订单ID，否则为 -1
//		"clientOrderId": "cancelMyOrder1",
//		"transactTime": 1684804350068,
//		"price": "2.00000000",
//		"origQty": "1.00000000",
//		"executedQty": "0.00000000",
//		"cummulativeQuoteQty": "0.00000000",
//		"status": "CANCELED",
//		"timeInForce": "GTC",
//		"type": "LIMIT",
//		"side": "BUY",
//		"selfTradePreventionMode": "NONE"
//	  }
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

// {
// 	"symbol": "LTCBTC",
// 	"orderId": "28",
// 	"origClientOrderId": "myOrder1",
// 	"clientOrderId": "cancelMyOrder1",
// 	"price": "1.00000000",
// 	"origQty": "10.00000000",
// 	"executedQty": "8.00000000",
// 	"cummulativeQuoteQty": "8.00000000",
// 	"status": "CANCELED",
// 	"timeInForce": "GTC",
// 	"type": "LIMIT",
// 	"side": "SELL",
// 	"isIsolated": true       // 是否是逐仓symbol交易
//   }

type SpotMarginOrderDeleteRes struct {
	Symbol                  string `json:"symbol"`                  // 交易对
	OrderId                 int64  `json:"orderId"`                 // 系统的订单ID
	OrigClientOrderId       string `json:"origClientOrderId"`       // 原始的客户端订单ID
	ClientOrderId           string `json:"clientOrderId"`           // 客户自己设置的ID
	Price                   string `json:"price"`                   // 订单价格
	OrigQty                 string `json:"origQty"`                 // 用户设置的原始订单数量
	ExecutedQty             string `json:"executedQty"`             // 交易的订单数量
	CummulativeQuoteQty     string `json:"cummulativeQuoteQty"`     // 累计交易的金额
	Status                  string `json:"status"`                  // 订单状态
	TimeInForce             string `json:"timeInForce"`             // 订单的时效方式
	Type                    string `json:"type"`                    // 订单类型， 比如市价单，现价单等
	Side                    string `json:"side"`                    // 订单方向，买还是卖
	IsIsolated              bool   `json:"isIsolated"`              // 是否是逐仓symbol交易
	SelfTradePreventionMode string `json:"selfTradePreventionMode"` // 自我交易预防模式
}

// [
//   {
//     "id": 28457,
//     "price": "4.00000100",
//     "qty": "12.00000000",
//     "time": 1499865549590, // 交易成交时间, 和websocket中的T一致.
//     "isBuyerMaker": true,
//     "isBestMatch": true
//   }
// ]

type SpotTradesRes []SpotTradesResRow
type SpotTradesResRow struct {
	Id           int64  `json:"id"`           // 交易ID
	Price        string `json:"price"`        // 交易价格
	Qty          string `json:"qty"`          // 交易数量
	Time         int64  `json:"time"`         // 交易成交时间, 和websocket中的T一致.
	IsBuyerMaker bool   `json:"isBuyerMaker"` // 是否是买方成交
	IsBestMatch  bool   `json:"isBestMatch"`  // 是否是最优成交
}

type SpotHistoricalTradesRes []SpotTradesResRow

// [
//   {
//     "a": 26129,         // 归集成交ID
//     "p": "0.01633102",  // 成交价
//     "q": "4.70443515",  // 成交量
//     "f": 27781,         // 被归集的首个成交ID
//     "l": 27781,         // 被归集的末个成交ID
//     "T": 1498793709153, // 成交时间
//     "m": true,          // 是否为主动卖出单
//     "M": true           // 是否为最优撮合单(可忽略，目前总为最优撮合)
//   }
// ]

type SpotAggTradesRes []SpotAggTradesResRow
type SpotAggTradesResRow struct {
	Id           int64  `json:"a"` // 归集成交ID
	Price        string `json:"p"` // 成交价格
	Qty          string `json:"q"` // 成交数量
	FirstTradeId int64  `json:"f"` // 被归集的首个成交ID
	LastTradeId  int64  `json:"l"` // 被归集的末个成交ID
	Time         int64  `json:"T"` // 成交时间
	IsBuyerMaker bool   `json:"m"` // 是否为主动卖出单
	IsBestMatch  bool   `json:"M"` // 是否为最优撮合单(可忽略，目前总为最优撮合)
}

//	{
//		"mins": 5,
//		"price": "9.35751834"
//	  }
type SpotAvgPriceRes struct {
	Mins  int64  `json:"mins"`  // 价格平均值计算时间
	Price string `json:"price"` // 价格平均值
}

//	{
//		"symbol": "BNBBTC",
//		"priceChange": "-94.99999800",
//		"priceChangePercent": "-95.960",
//		"weightedAvgPrice": "0.29628482",
//		"prevClosePrice": "0.10002000",
//		"lastPrice": "4.00000200",
//		"lastQty": "200.00000000",
//		"bidPrice": "4.00000000",
//		"bidQty": "100.00000000",
//		"askPrice": "4.00000200",
//		"askQty": "100.00000000",
//		"openPrice": "99.00000000",
//		"highPrice": "100.00000000",
//		"lowPrice": "0.10000000",
//		"volume": "8913.30000000",
//		"quoteVolume": "15.30000000",
//		"openTime": 1499783499040,
//		"closeTime": 1499869899040,
//		"firstId": 28385,   // 首笔成交id
//		"lastId": 28460,    // 末笔成交id
//		"count": 76         // 成交笔数
//	  }
type SpotTicker24hrRes []SpotTicker24hrResRow
type SpotTicker24hrResRow struct {
	Symbol             string `json:"symbol"`             // 交易对
	PriceChange        string `json:"priceChange"`        // 价格变动
	PriceChangePercent string `json:"priceChangePercent"` // 价格变动百分比
	WeightedAvgPrice   string `json:"weightedAvgPrice"`   // 平均价格
	PrevClosePrice     string `json:"prevClosePrice"`     // 前一日收盘价
	LastPrice          string `json:"lastPrice"`          // 最新成交价
	LastQty            string `json:"lastQty"`            // 最新成交量
	BidPrice           string `json:"bidPrice"`           // 当前最高买价
	BidQty             string `json:"bidQty"`             // 当前最高买价对应的量
	AskPrice           string `json:"askPrice"`           // 当前最低卖价
	AskQty             string `json:"askQty"`             // 当前最低卖价对应的量
	OpenPrice          string `json:"openPrice"`          // 24小时内第一次交易的价格
	HighPrice          string `json:"highPrice"`          // 24小时内最高成交价
	LowPrice           string `json:"lowPrice"`           // 24小时内最低成交加
	Volume             string `json:"volume"`             // 24小时内成交量
	QuoteVolume        string `json:"quoteVolume"`        // 24小时内成交额
	OpenTime           int64  `json:"openTime"`           // 统计开始时间
	CloseTime          int64  `json:"closeTime"`          // 统计结束时间
	FirstId            int64  `json:"firstId"`            // 首笔成交id
	LastId             int64  `json:"lastId"`             // 末笔成交id
	Count              int64  `json:"count"`              // 成交笔数
}

// {
// 	"symbol": "LTCBTC",
// 	"bidPrice": "4.00000000Decimal
// 	"bidQty": "431.00000000",
// 	"askPrice": "4.00000200",
// 	"askQty": "9.00000000"
//   }

type SpotTickerBookTickerRes []SpotTickerBookTickerResRow
type SpotTickerBookTickerResRow struct {
	Symbol   string `json:"symbol"`   // 交易对
	BidPrice string `json:"bidPrice"` // 当前最高买价
	BidQty   string `json:"bidQty"`   // 当前最高买价对应的量
	AskPrice string `json:"askPrice"` // 当前最低卖价
	AskQty   string `json:"askQty"`   // 当前最低卖价对应的量
}

//	{
//		"symbol":             "BNBBTC",
//		"priceChange":        "-8.00000000",  // 价格变化
//		"priceChangePercent": "-88.889",      // 价格变化百分比
//		"weightedAvgPrice":   "2.60427807",
//		"openPrice":          "9.00000000",
//		"highPrice":          "9.00000000",
//		"lowPrice":           "1.00000000",
//		"lastPrice":          "1.00000000",
//		"volume":             "187.00000000",
//		"quoteVolume":        "487.00000000",
//		"openTime":           1641859200000,  // ticker的开始时间
//		"closeTime":          1642031999999,  // ticker的结束时间
//		"firstId":            0,              // 统计时间内的第一笔trade id
//		"lastId":             60,
//		"count":              61              // 统计时间内交易笔数
//	  }
type SpotTickerRes []SpotTickerResRow

type SpotTickerResRow struct {
	Symbol             string `json:"symbol"`             // 交易对
	PriceChange        string `json:"priceChange"`        // 价格变动
	PriceChangePercent string `json:"priceChangePercent"` // 价格变动百分比
	WeightedAvgPrice   string `json:"weightedAvgPrice"`   // 平均价格
	OpenPrice          string `json:"openPrice"`          // 24小时内第一次交易的价格
	HighPrice          string `json:"highPrice"`          // 24小时内最高成交价
	LowPrice           string `json:"lowPrice"`           // 24小时内最低成交加
	LastPrice          string `json:"lastPrice"`          // 最新成交价
	Volume             string `json:"volume"`             // 24小时内成交量
	QuoteVolume        string `json:"quoteVolume"`        // 24小时内成交额
	OpenTime           int64  `json:"openTime"`           // 统计开始时间
	CloseTime          int64  `json:"closeTime"`          // 统计结束时间
	FirstId            int64  `json:"firstId"`            // 首笔成交id
	LastId             int64  `json:"lastId"`             // 末笔成交id
	Count              int64  `json:"count"`              // 成交笔数
}
