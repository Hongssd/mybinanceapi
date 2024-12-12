package mybinanceapi

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"strings"
	"time"
)

type Payload[T any] struct {
	resultChan chan T        //接收订阅结果的通道
	errChan    chan error    //接收订阅错误的通道
	closeChan  chan struct{} //接收订阅关闭的通道
}

func (p *Payload[T]) ResultChan() chan T {
	return p.resultChan
}
func (p *Payload[T]) ErrChan() chan error {
	return p.errChan
}
func (p *Payload[T]) CloseChan() chan struct{} {
	return p.closeChan
}

type WsSpotPayload struct {
	Ws                             *SpotWsStreamClient
	Id                             int64
	OutboundAccountPositionPayload *Payload[WsSpotPayloadOutboundAccountPosition]
	BalanceUpdatePayload           *Payload[WsSpotPayloadBalanceUpdate]
	ExecutionReportPayload         *Payload[WsSpotPayloadExecutionReport]
}
type WsFuturePayload struct {
	Ws                      *FutureWsStreamClient
	Id                      int64
	AccountUpdatePayload    *Payload[WsFuturePayloadAccountUpdate]
	OrderTradeUpdatePayload *Payload[WsFuturePayloadOrderTradeUpdate]
}
type WsSwapPayload struct {
	Ws                      *SwapWsStreamClient
	Id                      int64
	AccountUpdatePayload    *Payload[WsSwapPayloadAccountUpdate]
	OrderTradeUpdatePayload *Payload[WsSwapPayloadOrderTradeUpdate]
}
type WsPMContractPayload struct {
	Ws                      *PMContractStreamClient
	Id                      int64
	AccountUpdatePayload    *Payload[WsPMContractPayloadAccountUpdate]
	OrderTradeUpdatePayload *Payload[WsPMContractPayloadOrderTradeUpdate]
}
type WsPMMarginPayload struct {
	Ws                             *PMMarginStreamClient
	Id                             int64
	OutboundAccountPositionPayload *Payload[WsPMMarginPayloadOutboundAccountPosition]
	BalanceUpdatePayload           *Payload[WsPMMarginPayloadBalanceUpdate]
	ExecutionReportPayload         *Payload[WsPMMarginPayloadExecutionReport]
}

func (payload *WsSpotPayload) Close() {
	if _, ok := payload.Ws.wsSpotPayloadMap.Load(payload.Id); ok {
		payload.Ws.wsSpotPayloadMap.Delete(payload.Id)
		payload.OutboundAccountPositionPayload.closeChan <- struct{}{}
		payload.BalanceUpdatePayload.closeChan <- struct{}{}
		payload.ExecutionReportPayload.closeChan <- struct{}{}
	}
}
func (payload *WsFuturePayload) Close() {
	if _, ok := payload.Ws.wsFuturePayloadMap.Load(payload.Id); ok {
		payload.Ws.wsFuturePayloadMap.Delete(payload.Id)
		payload.AccountUpdatePayload.closeChan <- struct{}{}
		payload.OrderTradeUpdatePayload.closeChan <- struct{}{}
	}
}
func (payload *WsSwapPayload) Close() {
	if _, ok := payload.Ws.wsSwapPayloadMap.Load(payload.Id); ok {
		payload.Ws.wsSwapPayloadMap.Delete(payload.Id)
		payload.AccountUpdatePayload.closeChan <- struct{}{}
		payload.OrderTradeUpdatePayload.closeChan <- struct{}{}
	}
}
func (payload *WsPMContractPayload) Close() {
	if _, ok := payload.Ws.wsPMContractPayloadMap.Load(payload.Id); ok {
		payload.Ws.wsPMContractPayloadMap.Delete(payload.Id)
		payload.AccountUpdatePayload.closeChan <- struct{}{}
		payload.OrderTradeUpdatePayload.closeChan <- struct{}{}
	}
}
func (payload *WsPMMarginPayload) Close() {
	if _, ok := payload.Ws.wsPMMarginPayloadMap.Load(payload.Id); ok {
		payload.Ws.wsPMMarginPayloadMap.Delete(payload.Id)
		payload.OutboundAccountPositionPayload.closeChan <- struct{}{}
		payload.BalanceUpdatePayload.closeChan <- struct{}{}
		payload.ExecutionReportPayload.closeChan <- struct{}{}
	}
}

