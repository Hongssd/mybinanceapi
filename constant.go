package mybinanceapi

type AssetTransferType string

const (
	MAIN_UMFUTURE                 AssetTransferType = "MAIN_UMFUTURE"                 //现货钱包转向U本位合约钱包
	MAIN_CMFUTURE                 AssetTransferType = "MAIN_CMFUTURE"                 //现货钱包转向币本位合约钱包
	MAIN_MARGIN                   AssetTransferType = "MAIN_MARGIN"                   //现货钱包转向杠杆全仓钱包
	UMFUTURE_MAIN                 AssetTransferType = "UMFUTURE_MAIN"                 //U本位合约钱包转向现货钱包
	UMFUTURE_MARGIN               AssetTransferType = "UMFUTURE_MARGIN"               //U本位合约钱包转向杠杆全仓钱包
	CMFUTURE_MAIN                 AssetTransferType = "CMFUTURE_MAIN"                 //币本位合约钱包转向现货钱包
	MARGIN_MAIN                   AssetTransferType = "MARGIN_MAIN"                   //杠杆全仓钱包转向现货钱包
	MARGIN_UMFUTURE               AssetTransferType = "MARGIN_UMFUTURE"               //杠杆全仓钱包转向U本位合约钱包
	MARGIN_CMFUTURE               AssetTransferType = "MARGIN_CMFUTURE"               //杠杆全仓钱包转向币本位合约钱包
	CMFUTURE_MARGIN               AssetTransferType = "CMFUTURE_MARGIN"               //币本位合约钱包转向杠杆全仓钱包
	ISOLATEDMARGIN_MARGIN         AssetTransferType = "ISOLATEDMARGIN_MARGIN"         //杠杆逐仓钱包转向杠杆全仓钱包
	MARGIN_ISOLATEDMARGIN         AssetTransferType = "MARGIN_ISOLATEDMARGIN"         //杠杆全仓钱包转向杠杆逐仓钱包
	ISOLATEDMARGIN_ISOLATEDMARGIN AssetTransferType = "ISOLATEDMARGIN_ISOLATEDMARGIN" //杠杆逐仓钱包转向杠杆逐仓钱包
	MAIN_FUNDING                  AssetTransferType = "MAIN_FUNDING"                  //现货钱包转向资金钱包
	FUNDING_MAIN                  AssetTransferType = "FUNDING_MAIN"                  //资金钱包转向现货钱包
	FUNDING_UMFUTURE              AssetTransferType = "FUNDING_UMFUTURE"              //资金钱包转向U本位合约钱包
	UMFUTURE_FUNDING              AssetTransferType = "UMFUTURE_FUNDING"              //U本位合约钱包转向资金钱包
	MARGIN_FUNDING                AssetTransferType = "MARGIN_FUNDING"                //杠杆全仓钱包转向资金钱包
	FUNDING_MARGIN                AssetTransferType = "FUNDING_MARGIN"                //资金钱包转向杠杆全仓钱包
	FUNDING_CMFUTURE              AssetTransferType = "FUNDING_CMFUTURE"              //资金钱包转向币本位合约钱包
	CMFUTURE_FUNDING              AssetTransferType = "CMFUTURE_FUNDING"              //币本位合约钱包转向资金钱包
	MAIN_OPTION                   AssetTransferType = "MAIN_OPTION"                   //现货钱包转向期权钱包
	OPTION_MAIN                   AssetTransferType = "OPTION_MAIN"                   //期权钱包转向现货钱包
	UMFUTURE_OPTION               AssetTransferType = "UMFUTURE_OPTION"               //U本位合约钱包转向期权钱包
	OPTION_UMFUTURE               AssetTransferType = "OPTION_UMFUTURE"               //期权钱包转向U本位合约钱包
	MARGIN_OPTION                 AssetTransferType = "MARGIN_OPTION"                 //杠杆全仓钱包转向期权钱包
	OPTION_MARGIN                 AssetTransferType = "OPTION_MARGIN"                 //期权全仓钱包转向杠杆钱包
	FUNDING_OPTION                AssetTransferType = "FUNDING_OPTION"                //资金钱包转向期权钱包
	OPTION_FUNDING                AssetTransferType = "OPTION_FUNDING"                //期权钱包转向资金钱包
	MAIN_PORTFOLIO_MARGIN         AssetTransferType = "MAIN_PORTFOLIO_MARGIN"         //现货钱包转向统一账户钱包
	PORTFOLIO_MARGIN_MAIN         AssetTransferType = "PORTFOLIO_MARGIN_MAIN"         //统一账户钱包转向现货钱包
	MAIN_ISOLATED_MARGIN          AssetTransferType = "MAIN_ISOLATED_MARGIN"          //现货钱包转向逐仓账户钱包
	ISOLATED_MARGIN_MAIN          AssetTransferType = "ISOLATED_MARGIN_MAIN"          //逐仓钱包转向现货账户钱包

)
