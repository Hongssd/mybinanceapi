package mybinanceapi

// 账户接口

type PortfolioMarginBalanceResRow struct {
	Asset               string `json:"asset"`               // 资产
	TotalWalletBalance  string `json:"totalWalletBalance"`  // 钱包余额 =  全仓杠杆未锁定 + 全仓杠杆锁定 + u本位合约钱包余额 + 币本位合约钱包余额
	CrossMarginAsset    string `json:"crossMarginAsset"`    // 全仓资产 = 全仓杠杆未锁定 + 全仓杠杆锁定
	CrossMarginBorrowed string `json:"crossMarginBorrowed"` // 全仓杠杆借贷
	CrossMarginFree     string `json:"crossMarginFree"`     // 全仓杠杆未锁定
	CrossMarginInterest string `json:"crossMarginInterest"` // 全仓杠杆利息
	CrossMarginLocked   string `json:"crossMarginLocked"`   //全仓杠杆锁定
	UmWalletBalance     string `json:"umWalletBalance"`     // u本位合约钱包余额
	UmUnrealizedPNL     string `json:"umUnrealizedPNL"`     // u本位未实现盈亏
	CmWalletBalance     string `json:"cmWalletBalance"`     // 币本位合约钱包余额
	CmUnrealizedPNL     string `json:"cmUnrealizedPNL"`     // 币本位未实现盈亏
	UpdateTime          int64  `json:"updateTime"`
	NegativeBalance     string `json:"negativeBalance"`
}
type PortfolioMarginBalanceRes []PortfolioMarginBalanceResRow

type PortfolioMarginAccountRes struct {
	UniMMR                   string `json:"uniMMR"`                   // 统一账户维持保证金率
	AccountEquity            string `json:"accountEquity"`            // 以USD计价的账户权益
	ActualEquity             string `json:"actualEquity"`             // 不考虑质押率后的以USD计价账户权益
	AccountInitialMargin     string `json:"accountInitialMargin"`     //
	AccountMaintMargin       string `json:"accountMaintMargin"`       // 以USD计价统一账户维持保证金
	AccountStatus            string `json:"accountStatus"`            // 统一账户账户状态："NORMAL", "MARGIN_CALL", "SUPPLY_MARGIN", "REDUCE_ONLY", "ACTIVE_LIQUIDATION", "FORCE_LIQUIDATION", "BANKRUPTED"
	VirtualMaxWithdrawAmount string `json:"virtualMaxWithdrawAmount"` // 以USD计价的最大可转出
	TotalAvailableBalance    string `json:"totalAvailableBalance"`
	TotalMarginOpenLoss      string `json:"totalMarginOpenLoss"`
	UpdateTime               int64  `json:"updateTime"` // 更新时间
}

type PortfolioMarginMaxBorrowableRes struct {
	Amount      string `json:"amount"`      // 系统可借充足情况下用户账户当前最大可借额度
	BorrowLimit string `json:"borrowLimit"` // 平台限制的用户当前等级可以借的额度
}

type PortfolioMarginMaxWithdrawRes struct {
	Amount string `json:"amount"` // 系统可转出充足情况下用户账户当前最大可转出额度
}

type PortfolioMarginSetUmLeverageRes struct {
	Leverage         int    `json:"leverage"`         // 杠杆倍数
	MaxNotionalValue string `json:"maxNotionalValue"` // 最大名义价值
	Symbol           string `json:"symbol"`           // 交易对
}

type PortfolioMarginSetCmLeverageRes struct {
	Leverage         int    `json:"leverage"`         // 杠杆倍数
	MaxNotionalValue string `json:"maxNotionalValue"` // 最大名义价值
	Symbol           string `json:"symbol"`           // 交易对
}

type PortfolioMarginUmPositionSideDualPostRes struct{}
type PortfolioMarginUmPositionSideDualGetRes struct {
	DualSidePosition bool `json:"dualSidePosition"` // 是否开启双向持仓
}