func (ws *SpotWsStreamClient) CreatePayload() (*WsSpotPayload, error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return nil, err
	}
	id := node.Generate().Int64()
	payload := &WsSpotPayload{
		Ws: ws,
		Id: id,
		OutboundAccountPositionPayload: &Payload[WsSpotPayloadOutboundAccountPosition]{
			resultChan: make(chan WsSpotPayloadOutboundAccountPosition),
			errChan:    make(chan error),
			closeChan:  make(chan struct{}),
		},
		BalanceUpdatePayload: &Payload[WsSpotPayloadBalanceUpdate]{
			resultChan: make(chan WsSpotPayloadBalanceUpdate),
			errChan:    make(chan error),
			closeChan:  make(chan struct{}),
		},
		ExecutionReportPayload: &Payload[WsSpotPayloadExecutionReport]{
			resultChan: make(chan WsSpotPayloadExecutionReport),
			errChan:    make(chan error),
			closeChan:  make(chan struct{}),
		},
	}
	ws.wsSpotPayloadMap.Store(id, payload)
	return payload, nil
}
func (ws *FutureWsStreamClient) CreatePayload() (*WsFuturePayload, error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return nil, err
	}
	id := node.Generate().Int64()
	payload := &WsFuturePayload{
		Ws: ws,
		Id: id,
		AccountUpdatePayload: &Payload[WsFuturePayloadAccountUpdate]{
			resultChan: make(chan WsFuturePayloadAccountUpdate),
			errChan:    make(chan error),
			closeChan:  make(chan struct{}),
		},
		OrderTradeUpdatePayload: &Payload[WsFuturePayloadOrderTradeUpdate]{
			resultChan: make(chan WsFuturePayloadOrderTradeUpdate),
			errChan:    make(chan error),
			closeChan:  make(chan struct{}),
		},
	}
	ws.wsFuturePayloadMap.Store(id, payload)
	return payload, nil
}
func (ws *SwapWsStreamClient) CreatePayload() (*WsSwapPayload, error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return nil, err
	}
	id := node.Generate().Int64()
	payload := &WsSwapPayload{
		Ws: ws,
		Id: id,
		AccountUpdatePayload: &Payload[WsSwapPayloadAccountUpdate]{
			resultChan: make(chan WsSwapPayloadAccountUpdate),
			errChan:    make(chan error),
			closeChan:  make(chan struct{}),
		},
		OrderTradeUpdatePayload: &Payload[WsSwapPayloadOrderTradeUpdate]{
			resultChan: make(chan WsSwapPayloadOrderTradeUpdate),
			errChan:    make(chan error),
			closeChan:  make(chan struct{}),
		},
	}
	ws.wsSwapPayloadMap.Store(id, payload)
	return payload, nil
}

// PM
func (ws *PMContractStreamClient) CreatePayload() (*WsPMContractPayload, error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return nil, err
	}
	id := node.Generate().Int64()
	payload := &WsPMContractPayload{
		Ws: ws,
		Id: id,
		AccountUpdatePayload: &Payload[WsPMContractPayloadAccountUpdate]{
			resultChan: make(chan WsPMContractPayloadAccountUpdate),
			errChan:    make(chan error),
			closeChan:  make(chan struct{}),
		},
		OrderTradeUpdatePayload: &Payload[WsPMContractPayloadOrderTradeUpdate]{
			resultChan: make(chan WsPMContractPayloadOrderTradeUpdate),
			errChan:    make(chan error),
			closeChan:  make(chan struct{}),
		},
	}
	ws.wsPMContractPayloadMap.Store(id, payload)
	return payload, nil
}
func (ws *PMMarginStreamClient) CreatePayload() (*WsPMMarginPayload, error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return nil, err
	}
	id := node.Generate().Int64()
	payload := &WsPMMarginPayload{
		Ws: ws,
		Id: id,
		OutboundAccountPositionPayload: &Payload[WsPMMarginPayloadOutboundAccountPosition]{
			resultChan: make(chan WsPMMarginPayloadOutboundAccountPosition),
			errChan:    make(chan error),
			closeChan:  make(chan struct{}),
		},
		BalanceUpdatePayload: &Payload[WsPMMarginPayloadBalanceUpdate]{
			resultChan: make(chan WsPMMarginPayloadBalanceUpdate),
			errChan:    make(chan error),
			closeChan:  make(chan struct{}),
		},
		ExecutionReportPayload: &Payload[WsPMMarginPayloadExecutionReport]{
			resultChan: make(chan WsPMMarginPayloadExecutionReport),
			errChan:    make(chan error),
			closeChan:  make(chan struct{}),
		},
	}
	ws.wsPMMarginPayloadMap.Store(id, payload)
	return payload, nil
}

