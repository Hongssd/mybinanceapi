package mybinanceapi

type SwapPingRes struct {
}

type SwapServerTimeRes struct {
	ServerTime int64 `json:"serverTime"`
}

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
