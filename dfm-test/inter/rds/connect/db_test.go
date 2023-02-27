package connect

import (
	"context"
	"fmt"
	"testing"
)

func TestMy(t *testing.T) {
	rd := MakeRedis(GetRedisConfig())
	ctx := context.Background()
	defer rd.Close()
	if err := rd.LPush(ctx, "mykkk", 111).Err(); err != nil {
		fmt.Println("LPush1 err:", err)
	}
	if err := rd.LPush(ctx, "mykkk", 111).Err(); err != nil {
		fmt.Println("LPush2 err:", err)
	}
	if err := rd.LPush(ctx, "mykkk", 222).Err(); err != nil {
		fmt.Println("LPush3 err:", err)
	}

	if le, err := rd.LLen(ctx, "mykkk").Result(); err != nil {
		fmt.Println("LLen err:", err)
	} else {
		fmt.Println("le:", le)
	}
	if data, err := rd.LRange(ctx, "mykkk", 0, 22).Result(); err != nil {
		fmt.Println("LRange err:", err)
	} else {
		fmt.Println("data:", data)
	}

	if rei, err := rd.LRem(ctx, "mykkk", 1, 222).Result(); err != nil {
		fmt.Println("lrem1 err:", err)
	} else {
		fmt.Println("rei:", rei)
	}
	if err := rd.LRem(ctx, "mykkk", 1, 222).Err(); err != nil {
		fmt.Println("lrem2 err:", err)
	}
	if rei2, err := rd.LRem(ctx, "mykkk", 1, 222).Result(); err != nil {
		fmt.Println("lrem3 err:", err)
	} else {
		fmt.Println("rei2:", rei2)
	}

	if err := rd.Del(ctx, "mykkk").Err(); err != nil {
		fmt.Println("del err:", err)
	}
}
