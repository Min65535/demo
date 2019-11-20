package main

import (
	"github.com/elastic/go-elasticsearch/v7"
	"log"
	"github.com/dipperin/go-ms-toolkit/json"
	"net/http"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7/estransport"
	"os"
	"io/ioutil"
	"strings"
)

// ExtendedClient allows to call regular and Execute APIs.
//
type ExtendedClient struct {
	*elasticsearch.Client
}

func main() {
	//SearchRequest
	var c elasticsearch.Config
	c.Addresses = []string{"http://mygitlab:9200"}
	c.Logger = &estransport.ColorLogger{Output: os.Stdout, EnableRequestBody: true, EnableResponseBody: true}
	esClient, _ := elasticsearch.NewClient(c)
	resInfo, _ := esClient.Info()
	fmt.Println("resInfo:", resInfo)
	log.Println(elasticsearch.Version)
	esDao := ExtendedClient{Client: esClient}

	//req, _ := http.NewRequest("GET", "/_search", nil)
	reqStr := `{
		"query":{
			"match":{
				"order_no":"Qr20190723000001"
			}
		}
	}`
	req, _ := http.NewRequest("POST", "/approval_order_ind/duplicate_feature_data_phl/_search", strings.NewReader(reqStr))
	req.Header.Set("Content-Type", "application/json")
	res, err := esDao.Perform(req)
	if err != nil {
		panic(err)
	}

	fmt.Println("res status:", json.StringifyJson(res.Status), "head:", json.StringifyJson(res.Header))

	bodyRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("bodyRes:", string(bodyRes))

}
