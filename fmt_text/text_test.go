package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestU(t *testing.T) {
	t.Log(strings.Contains("failure", "f"))
	t.Log(strings.HasPrefix("failure", "fa"))
	t.Log(strings.HasPrefix("failure", "as"))
	t.Log(strings.HasSuffix("failure", "ae"))
	// 字串出现的位置
	t.Log(strings.Index("failure", "a"))
	t.Log(strings.LastIndex("failurea", "a"))      //子串
	t.Log(strings.LastIndexByte("feailurea", 'e')) //子字符
	// han := func(c rune) bool {
	// 	return unicode.Is(unicode.Han, c) // 汉字
	// }
	// fmt.Println(strings.IndexFunc("Hello, 世界", han))
	t.Log(Join("hello", []byte{'a', 'b'}))
	t.Log(strings.Join([]string{"a", "b"}, "11"))

	t.Log("ba" + strings.Repeat("na", 2))
	t.Log(strings.Replace("oink oink oink", "k", "ky", 2))
	t.Log(strings.Replace("oink oink oink", "oink", "moo", -1))
	t.Log(strings.ReplaceAll("oink oink oink", "oink", "moo"))
	//x := "!!!@@@你好,!@#$ Gophers###$$$"
	//fmt.Println(strings.Trim(x, "@#$!%^&*()_+=-"))
	s := strings.Builder{}
	s.WriteString("Hello")
	t.Log(s.String())
}

func Join(a string, b []byte) string {
	if len(b) == 0 {
		return a
	}

	buffer := bytes.NewBufferString(a)
	for _, v := range b {
		buffer.WriteByte(v)
	}
	return buffer.String()
}
