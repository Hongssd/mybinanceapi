package mybinanceapi

import "github.com/shopspring/decimal"

type KlinesMiddleRow [11]interface{}

type KlinesMiddle []KlinesMiddleRow

type KlinesResRow struct {
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
type KlinesRes []KlinesResRow

func (klinesMiddle *KlinesMiddle) ConvertToRes() KlinesRes {
	klines := []KlinesResRow{}
	for _, row := range *klinesMiddle {
		k := KlinesResRow{}
		k.StartTime = int64(row[0].(float64))
		k.Open = interfaceStringToFloat64(row[1])
		k.High = interfaceStringToFloat64(row[2])
		k.Low = interfaceStringToFloat64(row[3])
		k.Close = interfaceStringToFloat64(row[4])
		k.Volume = interfaceStringToFloat64(row[5])
		k.CloseTime = int64(row[6].(float64))
		k.TransactionVolume = interfaceStringToFloat64(row[7])
		k.TransactionNumber = int64(row[8].(float64))
		k.BuyTransactionVolume = interfaceStringToFloat64(row[9])
		k.BuyTransactionAmount = interfaceStringToFloat64(row[10])
		klines = append(klines, k)
	}
	return klines
}

type DepthGear struct {
	Price    string `json:"price"`
	Quantity string `json:"quantity"`
}

func (gear *DepthGear) ParseDecimal() (decimal.Decimal, decimal.Decimal, error) {
	price, err := decimal.NewFromString(gear.Price)
	if err != nil {
		return decimal.Zero, decimal.Zero, err
	}
	quantity, err := decimal.NewFromString(gear.Quantity)
	if err != nil {
		return decimal.Zero, decimal.Zero, err
	}
	return price, quantity, nil
}
