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
	TransactionVolume    float64 `json:"transaction_volume"`           //成交额
	TransactionNumber    int64   `json:"transaction_number"`           //成交笔数
	BuyTransactionVolume float64 `json:"buy_transaction_volume"`       //主动买入成交量
	BuyTransactionAmount float64 `json:"buy_transaction_amount"`       //主动买入成交额
	Confirm              bool    `json:"confirm"`                      //这根K线是否完结
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
		Confirm:              m["x"].(bool),
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

// 现货账户更新推送
type WsSpotPayloadOutboundAccountPosition struct {
	Event          string          `json:"e"` //事件类型
	Timestamp      int64           `json:"E"` //事件时间
	LastUpdateTime int64           `json:"u"` //账户末次更新时间戳
	Balances       []WsSpotBalance `json:"B"` //余额
}

type WsSpotBalance struct {
	Asset  string `json:"a"` //资产名称
	Free   string `json:"f"` //可用余额
	Locked string `json:"l"` //冻结余额
}

// 现货余额更新推送
type WsSpotPayloadBalanceUpdate struct {
	Event     string `json:"e"` //事件类型
	Timestamp int64  `json:"E"` //事件时间
	Asset     string `json:"a"` //资产名称
	Delta     string `json:"d"` //余额变化
	ClearTime int64  `json:"T"` //清算时间
}

type WsPMMarginPayloadBalanceUpdate struct {
	Event     string `json:"e"` //事件类型
	Timestamp int64  `json:"E"` //事件时间
	Asset     string `json:"a"` //资产名称
	Delta     string `json:"d"` //余额变化
	UpdateId  int64  `json:"U"` //事件更新ID
	ClearTime int64  `json:"T"` //清算时间
}

// 现货订单推送
type WsSpotPayloadExecutionReport struct {
	//标准字段
	Event               string `json:"e"` //事件类型
	Timestamp           int64  `json:"E"` //事件时间
	Symbol              string `json:"s"` //交易对
	ClientOrderId       string `json:"c"` //clientOrderId
	Side                string `json:"S"` //订单方向
	Type                string `json:"o"` //订单类型
	TimeInForce         string `json:"f"` //有效方式
	OrigQty             string `json:"q"` //订单原始数量
	Price               string `json:"p"` //订单原始价格
	StopPrice           string `json:"P"` //止盈止损单触发价格
	IcebergQty          string `json:"F"` //冰山订单数量
	OrderListId         int64  `json:"g"` //OCO订单 OrderListId
	OrigClientOrderId   string `json:"C"` //原始订单自定义ID(原始订单，指撤单操作的对象。撤单本身被视为另一个订单)
	ExecutionType       string `json:"x"` //本次事件的具体执行类型
	Status              string `json:"X"` //订单的当前状态
	RejectReason        string `json:"r"` //订单被拒绝的原因
	OrderId             int64  `json:"i"` //orderId
	LastQty             string `json:"l"` //订单末次成交量
	ExecutedQty         string `json:"z"` //订单累计已成交量
	LastPrice           string `json:"L"` //订单末次成交价格
	FeeQty              string `json:"n"` //手续费数量
	FeeAsset            string `json:"N"` //手续费资产类别
	TradeTime           int64  `json:"T"` //成交时间
	TradeId             int64  `json:"t"` //成交ID
	IsWorking           bool   `json:"w"` //订单是否在订单簿上？
	IsMaker             bool   `json:"m"` //该成交是作为挂单成交吗？
	OrderCreateTime     int64  `json:"O"` //订单创建时间
	CummulativeQuoteQty string `json:"Z"` //订单累计已成交金额
	LastQuoteQty        string `json:"Y"` //订单末次成交金额
	QuoteOrderQty       string `json:"Q"` //Quote Order Quantity
	SelfTradePrevention string `json:"V"` //SelfTradePreventionMode

	//特殊字段
	TrailingDelta              int64  `json:"d"`  //只出现在追踪止损订单中。
	TrailingTime               int64  `json:"D"`  //Trailing Time
	StrategyId                 int64  `json:"j"`  //如果在请求中添加了strategyId参数，则会出现。
	StrategyType               int64  `json:"J"`  //如果在请求中添加了strategyType参数，则会出现。
	PreventMatchId             int64  `json:"v"`  //被阻止撮合交易的ID; 这仅在订单因 STP 触发而过期时可见
	PreventedQuantity          string `json:"A"`  //只有在因为 STP 导致订单失效时可见。
	LastPreventedQuantity      string `json:"B"`  //Last Prevented Quantity
	TradeGroupId               int64  `json:"u"`  //Trade Group Id
	CounterOrderId             int64  `json:"U"`  //Counter Order Id
	CounterSymbol              string `json:"Cs"` //Counter Symbol
	PreventedExecutionQuantity string `json:"pl"` //Prevented Execution Quantity
	PreventedExecutionPrice    string `json:"pL"` //Prevented Execution Price
	PreventedExecutionQuoteQty string `json:"pY"` //Prevented Execution Quote Qty
	WorkingTime                int64  `json:"W"`  //Working Time; 订单被添加到 order book 的时间; 只有在订单在订单簿上时可见
	MatchType                  string `json:"b"`  //Match Type
	AllocationId               int64  `json:"a"`  //Allocation ID
	WorkingFloor               string `json:"k"`  //Working Floor; 只有在订单可能有分配时可见
	UsedSor                    bool   `json:"uS"` //UsedSor; 只有在订单使用 SOR 时可见
}

