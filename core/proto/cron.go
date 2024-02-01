package proto

// CronTask 定时任务
type CronTask struct {
	Name string // 定时任务名称
	Spec string // 定时任务规则
	Func func() // 定时任务函数
}
