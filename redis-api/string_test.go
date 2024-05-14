package main

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	redisClient *redis.Client
	once        sync.Once
)

type Data struct {
	UserName string `json:"username"`
	Colors   []int  `json:"colors"`
}

// GetredisClient 返回 Redis 客户端实例
func GetRedisClient() *redis.Client {
	once.Do(func() {
		redisClient = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})
		pong, err := redisClient.Ping(context.Background()).Result()
		fmt.Println("redis connect ping result:", pong)
		if err != nil {
			fmt.Println("redis connect ping failed, err:", err)
		}
	})
	return redisClient
}
func TestStr(t *testing.T) {
	GetRedisClient()
	ctx := context.Background()
	data := Data{
		UserName: "nihao",
		Colors:   []int{1, 2, 3},
	}
	dataBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println("json marshal failed, err:", err)
	}
	// 01.Set指令 key value 过期时间
	err = redisClient.Set(ctx, "key", dataBytes, 0).Err()
	if err != nil {
		panic(err)
	}
	// 02. Get指令 key 返回的是 *StringCmd  Result方法返回的是String
	val, err := redisClient.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}

	if val == "" {
		fmt.Println("key not exist")
		return
	}
	var res Data
	if err := json.Unmarshal([]byte(val), &res); err != nil {
		fmt.Println("json unmarshal failed, err:", err)
	}
	fmt.Printf("%+v\n", res)

	// 03. Del指令 key
	// err = redisClient.Del(ctx, "key").Err()
	// if err != nil {
	// 	panic(err)
	// }

	//04. 将key 的值加一或者减一 只能是整数
	err = redisClient.Set(ctx, "key_int", "1", 0).Err()
	if err != nil {
		panic(err)
	}
	// 如果没有key_int 自动创建 key_int并自动加1
	newVal, err := redisClient.Incr(ctx, "key_int").Result()
	// newVal, err = redisClient.Decr(ctx, "key_int").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("newVal:", newVal)

	// 05. SetEX指令 key value 过期时间
	err = redisClient.SetEX(ctx, "key1", "dada", 10*time.Second).Err()
	if err != nil {
		panic(err)
	}
	// 06. SetNX 分布式锁
	result, err := redisClient.SetNX(ctx, "lockKey", 1, 0).Result()
	if err != nil {
		panic(err)
	}
	if result {
		fmt.Println("setNX success")
	} else {
		fmt.Println("setNX failed")
	}

	// Lua 脚本用于原子性地检查锁的值并删除锁
	script := `
        if redis.call("GET", KEYS[1]) == ARGV[1] then
            return redis.call("DEL", KEYS[1])
        else
            return 0
        end
    `
	// 指定返回值的类型为 int64
	req, err := redisClient.Eval(ctx, script, []string{"lockKey"}, "1").Result()
	if err != nil {
		panic(err)
	}
	if req == int64(1) {
		fmt.Println("delete lock success")
	} else {
		fmt.Println("delete lock failed")
	}
}
