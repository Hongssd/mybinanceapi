package mybinanceapi

import "time"

// 账户接口
// binance PortfolioMarginGetBalance rest查询账户余额 (NONE)
func (client *PortfolioMarginRestClient) NewGetBalance() *PortfolioMarginBalanceApi {
	return &PortfolioMarginBalanceApi{
		client: client,
		req:    &PortfolioMarginBalanceReq{},
	}
}
func (api *PortfolioMarginBalanceApi) Do() (*PortfolioMarginBalanceRes, error) {
	api.Timestamp(time.Now().UnixMilli())
	url := binanceHandlerRequestApiWithSecret(PORTFOLIO_MARGIN, api.req, PortfolioMarginApi[PortfolioMarginGetBalance], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PortfolioMarginBalanceRes](api.client.c, url, GET)
}

// binance PortfolioMarginGetAccount rest查询账户信息 (NONE)
func (client *PortfolioMarginRestClient) NewGetAccount() *PortfolioMarginAccountApi {
	return &PortfolioMarginAccountApi{
		client: client,
		req:    &PortfolioMarginAccountReq{},
	}
}
func (api *PortfolioMarginAccountApi) Do() (*PortfolioMarginAccountRes, error) {
	api.Timestamp(time.Now().UnixMilli())
	url := binanceHandlerRequestApiWithSecret(PORTFOLIO_MARGIN, api.req, PortfolioMarginApi[PortfolioMarginGetAccount], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PortfolioMarginAccountRes](api.client.c, url, GET)
}

// binance PortfolioMarginGetMaxBorrowable rest查询最大可借贷额度
func (client *PortfolioMarginRestClient) NewGetMaxBorrowable() *PortfolioMarginMaxBorrowableApi {
	return &PortfolioMarginMaxBorrowableApi{
		client: client,
		req:    &PortfolioMarginMaxBorrowableReq{},
	}
}
func (api *PortfolioMarginMaxBorrowableApi) Do() (*PortfolioMarginMaxBorrowableRes, error) {
	api.Timestamp(time.Now().UnixMilli())
	url := binanceHandlerRequestApiWithSecret(PORTFOLIO_MARGIN, api.req, PortfolioMarginApi[PortfolioMarginGetMaxBorrowable], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PortfolioMarginMaxBorrowableRes](api.client.c, url, GET)
}

// binance PortfolioMarginGetMaxWithdraw rest查询最大可转出额度
func (client *PortfolioMarginRestClient) NewGetMaxWithdraw() *PortfolioMarginMaxWithdrawApi {
	return &PortfolioMarginMaxWithdrawApi{
		client: client,
		req:    &PortfolioMarginMaxWithdrawReq{},
	}
}
func (api *PortfolioMarginMaxWithdrawApi) Do() (*PortfolioMarginMaxWithdrawRes, error) {
	api.Timestamp(time.Now().UnixMilli())
	url := binanceHandlerRequestApiWithSecret(PORTFOLIO_MARGIN, api.req, PortfolioMarginApi[PortfolioMarginGetMaxWithdraw], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PortfolioMarginMaxWithdrawRes](api.client.c, url, GET)
}
