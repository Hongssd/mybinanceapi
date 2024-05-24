package mybinanceapi

type SwapApi int

const (
	//币合约账户接口
	SwapAccount              SwapApi = iota //账户信息 (USER_DATA)
	SwapPositionSideDualGet                 //查询持仓模式 (USER_DATA)
	SwapLeverageBracket                     //杠杆分层标准 (USER_DATA)
	SwapPositionSideDualPost                //更改持仓模式 (TRADE)
	SwapLeverage                            //调整开仓杠杆 (TRADE)
	SwapMarginType                          //变换逐全仓模式 (TRADE)

	//币合约订单接口
	SwapOpenOrders  //查询当前挂单 (USER_DATA)
	SwapAllOrders   //查询所有订单 (USER_DATA)
	SwapOrderPost   //下单 (TRADE)
	SwapOrderGet    //查询订单 (USER_DATA)
	SwapOrderDelete //撤销订单 (TRADE)

	//通用接口
	SwapPing         //测试连接
	SwapServerTime   //获取服务器时间
	SwapExchangeInfo //交易规则和交易对信息

	//行情接口
	SwapKlines      //获取K线数据
	SwapTickerPrice //获取交易对最新价格
	SwapDepth       //获取深度信息

	//Ws账户推送相关
	SwapListenKeyPost //生成listenKey (USER_STREAM)
	SwapListenKeyPut  //延长listenKey有效期 (USER_STREAM)
	SwapListenKeyDel  //关闭listenKey (USER_STREAM)
)

var SwapApiMap = map[SwapApi]string{

	//币合约账户接口
	SwapAccount:              "/dapi/v1/account",           //GET接口 账户信息 (USER_DATA)
	SwapPositionSideDualGet:  "/dapi/v1/positionSide/dual", //GET接口 (HMAC SHA256) 查询持仓模式(USER_DATA)
	SwapLeverageBracket:      "/dapi/v2/leverageBracket",   //GET接口 杠杆分层标准 (USER_DATA)
	SwapPositionSideDualPost: "/dapi/v1/positionSide/dual", //POST接口 (HMAC SHA256) 更改持仓模式(TRADE)
	SwapLeverage:             "/dapi/v1/leverage",          //POST接口 (HMAC SHA256) 调整开仓杠杆 (TRADE)
	SwapMarginType:           "/dapi/v1/marginType",        //POST接口 (HMAC SHA256) 变换逐全仓模式 (TRADE)

	//币合约订单接口
	SwapOpenOrders:  "/dapi/v1/openOrders", //GET接口 (HMAC SHA256) 查询当前挂单 (USER_DATA)
	SwapAllOrders:   "/dapi/v1/allOrders",  //GET接口 (HMAC SHA256) 查询所有订单 (USER_DATA)
	SwapOrderPost:   "/dapi/v1/order",      //POST接口 (HMAC SHA256) 下单 (TRADE)
	SwapOrderGet:    "/dapi/v1/order",      //GET接口 (HMAC SHA256) 查询订单 (USER_DATA)
	SwapOrderDelete: "/dapi/v1/order",      //DELETE接口 (HMAC SHA256) 撤销订单 (TRADE)

	//通用接口
	SwapPing:         "/dapi/v1/ping",         //GET接口 测试连接
	SwapServerTime:   "/dapi/v1/time",         //GET接口 获取服务器时间
	SwapExchangeInfo: "/dapi/v1/exchangeInfo", //GET接口 交易规则和交易对信息

	//行情接口
	SwapKlines:      "/dapi/v1/klines",       //GET接口 获取K线数据
	SwapTickerPrice: "/dapi/v1/ticker/price", //GET接口 获取交易对最新价格
	SwapDepth:       "/dapi/v1/depth",        //GET接口 获取深度信息

	//Ws账户推送相关
	SwapListenKeyPost: "/dapi/v1/listenKey", //POST接口 (HMAC SHA256) 创建listenKey (USER_STREAM)\
	SwapListenKeyPut:  "/dapi/v1/listenKey", //PUT接口 (HMAC SHA256) 延长listenKey有效期 (USER_STREAM)
	SwapListenKeyDel:  "/dapi/v1/listenKey", //DELETE接口 (HMAC SHA256) 关闭listenKey (USER_STREAM)
}
