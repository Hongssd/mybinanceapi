package mybinanceapi

type SpotAccountApiTradingStatusRes struct {
	Data SpotAccountApiTradingStatusData `json:"data"`
}
type SpotAccountApiTradingStatusData struct {
	IsLocked           bool                                        `json:"isLocked"`
	PlannedRecoverTime int64                                       `json:"plannedRecoverTime"`
	TriggerCondition   SpotAccountApiTradingStatusTriggerCondition `json:"triggerCondition"`
	UpdateTime         int64                                       `json:"updateTime"`
}
type SpotAccountApiTradingStatusTriggerCondition struct {
	GCR  int64 `json:"GCR"`
	IFER int64 `json:"IFER"`
	UFR  int64 `json:"UFR"`
}

type SpotAccountCommissionRates struct {
	Maker  string `json:"maker"`
	Taker  string `json:"taker"`
	Buyer  string `json:"buyer"`
	Seller string `json:"seller"`
}
type SpotBalance struct {
	Asset  string `json:"asset"`
	Free   string `json:"free"`
	Locked string `json:"locked"`
}
type SpotAccountRes struct {
	MakerCommission            int                        `json:"makerCommission"`
	TakerCommission            int                        `json:"takerCommission"`
	BuyerCommission            int                        `json:"buyerCommission"`
	SellerCommission           int                        `json:"sellerCommission"`
	CommissionRates            SpotAccountCommissionRates `json:"commissionRates"`
	CanTrade                   bool                       `json:"canTrade"`
	CanWithdraw                bool                       `json:"canWithdraw"`
	CanDeposit                 bool                       `json:"canDeposit"`
	Brokered                   bool                       `json:"brokered"`
	RequireSelfTradePrevention bool                       `json:"requireSelfTradePrevention"`
	PreventSor                 bool                       `json:"preventSor"`
	UpdateTime                 int64                      `json:"updateTime"`
	AccountType                string                     `json:"accountType"`
	Balance                    []SpotBalance              `json:"balances"`
	Permissions                []string                   `json:"permissions"`
	Uid                        int64                      `json:"uid"`
}

type SpotAssetTransferPostRes struct {
	TranId int64 `json:"tranId"`
}

type SpotAssetTransferGetRes struct {
	Total int                       `json:"total"`
	Rows  []SpotAssetTransferGetRow `json:"rows"`
}
type SpotAssetTransferGetRow struct {
	Asset     string `json:"asset"`
	Amount    string `json:"amount"`
	Type      string `json:"type"`
	Status    string `json:"status"`
	TranId    int64  `json:"tranId"`
	Timestamp int64  `json:"timestamp"`
}

type SpotAssetTradeFeeResRow struct {
	Symbol          string `json:"symbol"`
	MakerCommission string `json:"makerCommission"`
	TakerCommission string `json:"takerCommission"`
}
type SpotAssetTradeFeeRes []SpotAssetTradeFeeResRow
