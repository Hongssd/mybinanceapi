package mybinanceapi

type PortfolioMarginApi int

const (
	// 通用接口
	PortfolioMarginPing PortfolioMarginApi = iota // 测试服务器连通性
	// 账户信息接口
	PortfolioMarginGetBalance             // 查询账户余额
	PortfolioMarginGetAccount             // 查询账户信息
	PortfolioMarginGetMaxBorrowable       // 查询最大可借贷额度
	PortfolioMarginGetMaxWithdraw         // 查询最大可转出额度
	PortfolioMarginSetUmLeverage          // 设置UM开仓杠杆
	PortfolioMarginSetCmLeverage          // 设置CM开仓杠杆
	PortfolioMarginUmPositionSideDualPost // 更改UM持仓模式
	PortfolioMarginUmPositionSideDualGet  // 查询UM持仓模式
	PortfolioMarginCmPositionSideDualPost // 更改CM持仓模式
	PortfolioMarginCmPositionSideDualGet  // 查询CM持仓模式
	PortfolioMarginUmAccountV1            // 获取UM账户信息
	PortfolioMarginUmAccountV2            // 获取UM账户信息
	PortfolioMarginCmAccount              // 获取CM账户信息
	PortfolioMarginUmPositionRisk         // 查询UM持仓风险
	PortfolioMarginCmPositionRisk         // 查询CM持仓风险
	PortfolioMarginUmCommissionRate       // 查询UM手续费率
	PortfolioMarginCmCommissionRate       // 查询CM手续费率

	// 交易接口
	PortfolioMarginUmOrderPost               // UM下单
	PortfolioMarginUmOrderDelete             // UM撤单
	PortfolioMarginUmAllOpenOrdersDelete     // 撤销特定交易对当前全部UM挂单
	PortfolioMarginUmOrderPut                // UM修改订单
	PortfolioMarginUmOrderGet                // 查询UM订单
	PortfolioMarginUmOpenOrderGet            // 查询UM当前挂单
	PortfolioMarginUmOpenOrdersGet           // 查询UM当前全部挂单
	PortfolioMarginUmConditionalOrderPost    // UM条件下单
	PortfolioMarginUmConditionalOpenOrder    // 查询UM当前条件挂单
	PortfolioMarginUmConditionalOpenOrders   // 查询UM当前条件全部挂单
	PortfolioMarginUmConditionalOrderHistory // 查询UM条件订单历史

	PortfolioMarginCmOrderPost                      // CM下单
	PortfolioMarginCmOrderDelete                    // CM撤单
	PortfolioMarginCmAllOpenOrdersDelete            // 撤销特定交易对当前全部CM挂单
	PortfolioMarginCmOrderPut                       // CM修改订单
	PortfolioMarginCmOrderGet                       // 查询CM订单
	PortfolioMarginCmAllOrders                      // 查询CM所有订单
	PortfolioMarginCmOpenOrders                     // 查看当前全部CM挂单
	PortfolioMarginCmConditionalOrderPost           // CM条件下单
	PortfolioMarginCmConditionalOrderDelete         // 取消CM条件订单
	PortfolioMarginCmConditionalAllOpenOrdersDelete // 取消全部CM条件单
	PortfolioMarginCmConditionalOpenOrder           // 查询CM当前条件挂单

	PortfolioMarginMarginOrderPost    // 杠杆下单
	PortfolioMarginMarginOrderOcoPost // 杠杆OCO下单
)

