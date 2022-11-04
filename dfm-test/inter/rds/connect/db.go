package connect

import (
	"context"
	db_config "github.com/dipperin/go-ms-toolkit/db-config"
	"github.com/dipperin/go-ms-toolkit/qyenv"
	"github.com/go-redis/redis/v8"
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
