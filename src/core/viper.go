package core

import (
	"fmt"
	"log"
	"mayday/src/global"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

const (
	ConfigFile = "./config/config.yaml"
)

func Viper(path ...string) *viper.Viper {
	var config string
	if len(path) == 0 {
		config = ConfigFile
		log.Printf("正在使用读取配置文件,的路径为%v\n", ConfigFile)
	} else {
		config = path[0]
		log.Printf("正在使用func Viper()传递的值,config的路径为%v\n", config)
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("读取配置文件错误: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		log.Println("配置文件被修改:", e.Name)
		if err := v.Unmarshal(&global.GVA_CONFIG); err != nil {
			log.Println(err)
		}
	})

	if err := v.Unmarshal(&global.GVA_CONFIG); err != nil {
		log.Println(err)
	}

	return v
}
