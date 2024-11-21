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
