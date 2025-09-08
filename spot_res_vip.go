package mybinanceapi

//	{
//		"status": "OK",
//		"type": "GENERAL",
//		"code": "000000000",
//		"data": [
//		{
//		"endpoint": "WS-FAPI-MM"
//		"totalConnectionQuota": 600,
//		"connectionCount": 200,
//		"ipLimit": 350,
//		"addIpCount": 200,
//		"ipConnections": {
//		"192.168.1.0": 0, // 0 means added ok.
//		"192.168.1.1": -1, // -1 means In-progress for update
//		"192.168.1.2": -2, // -2 means Failed to update
//		"192.168.1.3": 10
//		}
//		},
//		{
//		"endpoint": "FAPI-MM"
//		"ipLimit": 350,
//		"addIpCount": 200,
//		"ipConnections": {
//		"192.168.1.0": 0, // 0 means added ok.
//		"192.168.1.1": -1, // -1 means In-progress for update
//		"192.168.1.2": -2, // -2 means Failed to update
//		"192.168.1.3": 10
//		}
//	   {
//		"endpoint": "DAPI-MM"
//		"ipLimit": 350,
//		"addIpCount": 200,
//		"ipConnections": {
//		"192.168.1.0": 0, // 0 means added ok.
//		"192.168.1.1": -1, // -1 means In-progress for update
//		"192.168.1.2": -2, // -2 means Failed to update
//		"192.168.1.3": 10
//		}
//		},
//		{
//		"endpoint": "WS-DAPI-MM"
//		"totalConnectionQuota": 600,
//		"connectionCount": 200,
//		"ipLimit": 350,
//		"addIpCount": 200,
//		"ipConnections": {
//		"192.168.1.0": 0, // 0 means added ok.
//		"192.168.1.1": -1, // -1 means In-progress for update
//		"192.168.1.2": -2, // -2 means Failed to update
//		"192.168.1.3": 10
//		}
//		},
//		{
//		"endpoint": "FSTREAM-MM"
//		"totalConnectionQuota": 600,
//		"connectionCount": 200,
//		"ipLimit": 350,
//		"addIpCount": 200,
//		"ipConnections": {
//		"192.168.1.0": 0, // 0 means added ok.
//		"192.168.1.1": -1, // -1 means In-progress for update
//		"192.168.1.2": -2, // -2 means Failed to update
//		"192.168.1.3": 10
//		}
//		}
//		{
//		"endpoint": "DSTREAM-MM"
//		"totalConnectionQuota": 600,
//		"connectionCount": 200,
//		"ipLimit": 350,
//		"addIpCount": 200,
//		"ipConnections": {
//		"192.168.1.0": 0, // 0 means added ok.
//		"192.168.1.1": -1, // -1 means In-progress for update
//		"192.168.1.2": -2, // -2 means Failed to update
//		"192.168.1.3": 10
//		}
//		}
//		]
//	   }
type SpotApiVipVipPortalFutureIpWhitelistsGetRes struct {
	Status string                                            `json:"status"`
	Type   string                                            `json:"type"`
	Code   string                                            `json:"code"`
	Data   []SpotApiVipVipPortalFutureIpWhitelistsGetResData `json:"data"`
}
type SpotApiVipVipPortalFutureIpWhitelistsGetResData struct {
	Endpoint             string         `json:"endpoint"`
	TotalConnectionQuota int            `json:"totalConnectionQuota"`
	ConnectionCount      int            `json:"connectionCount"`
	IpLimit              int            `json:"ipLimit"`
	AddIpCount           int            `json:"addIpCount"`
	IpConnections        map[string]int `json:"ipConnections"`
}

//	{
//		"status": "OK",
//		"type": "GENERAL",
//		"code": "000000000"
//	}
type SpotApiVipVipPortalFutureIpWhitelistsPostRes struct {
	Status string `json:"status"`
	Type   string `json:"type"`
	Code   string `json:"code"`
}

type SpotApiVipVipPortalFutureIpWhitelistsDeleteRes SpotApiVipVipPortalFutureIpWhitelistsPostRes
