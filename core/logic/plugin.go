package logic

import (
	"github.com/NothAmor/ZeroBot/core/common"
	"github.com/NothAmor/ZeroBot/core/matcher"
	"github.com/NothAmor/ZeroBot/core/proto"
)

// PluginHandler 寻找匹配插件执行
func PluginHandler(rcvMsg *proto.Msg) {
	for _, plugin := range common.Plugins {
		pluginInfo := plugin.GetPluginInfo()

		isMatch, args, err := matcher.Matcher(rcvMsg, *pluginInfo)
		if err != nil {
			common.Log.Errorf("匹配器错误: %s", err)
			continue
		}

		if isMatch {
			common.Log.Infof("消息匹配插件: %s", pluginInfo.Name)
			plugin.Handle(rcvMsg, args)
		}
	}
}
