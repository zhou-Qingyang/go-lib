package main

import (
	"context"
	"fmt"
	"testing"
)

func TestSet(t *testing.T) {
	GetRedisClient()
	ctx := context.Background()

	// 01. SAdd 添加元素到集合中
	err := redisClient.SAdd(ctx, "myset", "apple", "banana", "cherry").Err()
	if err != nil {
		panic(err)
	}

	// 02. SIsMember检查集合中是否存在某个元素
	exists, err := redisClient.SIsMember(ctx, "myset", "banana").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Is banana member of myset:", exists)

	// 03. SMembers获取集合中的所有成员
	members, err := redisClient.SMembers(ctx, "myset").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Members of myset:", members)

	// 04. SRem删除集合中的元素
	err = redisClient.SRem(ctx, "myset", "cherry").Err()
	if err != nil {
		panic(err)
	}
}
