package main

import (
	"bytes"
	"os"
	"testing"
)

func TestX(t *testing.T) {
	//Reader 接口
	data, err := os.ReadFile("a.txt")
	if err != nil {
		t.Fatal(err)
	}
	reader := bytes.NewReader(data)
	buf := make([]byte, len(data))
	n, _ := reader.Read(buf)
	t.Logf("读取字节数匹配期望 %d, 实际 %d\n", len(data), n)
	if n != len(data) {
		t.Fatalf("读取字节数不匹配，期望 %d, 实际 %d", len(data), n)
	}
	if !bytes.Equal(buf, data) {
		t.Fatal("读取的数据与文件内容不一致")
	}
}
