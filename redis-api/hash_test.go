package main

import (
	"context"
	"fmt"
	"testing"
)

func TestHash(t *testing.T) {
	GetRedisClient()
	ctx := context.Background()
	// 01 HashSet
	// HSet("myhash", "key1", "value1", "key2", "value2")
	// HSet("myhash", []string{"key1", "value1", "key2", "value2"})
	// HSet("myhash", map[string]interface{}{"key1": "value1", "key2": "value2"})
	// affectRows, err := redisClient.HSet(ctx, "hash_test", "key1", "value1", "key2", "value2").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("影响行数:", affectRows)

	// 02 HashGet
	// value, err := redisClient.HGet(ctx, "hash_test", "key1").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("key1的值:", value)

	// 03 HashGetAll
	// values, err := redisClient.HGetAll(ctx, "hash_test").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// for k, v := range values {
	// 	fmt.Printf("key:%s value:%s \n", k, v)
	// }

	// 04 HashExists
	// redisClient.HExists(ctx, "hash_test", "key1")

	// 05 HashDel
	// rowAffected, err := redisClient.HDel(ctx, "hash_test", "key1").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("影响行数:", rowAffected)

	// 06 HKeys
	strSlice, err := redisClient.HKeys(ctx, "hash_test").Result()
	if err != nil {
		panic(err)
	}
	for _, v := range strSlice {
		fmt.Println(v)
	}

	// 07 HVals
	strSlice1, err := redisClient.HVals(ctx, "hash_test").Result()
	if err != nil {
		panic(err)
	}
	for _, v := range strSlice1 {
		fmt.Println(v)
	}
}
