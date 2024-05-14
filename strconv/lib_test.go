package main

import (
	"strconv"
	"testing"
)

func TestXx(t *testing.T) {
	// 字符串转 int类型
	// func ParseInt(s string, base int, bitSize int) (i int64, err error)
	// func ParseUint(s string, base int, bitSize int) (n uint64, err error)
	// func Atoi(s string) (i int, err error)
	//ParseInt(s, 10, 0)
	strconv.ParseInt("123", 10, 64) //strconv.Atoi 的底层

	// int类型转字符串
	strconv.Itoa(123)
	s := strconv.FormatInt(123, 10) // 有符号整型转字符串
	//参数 bitSize 表示的是整数取值范围，或者说整数的具体类型。
	// 取值 0、8、16、32 和 64 分别代表 int、int8、int16、int32 和 int64。
	t.Log(s)
	//utf-8 和 utf-16 表示最小编码的字节长度 但是英文都非常普及 就没有utf-8那么推广开来
}
