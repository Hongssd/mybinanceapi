package mybinanceapi

type SwapTickerPriceRes []SwapTickerPriceResRow
type SwapTickerPriceResRow struct {
	Symbol string `json:"symbol"` // 交易对
	Ps     string `json:"ps"`     // 标的交易对
	Price  string `json:"price"`  // 价格
	Time   int64  `json:"time"`   // 时间
}

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