// U本位合约余额和仓位推送
type WsFuturePayloadAccountUpdate struct {
	Event     string         `json:"e"` //事件类型
	Timestamp int64          `json:"E"` //事件时间
	TradeTime int64          `json:"T"` //撮合时间
	Action    WsFutureAction `json:"a"` // 账户更新事件
}

type WsFutureAction struct {
	Reason   string             `json:"m"` //事件推出原因
	Balance  []WsFutureBalance  `json:"B"` //余额信息
	Position []WsFuturePosition `json:"P"` //持仓信息
}

type WsFutureBalance struct {
	Asset              string `json:"a"`  //资产名称
	WalletBalance      string `json:"wb"` //钱包余额
	CrossWalletBalance string `json:"cw"` //除去逐仓仓位保证金的钱包余额
	BalanceChange      string `json:"bc"` //除去盈亏与交易手续费以外的钱包余额改变量
}

type WsFuturePosition struct {
	Symbol             string `json:"s"`   //交易对
	PositionAmount     string `json:"pa"`  //仓位
	EntryPrice         string `json:"ep"`  //入仓价格
	BreakEvenPrice     string `json:"bep"` //盈亏平衡价
	CumulativeRealized string `json:"cr"`  //(费前)累计实现损益
	UnrealizedProfit   string `json:"up"`  //持仓未实现盈亏
	MarginType         string `json:"mt"`  //保证金模式
	InitialMargin      string `json:"iw"`  //若为逐仓，仓位保证金
	PositionSide       string `json:"ps"`  //持仓方向
}

// U本位合约订单/交易 更新推送
type WsFuturePayloadOrderTradeUpdate struct {
	Event     string        `json:"e"` //事件类型
	Timestamp int64         `json:"E"` //事件时间
	TradeTime int64         `json:"T"` //撮合时间
	Order     WsFutureOrder `json:"o"` //订单信息
}

