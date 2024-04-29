package main

import (
	"bufio"
	"bytes"
	"os"
)

func main() {
	println("请输入11个字节的数据：")
	scanner := bufio.NewScanner(os.Stdin)
	// 获取用户输入的数据
	input := scanner.Bytes()
	// 截取前 11 个字节的数据
	var data []byte
	if scanner.Scan() {
		// 获取用户输入的数据
		input = scanner.Bytes()
		data = input[:11]
	}
	reader := bytes.NewReader(data)
	// 读取 11 个字节的数据
	target := make([]byte, 11)
	_, err := reader.Read(target)
	if err != nil {
		panic(err)
	}
	// 打印读取到的数据
	println("输入的内容:\n", string(target))
}
