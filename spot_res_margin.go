package mybinanceapi

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

type MarginMaxBorrowableRes struct {
	Amount      string `json:"amount"`      //系统可借充足情况下用户账户当前最大可借额度
	BorrowLimit string `json:"borrowLimit"` //平台限制的用户当前等级可以借的额度
}

type MarginMaxTransferableRes struct {
	Amount string `json:"amount"` //用户账户当前最大可转出额度
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
type MarginInterestHistoryRes struct {
	Rows  []MarginInterestHistoryRow `json:"rows"`
	Total int                        `json:"total"`
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

type SpotMarginOrderDeleteRes struct {
	Symbol                  string `json:"symbol"`                  // 交易对
	OrderId                 string `json:"orderId"`                 // 系统的订单ID
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

type MarginAllOrdersRes []MarginOrderOrder

type MarginOpenOrdersRes []MarginOrderOrder

type TranRes struct {
	TranId int64 `json:"tranId"` //transaction id
}
type MarginTransferRes TranRes

type MarginIsolatedTransferRes TranRes

type MarginLoanRes TranRes

type MarginRepayRes TranRes
