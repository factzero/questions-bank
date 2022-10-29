package core

import (
	"flag"
	"fmt"
	"os"
	"server/global"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

const ConfigEnv = "GVA_CONFIG"
const ConfigDefaultFile = "config.yaml"

func Viper(path ...string) *viper.Viper {
	var cfgFile string

	if len(path) == 0 {
		flag.StringVar(&cfgFile, "c", "", "choose config file.")
		flag.Parse()
		if cfgFile == "" { // 判断命令行参数是否为空
			if configEnv := os.Getenv(ConfigEnv); configEnv == "" { // 判断 internal.ConfigEnv 常量存储的环境变量是否为空
				switch gin.Mode() {
				case gin.DebugMode:
					cfgFile = ConfigDefaultFile
					fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.EnvGinMode, ConfigDefaultFile)
				}
			} else {
				cfgFile = configEnv
				fmt.Printf("您正在使用%s环境变量,config的路径为%s\n", ConfigEnv, cfgFile)
			}
		} else { // 命令行参数不为空 将值赋值于cfgFile
			fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%s\n", cfgFile)
		}
	} else { // 函数传递的可变参数的第一个值赋值于config
		cfgFile = path[0]
		fmt.Printf("您正在使用func Viper()传递的值,config的路径为%s\n", cfgFile)
	}

	v := viper.New()
	v.SetConfigFile(cfgFile)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		fmt.Println("Fatal error viper read config file:", err)
		panic(err)
	}

	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed:", in.Name)
		if err := v.Unmarshal(&global.GVA_CONFIG); err != nil {
			fmt.Println(err)
		}
	})

	if err := v.Unmarshal(&global.GVA_CONFIG); err != nil {
		fmt.Println(err)
	}

	return v
}
