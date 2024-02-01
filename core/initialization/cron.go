package initialization

import (
	"github.com/NothAmor/ZeroBot/core/common"
	"github.com/NothAmor/ZeroBot/core/proto"
	"github.com/robfig/cron/v3"
)

// initCron 初始化定时任务
func initCron() {
	common.Cron = cron.New()
}

// ZeroBotRegisterCron 注册定时任务
func ZeroBotRegisterCron(tasks ...*proto.CronTask) {
	for _, task := range tasks {
		_, err := common.Cron.AddFunc(task.Spec, task.Func)
		if err != nil {
			common.Log.Errorf("Failed to add cron task: %v", err)
			return
		}

		common.Log.Infof("定时任务: [%s] 已加载, 任务执行时间: [%s]", task.Name, task.Spec)
	}
}
