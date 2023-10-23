package binanceservice

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

// GetRequest 发送 get 请求
func GetRequest(url string, isGzip bool) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	if isGzip { // 请求 header 添加 gzip
		req.Header.Add("Content-Encoding", "gzip")
		req.Header.Add("Accept-Encoding", "gzip")
	}
	req.Close = true

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
	return data, err
}

func GetRequestWithHeader(url string, headerMap map[string]string, isGzip bool) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
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
	return data, err
}

func PostRequestWithHeader(url string, reqBody []byte, headerMap map[string]string, isGzip bool) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
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
	return data, err
}

type MyBinance struct {
}

const (
	BINANCE_API_SPOT_HTTP      = "api.binance.com"
	BINANCE_API_FUTURE_HTTP    = "fapi.binance.com"
	BINANCE_API_SWAP_HTTP      = "dapi.binance.com"
	BINANCE_API_SPOT_WEBSOCKET = "wss://stream.binance.com:9443/ws"
	IS_GZIP                    = true
)

type ApiType int

const (
	SPOT ApiType = iota
	FUTURE
	SWAP
)

type Client struct {
	ApiKey    string
	ApiSecret string
}

type SpotRestClient struct {
	c *Client
}
type FutureRestClient struct {
	c *Client
}
type SwapRestClient struct {
	c *Client
}

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

// 通用鉴权POST接口调用
func binanceCallApiWithSecretPost[T any](client *Client, url string, reqBody []byte) (*T, error) {
	body, err := PostRequestWithHeader(url, []byte{}, map[string]string{"X-MBX-APIKEY": client.ApiKey}, IS_GZIP)
	if err != nil {
		return nil, err
	}
	res, err := handlerCommonRest[T](body)
	if err != nil {
		return nil, err
	}
	return &res.Result, res.handlerError()
}

// 通用鉴权GET接口调用
func binanceCallApiWithSecretGet[T any](client *Client, url string) (*T, error) {
	body, err := GetRequestWithHeader(url, map[string]string{"X-MBX-APIKEY": client.ApiKey}, IS_GZIP)
	if err != nil {
		return nil, err
	}
	// log.Warn(string(body))
	res, err := handlerCommonRest[T](body)
	if err != nil {
		return nil, err
	}
	return &res.Result, res.handlerError()
}

// 通用鉴权POST接口封装
func binanceHandlerRequestApiWithSecretPost[T any](apiType ApiType, request *T, name, secret string) (string, []byte) {
	query := binanceHandlerReq(request)
	log.Debug(query)
	sign := HmacSha256(secret, query)

	u := url.URL{
		Scheme:   "https",
		Host:     BinanceGetRestHostByApiType(apiType),
		Path:     name,
		RawQuery: query + "&signature=" + sign,
	}
	reqBody, _ := json.Marshal(request)
	return u.String(), reqBody
}

// 通用鉴权GET接口封装
func binanceHandlerRequestApiWithSecretGet[T any](apiType ApiType, request *T, name, secret string) string {
	query := binanceHandlerReq(request)
	sign := HmacSha256(secret, query)
	u := url.URL{
		Scheme:   "https",
		Host:     BinanceGetRestHostByApiType(apiType),
		Path:     name,
		RawQuery: query + "&signature=" + sign,
	}
	log.Debug(u.String())
	return u.String()
}

// 通用非鉴权GET接口封装
func binanceHandlerRequestApiGet[T any](apiType ApiType, request *T, name string) string {
	query := binanceHandlerReq(request)
	u := url.URL{
		Scheme:   "https",
		Host:     BinanceGetRestHostByApiType(apiType),
		Path:     name,
		RawQuery: query,
	}
	log.Debug(u.String())
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
		return BINANCE_API_SPOT_HTTP
	case FUTURE:
		return BINANCE_API_FUTURE_HTTP
	case SWAP:
		return BINANCE_API_SWAP_HTTP
	default:
		return ""
	}
}
