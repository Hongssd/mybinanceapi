package mybinanceapi

import "github.com/shopspring/decimal"

type SpotMarginAllPairsReq struct {
	Recvwindow *int64 `json:"recvWindow"`
	Timestamp  *int64 `json:"timestamp"`
}
type SpotMarginAllPairsApi struct {
	client *SpotRestClient
	req    *SpotMarginAllPairsReq
}

func (api *SpotMarginAllPairsApi) Recvwindow(Recvwindow int64) *SpotMarginAllPairsApi {
	api.req.Recvwindow = GetPointer(Recvwindow)
	return api
}
func (api *SpotMarginAllPairsApi) Timestamp(Timestamp int64) *SpotMarginAllPairsApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SpotMarginIsolatedAllPairsReq struct {
	Recvwindow *int64 `json:"recvWindow"`
	Timestamp  *int64 `json:"timestamp"`
}
type SpotMarginIsolatedAllPairsApi struct {
	client *SpotRestClient
	req    *SpotMarginIsolatedAllPairsReq
}

func (api *SpotMarginIsolatedAllPairsApi) Recvwindow(Recvwindow int64) *SpotMarginIsolatedAllPairsApi {
	api.req.Recvwindow = GetPointer(Recvwindow)
	return api
}
func (api *SpotMarginIsolatedAllPairsApi) Timestamp(Timestamp int64) *SpotMarginIsolatedAllPairsApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SpotMarginAccountReq struct {
	Recvwindow *int64 `json:"recvWindow"`
	Timestamp  *int64 `json:"timestamp"`
}
type SpotMarginAccountApi struct {
	client *SpotRestClient
	req    *SpotMarginAccountReq
}

func (api *SpotMarginAccountApi) Recvwindow(Recvwindow int64) *SpotMarginAccountApi {
	api.req.Recvwindow = GetPointer(Recvwindow)
	return api
}
func (api *SpotMarginAccountApi) Timestamp(Timestamp int64) *SpotMarginAccountApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SpotMarginIsolatedAccountReq struct {
	Symbols    *string `json:"symbols"` //NO	最多可以传5个symbol; 由","分隔的字符串表示. e.g. "BTCUSDT,BNBUSDT,ADAUSDT"
	Recvwindow *int64  `json:"recvWindow"`
	Timestamp  *int64  `json:"timestamp"`
}
type SpotMarginIsolatedAccountApi struct {
	client *SpotRestClient
	req    *SpotMarginIsolatedAccountReq
}

func (api *SpotMarginIsolatedAccountApi) Symbols(Symbols string) *SpotMarginIsolatedAccountApi {
	api.req.Symbols = GetPointer(Symbols)
	return api
}
func (api *SpotMarginIsolatedAccountApi) Recvwindow(Recvwindow int64) *SpotMarginIsolatedAccountApi {
	api.req.Recvwindow = GetPointer(Recvwindow)
	return api
}
func (api *SpotMarginIsolatedAccountApi) Timestamp(Timestamp int64) *SpotMarginIsolatedAccountApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SpotMarginMaxBorrowableReq struct {
	Asset          *string `json:"asset"`          //YES 币种, 比如, BTC
	IsolatedSymbol *string `json:"isolatedSymbol"` //NO 逐仓交易对，适用于逐仓查询
	Recvwindow     *int64  `json:"recvWindow"`     //NO 赋值不能大于 60000
	Timestamp      *int64  `json:"timestamp"`      //YES
}
type SpotMarginMaxBorrowableApi struct {
	client *SpotRestClient
	req    *SpotMarginMaxBorrowableReq
}

