package mybinanceapi

import (
	"compress/gzip"
	"crypto/tls"
	"errors"
	"github.com/robfig/cron/v3"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type RestProxy struct {
	ProxyUrl string //代理的协议IP端口URL

	SpotWeight   ProxyWeight
	FutureWeight ProxyWeight
	SwapWeight   ProxyWeight
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

func isUseProxy() bool {
	return UseProxy
}

func init() {
	c := cron.New(cron.WithSeconds())
	//每分钟0秒权重清零，状态恢复
	_, err := c.AddFunc("0 */1 * * * *", func() {
		for _, proxy := range proxyList {
			proxy.SpotWeight.restore()
			proxy.FutureWeight.restore()
			proxy.SwapWeight.restore()
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

func RequestWithHeader(urlStr string, method string, headerMap map[string]string, isGzip bool) ([]byte, error) {
	req, err := http.NewRequest(method, urlStr, nil)
	if err != nil {
		return nil, err
	}
	for k, v := range headerMap {
		req.Header.Set(k, v)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{
		Timeout: httpTimeout,
	}
	if isGzip { // 请求 header 添加 gzip
		req.Header.Add("Content-Encoding", "gzip")
		req.Header.Add("Accept-Encoding", "gzip")
	}

	log.Debug(method, ": ", req.URL.String())

	var currentProxy *RestProxy
	var currentProxyWeight *ProxyWeight
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
		default:
			return nil, errors.New("request host error")
		}

		currentProxy, currentProxyWeight, _ = getBestProxyAndWeight(apiType)
		if currentProxy == nil || currentProxyWeight.UsedWeight >= maxWeight {
			currentProxyWeight.Is429Limited = true
			return nil, errors.New("all proxy ip weight limit reached")
		}

		url_i := url.URL{}
		bestProxy, _ := url_i.Parse(currentProxy.ProxyUrl)

		reqProxy := &http.Transport{}
		reqProxy.Proxy = http.ProxyURL(bestProxy)                        // set proxy
		reqProxy.TLSClientConfig = &tls.Config{InsecureSkipVerify: true} //set ssl

		client.Transport = reqProxy
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
