package main

import (
	"fmt"
	"github.com/dipperin/go-ms-toolkit/json"
	"testing"
)

func TestMY(t *testing.T) {
	w := 2 | 4 | 8 | 64
	s := 16 | 32
	fmt.Println(w & s)
	fmt.Println(w&s != 0)

	// var m map[string]interface{}
	var str = `{
  "service_fee_rate": 0.005,
  "usdt_user_transfer_service_rate": 0.01,
  "usdt_user_transfer_minimum_service_fee": 1,
  "usdt_user_transfer_one_time_limit_min": 100,
  "usdt_withdrawal_service_rate": 0.03,
  "usdt_withdrawal_minimum_service_fee": 5,
  "ios_app_version": "1.0.3",
  "sgw_api_key": "UH6TGI21Z59OFL0MVXCS7NW8DQY3AJKEBRP4",
  "android_app_information": {
    "version_code": "56",
    "build_version": "1.5.1",
    "force": "true",
    "build_update_description": "1.修复已知问题,优化用户体验",
    "downloadURL": "http://xg.apk.n00fz7jmqs2w81wi1ml.ppiaas.com/app-lin-prod-release-201118-56.apk"
  }
}`
	tmp := make(map[string]interface{})
	if err := json.ParseJson(str, &tmp); err != nil {
		fmt.Println("app config parse json fail", err.Error())
		return
	}
	key := "sgw_api_key"
	result, ok := tmp[key]
	if !ok {
		fmt.Println("AppConfigValue##key not found", key, str)
		return
	}
	var apiKey string
	apiKey = result.(string)

	fmt.Println("api:", apiKey)

}