func (api *SpotMarginMaxBorrowableApi) Asset(Asset string) *SpotMarginMaxBorrowableApi {
	api.req.Asset = GetPointer(Asset)
	return api
}
func (api *SpotMarginMaxBorrowableApi) IsolatedSymbol(IsolatedSymbol string) *SpotMarginMaxBorrowableApi {
	api.req.IsolatedSymbol = GetPointer(IsolatedSymbol)
	return api
}
func (api *SpotMarginMaxBorrowableApi) Recvwindow(Recvwindow int64) *SpotMarginMaxBorrowableApi {
	api.req.Recvwindow = GetPointer(Recvwindow)
	return api
}
func (api *SpotMarginMaxBorrowableApi) Timestamp(Timestamp int64) *SpotMarginMaxBorrowableApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SpotMarginMaxTransferableReq struct {
	Asset          *string `json:"asset"`          //YES 币种, 比如, BTC
	IsolatedSymbol *string `json:"isolatedSymbol"` //NO 逐仓交易对，适用于逐仓查询
	Recvwindow     *int64  `json:"recvWindow"`     //NO 赋值不能大于 60000
	Timestamp      *int64  `json:"timestamp"`      //YES
}
type SpotMarginMaxTransferableApi struct {
	client *SpotRestClient
	req    *SpotMarginMaxTransferableReq
}

func (api *SpotMarginMaxTransferableApi) Asset(Asset string) *SpotMarginMaxTransferableApi {
	api.req.Asset = GetPointer(Asset)
	return api
}
func (api *SpotMarginMaxTransferableApi) IsolatedSymbol(IsolatedSymbol string) *SpotMarginMaxTransferableApi {
	api.req.IsolatedSymbol = GetPointer(IsolatedSymbol)
	return api
}
func (api *SpotMarginMaxTransferableApi) Recvwindow(Recvwindow int64) *SpotMarginMaxTransferableApi {
	api.req.Recvwindow = GetPointer(Recvwindow)
	return api
}
func (api *SpotMarginMaxTransferableApi) Timestamp(Timestamp int64) *SpotMarginMaxTransferableApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SpotMarginInterestHistoryReq struct {
	Asset          *string `json:"asset"`          //NO 币种, 比如, BTC
	IsolatedSymbol *string `json:"isolatedSymbol"` //NO 逐仓交易对，适用于逐仓查询
	StartTime      *int64  `json:"startTime"`      //NO
	EndTime        *int64  `json:"endTime"`        //NO
	Current        *int64  `json:"current"`        //NO	当前查询页。 开始值 1. 默认:1
	Size           *int64  `json:"size"`           //NO	默认:10 最大:100
	Archived       *string `json:"archived"`       //NO	默认: false. 查询6个月以前的数据，需要设为 true
	Recvwindow     *int64  `json:"recvWindow"`     //NO 赋值不能大于 60000
	Timestamp      *int64  `json:"timestamp"`      //YES
}
type SpotMarginInterestHistoryApi struct {
	client *SpotRestClient
	req    *SpotMarginInterestHistoryReq
}

func (api *SpotMarginInterestHistoryApi) Asset(Asset string) *SpotMarginInterestHistoryApi {
	api.req.Asset = GetPointer(Asset)
	return api
}
func (api *SpotMarginInterestHistoryApi) IsolatedSymbol(IsolatedSymbol string) *SpotMarginInterestHistoryApi {
	api.req.IsolatedSymbol = GetPointer(IsolatedSymbol)
	return api
}
func (api *SpotMarginInterestHistoryApi) StartTime(StartTime int64) *SpotMarginInterestHistoryApi {
	api.req.StartTime = GetPointer(StartTime)
	return api
}
func (api *SpotMarginInterestHistoryApi) EndTime(EndTime int64) *SpotMarginInterestHistoryApi {
	api.req.EndTime = GetPointer(EndTime)
	return api
}
func (api *SpotMarginInterestHistoryApi) Current(Current int64) *SpotMarginInterestHistoryApi {
	api.req.Current = GetPointer(Current)
	return api
}
func (api *SpotMarginInterestHistoryApi) Size(Size int64) *SpotMarginInterestHistoryApi {
	api.req.Size = GetPointer(Size)
	return api
}
func (api *SpotMarginInterestHistoryApi) Archived(Archived string) *SpotMarginInterestHistoryApi {
	api.req.Archived = GetPointer(Archived)
	return api
}
func (api *SpotMarginInterestHistoryApi) Recvwindow(Recvwindow int64) *SpotMarginInterestHistoryApi {
	api.req.Recvwindow = GetPointer(Recvwindow)
	return api
}
func (api *SpotMarginInterestHistoryApi) Timestamp(Timestamp int64) *SpotMarginInterestHistoryApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SpotMarginOrderGetReq struct {
	Symbol            *string `json:"symbol"`            //NO
	IsIsolated        *string `json:"isIsolated"`        //NO 	是否逐仓杠杆，"TRUE", "FALSE", 默认 "FALSE"
	OrderId           *int64  `json:"orderId"`           //NO
	OrigClientOrderId *string `json:"origClientOrderId"` //NO
	Recvwindow        *int64  `json:"recvWindow"`        //NO 赋值不能大于 60000
	Timestamp         *int64  `json:"timestamp"`         //YES
}
type SpotMarginOrderGetApi struct {
	client *SpotRestClient
	req    *SpotMarginOrderGetReq
}

