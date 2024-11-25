package mybinanceapi

type SpotBrokerSubAccountPostReq struct {
	Tag        *string `json:"tag,omitempty"`        // NO 经纪商代码
	RecvWindow *int64  `json:"recvWindow,omitempty"` // NO 接收窗口
	Timestamp  *int64  `json:"timestamp,omitempty"`  // YES 时间戳
}

// 经纪商代码
func (api *SpotBrokerSubAccountPostApi) Tag(tag string) *SpotBrokerSubAccountPostApi {
	api.req.Tag = GetPointer(tag)
	return api
}

// 接收窗口
func (api *SpotBrokerSubAccountPostApi) RecvWindow(recvWindow int64) *SpotBrokerSubAccountPostApi {
	api.req.RecvWindow = GetPointer(recvWindow)
	return api
}

// 时间戳
func (api *SpotBrokerSubAccountPostApi) Timestamp(timestamp int64) *SpotBrokerSubAccountPostApi {
	api.req.Timestamp = GetPointer(timestamp)
	return api
}

type SpotBrokerSubAccountPostApi struct {
	client *SpotRestClient
	req    *SpotBrokerSubAccountPostReq
}

type SpotBrokerSubAccountGetReq struct {
	SubAccountId *string `json:"subAccountId,omitempty"` // NO 子账户ID
	Page         *int    `json:"page,omitempty"`         // NO 页码
	Size         *int    `json:"size,omitempty"`         // NO 每页数量
	RevcWindow   *int64  `json:"recvWindow,omitempty"`   // NO 接收窗口
	Timestamp    *int64  `json:"timestamp,omitempty"`    // YES 时间戳
}

// 子账户ID
func (api *SpotBrokerSubAccountGetApi) SubAccountId(subAccountId string) *SpotBrokerSubAccountGetApi {
	api.req.SubAccountId = GetPointer(subAccountId)
	return api
}

// 页码
func (api *SpotBrokerSubAccountGetApi) Page(page int) *SpotBrokerSubAccountGetApi {
	api.req.Page = GetPointer(page)
	return api
}

// 每页数量
func (api *SpotBrokerSubAccountGetApi) Size(size int) *SpotBrokerSubAccountGetApi {
	api.req.Size = GetPointer(size)
	return api
}

// 接收窗口
func (api *SpotBrokerSubAccountGetApi) RevcWindow(recvWindow int64) *SpotBrokerSubAccountGetApi {
	api.req.RevcWindow = GetPointer(recvWindow)
	return api
}

// 时间戳
func (api *SpotBrokerSubAccountGetApi) Timestamp(timestamp int64) *SpotBrokerSubAccountGetApi {
	api.req.Timestamp = GetPointer(timestamp)
	return api
}

type SpotBrokerSubAccountGetApi struct {
	client *SpotRestClient
	req    *SpotBrokerSubAccountGetReq
}

type SpotBrokerSubAccountFuturesReq struct {
	SubAccountId *string `json:"subAccountId,omitempty"` // YES 子账户ID
	Futures      *bool   `json:"futures,omitempty"`      // YES 是否开通合约 only true for now
	RecvWindow   *int64  `json:"recvWindow,omitempty"`   // NO 接收窗口
	Timestamp    *int64  `json:"timestamp,omitempty"`    // YES 时间戳
}

// 子账户ID
func (api *SpotBrokerSubAccountFuturesApi) SubAccountId(subAccountId string) *SpotBrokerSubAccountFuturesApi {
	api.req.SubAccountId = GetPointer(subAccountId)
	return api
}

// 是否开通合约 only true for now
func (api *SpotBrokerSubAccountFuturesApi) Futures(futures bool) *SpotBrokerSubAccountFuturesApi {
	api.req.Futures = GetPointer(futures)
	return api
}

// 接收窗口
func (api *SpotBrokerSubAccountFuturesApi) RecvWindow(recvWindow int64) *SpotBrokerSubAccountFuturesApi {
	api.req.RecvWindow = GetPointer(recvWindow)
	return api
}