type SpotWsType int

const (
	SPOT_WS_TYPE SpotWsType = iota
	SPOT_MARGIN_WS_TYPE
	SPOT_ISOLATED_MARGIN_WS_TYPE
)

func (ws *SpotWsStreamClient) ConvertToAccountWs(apiKey string, apiSecret string, spotWsType SpotWsType, isolatedSymbol ...string) (*SpotWsStreamClient, error) {
	ws.wsStreamPath = WS_ACCOUNT_PATH
	ws.apiKey = apiKey
	ws.apiSecret = apiSecret
	ws.spotWsType = spotWsType
	ws.isListenWs = true
	if len(isolatedSymbol) > 0 {
		ws.isolatedSymbol = isolatedSymbol[0]
	}
	b := MyBinance{}
	ws.client = b.NewSpotRestClient(apiKey, apiSecret)

	err := ws.listenKeyPost()
	if err != nil {
		return nil, err
	}
	//创建一个协程定时刷新listenKey，如果已存在旧的刷新协程则不再创建
	if ws.listenKeyRefreshStopChan == nil {
		stopChan := make(chan struct{})
		ws.listenKeyRefreshStopChan = &stopChan
		go func() {
			for {
				select {
				case <-time.After(ListenKeyRefreshInterval):
					err := ws.listenKeyPut()
					for err != nil {
						log.Error(err)
						time.Sleep(5 * time.Second)
						if strings.Contains(err.Error(), "-1125") {
							//如果是-1125错误，则Post更新
							err = ws.listenKeyPost()
						} else {
							err = ws.listenKeyPut()
						}
					}
				case <-*ws.listenKeyRefreshStopChan:
					return
				}
			}
		}()
	}

	return ws, nil
}
func (ws *SpotWsStreamClient) listenKeyPost() error {
	//申请新的listenKey
	switch ws.spotWsType {
	case SPOT_WS_TYPE:
		res, err := ws.client.NewSpotUserDataStreamPost().Do()
		if err != nil {
			log.Error(err)
			return err
		}
		ws.listenKey = res.ListenKey
		log.Debug("listenKey:", ws.listenKey)
	case SPOT_MARGIN_WS_TYPE:
		res, err := ws.client.NewSpotMarginUserDataStreamPost().Do()
		if err != nil {
			log.Error(err)
			return err
		}
		ws.listenKey = res.ListenKey
		log.Debug("listenKey:", ws.listenKey)
	case SPOT_ISOLATED_MARGIN_WS_TYPE:
		res, err := ws.client.NewSpotMarginIsolatedUserDataStreamPost().Symbol(ws.isolatedSymbol).Do()
		if err != nil {
			log.Error(err)
			return err
		}
		ws.listenKey = res.ListenKey
		log.Debug("listenKey:", ws.listenKey)
	default:
		return fmt.Errorf("spotWsType is not support")
	}
	return nil
}
func (ws *SpotWsStreamClient) listenKeyPut() error {
	//申请新的listenKey
	switch ws.spotWsType {
	case SPOT_WS_TYPE:
		_, err := ws.client.NewSpotUserDataStreamPut().ListenKey(ws.listenKey).Do()
		if err != nil {
			log.Error(err)
			return err
		}
	case SPOT_MARGIN_WS_TYPE:
		_, err := ws.client.NewSpotMarginUserDataStreamPut().ListenKey(ws.listenKey).Do()
		if err != nil {
			log.Error(err)
			return err
		}

	case SPOT_ISOLATED_MARGIN_WS_TYPE:
		_, err := ws.client.NewSpotMarginIsolatedUserDataStreamPut().Symbol(ws.isolatedSymbol).ListenKey(ws.listenKey).Do()
		if err != nil {
			log.Error(err)
			return err
		}
	default:
		return fmt.Errorf("spotWsType is not support")
	}
	return nil
}
func (ws *SpotWsStreamClient) listenKeyDelete() error {
	switch ws.spotWsType {
	case SPOT_WS_TYPE:
		_, err := ws.client.NewSpotUserDataStreamDelete().ListenKey(ws.listenKey).Do()
		if err != nil {
			log.Error(err)
			return err
		}
	case SPOT_MARGIN_WS_TYPE:
		_, err := ws.client.NewSpotMarginUserDataStreamDelete().ListenKey(ws.listenKey).Do()
		if err != nil {
			log.Error(err)
			return err
		}

	case SPOT_ISOLATED_MARGIN_WS_TYPE:
		_, err := ws.client.NewSpotMarginIsolatedUserDataStreamDelete().Symbol(ws.isolatedSymbol).ListenKey(ws.listenKey).Do()
		if err != nil {
			log.Error(err)
			return err
		}
	default:
		return fmt.Errorf("spotWsType is not support")
	}
	return nil
}

