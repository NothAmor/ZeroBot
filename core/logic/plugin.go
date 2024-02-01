package logic

import (
	"github.com/NothAmor/ZeroBot/core/common"
	"github.com/spf13/cast"
)

// PluginHandler 寻找匹配插件执行
func PluginHandler(msgType string, msg []byte) {
	for _, plugin := range common.Plugins {
		pluginInfo := plugin.GetPluginInfo()
		if pluginInfo.Rule == "private" {
			if plugin.Matcher(msgType, cast.ToString(msg)) {
				common.Log.Infof("Matched plugin: %s", pluginInfo.Name)
				plugin.Handle(cast.ToString(msg))
			}
		}
	}
}
