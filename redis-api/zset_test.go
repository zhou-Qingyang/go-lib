package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/go-redis/redis/v8"
)

func TestZSet(t *testing.T) {
	GetRedisClient()
	// 向 Zset 集合中添加元素
	err := redisClient.ZAdd(context.Background(), "myzset", &redis.Z{Score: 1, Member: "member1"}, &redis.Z{Score: 2, Member: "member2"}).Err()
	if err != nil {
		panic(err)
	}

	// 获取 Zset 集合的所有元素
	members, err := redisClient.ZRange(context.Background(), "myzset", 0, -1).Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("Zset 集合中的元素：", members)

	// 获取 Zset 集合中指定范围内的元素
	membersWithScores, err := redisClient.ZRangeWithScores(context.Background(), "myzset", 0, -1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Zset 集合中的元素及其分数：")
	for _, member := range membersWithScores {
		fmt.Println(member.Member, "-", member.Score)
	}

	// 检查元素是否存在于 Zset 集合中
	exists, err := redisClient.ZRank(context.Background(), "myzset", "member1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("元素 member1 在 Zset 集合中的排名：", exists)
}
