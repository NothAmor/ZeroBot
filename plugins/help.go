package plugins

import (
	"fmt"

	"github.com/NothAmor/ZeroBot/core/common"
	"github.com/NothAmor/ZeroBot/core/common/plugin"
	"github.com/NothAmor/ZeroBot/core/proto"
	"github.com/NothAmor/ZeroBot/core/utils/qq"
)

// Help 帮助插件
type Help struct {
	plugin.PluginInfo // 插件信息(必须)
}

// GetPluginInfo 插件信息
func (h *Help) GetPluginInfo() *plugin.PluginInfo {
	return &h.PluginInfo
}

// Init 初始化
func (h *Help) Init() {
	h.Name = "帮助"                       // 插件名
	h.Description = "帮助插件"              // 插件描述
	h.Usage = "帮助,help"                 // 插件用法(,)分隔
	h.UsageType = proto.UsageTypePrefix // 插件用法类型
	h.Rule = proto.RuleTypeAll          // 插件规则
}

// Handle 处理器
func (h *Help) Handle(msg *proto.Msg, args []string) {
	helpContent := fmt.Sprintf(`你好, 我是%s, 一个基于Go语言的QQ机器人框架。
功能列表:
1. roll
	使用方法: @Bot Amor roll 麻辣烫 过桥米线 烤肉拌饭`, common.Config.Bot.Name)
	qq.ReplyText(msg, helpContent)
}
