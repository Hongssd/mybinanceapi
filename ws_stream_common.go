package mybinanceapi

import "sync"

type SpotWsStreamClient struct {
	WsStreamClient
	spotWsType     SpotWsType
	isolatedSymbol string
	client         *SpotRestClient

	isUserDataSubscribe bool
}
type FutureWsStreamClient struct {
	WsStreamClient
	client *FutureRestClient
}
type SwapWsStreamClient struct {
	WsStreamClient
	client *SwapRestClient
}
type PMContractStreamClient struct {
	WsStreamClient
	client *PortfolioMarginRestClient
}
type PMMarginStreamClient struct {
	WsStreamClient
	client *PortfolioMarginRestClient
}

func (*MyBinance) NewSpotWsStreamClient() *SpotWsStreamClient {
	ws := &SpotWsStreamClient{
		WsStreamClient: WsStreamClient{
			apiType:       SPOT,
			isGzip:        false,
			reSubscribeMu: &sync.Mutex{},
			wsStreamPath:  WS_STREAM_PATH,
		},
	}
	ws.initStructs()
	return ws
}

// NewFuturePublicWsStreamClient 创建 USDT-M 行情客户端，连接 fstream /public/stream（盘口、bookTicker、depth 等）。
func (*MyBinance) NewFuturePublicWsStreamClient() *FutureWsStreamClient {
	ws := &FutureWsStreamClient{
		WsStreamClient: WsStreamClient{
			apiType:             FUTURE,
			isGzip:              false,
			reSubscribeMu:       &sync.Mutex{},
			wsStreamPath:        WS_STREAM_PATH,
			futureStreamTier:    FutureStreamTierPublic,
			privateStreamEvents: nil,
		},
	}
	ws.initStructs()
	return ws
}

// NewFutureMarketWsStreamClient 创建 USDT-M 行情客户端，连接 fstream /market/stream（K 线、ticker、aggTrade 等除 public 以外的市场数据）。
func (*MyBinance) NewFutureMarketWsStreamClient() *FutureWsStreamClient {
	ws := &FutureWsStreamClient{
		WsStreamClient: WsStreamClient{
			apiType:             FUTURE,
			isGzip:              false,
			reSubscribeMu:       &sync.Mutex{},
			wsStreamPath:        WS_STREAM_PATH,
			futureStreamTier:    FutureStreamTierMarket,
			privateStreamEvents: nil,
		},
	}
	ws.initStructs()
	return ws
}

// Deprecated: 请使用 NewFuturePublicWsStreamClient 订阅 depth、bookTicker 等，或使用 NewFutureMarketWsStreamClient 订阅其余市场流。本函数等价于 NewFutureMarketWsStreamClient。
func (*MyBinance) NewFutureWsStreamClient() *FutureWsStreamClient {
	ws := &FutureWsStreamClient{
		WsStreamClient: WsStreamClient{
			apiType:             FUTURE,
			isGzip:              false,
			reSubscribeMu:       &sync.Mutex{},
			wsStreamPath:        WS_STREAM_PATH,
			futureStreamTier:    FutureStreamTierMarket,
			privateStreamEvents: nil,
		},
	}
	ws.initStructs()
	return ws
}

// NewFutureUserWsStreamClient 创建 USDT-M 用户数据 WebSocket 客户端（/private/stream），等价于 NewFutureMarketWsStreamClient 占位后立刻 ConvertToAccountWs。
// privateEvents 非空时覆盖 DefaultFuturePrivateStreamEvents 写入 query。
func (*MyBinance) NewFutureUserWsStreamClient(apiKey, apiSecret string, privateEvents ...string) (*FutureWsStreamClient, error) {
	ws := &FutureWsStreamClient{
		WsStreamClient: WsStreamClient{
			apiType:             FUTURE,
			isGzip:              false,
			reSubscribeMu:       &sync.Mutex{},
			wsStreamPath:        WS_STREAM_PATH,
			futureStreamTier:    FutureStreamTierMarket,
			privateStreamEvents: nil,
		},
	}
	ws.initStructs()
	return ws.ConvertToAccountWs(apiKey, apiSecret, privateEvents...)
}
func (*MyBinance) NewSwapWsStreamClient() *SwapWsStreamClient {
	ws := &SwapWsStreamClient{
		WsStreamClient: WsStreamClient{
			apiType:       SWAP,
			isGzip:        false,
			reSubscribeMu: &sync.Mutex{},
			wsStreamPath:  WS_STREAM_PATH,
		},
	}
	ws.initStructs()
	return ws
}
func (*MyBinance) NewPMContractStreamClient() *PMContractStreamClient {
	ws := &PMContractStreamClient{
		WsStreamClient: WsStreamClient{
			apiType:       PORTFOLIO_MARGIN,
			isGzip:        false,
			reSubscribeMu: &sync.Mutex{},
			wsStreamPath:  WS_PM_STREAM_PATH,
		},
	}
	ws.initStructs()
	return ws
}
func (*MyBinance) NewPMMarginStreamClient() *PMMarginStreamClient {
	ws := &PMMarginStreamClient{
		WsStreamClient: WsStreamClient{
			apiType:       PORTFOLIO_MARGIN,
			isGzip:        false,
			reSubscribeMu: &sync.Mutex{},
			wsStreamPath:  WS_PM_STREAM_PATH,
		},
	}
	ws.initStructs()
	return ws
}

func (ws *SpotWsStreamClient) Close() error {
	if ws.isListenWs {
		err := ws.listenKeyDelete()
		if err != nil {
			return err
		}
		*ws.listenKeyRefreshStopChan <- struct{}{}
	}
	return ws.WsStreamClient.Close()
}
func (ws *FutureWsStreamClient) Close() error {
	if ws.isListenWs {
		err := ws.listenKeyDelete()
		if err != nil {
			return err
		}
		*ws.listenKeyRefreshStopChan <- struct{}{}
	}
	return ws.WsStreamClient.Close()
}
func (ws *SwapWsStreamClient) Close() error {
	if ws.isListenWs {
		err := ws.listenKeyDelete()
		if err != nil {
			return err
		}
		*ws.listenKeyRefreshStopChan <- struct{}{}
	}
	return ws.WsStreamClient.Close()
}
func (ws *PMMarginStreamClient) Close() error {
	if ws.isListenWs {
		err := ws.listenKeyDelete()
		if err != nil {
			return err
		}
		*ws.listenKeyRefreshStopChan <- struct{}{}
	}
	return ws.WsStreamClient.Close()
}