type WsFutureOrder struct {
	Symbol         string `json:"s"`   //交易对
	ClientOrderId  string `json:"c"`   //clientOrderId
	Side           string `json:"S"`   //订单方向
	Type           string `json:"o"`   //订单类型
	TimeInForce    string `json:"f"`   //有效方式
	OrigQty        string `json:"q"`   //订单原始数量
	Price          string `json:"p"`   //订单原始价格
	AvgPrice       string `json:"ap"`  //订单平均价格
	StopPrice      string `json:"sp"`  //止盈止损单触发价格
	ExecutionType  string `json:"x"`   //本次事件的具体执行类型
	Status         string `json:"X"`   //订单的当前状态
	OrderId        int64  `json:"i"`   //orderId
	LastQty        string `json:"l"`   //订单末次成交量
	ExecutedQty    string `json:"z"`   //订单累计已成交量
	LastPrice      string `json:"L"`   //订单末次成交价格
	FeeAsset       string `json:"N"`   //手续费资产类别
	FeeQty         string `json:"n"`   //手续费数量
	TradeTime      int64  `json:"T"`   //成交时间
	TradeId        int64  `json:"t"`   //成交ID
	BuyNetValue    string `json:"b"`   //买单净值
	SellNetValue   string `json:"a"`   //卖单净值
	IsMaker        bool   `json:"m"`   //该成交是作为挂单成交吗？
	IsReduceOnly   bool   `json:"R"`   //是否是只减仓单
	TriggerType    string `json:"wt"`  //触发价类型
	OrigType       string `json:"ot"`  //原始订单类型
	PositionSide   string `json:"ps"`  //持仓方向
	IsClose        bool   `json:"cp"`  //是否为触发平仓单; 仅在条件订单情况下会推送此字段
	ActivePrice    string `json:"AP"`  //追踪止损激活价格, 仅在追踪止损单时会推送此字段
	CallbackRate   string `json:"cr"`  //追踪止损回调比例, 仅在追踪止损单时会推送此字段
	PriceProtect   bool   `json:"pP"`  //是否开启条件单触发保护
	RealizedProfit string `json:"rp"`  //该交易实现盈亏
	PreventMode    string `json:"V"`   //自成交防止模式
	PriceMatch     string `json:"pm"`  //价格匹配模式
	GoodTillDate   int64  `json:"gtd"` //TIF为GTD的订单自动取消时间
}

// 币本位合约账户更新推送
type WsSwapPayloadAccountUpdate struct {
	Event        string       `json:"e"` //事件类型
	Timestamp    int64        `json:"E"` //事件时间
	TradeTime    int64        `json:"T"` //撮合时间
	AccountAlias string       `json:"i"` //账户唯一识别码 accountAlias
	Action       WsSwapAction `json:"a"` // 账户更新事件
}

type WsSwapAction struct {
	Reason   string           `json:"m"` //事件推出原因
	Balance  []WsSwapBalance  `json:"B"` //余额信息
	Position []WsSwapPosition `json:"P"` //持仓信息
}

type WsSwapBalance struct {
	Asset              string `json:"a"`  //资产名称
	WalletBalance      string `json:"wb"` //钱包余额
	CrossWalletBalance string `json:"cw"` //除去逐仓仓位保证金的钱包余额
	BalanceChange      string `json:"bc"` //除去盈亏与交易手续费以外的钱包余额改变量
}

type WsSwapPosition struct {
	Symbol             string `json:"s"`   //交易对
	PositionAmount     string `json:"pa"`  //仓位
	EntryPrice         string `json:"ep"`  //入仓价格
	BreakEvenPrice     string `json:"bep"` //盈亏平衡价
	CumulativeRealized string `json:"cr"`  //(费前)累计实现损益
	UnrealizedProfit   string `json:"up"`  //持仓未实现盈亏
	MarginType         string `json:"mt"`  //保证金模式
	InitialMargin      string `json:"iw"`  //若为逐仓，仓位保证金
	PositionSide       string `json:"ps"`  //持仓方向
}

// {

