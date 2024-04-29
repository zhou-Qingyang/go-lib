package main

import (
	"fmt"
	"testing"
)

func TestY(t *testing.T) {
	type User struct {
		Name string
	}
	u := User{Name: "tang"}
	//Printf 格式化输出
	fmt.Printf("%+v\n", u)   //格式化输出结构 结构体名
	fmt.Printf("%v\n", u)    //输出值的 Go 语言表示方法
	fmt.Printf("%#v\n", u)   //输出值的 Go 语言表示方法 包名加上结构体名
	fmt.Printf("%T\n", u)    //输出值的类型的 Go 语言表示
	fmt.Printf("%t\n", true) //输出值的 true 或 false
	fmt.Printf("%b\n", 1024) //二进制表示
	fmt.Printf("%c\n", 97)   //数值对应的 Unicode 编码字符
	//Unicode 是一种字符集 可能世界上绝大多数使用字符
	//UTF-8 变成编码方式 使用 1-4个字节
}
