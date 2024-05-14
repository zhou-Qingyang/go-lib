package main

import (
	"context"
	"fmt"
	"testing"
)

func TestBitMap(t *testing.T) {
	GetRedisClient()
	// 设置用户ID和日期
	userID := "user123"
	today := "2024-05-15" // 假设今天是 2024-05-14
	// 将用户签到日期对应的位设置为 1
	err := redisClient.SetBit(context.Background(), fmt.Sprintf("signin:%s:%s", userID, today), 0, 1).Err()
	if err != nil {
		panic(err)
	}

	// 检查用户在指定日期是否已经签到
	isSigned, err := redisClient.GetBit(context.Background(), fmt.Sprintf("signin:%s:%s", userID, today), 0).Result()
	if err != nil {
		panic(err)
	}
	if isSigned == 1 {
		fmt.Printf("用户 %s 在 %s 已签到\n", userID, today)
	} else {
		fmt.Printf("用户 %s 在 %s 未签到\n", userID, today)
	}
}
