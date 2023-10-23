package binanceservice

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
