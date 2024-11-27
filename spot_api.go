package mybinanceapi

type SpotApi int

const (
	//子母账户接口
	SpotSubAccountList                             SpotApi = iota //GET接口 查询子账户列表(适用主账户)
	SpotSubAccountUniversalTransferHistory                        //GET接口 查询子母账户万能划转历史查询
	SpotSubAccountAssets                                          //GET接口 查询子账户资产 (适用主账户)(USER_DATA)
	SpotSubAccountFuturesAccount                                  //GET接口 查询子账户Futures账户详情V2 (适用主账户)
	SpotSubAccountApiIpRestriction                                //GET接口 查询子账户API Key IP白名单 (适用母账户)
	SpotSubAccountTransferSubUserHistory                          //GET接口 查询子账户划转历史 (仅适用子账户)
	SpotManagedSubAccountQueryTransLog                            //GET接口 查询托管子账户的划转记录(适用交易团队子账户)(USER_DATA)
	SpotSubAccountVirtualSubAccount                               //POST接口 创建虚拟子账户(适用主账户)
	SpotSubAccountUniversalTransfer                               //POST接口 子母账户万能划转 (适用主账户)
	SpotSubAccountFuturesEnable                                   //POST接口 为子账户开通Futures (适用主账户)
	SpotSubAccountMarginEnable                                    //POST接口 为子账户开通杠杆 (适用主账户)
	SpotSubAccountSubAccountApiIpRestrictionPost                  //POST接口 为子账户API Key增加IP白名单(适用母账户)
	SpotSubAccountSubAccountApiIpRestrictionGet                   //GET接口 查询子账户API Key IP白名单(适用母账户)
	SpotSubAccountSubAccountApiIpRestrictionDelete                //DELETE接口 删除子账户API Key IP白名单(适用母账户)
	SpotSubAccountCapitalDepositSubAddress                        //GET接口 获取子账户充值地址 (适用主账户)
	SpotSubAccountCapitalDepositSubHisrec                         //GET接口 获取子账户充值历史 (适用主账户)
	SpotSubAccountFuturesPositionRisk                             //GET接口 查询子账户合约持仓信息 (适用主账户)
	SpotSubAccountSpotSummary                                     //GET接口 查询子账户现货汇总账户 (适用主账户)

	//杠杆账户接口
	SpotMarginAllPairs         //GET接口 获取所有全仓杠杆交易对(MARKET_DATA)
	SpotMarginIsolatedAllPairs //GET接口 获取所有逐仓杠杆交易对(USER_DATA)
	SpotMarginAccount          //GET接口 查询全仓杠杆账户详情 (USER_DATA)
	SpotMarginIsolatedAccount  //GET接口 查询逐仓杠杆账户详情 (USER_DATA)
	SpotMarginMaxBorrowable    //GET接口 查询账户最大可借贷额度(USER_DATA)
	SpotMarginMaxTransferable  //GET接口 查询最大可转出额 (USER_DATA)
	SpotMarginInterestHistory  //GET接口 获取利息历史 (USER_DATA)

	SpotMarginOrderGet    //GET接口 查询杠杆账户订单 (USER_DATA)
	SpotMarginOrderPost   //POST接口 杠杆账户下单 (TRADE)
	SpotMarginOrderDelete //DELETE接口 撤销订单 (TRADE)
	SpotMarginAllOrders   //GET接口 查询杠杆账户所有订单 (USER_DATA)
	SpotMarginOpenOrders  //GET接口 查询杠杆账户挂单记录 (USER_DATA)

	SpotMarginTransfer         //POST接口 全仓杠杆账户划转 (MARGIN)
	SpotMarginIsolatedTransfer //POST接口 逐仓杠杆账户划转 (MARGIN)
	SpotMarginLoan             //POST接口 杠杆账户借贷 (MARGIN) 支持逐仓和全仓
	SpotMarginRepay            //POST接口 杠杆账户归还借贷 (MARGIN) 支持逐仓和全仓

	SpotMarginMaxLeverage        //POST接口 调整全仓最大杠杆倍数（USER_DATA) 仅全仓
	SpotMarginTradeCoeff         //GET 获取用户个人杠杆账户信息汇总（USER_DATA）全仓
	SpotMarginCrossMarginData    //GET接口 获取全仓杠杆利率及限额 (USER_DATA)
	SpotMarginIsolatedMarginData //GET接口 获取逐仓杠杆利率及限额 (USER_DATA)
	SpotMarginIsolatedMarginTier // GET接口 获取逐仓杠杆档位信息 (USER_DATA)

	//现货账户接口
	SpotAccountApiTradingStatus //GET接口 账户API交易状态(USER_DATA)
	SpotAccount                 //GET接口 账户信息 (USER_DATA)
	SpotAssetGetFundingAsset    //POST接口 资金账户 (USER_DATA)
	SpotAssetTransferPost       //POST接口 用户万向划转 (USER_DATA)
	SpotAssetTransferGet        //GET接口 查询用户万向划转历史 (USER_DATA)
	SpotAssetTradeFee           //GET接口 查询用户交易手续费率 (USER_DATA)

	//现货订单接口
	SpotOpenOrders         //GET接口 查询当前挂单 (USER_DATA)
	SpotAllOrders          //GET接口 查询所有订单 (USER_DATA)
	SpotOrderGet           //GET接口 查询订单 (USER_DATA)
	SpotOrderPost          //POST接口 下单 (TRADE)
	SpotOrderDelete        //DELETE接口 撤销订单 (TRADE)
	SpotOrderCancelReplace //POST接口 撤消挂单再下单 (TRADE)
	SpotMyTrades           //GET接口 账户成交历史 (USER_DATA)

	//通用接口
	SpotPing         //GET接口 测试服务器连通性
	SpotServerTime   //GET接口 获取服务器时间
	SpotExchangeInfo //GET接口 获取交易规则和交易对信息。

	//行情接口
	SpotKlines           //GET接口 K线数据
	SpotTickerPrice      //GET接口 获取交易对最新价格
	SpotDepth            //GET接口 获取深度信息
	SpotTrades           //GET接口 近期成交列表
	SpotHistoricalTrades //GET接口 历史成交记录
	SpotAggTrades        //GET接口 近期成交(归集)
	SpotAvgPrice         //GET接口 当前平均价格
	SpotUiKlines         //GET接口 UIK线数据
	SpotTicker24hr       //GET接口 24hr 价格变动情况
	SpotTickerBookTicker //GET接口 当前最优挂单
	SpotTicker           //GET接口 滚动窗口价格变动统计

	//Ws账户推送相关
	SpotUserDataStreamPost   //POST接口    (现货账户)生成 Listen Key (USER_STREAM)
	SpotUserDataStreamPut    //PUT接口    (现货账户)延长 Listen Key 有效期 (USER_STREAM)
	SpotUserDataStreamDelete //DELETE接口 (现货账户)关闭 Listen Key (USER_STREAM)

	SpotMarginUserDataStreamPost   //POST接口   (杠杆账户)生成 Listen Key (USER_STREAM)
	SpotMarginUserDataStreamPut    //PUT接口    (杠杆账户)延长 Listen Key 有效期 (USER_STREAM)
	SpotMarginUserDataStreamDelete //DELETE接口 (杠杆账户)关闭 Listen Key (USER_STREAM)

	SpotMarginIsolatedUserDataStreamPost   //POST接口   (逐仓杠杆账户)生成 Listen Key (USER_STREAM)
	SpotMarginIsolatedUserDataStreamPut    //PUT接口    (逐仓杠杆账户)延长 Listen Key 有效期 (USER_STREAM)
	SpotMarginIsolatedUserDataStreamDelete //DELETE接口 (逐仓杠杆账户)关闭 Listen Key (USER_STREAM)

	//经纪商子母账号接口
	SpotBrokerSubAccountPost                        //POST接口 创建经纪商子账户
	SpotBrokerSubAccountGet                         //GET接口 查询经纪商子账户
	SpotBrokerSubAccountApiPost                     //POST 创建经纪商子账号ApiKey
	SpotBrokerSubAccountApiGet                      //GET 查询经纪商子账号ApiKey
	SpotBrokerSubAccountApiDelete                   //DELETE 删除经纪商子账号ApiKey
	SpotBrokerSubAccountFutures                     //POST接口 授权经纪商子账户Futures权限
	SpotBrokerSubAccountPermissionUniversalTransfer //POST接口 授权经纪商子账户万能划转权限
	SpotBrokerSubAccountApiIpRestrictionPost        //POST接口 更新经纪商子账户API Key IP白名单
	SpotBrokerSubAccountApiIpRestrictionGet         //GET接口 查询经纪商子账户API Key IP白名单
	SpotBrokerSubAccountApiIpRestrictionDelete      //DELETE接口 删除经纪商子账户API Key IP白名单
	SpotBrokerSubAccountDepositHist                 //GET接口 查询经纪商子账户充值历史
	SpotBrokerUniversalTransferPost                 //POST接口 经纪商子账户万能划转
	SpotBrokerUniversalTransferGet                  //GET接口 查询经纪商子账户万能划转历史
)