// 	"e":"ORDER_TRADE_UPDATE",         // 事件类型
// 	"E":1568879465651,                // 事件时间
// 	"T":1568879465650,                // 撮合时间
// 	"i": "SfsR",                          // 账户唯一识别码 accountAlias
// 	"o":{
// 	  "s":"BTCUSD_200925",                    // 交易对
// 	  "c":"TEST",                     // 客户端自定订单ID
// 		// 特殊的自定义订单ID:
// 		// "autoclose-"开头的字符串: 系统强平订单
// 		// "delivery-"开头的字符串: 系统交割平仓单
// 		// "delivery_autoclose-": 下架或交割的结算订单
// 	  "S":"SELL",                     // 订单方向
// 	  "o":"LIMIT",                    // 订单类型
// 	  "f":"GTC",                      // 有效方式
// 	  "q":"0.001",                    // 订单原始数量
// 	  "p":"9910",                     // 订单原始价格
// 	  "ap":"0",                       // 订单平均价格
// 	  "sp":"0",                       // 订单停止价格
// 	  "x":"NEW",                      // 本次事件的具体执行类型
// 	  "X":"NEW",                      // 订单的当前状态
// 	  "i":8886774,                    // 订单ID
// 	  "l":"0",                        // 订单末次成交量
// 	  "z":"0",                        // 订单累计已成交量
// 	  "L":"0",                        // 订单末次成交价格
// 	  "ma": "BTC",                    // 保证金资产类型
// 	  "N": "BTC",                     // 该成交手续费资产类型
// 	  "n": "0",                       // 该成交手续费数量
// 	  "T":1568879465650,              // 成交时间
// 	  "t":0,                          // 成交ID
// 	  "rp": "0",                      // 该成交已实现盈亏
// 	  "b":"0",                        // 买单净值
// 	  "a":"9.91",                     // 卖单净值
// 	  "m": false,                     // 该成交是作为挂单成交吗？
// 	  "R":false   ,                   // 是否是只减仓单
// 	  "wt": "CONTRACT_PRICE",         // 触发价类型
// 	  "ot": "LIMIT",                  // 原始订单类型
// 	  "ps":"LONG",                        // 持仓方向
// 	  "cp":false,                     // 是否为触发平仓单
// 	  "AP":"7476.89",                 // 追踪止损激活价格
// 	  "cr":"5.0",                     // 追踪止损回调比例
// 	  "pP": false                     // 是否开启条件单触发保护

// 	}

//	}
//
// 币本位合约订单/交易 更新推送
type WsSwapPayloadOrderTradeUpdate struct {
	Event        string      `json:"e"` //事件类型
	Timestamp    int64       `json:"E"` //事件时间
	TradeTime    int64       `json:"T"` //撮合时间
	AccountAlias string      `json:"i"` //账户唯一识别码 accountAlias
	Order        WsSwapOrder `json:"o"` //订单信息
}

type WsSwapOrder struct {
	Symbol         string `json:"s"`  //交易对
	ClientOrderId  string `json:"c"`  //clientOrderId
	Side           string `json:"S"`  //订单方向
	Type           string `json:"o"`  //订单类型
	TimeInForce    string `json:"f"`  //有效方式
	OrigQty        string `json:"q"`  //订单原始数量
	Price          string `json:"p"`  //订单原始价格
	AvgPrice       string `json:"ap"` //订单平均价格
	StopPrice      string `json:"sp"` //止盈止损单触发价格
	ExecutionType  string `json:"x"`  //本次事件的具体执行类型
	Status         string `json:"X"`  //订单的当前状态
	OrderId        int64  `json:"i"`  //orderId
	LastQty        string `json:"l"`  //订单末次成交量
	ExecutedQty    string `json:"z"`  //订单累计已成交量
	LastPrice      string `json:"L"`  //订单末次成交价格
	MarginAsset    string `json:"ma"` //保证金资产类型
	FeeAsset       string `json:"N"`  //手续费资产类别
	FeeQty         string `json:"n"`  //手续费数量
	TradeTime      int64  `json:"T"`  //成交时间
	TradeId        int64  `json:"t"`  //成交ID
	RealizedProfit string `json:"rp"` //该交易实现盈亏
	BuyNetValue    string `json:"b"`  //买单净值
	SellNetValue   string `json:"a"`  //卖单净值
	IsMaker        bool   `json:"m"`  //该成交是作为挂单成交吗？
	IsReduceOnly   bool   `json:"R"`  //是否是只减仓单
	TriggerType    string `json:"wt"` //触发价类型
	OrigType       string `json:"ot"` //原始订单类型
	PositionSide   string `json:"ps"` //持仓方向
	IsClose        bool   `json:"cp"` //是否为触发平仓单
	ActivePrice    string `json:"AP"` //追踪止损激活价格
	CallbackRate   string `json:"cr"` //追踪止损回调比例
	PriceProtect   bool   `json:"pP"` //是否开启条件单触发保护
}