func (api *SpotMarginOrderGetApi) Symbol(Symbol string) *SpotMarginOrderGetApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *SpotMarginOrderGetApi) IsIsolated(IsIsolated string) *SpotMarginOrderGetApi {
	api.req.IsIsolated = GetPointer(IsIsolated)
	return api
}
func (api *SpotMarginOrderGetApi) OrderId(OrderId int64) *SpotMarginOrderGetApi {
	api.req.OrderId = GetPointer(OrderId)
	return api
}
func (api *SpotMarginOrderGetApi) OrigClientOrderId(OrigClientOrderId string) *SpotMarginOrderGetApi {
	api.req.OrigClientOrderId = GetPointer(OrigClientOrderId)
	return api
}
func (api *SpotMarginOrderGetApi) Recvwindow(Recvwindow int64) *SpotMarginOrderGetApi {
	api.req.Recvwindow = GetPointer(Recvwindow)
	return api
}
func (api *SpotMarginOrderGetApi) Timestamp(Timestamp int64) *SpotMarginOrderGetApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SpotMarginOrderPostReq struct {
	Symbol                  *string          `json:"symbol"`                  //YES
	IsIsolated              *string          `json:"isIsolated"`              //NO	是否逐仓杠杆，"TRUE", "FALSE", 默认 "FALSE"
	Side                    *string          `json:"side"`                    //YES	BUY SELL
	Type                    *string          `json:"type"`                    //YES	详见枚举定义：订单类型
	Quantity                *decimal.Decimal `json:"quantity"`                //NO
	QuoteOrderQty           *decimal.Decimal `json:"quoteOrderQty"`           //NO
	Price                   *decimal.Decimal `json:"price"`                   //NO
	StopPrice               *decimal.Decimal `json:"stopPrice"`               //NO	与STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT, 和 TAKE_PROFIT_LIMIT 订单一起使用.
	NewClientOrderId        *string          `json:"newClientOrderId"`        //NO	客户自定义的唯一订单ID。若未发送自动生成。
	IcebergQty              *decimal.Decimal `json:"icebergQty"`              //NO	与 LIMIT, STOP_LOSS_LIMIT, 和 TAKE_PROFIT_LIMIT 一起使用创建 iceberg 订单.
	NewOrderRespType        *string          `json:"newOrderRespType"`        //NO	设置响应: JSON. ACK, RESULT, 或 FULL; MARKET 和 LIMIT 订单类型默认为 FULL, 所有其他订单默认为 ACK.
	SideEffectType          *string          `json:"sideEffectType"`          //NO	NO_SIDE_EFFECT, MARGIN_BUY, AUTO_REPAY;默认为 NO_SIDE_EFFECT.
	TimeInForce             *string          `json:"timeInForce"`             //NO	GTC,IOC,FOK
	SelfTradePreventionMode *string          `json:"selfTradePreventionMode"` //NO	允许的 ENUM 取决于交易对的配置。支持的值有 EXPIRE_TAKER，EXPIRE_MAKER，EXPIRE_BOTH，NONE
	AutoRepayAtCancel       *bool            `json:"autoRepayAtCancel"`       //NO	只有在自动借款单或者自动借还单生效，true表示的是撤单后需要把订单产生的借款归还，默认为true
	RecvWindow              *int64           `json:"recvWindow"`              //NO	赋值不能大于 60000
	Timestamp               *int64           `json:"timestamp"`               //YES
}
type SpotMarginOrderPostApi struct {
	client *SpotRestClient
	req    *SpotMarginOrderPostReq
}

