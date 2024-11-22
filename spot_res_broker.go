package mybinanceapi

type SpotBrokerSubAccountGetResRow struct {
	SubAccountId          string  `json:"subAccountId"`          // 子账户ID
	Email                 string  `json:"email"`                 // 邮箱
	Tag                   string  `json:"tag"`                   // 标签
	MakerCommission       float64 `json:"makerCommission"`       // maker手续费
	TakerCommission       float64 `json:"takerCommission"`       // taker手续费
	MarginMakerCommission float64 `json:"marginMakerCommission"` // 杠杆maker手续费
	MarginTakerCommission float64 `json:"marginTakerCommission"` // 杠杆taker手续费
	CreateTime            int64   `json:"createTime"`            // 创建时间
}
type SpotBrokerSubAccountGetRes []SpotBrokerSubAccountGetResRow