var SpotApiMap = map[SpotApi]string{
	//子母账户接口
	SpotSubAccountList:                             "/sapi/v1/sub-account/list",                        //GET接口 查询子账户列表(适用主账户)
	SpotSubAccountUniversalTransferHistory:         "/sapi/v1/sub-account/universalTransfer",           //GET接口 子母账户万能划转历史查询 (适用主账户)
	SpotSubAccountAssets:                           "/sapi/v4/sub-account/assets",                      //GET接口 查询子账户资产(适用主账户)(USER_DATA)
	SpotSubAccountFuturesAccount:                   "/sapi/v2/sub-account/futures/account",             //GET接口 查询子账户Futures账户详情V2 (适用主账户)
	SpotSubAccountApiIpRestriction:                 "/sapi/v1/sub-account/subAccountApi/ipRestriction", //GET接口 查询子账户API Key IP白名单 (适用母账户)
	SpotSubAccountTransferSubUserHistory:           "/sapi/v1/sub-account/transfer/subUserHistory",     //GET 查询子账户划转历史 (仅适用子账户)
	SpotManagedSubAccountQueryTransLog:             "/sapi/v1/managed-subaccount/query-trans-log",      //GET接口 查询托管子账户的划转记录(适用交易团队子账户)(USER_DATA)
	SpotSubAccountVirtualSubAccount:                "/sapi/v1/sub-account/virtualSubAccount",           //POST接口 创建虚拟子账户(适用主账户)
	SpotSubAccountUniversalTransfer:                "/sapi/v1/sub-account/universalTransfer",           //POST接口 子母账户万能划转 (适用主账户)
	SpotSubAccountFuturesEnable:                    "/sapi/v1/sub-account/futures/enable",              //POST接口 为子账户开通Futures (适用主账户)
	SpotSubAccountMarginEnable:                     "/sapi/v1/sub-account/margin/enable",               //POST接口 为子账户开通杠杆 (适用主账户)
	SpotSubAccountSubAccountApiIpRestrictionPost:   "/sapi/v2/sub-account/subAccountApi/ipRestriction", //POST接口 为子账户API Key增加IP白名单(适用母账户)
	SpotSubAccountSubAccountApiIpRestrictionGet:    "/sapi/v1/sub-account/subAccountApi/ipRestriction", //GET接口 查询子账户API Key IP白名单(适用母账户)
	SpotSubAccountSubAccountApiIpRestrictionDelete: "/sapi/v1/sub-account/subAccountApi/ipRestriction", //DELETE接口 删除子账户API Key IP白名单(适用母账户)
	SpotSubAccountCapitalDepositSubAddress:         "/sapi/v1/capital/deposit/subAddress",              //GET接口 获取子账户充值地址 (适用主账户)
	SpotSubAccountCapitalDepositSubHisrec:          "/sapi/v1/capital/deposit/subHisrec",               //GET接口 获取子账户充值历史 (适用主账户)
	SpotSubAccountFuturesPositionRisk:              "/sapi/v1/sub-account/futures/positionRisk",        //GET接口 查询子账户合约持仓信息 (适用主账户)
	SpotSubAccountSpotSummary:                      "/sapi/v1/sub-account/spotSummary",                 //GET接口 查询子账户现货汇总账户 (适用主账户)

	//杠杆账户接口
	SpotMarginAllPairs:         "/sapi/v1/margin/allPairs",          //GET 获取所有全仓杠杆交易对(MARKET_DATA)
	SpotMarginIsolatedAllPairs: "/sapi/v1/margin/isolated/allPairs", //GET 获取所有逐仓杠杆交易对(USER_DATA)
	SpotMarginAccount:          "/sapi/v1/margin/account",           //GET 查询全仓杠杆账户详情 (USER_DATA)
	SpotMarginIsolatedAccount:  "/sapi/v1/margin/isolated/account",  //GET 查询逐仓杠杆账户详情 (USER_DATA)
	SpotMarginMaxBorrowable:    "/sapi/v1/margin/maxBorrowable",     //GET 查询账户最大可借贷额度(USER_DATA)
	SpotMarginMaxTransferable:  "/sapi/v1/margin/maxTransferable",   //GET 查询最大可转出额 (USER_DATA)
	SpotMarginInterestHistory:  "/sapi/v1/margin/interestHistory",   //GET 获取利息历史 (USER_DATA)

	SpotMarginOrderGet:    "/sapi/v1/margin/order",      //GET 查询杠杆账户订单 (USER_DATA)
	SpotMarginOrderPost:   "/sapi/v1/margin/order",      // POST /sapi/v1/margin/order (HMAC SHA256) 杠杆账户下单 (TRADE)
	SpotMarginOrderDelete: "/sapi/v1/margin/order",      //DELETE 撤销杠杆账户订单 (TRADE)
	SpotMarginAllOrders:   "/sapi/v1/margin/allOrders",  //GET 查询杠杆账户所有订单 (USER_DATA)
	SpotMarginOpenOrders:  "/sapi/v1/margin/openOrders", //GET 查询杠杆账户挂单记录 (USER_DATA)

	SpotMarginTransfer:         "/sapi/v1/margin/transfer",          //POST  全仓杠杆账户划转 (MARGIN)
	SpotMarginIsolatedTransfer: "/sapi/v1/margin/isolated/transfer", //POST  逐仓杠杆账户划转 (MARGIN)
	SpotMarginLoan:             "/sapi/v1/margin/loan",              //POST  杠杆账户借贷 (MARGIN) 支持逐仓和全仓
	SpotMarginRepay:            "/sapi/v1/margin/repay",             //POST  杠杆账户归还借贷 (MARGIN) 支持逐仓和全仓

	SpotMarginMaxLeverage:        "/sapi/v1/margin/max-leverage",       //POST 调整全仓最大杠杆倍数（USER_DATA)
	SpotMarginTradeCoeff:         "/sapi/v1/margin/tradeCoeff",         // GET 获取用户个人杠杆账户信息汇总（USER_DATA）全仓
	SpotMarginCrossMarginData:    "/sapi/v1/margin/crossMarginData",    //GET接口 获取全仓杠杆利率及限额 (USER_DATA)
	SpotMarginIsolatedMarginData: "/sapi/v1/margin/isolatedMarginData", // GET接口 获取逐仓杠杆利率及限额 (USER_DATA)
	SpotMarginIsolatedMarginTier: "/sapi/v1/margin/isolatedMarginTier", // GET接口 获取逐仓杠杆档位信息 (USER_DATA)

	//现货账户接口
	SpotAccountApiTradingStatus: "/sapi/v1/account/apiTradingStatus", //GET接口 账户API交易状态(USER_DATA)
	SpotAccount:                 "/api/v3/account",                   //GET接口 账户信息 (USER_DATA)
	SpotAssetGetFundingAsset:    "/sapi/v1/asset/get-funding-asset",  //POST接口 资金账户 (USER_DATA)
	SpotAssetTransferPost:       "/sapi/v1/asset/transfer",           //POST接口 用户万向划转 (USER_DATA)
	SpotAssetTransferGet:        "/sapi/v1/asset/transfer",           //GET接口 查询用户万向划转历史 (USER_DATA)
	SpotAssetTradeFee:           "/sapi/v1/asset/tradeFee",           //GET接口 查询用户交易手续费率 (USER_DATA)

	//现货订单接口
	SpotOpenOrders: "/api/v3/openOrders", //GET接口 查询当前挂单 (USER_DATA)
	SpotAllOrders:  "/api/v3/allOrders",  //GET接口 查询所有订单 (USER_DATA)

	SpotOrderGet:           "/api/v3/order",               //GET接口 查询订单 (USER_DATA)
	SpotOrderPost:          "/api/v3/order",               //POST 现货下单 (TRADE)
	SpotOrderDelete:        "/api/v3/order",               //DELETE接口 撤销订单 (TRADE)
	SpotOrderCancelReplace: "/api/v3/order/cancelReplace", //POST接口 撤消挂单再下单 (TRADE)
	SpotMyTrades:           "/api/v3/myTrades",            //GET接口 账户成交历史 (USER_DATA)

	//通用接口
	SpotPing:         "/api/v3/ping",         //GET接口 测试连通性
	SpotServerTime:   "/api/v3/time",         //GET接口 获取服务器时间
	SpotExchangeInfo: "/api/v3/exchangeInfo", //GET接口 获取交易规范

	//行情接口
	SpotTickerPrice:      "/api/v3/ticker/price",      //GET接口 获取交易对最新价格
	SpotKlines:           "/api/v3/klines",            //GET接口 获取K线数据
	SpotDepth:            "/api/v3/depth",             //GET接口 获取深度信息
	SpotTrades:           "/api/v3/trades",            //GET接口 近期成交列表
	SpotHistoricalTrades: "/api/v3/historicalTrades",  //GET接口 历史成交记录
	SpotAggTrades:        "/api/v3/aggTrades",         //GET接口 近期成交(归集)
	SpotAvgPrice:         "/api/v3/avgPrice",          //GET接口 当前平均价格
	SpotUiKlines:         "/api/v3/uiKlines",          //GET接口 UIK线数据
	SpotTicker24hr:       "/api/v3/ticker/24hr",       //GET接口 24hr 价格变动情况
	SpotTickerBookTicker: "/api/v3/ticker/bookTicker", //GET接口 当前最优挂单
	SpotTicker:           "/api/v3/ticker",            //GET接口 滚动窗口价格变动统计

	//Ws账户推送相关
	SpotUserDataStreamPost:   "/api/v3/userDataStream", //POST接口 (现货账户)生成 Listen Key (USER_STREAM)
	SpotUserDataStreamPut:    "/api/v3/userDataStream", //PUT接口 (现货账户)延长 Listen Key 有效期 (USER_STREAM)
	SpotUserDataStreamDelete: "/api/v3/userDataStream", //DELETE接口 (现货账户)关闭 Listen Key (USER_STREAM)

	SpotMarginUserDataStreamPost:   "/sapi/v1/userDataStream", //POST接口 (杠杆账户)生成 Listen Key (USER_STREAM)
	SpotMarginUserDataStreamPut:    "/sapi/v1/userDataStream", //PUT接口 (杠杆账户)延长 Listen Key 有效期 (USER_STREAM)
	SpotMarginUserDataStreamDelete: "/sapi/v1/userDataStream", //DELETE接口 (杠杆账户)关闭 Listen Key (USER_STREAM)

	SpotMarginIsolatedUserDataStreamPost:   "/sapi/v1/userDataStream/isolated", //POST接口 (逐仓杠杆账户)生成 Listen Key (USER_STREAM)
	SpotMarginIsolatedUserDataStreamPut:    "/sapi/v1/userDataStream/isolated", //PUT接口 (逐仓杠杆账户)延长 Listen Key 有效期 (USER_STREAM)
	SpotMarginIsolatedUserDataStreamDelete: "/sapi/v1/userDataStream/isolated", //DELETE接口 (逐仓杠杆账户)关闭 Listen Key (USER_STREAM)

	// 经纪商子母账号接口
	SpotBrokerSubAccountPost:                        "/sapi/v1/broker/subAccount",                              //POST接口 创建经纪商子账户
	SpotBrokerSubAccountGet:                         "/sapi/v1/broker/subAccount",                              //GET接口 查询经纪商子账户
	SpotBrokerSubAccountApiPost:                     "/sapi/v1/broker/subAccountApi",                           //POST创建经纪商子账号ApiKey
	SpotBrokerSubAccountApiGet:                      "/sapi/v1/broker/subAccountApi",                           //GET查询经纪商子账号ApiKey
	SpotBrokerSubAccountApiDelete:                   "/sapi/v1/broker/subAccountApi",                           //DELETE删除经纪商子账号ApiKey
	SpotBrokerSubAccountFutures:                     "/sapi/v1/broker/subAccount/futures",                      //POST接口 授权经纪商子账户Futures权限
	SpotBrokerSubAccountPermissionUniversalTransfer: "/sapi/v1/broker/subAccount/permission/universalTransfer", //POST接口 授权经纪商子账户万能划转权限
	SpotBrokerSubAccountApiIpRestrictionPost:        "/sapi/v1/broker/subAccountApi/ipRestriction",             //POST接口 更新经纪商子账户API Key IP白名单
	SpotBrokerSubAccountApiIpRestrictionGet:         "/sapi/v1/broker/subAccountApi/ipRestriction",             //GET接口 查询经纪商子账户API Key IP白名单
	SpotBrokerSubAccountApiIpRestrictionDelete:      "/sapi/v1/broker/subAccountApi/ipRestriction",             //DELETE接口 删除经纪商子账户API Key IP白名单
	SpotBrokerSubAccountDepositHist:                 "/sapi/v1/broker/subAccount/depositHist",                  //GET接口 查询经纪商子账户充值历史
	SpotBrokerUniversalTransferPost:                 "/sapi/v1/broker/universalTransfer",                       //POST接口 经纪商子账户万能划转
	SpotBrokerUniversalTransferGet:                  "/sapi/v1/broker/universalTransfer",                       //GET接口 查询经纪商子账户万能划转历史
}
