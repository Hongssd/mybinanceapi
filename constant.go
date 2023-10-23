package mybinanceapi

type AssertTransferType string

const (
	MAIN_UMFUTURE                 AssertTransferType = "MAIN_UMFUTURE"                 //现货钱包转向U本位合约钱包
	MAIN_CMFUTURE                 AssertTransferType = "MAIN_CMFUTURE"                 //现货钱包转向币本位合约钱包
	MAIN_MARGIN                   AssertTransferType = "MAIN_MARGIN"                   //现货钱包转向杠杆全仓钱包
	UMFUTURE_MAIN                 AssertTransferType = "UMFUTURE_MAIN"                 //U本位合约钱包转向现货钱包
	UMFUTURE_MARGIN               AssertTransferType = "UMFUTURE_MARGIN"               //U本位合约钱包转向杠杆全仓钱包
	CMFUTURE_MAIN                 AssertTransferType = "CMFUTURE_MAIN"                 //币本位合约钱包转向现货钱包
	MARGIN_MAIN                   AssertTransferType = "MARGIN_MAIN"                   //杠杆全仓钱包转向现货钱包
	MARGIN_UMFUTURE               AssertTransferType = "MARGIN_UMFUTURE"               //杠杆全仓钱包转向U本位合约钱包
	MARGIN_CMFUTURE               AssertTransferType = "MARGIN_CMFUTURE"               //杠杆全仓钱包转向币本位合约钱包
	CMFUTURE_MARGIN               AssertTransferType = "CMFUTURE_MARGIN"               //币本位合约钱包转向杠杆全仓钱包
	ISOLATEDMARGIN_MARGIN         AssertTransferType = "ISOLATEDMARGIN_MARGIN"         //杠杆逐仓钱包转向杠杆全仓钱包
	MARGIN_ISOLATEDMARGIN         AssertTransferType = "MARGIN_ISOLATEDMARGIN"         //杠杆全仓钱包转向杠杆逐仓钱包
	ISOLATEDMARGIN_ISOLATEDMARGIN AssertTransferType = "ISOLATEDMARGIN_ISOLATEDMARGIN" //杠杆逐仓钱包转向杠杆逐仓钱包
	MAIN_FUNDING                  AssertTransferType = "MAIN_FUNDING"                  //现货钱包转向资金钱包
	FUNDING_MAIN                  AssertTransferType = "FUNDING_MAIN"                  //资金钱包转向现货钱包
	FUNDING_UMFUTURE              AssertTransferType = "FUNDING_UMFUTURE"              //资金钱包转向U本位合约钱包
	UMFUTURE_FUNDING              AssertTransferType = "UMFUTURE_FUNDING"              //U本位合约钱包转向资金钱包
	MARGIN_FUNDING                AssertTransferType = "MARGIN_FUNDING"                //杠杆全仓钱包转向资金钱包
	FUNDING_MARGIN                AssertTransferType = "FUNDING_MARGIN"                //资金钱包转向杠杆全仓钱包
	FUNDING_CMFUTURE              AssertTransferType = "FUNDING_CMFUTURE"              //资金钱包转向币本位合约钱包
	CMFUTURE_FUNDING              AssertTransferType = "CMFUTURE_FUNDING"              //币本位合约钱包转向资金钱包
	MAIN_OPTION                   AssertTransferType = "MAIN_OPTION"                   //现货钱包转向期权钱包
	OPTION_MAIN                   AssertTransferType = "OPTION_MAIN"                   //期权钱包转向现货钱包
	UMFUTURE_OPTION               AssertTransferType = "UMFUTURE_OPTION"               //U本位合约钱包转向期权钱包
	OPTION_UMFUTURE               AssertTransferType = "OPTION_UMFUTURE"               //期权钱包转向U本位合约钱包
	MARGIN_OPTION                 AssertTransferType = "MARGIN_OPTION"                 //杠杆全仓钱包转向期权钱包
	OPTION_MARGIN                 AssertTransferType = "OPTION_MARGIN"                 //期权全仓钱包转向杠杆钱包
	FUNDING_OPTION                AssertTransferType = "FUNDING_OPTION"                //资金钱包转向期权钱包
	OPTION_FUNDING                AssertTransferType = "OPTION_FUNDING"                //期权钱包转向资金钱包
	MAIN_PORTFOLIO_MARGIN         AssertTransferType = "MAIN_PORTFOLIO_MARGIN"         //现货钱包转向统一账户钱包
	PORTFOLIO_MARGIN_MAIN         AssertTransferType = "PORTFOLIO_MARGIN_MAIN"         //统一账户钱包转向现货钱包
	MAIN_ISOLATED_MARGIN          AssertTransferType = "MAIN_ISOLATED_MARGIN"          //现货钱包转向逐仓账户钱包
	ISOLATED_MARGIN_MAIN          AssertTransferType = "ISOLATED_MARGIN_MAIN"          //逐仓钱包转向现货账户钱包

)
