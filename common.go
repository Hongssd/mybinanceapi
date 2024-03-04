package mybinanceapi

import (
	"bytes"
	"compress/gzip"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"io"

	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"

	jsoniter "github.com/json-iterator/go"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

const (
	BIT_BASE_10 = 10
	BIT_SIZE_64 = 64
	BIT_SIZE_32 = 32
)

type RequestType string

const (
	GET    = "GET"
	POST   = "POST"
	DELETE = "DELETE"
	PUT    = "PUT"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

var log = logrus.New()

func SetLogger(logger *logrus.Logger) {
	log = logger
}

func GetPointer[T any](v T) *T {
	return &v
}

func HmacSha256(secret, data string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}

// Request 发送请求
func Request(url string, method string, isGzip bool) ([]byte, error) {
	return RequestWithHeader(url, method, map[string]string{}, isGzip)
}

func RequestWithHeader(url string, method string, headerMap map[string]string, isGzip bool) ([]byte, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	for k, v := range headerMap {
		req.Header.Set(k, v)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	if isGzip { // 请求 header 添加 gzip
		req.Header.Add("Content-Encoding", "gzip")
		req.Header.Add("Accept-Encoding", "gzip")
	}
	req.Close = true

	log.Debug(req.URL.String())

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body := resp.Body
	if resp.Header.Get("Content-Encoding") == "gzip" {
		body, err = gzip.NewReader(resp.Body)
		if err != nil {
			log.Error(err)
			return nil, err
		}
	}
	data, err := io.ReadAll(body)
	log.Debug(string(data))
	return data, err
}

type MyBinance struct {
}

const (
	BINANCE_API_SPOT_HTTP        = "api.binance.com"
	BINANCE_API_FUTURE_HTTP      = "fapi.binance.com"
	BINANCE_API_SWAP_HTTP        = "dapi.binance.com"
	TEST_BINANCE_API_SPOT_HTTP   = "testnet.binance.vision"
	TEST_BINANCE_API_FUTURE_HTTP = "testnet.binancefuture.com"
	TEST_BINANCE_API_SWAP_HTTP   = "testnet.binancefuture.com"
	IS_GZIP                      = true
)

type NetType int

const (
	MAIN_NET NetType = iota
	TEST_NET
)

var NowNetType = MAIN_NET

func SetNetType(netType NetType) {
	NowNetType = netType
}

type ApiType int

const (
	SPOT ApiType = iota
	FUTURE
	SWAP
)

func (apiType *ApiType) String() string {
	switch *apiType {
	case SPOT:
		return "SPOT"
	case FUTURE:
		return "FUTURE"
	case SWAP:
		return "SWAP"
	}
	return ""
}

type Client struct {
	ApiKey    string
	ApiSecret string
}
type RestClient struct {
	c *Client
}

type SpotRestClient RestClient
type FutureRestClient RestClient
type SwapRestClient RestClient

func (*MyBinance) NewSpotRestClient(apiKey string, apiSecret string) *SpotRestClient {
	client := &SpotRestClient{
		&Client{
			ApiKey:    apiKey,
			ApiSecret: apiSecret,
		},
	}
	return client
}
func (*MyBinance) NewFutureRestClient(apiKey string, apiSecret string) *FutureRestClient {
	client := &FutureRestClient{
		&Client{
			ApiKey:    apiKey,
			ApiSecret: apiSecret,
		},
	}
	return client
}
func (*MyBinance) NewSwapRestClient(apiKey string, apiSecret string) *SwapRestClient {
	client := &SwapRestClient{
		&Client{
			ApiKey:    apiKey,
			ApiSecret: apiSecret,
		},
	}
	return client
}

var serverTimeDelta int64 = 0

func setServerTimeDelta(delta int64) {
	serverTimeDelta = delta
}

// 通用鉴权接口调用
func binanceCallApiWithSecret[T any](client *Client, url, method string) (*T, error) {
	body, err := RequestWithHeader(url, method, map[string]string{"X-MBX-APIKEY": client.ApiKey}, IS_GZIP)
	if err != nil {
		return nil, err
	}
	res, err := handlerCommonRest[T](body)
	if err != nil {
		return nil, err
	}
	return &res.Result, res.handlerError()
}

// 通用鉴权接口封装
func binanceHandlerRequestApiWithSecret[T any](apiType ApiType, request *T, name, secret string) string {
	query := binanceHandlerReq(request)
	sign := HmacSha256(secret, query)

	u := url.URL{
		Scheme:   "https",
		Host:     BinanceGetRestHostByApiType(apiType),
		Path:     name,
		RawQuery: query + "&signature=" + sign,
	}
	// log.Debug(u.RequestURI() + "---" + u.Query().Encode())
	// log.Debug(u.String())
	return u.String()
}

// 通用非鉴权接口封装
func binanceHandlerRequestApi[T any](apiType ApiType, request *T, name string) string {
	query := binanceHandlerReq(request)
	u := url.URL{
		Scheme:   "https",
		Host:     BinanceGetRestHostByApiType(apiType),
		Path:     name,
		RawQuery: query,
	}
	// log.Debug(u.String())
	return u.String()
}

func binanceHandlerReq[T any](req *T) string {
	var paramBuffer bytes.Buffer

	t := reflect.TypeOf(req)
	v := reflect.ValueOf(req)
	if v.IsNil() {
		return ""
	}
	t = t.Elem()
	v = v.Elem()
	count := v.NumField()
	for i := 0; i < count; i++ {
		paramName := t.Field(i).Tag.Get("json")
		switch v.Field(i).Elem().Kind() {
		case reflect.String:
			paramBuffer.WriteString(paramName + "=" + v.Field(i).Elem().String() + "&")
		case reflect.Int, reflect.Int64:
			paramBuffer.WriteString(paramName + "=" + strconv.FormatInt(v.Field(i).Elem().Int(), BIT_BASE_10) + "&")
		case reflect.Float32, reflect.Float64:
			paramBuffer.WriteString(paramName + "=" + decimal.NewFromFloat(v.Field(i).Elem().Float()).String() + "&")
		case reflect.Bool:
			paramBuffer.WriteString(paramName + "=" + strconv.FormatBool(v.Field(i).Elem().Bool()) + "&")
		case reflect.Struct:
			sv := reflect.ValueOf(v.Field(i).Interface())
			ToStringMethod := sv.MethodByName("String")
			params := make([]reflect.Value, 0)
			result := ToStringMethod.Call(params)
			paramBuffer.WriteString(paramName + "=" + result[0].String() + "&")
		case reflect.Slice:
			s := v.Field(i).Interface()
			d, _ := json.Marshal(s)
			paramBuffer.WriteString(paramName + "=" + url.QueryEscape(string(d)) + "&")
		case reflect.Invalid:
		default:
			log.Errorf("req type error %s:%s", paramName, v.Field(i).Elem().Kind())
		}
	}
	return strings.TrimRight(paramBuffer.String(), "&")
}

func BinanceGetRestHostByApiType(apiType ApiType) string {
	switch apiType {
	case SPOT:
		switch NowNetType {
		case MAIN_NET:
			return BINANCE_API_SPOT_HTTP
		case TEST_NET:
			return TEST_BINANCE_API_SPOT_HTTP
		}
	case FUTURE:
		switch NowNetType {
		case MAIN_NET:
			return BINANCE_API_FUTURE_HTTP
		case TEST_NET:
			return TEST_BINANCE_API_FUTURE_HTTP
		}
	case SWAP:
		switch NowNetType {
		case MAIN_NET:
			return BINANCE_API_SWAP_HTTP
		case TEST_NET:
			return TEST_BINANCE_API_SWAP_HTTP
		}
	}
	return ""
}

func interfaceStringToFloat64(inter interface{}) float64 {
	return stringToFloat64(inter.(string))
}

func interfaceStringToInt64(inter interface{}) int64 {
	return int64(inter.(float64))
}

func stringToFloat64(str string) float64 {
	f, _ := strconv.ParseFloat(str, BIT_SIZE_64)
	return f
}
