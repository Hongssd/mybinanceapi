package mybinanceapi

import (
	"fmt"
	"time"
)

func (ws *SpotWsStreamClient) SubscribeUserDataStream() (*WsApiResult[SpotWsApiUserDataStreamResult], error) {
	type Timestamp struct {
		Timestamp *int64 `json:"timestamp"`
	}

	type WsSpotSubscribeUserDataStreamReq struct {
		WsApiReqExtend
		Timestamp
	}

	type WsSpotMarginUserListenTokenPostReq struct {
		ListenToken *string `json:"listenToken"`
	}

	switch ws.spotWsType {
	case SPOT_WS_TYPE:
		// 现货：使用签名 + timestamp 订阅，不需要 listenToken 刷新
		nowTimestamp := time.Now().UnixMilli() + serverTimeDelta

		timestamp := Timestamp{
			Timestamp: &nowTimestamp,
		}

		req := WsSpotSubscribeUserDataStreamReq{
			WsApiReqExtend: WsApiReqExtend{
				ApiKey:    ws.apiKey,
				Signature: signWsApi(ws.apiKey, ws.apiSecret, &timestamp),
			},
			Timestamp: timestamp,
		}
		return sendApiMsg[WsSpotSubscribeUserDataStreamReq, SpotWsApiUserDataStreamResult](&ws.WsStreamClient, 0, "userDataStream.subscribe.signature", req)

	case SPOT_MARGIN_WS_TYPE:
		// 全仓杠杆：先通过 REST 创建 listenToken，再用 listenToken 订阅
		res, err := ws.client.NewSpotMarginUserListenTokenPost().Do()
		if err != nil {
			return nil, err
		}

		// 启动协程定期刷新 listenToken（仅首次订阅时创建）
		ws.startMarginListenTokenRefreshLoop(res, func() (*SpotMarginUserListenTokenPostRes, error) {
			return ws.client.NewSpotMarginUserListenTokenPost().Do()
		})

		req := WsSpotMarginUserListenTokenPostReq{
			ListenToken: &res.Token,
		}
		return sendApiMsg[WsSpotMarginUserListenTokenPostReq, SpotWsApiUserDataStreamResult](&ws.WsStreamClient, 0, "userDataStream.subscribe.listenToken", req)

	case SPOT_ISOLATED_MARGIN_WS_TYPE:
		// 逐仓杠杆：为指定 symbol 创建 listenToken
		res, err := ws.client.NewSpotMarginUserListenTokenPost().
			Symbol(ws.isolatedSymbol).
			IsIsolated(true).
			Do()
		if err != nil {
			return nil, err
		}

		// 启动协程定期刷新该逐仓交易对的 listenToken
		ws.startMarginListenTokenRefreshLoop(res, func() (*SpotMarginUserListenTokenPostRes, error) {
			return ws.client.NewSpotMarginUserListenTokenPost().
				Symbol(ws.isolatedSymbol).
				IsIsolated(true).
				Do()
		})

		req := WsSpotMarginUserListenTokenPostReq{
			ListenToken: &res.Token,
		}
		return sendApiMsg[WsSpotMarginUserListenTokenPostReq, SpotWsApiUserDataStreamResult](&ws.WsStreamClient, 0, "userDataStream.subscribe.listenToken", req)

	default:
		return nil, fmt.Errorf("spotWsType is not support")
	}
}

// startMarginListenTokenRefreshLoop 启动协程，按照 expirationTime 周期性刷新杠杆账户 listenToken。
// 注意：这里只是定期调用 REST 刷新/重建 listenToken，以避免有效期过期；
// 订阅本身仍由 SubscribeUserDataStream 在连接建立/重连时处理。
func (ws *SpotWsStreamClient) startMarginListenTokenRefreshLoop(
	initial *SpotMarginUserListenTokenPostRes,
	builder func() (*SpotMarginUserListenTokenPostRes, error),
) {
	// 仅在首次订阅时启动刷新协程，避免重复创建
	if ws.isUserDataSubscribe {
		return
	}
	ws.isUserDataSubscribe = true

	go func() {
		res := initial
		for {
			// 计算距离过期的时间，预留 1 分钟安全窗口
			now := time.Now().UnixMilli()
			delayMs := res.ExpirationTime - now - 60*1000
			if delayMs < 10*1000 {
				// 如果服务器时间与本地略有偏差或 validity 较短，至少等待 10 秒再刷新
				delayMs = 10 * 1000
			}

			time.Sleep(time.Duration(delayMs) * time.Millisecond)

			newRes, err := builder()
			if err != nil {
				log.Error("refresh margin listenToken error: ", err)
				// 出错时稍等一会儿再重试，避免紧急重试打爆接口
				time.Sleep(10 * time.Second)
				continue
			}

			log.Debugf("refresh margin listenToken success, oldToken=%s newToken=%s expirationTime=%d",
				res.Token, newRes.Token, newRes.ExpirationTime)
			res = newRes
		}
	}()
}
