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
