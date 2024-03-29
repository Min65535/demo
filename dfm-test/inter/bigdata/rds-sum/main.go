package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/dipperin/go-ms-toolkit/json"
	"github.com/go-redis/redis/v8"
	"time"
)

var (
	in = flag.String("in", "[111]", "[111]")
)

// addr: 127.0.0.1:6379
// password: 123456
// db: 0
// poolSize: 100
// read_timeout: 10000
// write_timeout: 20000
func main() {
	flag.Parse()
	var arr []uint64
	if err := json.ParseJson(*in, &arr); err != nil {
		fmt.Println("ParseJson err:", err)
		return
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		DB:       0,   // use default DB
		PoolSize: 100, // 连接池大小
		Password: "123456",
	})

	ctx1, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := rdb.Ping(ctx1).Result()

	if err != nil {
		fmt.Println("Failed to open redis with err:", err)
		panic(err)
	}
	ctx := context.Background()
	if len(arr) > 0 {
		for _, uid := range arr {
			re, errZs := rdb.ZScore(ctx, "game_room::status", fmt.Sprintf("%d", uid)).Result()
			switch errZs {
			case redis.Nil:
			case nil:
			default:
				fmt.Println("ZScore err:", err)
				continue
			}
			if re > 0 {
				fmt.Println(uid)
				if e2 := rdb.ZRem(ctx, "game_room::status", fmt.Sprintf("%d", uid)).Err(); e2 != nil {
					fmt.Println("ZRem uid err: ", uid, err)
				}
			}
		}
	}

}
