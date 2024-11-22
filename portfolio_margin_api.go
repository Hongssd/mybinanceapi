package mybinanceapi

const (
	// 通用接口
	PortfolioMarginPing = "/papi/v1/ping" // 测试服务器连通性
	// 账户信息接口
	PortfolioMarginGetBalance            = "/papi/v1/balance"              // 查询账户余额
	PortfolioMarginGetAccount            = "/papi/v1/account"              // 查询账户信息
	PortfolioMarginGetMaxBorrowable      = "/papi/v1/margin/maxBorrowable" // 查询最大可借贷额度
	PortfolioMarginGetMaxWithdraw        = "/papi/v1/margin/maxWithdraw"   // 查询最大可转出额度
	PortfolioMarginSetUmLeverage         = "/papi/v1/um/leverage"          // 设置UM开仓杠杆
	PortfolioMarginSetCmLeverage         = "/papi/v1/cm/leverage"          // 设置CM开仓杠杆
	PortfolioMarginSetUmPositionSideDual = "/papi/v1/um/positionSide/dual" // 设置UM双向持仓模式
	PortfolioMarginGetCmPositionSideDual = "/papi/v1/cm/positionSide/dual" // 设置CM双向持仓模式
)

var PortfolioMarginApi = map[string]string{
	// 通用接口
	PortfolioMarginPing: "/papi/v1/ping", // 测试服务器连通性
	// 账户信息接口
	PortfolioMarginGetBalance:            "/papi/v1/balance",              // 查询账户余额
	PortfolioMarginGetAccount:            "/papi/v1/account",              // 查询账户信息
	PortfolioMarginGetMaxBorrowable:      "/papi/v1/margin/maxBorrowable", // 查询最大可借贷额度
	PortfolioMarginGetMaxWithdraw:        "/papi/v1/margin/maxWithdraw",   // 查询最大可转出额度
	PortfolioMarginSetUmLeverage:         "/papi/v1/um/leverage",          // 设置UM开仓杠杆
	PortfolioMarginSetCmLeverage:         "/papi/v1/cm/leverage",          // 设置CM开仓杠杆
	PortfolioMarginSetUmPositionSideDual: "/papi/v1/um/positionSide/dual", // 设置UM双向持仓模式
	PortfolioMarginGetCmPositionSideDual: "/papi/v1/cm/positionSide/dual", // 查询CM双向持仓模式
}