// 统一账户合约账户更新
type WsPMContractPayloadAccountUpdate struct {
	Event        string             `json:"e"`  //事件类型
	EventLine    string             `json:"fs"` //事件业务线
	Timestamp    int64              `json:"E"`  //事件时间
	TradeTime    int64              `json:"T"`  //撮合时间
	AccountAlias string             `json:"i"`  //账户唯一识别码 accountAlias
	Action       WsPMContractAction `json:"a"`  // 账户更新事件
}

type WsPMContractAction struct {
	Reason   string                 `json:"m"` //事件推出原因
	Balance  []WsPMContractBalance  `json:"B"` //余额信息
	Position []WsPMContractPosition `json:"P"` //持仓信息
}

type WsPMContractBalance struct {
	Asset              string `json:"a"`  //资产名称
	WalletBalance      string `json:"wb"` //钱包余额
	CrossWalletBalance string `json:"cw"` //除去逐仓仓位保证金的钱包余额
	BalanceChange      string `json:"bc"` //除去盈亏与交易手续费以外的钱包余额改变量
}

type WsPMContractPosition struct {
	Symbol             string `json:"s"`   //交易对
	PositionAmount     string `json:"pa"`  //仓位
	EntryPrice         string `json:"ep"`  //入仓价格
	CumulativeRealized string `json:"cr"`  //(费前)累计实现损益
	UnrealizedProfit   string `json:"up"`  //持仓未实现盈亏
	PositionSide       string `json:"ps"`  //持仓方向
	BreakEvenPrice     string `json:"bep"` //盈亏平衡价
}

// 统一账户合约订单/交易 更新推送
type WsPMContractPayloadOrderTradeUpdate struct {
	Event     string            `json:"e"`  //事件类型
	Timestamp int64             `json:"E"`  //事件时间
	TradeTime int64             `json:"T"`  //撮合时间
	EventLine string            `json:"fs"` //事件业务线
	Order     WsPMContractOrder `json:"o"`  //订单信息
}

type WsPMContractOrder struct {
	Symbol         string `json:"s"`   //交易对
	ClientOrderId  string `json:"c"`   //clientOrderId
	Side           string `json:"S"`   //订单方向
	Type           string `json:"o"`   //订单类型
	TimeInForce    string `json:"f"`   //有效方式
	OrigQty        string `json:"q"`   //订单原始数量
	Price          string `json:"p"`   //订单原始价格
	AvgPrice       string `json:"ap"`  //订单平均价格
	StopPrice      string `json:"sp"`  //止盈止损单触发价格
	ExcecutionType string `json:"x"`   //本次事件的具体执行类型
	Status         string `json:"X"`   //订单的当前状态
	OrderId        int64  `json:"i"`   //orderId
	LastQty        string `json:"l"`   //订单末次成交量
	ExecutedQty    string `json:"z"`   //订单累计已成交量
	LastPrice      string `json:"L"`   //订单末次成交价格
	FeeAsset       string `json:"N"`   //手续费资产类别
	FeeQty         string `json:"n"`   //手续费数量
	TradeTime      int64  `json:"T"`   //成交时间
	TradeId        int64  `json:"t"`   //成交ID
	BuyNetValue    string `json:"b"`   //买单净值
	SellNetValue   string `json:"a"`   //卖单净值
	IsMaker        bool   `json:"m"`   //该成交是作为挂单成交吗？
	IsReduceOnly   bool   `json:"R"`   //是否是只减仓单
	PositionSide   string `json:"ps"`  //持仓方向
	RealizedProfit string `json:"rp"`  //该交易实现盈亏
	StrategyType   string `json:"st"`  //策略单类型
	StrategyId     int64  `json:"si"`  //策略单ID
	PreventMode    string `json:"V"`   //自成交防止模式
	GoodTillDate   int64  `json:"gtd"` //TIF为GTD的订单自动取消时间
}

