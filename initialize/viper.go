package initialize

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/TravisRoad/gotracker/global"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

const (
	configDevFile  = "config.dev.yaml"
	configProdFile = "config.prod.yaml"
)

// 命令行优先级 > 环境变量 > 配置文件
func initViper() *viper.Viper {
	var configFile string

	flag.StringVar(&configFile, "config", "", "choose config file")
	flag.Parse()

	mode := os.Getenv("MODE")

	global.Mode = mode
	switch mode {
	case "TEST":
		_, file, _, _ := runtime.Caller(0)
		global.WorkspacePath = filepath.Dir(filepath.Dir(file))

	case "DEBUG":
	case "PROD":
	default:
		global.Mode = "DEBUG" // 默认为 DEBUG
	}

	if configFile == "" { // 未在命令行设置配置文件路径
		switch global.Mode {
		case "TEST":
			configFile = filepath.Join(global.WorkspacePath, configDevFile)
		case "DEBUG":
			configFile = configDevFile
		case "PROD":
			configFile = configProdFile
		default:
			panic("mode error")
		}
	}
	fmt.Printf("正在使用 %s 模式，config 路径为 %s\n", global.Mode, configFile)

	v := viper.New()
	v.SetConfigFile(configFile)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("fatal config file error: %s", err.Error()))
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed: ", e.Name)
		if err = v.Unmarshal(&global.Config); err != nil {
			fmt.Println("parse config file error: ", err.Error())
		}
	})

	// 解析到 global.Config 上
	if err = v.Unmarshal(&global.Config); err != nil {
		fmt.Println("parse config file error: ", err.Error())
	}

	return v
}