// 时间戳
func (api *SpotBrokerSubAccountFuturesApi) Timestamp(timestamp int64) *SpotBrokerSubAccountFuturesApi {
	api.req.Timestamp = GetPointer(timestamp)
	return api
}

type SpotBrokerSubAccountFuturesApi struct {
	client *SpotRestClient
	req    *SpotBrokerSubAccountFuturesReq
}

type SpotBrokerSubAccountApiPostReq struct {
	SubAccountId *string `json:"subAccountId,omitempty"` // YES 子账户ID
	CanTrade     *bool   `json:"canTrade,omitempty"`     // YES ENUM 是否允许交易
	MarginTrade  *bool   `json:"marginTrade,omitempty"`  // YES ENUM 是否允许杠杆交易
	FuturesTrade *bool   `json:"futuresTrade,omitempty"` // YES ENUM 是否允许合约交易
	RecvWindow   *int64  `json:"recvWindow,omitempty"`   // NO 接收窗口
	Timestamp    *int64  `json:"timestamp,omitempty"`    // YES 时间戳
}

// 子账户ID
func (api *SpotBrokerSubAccountApiPostApi) SubAccountId(subAccountId string) *SpotBrokerSubAccountApiPostApi {
	api.req.SubAccountId = GetPointer(subAccountId)
	return api
}

// 是否允许交易
func (api *SpotBrokerSubAccountApiPostApi) CanTrade(canTrade bool) *SpotBrokerSubAccountApiPostApi {
	api.req.CanTrade = GetPointer(canTrade)
	return api
}

// 是否允许杠杆交易
func (api *SpotBrokerSubAccountApiPostApi) MarginTrade(marginTrade bool) *SpotBrokerSubAccountApiPostApi {
	api.req.MarginTrade = GetPointer(marginTrade)
	return api
}

// 是否允许合约交易
func (api *SpotBrokerSubAccountApiPostApi) FuturesTrade(futuresTrade bool) *SpotBrokerSubAccountApiPostApi {
	api.req.FuturesTrade = GetPointer(futuresTrade)
	return api
}

// 接收窗口
func (api *SpotBrokerSubAccountApiPostApi) RecvWindow(recvWindow int64) *SpotBrokerSubAccountApiPostApi {
	api.req.RecvWindow = GetPointer(recvWindow)
	return api
}

// 时间戳
func (api *SpotBrokerSubAccountApiPostApi) Timestamp(timestamp int64) *SpotBrokerSubAccountApiPostApi {
	api.req.Timestamp = GetPointer(timestamp)
	return api
}

type SpotBrokerSubAccountApiPostApi struct {
	client *SpotRestClient
	req    *SpotBrokerSubAccountApiPostReq
}
type SpotBrokerSubAccountApiGetReq struct {
	SubAccountId     *string `json:"subAccountId,omitempty"`     // YES 子账户ID
	SubAccountApiKey *string `json:"subAccountApiKey,omitempty"` // NO 子账户API Key
	Page             *int    `json:"page,omitempty"`             // NO 页码
	Size             *int    `json:"size,omitempty"`             // NO 每页数量
	RecvWindow       *int64  `json:"recvWindow,omitempty"`       // NO 接收窗口
	Timestamp        *int64  `json:"timestamp,omitempty"`        // YES 时间戳
}

// 子账户ID
func (api *SpotBrokerSubAccountApiGetApi) SubAccountId(subAccountId string) *SpotBrokerSubAccountApiGetApi {
	api.req.SubAccountId = GetPointer(subAccountId)
	return api
}

// 子账户API Key
func (api *SpotBrokerSubAccountApiGetApi) SubAccountApiKey(subAccountApiKey string) *SpotBrokerSubAccountApiGetApi {
	api.req.SubAccountApiKey = GetPointer(subAccountApiKey)
	return api
}

// 页码 default 1
func (api *SpotBrokerSubAccountApiGetApi) Page(page int) *SpotBrokerSubAccountApiGetApi {
	api.req.Page = GetPointer(page)
	return api
}

