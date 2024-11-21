package mybinanceapi

const (
	// 通用接口
	PortfolioMarginPing = "/papi/v1/ping" // 测试服务器连通性
	// 账户信息接口
	PortfolioMarginBalance = "/papi/v1/balance" // 查询账户余额
)

var PortfolioMarginApi = map[string]string{
	// 通用接口
	PortfolioMarginPing: "/papi/v1/ping", // 测试服务器连通性
	// 账户信息接口
	PortfolioMarginBalance: "/papi/v1/balance", // 查询账户余额

}