func (api *SpotMarginOrderPostApi) Symbol(Symbol string) *SpotMarginOrderPostApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *SpotMarginOrderPostApi) IsIsolated(IsIsolated string) *SpotMarginOrderPostApi {
	api.req.IsIsolated = GetPointer(IsIsolated)
	return api
}
func (api *SpotMarginOrderPostApi) Side(Side string) *SpotMarginOrderPostApi {
	api.req.Side = GetPointer(Side)
	return api
}
func (api *SpotMarginOrderPostApi) Type(Type string) *SpotMarginOrderPostApi {
	api.req.Type = GetPointer(Type)
	return api
}
func (api *SpotMarginOrderPostApi) Quantity(Quantity decimal.Decimal) *SpotMarginOrderPostApi {
	api.req.Quantity = GetPointer(Quantity)
	return api
}
func (api *SpotMarginOrderPostApi) QuoteOrderQty(QuoteOrderQty decimal.Decimal) *SpotMarginOrderPostApi {
	api.req.QuoteOrderQty = GetPointer(QuoteOrderQty)
	return api
}
func (api *SpotMarginOrderPostApi) Price(Price decimal.Decimal) *SpotMarginOrderPostApi {
	api.req.Price = GetPointer(Price)
	return api
}
func (api *SpotMarginOrderPostApi) StopPrice(StopPrice decimal.Decimal) *SpotMarginOrderPostApi {
	api.req.StopPrice = GetPointer(StopPrice)
	return api
}
func (api *SpotMarginOrderPostApi) NewClientOrderId(NewClientOrderId string) *SpotMarginOrderPostApi {
	api.req.NewClientOrderId = GetPointer(NewClientOrderId)
	return api
}
func (api *SpotMarginOrderPostApi) IcebergQty(IcebergQty decimal.Decimal) *SpotMarginOrderPostApi {
	api.req.IcebergQty = GetPointer(IcebergQty)
	return api
}
func (api *SpotMarginOrderPostApi) NewOrderRespType(NewOrderRespType string) *SpotMarginOrderPostApi {
	api.req.NewOrderRespType = GetPointer(NewOrderRespType)
	return api
}
func (api *SpotMarginOrderPostApi) SideEffectType(SideEffectType string) *SpotMarginOrderPostApi {
	api.req.SideEffectType = GetPointer(SideEffectType)
	return api
}
func (api *SpotMarginOrderPostApi) TimeInForce(TimeInForce string) *SpotMarginOrderPostApi {
	api.req.TimeInForce = GetPointer(TimeInForce)
	return api
}
func (api *SpotMarginOrderPostApi) SelfTradePreventionMode(SelfTradePreventionMode string) *SpotMarginOrderPostApi {
	api.req.SelfTradePreventionMode = GetPointer(SelfTradePreventionMode)
	return api
}
func (api *SpotMarginOrderPostApi) AutoRepayAtCancel(AutoRepayAtCancel bool) *SpotMarginOrderPostApi {
	api.req.AutoRepayAtCancel = GetPointer(AutoRepayAtCancel)
	return api
}
func (api *SpotMarginOrderPostApi) RecvWindow(RecvWindow int64) *SpotMarginOrderPostApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *SpotMarginOrderPostApi) Timestamp(Timestamp int64) *SpotMarginOrderPostApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SpotMarginOrderDeleteReq struct {
	Symbol            *string `json:"symbol"`            //YES
	IsIsolated        *string `json:"isIsolated"`        //NO	是否逐仓杠杆，"TRUE", "FALSE", 默认 "FALSE"
	OrderId           *int64  `json:"orderId"`           //NO
	OrigClientOrderId *string `json:"origClientOrderId"` //NO
	NewClientOrderId  *string `json:"newClientOrderId"`  //NO	用于唯一识别此撤销订单，默认自动生成。
	RecvWindow        *int64  `json:"recvWindow"`        //NO	T赋值不能大于 60000
	Timestamp         *int64  `json:"timestamp"`         //YES
}
type SpotMarginOrderDeleteApi struct {
	client *SpotRestClient
	req    *SpotMarginOrderDeleteReq
}