// 每页数量 default 500, max 500
func (api *SpotBrokerSubAccountApiGetApi) Size(size int) *SpotBrokerSubAccountApiGetApi {
	api.req.Size = GetPointer(size)
	return api
}

// 接收窗口
func (api *SpotBrokerSubAccountApiGetApi) RecvWindow(recvWindow int64) *SpotBrokerSubAccountApiGetApi {
	api.req.RecvWindow = GetPointer(recvWindow)
	return api
}

// 时间戳
func (api *SpotBrokerSubAccountApiGetApi) Timestamp(timestamp int64) *SpotBrokerSubAccountApiGetApi {
	api.req.Timestamp = GetPointer(timestamp)
	return api
}

type SpotBrokerSubAccountApiGetApi struct {
	client *SpotRestClient
	req    *SpotBrokerSubAccountApiGetReq
}

type SpotBrokerSubAccountApiDeleteReq struct {
	SubAccountId     *string `json:"subAccountId,omitempty"`     // YES 子账户ID
	SubAccountApiKey *string `json:"subAccountApiKey,omitempty"` // YES 子账户API Key
	RecvWindow       *int64  `json:"recvWindow,omitempty"`       // NO 接收窗口
	Timestamp        *int64  `json:"timestamp,omitempty"`        // YES 时间戳
}

// 子账户ID
func (api *SpotBrokerSubAccountApiDeleteApi) SubAccountId(subAccountId string) *SpotBrokerSubAccountApiDeleteApi {
	api.req.SubAccountId = GetPointer(subAccountId)
	return api
}

// 子账户API Key
func (api *SpotBrokerSubAccountApiDeleteApi) SubAccountApiKey(subAccountApiKey string) *SpotBrokerSubAccountApiDeleteApi {
	api.req.SubAccountApiKey = GetPointer(subAccountApiKey)
	return api
}

// 接收窗口
func (api *SpotBrokerSubAccountApiDeleteApi) RecvWindow(recvWindow int64) *SpotBrokerSubAccountApiDeleteApi {
	api.req.RecvWindow = GetPointer(recvWindow)
	return api
}

// 时间戳
func (api *SpotBrokerSubAccountApiDeleteApi) Timestamp(timestamp int64) *SpotBrokerSubAccountApiDeleteApi {
	api.req.Timestamp = GetPointer(timestamp)
	return api
}

type SpotBrokerSubAccountApiDeleteApi struct {
	client *SpotRestClient
	req    *SpotBrokerSubAccountApiDeleteReq
}

type SpotBrokerSubAccountPermissionUniversalTransferReq struct {
	SubAccountId         *int64  `json:"subAccountId,omitempty"`         // YES 子账户ID
	SubAccountApiKey     *string `json:"subAccountApiKey,omitempty"`     // YES 子账户API Key
	CanUniversalTransfer *bool   `json:"canUniversalTransfer,omitempty"` // YES 是否可以万能划转 true or false
	RecvWindow           *int64  `json:"recvWindow,omitempty"`           // NO 接收窗口
	Timestamp            *int64  `json:"timestamp,omitempty"`            // YES 时间戳
}

// 子账户ID
func (api *SpotBrokerSubAccountPermissionUniversalTransferApi) SubAccountId(subAccountId int64) *SpotBrokerSubAccountPermissionUniversalTransferApi {
	api.req.SubAccountId = GetPointer(subAccountId)
	return api
}

// 子账户API Key
func (api *SpotBrokerSubAccountPermissionUniversalTransferApi) SubAccountApiKey(subAccountApiKey string) *SpotBrokerSubAccountPermissionUniversalTransferApi {
	api.req.SubAccountApiKey = GetPointer(subAccountApiKey)
	return api
}

