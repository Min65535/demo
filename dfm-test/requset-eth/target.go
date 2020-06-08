package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func HttpReq(r *http.Request) ([]byte, error) {
	httpClient := &http.Client{}
	resp, err := httpClient.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		return nil, err2
	}
	return b, nil
}

func main() {
	//r :=resty.New()
	//r.HostURL= ""
	var (
		host   string
		scheme string
		num    int
	)
	flag.StringVar(&host, "host", "", "host")
	flag.StringVar(&scheme, "scheme", "", "scheme")
	flag.IntVar(&num, "num", 1, "num")
	flag.Parse()
	if host != "" && scheme != "" {
		r := &http.Request{
			URL:    &url.URL{Host: host, Scheme: scheme},
			Method: "GET",
		}

		for i := 0; i < num; i++ {
			resp, err := HttpReq(r)
			if err != nil {
				fmt.Println(i+1, "err:", err)
				return
			}
			fmt.Println(i+1, "resp:", string(resp))
		}
	}
}
