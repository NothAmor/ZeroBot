package main

import (
	"runtime/debug"

	"github.com/NothAmor/ZeroBot/core/common"
	"github.com/NothAmor/ZeroBot/core/initialization"
	"github.com/NothAmor/ZeroBot/core/proto"
	"github.com/NothAmor/ZeroBot/core/utils/notify"
	"github.com/NothAmor/ZeroBot/plugins"
	"github.com/NothAmor/ZeroBot/plugins/help"
	"github.com/spf13/cast"
)

func main() {
	// 初始化ZeroBot框架
	err := initialization.ZeroBotInit()
	if err != nil {
		panic(err)
	}

	// 注册插件
	initialization.ZeroBotRegisterPlugins(
		&help.Help{}, // 帮助插件
	)

	// 注册定时任务
	initialization.ZeroBotRegisterCron(
		&proto.CronTask{Name: "早上好", Spec: "02 11 * * *", Func: plugins.Morning}, // 早上好
	)

	// panic recover
	defer recovery()

	// 启动ZeroBot框架
	initialization.ZeroBotStart()
}

func recovery() {
	if rec := recover(); rec != nil {
		common.Log.Errorln("Panic Panic occur")
		if err, ok := rec.(error); ok {
			common.Log.Errorf("PanicRecover Unhandled error: %v\n stack:%v", err.Error(), cast.ToString(debug.Stack()))
		} else {
			common.Log.Errorf("PanicRecover Panic: %v\n stack:%v", rec, cast.ToString(debug.Stack()))
		}
		notify.Notify("ZeroBot Panic", cast.ToString(debug.Stack()))
	}
}
