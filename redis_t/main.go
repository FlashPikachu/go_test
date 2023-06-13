package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func main() {
	ctx := context.Background()
	err := rdb.Set(ctx, "goredistestkey", "goredistestvalue", 0).Err()
	if err != nil {
		panic(any(err))
	}
	result, err := rdb.Get(ctx, "goredistestkey").Result()
	if err != nil {
		panic(any(err))
	}
	println("what the fkk")
	println(result)

	ret, err := rdb.Do(ctx, "get", "goredistestkey").Result()
	if err != nil {
		panic(any(err))
	}
	fmt.Println(ret.(string))
}

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "47.110.240.147:63792",
		Password: "valhalla", // 没有密码，默认值
		DB:       0,          // 默认DB 0
	})
}
