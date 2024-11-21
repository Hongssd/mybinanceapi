package mybinanceapi

import "time"

// 账户接口
// binance PortfolioMarginBalance rest查询账户余额 (NONE)
func (client *PortfolioMarginRestClient) NewBalance() *PortfolioMarginBalanceApi {
	return &PortfolioMarginBalanceApi{
		client: client,
		req:    &PortfolioMarginBalanceReq{},
	}
}
func (api *PortfolioMarginBalanceApi) Do() (*PortfolioMarginBalanceRes, error) {
	api.Timestamp(time.Now().UnixMilli())
	url := binanceHandlerRequestApiWithSecret(PORTFOLIO_MARGIN, api.req, PortfolioMarginApi[PortfolioMarginBalance], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PortfolioMarginBalanceRes](api.client.c, url, GET)
}
