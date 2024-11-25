package mybinanceapi

type SpotBrokerSubAccountPostRes struct {
	SubaccountId string `json:"subaccountId"` // 子账户ID
	Email        string `json:"email"`        // 邮箱
	Tag          string `json:"tag"`          // 标签
}
type SpotBrokerSubAccountGetResRow struct {
	SubaccountId          string  `json:"subaccountId"`          // 子账户ID
	Email                 string  `json:"email"`                 // 邮箱
	Tag                   string  `json:"tag"`                   // 标签
	MakerCommission       float64 `json:"makerCommission"`       // maker手续费
	TakerCommission       float64 `json:"takerCommission"`       // taker手续费
	MarginMakerCommission float64 `json:"marginMakerCommission"` // 杠杆maker手续费
	MarginTakerCommission float64 `json:"marginTakerCommission"` // 杠杆taker手续费
	CreateTime            int64   `json:"createTime"`            // 创建时间
}
type SpotBrokerSubAccountGetRes []SpotBrokerSubAccountGetResRow

type SpotBrokerSubAccountFuturesRes struct {
	SubaccountId  string `json:"subaccountId"`  // 子账户ID
	EnableFutures bool   `json:"enableFutures"` // 是否开通合约
	UpdateTime    int64  `json:"updateTime"`    // 更新时间
}

type SpotBrokerSubAccountApiPostRes struct {
	SubaccountId string `json:"subaccountId"` // 子账户ID
	ApiKey       string `json:"apiKey"`       // API Key
	SecretKey    string `json:"secretKey"`    // Secret Key
	CanTrade     bool   `json:"canTrade"`     // 是否可以交易
	MarginTrade  bool   `json:"marginTrade"`  // 是否可以杠杆交易
	FuturesTrade bool   `json:"futuresTrade"` // 是否可以合约交易
}
type SpotBrokerSubAccountApiGetResRow struct {
	ApiKey       string `json:"apiKey"`       // API Key
	SubaccountId string `json:"subaccountId"` // 子账户ID
	CanTrade     bool   `json:"canTrade"`     // 是否可以交易
	MarginTrade  bool   `json:"marginTrade"`  // 是否可以杠杆交易
	FuturesTrade bool   `json:"futuresTrade"` // 是否可以合约交易
}
type SpotBrokerSubAccountApiGetRes []SpotBrokerSubAccountApiGetResRow
type SpotBrokerSubAccountApiDeleteRes struct{}
type SpotBrokerSubAccountPermissionUniversalTransferRes struct {
	SubaccountId         string `json:"subaccountId"`         // 子账户ID
	Apikey               string `json:"apikey"`               // API Key
	CanUniversalTransfer bool   `json:"canUniversalTransfer"` // 是否可以进行资产划转
}

type SpotBrokerSubAccountApiIpRestrictionPostRes struct {
	Status     string   `json:"status"` // IP Restriction status. 1 = IP Unrestricted. 2 = Restrict access to trusted IPs only.
	ApiKey     string   `json:"apiKey"`
	IpList     []string `json:"ipList"`
	UpdateTime int64    `json:"updateTime"`
}

type SpotBrokerSubAccountApiIpRestrictionGetRes struct {
	SubaccountId string   `json:"subaccountId"` // 子账户ID
	IpRestrict   bool     `json:"ipRestrict"`   // 是否开启IP限制
	Apikey       string   `json:"apikey"`       // API Key
	IpList       []string `json:"ipList"`       // IP白名单
	UpdateTime   int64    `json:"updateTime"`   // 更新时间
}

type SpotBrokerSubAccountApiIpRestrictionDeleteRes struct {
	Status     string   `json:"status"` // IP Restriction status. 1 = IP Unrestricted. 2 = Restrict access to trusted IPs only.
	Apikey     string   `json:"apikey"`
	IpList     []string `json:"ipList"`
	UpdateTime int64    `json:"updateTime"`
}

type SpotBrokerSubAccountDepositHistResRow struct {
	DepositId        int64  `json:"depositId"`
	SubAccountId     string `json:"subAccountId"`
	Address          string `json:"address"`
	AddressTag       string `json:"addressTag"`
	Amount           string `json:"amount"`
	Coin             string `json:"coin"`
	InsertTime       int64  `json:"insertTime"`
	TransferType     int    `json:"transferType"`
	Network          string `json:"network"`
	Status           int    `json:"status"`
	TxId             string `json:"txId"`
	SourceAddress    string `json:"sourceAddress"`
	ConfirmTimes     string `json:"confirmTimes"`
	SelfReturnStatus int    `json:"selfReturnStatus"`
}
type SpotBrokerSubAccountDepositHistRes []SpotBrokerSubAccountDepositHistResRow

type SpotBrokerUniversalTransferPostRes struct {
	TxnId        int64  `json:"txnId"`        // 交易ID
	ClientTranId string `json:"clientTranId"` // 客户端交易ID
}
type SpotBrokerUniversalTransferGetResRow struct {
	ToId            string `json:"toId"`            // 目标账户ID
	Asset           string `json:"asset"`           // 资产
	Qty             string `json:"qty"`             // 数量
	Time            int64  `json:"time"`            // 时间
	Status          string `json:"status"`          // 状态
	TxnId           string `json:"txnId"`           // 交易ID
	ClientTranId    string `json:"clientTranId"`    // 客户端交易ID
	FromAccountType string `json:"fromAccountType"` // 转出账户类型
	ToAccountType   string `json:"toAccountType"`   // 转入账户类型
}
