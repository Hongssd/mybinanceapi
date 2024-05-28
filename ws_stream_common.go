package mybinanceapi

import "sync"

type SpotWsStreamClient struct {
	WsStreamClient
	spotWsType     SpotWsType
	isolatedSymbol string
	client         *SpotRestClient
}
type FutureWsStreamClient struct {
	WsStreamClient
	client *FutureRestClient
}
type SwapWsStreamClient struct {
	WsStreamClient
	client *SwapRestClient
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
func (*MyBinance) NewFutureWsStreamClient() *FutureWsStreamClient {
	ws := &FutureWsStreamClient{
		WsStreamClient: WsStreamClient{
			apiType:       FUTURE,
			isGzip:        false,
			reSubscribeMu: &sync.Mutex{},
			wsStreamPath:  WS_STREAM_PATH,
		},
	}
	ws.initStructs()
	return ws
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
