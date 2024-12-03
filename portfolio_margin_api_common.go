package mybinanceapi

// GET Portfolio margin ping REST 测试服务器连通性 (NONE)
func (client *PortfolioMarginRestClient) NewPing() *PortfolioMarginPingApi {
	return &PortfolioMarginPingApi{
		client: client,
		req:    &PortfolioMarginPingReq{},
	}
}

func (api *PortfolioMarginPingApi) Do() (*PortfolioMarginPingRes, error) {
	url := binanceHandlerRequestApi(PORTFOLIO_MARGIN, api.req, PortfolioMarginApiMap[PortfolioMarginPing])
	return binanceCallApiWithSecret[PortfolioMarginPingRes](api.client.c, url, GET)
}