func (api *SpotMarginOrderDeleteApi) Symbol(Symbol string) *SpotMarginOrderDeleteApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *SpotMarginOrderDeleteApi) IsIsolated(IsIsolated string) *SpotMarginOrderDeleteApi {
	api.req.IsIsolated = GetPointer(IsIsolated)
	return api
}
func (api *SpotMarginOrderDeleteApi) OrderId(OrderId int64) *SpotMarginOrderDeleteApi {
	api.req.OrderId = GetPointer(OrderId)
	return api
}
func (api *SpotMarginOrderDeleteApi) OrigClientOrderId(OrigClientOrderId string) *SpotMarginOrderDeleteApi {
	api.req.OrigClientOrderId = GetPointer(OrigClientOrderId)
	return api
}
func (api *SpotMarginOrderDeleteApi) NewClientOrderId(NewClientOrderId string) *SpotMarginOrderDeleteApi {
	api.req.NewClientOrderId = GetPointer(NewClientOrderId)
	return api
}
func (api *SpotMarginOrderDeleteApi) RecvWindow(RecvWindow int64) *SpotMarginOrderDeleteApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *SpotMarginOrderDeleteApi) Timestamp(Timestamp int64) *SpotMarginOrderDeleteApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SpotMarginAllOrdersReq struct {
	Symbol     *string `json:"symbol"`     //YES
	IsIsolated *string `json:"isIsolated"` //NO 	是否逐仓杠杆，"TRUE", "FALSE", 默认 "FALSE"
	OrderId    *int64  `json:"orderId"`    //NO
	StartTime  *int64  `json:"startTime"`  //NO
	EndTime    *int64  `json:"endTime"`    //NO
	Limit      *int64  `json:"limit"`      //NO	默认 500;最大500.
	Recvwindow *int64  `json:"recvWindow"` //NO 赋值不能大于 60000
	Timestamp  *int64  `json:"timestamp"`  //YES
}
type SpotMarginAllOrdersApi struct {
	client *SpotRestClient
	req    *SpotMarginAllOrdersReq
}

func (api *SpotMarginAllOrdersApi) Symbol(Symbol string) *SpotMarginAllOrdersApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *SpotMarginAllOrdersApi) IsIsolated(IsIsolated string) *SpotMarginAllOrdersApi {
	api.req.IsIsolated = GetPointer(IsIsolated)
	return api
}
func (api *SpotMarginAllOrdersApi) OrderId(OrderId int64) *SpotMarginAllOrdersApi {
	api.req.OrderId = GetPointer(OrderId)
	return api
}
func (api *SpotMarginAllOrdersApi) StartTime(StartTime int64) *SpotMarginAllOrdersApi {
	api.req.StartTime = GetPointer(StartTime)
	return api
}
func (api *SpotMarginAllOrdersApi) EndTime(EndTime int64) *SpotMarginAllOrdersApi {
	api.req.EndTime = GetPointer(EndTime)
	return api
}
func (api *SpotMarginAllOrdersApi) Limit(Limit int64) *SpotMarginAllOrdersApi {
	api.req.Limit = GetPointer(Limit)
	return api
}
func (api *SpotMarginAllOrdersApi) Recvwindow(Recvwindow int64) *SpotMarginAllOrdersApi {
	api.req.Recvwindow = GetPointer(Recvwindow)
	return api
}
func (api *SpotMarginAllOrdersApi) Timestamp(Timestamp int64) *SpotMarginAllOrdersApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SpotMarginOpenOrdersReq struct {
	Symbol     *string `json:"symbol"`     //NO
	IsIsolated *string `json:"isIsolated"` //NO 	是否逐仓杠杆，"TRUE", "FALSE", 默认 "FALSE"
	Recvwindow *int64  `json:"recvWindow"` //NO 赋值不能大于 60000
	Timestamp  *int64  `json:"timestamp"`  //YES
}
type SpotMarginOpenOrdersApi struct {
	client *SpotRestClient
	req    *SpotMarginOpenOrdersReq
}