// 是否可以万能划转 true or false
func (api *SpotBrokerSubAccountPermissionUniversalTransferApi) CanUniversalTransfer(canUniversalTransfer bool) *SpotBrokerSubAccountPermissionUniversalTransferApi {
	api.req.CanUniversalTransfer = GetPointer(canUniversalTransfer)
	return api
}

// 接收窗口
func (api *SpotBrokerSubAccountPermissionUniversalTransferApi) RecvWindow(recvWindow int64) *SpotBrokerSubAccountPermissionUniversalTransferApi {
	api.req.RecvWindow = GetPointer(recvWindow)
	return api
}

// 时间戳
func (api *SpotBrokerSubAccountPermissionUniversalTransferApi) Timestamp(timestamp int64) *SpotBrokerSubAccountPermissionUniversalTransferApi {
	api.req.Timestamp = GetPointer(timestamp)
	return api
}

type SpotBrokerSubAccountPermissionUniversalTransferApi struct {
	client *SpotRestClient
	req    *SpotBrokerSubAccountPermissionUniversalTransferReq
}

// subAccountId	STRING	YES
// subAccountApiKey	STRING	YES
// status	STRING	YES	IP Restriction status. 1 = IP Unrestricted. 2 = Restrict access to trusted IPs only.
// ipAddress	STRING	NO	Insert static IP in batch, separated by commas.
// recvWindow	LONG	NO
// timestamp	LONG	YES
type SpotBrokerSubAccountApiIpRestrictionReq struct {
	SubAccountId     *string `json:"subAccountId,omitempty"`     // YES 子账户ID
	SubAccountApiKey *string `json:"subAccountApiKey,omitempty"` // YES 子账户API Key
	Status           *string `json:"status,omitempty"`           // YES IP Restriction status. 1 = IP Unrestricted. 2 = Restrict access to trusted IPs only.
	IpAddress        *string `json:"ipAddress,omitempty"`        // NO Insert static IP in batch, separated by commas.
	RecvWindow       *int64  `json:"recvWindow,omitempty"`       // NO 接收窗口
	Timestamp        *int64  `json:"timestamp,omitempty"`        // YES 时间戳
}

// 子账户ID
func (api *SpotBrokerSubAccountApiIpRestrictionApi) SubAccountId(subAccountId string) *SpotBrokerSubAccountApiIpRestrictionApi {
	api.req.SubAccountId = GetPointer(subAccountId)
	return api
}

// 子账户API Key
func (api *SpotBrokerSubAccountApiIpRestrictionApi) SubAccountApiKey(subAccountApiKey string) *SpotBrokerSubAccountApiIpRestrictionApi {
	api.req.SubAccountApiKey = GetPointer(subAccountApiKey)
	return api
}

// IP Restriction status. 1 = IP Unrestricted. 2 = Restrict access to trusted IPs only.
func (api *SpotBrokerSubAccountApiIpRestrictionApi) Status(status string) *SpotBrokerSubAccountApiIpRestrictionApi {
	api.req.Status = GetPointer(status)
	return api
}

// Insert static IP in batch, separated by commas.
func (api *SpotBrokerSubAccountApiIpRestrictionApi) IpAddress(ipAddress string) *SpotBrokerSubAccountApiIpRestrictionApi {
	api.req.IpAddress = GetPointer(ipAddress)
	return api
}

// 接收窗口
func (api *SpotBrokerSubAccountApiIpRestrictionApi) RecvWindow(recvWindow int64) *SpotBrokerSubAccountApiIpRestrictionApi {
	api.req.RecvWindow = GetPointer(recvWindow)
	return api
}

// 时间戳
func (api *SpotBrokerSubAccountApiIpRestrictionApi) Timestamp(timestamp int64) *SpotBrokerSubAccountApiIpRestrictionApi {
	api.req.Timestamp = GetPointer(timestamp)
	return api
}

type SpotBrokerSubAccountApiIpRestrictionApi struct {
	client *SpotRestClient
	req    *SpotBrokerSubAccountApiIpRestrictionReq
}

