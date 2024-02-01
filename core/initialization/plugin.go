package initialization

import "github.com/NothAmor/ZeroBot/core/common"

// ZeroBotRegisterPlugins 注册插件
func ZeroBotRegisterPlugins(plugin ...common.Accessor) {
	for _, plugin := range plugin {
		plugin.Init()
		pluginInfo := plugin.GetPluginInfo()
		common.Log.Infof("插件: [%s] 已加载", pluginInfo.Name)
		common.Plugins = append(common.Plugins, plugin)
	}
}
