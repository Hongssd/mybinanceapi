package binanceservice

import (
	"fmt"
	"strings"
)

type BinanceErrorRes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type BinanceRestRes[T any] struct {
	Result          T //请求结果
	BinanceErrorRes   //错误信息
}

func handlerCommonRest[T any](data []byte) (*BinanceRestRes[T], error) {
	res := &BinanceRestRes[T]{}
	var err error
	// log.Warn(string(data))
	if strings.Contains(string(data), "code") {
		err = json.Unmarshal(data, res)
		if err != nil {
			log.Error("rest返回值获取失败", err)
		}
	} else {
		var result T
		err = json.Unmarshal(data, &result)
		if err != nil {
			log.Error("rest返回值获取失败", err)
		}
		res.Result = result
	}
	return res, err
}
func (err *BinanceErrorRes) handlerError() error {
	if err.Code != 0 && err.Code != 200 {
		return fmt.Errorf("request error:[code:%v][message:%v]", err.Code, err.Msg)
	} else {
		return nil
	}

}
