package mybinanceapi

import (
	"time"
)

type WsApiReqExtend struct {
	ApiKey    string `json:"apiKey"`
	Signature string `json:"signature"`
}

func signWsApi[T any](apiKey, apiSecret string, paramsReq *T) string {
	firstString := "apiKey=" + apiKey
	query := binanceHandlerWsApiReq(paramsReq)
	firstString += "&" + query
	sign := HmacSha256(apiSecret, firstString)
	return sign
}

func (ws *SpotWsStreamClient) CreateOrder(api *SpotOrderPostApi) (*WsApiResult[SpotOrderPostRes], error) {
	api.Timestamp(time.Now().UnixMilli())

	type WsSpotOrderPostReq struct {
		WsApiReqExtend
		SpotOrderPostReq
	}
	req := WsSpotOrderPostReq{
		WsApiReqExtend: WsApiReqExtend{
			ApiKey:    ws.apiKey,
			Signature: signWsApi(ws.apiKey, ws.apiSecret, api.req),
		},
		SpotOrderPostReq: *api.req,
	}
	return sendApiMsg[WsSpotOrderPostReq, SpotOrderPostRes](&ws.WsStreamClient, 0, "order.place", req)
}
func (ws *SpotWsStreamClient) CancelOrder(api *SpotOrderDeleteApi) (*WsApiResult[SpotOrderDeleteRes], error) {
	api.Timestamp(time.Now().UnixMilli())

	type WsSpotOrderDeleteReq struct {
		WsApiReqExtend
		SpotOrderDeleteReq
	}
	req := WsSpotOrderDeleteReq{
		WsApiReqExtend: WsApiReqExtend{
			ApiKey:    ws.apiKey,
			Signature: signWsApi(ws.apiKey, ws.apiSecret, api.req),
		},
		SpotOrderDeleteReq: *api.req,
	}
	return sendApiMsg[WsSpotOrderDeleteReq, SpotOrderDeleteRes](&ws.WsStreamClient, 0, "order.cancel", req)
}
func (ws *SpotWsStreamClient) CancelReplaceOrder(api *SpotOrderCancelReplaceApi) (*WsApiResult[SpotOrderCancelReplaceRes], error) {
	api.Timestamp(time.Now().UnixMilli())

	type WsSpotOrderCancelReplaceReq struct {
		WsApiReqExtend
		SpotOrderCancelReplaceReq
	}
	req := WsSpotOrderCancelReplaceReq{
		WsApiReqExtend: WsApiReqExtend{
			ApiKey:    ws.apiKey,
			Signature: signWsApi(ws.apiKey, ws.apiSecret, api.req),
		},
		SpotOrderCancelReplaceReq: *api.req,
	}
	return sendApiMsg[WsSpotOrderCancelReplaceReq, SpotOrderCancelReplaceRes](&ws.WsStreamClient, 0, "order.cancelReplace", req)
}
func (ws *SpotWsStreamClient) QueryOrder(api *SpotOrderGetApi) (*WsApiResult[SpotOrderGetRes], error) {
	api.Timestamp(time.Now().UnixMilli())

	type WsSpotOrderGetReq struct {
		WsApiReqExtend
		SpotOrderGetReq
	}
	req := WsSpotOrderGetReq{
		WsApiReqExtend: WsApiReqExtend{
			ApiKey:    ws.apiKey,
			Signature: signWsApi(ws.apiKey, ws.apiSecret, api.req),
		},
		SpotOrderGetReq: *api.req,
	}
	return sendApiMsg[WsSpotOrderGetReq, SpotOrderGetRes](&ws.WsStreamClient, 0, "order.status", req)
}

func (ws *FutureWsStreamClient) CreateOrder(api *FutureOrderPostApi) (*WsApiResult[FutureOrderPostRes], error) {
	api.Timestamp(time.Now().UnixMilli())

	type WsFutureOrderPostReq struct {
		WsApiReqExtend
		FutureOrderPostReq
	}
	req := WsFutureOrderPostReq{
		WsApiReqExtend: WsApiReqExtend{
			ApiKey:    ws.apiKey,
			Signature: signWsApi(ws.apiKey, ws.apiSecret, api.req),
		},
		FutureOrderPostReq: *api.req,
	}
	return sendApiMsg[WsFutureOrderPostReq, FutureOrderPostRes](&ws.WsStreamClient, 0, "order.place", req)
}
func (ws *FutureWsStreamClient) CancelOrder(api *FutureOrderDeleteApi) (*WsApiResult[FutureOrderDeleteRes], error) {
	api.Timestamp(time.Now().UnixMilli())

	type WsFutureOrderDeleteReq struct {
		WsApiReqExtend
		FutureOrderDeleteReq
	}
	req := WsFutureOrderDeleteReq{
		WsApiReqExtend: WsApiReqExtend{
			ApiKey:    ws.apiKey,
			Signature: signWsApi(ws.apiKey, ws.apiSecret, api.req),
		},
		FutureOrderDeleteReq: *api.req,
	}
	return sendApiMsg[WsFutureOrderDeleteReq, FutureOrderDeleteRes](&ws.WsStreamClient, 0, "order.cancel", req)
}
func (ws *FutureWsStreamClient) AmendOrder(api *FutureOrderPutApi) (*WsApiResult[FutureOrderPutRes], error) {
	api.Timestamp(time.Now().UnixMilli())

	type WsFutureOrderPutReq struct {
		WsApiReqExtend
		FutureOrderPutReq
	}
	req := WsFutureOrderPutReq{
		WsApiReqExtend: WsApiReqExtend{
			ApiKey:    ws.apiKey,
			Signature: signWsApi(ws.apiKey, ws.apiSecret, api.req),
		},
		FutureOrderPutReq: *api.req,
	}
	return sendApiMsg[WsFutureOrderPutReq, FutureOrderPutRes](&ws.WsStreamClient, 0, "order.modify", req)
}
func (ws *FutureWsStreamClient) QueryOrder(api *FutureOrderGetApi) (*WsApiResult[FutureOrderGetRes], error) {
	api.Timestamp(time.Now().UnixMilli())

	type WsFutureOrderGetReq struct {
		WsApiReqExtend
		FutureOrderGetReq
	}
	req := WsFutureOrderGetReq{
		WsApiReqExtend: WsApiReqExtend{
			ApiKey:    ws.apiKey,
			Signature: signWsApi(ws.apiKey, ws.apiSecret, api.req),
		},
		FutureOrderGetReq: *api.req,
	}
	return sendApiMsg[WsFutureOrderGetReq, FutureOrderGetRes](&ws.WsStreamClient, 0, "order.status", req)
}
