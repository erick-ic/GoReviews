package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {
	// 1. 创建客户端（连接本地 Redis，默认 DB 0）
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis 地址
		Password: "",               // 没有密码就留空
		DB:       0,                // 使用默认数据库 0
	})

	// 2. 使用 context（推荐始终传递）
	ctx := context.Background()

	// 3. 写入数据（String 类型）
	err := rdb.Set(ctx, "name", "Alice", 0).Err() // 0 表示永不过期
	if err != nil {
		panic(err)
	}

	// 4. 读取数据
	val, err := rdb.Get(ctx, "name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("name:", val) // 输出 Alice

	// 5. 删除键
	delCount, err := rdb.Del(ctx, "name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("删除键数量:", delCount)
}
