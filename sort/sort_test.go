package main

import (
	"sort"
	"testing"
)

// 学生成绩结构体
type StuScore struct {
	name  string // 姓名
	score int    // 成绩
}

type StuScores []StuScore

// Len()
func (s StuScores) Len() int {
	return len(s)
}

// Less(): 成绩将有低到高排序
// 如果 i 索引的数据小于 j 索引的数据，
// 返回 true，且不会调用下面的 Swap()，即数据升序排序。
func (s StuScores) Less(i, j int) bool {
	return s[i].score < s[j].score
}

// Swap()
func (s StuScores) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func TestXxx(t *testing.T) {
	// 自定义结构体 StuScore 需要先实现 sort.Interface 接口
	stus := StuScores{
		{"alan", 95},
		{"hikerell", 91},
		{"acmfly", 96},
		{"leao", 90},
	}
	sort.Sort(stus)
	t.Log(stus)

	// 原生int32 string float64
	sort.Ints([]int{1, 3, 2, 4, 5})
	//1.降序排序
	sort.Sort(sort.Reverse(sort.IntSlice([]int{1, 3, 2, 4, 5})))
	//2.SearchInts() 的使用条件为：切片 a 已经升序排序 以下是一个错误使用的例子

	//对于interface
	people := []struct {
		Name string
		Age  int
	}{
		{"Gopher", 7},
		{"Alice", 55},
		{"Vera", 24},
		{"Bob", 75},
	}
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	t.Log(people)
	//1 sort.Slice
	//2 sort.SliceStable
	//3 sort.SliceIsSorted //判断是否已排序
	//4 sort.Search //查找某个元素 先执行升序排序
}