func (ws *FutureWsStreamClient) ConvertToAccountWs(apiKey string, apiSecret string) (*FutureWsStreamClient, error) {
	ws.wsStreamPath = WS_ACCOUNT_PATH
	ws.apiKey = apiKey
	ws.apiSecret = apiSecret
	ws.isListenWs = true
	b := MyBinance{}
	ws.client = b.NewFutureRestClient(apiKey, apiSecret)

	err := ws.listenKeyPost()
	if err != nil {
		return nil, err
	}
	//创建一个协程定时刷新listenKey，如果已存在旧的刷新协程则不再创建
	if ws.listenKeyRefreshStopChan == nil {
		stopChan := make(chan struct{})
		ws.listenKeyRefreshStopChan = &stopChan
		go func() {
			for {
				select {
				case <-time.After(ListenKeyRefreshInterval):
					err := ws.listenKeyPut()
					for err != nil {
						log.Error(err)
						time.Sleep(5 * time.Second)
						if strings.Contains(err.Error(), "-1125") {
							//如果是-1125错误，则Post更新
							err = ws.listenKeyPost()
						} else {
							err = ws.listenKeyPut()
						}
					}
				case <-*ws.listenKeyRefreshStopChan:
					return
				}
			}
		}()
	}

	return ws, nil
}
func (ws *FutureWsStreamClient) listenKeyPost() error {
	res, err := ws.client.NewFutureListenKeyPost().Do()
	if err != nil {
		log.Error(err)
		return err
	}
	ws.listenKey = res.ListenKey
	log.Debug("listenKey:", ws.listenKey)
	return nil
}
func (ws *FutureWsStreamClient) listenKeyPut() error {
	_, err := ws.client.NewFutureListenKeyPut().Do()
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
func (ws *FutureWsStreamClient) listenKeyDelete() error {
	_, err := ws.client.NewFutureListenKeyDelete().Do()
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (ws *SwapWsStreamClient) ConvertToAccountWs(apiKey string, apiSecret string) (*SwapWsStreamClient, error) {
	ws.wsStreamPath = WS_ACCOUNT_PATH
	ws.apiKey = apiKey
	ws.apiSecret = apiSecret
	ws.isListenWs = true
	b := MyBinance{}
	ws.client = b.NewSwapRestClient(apiKey, apiSecret)

	err := ws.listenKeyPost()
	if err != nil {
		return nil, err
	}

	//创建一个协程定时刷新listenKey，如果已存在旧的刷新协程则不再创建
	if ws.listenKeyRefreshStopChan == nil {
		stopChan := make(chan struct{})
		ws.listenKeyRefreshStopChan = &stopChan
		go func() {
			for {
				select {
				case <-time.After(ListenKeyRefreshInterval):
					err := ws.listenKeyPut()
					for err != nil {
						log.Error(err)
						time.Sleep(5 * time.Second)
						if strings.Contains(err.Error(), "-1125") {
							//如果是-1125错误，则Post更新
							err = ws.listenKeyPost()
						} else {
							err = ws.listenKeyPut()
						}
					}
				case <-*ws.listenKeyRefreshStopChan:
					return
				}
			}
		}()
	}

	return ws, nil
}
func (ws *SwapWsStreamClient) listenKeyPost() error {
	res, err := ws.client.NewSwapListenKeyPost().Do()
	if err != nil {
		log.Error(err)
		return err
	}
	ws.listenKey = res.ListenKey
	log.Debug("listenKey:", ws.listenKey)
	return nil
}
func (ws *SwapWsStreamClient) listenKeyPut() error {
	_, err := ws.client.NewSwapListenKeyPut().Do()
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
func (ws *SwapWsStreamClient) listenKeyDelete() error {
	_, err := ws.client.NewSwapListenKeyDelete().Do()
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (ws *PMContractStreamClient) ConvertToAccountWs(apiKey string, apiSecret string) (*PMContractStreamClient, error) {
	ws.wsStreamPath = WS_PM_STREAM_PATH
	ws.apiKey = apiKey
	ws.apiSecret = apiSecret
	ws.isListenWs = true
	b := MyBinance{}
	ws.client = b.NewPortfolioMarginClient(apiKey, apiSecret)

	err := ws.listenKeyPost()
	if err != nil {
		return nil, err
	}

	//创建一个协程定时刷新listenKey，如果已存在旧的刷新协程则不再创建
	if ws.listenKeyRefreshStopChan == nil {
		stopChan := make(chan struct{})
		ws.listenKeyRefreshStopChan = &stopChan
		go func() {
			for {
				select {
				case <-time.After(ListenKeyRefreshInterval):
					err := ws.listenKeyPut()
					for err != nil {
						log.Error(err)
						time.Sleep(5 * time.Second)
						if strings.Contains(err.Error(), "-1125") {
							//如果是-1125错误，则Post更新
							err = ws.listenKeyPost()
						} else {
							err = ws.listenKeyPut()
						}
					}
				case <-*ws.listenKeyRefreshStopChan:
					return
				}
			}
		}()
	}

	return ws, nil
}
func (ws *PMContractStreamClient) listenKeyPost() error {
	res, err := ws.client.NewPortfolioMarginListenKeyPost().Do()
	if err != nil {
		log.Error(err)
		return err
	}
	ws.listenKey = res.ListenKey
	log.Debug("listenKey:", ws.listenKey)
	return nil
}
func (ws *PMContractStreamClient) listenKeyPut() error {
	_, err := ws.client.NewPortfolioMarginListenKeyPut().Do()
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
func (ws *PMContractStreamClient) listenKeyDelete() error {
	_, err := ws.client.NewPortfolioMarginListenKeyDel().Do()
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (ws *PMMarginStreamClient) ConvertToAccountWs(apiKey string, apiSecret string) (*PMMarginStreamClient, error) {
	ws.wsStreamPath = WS_PM_STREAM_PATH
	ws.apiKey = apiKey
	ws.apiSecret = apiSecret
	ws.isListenWs = true
	b := MyBinance{}
	ws.client = b.NewPortfolioMarginClient(apiKey, apiSecret)

	err := ws.listenKeyPost()
	if err != nil {
		return nil, err
	}

	// 创建一个协程定时刷新listenKey，如果已存在旧的刷新协程则不再创建
	if ws.listenKeyRefreshStopChan == nil {
		stopChan := make(chan struct{})
		ws.listenKeyRefreshStopChan = &stopChan
		go func() {
			for {
				select {
				case <-time.After(ListenKeyRefreshInterval):
					err := ws.listenKeyPut()
					for err != nil {
						log.Error(err)
						time.Sleep(5 * time.Second)
						if strings.Contains(err.Error(), "-1125") {
							// 如果是-1125错误，则Post更新
							err = ws.listenKeyPost()
						} else {
							err = ws.listenKeyPut()
						}
					}
				case <-*ws.listenKeyRefreshStopChan:
					return
				}
			}
		}()
	}

	return ws, nil
}
func (ws *PMMarginStreamClient) listenKeyPut() error {
	_, err := ws.client.NewPortfolioMarginListenKeyPut().Do()
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
func (ws *PMMarginStreamClient) listenKeyPost() error {
	res, err := ws.client.NewPortfolioMarginListenKeyPost().Do()
	if err != nil {
		log.Error(err)
		return err
	}
	ws.listenKey = res.ListenKey
	log.Debug("listenKey:", ws.listenKey)
	return nil
}
func (ws *PMMarginStreamClient) listenKeyDelete() error {
	_, err := ws.client.NewPortfolioMarginListenKeyDel().Do()
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
