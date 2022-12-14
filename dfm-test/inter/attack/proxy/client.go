package proxy

import (
	"crypto/tls"
	"fmt"
	"github.com/dipperin/go-ms-toolkit/json"
	"github.com/go-resty/resty/v2"
	"math/rand"
	"net/http"
	"time"
)

var AgentList = []string{
	"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:79.0) Gecko/20100101 Firefox/79.0",
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.105 Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/35.0.1916.153 Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.1; WOW64; rv:30.0) Gecko/20100101 Firefox/30.0",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_2) AppleWebKit/537.75.14 (KHTML, like Gecko) Version/7.0.3 Safari/537.75.14",
	"Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; Win64; x64; Trident/6.0)",
	"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; .NET CLR 1.1.4322; .NET CLR 2.0.50727)",
	"Mozilla/5.0 (compatible; Konqueror/3.5; Linux) KHTML/3.5.5 (like Gecko) (Kubuntu)",
	"Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.8.0.12) Gecko/20070731 Ubuntu/dapper-security Firefox/1.5.0.12",
	"Lynx/2.8.5rel.1 libwww-FM/2.14 SSL-MM/1.4.1 GNUTLS/1.2.9",
	"Mozilla/5.0 (X11; Linux i686) AppleWebKit/535.7 (KHTML, like Gecko) Ubuntu/11.04 Chromium/16.0.912.77 Chrome/16.0.912.77 Safari/535.7",
	"Mozilla/5.0 (X11; Ubuntu; Linux i686; rv:10.0) Gecko/20100101 Firefox/10.0",
	"Mozilla/5.0 (Windows; U; Windows NT 5.1; it; rv:1.8.1.11) Gecko/20071127 Firefox/2.0.0.11",
	"Opera/9.25 (Windows NT 5.1; U; en)",
	"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3947.100 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.67 Safari/537.36 Edg/87.0.664.55",
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.66 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.67 Safari/537.36",
	"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:83.0) Gecko/20100101 Firefox/83.0",
}

var (
	UrlList []string
)

type Proxy struct {
	url    string
	client *resty.Client
}

func NewProxy(url string) *Proxy {
	InitProxyUrl()
	r := &Proxy{url: url}
	r.client = resty.New()
	r.client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	return r
}

func InitProxyUrl() {
	UrlList = make([]string, 0)
	// data := `["10.8.0.13:6666","10.8.0.13:6667","10.8.0.13:6668","10.8.0.13:6669","10.8.0.13:6670",""]`
	data := `[]`
	if err := json.ParseJson(data, &UrlList); err != nil {
		return
	}
}

func (r *Proxy) Get(api string, query string, respData interface{}) error {
	agentIndex := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(AgentList))

	sleep := rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(1000) + 1
	// 限速
	time.Sleep(time.Duration(sleep) * time.Millisecond)
	var proUrl string
	if len(UrlList) > 0 {
		urlIndex := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(UrlList))
		proUrl = UrlList[urlIndex]
	}
	rct := r.client
	if proUrl != "" {
		rct = rct.SetProxy(proUrl)
	}
	resp, err := rct.R().
		SetHeader("accept", "application/json, text/plain, */*").
		SetHeader("User-Agent", AgentList[agentIndex]).
		SetQueryString(query).
		Get(r.url + api)

	if err != nil {
		return err
	}

	if err = json.ParseJsonFromBytes(resp.Body(), respData); err != nil {
		return err
	}

	switch resp.StatusCode() {
	case http.StatusOK, http.StatusBadRequest, http.StatusForbidden:
		return nil
	default:
		return fmt.Errorf("request api: %s get resp status code: %d, raw resp: %s", api, resp.StatusCode(), string(resp.Body()))
	}
}

func (r *Proxy) Post(api string, req interface{}, respData interface{}) error {
	agentIndex := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(AgentList))
	urlIndex := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(UrlList))
	sleep := rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(1000) + 1
	// 限速
	time.Sleep(time.Duration(sleep) * time.Millisecond)
	proUrl := UrlList[urlIndex]
	rct := r.client
	if proUrl != "" {
		rct = rct.SetProxy(proUrl)
	}
	resp, err := rct.R().
		SetHeader("accept", "application/json, text/plain, */*").
		SetHeader("User-Agent", AgentList[agentIndex]).
		SetBody(json.StringifyJsonToBytes(req)).
		Post(r.url + api)

	if err != nil {
		return err
	}

	if err = json.ParseJsonFromBytes(resp.Body(), respData); err != nil {
		return err
	}

	switch resp.StatusCode() {
	case http.StatusOK, http.StatusBadRequest, http.StatusForbidden:
		return nil
	default:
		return fmt.Errorf("request api: %s get resp status code: %d, raw resp: %s", api, resp.StatusCode(), string(resp.Body()))
	}
}