func (api *SpotMarginOpenOrdersApi) Symbol(Symbol string) *SpotMarginOpenOrdersApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *SpotMarginOpenOrdersApi) IsIsolated(IsIsolated string) *SpotMarginOpenOrdersApi {
	api.req.IsIsolated = GetPointer(IsIsolated)
	return api
}
func (api *SpotMarginOpenOrdersApi) Recvwindow(Recvwindow int64) *SpotMarginOpenOrdersApi {
	api.req.Recvwindow = GetPointer(Recvwindow)
	return api
}
func (api *SpotMarginOpenOrdersApi) Timestamp(Timestamp int64) *SpotMarginOpenOrdersApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SpotMarginTransferReq struct {
	Asset      *string          `json:"asset"`      //YES 被划转的资产, 比如, BTC
	Amount     *decimal.Decimal `json:"amount"`     //YES 划转数量
	Type       *int             `json:"type"`       //YES 1: 主账户向全仓杠杆账户划转 2: 全仓杠杆账户向主账户划转
	Recvwindow *int64           `json:"recvWindow"` //NO 赋值不能大于 60000
	Timestamp  *int64           `json:"timestamp"`  //YES
}
type SpotMarginTransferApi struct {
	client *SpotRestClient
	req    *SpotMarginTransferReq
}

func (api *SpotMarginTransferApi) Asset(Asset string) *SpotMarginTransferApi {
	api.req.Asset = GetPointer(Asset)
	return api
}
func (api *SpotMarginTransferApi) Amount(Amount decimal.Decimal) *SpotMarginTransferApi {
	api.req.Amount = GetPointer(Amount)
	return api
}
func (api *SpotMarginTransferApi) Type(Type int) *SpotMarginTransferApi {
	api.req.Type = GetPointer(Type)
	return api
}
func (api *SpotMarginTransferApi) Recvwindow(Recvwindow int64) *SpotMarginTransferApi {
	api.req.Recvwindow = GetPointer(Recvwindow)
	return api
}
func (api *SpotMarginTransferApi) Timestamp(Timestamp int64) *SpotMarginTransferApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SpotMarginIsolatedTransferReq struct {
	Asset      *string          `json:"asset"`      //YES 被划转的资产, 比如, BTC
	Symbol     *string          `json:"symbol"`     //YES 逐仓 symbol
	TransFrom  *string          `json:"transFrom"`  //YES "SPOT", "ISOLATED_MARGIN"
	TransTo    *string          `json:"transTo"`    //YES "SPOT", "ISOLATED_MARGIN"
	Amount     *decimal.Decimal `json:"amount"`     //YES 划转数量
	Recvwindow *int64           `json:"recvWindow"` //NO 赋值不能大于 60000
	Timestamp  *int64           `json:"timestamp"`  //YES
}
type SpotMarginIsolatedTransferApi struct {
	client *SpotRestClient
	req    *SpotMarginIsolatedTransferReq
}

