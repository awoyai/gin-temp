package core

import (
	"flag"
	"fmt"

	"github.com/awoyai/gin-temp/global"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Viper //
// 优先级: 命令行 > 环境变量 > 默认值
func Viper() *viper.Viper {
	var config string
	flag.StringVar(&config, "c", "", "choose config file.")
	flag.Parse()

	if config == "" { // 判断命令行参数是否为空
		config = "./config.yaml"
	} else { // 命令行参数不为空 将值赋值于config
		fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%s\n", config)
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.CONFIG); err != nil {
			fmt.Println(err)
		}
		fmt.Printf("changed config: %+v\n", global.CONFIG)
	})
	if err = v.Unmarshal(&global.CONFIG); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("config: %+v\n", global.CONFIG)
	return v
}
