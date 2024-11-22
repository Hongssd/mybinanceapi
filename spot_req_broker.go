package mybinanceapi

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
