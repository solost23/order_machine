package main

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"order_machine/configs"
	"order_machine/internal/server"
	"os"
	"path"
)

var (
	WebConfigPath = "configs/conf.yml"
	version       = "__BUILD_VERSION_"
	execDir       string
	st, v, V      bool
)

func main() {
	flag.StringVar(&execDir, "d", ".", "项目目录")
	flag.BoolVar(&v, "v", false, "查看版本号")
	flag.BoolVar(&V, "V", false, "查看版本号")
	flag.BoolVar(&st, "s", false, "项目状态")
	flag.Parse()
	if v || V {
		fmt.Println(version)
		os.Exit(-1)
	}
	// run
	serverConfig, err := InitConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	InitLogger()
	server.Run(serverConfig)
}

func InitConfig() (serverConfig *configs.ServerConfig, err error) {
	serverConfig = new(configs.ServerConfig)
	confPath := path.Join(execDir, WebConfigPath)
	v := viper.New()
	v.SetConfigFile(confPath)
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	if err := v.Unmarshal(serverConfig); err != nil {
		return nil, err
	}
	return serverConfig, nil
}

func InitLogger() {

}