type WsPMMarginBalance struct {
	Asset  string `json:"a"` //资产名称
	Free   string `json:"f"` //可用余额
	Locked string `json:"l"` //冻结余额
}

// 统一账户杠杆账户更新
type WsPMMarginPayloadOutboundAccountPosition struct {
	Event          string              `json:"e"` //事件类型
	Timestamp      int64               `json:"E"` //事件时间
	LastUpdateTime int64               `json:"u"` //账户末次更新时间戳
	TimeUpdateId   int64               `json:"U"` //时间更新ID
	Balances       []WsPMMarginBalance `json:"B"` //余额
}

type WsPMMarginPayloadExecutionReport struct {
	Event               string `json:"e"` //事件类型
	Timestamp           int64  `json:"E"` //事件时间
	Symbol              string `json:"s"` //交易对
	ClientOrderId       string `json:"c"` //clientOrderId
	Side                string `json:"S"` //订单方向
	Type                string `json:"o"` //订单类型
	TimeInForce         string `json:"f"` //有效方式
	OrigQty             string `json:"q"` //订单原始数量
	Price               string `json:"p"` //订单原始价格
	StopPrice           string `json:"P"` //止盈止损单触发价格
	IcebergQty          string `json:"F"` //冰山订单数量
	OrderListId         int64  `json:"g"` //OCO订单 OrderListId
	OrigClientOrderId   string `json:"C"` //原始订单自定义ID(原始订单，指撤单操作的对象。撤单本身被视为另一个订单)
	ExecutionType       string `json:"x"` //本次事件的具体执行类型
	Status              string `json:"X"` //订单的当前状态
	RejectReason        string `json:"r"` //订单被拒绝的原因
	OrderId             int64  `json:"i"` //orderId
	LastQty             string `json:"l"` //订单末次成交量
	ExecutedQty         string `json:"z"` //订单累计已成交量
	LastPrice           string `json:"L"` //订单末次成交价格
	FeeQty              string `json:"n"` //手续费数量
	FeeAsset            string `json:"N"` //手续费资产类别
	TradeTime           int64  `json:"T"` //成交时间
	TradeId             int64  `json:"t"` //成交ID
	PreventMatchId      int64  `json:"v"` //被阻止撮合交易的ID; 这仅在订单因 STP 触发而过期时可见
	IsWorking           bool   `json:"w"` //订单是否在订单簿上？
	IsMaker             bool   `json:"m"` //该成交是作为挂单成交吗？
	OrderCreateTime     int64  `json:"O"` //订单创建时间
	CummulativeQuoteQty string `json:"Z"` //订单累计已成交金额
	LastQuoteQty        string `json:"Y"` //订单末次成交金额
	QuoteOrderQty       string `json:"Q"` //Quote Order Quantity
	WorkingTime         int64  `json:"W"` //Working Time; 订单被添加到 order book 的时间
	SelfTradePrevention string `json:"V"` //SelfTradePreventionMode
}

func HandleWsPayloadResult[T any](data []byte) (*T, error) {
	var result T
	err := json.Unmarshal(data, &result)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &result, nil
}
