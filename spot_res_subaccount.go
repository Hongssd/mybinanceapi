package mybinanceapi

type SubAccountResSubAccount struct {
	Email                       string `json:"email"`
	IsFreeze                    bool   `json:"isFreeze"`
	CreateTime                  int64  `json:"createTime"`
	IsManagedSubAccount         bool   `json:"isManagedSubAccount"`
	IsAssetManagementSubAccount bool   `json:"isAssetManagementSubAccount"`
}
type SubAccountListRes struct {
	SubAccounts []SubAccountResSubAccount `json:"subAccounts"`
}

type SubAccountUniversalTransferHistoryResResult struct {
	TranId          int64  `json:"tranId"`
	FromEmail       string `json:"fromEmail"`
	ToEmail         string `json:"toEmail"`
	Asset           string `json:"asset"`
	Amount          string `json:"amount"`
	CreateTimeStamp int64  `json:"createTimeStamp"`
	FromAccountType string `json:"fromAccountType"`
	ToAccountType   string `json:"toAccountType"`
	Status          string `json:"status"`
	ClientTranId    string `json:"clientTranId"`
}
type SubAccountUniversalTransferHistoryRes struct {
	Result     []SubAccountUniversalTransferHistoryResResult `json:"result"`
	TotalCount int                                           `json:"totalCount"`
}

type SubAccountAssetsRes struct {
	Balance []SpotBalance `json:"balances"`
}

type SubAccountFuturesAccountAssets struct {
	Asset                  string `json:"asset"`
	InitialMargin          string `json:"initialMargin"`
	MaintenanceMargin      string `json:"maintenanceMargin"`
	MarginBalance          string `json:"marginBalance"`
	MaxWithdrawAmount      string `json:"maxWithdrawAmount"`
	OpenOrderInitialMargin string `json:"openOrderInitialMargin"`
	PositionInitialMargin  string `json:"positionInitialMargin"`
	UnrealizedProfit       string `json:"unrealizedProfit"`
	WalletBalance          string `json:"walletBalance"`
}
type SubAccountFuturesAccountAccountResp struct {
	Email                       string                           `json:"email"`
	Assets                      []SubAccountFuturesAccountAssets `json:"assets"`
	CanDeposit                  bool                             `json:"canDeposit"`
	CanTrade                    bool                             `json:"canTrade"`
	CanWithdraw                 bool                             `json:"canWithdraw"`
	FeeTier                     int                              `json:"feeTier"`
	UpdateTime                  int64                            `json:"updateTime"`
	MaxWithdrawAmount           string                           `json:"maxWithdrawAmount"`
	TotalInitialMargin          string                           `json:"totalInitialMargin"`
	TotalMaintenanceMargin      string                           `json:"totalMaintenanceMargin"`
	TotalMarginBalance          string                           `json:"totalMarginBalance"`
	TotalOpenOrderInitialMargin string                           `json:"totalOpenOrderInitialMargin"`
	TotalPositionInitialMargin  string                           `json:"totalPositionInitialMargin"`
	TotalUnrealizedProfit       string                           `json:"totalUnrealizedProfit"`
	TotalWalletBalance          string                           `json:"totalWalletBalance"`
}
type SubAccountFuturesAccountRes struct {
	FutureAccountResp   SubAccountFuturesAccountAccountResp `json:"futureAccountResp"`
	DeliveryAccountResp SubAccountFuturesAccountAccountResp `json:"deliveryAccountResp"`
}

type SpotSubAccountApiIpRestrictionRes struct {
	IpRestrict string   `json:"ipRestrict"`
	IpList     []string `json:"ipList"`
	UpdateTime int64    `json:"updateTime"`
	ApiKey     string   `json:"apiKey"`
}

type SpotSubAccountTransferSubUserHistoryRes []SpotSubAccountTransferSubUserHistoryResRow
type SpotSubAccountTransferSubUserHistoryResRow struct {
	CounterParty    string `json:"counterParty"`
	Email           string `json:"email"`
	Type            int64  `json:"type"` // 1 for transfer in , 2 for transfer out
	Asset           string `json:"asset"`
	Qty             string `json:"qty"`
	FromAccountType string `json:"fromAccountType"`
	ToAccountType   string `json:"toAccountType"`
	Status          string `json:"status"` // status: PROCESS / SUCCESS / FAILURE
	TranId          int64  `json:"tranId"`
	Time            int64  `json:"time"`
}

type ManagedSubAccountQueryTransLogRes struct {
	ManagerSubTransferHistoryVos []ManagerSubTransferHistoryVos `json:"managerSubTransferHistoryVos"`
}
type ManagerSubTransferHistoryVos struct {
	FromEmail       string `json:"fromEmail"`
	FromAccountType string `json:"fromAccountType"`
	ToEmail         string `json:"toEmail"`
	ToAccountType   string `json:"toAccountType"`
	Asset           string `json:"asset"`
	Amount          string `json:"amount"`
	ScheduledData   int64  `json:"scheduledData"`
	CreateTime      int64  `json:"createTime"`
	Status          string `json:"status"`
	TranId          int64  `json:"tranId"`
}

type SubAccountVirtualSubAccountRes struct {
	Email string `json:"email"`
}

type SubAccountUniversalTransferRes struct {
	TranId       int64  `json:"tranId"`
	ClientTranId string `json:"clientTranId"`
}

type SubAccountFuturesEnableRes struct {
	Email            string `json:"email"`
	IsFuturesEnabled bool   `json:"isFuturesEnabled"`
}

type SpotSubAccountMarginEnableRes struct {
	Email           string `json:"email"`
	IsMarginEnabled bool   `json:"isMarginEnabled"`
}

type SpotSubAccountSubAccountApiIpRestrictionPostRes struct {
	Status     string   `json:"status"`     // IP Restriction status. 1 = IP Unrestricted. 2 = Restrict access to trusted IPs only.
	ApiKey     string   `json:"apiKey"`     // API Key
	IpList     []string `json:"ipList"`     // IP白名单
	UpdateTime int64    `json:"updateTime"` // 更新时间
}
type SpotSubAccountSubAccountApiIpRestrictionGetRes struct {
	IpRestrict string   `json:"ipRestrict"` // 是否开启IP限制 "true" or "false"
	ApiKey     string   `json:"apiKey"`     // API Key
	IpList     []string `json:"ipList"`     // IP白名单
	UpdateTime int64    `json:"updateTime"` // 更新时间
}
type SpotSubAccountSubAccountApiIpRestrictionDeleteRes struct {
	IpRestrict string   `json:"ipRestrict"` // 是否开启IP限制 "true" or "false"
	ApiKey     string   `json:"apiKey"`     // API Key
	IpList     []string `json:"ipList"`     // IP白名单
	UpdateTime int64    `json:"updateTime"` // 更新时间
}

type SpotSubAccountCapitalDepositSubHisrecResRow struct {
	Id            string `json:"id"`
	Amount        string `json:"amount"`
	Coin          string `json:"coin"`
	Network       string `json:"network"`
	Status        int    `json:"status"`
	Address       string `json:"address"`
	AddressTag    string `json:"addressTag"`
	TxId          string `json:"txId"`
	InsertTime    int64  `json:"insertTime"`
	TransferType  int    `json:"transferType"`
	ConfirmTimes  string `json:"confirmTimes"`
	UnlockConfirm int    `json:"unlockConfirm"`
	WalletType    int    `json:"walletType"`
}
type SpotSubAccountCapitalDepositSubHisrecRes []SpotSubAccountCapitalDepositSubHisrecResRow
