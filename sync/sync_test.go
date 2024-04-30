package main

import (
	"sync"
	"testing"
	"time"
)

var m *sync.RWMutex
var val = 0

func TestXxx(t *testing.T) {
	// 进程是非常轻量级的 goroutine 占用资源不多
	// 但是如果要保证顺序执行，需要使用 channel
	// 同时可以使用 共享内存 锁的方式 来实现协同执行

	// m = new(sync.RWMutex)
	// go read(1)
	// go write(2)
	// go read(3)
	// time.Sleep(5 * time.Second)

	// 01 锁懂了  读取不加锁可能 写锁先执行
	// 02 临时对象池 不懂
	// 03 Once  多个协程同时启动 但是必须只要执行一个方法
	// 04 WaitGroup And Cond跟唤醒线程有关系
	// var once sync.Once
	// ch := make(chan bool)
	// onceBody := func() {
	// 	fmt.Println("Only once")
	// }
	// for i := 0; i < 10; i++ {
	// 	go func(d int) {
	// 		once.Do(onceBody)
	// 		fmt.Println(d)
	// 	}(i)
	// }
	// <-ch

}

func read(i int) {
	m.RLock()
	time.Sleep(1 * time.Second)
	println("val: ", val)
	time.Sleep(1 * time.Second)
	m.RUnlock()
}

func write(i int) {
	m.Lock()
	val = 10
	time.Sleep(1 * time.Second)
	m.Unlock()
}
