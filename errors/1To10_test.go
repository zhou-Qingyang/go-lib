package main

import (
	"fmt"
	"sync"
	"testing"
)

func Test1(t *testing.T) {
	//01 返回值屏蔽
	//02 context静止包问题
	//03 禁止退出
	//04 map无序集合
	//05 main函数提前退出问题
	//06 循环依赖问题
	//07 switch 默认执行完case 直接跳出 fallthrough则强制执行 case ' ', '\t': 或者带上多个条件
	//08 海豹符号只能函数内部使用
	//09 使用断言前先加上判断
	//10 goRoutine 异常处理 导致主线程无法生效
	//11 Slice作为参数传递是值传递内容相同
	//12 omitempty json 序列化的时候 为空或者0的字段直接不序列化
	//13 控制并发线程 waitGroup channel waitGroup + channel 线程池
	//14 无法给显示变量使用 nil赋值 var x = nil 只有 interface map channel slice function pointer
	//15 map 和 slice 初始化 make默认搭配map使用 map nil 使用直接panic
	//16 结构体值的比较 == 来比较结构体变量，前提是两个结构体的成员都是可比  但是如果结构体中存在 map或者slice这种 无法比较
	//17 for range 迭代切片的时候 要根据索引修改值 不然不会修改原来的值 因为迭代出来的是值拷贝
	//18 slice的拷贝 如果使用的是 slice[:3]这种方式 仍然会保留之前slice的容量 所以通常使用copy
	//19 interface{} json反序列化出来是一个 float64

	// t.Log(Hello(""))
	// t.Log(Hello("11"))
	// InterfaceAssert()
	// Has()
	// SliceTest()
}

func SliceTest() {
	tmp := make([]int, 0, 5)
	fmt.Printf("%p\n", &tmp)
	fmt.Printf("%v %d %d %p\n", tmp, len(tmp), cap(tmp), tmp)
	change(tmp)
	fmt.Printf("%v %d %d %p\n", tmp, len(tmp), cap(tmp), tmp)
}
func change(tmp []int) {
	fmt.Printf("%p\n", &tmp)
	// 发生扩容了
	tmp = append(tmp, 6)
	fmt.Printf("%v %d %d %p\n", tmp, len(tmp), cap(tmp), tmp)
}

// 01 返回屏蔽  技能
func Hello(name string) (str string) {
	str = "测试"
	if name != "" {
		return name + "hello"
	}
	return
}

// 03 禁止退出
func ForEver() {
	//01. select{}

	// 02. defer func() {
	// 	for {
	// 	}
	// }()

	// 03. defer func() {
	// 	select {}
	// }()
}

// 09 使用断言前先加上判断
func InterfaceAssert() {
	var v interface{}
	v = 1
	switch v.(type) {
	case string:
		fmt.Printf("value:%s \n", v.(string))
	case int:
		fmt.Printf("value:%d \n", v.(int))
	case int32:
		fmt.Printf("value:%d \n", v.(int32))
	}
}

func Has() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("err:", err)
			}
			wg.Done()
		}()
		fmt.Println("Hello word")
		panic("routine errors")
	}()
	wg.Wait()
	fmt.Println("word")
}
