package viper

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/spf13/viper"
)

func TestViper(t *testing.T) {
	// 单个对象依次读取
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		//log.Fatal("read config failed: %v", err)
		fmt.Println(err)
	}
	fmt.Println(viper.Get("Mysql"))
	fmt.Println("Mysql password:", viper.Get("Mysql.password"))

	viper.SetConfigType("yaml")

	// 读取配置:1 从字节流当中
	tomlConfig := []byte(`
	Mysql:
	 user_name: qiangzhou1
	 password: 123456
	Redis:
	 ip: 123.23.23.23
	 port: 64379
	 userName: admin
	`)

	err = viper.ReadConfig(bytes.NewBuffer(tomlConfig))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(":", viper.Get("Redis.port"))

	// UnMarshal读取配置:2.从结构体有
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	var c Config
	err = viper.Unmarshal(&c)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("config message:%+v \n", c)

	//3. 监听文件配置
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	viper.WatchConfig()
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}

	viper.WatchConfig()
	fmt.Println("redis port before sleep: ", viper.Get("redis.port"))
	time.Sleep(time.Second * 10)
	fmt.Println("redis port after sleep: ", viper.Get("redis.port"))
}
