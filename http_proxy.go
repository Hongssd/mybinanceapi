package mybinanceapi

import (
	"bytes"
	"compress/gzip"
	"crypto/tls"
	"errors"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"time"

	"github.com/robfig/cron/v3"
)

type RestProxy struct {
	ProxyUrl string //代理的协议IP端口URL

	SpotWeight      ProxyWeight
	FutureWeight    ProxyWeight
	SwapWeight      ProxyWeight
	PortfolioMargin ProxyWeight
}

type ProxyWeight struct {
	UsedWeight   int  //已用权重
	Is429Limited bool //是否被限制
}

func (w *ProxyWeight) restore() {
	w.UsedWeight = 0
	w.Is429Limited = false
}

var proxyList = []*RestProxy{}

var UseProxy = false
var WsUseProxy = false

// ConnectionPoolConfig 连接池配置
type ConnectionPoolConfig struct {
	Enable                 bool          // 是否启用连接池（默认：true）
	MaxIdleConns           int           // 最大空闲连接数（默认：100）
	MaxIdleConnsPerHost    int           // 每个 host 的最大空闲连接数（默认：10）
	IdleConnTimeout        time.Duration // 空闲连接超时时间（默认：90秒）
	TLSHandshakeTimeout    time.Duration // TLS 握手超时（默认：10秒）
	ResponseHeaderTimeout  time.Duration // 等待响应头超时（默认：10秒）
	ExpectContinueTimeout  time.Duration // 期望继续超时（默认：1秒）
	ForceAttemptHTTP2      bool          // 是否尝试使用 HTTP/2（默认：true）
	DisableCompression     bool          // 是否禁用压缩（默认：false）
	DisableKeepAlives      bool          // 是否禁用 Keep-Alive（默认：false）
	MaxConnsPerHost        int           // 每个 host 的最大连接数，0 表示不限制（默认：0）
	MaxResponseHeaderBytes int64         // 最大响应头字节数，0 表示使用默认值（默认：0）
}

// 默认连接池配置
var defaultPoolConfig = ConnectionPoolConfig{
	Enable:                 true,
	MaxIdleConns:           100,
	MaxIdleConnsPerHost:    10,
	IdleConnTimeout:        90 * time.Second,
	TLSHandshakeTimeout:    10 * time.Second,
	ResponseHeaderTimeout:  10 * time.Second,
	ExpectContinueTimeout:  1 * time.Second,
	ForceAttemptHTTP2:      true,
	DisableCompression:     false,
	DisableKeepAlives:      false,
	MaxConnsPerHost:        0,
	MaxResponseHeaderBytes: 0,
}

// 当前使用的连接池配置
var currentPoolConfig = defaultPoolConfig

// 连接池管理 - 用于 TCP 连接复用
var (
	// 默认 Transport，用于不使用代理的场景
	defaultTransport *http.Transport
	// 默认 Client，用于不使用代理的场景
	defaultClient *http.Client
	// 代理 Transport 缓存，key 为代理 URL
	proxyTransportCache = make(map[string]*http.Transport)
	// 代理 Client 缓存，key 为代理 URL
	proxyClientCache = make(map[string]*http.Client)
	// 缓存锁，用于并发安全
	transportCacheMutex sync.RWMutex
)

// 初始化默认 Transport 和 Client
func initDefaultClient() {
	defaultTransport = createTransport(nil)
	defaultClient = &http.Client{
		Transport: defaultTransport,
		Timeout:   httpTimeout,
	}
}

