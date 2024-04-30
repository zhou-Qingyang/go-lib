package main

import (
	"fmt"
	"os"
	"testing"
	"unicode/utf8"
)

func TestX1(t *testing.T) {
	// strings.NewReader  bytes.Reader
	//ReaderAt
	// data := make([]byte, 20)
	// reader := strings.NewReader("Go语言学习之路")
	// reader.ReadAt(data, 2)
	// fmt.Println(string(data))
	chineseChar := 'G'
	byteCount := utf8.RuneLen(chineseChar)
	fmt.Printf("%v 占据的字节数为：%d\n", chineseChar, byteCount)
	file, err := os.Create("writeAt.txt")
	// os.Open("writeAt.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString("Golang文中社区——这里是多余") //18
	n, err := file.WriteAt([]byte("Go语言中文网"), 24)
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
}
