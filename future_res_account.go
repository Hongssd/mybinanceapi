package mybinanceapi

import "github.com/shopspring/decimal"

type FutureAccountResAsset struct {
	Asset                  string `json:"asset"`                  //资产
	WalletBalance          string `json:"walletBalance"`          //余额
	UnrealizedProfit       string `json:"unrealizedProfit"`       // 未实现盈亏
	MarginBalance          string `json:"marginBalance"`          // 保证金余额
	MaintMargin            string `json:"maintMargin"`            // 维持保证金
	InitialMargin          string `json:"initialMargin"`          // 当前所需起始保证金
	PositionInitialMargin  string `json:"positionInitialMargin"`  // 持仓所需起始保证金(基于最新标记价格)
	OpenOrderInitialMargin string `json:"openOrderInitialMargin"` // 当前挂单所需起始保证金(基于最新标记价格)
	CrossWalletBalance     string `json:"crossWalletBalance"`     //全仓账户余额
	CrossUnPnl             string `json:"crossUnPnl"`             // 全仓持仓未实现盈亏
	AvailableBalance       string `json:"availableBalance"`       // 可用余额
	MaxWithdrawAmount      string `json:"maxWithdrawAmount"`      // 最大可转出余额
	MarginAvailable        bool   `json:"marginAvailable"`        // 是否可用作联合保证金
	UpdateTime             int64  `json:"updateTime"`             //更新时间
}
type FutureAccountResPosition struct {
	//根据用户持仓模式展示持仓方向，即单向模式下只返回BOTH持仓情况，双向模式下只返回 LONG 和 SHORT 持仓情况
	Symbol                 string `json:"symbol"`                 // 交易对
	InitialMargin          string `json:"initialMargin"`          // 当前所需起始保证金(基于最新标记价格)
	MaintMargin            string `json:"maintMargin"`            //维持保证金
	UnrealizedProfit       string `json:"unrealizedProfit"`       // 持仓未实现盈亏
	PositionInitialMargin  string `json:"positionInitialMargin"`  // 持仓所需起始保证金(基于最新标记价格)
	OpenOrderInitialMargin string `json:"openOrderInitialMargin"` // 当前挂单所需起始保证金(基于最新标记价格)
	Leverage               string `json:"leverage"`               // 杠杆倍率
	Isolated               bool   `json:"isolated"`               // 是否是逐仓模式
	EntryPrice             string `json:"entryPrice"`             // 持仓成本价
	MaxNotional            string `json:"maxNotional"`            // 当前杠杆下用户可用的最大名义价值
	BidNotional            string `json:"bidNotional"`            // 买单净值，忽略
	AskNotional            string `json:"askNotional"`            // 卖单净值，忽略
	PositionSide           string `json:"positionSide"`           // 持仓方向
	PositionAmt            string `json:"positionAmt"`            // 持仓数量
	UpdateTime             int64  `json:"updateTime"`             // 更新时间
}
type FutureAccountRes struct {
	FeeTier                     int64                      `json:"feeTier"`     // 手续费等级
	CanTrade                    bool                       `json:"canTrade"`    // 是否可以交易
	CanDeposit                  bool                       `json:"canDeposit"`  // 是否可以入金
	CanWithdraw                 bool                       `json:"canWithdraw"` // 是否可以出金
	UpdateTime                  int64                      `json:"updateTime"`  // 保留字段，请忽略
	MultiAssetsMargin           bool                       `json:"multiAssetsMargin"`
	TotalInitialMargin          string                     `json:"totalInitialMargin"`          // 当前所需起始保证金总额(存在逐仓请忽略), 仅计算u
	TotalMaintMargin            string                     `json:"totalMaintMargin"`            // 维持保证金总额, 仅计算usdt资产
	TotalWalletBalance          string                     `json:"totalWalletBalance"`          // 账户总余额, 仅计算usdt资产
	TotalUnrealizedProfit       string                     `json:"totalUnrealizedProfit"`       // 持仓未实现盈亏总额, 仅计算usdt资产
	TotalMarginBalance          string                     `json:"totalMarginBalance"`          // 保证金总余额, 仅计算usdt资产
	TotalPositionInitialMargin  string                     `json:"totalPositionInitialMargin"`  // 持仓所需起始保证金(基于最新标记价格), 仅计算usdt
	TotalOpenOrderInitialMargin string                     `json:"totalOpenOrderInitialMargin"` // 当前挂单所需起始保证金(基于最新标记价格), 仅计算
	TotalCrossWalletBalance     string                     `json:"totalCrossWalletBalance"`     // 全仓账户余额, 仅计算usdt资产
	TotalCrossUnPnl             string                     `json:"totalCrossUnPnl"`             // 全仓持仓未实现盈亏总额, 仅计算usdt资产
	AvailableBalance            string                     `json:"availableBalance"`            // 可用余额, 仅计算usdt资产
	MaxWithdrawAmount           string                     `json:"maxWithdrawAmount"`           // 最大可转出余额, 仅计算usdt资产
	Assets                      []FutureAccountResAsset    `json:"assets"`
	Positions                   []FutureAccountResPosition `json:"positions"` // 头寸，将返回所有市场symbol。
}

