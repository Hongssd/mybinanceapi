package mybinanceapi

import "github.com/shopspring/decimal"

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

type SpotBrokerSubAccountDepositHistReq struct {
	SubAccountId *string `json:"subAccountId,omitempty"` // NO
	Coin         *string `json:"coin,omitempty"`         // NO
	Status       *int    `json:"status,omitempty"`       // NO 0(0:pending,6: credited but cannot withdraw, 1:success)
	StartTime    *int64  `json:"startTime,omitempty"`    // NO Default: 7 days from current timestamp
	EndTime      *int64  `json:"endTime,omitempty"`      // NO Default: present timestamp
	Limit        *int    `json:"limit,omitempty"`        // NO Default：500
	Offset       *int    `json:"offset,omitempty"`       // NO Default：0
	RecvWindow   *int64  `json:"recvWindow,omitempty"`   // NO
	Timestamp    *int64  `json:"timestamp,omitempty"`    // YES
}

// NO 子账户ID
func (api *SpotBrokerSubAccountDepositHistApi) SubAccountId(subAccountId string) *SpotBrokerSubAccountDepositHistApi {
	api.req.SubAccountId = GetPointer(subAccountId)
	return api
}

// NO
func (api *SpotBrokerSubAccountDepositHistApi) Coin(coin string) *SpotBrokerSubAccountDepositHistApi {
	api.req.Coin = GetPointer(coin)
	return api
}

// NO 0(0:pending,6: credited but cannot withdraw, 1:success)
func (api *SpotBrokerSubAccountDepositHistApi) Status(status int) *SpotBrokerSubAccountDepositHistApi {
	api.req.Status = GetPointer(status)
	return api
}

// NO Default: 7 days from current timestamp
func (api *SpotBrokerSubAccountDepositHistApi) StartTime(startTime int64) *SpotBrokerSubAccountDepositHistApi {
	api.req.StartTime = GetPointer(startTime)
	return api
}

// NO Default: present timestamp
func (api *SpotBrokerSubAccountDepositHistApi) EndTime(endTime int64) *SpotBrokerSubAccountDepositHistApi {
	api.req.EndTime = GetPointer(endTime)
	return api
}

// NO Default：500
func (api *SpotBrokerSubAccountDepositHistApi) Limit(limit int) *SpotBrokerSubAccountDepositHistApi {
	api.req.Limit = GetPointer(limit)
	return api
}

// NO Default：0
func (api *SpotBrokerSubAccountDepositHistApi) Offset(offset int) *SpotBrokerSubAccountDepositHistApi {
	api.req.Offset = GetPointer(offset)
	return api
}

// NO
func (api *SpotBrokerSubAccountDepositHistApi) RecvWindow(recvWindow int64) *SpotBrokerSubAccountDepositHistApi {
	api.req.RecvWindow = GetPointer(recvWindow)
	return api
}

// YES
func (api *SpotBrokerSubAccountDepositHistApi) Timestamp(timestamp int64) *SpotBrokerSubAccountDepositHistApi {
	api.req.Timestamp = GetPointer(timestamp)
	return api
}

type SpotBrokerSubAccountDepositHistApi struct {
	client *SpotRestClient
	req    *SpotBrokerSubAccountDepositHistReq
}

type SpotBrokerUniversalTransferPostReq struct {
	FromId          *string          `json:"fromId,omitempty"`          // NO
	ToId            *string          `json:"toId,omitempty"`            // NO
	FromAccountType *string          `json:"fromAccountType,omitempty"` // YES SPOT,USDT_FUTURE,COIN_FUTURE
	ToAccountType   *string          `json:"toAccountType,omitempty"`   // YES SPOT,USDT_FUTURE,COIN_FUTURE
	ClientTranId    *string          `json:"clientTranId,omitempty"`    // NO Client transfer id, must be unique. The max length is 32 characters
	Asset           *string          `json:"asset,omitempty"`           // YES
	Amount          *decimal.Decimal `json:"amount,omitempty"`          // YES
	RecvWindow      *int64           `json:"recvWindow,omitempty"`      // NO
	Timestamp       *int64           `json:"timestamp,omitempty"`       // YES
}

// NO
func (api *SpotBrokerUniversalTransferPostApi) FromId(fromId string) *SpotBrokerUniversalTransferPostApi {
	api.req.FromId = GetPointer(fromId)
	return api
}

// NO
func (api *SpotBrokerUniversalTransferPostApi) ToId(toId string) *SpotBrokerUniversalTransferPostApi {
	api.req.ToId = GetPointer(toId)
	return api
}

// YES SPOT,USDT_FUTURE,COIN_FUTURE
func (api *SpotBrokerUniversalTransferPostApi) FromAccountType(fromAccountType string) *SpotBrokerUniversalTransferPostApi {
	api.req.FromAccountType = GetPointer(fromAccountType)
	return api
}

