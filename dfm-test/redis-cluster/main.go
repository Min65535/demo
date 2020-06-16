package main

import (
	"flag"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

func makeRedisCluster(opts *redis.ClusterOptions) *redis.ClusterClient {
	cli := redis.NewClusterClient(opts)
	if err := cli.Ping().Err(); err != nil {
		panic(err)
	}
	return cli
}

func getRedisClusterConfig() *redis.ClusterOptions {
	return &redis.ClusterOptions{
		Addrs: []string{
			//"172.29.58.184:6380",
			//"172.29.58.184:6381",
			//"172.29.59.133:6382",
			//"172.29.59.133:6383",
			//"172.29.59.23:6384",
			//"172.29.59.23:6385",
			"192.168.199.3:6380",
			"192.168.199.3:6381",
			"192.168.199.3:6382",
			"192.168.199.7:6383",
			"192.168.199.7:6384",
			"192.168.199.7:6385",
		},
		MinIdleConns: 10,
		PoolSize:     100,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 3 * time.Second,
	}
}

func main() {
	var (
		set   string
		value string
		get   string
	)
	flag.StringVar(&set, "set", "", "key set")
	flag.StringVar(&value, "value", "", "key value")
	flag.StringVar(&get, "get", "", "get key")
	flag.Parse()
	cli := makeRedisCluster(getRedisClusterConfig())

	switch {
	case get != "":
		getStr := cli.Get(get).Val()
		fmt.Println(`get the value of the key "`+get+`" is:`, getStr)
	case set != "":
		if value == "" {
			fmt.Println("the value is empty")
		}
		setVal, err := cli.Set(set, value, 30*time.Second).Result()
		if err != nil {
			fmt.Println("fail to set the value:", err)
			return
		}
		fmt.Println("setVal:", setVal)
	}
}
