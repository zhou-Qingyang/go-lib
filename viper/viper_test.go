package viper

import (
	"fmt"
	"testing"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func TestViper(t *testing.T) {
	var cf Config
	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath(".")
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&cf); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&cf); err != nil {
		panic(err)
	}
	fmt.Printf("%+v \n", cf)
}
