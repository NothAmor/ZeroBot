package initialization

import (
	"github.com/NothAmor/ZeroBot/core/common"
	"github.com/NothAmor/ZeroBot/core/common/plugin"
)

// ZeroBotRegisterPlugins 注册插件
func ZeroBotRegisterPlugins(plugin ...plugin.Accessor) {
	for _, plugin := range plugin {
		plugin.Init()
		pluginInfo := plugin.GetPluginInfo()
		common.Log.Infof("插件: [%s] 已加载", pluginInfo.Name)
		common.Plugins = append(common.Plugins, plugin)
	}
}
