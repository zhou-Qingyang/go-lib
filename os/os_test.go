package main

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestXxx(t *testing.T) {
	//const (
	// 	O_RDONLY int = syscall.O_RDONLY // 只读模式打开文件
	// 	O_WRONLY int = syscall.O_WRONLY // 只写模式打开文件
	// 	O_RDWR   int = syscall.O_RDWR   // 读写模式打开文件
	// 	O_APPEND int = syscall.O_APPEND // 写操作时将数据附加到文件尾部
	// 	O_CREATE int = syscall.O_CREAT  // 如果不存在将创建一个新文件
	// 	O_EXCL   int = syscall.O_EXCL   // 和 O_CREATE 配合使用，文件必须不存在
	// 	O_SYNC   int = syscall.O_SYNC   // 打开文件用于同步 I/O
	// 	O_TRUNC  int = syscall.O_TRUNC  // 如果可能，打开时清空文件
	// )
	file, err := os.OpenFile("1.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC|os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("打开文件错误:", err)
		return
	}
	defer file.Close() // 在函数结束时关闭文件

	// 写入数据到文件
	_, err = file.WriteString("dada1")
	if err != nil {
		fmt.Println("写入文件错误:", err)
		return
	}
	// 读取文件内容
	var b = make([]byte, 5) // 创建一个缓冲区用于存储读取的数据
	_, err = file.Read(b)
	if err != nil {
		fmt.Println("读取文件错误:", err)
		return
	}
	fmt.Println("文件内容:", string(b))
	// 01 file.Openfile
	// 02 file.Read  || file.ReadAt  还是用 ReadFile
	// 03 file.Seek
	// 04 file.Write || file.WriteAt
}

func TestOpenFile(t *testing.T) {
	// 直接使用 io读取有
	data, err := os.ReadFile("1.txt")
	if err != nil {
		t.Fatal(err)
	}
	reader := bytes.NewReader(data)
	buf := make([]byte, len(data))
	reader.Read(buf)
	fmt.Println(string(buf))
}

func TestReadFile(t *testing.T) {
	file, err := os.OpenFile("1.txt", os.O_RDONLY, 0644)
	if err != nil {
		t.Fatal(err)
	}
	var b = make([]byte, 5) // 创建一个缓冲区用于存储读取的数据
	_, err = file.Read(b)
	if err != nil {
		fmt.Println("读取文件错误:", err)
		return
	}
	fmt.Println("文件内容:", string(b))
}

func TestFile(t *testing.T) {
	_, err := os.Create("1.txt")
	if err != nil {
		t.Fatal(err)
	}
}
