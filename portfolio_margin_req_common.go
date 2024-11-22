package mybinanceapi

type PortfolioMarginPingReq struct{}
type PortfolioMarginPingApi struct {
	client *PortfolioMarginRestClient
	req    *PortfolioMarginPingReq
}