type PortfolioMarginCmPositionSideDualPostRes struct{}
type PortfolioMarginCmPositionSideDualGetRes struct {
	DualSidePosition bool `json:"dualSidePosition"` // 是否开启双向持仓
}

type PortfolioMarginUmAccountV1Res struct {
	Assets []struct {
		Asset                  string `json:"asset"`                  // 资产
		CrossWalletBalance     string `json:"crossWalletBalance"`     // 全仓账户余额
		CrossUnPnl             string `json:"crossUnPnl"`             // 全仓持仓未实现盈亏
		MaintMargin            string `json:"maintMargin"`            // 维持保证金
		InitialMargin          string `json:"initialMargin"`          // 当前所需起始保证金
		PositionInitialMargin  string `json:"positionInitialMargin"`  // 持仓所需起始保证金(基于最新标记价格)
		OpenOrderInitialMargin string `json:"openOrderInitialMargin"` // 当前挂单所需起始保证金(基于最新标记价格)
		UpdateTime             int64  `json:"updateTime"`             // 更新时间
	} `json:"assets"` // 资产
	Positions []struct {
		Symbol                 string `json:"symbol"`                 // 交易对
		InitialMargin          string `json:"initialMargin"`          // 当前所需起始保证金(基于最新标记价格)
		MaintMargin            string `json:"maintMargin"`            // 维持保证金
		UnrealizedProfit       string `json:"unrealizedProfit"`       // 持仓未实现盈亏
		PositionInitialMargin  string `json:"positionInitialMargin"`  // 持仓所需起始保证金(基于最新标记价格)
		OpenOrderInitialMargin string `json:"openOrderInitialMargin"` // 当前挂单所需起始保证金(基于最新标记价格)
		Leverage               string `json:"leverage"`               // 杠杆倍率
		EntryPrice             string `json:"entryPrice"`             // 持仓成本价
		MaxNotional            string `json:"maxNotional"`            // 当前杠杆下用户可用的最大名义价值
		BidNotional            string `json:"bidNotional"`            // 买单净值
		AskNotional            string `json:"askNotional"`            // 卖单净值
		PositionSide           string `json:"positionSide"`           // 持仓方向
		PositionAmt            string `json:"positionAmt"`            // 持仓数量
		UpdateTime             int64  `json:"updateTime"`             // 更新时间
	} `json:"positions"` // 头寸，将返回所有市场symbol。
}

type PortfolioMarginUmAccountV2Res struct {
	Assets []struct {
		Asset                  string `json:"asset"`                  // 资产
		CrossWalletBalance     string `json:"crossWalletBalance"`     // 全仓账户余额
		CrossUnPnl             string `json:"crossUnPnl"`             // 全仓持仓未实现盈亏
		MaintMargin            string `json:"maintMargin"`            // 维持保证金
		InitialMargin          string `json:"initialMargin"`          // 当前所需起始保证金
		PositionInitialMargin  string `json:"positionInitialMargin"`  // 持仓所需起始保证金(基于最新标记价格)
		OpenOrderInitialMargin string `json:"openOrderInitialMargin"` // 当前挂单所需起始保证金(基于最新标记价格)
		UpdateTime             int64  `json:"updateTime"`             // 更新时间
	} `json:"assets"` // 资产
	Positions []struct {
		Symbol           string `json:"symbol"`           // 交易对
		InitialMargin    string `json:"initialMargin"`    // 当前所需起始保证金(基于最新标记价格)
		MaintMargin      string `json:"maintMargin"`      // 维持保证金
		UnrealizedProfit string `json:"unrealizedProfit"` // 持仓未实现盈亏
		PositionSide     string `json:"positionSide"`     // 持仓方向
		PositionAmt      string `json:"positionAmt"`      // 持仓数量
		UpdateTime       int64  `json:"updateTime"`       // 更新时间
		Notional         string `json:"notional"`         // 名义价值
	} `json:"positions"` // 头寸，将返回所有市场symbol。
}