func (api *SpotMarginIsolatedTransferApi) Asset(Asset string) *SpotMarginIsolatedTransferApi {
	api.req.Asset = GetPointer(Asset)
	return api
}
func (api *SpotMarginIsolatedTransferApi) Symbol(Symbol string) *SpotMarginIsolatedTransferApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *SpotMarginIsolatedTransferApi) TransFrom(TransFrom string) *SpotMarginIsolatedTransferApi {
	api.req.TransFrom = GetPointer(TransFrom)
	return api
}
func (api *SpotMarginIsolatedTransferApi) TransTo(TransTo string) *SpotMarginIsolatedTransferApi {
	api.req.TransTo = GetPointer(TransTo)
	return api
}
func (api *SpotMarginIsolatedTransferApi) Amount(Amount decimal.Decimal) *SpotMarginIsolatedTransferApi {
	api.req.Amount = GetPointer(Amount)
	return api
}
func (api *SpotMarginIsolatedTransferApi) Recvwindow(Recvwindow int64) *SpotMarginIsolatedTransferApi {
	api.req.Recvwindow = GetPointer(Recvwindow)
	return api
}
func (api *SpotMarginIsolatedTransferApi) Timestamp(Timestamp int64) *SpotMarginIsolatedTransferApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SpotMarginLoanReq struct {
	Asset      *string          `json:"asset"`      //YES 被划转的资产, 比如, BTC
	IsIsolated *string          `json:"isIsolated"` //NO 是否逐仓杠杆，"TRUE", "FALSE", 默认 "FALSE"
	Symbol     *string          `json:"symbol"`     //NO 逐仓交易对，配合逐仓使用
	Amount     *decimal.Decimal `json:"amount"`     //YES 划转数量
	Recvwindow *int64           `json:"recvWindow"` //NO 赋值不能大于 60000
	Timestamp  *int64           `json:"timestamp"`  //YES
}
type SpotMarginLoanApi struct {
	client *SpotRestClient
	req    *SpotMarginLoanReq
}

func (api *SpotMarginLoanApi) Asset(Asset string) *SpotMarginLoanApi {
	api.req.Asset = GetPointer(Asset)
	return api
}
func (api *SpotMarginLoanApi) IsIsolated(IsIsolated string) *SpotMarginLoanApi {
	api.req.IsIsolated = GetPointer(IsIsolated)
	return api
}
func (api *SpotMarginLoanApi) Symbol(Symbol string) *SpotMarginLoanApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *SpotMarginLoanApi) Amount(Amount decimal.Decimal) *SpotMarginLoanApi {
	api.req.Amount = GetPointer(Amount)
	return api
}
func (api *SpotMarginLoanApi) Recvwindow(Recvwindow int64) *SpotMarginLoanApi {
	api.req.Recvwindow = GetPointer(Recvwindow)
	return api
}
func (api *SpotMarginLoanApi) Timestamp(Timestamp int64) *SpotMarginLoanApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type SpotMarginRepayReq struct {
	Asset      *string          `json:"asset"`      //YES 被划转的资产, 比如, BTC
	IsIsolated *string          `json:"isIsolated"` //NO 是否逐仓杠杆，"TRUE", "FALSE", 默认 "FALSE"
	Symbol     *string          `json:"symbol"`     //NO 逐仓交易对，配合逐仓使用
	Amount     *decimal.Decimal `json:"amount"`     //YES 划转数量
	Recvwindow *int64           `json:"recvWindow"` //NO 赋值不能大于 60000
	Timestamp  *int64           `json:"timestamp"`  //YES
}
type SpotMarginRepayApi struct {
	client *SpotRestClient
	req    *SpotMarginRepayReq
}

func (api *SpotMarginRepayApi) Asset(Asset string) *SpotMarginRepayApi {
	api.req.Asset = GetPointer(Asset)
	return api
}
func (api *SpotMarginRepayApi) IsIsolated(IsIsolated string) *SpotMarginRepayApi {
	api.req.IsIsolated = GetPointer(IsIsolated)
	return api
}
func (api *SpotMarginRepayApi) Symbol(Symbol string) *SpotMarginRepayApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *SpotMarginRepayApi) Amount(Amount decimal.Decimal) *SpotMarginRepayApi {
	api.req.Amount = GetPointer(Amount)
	return api
}
func (api *SpotMarginRepayApi) Recvwindow(Recvwindow int64) *SpotMarginRepayApi {
	api.req.Recvwindow = GetPointer(Recvwindow)
	return api
}
func (api *SpotMarginRepayApi) Timestamp(Timestamp int64) *SpotMarginRepayApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}
