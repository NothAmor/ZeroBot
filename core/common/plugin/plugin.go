package plugin

import "github.com/NothAmor/ZeroBot/core/proto"

// PluginInfo 插件信息
type PluginInfo struct {
	Name        string // 插件名称
	Description string // 插件描述
	Usage       string // 插件用法
	UsageType   string // 插件用法类型
	Rule        string // 插件规则
}

// Accessor 插件接口
type Accessor interface {
	GetPluginInfo() *PluginInfo           // 插件信息
	Init()                                // 初始化
	Handle(msg *proto.Msg, args []string) // 处理器
}
