package mybinanceapi

import (
	"strings"
	"time"
)

type WsKline struct {
	Timestamp            int64   `json:"timestamp"`
	AccountType          string  `json:"account_type"`                 //K线类型 现货:spot 币合约:swap u合约:future
	Symbol               string  `json:"symbol"`                       //交易对
	Interval             string  `json:"interval"`                     //K线间隔
	StartTime            int64   `json:"start_time" gorm:"primaryKey"` //开盘时间
	Open                 float64 `json:"open"`                         //开盘价
	High                 float64 `json:"high"`                         //最高价
	Low                  float64 `json:"low"`                          //最低价
	Close                float64 `json:"close"`                        //收盘价
	Volume               float64 `json:"volume"`                       //成交量
	CloseTime            int64   `json:"close_time"`                   //收盘时间
	TransactionVolume    float64 `json:"transaction_volume"`           // 成交额
	TransactionNumber    int64   `json:"transaction_number"`           //成交笔数
	BuyTransactionVolume float64 `json:"buy_transaction_volume"`       //主动买入成交量
	BuyTransactionAmount float64 `json:"buy_transaction_amount"`       //主动买入成交额
}

func HandleWsCombinedKline(apiType ApiType, data []byte) (*WsKline, error) {
	_, ddata, _ := handlerWsCombined(data)
	return HandleWsKline(apiType, ddata)
}

func handlerWsCombined(data []byte) (string, []byte, error) {
	all := make(map[string]interface{})
	err := json.Unmarshal(data, &all)
	if err != nil {
		log.Error(err)
		return "", nil, err
	}
	mdata := (all["data"]).(map[string]interface{})
	ddata, err := json.Marshal(mdata)
	if err != nil {
		log.Error(err)
		return "", nil, err
	}
	stream := all["stream"].(string)
	return stream, ddata, nil
}

func HandleWsKline(apiType ApiType, data []byte) (*WsKline, error) {

	all := make(map[string]interface{})
	err := json.Unmarshal(data, &all)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	timestamp := interfaceStringToInt64(all["E"])
	m := all["k"].(map[string]interface{})
	myKline := HandleWsKlineMap(apiType, m)
	myKline.Timestamp = timestamp
	return myKline, nil
}

func HandleWsKlineMap(apiType ApiType, m map[string]interface{}) *WsKline {
	myKline := WsKline{
		AccountType:          apiType.String(),
		Symbol:               m["s"].(string),
		Interval:             m["i"].(string),
		StartTime:            int64(m["t"].(float64)),
		Open:                 interfaceStringToFloat64(m["o"]),
		High:                 interfaceStringToFloat64(m["h"]),
		Low:                  interfaceStringToFloat64(m["l"]),
		Close:                interfaceStringToFloat64(m["c"]),
		Volume:               interfaceStringToFloat64(m["v"]),
		CloseTime:            int64(m["T"].(float64)),
		TransactionVolume:    interfaceStringToFloat64(m["q"]),
		TransactionNumber:    int64(m["n"].(float64)),
		BuyTransactionVolume: interfaceStringToFloat64(m["V"]),
		BuyTransactionAmount: interfaceStringToFloat64(m["Q"]),
	}
	return &myKline
}

type WsDepth struct {
	Timestamp    int64        `json:"timestamp"`
	Level        string       `json:"level"`
	USpeed       string       `json:"u_speed"`
	AccountType  string       `json:"account_type"`
	Symbol       string       `json:"symbol"`
	UpperU       int64        `json:"upper_u"`
	LowerU       int64        `json:"lower_u"`
	PreU         int64        `json:"pre_u"`
	LastUpdateID int64        `json:"last_update_id"`
	Bids         []PriceLevel `json:"bids"`
	Asks         []PriceLevel `json:"asks"`
}
type PriceLevel struct {
	Price    float64 `json:"price"`
	Quantity float64 `json:"quantity"`
}

func HandleWsCombinedDepth(apiType ApiType, data []byte) (*WsDepth, error) {

	streams, ddata, _ := handlerWsCombined(data)
	split := strings.Split(streams, "@")
	symbol := strings.ToUpper(split[0])

	var isLevelDepth bool
	if split[1] == "depth" {
		isLevelDepth = false
	} else {
		isLevelDepth = true
	}
	if isLevelDepth {
		Level := ""
		uSpeed := ""
		if len(split) >= 2 {
			Level = split[1][5:len(split[1])]
		}
		if len(split) == 3 {
			uSpeed = split[2]
		}
		d, err := HandleDepth(apiType, symbol, ddata, isLevelDepth)
		if err != nil {
			return d, err
		}
		d.USpeed = uSpeed
		d.Level = Level
		return d, err
	} else {
		uSpeed := ""
		if len(split) == 3 {
			uSpeed = split[2]
		}
		d, err := HandleDepth(apiType, symbol, ddata, isLevelDepth)
		if err != nil {
			return d, err
		}
		d.USpeed = uSpeed
		return d, err
	}

}