// YES SPOT,USDT_FUTURE,COIN_FUTURE
func (api *SpotBrokerUniversalTransferPostApi) ToAccountType(toAccountType string) *SpotBrokerUniversalTransferPostApi {
	api.req.ToAccountType = GetPointer(toAccountType)
	return api
}

// NO Client transfer id, must be unique. The max length is 32 characters
func (api *SpotBrokerUniversalTransferPostApi) ClientTranId(clientTranId string) *SpotBrokerUniversalTransferPostApi {
	api.req.ClientTranId = GetPointer(clientTranId)
	return api
}

// YES
func (api *SpotBrokerUniversalTransferPostApi) Asset(asset string) *SpotBrokerUniversalTransferPostApi {
	api.req.Asset = GetPointer(asset)
	return api
}

// YES
func (api *SpotBrokerUniversalTransferPostApi) Amount(amount decimal.Decimal) *SpotBrokerUniversalTransferPostApi {
	api.req.Amount = GetPointer(amount)
	return api
}

// NO
func (api *SpotBrokerUniversalTransferPostApi) RecvWindow(recvWindow int64) *SpotBrokerUniversalTransferPostApi {
	api.req.RecvWindow = GetPointer(recvWindow)
	return api
}

// YES
func (api *SpotBrokerUniversalTransferPostApi) Timestamp(timestamp int64) *SpotBrokerUniversalTransferPostApi {
	api.req.Timestamp = GetPointer(timestamp)
	return api
}

type SpotBrokerUniversalTransferPostApi struct {
	client *SpotRestClient
	req    *SpotBrokerUniversalTransferPostReq
}

type SpotBrokerUniversalTransferGetReq struct {
	FromId        *string `json:"fromId,omitempty"`        // NO
	ToId          *string `json:"toId,omitempty"`          // NO
	ClientTranId  *string `json:"clientTranId,omitempty"`  // NO client transfer id
	StartTime     *int64  `json:"startTime,omitempty"`     // NO
	EndTime       *int64  `json:"endTime,omitempty"`       // NO
	Page          *int    `json:"page,omitempty"`          // NO default 1
	Limit         *int    `json:"limit,omitempty"`         // NO default 500, max 500
	ShowAllStatus *bool   `json:"showAllStatus,omitempty"` // NO TRUE or FALSE
	RecvWindow    *int64  `json:"recvWindow,omitempty"`    // NO
	Timestamp     *int64  `json:"timestamp,omitempty"`     // YES
}

// NO
func (api *SpotBrokerUniversalTransferGetApi) FromId(fromId string) *SpotBrokerUniversalTransferGetApi {
	api.req.FromId = GetPointer(fromId)
	return api
}

// NO
func (api *SpotBrokerUniversalTransferGetApi) ToId(toId string) *SpotBrokerUniversalTransferGetApi {
	api.req.ToId = GetPointer(toId)
	return api
}

// NO client transfer id
func (api *SpotBrokerUniversalTransferGetApi) ClientTranId(clientTranId string) *SpotBrokerUniversalTransferGetApi {
	api.req.ClientTranId = GetPointer(clientTranId)
	return api
}

// NO
func (api *SpotBrokerUniversalTransferGetApi) StartTime(startTime int64) *SpotBrokerUniversalTransferGetApi {
	api.req.StartTime = GetPointer(startTime)
	return api

}

// NO
func (api *SpotBrokerUniversalTransferGetApi) EndTime(endTime int64) *SpotBrokerUniversalTransferGetApi {
	api.req.EndTime = GetPointer(endTime)
	return api
}

// NO default 1
func (api *SpotBrokerUniversalTransferGetApi) Page(page int) *SpotBrokerUniversalTransferGetApi {
	api.req.Page = GetPointer(page)
	return api
}

// NO default 500, max 500
func (api *SpotBrokerUniversalTransferGetApi) Limit(limit int) *SpotBrokerUniversalTransferGetApi {
	api.req.Limit = GetPointer(limit)
	return api
}

// NO TRUE or FALSE
func (api *SpotBrokerUniversalTransferGetApi) ShowAllStatus(showAllStatus bool) *SpotBrokerUniversalTransferGetApi {
	api.req.ShowAllStatus = GetPointer(showAllStatus)
	return api
}

// NO
func (api *SpotBrokerUniversalTransferGetApi) RecvWindow(recvWindow int64) *SpotBrokerUniversalTransferGetApi {
	api.req.RecvWindow = GetPointer(recvWindow)
	return api
}

// YES
func (api *SpotBrokerUniversalTransferGetApi) Timestamp(timestamp int64) *SpotBrokerUniversalTransferGetApi {
	api.req.Timestamp = GetPointer(timestamp)
	return api
}

type SpotBrokerUniversalTransferGetApi struct {
	client *SpotRestClient
	req    *SpotBrokerUniversalTransferGetReq
}