type FuturePositionSideDualGetRes struct {
	DualSidePosition bool `json:"dualSidePosition"` // "true": 双向持仓模式；"false": 单向持仓模式
}

type FutureMultiAssetsMarginGetRes struct {
	MultiAssetsMargin bool `json:"multiAssetsMargin"` // "true": 联合保证金模式开启；"false": 联合保证金模式关闭
}

type FuturePositionSideDualPostRes FutureCommonPostRes

type FutureMultiAssetsMarginPostRes FutureCommonPostRes

type FutureLeverageRes struct {
	Leverage         int64  `json:"leverage"`         // 杠杆倍数
	MaxNotionalValue string `json:"maxNotionalValue"` // 当前杠杆倍数下允许的最大名义价值
	Symbol           string `json:"symbol"`           // 交易对
}

type FutureMarginTypeRes FutureCommonPostRes

type FutureLeverageBracketRes []FutureLeverageBracketResResult
type FutureLeverageBracketResResult struct {
	Symbol   string                                   `json:"symbol"`   // 交易对
	Brackets []FutureLeverageBracketResResultBrackets `json:"brackets"` // 交易对
}
type FutureLeverageBracketResResultBrackets struct {
	Bracket          int64           `json:"bracket"`          // 层级
	InitialLeverage  int64           `json:"initialLeverage"`  // 该层允许的最高初始杠杆倍数
	NotionalCap      int64           `json:"notionalCap"`      // 该层对应的名义价值上限
	NotionalFloor    int64           `json:"notionalFloor"`    // 该层对应的名义价值下限
	MaintMarginRatio decimal.Decimal `json:"maintMarginRatio"` // 该层对应的维持保证金率
	Cum              decimal.Decimal `json:"cum"`              // 速算数
}

type FuturePositionRiskRes []FuturePositionRiskRow
type FuturePositionRiskRow struct {
	EntryPrice       string `json:"entryPrice"`       //开仓均价
	BreakEvenPrice   string `json:"breakEvenPrice"`   //盈亏平衡价
	MarginType       string `json:"marginType"`       //逐仓模式或全仓模式
	IsAutoAddMargin  string `json:"isAutoAddMargin"`  //是否自动追加保证金
	IsolatedMargin   string `json:"isolatedMargin"`   //逐仓保证金
	Leverage         string `json:"leverage"`         //当前杠杆倍数
	LiquidationPrice string `json:"liquidationPrice"` //参考强平价格
	MarkPrice        string `json:"markPrice"`        //当前标记价格
	MaxNotionalValue string `json:"maxNotionalValue"` //当前杠杆倍数允许的名义价值上限
	PositionAmt      string `json:"positionAmt"`      //头寸数量，符号代表多空方向, 正数为多，负数为空
	Notional         string `json:"notional"`         //名义价值
	IsolatedWallet   string `json:"isolatedWallet"`   //逐仓钱包余额
	Symbol           string `json:"symbol"`           //交易对
	UnRealizedProfit string `json:"unRealizedProfit"` //持仓未实现盈亏
	PositionSide     string `json:"positionSide"`     //持仓方向
	UpdateTime       int64  `json:"updateTime"`       //更新时间
}