type PortfolioMarginCmAccountRes struct {
	Assets []struct {
		Asset                  string `json:"asset"`                  // 资产
		CrossWalletBalance     string `json:"crossWalletBalance"`     // 全仓账户余额
		CrossUnPnl             string `json:"crossUnPnl"`             // 全仓持仓未实现盈亏
		MaintMargin            string `json:"maintMargin"`            // 维持保证金
		InitialMargin          string `json:"initialMargin"`          // 当前所需起始保证金
		PositionInitialMargin  string `json:"positionInitialMargin"`  // 持仓所需起始保证金(基于最新标记价格)
		OpenOrderInitialMargin string `json:"openOrderInitialMargin"` // 当前挂单所需起始保证金(基于最新标记价格)
		UpdateTime             int64  `json:"updateTime"`             // 更新时间
	} `json:"assets"`
	Positions []struct {
		Symbol                 string `json:"symbol"`                 // 交易对
		PositionAmt            string `json:"positionAmt"`            // 持仓数量
		InitialMargin          string `json:"initialMargin"`          // 当前所需起始保证金(基于最新标记价格)
		MaintMargin            string `json:"maintMargin"`            // 维持保证金
		UnrealizedProfit       string `json:"unrealizedProfit"`       // 持仓未实现盈亏
		PositionInitialMargin  string `json:"positionInitialMargin"`  // 持仓所需起始保证金(基于最新标记价格)
		OpenOrderInitialMargin string `json:"openOrderInitialMargin"` // 当前挂单所需起始保证金(基于最新标记价格)
		Leverage               string `json:"leverage"`               // 杠杆倍率
		PositionSide           string `json:"positionSide"`           // 持仓方向
		EntryPrice             string `json:"entryPrice"`             // 持仓成本价
		MaxQty                 string `json:"maxQty"`                 // 当前杠杆下用户可用的最大名义价值
		UpdateTime             int64  `json:"updateTime"`             // 更新时间
	}
}

type PortfolioMarginUmPositionRiskResRow struct {
	EntryPrice       string `json:"entryPrice"`       // 持仓成本价
	Leverage         string `json:"leverage"`         // 杠杆倍率
	MarkPrice        string `json:"markPrice"`        // 标记价格
	MaxNotionalValue string `json:"maxNotionalValue"` // 最大名义价值
	PositionAmt      string `json:"positionAmt"`      // 持仓数量
	Notional         string `json:"notional"`         // 名义价值
	Symbol           string `json:"symbol"`           // 交易对
	UnRealizedProfit string `json:"unRealizedProfit"` // 未实现盈亏
	LiquidationPrice string `json:"liquidationPrice"` // 强平价格
	PositionSide     string `json:"positionSide"`     // 持仓方向
	UpdateTime       int64  `json:"updateTime"`       // 更新时间
}

type PortfolioMarginUmPositionRiskRes []PortfolioMarginUmPositionRiskResRow

type PortfolioMarginCmPositionRiskResRow struct {
	Symbol           string `json:"symbol"`           // 交易对
	PositionAmt      string `json:"positionAmt"`      // 持仓数量
	EntryPrice       string `json:"entryPrice"`       // 持仓成本价
	MarkPrice        string `json:"markPrice"`        // 标记价格
	UnRealizedProfit string `json:"unRealizedProfit"` // 未实现盈亏
	LiquidationPrice string `json:"liquidationPrice"` // 强平价格
	Leverage         string `json:"leverage"`         // 杠杆倍率
	PositionSide     string `json:"positionSide"`     // 持仓方向
	UpdateTime       int64  `json:"updateTime"`       // 更新时间
	MaxQty           string `json:"maxQty"`           // 当前杠杆下用户可用的最大名义价值
	NotionalValue    string `json:"notionalValue"`    // 名义价值
}
type PortfolioMarginCmPositionRiskRes []PortfolioMarginCmPositionRiskResRow

