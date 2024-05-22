package mybinanceapi

type FutureCommonPostRes struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
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