var PortfolioMarginApiMap = map[PortfolioMarginApi]string{
	// 通用接口
	PortfolioMarginPing: "/papi/v1/ping", // 测试服务器连通性

	// 账户信息接口
	PortfolioMarginGetBalance:             "/papi/v1/balance",              // 查询账户余额
	PortfolioMarginGetAccount:             "/papi/v1/account",              // 查询账户信息
	PortfolioMarginGetMaxBorrowable:       "/papi/v1/margin/maxBorrowable", // 查询最大可借贷额度
	PortfolioMarginGetMaxWithdraw:         "/papi/v1/margin/maxWithdraw",   // 查询最大可转出额度
	PortfolioMarginSetUmLeverage:          "/papi/v1/um/leverage",          // 设置UM开仓杠杆
	PortfolioMarginSetCmLeverage:          "/papi/v1/cm/leverage",          // 设置CM开仓杠杆
	PortfolioMarginUmPositionSideDualPost: "/papi/v1/um/positionSide/dual", // 更改UM持仓模式
	PortfolioMarginUmPositionSideDualGet:  "/papi/v1/um/positionSide/dual", // 查询UM持仓模式
	PortfolioMarginCmPositionSideDualPost: "/papi/v1/cm/positionSide/dual", // 更改CM持仓模式
	PortfolioMarginCmPositionSideDualGet:  "/papi/v1/cm/positionSide/dual", // 查询CM持仓模式
	PortfolioMarginUmAccountV1:            "/papi/v1/um/account",           // 获取UM账户信息
	PortfolioMarginUmAccountV2:            "/papi/v2/um/account",           // 获取UM账户信息
	PortfolioMarginCmAccount:              "/papi/v1/cm/account",           // 获取CM账户信息
	PortfolioMarginUmPositionRisk:         "/papi/v1/um/positionRisk",      // 查询UM持仓风险
	PortfolioMarginCmPositionRisk:         "/papi/v1/cm/positionRisk",      // 查询CM持仓风险
	PortfolioMarginUmCommissionRate:       "/papi/v1/um/commissionRate",    // 查询UM手续费率
	PortfolioMarginCmCommissionRate:       "/papi/v1/cm/commissionRate",    // 查询CM手续费率

	//交易接口
	PortfolioMarginUmOrderPost:               "/papi/v1/um/order",                    // UM下单
	PortfolioMarginUmOrderDelete:             "/papi/v1/um/order",                    // UM撤单
	PortfolioMarginUmAllOpenOrdersDelete:     "/papi/v1/um/allOpenOrders",            // 撤销特定交易对当前全部UM挂单
	PortfolioMarginUmOrderPut:                "/papi/v1/um/order",                    // UM修改订单
	PortfolioMarginUmOrderGet:                "/papi/v1/um/order",                    // 查询UM订单
	PortfolioMarginUmOpenOrderGet:            "/papi/v1/um/openOrder",                // 查询UM当前挂单
	PortfolioMarginUmOpenOrdersGet:           "/papi/v1/um/openOrders",               // 查询UM当前全部挂单
	PortfolioMarginUmConditionalOrderPost:    "/papi/v1/um/conditional/order",        // UM条件下单
	PortfolioMarginUmConditionalOpenOrder:    "/papi/v1/um/conditional/openOrder",    // 查询UM当前条件挂单
	PortfolioMarginUmConditionalOpenOrders:   "/papi/v1/um/conditional/openOrders",   // 查询UM当前条件全部挂单
	PortfolioMarginUmConditionalOrderHistory: "/papi/v1/um/conditional/orderHistory", // 查询UM条件订单历史

	PortfolioMarginCmOrderPost:                      "/papi/v1/cm/order",                     // CM下单
	PortfolioMarginCmOrderDelete:                    "/papi/v1/cm/order",                     // CM撤单
	PortfolioMarginCmAllOpenOrdersDelete:            "/papi/v1/cm/allOpenOrders",             // 撤销特定交易对当前全部CM挂单
	PortfolioMarginCmOrderPut:                       "/papi/v1/cm/order",                     // CM修改订单
	PortfolioMarginCmOrderGet:                       "/papi/v1/cm/order",                     // 查询CM订单
	PortfolioMarginCmAllOrders:                      "/papi/v1/cm/allOrders",                 // 查询CM所有订单
	PortfolioMarginCmOpenOrders:                     "/papi/v1/cm/openOrders",                // 查看当前全部CM挂单（带symbol或不带symbol)
	PortfolioMarginCmConditionalOrderPost:           "/papi/v1/cm/conditional/order",         // CM条件下单
	PortfolioMarginCmConditionalOrderDelete:         "/papi/v1/cm/conditional/order",         // 取消CM条件订单
	PortfolioMarginCmConditionalAllOpenOrdersDelete: "/papi/v1/cm/conditional/allOpenOrders", // 取消全部CM条件单
	PortfolioMarginCmConditionalOpenOrder:           "/papi/v1/cm/conditional/openOrder",     // 查询CM当前条件挂单

	PortfolioMarginMarginOrderPost:    "/papi/v1/margin/order",     // 杠杆下单
	PortfolioMarginMarginOrderOcoPost: "/papi/v1/margin/order/oco", // 杠杆OCO下单
}
