package main

import (
	"flag"
	"fmt"
	"github.com/dipperin/go-ms-toolkit/qyenv"
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

var clusterAddrArray []string

func GetRedisClusterAddrArray() []string {
	if len(clusterAddrArray) != 0 {
		return clusterAddrArray
	}
	useDocker := qyenv.GetUseDocker()
	if useDocker == 1 {
		fmt.Println("采用的是dev容器的配置")
		var array = []string{
			// "redis-cluster-node-1.default.svc.cluster.local:6379",
			// "redis-cluster-node-2.default.svc.cluster.local:6379",
			// "redis-cluster-node-3.default.svc.cluster.local:6379",
			// "redis-cluster-node-4.default.svc.cluster.local:6379",
			// "redis-cluster-node-5.default.svc.cluster.local:6379",
			// "redis-cluster-node-6.default.svc.cluster.local:6379",
			"redis-cluster-service.default.svc.cluster.local:6379",
		}
		clusterAddrArray = append(clusterAddrArray, array...)
	} else if useDocker == 2 {
		fmt.Println("采用的是prod生产的配置")
		var array = []string{
			"redis-cluster-node-1.default.svc.cluster.local:6379",
			"redis-cluster-node-2.default.svc.cluster.local:6379",
			"redis-cluster-node-3.default.svc.cluster.local:6379",
			"redis-cluster-node-4.default.svc.cluster.local:6379",
			"redis-cluster-node-5.default.svc.cluster.local:6379",
			"redis-cluster-node-6.default.svc.cluster.local:6379",
		}
		clusterAddrArray = append(clusterAddrArray, array...)
	} else {
		fmt.Println("采用的是非docker环境的配置,请自己组建集群")
		// clusterAddrArray = append(clusterAddrArray, "127.0.0.1:6379")
	}
	return clusterAddrArray
}

func GetRedisClusterConfig() *redis.ClusterOptions {
	return &redis.ClusterOptions{
		// Addrs: []string{
		//	//"172.29.58.184:6380",
		//	//"172.29.58.184:6381",
		//	//"172.29.59.133:6382",
		//	//"172.29.59.133:6383",
		//	//"172.29.59.23:6384",
		//	//"172.29.59.23:6385",
		//	//"192.168.199.3:6380",
		//	//"192.168.199.3:6381",
		//	//"192.168.199.3:6382",
		//	//"192.168.199.7:6383",
		//	//"192.168.199.7:6384",
		//	//"192.168.199.7:6385",
		//
		//	"redis-cluster-node-1.default.svc.cluster.local:6380",
		//	"redis-cluster-node-2.default.svc.cluster.local:6382",
		//	"redis-cluster-node-3.default.svc.cluster.local:6384",
		//	"redis-cluster-node-4.default.svc.cluster.local:6381",
		//	"redis-cluster-node-5.default.svc.cluster.local:6383",
		//	"redis-cluster-node-6.default.svc.cluster.local:6385",
		// },
		Addrs:        GetRedisClusterAddrArray(),
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
	cli := makeRedisCluster(GetRedisClusterConfig())

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
