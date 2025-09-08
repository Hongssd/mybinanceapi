package mybinanceapi

// Parameters:
// Name Type Mandatory Description
// RecvWindow LONG NO No more than 60000
// timestamp LONG YES
// endpoint STRING NO FAPI-MM/DAPI-MM/WS-FAPI-MM/WS-DAPI-MM/FSTREAM-MM/DSTREAM-MM
type SpotApiVipVipPortalFutureIpWhitelistsGetReq struct {
	Endpoint   *string `json:"endpoint"`   //NO FAPI-MM or FSTREAM-MM
	RecvWindow *int64  `json:"recvWindow"` //NO
	Timestamp  *int64  `json:"timestamp"`  //YES
}

type SpotApiVipVipPortalFutureIpWhitelistsGetApi struct {
	client *SpotRestClient
	req    *SpotApiVipVipPortalFutureIpWhitelistsGetReq
}

//FAPI-MM/DAPI-MM/WS-FAPI-MM/WS-DAPI-MM/FSTREAM-MM/DSTREAM-MM
func (api *SpotApiVipVipPortalFutureIpWhitelistsGetApi) Endpoint(Endpoint string) *SpotApiVipVipPortalFutureIpWhitelistsGetApi {
	api.req.Endpoint = GetPointer(Endpoint)
	return api
}

func (api *SpotApiVipVipPortalFutureIpWhitelistsGetApi) RecvWindow(RecvWindow int64) *SpotApiVipVipPortalFutureIpWhitelistsGetApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *SpotApiVipVipPortalFutureIpWhitelistsGetApi) Timestamp(Timestamp int64) *SpotApiVipVipPortalFutureIpWhitelistsGetApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

// Parameters:
// Name Type Mandatory Description
// RecvWindow LONG NO No more than 60000
// timestamp LONG YES
// endpoint STRING NO FAPI-MM/DAPI-MM/WS-FAPI-MM/WS-DAPI-MM/FSTREAM-MM/DSTREAM-MM
// ipAddresses STRING[] NO
type SpotApiVipVipPortalFutureIpWhitelistsPostReq struct {
	Endpoint    *string   `json:"endpoint"`    //NO FAPI-MM or FSTREAM-MM
	IpAddresses *[]string `json:"ipAddresses"` //NO
	RecvWindow  *int64    `json:"recvWindow"`  //NO
	Timestamp   *int64    `json:"timestamp"`   //YES
}

type SpotApiVipVipPortalFutureIpWhitelistsPostApi struct {
	client *SpotRestClient
	req    *SpotApiVipVipPortalFutureIpWhitelistsPostReq
}

//FAPI-MM/DAPI-MM/WS-FAPI-MM/WS-DAPI-MM/FSTREAM-MM/DSTREAM-MM
func (api *SpotApiVipVipPortalFutureIpWhitelistsPostApi) Endpoint(Endpoint string) *SpotApiVipVipPortalFutureIpWhitelistsPostApi {
	api.req.Endpoint = GetPointer(Endpoint)
	return api
}
func (api *SpotApiVipVipPortalFutureIpWhitelistsPostApi) IpAddresses(IpAddresses []string) *SpotApiVipVipPortalFutureIpWhitelistsPostApi {
	api.req.IpAddresses = GetPointer(IpAddresses)
	return api
}
func (api *SpotApiVipVipPortalFutureIpWhitelistsPostApi) RecvWindow(RecvWindow int64) *SpotApiVipVipPortalFutureIpWhitelistsPostApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *SpotApiVipVipPortalFutureIpWhitelistsPostApi) Timestamp(Timestamp int64) *SpotApiVipVipPortalFutureIpWhitelistsPostApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

// Parameters:
// Name Type Mandatory Description
// RecvWindow LONG NO No more than 60000
// timestamp LONG YES
// endpoint STRING NO FAPI-MM/DAPI-MM/WS-FAPI-MM/WS-DAPI-MM/FSTREAM-MM/DSTREAM-MM
// ipAddresses STRING[] NO
type SpotApiVipVipPortalFutureIpWhitelistsDeleteReq struct {
	Endpoint    *string   `json:"endpoint"`    //NO FAPI-MM or FSTREAM-MM
	IpAddresses *[]string `json:"ipAddresses"` //NO
	RecvWindow  *int64    `json:"recvWindow"`  //NO
	Timestamp   *int64    `json:"timestamp"`   //YES
}

type SpotApiVipVipPortalFutureIpWhitelistsDeleteApi struct {
	client *SpotRestClient
	req    *SpotApiVipVipPortalFutureIpWhitelistsDeleteReq
}

//FAPI-MM/DAPI-MM/WS-FAPI-MM/WS-DAPI-MM/FSTREAM-MM/DSTREAM-MM
func (api *SpotApiVipVipPortalFutureIpWhitelistsDeleteApi) Endpoint(Endpoint string) *SpotApiVipVipPortalFutureIpWhitelistsDeleteApi {
	api.req.Endpoint = GetPointer(Endpoint)
	return api
}

func (api *SpotApiVipVipPortalFutureIpWhitelistsDeleteApi) IpAddresses(IpAddresses []string) *SpotApiVipVipPortalFutureIpWhitelistsDeleteApi {
	api.req.IpAddresses = GetPointer(IpAddresses)
	return api
}
func (api *SpotApiVipVipPortalFutureIpWhitelistsDeleteApi) RecvWindow(RecvWindow int64) *SpotApiVipVipPortalFutureIpWhitelistsDeleteApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *SpotApiVipVipPortalFutureIpWhitelistsDeleteApi) Timestamp(Timestamp int64) *SpotApiVipVipPortalFutureIpWhitelistsDeleteApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}
