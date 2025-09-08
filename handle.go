package mybinanceapi

import (
	"fmt"
	"strconv"
	"strings"
)

type BinanceErrorRes struct {
	Code interface{} `json:"code"`
	Msg  string      `json:"msg"`
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

		if res.Code != nil {

			var result T
			err = json.Unmarshal(data, &result)
			if err != nil {
				return res, nil
			}
			res.Result = result
		}
	} else {
		var result T
		err = json.Unmarshal(data, &result)
		if err != nil {
			log.Error("rest返回值序列化错误", err)
		}
		res.Result = result
	}
	// log.Warn(res)
	return res, err
}
func (err *BinanceErrorRes) handlerError() error {
	if codeF64, ok := err.Code.(float64); ok {
		codeInt := int(codeF64)
		if codeInt != 0 && codeInt != 200 {
			return fmt.Errorf("request error:[code:%v][message:%v]", err.Code, err.Msg)
		} else {
			return nil
		}
	}
	if codeStr, ok := err.Code.(string); ok {
		codeInt, _ := strconv.Atoi(codeStr)
		if codeInt != 0 && codeInt != 200 {
			return fmt.Errorf("request error:[code:%v][message:%v]", err.Code, err.Msg)
		} else {
			return nil
		}
	}
	return nil
}
