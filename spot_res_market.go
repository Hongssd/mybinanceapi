package mybinanceapi

type SpotTickerPriceResRow struct {
	Symbol string `json:"symbol"` // 交易对
	Price  string `json:"price"`  // 价格
}
type SpotTickerPriceRes []SpotTickerPriceResRow

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

type SpotAvgPriceRes struct {
	Mins  int64  `json:"mins"`  // 价格平均值计算时间
	Price string `json:"price"` // 价格平均值
}

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

type SpotTickerBookTickerRes []SpotTickerBookTickerResRow
type SpotTickerBookTickerResRow struct {
	Symbol   string `json:"symbol"`   // 交易对
	BidPrice string `json:"bidPrice"` // 当前最高买价
	BidQty   string `json:"bidQty"`   // 当前最高买价对应的量
	AskPrice string `json:"askPrice"` // 当前最低卖价
	AskQty   string `json:"askQty"`   // 当前最低卖价对应的量
}

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
