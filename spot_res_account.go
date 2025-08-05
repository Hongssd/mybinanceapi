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

type SpotAssetGetFundingAssetRes []SpotAssetGetFundingAssetResRow
type SpotAssetGetFundingAssetResRow struct {
	Asset        string `json:"asset"`
	Free         string `json:"free"`         // 可用余额
	Locked       string `json:"locked"`       // 锁定资金
	Freeze       string `json:"freeze"`       //冻结资金
	Withdrawing  string `json:"withdrawing"`  // 提币
	BtcValuation string `json:"btcValuation"` // btc估值
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

type SpotCapitalDepositAddressRes struct {
	Address string `json:"address"`
	Coin    string `json:"coin"`
	Tag     string `json:"tag"`
	Url     string `json:"url"`
}
type SpotCapitalDepositHisrecResRow struct {
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
type SpotCapitalDepositHisrecRes []SpotCapitalDepositHisrecResRow

type SpotAssetDustBtcRes struct {
	Details            []SpotAssetDustBtcDetail `json:"details"`
	TotalTransferBTC   string                   `json:"totalTransferBTC"`
	TotalTransferBNB   string                   `json:"totalTransferBNB"`
	DribbletPercentage string                   `json:"dribbletPercentage"`
}

type SpotAssetDustBtcDetail struct {
	Asset            string `json:"asset"`
	AssetFullName    string `json:"assetFullName"`
	AmountFree       string `json:"amountFree"`
	ToBTC            string `json:"toBTC"`
	ToBNB            string `json:"toBNB"`
	ToBNBOffExchange string `json:"toBNBOffExchange"`
	Exchange         string `json:"exchange"`
}

type SpotAssetDustRes struct {
	TotalServiceCharge string                `json:"totalServiceCharge"`
	TotalTransfered    string                `json:"totalTransfered"`
	TransferResult     []SpotAssetDustResRow `json:"transferResult"`
}

type SpotAssetDustResRow struct {
	Amount              string `json:"amount"`
	FromAsset           string `json:"fromAsset"`
	OperateTime         int64  `json:"operateTime"`
	ServiceChargeAmount string `json:"serviceChargeAmount"`
	TranId              int64  `json:"tranId"`
	TransferedAmount    string `json:"transferedAmount"`
}

// [
//
//	{
//	  "asset": "ETH",
//	  "interest": "0.00083334",
//	  "principal": "0.001",
//	  "liabilityAsset": "USDT",
//	  "liabilityQty": 0.3552
//	}
//
// ]
type SpotMarginExchangeSmallLiabilityGetRes []SpotMarginExchangeSmallLiabilityGetResRow
type SpotMarginExchangeSmallLiabilityGetResRow struct {
	Asset          string `json:"asset"`
	Interest       string `json:"interest"`
	Principal      string `json:"principal"`
	LiabilityAsset string `json:"liabilityAsset"`
	LiabilityQty   string `json:"liabilityQty"`
}

type SpotMarginExchangeSmallLiabilityPostRes struct{}

// [
//         {
//             "subaccountId":"1",
//             "income": "0.02063898",
//             "asset":"BTC",
//             "symbol": "ETHBTC",
//             "tradeId": 123456,
//             "time":1544433328000,
//             "status": 1
//         },
//         {
//             "subaccountId":"2",
//             "income": "1.2063898",
//             "asset":"USDT",
//             "symbol": "BTCUSDT",
//             "tradeId": 223456,
//             "time":1581580800000,
//             "status": 1
//         }

// ]

type SpotBrokerRebateRecentRecordRes []SpotBrokerRebateRecentRecordResRow
type SpotBrokerRebateRecentRecordResRow struct {
	SubAccountId string `json:"subaccountId"`
	Income       string `json:"income"`
	Asset        string `json:"asset"`
	Symbol       string `json:"symbol"`
	TradeId      int64  `json:"tradeId"`
	Time         int64  `json:"time"`
	Status       int64  `json:"status"`
}

// [
//
//	{
//	   "subaccountId": "1",
//	   "income": "0.02063898",
//	   "asset": "USDT",
//	   "symbol": "ETHUSDT",
//	   "tradeId": 123456,
//	   "time": 1544433328000,
//	   "status": 1
//	 },
//	 {
//	   "subaccountId": "2",
//	   "income": "0.02060008",
//	   "asset": "USDT",
//	   "symbol": "BTCUSDT",
//	   "tradeId": 223456,
//	   "time": 1544433328000,
//	   "status": 1
//	 }
//
// ]
type SpotBrokerRebateFuturesRecentRecordRes []SpotBrokerRebateFuturesRecentRecordResRow
type SpotBrokerRebateFuturesRecentRecordResRow struct {
	SubAccountId string `json:"subaccountId"`
	Income       string `json:"income"`
	Asset        string `json:"asset"`
	Symbol       string `json:"symbol"`
	TradeId      int64  `json:"tradeId"`
	Time         int64  `json:"time"`
	Status       int64  `json:"status"`
}
