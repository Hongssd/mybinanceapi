package mybinanceapi

type FutureApi int

const (
	//账户接口
	FutureAccount               FutureApi = iota //GET接口 账户信息V2 (USER_DATA)
	FuturePositionSideDualGet                    //GET接口 (HMAC SHA256)查询持仓模式(USER_DATA)
	FutureMultiAssetsMarginGet                   //GET接口 (HMAC SHA256)查询联合保证金模式(USER_DATA)
	FuturePositionSideDualPost                   //POST接口 (HMAC SHA256)更改持仓模式(TRADE)
	FutureMultiAssetsMarginPost                  //POST接口 (HMAC SHA256)更改联合保证金模式(TRADE)
	FutureLeverage                               //POST接口 (HMAC SHA256)调整开仓杠杆 (TRADE)
	FutureMarginType                             //POST接口 (HMAC SHA256)变换逐全仓模式 (TRADE)
	FutureLeverageBracket                        //GET接口 杠杆分层标准 (USER_DATA)
	FuturePositionRisk                           //GET接口 用户持仓风险V2 (USER_DATA)
	FutureIncomeAsyn                             //GET接口 获取合约资金流水下载Id (USER_DATA)
	FutureIncomeAsynId                           //GET接口 通过下载Id获取合约资金流水下载链接(USER_DATA)

	//交易接口
	FutureOpenOrders        //GET接口 (HMAC SHA256)查询当前挂单 (USER_DATA)
	FutureAllOrders         //GET接口 (HMAC SHA256)查询所有订单 (USER_DATA)
	FutureOrderPost         //POST接口 (HMAC SHA256)下单 (TRADE)
	FutureOrderPut          //PUT接口 (HMAC SHA256)修改订单 (TRADE)
	FutureOrderGet          //GET接口 (HMAC SHA256)查询订单 (USER_DATA)
	FutureOrderDelete       //DELETE接口 (HMAC SHA256)撤销订单 (TRADE)
	FutureBatchOrdersPost   //POST接口 (HMAC SHA256)批量下单 (TRADE)
	FutureBatchOrdersPut    //PUT接口 (HMAC SHA256)批量修改订单 (TRADE)
	FutureBatchOrdersDelete //DELETE接口 (HMAC SHA256)批量撤销订单 (TRADE)
	FutureUserTrades        //GET接口 (HMAC SHA256)账户成交历史 (USER_DATA)
	FutureCommissionRate    //GET接口 (HMAC SHA256)查询用户当前的手续费率

	//通用接口
	FuturePing         //GET接口 测试服务器连通性
	FutureServerTime   //GET接口 获取服务器时间
	FutureExchangeInfo //GET接口 交易规则和交易对信息

	//行情接口
	FutureKlines           //K线数据
	FutureDepth            //深度信息
	FutureTrades           //最新成交
	FutureHistoricalTrades //历史成交
	FutureAggTrades        //近期成交(归集)
	FuturePremiumIndex     //最新标记价格和资金费率
	FutureFundingRate      //查询资金费率历史
	FutureFundingInfo      //查询资金费率信息
	FutureTicker24hr       //24hr价格变动情况
	FutureTickerPrice      //最新价格
	FutureTickerBookTicker //当前最优挂单
	FutureDataBasis        //基差数据

	//Ws账户推送相关接口
	FutureListenKeyPost   //生成listenKey (USER_STREAM)
	FutureListenKeyPut    //延长listenKey有效期 (USER_STREAM)
	FutureListenKeyDelete //关闭listenKey (USER_STREAM)

)