func HandleDepth(apiType ApiType, symbol string, data []byte, isLeveDepth bool) (*WsDepth, error) {

	all := make(map[string]interface{})
	err := json.Unmarshal(data, &all)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	timestamp := time.Now().UnixMilli()
	bidsKey := "b"
	asksKey := "a"
	lastUpdateIdKey := "u"

	if isLeveDepth && apiType == SPOT {
		bidsKey = "bids"
		asksKey = "asks"
		lastUpdateIdKey = "lastUpdateId"
	} else {
		timestamp = interfaceStringToInt64(all["E"])
	}

	bids := []PriceLevel{}
	asks := []PriceLevel{}

	bidsArray := all[bidsKey].([]interface{})
	asksArray := all[asksKey].([]interface{})

	for _, b := range bidsArray {
		bid := PriceLevel{
			Price:    interfaceStringToFloat64((b.([]interface{}))[0]),
			Quantity: interfaceStringToFloat64((b.([]interface{}))[1]),
		}
		bids = append(bids, bid)
	}

	for _, a := range asksArray {
		ask := PriceLevel{
			Price:    interfaceStringToFloat64((a.([]interface{}))[0]),
			Quantity: interfaceStringToFloat64((a.([]interface{}))[1]),
		}
		asks = append(asks, ask)
	}
	var upperU, lowerU, preU int64

	if isLeveDepth && apiType == SPOT {
		upperU, lowerU, preU = 0, 0, 0
	} else {
		switch apiType {
		case SPOT:
			upperU = interfaceStringToInt64(all["U"])
			lowerU = interfaceStringToInt64(all["u"])
		case FUTURE, SWAP:
			upperU = interfaceStringToInt64(all["U"])
			lowerU = interfaceStringToInt64(all["u"])
			preU = interfaceStringToInt64(all["pu"])
		default:
			upperU, lowerU, preU = 0, 0, 0
		}
	}

	myDepth := WsDepth{
		UpperU:       upperU,
		LowerU:       lowerU,
		PreU:         preU,
		LastUpdateID: int64(all[lastUpdateIdKey].(float64)),
		AccountType:  apiType.String(),
		Symbol:       symbol,
		Timestamp:    timestamp,
		Bids:         bids,
		Asks:         asks,
	}
	return &myDepth, nil
}

type WsAggTrade struct {
	Timestamp   int64   `json:"timestamp"`                //事件时间
	Symbol      string  `json:"symbol"`                   //交易对
	AccountType string  `json:"account_type"`             //账户类型 SPOT 现货 FUTURE U合约 SWAP 币本位合约
	AId         int64   `gorm:"primaryKey " json:"a_id" ` // 归集交易ID
	Price       float64 `json:"price"`                    //成交价
	Quantity    float64 `json:"quantity"`                 //成交量
	First       int64   `json:"first"`                    //被归集的首个交易ID
	Last        int64   `json:"last"`                     //被归集的末次交易ID
	TradeTime   int64   `json:"trade_time"`               //成交时间
	IsMarket    bool    `json:"is_market"`                //买方是否做市
}

func HandleWsCombinedAggTrade(apiType ApiType, data []byte) (*WsAggTrade, error) {
	_, ddata, _ := handlerWsCombined(data)
	return HandleWsAggTrade(apiType, ddata)
}
func HandleWsAggTrade(apiType ApiType, data []byte) (*WsAggTrade, error) {
	all := make(map[string]interface{})
	err := json.Unmarshal(data, &all)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	myAggTrade := HandleWsAggTradeMap(apiType, all)
	return myAggTrade, nil
}
func HandleWsAggTradeMap(apiType ApiType, m map[string]interface{}) *WsAggTrade {
	myAggTrade := WsAggTrade{
		Timestamp:   interfaceStringToInt64(m["E"].(float64)),
		AccountType: apiType.String(),
		Symbol:      m["s"].(string),
		AId:         interfaceStringToInt64(m["a"].(float64)),
		Price:       interfaceStringToFloat64(m["p"]),
		Quantity:    interfaceStringToFloat64(m["q"]),
		First:       interfaceStringToInt64(m["f"]),
		Last:        interfaceStringToInt64(m["l"]),
		TradeTime:   interfaceStringToInt64(m["T"]),
		IsMarket:    m["m"].(bool),
	}
	return &myAggTrade
}
