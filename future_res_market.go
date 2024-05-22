package mybinanceapi

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

type FutureFundingRateRes []FutureFundingRateResRow
type FutureFundingRateResRow struct {
	Symbol      string `json:"symbol"`      // 交易对
	FundingTime int64  `json:"fundingTime"` // 资金费率
	FundingRate string `json:"fundingRate"` // 资金费时间
	MarkPrice   string `json:"markPrice"`   // 资金费对应标记价格
}

type FutureFundingInfoRes []FutureFundingInfoResRow
type FutureFundingInfoResRow struct {
	Symbol                   string `json:"symbol"`                   // 交易对
	AdjustedFundingRateCap   string `json:"adjustedFundingRateCap"`   // 资金费率上限
	AdjustedFundingRateFloor string `json:"adjustedFundingRateFloor"` // 资金费率下限
	FundingIntervalHours     int64  `json:"fundingIntervalHours"`     // 资金费率间隔时间
	Disclaimer               bool   `json:"disclaimer"`
}

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

type FutureTickerPriceRes []FutureTickerPriceResRow
type FutureTickerPriceResRow struct {
	Symbol string `json:"symbol"` // 交易对
	Price  string `json:"price"`  // 价格
	Time   int64  `json:"time"`   // 撮合引擎时间
}

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
