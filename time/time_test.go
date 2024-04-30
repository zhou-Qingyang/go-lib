package main

import (
	"testing"
	"time"
)

func TestXxx(t *testing.T) {
	// time 包中提供两个 时间地区
	// Local 和 UTC Local 代表当前系统本地时区；UTC 代表通用协调时间，也就是零时区
	// 默认使用UTC
	// Unix 时间是指从 1970 1月1日开始计算的秒
	// location, _ := time.LoadLocation("Local")
	// location1, _ := time.LoadLocation("UTC")
	// t.Log(location.String())
	// t.Log(location1.String())

	//time 默认纳秒级别
	// 1us = 1000ns
	// 1ms = 1000us
	// 1s = 1000ms
	// 获取当前时间的 UTC 时间
	// utcTime := time.Now().UTC()
	// // 获取当前时间的本地时间
	// localTime := time.Now().Local()
	// unix := time.Now().Unix()
	// t.Log(utcTime)
	// t.Log(localTime)
	// t.Log(unix) //1714448604

	// 时间的解析
	// 如果使用 time.Parse() 方法，需要提供时区
	// time.Now  默认是Local 时区  time.Parse UTC
	// 所以默认使用 time.ParseInLocation

	t.Log(time.Local)
	tdd1 := time.Unix(time.Now().Unix(), 0) //根据unix 时间戳获取时间结构体
	t.Log(tdd1.Format("2006-01-02 15:04:05"))
	// 01. time.Unix
	// 02. time.Format
	// 03. time.ParseInLocation || time.Parse
	// 04. time.Date
	e := time.Now()
	// 使用 time.Date 创建一个表示 2024 年 4 月 30 日 12 点的时间对象
	fm := time.Date(e.Year(), e.Month(), e.Day(), 12, 0, 0, 0, time.Local)
	t.Log(fm.Format("2006-01-02 15:04:05"))
}

func TestTime2(t *testing.T) {
	// Timer
}
