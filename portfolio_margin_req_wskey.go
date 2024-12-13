package mybinanceapi

type PortfolioMarginListenKeyPostReq struct{}

type PortfolioMarginListenKeyPostApi struct {
	client *PortfolioMarginRestClient
	req    *PortfolioMarginListenKeyPostReq
}

type PortfolioMarginListenKeyPutReq struct{}

type PortfolioMarginListenKeyPutApi struct {
	client *PortfolioMarginRestClient
	req    *PortfolioMarginListenKeyPutReq
}

type PortfolioMarginListenKeyDelReq struct{}

type PortfolioMarginListenKeyDelApi struct {
	client *PortfolioMarginRestClient
	req    *PortfolioMarginListenKeyDelReq
}
