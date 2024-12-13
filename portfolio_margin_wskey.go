package mybinanceapi

// Ws账户推送相关
// binance PortfolioMargin PortfolioMarginListenKeyPost rest创建listenKey (USER_STREAM)
func (client *PortfolioMarginRestClient) NewPortfolioMarginListenKeyPost() *PortfolioMarginListenKeyPostApi {
	return &PortfolioMarginListenKeyPostApi{
		client: client,
		req:    &PortfolioMarginListenKeyPostReq{},
	}
}

func (api *PortfolioMarginListenKeyPostApi) Do() (*PortfolioMarginListenKeyPostRes, error) {
	url := binanceHandlerRequestApiWithSecret(PORTFOLIO_MARGIN, api.req, PortfolioMarginApiMap[PortfolioMarginListenKeyPost], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PortfolioMarginListenKeyPostRes](api.client.c, url, POST)
}

// binance PortfolioMargin PortfolioMarginListenKeyPut rest延长listenKey有效期 (USER_STREAM)
func (client *PortfolioMarginRestClient) NewPortfolioMarginListenKeyPut() *PortfolioMarginListenKeyPutApi {
	return &PortfolioMarginListenKeyPutApi{
		client: client,
		req:    &PortfolioMarginListenKeyPutReq{},
	}
}

func (api *PortfolioMarginListenKeyPutApi) Do() (*PortfolioMarginListenKeyPutRes, error) {
	url := binanceHandlerRequestApiWithSecret(PORTFOLIO_MARGIN, api.req, PortfolioMarginApiMap[PortfolioMarginListenKeyPut], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PortfolioMarginListenKeyPutRes](api.client.c, url, PUT)
}

// binance PortfolioMargin PortfolioMarginListenKeyDel rest关闭listenKey (USER_STREAM)
func (client *PortfolioMarginRestClient) NewPortfolioMarginListenKeyDel() *PortfolioMarginListenKeyDelApi {
	return &PortfolioMarginListenKeyDelApi{
		client: client,
		req:    &PortfolioMarginListenKeyDelReq{},
	}
}

func (api *PortfolioMarginListenKeyDelApi) Do() (*PortfolioMarginListenKeyDelRes, error) {
	url := binanceHandlerRequestApiWithSecret(PORTFOLIO_MARGIN, api.req, PortfolioMarginApiMap[PortfolioMarginListenKeyDel], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PortfolioMarginListenKeyDelRes](api.client.c, url, DELETE)
}
