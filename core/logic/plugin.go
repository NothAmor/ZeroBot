package logic

import (
	"github.com/NothAmor/ZeroBot/core/common"
	"github.com/NothAmor/ZeroBot/core/proto"
)

// PluginHandler 寻找匹配插件执行
func PluginHandler(rcvMsg *proto.Msg) {
	for _, plugin := range common.Plugins {
		pluginInfo := plugin.GetPluginInfo()

		if plugin.Matcher(rcvMsg.MessageType, rcvMsg.RawMessage) {
			common.Log.Infof("消息匹配插件: %s", pluginInfo.Name)
			plugin.Handle(rcvMsg)
		}
	}
}
