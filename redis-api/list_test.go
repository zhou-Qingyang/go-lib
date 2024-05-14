package main

import (
	"testing"
)

func TestList(t *testing.T) {
	// 01. LPush指令实现 返回影响行数
	// for _, v := range []int{1, 2, 3} {
	// 	_, err := redisClient.LPush(context.TODO(), "redis_stack", v).Result()
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }

	// 02. LPop指令实现 拿到存放的字符串类型的数据
	// for i := 0; i < 3; i++ {
	// 	res, err := redisClient.LPop(context.TODO(), "redis_stack").Result()
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println(res)
	// }

	// 入栈操作
	// 03. RPop指令实现
	// for _, v := range []int{1, 2, 3} {
	// 	_, err := redisClient.RPush(context.TODO(), "redis_queue", v).Result()
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }
	// 出栈操作
	// 04. LPop
	// for i := 0; i < 3; i++ {
	// 	res, err := redisClient.LPop(context.TODO(), "redis_queue").Result()
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println(res)
	// }

}
