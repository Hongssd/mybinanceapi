package mybinanceapi

// 账户接口
type PortfolioMarginBalanceReq struct {
	Asset      *string `json:"asset,omitempty"`      // NO 资产
	RevcWindow *int64  `json:"recvWindow,omitempty"` // NO 接收窗口
	Timestamp  *int64  `json:"timestamp,omitempty"`  // YES 时间戳
}

// 资产
func (api *PortfolioMarginBalanceApi) Asset(asset string) *PortfolioMarginBalanceApi {
	api.req.Asset = GetPointer(asset)
	return api
}

// 接收窗口
func (api *PortfolioMarginBalanceApi) RevcWindow(recvWindow int64) *PortfolioMarginBalanceApi {
	api.req.RevcWindow = GetPointer(recvWindow)
	return api
}

// 时间戳
func (api *PortfolioMarginBalanceApi) Timestamp(timestamp int64) *PortfolioMarginBalanceApi {
	api.req.Timestamp = GetPointer(timestamp)
	return api
}

type PortfolioMarginBalanceApi struct {
	client *PortfolioMarginRestClient
	req    *PortfolioMarginBalanceReq
}
