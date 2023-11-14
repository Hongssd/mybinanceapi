package mybinanceapi

type WsKline struct {
	Version              float64 `json:"version"`
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

func HandleWsCombinedKlineGzip(apiType ApiType, data []byte) (*WsKline, error) {
	return HandleWsKline(apiType, data)
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
		Version:              float64(1.0),
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
