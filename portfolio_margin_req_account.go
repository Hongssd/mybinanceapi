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

type PortfolioMarginAccountReq struct {
	RevcWindow *int64 `json:"recvWindow,omitempty"` // NO 接收窗口
	Timestamp  *int64 `json:"timestamp,omitempty"`  // YES 时间戳
}

// 接收窗口
func (api *PortfolioMarginAccountApi) RevcWindow(recvWindow int64) *PortfolioMarginAccountApi {
	api.req.RevcWindow = GetPointer(recvWindow)
	return api
}

// 时间戳
func (api *PortfolioMarginAccountApi) Timestamp(timestamp int64) *PortfolioMarginAccountApi {
	api.req.Timestamp = GetPointer(timestamp)
	return api
}

type PortfolioMarginAccountApi struct {
	client *PortfolioMarginRestClient
	req    *PortfolioMarginAccountReq
}

type PortfolioMarginMaxBorrowableReq struct {
	Asset      *string `json:"asset,omitempty"`      // YES 资产
	RevcWindow *int64  `json:"recvWindow,omitempty"` // NO 接收窗口
	Timestamp  *int64  `json:"timestamp,omitempty"`  // YES 时间戳
}

// 资产
func (api *PortfolioMarginMaxBorrowableApi) Asset(asset string) *PortfolioMarginMaxBorrowableApi {
	api.req.Asset = GetPointer(asset)
	return api
}

// 接收窗口
func (api *PortfolioMarginMaxBorrowableApi) RevcWindow(recvWindow int64) *PortfolioMarginMaxBorrowableApi {
	api.req.RevcWindow = GetPointer(recvWindow)
	return api
}

// 时间戳
func (api *PortfolioMarginMaxBorrowableApi) Timestamp(timestamp int64) *PortfolioMarginMaxBorrowableApi {
	api.req.Timestamp = GetPointer(timestamp)
	return api
}

type PortfolioMarginMaxBorrowableApi struct {
	client *PortfolioMarginRestClient
	req    *PortfolioMarginMaxBorrowableReq
}

type PortfolioMarginMaxWithdrawReq struct {
	Asset      *string `json:"asset,omitempty"`      // YES 资产
	RevcWindow *int64  `json:"recvWindow,omitempty"` // NO 接收窗口
	Timestamp  *int64  `json:"timestamp,omitempty"`  // YES 时间戳
}

// 资产
func (api *PortfolioMarginMaxWithdrawApi) Asset(asset string) *PortfolioMarginMaxWithdrawApi {
	api.req.Asset = GetPointer(asset)
	return api
}

// 接收窗口
func (api *PortfolioMarginMaxWithdrawApi) RevcWindow(recvWindow int64) *PortfolioMarginMaxWithdrawApi {
	api.req.RevcWindow = GetPointer(recvWindow)
	return api
}

// 时间戳
func (api *PortfolioMarginMaxWithdrawApi) Timestamp(timestamp int64) *PortfolioMarginMaxWithdrawApi {
	api.req.Timestamp = GetPointer(timestamp)
	return api
}

type PortfolioMarginMaxWithdrawApi struct {
	client *PortfolioMarginRestClient
	req    *PortfolioMarginMaxWithdrawReq
}
