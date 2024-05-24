package mybinanceapi

type SpotUserDataStreamPostRes struct {
	ListenKey string `json:"listenKey"`
}

type SpotUserDataStreamPutRes struct{}

type SpotUserDataStreamDeleteRes struct{}

type SpotMarginUserDataStreamPostRes struct {
	ListenKey string `json:"listenKey"`
}

type SpotMarginUserDataStreamPutRes struct{}

type SpotMarginUserDataStreamDeleteRes struct{}

type SpotMarginIsolatedUserDataStreamPostRes struct {
	ListenKey string `json:"listenKey"`
}

type SpotMarginIsolatedUserDataStreamPutRes struct{}

type SpotMarginIsolatedUserDataStreamDeleteRes struct{}