// 根据配置创建 Transport
func createTransport(proxyURL *url.URL) *http.Transport {
	transport := &http.Transport{
		MaxIdleConns:           currentPoolConfig.MaxIdleConns,
		MaxIdleConnsPerHost:    currentPoolConfig.MaxIdleConnsPerHost,
		IdleConnTimeout:        currentPoolConfig.IdleConnTimeout,
		TLSHandshakeTimeout:    currentPoolConfig.TLSHandshakeTimeout,
		ResponseHeaderTimeout:  currentPoolConfig.ResponseHeaderTimeout,
		ExpectContinueTimeout:  currentPoolConfig.ExpectContinueTimeout,
		ForceAttemptHTTP2:      currentPoolConfig.ForceAttemptHTTP2,
		DisableCompression:     currentPoolConfig.DisableCompression,
		DisableKeepAlives:      currentPoolConfig.DisableKeepAlives,
		MaxConnsPerHost:        currentPoolConfig.MaxConnsPerHost,
		MaxResponseHeaderBytes: currentPoolConfig.MaxResponseHeaderBytes,
	}

	if proxyURL != nil {
		transport.Proxy = http.ProxyURL(proxyURL)
		transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	return transport
}

// 获取或创建代理 Client
func getProxyClient(proxyUrl string) (*http.Client, error) {
	// 先尝试读锁获取缓存
	transportCacheMutex.RLock()
	if client, exists := proxyClientCache[proxyUrl]; exists {
		transportCacheMutex.RUnlock()
		return client, nil
	}
	transportCacheMutex.RUnlock()

	// 缓存不存在，使用写锁创建新的
	transportCacheMutex.Lock()
	defer transportCacheMutex.Unlock()

	// 双重检查，防止并发创建
	if client, exists := proxyClientCache[proxyUrl]; exists {
		return client, nil
	}

	// 解析代理 URL
	url_i := url.URL{}
	proxyURL, err := url_i.Parse(proxyUrl)
	if err != nil {
		return nil, err
	}

	// 使用统一的配置创建 Transport
	transport := createTransport(proxyURL)

	// 创建新的 Client
	client := &http.Client{
		Transport: transport,
		Timeout:   httpTimeout,
	}

	// 缓存
	proxyTransportCache[proxyUrl] = transport
	proxyClientCache[proxyUrl] = client

	return client, nil
}

// SetConnectionPool 设置连接池配置
// 参数：
//   - config: 连接池配置，传入 nil 则重置为默认配置
//
// 示例：
//
//	// 启用连接池并使用默认配置
//	mybinanceapi.SetConnectionPool(nil)
//
//	// 自定义配置
//	config := mybinanceapi.ConnectionPoolConfig{
//		Enable:              true,
//		MaxIdleConns:        200,
//		MaxIdleConnsPerHost: 20,
//		IdleConnTimeout:     120 * time.Second,
//	}
//	mybinanceapi.SetConnectionPool(&config)
//
//	// 禁用连接池（恢复原有每次创建新连接的行为）
//	config := mybinanceapi.ConnectionPoolConfig{Enable: false}
//	mybinanceapi.SetConnectionPool(&config)
func SetConnectionPool(config *ConnectionPoolConfig) {
	transportCacheMutex.Lock()
	defer transportCacheMutex.Unlock()

	// 如果传入 nil，使用默认配置
	if config == nil {
		currentPoolConfig = defaultPoolConfig
	} else {
		// 使用用户提供的配置，但对未设置的值使用默认值
		if config.MaxIdleConns <= 0 {
			config.MaxIdleConns = defaultPoolConfig.MaxIdleConns
		}
		if config.MaxIdleConnsPerHost <= 0 {
			config.MaxIdleConnsPerHost = defaultPoolConfig.MaxIdleConnsPerHost
		}
		if config.IdleConnTimeout <= 0 {
			config.IdleConnTimeout = defaultPoolConfig.IdleConnTimeout
		}
		if config.TLSHandshakeTimeout <= 0 {
			config.TLSHandshakeTimeout = defaultPoolConfig.TLSHandshakeTimeout
		}
		if config.ResponseHeaderTimeout <= 0 {
			config.ResponseHeaderTimeout = defaultPoolConfig.ResponseHeaderTimeout
		}
		if config.ExpectContinueTimeout <= 0 {
			config.ExpectContinueTimeout = defaultPoolConfig.ExpectContinueTimeout
		}

		currentPoolConfig = *config
	}

	// 清理现有的连接池缓存，下次请求时会使用新配置重新创建
	if defaultTransport != nil {
		defaultTransport.CloseIdleConnections()
	}
	for _, transport := range proxyTransportCache {
		if transport != nil {
			transport.CloseIdleConnections()
		}
	}

	// 清空缓存，强制重新创建
	proxyTransportCache = make(map[string]*http.Transport)
	proxyClientCache = make(map[string]*http.Client)

	// 重新初始化默认 Client
	initDefaultClient()
}

// GetConnectionPoolConfig 获取当前的连接池配置（返回副本）
func GetConnectionPoolConfig() ConnectionPoolConfig {
	transportCacheMutex.RLock()
	defer transportCacheMutex.RUnlock()
	return currentPoolConfig
}

// EnableConnectionPool 启用连接池（快捷方法）
func EnableConnectionPool() {
	config := GetConnectionPoolConfig()
	config.Enable = true
	SetConnectionPool(&config)
}

// DisableConnectionPool 禁用连接池（快捷方法）
// 禁用后将恢复原有的每次请求创建新连接的行为
func DisableConnectionPool() {
	config := GetConnectionPoolConfig()
	config.Enable = false
	SetConnectionPool(&config)
}

// CloseAllIdleConnections 关闭所有空闲连接（可选，用于程序退出时优雅关闭）
func CloseAllIdleConnections() {
	transportCacheMutex.Lock()
	defer transportCacheMutex.Unlock()

	// 关闭默认 Transport 的空闲连接
	if defaultTransport != nil {
		defaultTransport.CloseIdleConnections()
	}

	// 关闭所有代理 Transport 的空闲连接
	for _, transport := range proxyTransportCache {
		if transport != nil {
			transport.CloseIdleConnections()
		}
	}
}

func GetCurrentProxyList() []*RestProxy {
	return proxyList
}

func SetUseProxy(useProxy bool, proxyUrls ...string) {
	UseProxy = useProxy
	var newProxyList []*RestProxy
	for _, proxyUrl := range proxyUrls {
		newProxyList = append(newProxyList, &RestProxy{
			ProxyUrl: proxyUrl,
		})
	}
	proxyList = newProxyList
}

func SetWsUseProxy(useProxy bool) error {
	if !UseProxy {
		return errors.New("please set UseProxy first")
	}
	WsUseProxy = useProxy
	return nil
}

func isUseProxy() bool {
	return UseProxy
}

func init() {
	// 初始化默认的 HTTP Client（启用连接复用）
	initDefaultClient()

	c := cron.New(cron.WithSeconds())
	//每分钟0秒权重清零，状态恢复
	_, err := c.AddFunc("0 */1 * * * *", func() {
		for _, proxy := range proxyList {
			proxy.SpotWeight.restore()
			proxy.FutureWeight.restore()
			proxy.SwapWeight.restore()
			proxy.PortfolioMargin.restore()
		}
	})
	if err != nil {
		log.Error(err)
		return
	}
	c.Start()
}

// 获取最佳代理，当前全部代理权重已用完则返回错误
func getBestProxyAndWeight(apiType ApiType) (*RestProxy, *ProxyWeight, error) {

	var minWeigthProxy *RestProxy
	var minWeight *ProxyWeight
	for _, proxy := range proxyList {
		var proxyWeight *ProxyWeight
		switch apiType {
		case SPOT:
			proxyWeight = &proxy.SpotWeight
		case FUTURE:
			proxyWeight = &proxy.FutureWeight
		case SWAP:
			proxyWeight = &proxy.SwapWeight
		case PORTFOLIO_MARGIN:
			proxyWeight = &proxy.PortfolioMargin
		default:
			return nil, nil, errors.New("apiType error")
		}
		if proxyWeight.Is429Limited {
			continue
		}
		if minWeigthProxy == nil {
			minWeigthProxy = proxy
			minWeight = proxyWeight
			continue
		}
		if proxyWeight.UsedWeight < minWeight.UsedWeight {
			minWeigthProxy = proxy
			minWeight = proxyWeight
		}
	}

	return minWeigthProxy, minWeight, nil
}

// 获取随机代理
func getRandomProxy() (*RestProxy, error) {
	length := len(proxyList)
	if length == 0 {
		return nil, errors.New("proxyList is empty")
	}

	return proxyList[rand.Intn(length)], nil
}
func RequestWithHeader(urlStr string, method string, headerMap map[string]string, isGzip bool) ([]byte, error) {
	return RequestWithHeaderAndBody(urlStr, method, headerMap, []byte{}, isGzip)
}

func RequestWithHeaderAndBody(urlStr string, method string, headerMap map[string]string, reqBody []byte, isGzip bool) ([]byte, error) {
	req, err := http.NewRequest(method, urlStr, nil)
	if err != nil {
		return nil, err
	}
	for k, v := range headerMap {
		req.Header.Set(k, v)
	}
	//req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if isGzip { // 请求 header 添加 gzip
		req.Header.Add("Content-Encoding", "gzip")
		req.Header.Add("Accept-Encoding", "gzip")
	}

	log.Debug(method, ": ", req.URL.String())

	var client *http.Client
	var currentProxy *RestProxy
	var currentProxyWeight *ProxyWeight

	// 检查是否启用连接池
	if currentPoolConfig.Enable {
		// 使用连接池的 client（启用 TCP 复用）
		if UseProxy {
			apiType := FUTURE
			maxWeight := 0
			switch req.Host {
			case BinanceGetRestHostByApiType(SPOT):
				apiType = SPOT
				maxWeight = 2400
			case BinanceGetRestHostByApiType(FUTURE):
				apiType = FUTURE
				maxWeight = 2400
			case BinanceGetRestHostByApiType(SWAP):
				apiType = SWAP
				maxWeight = 2400
			case BinanceGetRestHostByApiType(PORTFOLIO_MARGIN):
				apiType = PORTFOLIO_MARGIN
				maxWeight = 2400
			default:
				return nil, errors.New("request host error")
			}

			currentProxy, currentProxyWeight, _ = getBestProxyAndWeight(apiType)
			if currentProxy == nil || currentProxyWeight == nil || currentProxyWeight.UsedWeight >= maxWeight {
				return nil, errors.New("all proxy ip weight limit reached")
			}

			// 从缓存中获取或创建代理 Client（启用连接复用）
			client, err = getProxyClient(currentProxy.ProxyUrl)
			if err != nil {
				return nil, err
			}
		} else {
			// 使用默认 Client（启用连接复用）
			client = defaultClient
		}
	} else {
		// 禁用连接池：每次创建新的 client（原有行为）
		client = &http.Client{
			Timeout: httpTimeout,
		}
		if UseProxy {
			apiType := FUTURE
			maxWeight := 0
			switch req.Host {
			case BinanceGetRestHostByApiType(SPOT):
				apiType = SPOT
				maxWeight = 2400
			case BinanceGetRestHostByApiType(FUTURE):
				apiType = FUTURE
				maxWeight = 2400
			case BinanceGetRestHostByApiType(SWAP):
				apiType = SWAP
				maxWeight = 2400
			case BinanceGetRestHostByApiType(PORTFOLIO_MARGIN):
				apiType = PORTFOLIO_MARGIN
				maxWeight = 2400
			default:
				return nil, errors.New("request host error")
			}

			currentProxy, currentProxyWeight, _ = getBestProxyAndWeight(apiType)
			if currentProxy == nil || currentProxyWeight == nil || currentProxyWeight.UsedWeight >= maxWeight {
				return nil, errors.New("all proxy ip weight limit reached")
			}

			url_i := url.URL{}
			bestProxy, _ := url_i.Parse(currentProxy.ProxyUrl)

			reqProxy := &http.Transport{}
			reqProxy.Proxy = http.ProxyURL(bestProxy)
			reqProxy.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

			client.Transport = reqProxy
		}
	}

	if len(reqBody) > 0 {
		req.Body = io.NopCloser(bytes.NewBuffer(reqBody))
		log.Warn(string(reqBody))
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Error(err)
		}
	}(resp.Body)

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
	log.Debug(resp.Header)
	if isUseProxy() {
		//回填权重
		if resp.Header.Get("X-MBX-USED-WEIGHT-1M") != "" {
			usedWeight, err := strconv.Atoi(resp.Header.Get("X-MBX-USED-WEIGHT-1M"))
			if err != nil {
				log.Error(err)
			}
			if usedWeight > currentProxyWeight.UsedWeight {
				currentProxyWeight.UsedWeight = usedWeight
			}
		}

		//回填是否429限制
		if resp.StatusCode == 429 {
			currentProxyWeight.Is429Limited = true
		}
	}
	return data, err
}
