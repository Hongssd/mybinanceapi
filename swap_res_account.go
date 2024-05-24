package mybinanceapi

import "github.com/shopspring/decimal"

type SwapCommonPostRes struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

type SwapAccountRes struct {
	Assets []struct {
		Asset                  string `json:"asset"`                  // 资产名
		WalletBalance          string `json:"walletBalance"`          // 账户余额
		UnrealizedProfit       string `json:"unrealizedProfit"`       // 全部持仓未实现盈亏
		MarginBalance          string `json:"marginBalance"`          // 保证金余额
		MaintMargin            string `json:"maintMargin"`            // 维持保证金
		InitialMargin          string `json:"initialMargin"`          // 当前所需起始保证金(按最新标标记价格)
		PositionInitialMargin  string `json:"positionInitialMargin"`  // 当前所需持仓起始保证金(按最新标标记价格)
		OpenOrderInitialMargin string `json:"openOrderInitialMargin"` // 当前所需挂单起始保证金(按最新标标记价格)
		MaxWithdrawAmount      string `json:"maxWithdrawAmount"`      // 最大可提款金额
		CrossWalletBalance     string `json:"crossWalletBalance"`     // 可用于全仓的账户余额
		CrossUnPnl             string `json:"crossUnPnl"`             // 所有全仓持仓的未实现盈亏
		AvailableBalance       string `json:"availableBalance"`       // 可用下单余额
	} `json:"assets"` // 资产内容

	Positions []struct {
		Symbol                 string `json:"symbol"`                 // 交易对
		PositionAmt            string `json:"positionAmt"`            // 持仓数量
		InitialMargin          string `json:"initialMargin"`          // 当前所需起始保证金(按最新标标记价格)
		MaintMargin            string `json:"maintMargin"`            // 持仓维持保证金
		UnrealizedProfit       string `json:"unrealizedProfit"`       // 持仓未实现盈亏
		PositionInitialMargin  string `json:"positionInitialMargin"`  // 当前所需持仓起始保证金(按最新标标记价格)
		OpenOrderInitialMargin string `json:"openOrderInitialMargin"` // 当前所需挂单起始保证金(按最新标标记价格)
		Leverage               string `json:"leverage"`               // 杠杆倍率
		Isolated               bool   `json:"isolated"`               // 是否是逐仓模式
		PositionSide           string `json:"positionSide"`           // 持仓方向
		EntryPrice             string `json:"entryPrice"`             // 平均持仓成本
		UpdateTime             int64  `json:"updateTime"`             // 最新更新时间
		MaxQty                 string `json:"maxQty"`                 // 当前杠杆下最大可开仓数(标的数量)
	} `json:"positions"` // 头寸
	CanDeposit  bool  `json:"canDeposit"`  // 是否可以入金
	CanTrade    bool  `json:"canTrade"`    // 是否可以交易
	CanWithdraw bool  `json:"canWithdraw"` // 是否可以出金
	FeeTier     int64 `json:"feeTier"`     // 手续费等级
	UpdateTime  int64 `json:"updateTime"`
}

type SwapPositionSideDualGetRes struct {
	DualSidePosition bool `json:"dualSidePosition"` // "true": 双向持仓模式；"false": 单向持仓模式
}

type SwapLeverageBracketRes []SwapLeverageBracketResResult
type SwapLeverageBracketResResult struct {
	Symbol   string                                 `json:"symbol"`   // 交易对
	Brackets []SwapLeverageBracketResResultBrackets `json:"brackets"` // 杠杆档位
}
type SwapLeverageBracketResResultBrackets struct {
	Bracket          int64           `json:"bracket"`          // 层级
	InitialLeverage  int64           `json:"initialLeverage"`  // 该层允许的最高初始杠杆倍数
	QtyCap           int64           `json:"qtyCap"`           // 该层对应的数量上限
	QtyFloor         int64           `json:"qtyFloor"`         // 该层对应的数量下限
	MaintMarginRatio decimal.Decimal `json:"maintMarginRatio"` // 该层对应的维持保证金率
	Cum              decimal.Decimal `json:"cum"`              // 速算数
}

type SwapPositionSideDualPostRes SwapCommonPostRes

type SwapLeverageRes struct {
	Leverage         int64  `json:"leverage"`         // 杠杆倍数
	MaxNotionalValue string `json:"maxNotionalValue"` // 当前杠杆倍数下允许的最大名义价值
	Symbol           string `json:"symbol"`           // 交易对
}

type SwapMarginTypeRes SwapCommonPostRes
