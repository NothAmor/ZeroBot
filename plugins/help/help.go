package help

import (
	"github.com/NothAmor/ZeroBot/core/common"
	"github.com/NothAmor/ZeroBot/core/proto"
	"github.com/NothAmor/ZeroBot/core/utils/qq"
)

// Help 帮助插件
type Help struct {
	Name        string // 插件名称
	Description string // 插件描述
	Usage       string // 插件用法
	UsageType   string // 插件用法类型
	Rule        string // 插件规则
}

// GetPluginInfo 插件信息
func (h *Help) GetPluginInfo() *common.PluginInfo {
	return &common.PluginInfo{
		Name:        h.Name,
		Description: h.Description,
		Usage:       h.Usage,
		UsageType:   h.UsageType,
		Rule:        h.Rule,
	}
}

// Init 初始化
func (h *Help) Init() {
	h.Name = "帮助"
	h.Description = "帮助插件"
	h.Usage = "帮助/help"
	h.UsageType = "command"
	h.Rule = "private"
}

// Matcher 匹配器
func (h *Help) Matcher(rule, msg string) bool {
	if (rule == "private" || rule == "group") && (msg == "help" || msg == "帮助") {
		return true
	}
	return false
}

// Handle 处理器
func (h *Help) Handle(msg *proto.Msg) {
	helpContent := `你好, 我是ZeroBot, 一个基于Go语言的QQ机器人框架。`
	qq.ReplyText(msg, helpContent)
}
