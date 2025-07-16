package mybinanceapi

import (
	"fmt"
	"strconv"
	"strings"
)

type BinanceErrorRes struct {
	Code    interface{} `json:"code"`
	CodeInt int         `json:"codeInt"`
	CodeStr string      `json:"codeStr"`
	Msg     string      `json:"msg"`
}

type BinanceRestRes[T any] struct {
	Result          T //请求结果
	BinanceErrorRes   //错误信息
}

func handlerCommonRest[T any](data []byte) (*BinanceRestRes[T], error) {
	res := &BinanceRestRes[T]{}
	var err error
	// log.Warn(string(data))
	if strings.Contains(string(data), "code") && !strings.HasPrefix(string(data), "[") {
		err = json.Unmarshal(data, res)
		if err != nil {
			log.Error("rest返回值获取失败", err)
		}

		if codeF64, ok := res.Code.(float64); ok {
			res.CodeInt = int(codeF64)
			res.CodeStr = fmt.Sprintf("%v", codeF64)
		} else if codeInt, ok := res.Code.(int); ok {
			res.CodeInt = codeInt
			res.CodeStr = fmt.Sprintf("%v", codeInt)
		} else if codeStr, ok := res.Code.(string); ok {
			codeInt, err := strconv.Atoi(codeStr)
			if err == nil {
				res.CodeInt = codeInt
				res.CodeStr = codeStr
			} else {
				res.CodeStr = codeStr
			}
		} else {
			res.CodeStr = fmt.Sprintf("%v", res.Code)
		}
	} else {
		var result T
		err = json.Unmarshal(data, &result)
		if err != nil {
			log.Error("rest返回值序列化错误", err)
		}
		res.Result = result
	}
	return res, err
}
func (err *BinanceErrorRes) handlerError() error {
	if err.CodeInt != 0 && err.CodeInt != 200 {
		return fmt.Errorf("request error:[code:%v][message:%v]", err.Code, err.Msg)
	} else {
		return nil
	}

}
