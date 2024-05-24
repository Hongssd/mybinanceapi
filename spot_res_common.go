package mybinanceapi

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
