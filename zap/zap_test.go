package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestZapLog(t *testing.T) {
	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath(".")
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	if err = v.Unmarshal(&GlobalConf); err != nil {
		panic(err)
	}

	if ok, _ := PathExists(GlobalConf.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", GlobalConf.Zap.Director)
		_ = os.Mkdir(GlobalConf.Zap.Director, os.ModePerm)
	}
	levels := GlobalConf.Zap.Levels()
	length := len(levels)
	cores := make([]zapcore.Core, 0, length)
	for i := 0; i < length; i++ {
		core := NewZapCore(levels[i])
		cores = append(cores, core)
	}
	logger := zap.New(zapcore.NewTee(cores...))
	if GlobalConf.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	logger.Error("这样子报错篇？")
	logger.Info("hello world2223")
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, err
	}
	return false, nil
}