type PortfolioMarginUmCommissionRateRes struct {
	Symbol              string `json:"symbol"`              // 交易对
	MakerCommissionRate string `json:"makerCommissionRate"` // maker手续费率
	TakerCommissionRate string `json:"takerCommissionRate"` // taker手续费率
}

type PortfolioMarginCmCommissionRateRes struct {
	Symbol              string `json:"symbol"`              // 交易对
	MakerCommissionRate string `json:"makerCommissionRate"` // maker手续费率
	TakerCommissionRate string `json:"takerCommissionRate"` // taker手续费率
}

type PortfolioMarginAutoCollectionRes struct {
	Msg string `json:"msg"` // 消息
}
type PortfolioMarginAssetCollectionRes struct {
	Msg string `json:"msg"` // 消息
}

// [
//     {
//         "symbol": "",               // 交易对，仅针对涉及交易对的资金流
//         "incomeType": "TRANSFER",   // 资金流类型
//         "income": "-0.37500000",    // 资金流数量，正数代表流入，负数代表流出
//         "asset": "USDT",            // 资产内容
//         "info":"TRANSFER",          // 备注信息，取决于流水类型
//         "time": 1570608000000,
//         "tranId":"9689322392",      // 划转ID
//         "tradeId":""                // 引起流水产生的原始交易ID
//     },
//     {
//         "symbol": "BTCUSDT",
//         "incomeType": "COMMISSION",
//         "income": "-0.01000000",
//         "asset": "USDT",
//         "info":"COMMISSION",
//         "time": 1570636800000,
//         "tranId":"9689322392",
//         "tradeId":"2059192"
//     }
// ]
type PortfolioMarginUmIncomeResRow struct {
	Symbol     string `json:"symbol"`     // 交易对，仅针对涉及交易对的资金流
	IncomeType string `json:"incomeType"` // 资金流类型
	Income     string `json:"income"`     // 资金流数量，正数代表流入，负数代表流出
	Asset      string `json:"asset"`      // 资产内容
	Info       string `json:"info"`       // 备注信息，取决于流水类型
	Time       int64  `json:"time"`       // 时间
	TranId     int64  `json:"tranId"`     // 划转ID
	TradeId    string `json:"tradeId"`    // 引起流水产生的原始交易ID
}
type PortfolioMarginUmIncomeRes []PortfolioMarginUmIncomeResRow

// {
// 	"avgCostTimestampOfLast30d":7241837, //过去30天平均数据下载时间
//   	"downloadId":"546975389218332672",   //下载Id
// }
type PortfolioMarginUmIncomeAsynRes struct {
	AvgCostTimestampOfLast30d int64  `json:"avgCostTimestampOfLast30d"` // 过去30天平均数据下载时间
	DownloadId                string `json:"downloadId"`                // 下载Id
}

// {
// 	"downloadId":"545923594199212032", // 下载Id
//   	"status":"completed",     // 状态，枚举类型：completed 已完成，processing 处理中
//   	"url":"www.binance.com",  // 适配该笔ID请求的下载链接
//     "s3Link": null,
//   	"notified":true,          // 忽略
//   	"expirationTimestamp":1645009771000,  // 晚于该时间戳之后链接将自动失效
//   	"isExpired":null,
// }
type PortfolioMarginUmIncomeAsynIdRes struct {
	DownloadId          string `json:"downloadId"`          // 下载Id
	Status              string `json:"status"`              // 状态，枚举类型：completed 已完成，processing 处理中
	Url                 string `json:"url"`                 // 适配该笔ID请求的下载链接
	S3Link              string `json:"s3Link"`              // 忽略
	Notified            bool   `json:"notified"`            // 忽略
	ExpirationTimestamp int64  `json:"expirationTimestamp"` // 晚于该时间戳之后链接将自动失效
	IsExpired           bool   `json:"isExpired"`           // 忽略
}

type PortfolioMarginRepayFuturesNegativeBalanceRes struct {
	Msg string `json:"msg"` // 消息
}
