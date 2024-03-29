package initialization

import (
	"fmt"
	"io/ioutil"

	"github.com/NothAmor/ZeroBot/core/common"
	"gopkg.in/yaml.v3"
)

// 启动ZeroBot框架
func ZeroBotStart() {

	// 启动定时任务
	common.Cron.Start()

	// 启动Gin Web
	if err := common.Web.Run(fmt.Sprintf(":%d", common.Config.System.Port)); err != nil {
		panic(err)
	}
}

// 初始化ZeroBot框架
func ZeroBotInit() (err error) {
	// 初始化日志
	err = initLog()
	if err != nil {
		return
	}

	// 初始化配置文件
	err = initConfig()
	if err != nil {
		return
	}

	// 初始化全局变量
	initCommonVar()

	// 初始化定时任务
	initCron()

	// 初始化Gin Web
	err = initGin()
	if err != nil {
		return
	}

	fmt.Println(common.LOGO)
	fmt.Printf("ZeroBot %s\n", common.VERSION)
	fmt.Println("ZeroBot框架初始化完成")
	fmt.Println("欢迎使用ZeroBot框架")

	return
}

// 初始化配置文件
func initConfig() (err error) {
	// 读取配置文件
	configFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		return
	}

	err = yaml.Unmarshal(configFile, &common.Config)
	if err != nil {
		return
	}

	return
}

// 初始化全局变量
func initCommonVar() {
	common.VERSION = common.Config.System.Version
}