var FutureApiMap = map[FutureApi]string{

	//账户接口
	FutureAccount:               "/fapi/v2/account",           //GET接口 账户信息V2 (USER_DATA)
	FuturePositionSideDualGet:   "/fapi/v1/positionSide/dual", //GET接口 (HMAC SHA256)查询持仓模式(USER_DATA)
	FutureMultiAssetsMarginGet:  "/fapi/v1/multiAssetsMargin", //GET接口 (HMAC SHA256)查询联合保证金模式(USER_DATA)
	FuturePositionSideDualPost:  "/fapi/v1/positionSide/dual", //POST接口 (HMAC SHA256)更改持仓模式(TRADE)
	FutureMultiAssetsMarginPost: "/fapi/v1/multiAssetsMargin", //POST接口 (HMAC SHA256)更改联合保证金模式(TRADE)
	FutureLeverage:              "/fapi/v1/leverage",          //POST接口 (HMAC SHA256)调整开仓杠杆 (TRADE)
	FutureMarginType:            "/fapi/v1/marginType",        //POST接口 (HMAC SHA256)变换逐全仓模式 (TRADE)
	FutureLeverageBracket:       "/fapi/v1/leverageBracket",   //GET接口 杠杆分层标准 (USER_DATA)
	FuturePositionRisk:          "/fapi/v2/positionRisk",      //GET接口 用户持仓风险V2 (USER_DATA)
	FutureIncomeAsyn:            "/fapi/v1/income/asyn",       //GET接口 获取合约资金流水下载Id (USER_DATA)
	FutureIncomeAsynId:          "/fapi/v1/income/asyn/id",    //GET接口 通过下载Id获取合约资金流水下载链接(USER_DATA)

	//交易接口
	FutureOpenOrders:        "/fapi/v1/openOrders",     //GET接口 (HMAC SHA256)查询当前挂单 (USER_DATA)
	FutureAllOrders:         "/fapi/v1/allOrders",      //GET接口 (HMAC SHA256)查询所有订单 (USER_DATA)
	FutureOrderPost:         "/fapi/v1/order",          //POST接口 (HMAC SHA256)下单 (TRADE)
	FutureOrderPut:          "/fapi/v1/order",          //PUT接口 (HMAC SHA256)修改订单 (TRADE)
	FutureOrderGet:          "/fapi/v1/order",          //GET接口 (HMAC SHA256)查询订单 (USER_DATA)
	FutureOrderDelete:       "/fapi/v1/order",          //DELETE接口 (HMAC SHA256)撤销订单 (TRADE)
	FutureBatchOrdersPost:   "/fapi/v1/batchOrders",    //POST接口 (HMAC SHA256)批量下单 (TRADE)
	FutureBatchOrdersPut:    "/fapi/v1/batchOrders",    //PUT接口 (HMAC SHA256)批量修改订单 (TRADE)
	FutureBatchOrdersDelete: "/fapi/v1/batchOrders",    //DELETE接口 (HMAC SHA256)批量撤销订单 (TRADE)
	FutureUserTrades:        "/fapi/v1/userTrades",     //GET接口 (HMAC SHA256)账户成交历史 (USER_DATA)
	FutureCommissionRate:    "/fapi/v1/commissionRate", //GET接口 (HMAC SHA256)查询用户当前的手续费率

	//通用接口
	FuturePing:         "/fapi/v1/ping",         //GET接口 测试服务器连通性
	FutureServerTime:   "/fapi/v1/time",         //GET接口 获取服务器时间
	FutureExchangeInfo: "/fapi/v1/exchangeInfo", //GET接口 交易规则和交易对信息

	//行情接口
	FutureKlines:           "/fapi/v1/klines",            //GET接口 K线数据
	FutureDepth:            "/fapi/v1/depth",             //GET接口 深度信息
	FutureTrades:           "/fapi/v1/trades",            //GET接口 最新成交
	FutureHistoricalTrades: "/fapi/v1/historicalTrades",  //GET接口 查询历史成交(MARKET_DATA)
	FutureAggTrades:        "/fapi/v1/aggTrades",         //GET接口 近期成交(归集)
	FuturePremiumIndex:     "/fapi/v1/premiumIndex",      //GET接口 最新标记价格和资金费率
	FutureFundingRate:      "/fapi/v1/fundingRate",       //GET接口 查询资金费率历史
	FutureFundingInfo:      "/fapi/v1/fundingInfo",       //GET接口 查询资金费率信息
	FutureTicker24hr:       "/fapi/v1/ticker/24hr",       //GET接口 24hr价格变动情况
	FutureTickerPrice:      "/fapi/v1/ticker/price",      //GET接口 最新价格
	FutureTickerBookTicker: "/fapi/v1/ticker/bookTicker", //GET接口 当前最优挂单
	FutureDataBasis:        "/futures/data/basis",        //GET接口 基差数据

	//Ws账户推送相关接口
	FutureListenKeyPost:   "/fapi/v1/listenKey", //POST接口 生成listenKey (USER_STREAM)
	FutureListenKeyPut:    "/fapi/v1/listenKey", //PUT接口 延长listenKey有效期 (USER_STREAM)
	FutureListenKeyDelete: "/fapi/v1/listenKey", //DELETE接口 关闭listenKey (USER_STREAM)

}