type SpotBrokerSubAccountApiIpRestrictionGetReq struct {
	SubAccountId     *string `json:"subAccountId,omitempty"`     // YES 子账户ID
	SubAccountApiKey *string `json:"subAccountApiKey,omitempty"` // YES 子账户API Key
	RecvWindow       *int64  `json:"recvWindow,omitempty"`       // NO 接收窗口
	Timestamp        *int64  `json:"timestamp,omitempty"`        // YES 时间戳
}

// 子账户ID
func (api *SpotBrokerSubAccountApiIpRestrictionGetApi) SubAccountId(subAccountId string) *SpotBrokerSubAccountApiIpRestrictionGetApi {
	api.req.SubAccountId = GetPointer(subAccountId)
	return api
}

// 子账户API Key
func (api *SpotBrokerSubAccountApiIpRestrictionGetApi) SubAccountApiKey(subAccountApiKey string) *SpotBrokerSubAccountApiIpRestrictionGetApi {
	api.req.SubAccountApiKey = GetPointer(subAccountApiKey)
	return api
}

// 接收窗口
func (api *SpotBrokerSubAccountApiIpRestrictionGetApi) RecvWindow(recvWindow int64) *SpotBrokerSubAccountApiIpRestrictionGetApi {
	api.req.RecvWindow = GetPointer(recvWindow)
	return api
}

// 时间戳
func (api *SpotBrokerSubAccountApiIpRestrictionGetApi) Timestamp(timestamp int64) *SpotBrokerSubAccountApiIpRestrictionGetApi {
	api.req.Timestamp = GetPointer(timestamp)
	return api
}

type SpotBrokerSubAccountApiIpRestrictionGetApi struct {
	client *SpotRestClient
	req    *SpotBrokerSubAccountApiIpRestrictionGetReq
}

type SpotBrokerSubAccountApiIpRestrictionDeleteReq struct {
	SubAccountId     *string `json:"subAccountId,omitempty"`     // YES 子账户ID
	SubAccountApiKey *string `json:"subAccountApiKey,omitempty"` // YES 子账户API Key
	IpAddress        *string `json:"ipAddress,omitempty"`        // NO Insert static IP in batch, separated by commas.
	RecvWindow       *int64  `json:"recvWindow,omitempty"`       // NO 接收窗口
	Timestamp        *int64  `json:"timestamp,omitempty"`        // YES 时间戳
}

// 子账户ID
func (api *SpotBrokerSubAccountApiIpRestrictionDeleteApi) SubAccountId(subAccountId string) *SpotBrokerSubAccountApiIpRestrictionDeleteApi {
	api.req.SubAccountId = GetPointer(subAccountId)
	return api
}

// 子账户API Key
func (api *SpotBrokerSubAccountApiIpRestrictionDeleteApi) SubAccountApiKey(subAccountApiKey string) *SpotBrokerSubAccountApiIpRestrictionDeleteApi {
	api.req.SubAccountApiKey = GetPointer(subAccountApiKey)
	return api
}

// Insert static IP in batch, separated by commas.
func (api *SpotBrokerSubAccountApiIpRestrictionDeleteApi) IpAddress(ipAddress string) *SpotBrokerSubAccountApiIpRestrictionDeleteApi {
	api.req.IpAddress = GetPointer(ipAddress)
	return api
}

// 接收窗口
func (api *SpotBrokerSubAccountApiIpRestrictionDeleteApi) RecvWindow(recvWindow int64) *SpotBrokerSubAccountApiIpRestrictionDeleteApi {
	api.req.RecvWindow = GetPointer(recvWindow)
	return api
}

// 时间戳
func (api *SpotBrokerSubAccountApiIpRestrictionDeleteApi) Timestamp(timestamp int64) *SpotBrokerSubAccountApiIpRestrictionDeleteApi {
	api.req.Timestamp = GetPointer(timestamp)
	return api
}

type SpotBrokerSubAccountApiIpRestrictionDeleteApi struct {
	client *SpotRestClient
	req    *SpotBrokerSubAccountApiIpRestrictionDeleteReq
}
