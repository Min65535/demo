package main

import (
	"context"
	"fmt"
	db_config "github.com/dipperin/go-ms-toolkit/db-config"
	"github.com/dipperin/go-ms-toolkit/qyenv"
	"github.com/go-redis/redis/v8"
	"github.com/min65535/demo/dfm-test/inter/rds/locker"
	"time"
)

func GetRedisConfig() *redis.Options {
	var db int
	runEnv := qyenv.GetDBEnv()
	switch runEnv {
	case "prod":
		db = 9
	case "preprod":
		db = 3
	case "test":
		db = 2
	default:
		db = 1
	}
	return &redis.Options{
		Addr: db_config.GetAppConfig().RedisUrl,
		DB:   db,
	}
}

func MakeRedis(opts *redis.Options) *redis.Client {
	ctx := context.Background()
	cli := redis.NewClient(opts)
	if err := cli.Ping(ctx).Err(); err != nil {
		panic(err)
	}
	return cli
}

func main() {

	rd := MakeRedis(GetRedisConfig())
	defer rd.Close()
	//
	// sign, err := rd.SetNX("dfm-sss", true, time.Second*30).Result()
	// if err != nil {
	// 	fmt.Println("err:", err.Error())
	// 	// return
	// }
	// fmt.Println("sign:", sign)
	ctx := context.Background()
	woLocker := locker.New(rd, locker.WithExpiration(30*time.Second), locker.WithPrefix("WOrderLocker_trade"), locker.WithMaxSpin(1))
	err2 := woLocker.Lock(ctx, "dfm-ss")
	if err2 != nil {
		fmt.Println("err2:", err2.Error())
		return
	}
	defer woLocker.Unlock(ctx, "dfm-ss")
	rd.Del(ctx, "dfm-ss")

}
