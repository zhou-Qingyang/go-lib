package main

import (
	"context"
	"fmt"
	"testing"
)

func TestHyper(t *testing.T) {
	GetRedisClient()
	// 定义两个 HyperLogLog 键名
	hllKey1 := "online_ips_1"
	hllKey2 := "online_ips_2"

	// 模拟两个 HyperLogLog 键的统计结果，这里只是为了示例，实际中可以从 Redis 中获取
	// 注意：合并操作不会修改原始键，所以这里直接将统计结果保存到新的键中
	for _, ip := range []string{"192.168.0.1", "192.168.0.2", "192.168.0.3"} {
		err := redisClient.PFAdd(context.Background(), hllKey1, ip).Err()
		if err != nil {
			panic(err)
		}
	}

	for _, ip := range []string{"192.168.0.3", "192.168.0.4", "192.168.0.5"} {
		err := redisClient.PFAdd(context.Background(), hllKey2, ip).Err()
		if err != nil {
			panic(err)
		}
	}

	// 合并两个 HyperLogLog 键的统计结果到一个新的 HyperLogLog 键中
	err := redisClient.PFMerge(context.Background(), "merged_online_ips", hllKey1, hllKey2).Err()
	if err != nil {
		panic(err)
	}

	// 获取合并后的 HyperLogLog 键中统计的不同 IP 数量（估算值）
	count, err := redisClient.PFCount(context.Background(), "merged_online_ips").Result()
	if err != nil {
		panic(err)
	}

	fmt.Printf("合并后的在线 IP 数量（估算值）：%d\n", count)
}
