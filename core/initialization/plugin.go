package initialization

import "github.com/NothAmor/ZeroBot/core/common"

// ZeroBotRegisterPlugins 注册插件
func ZeroBotRegisterPlugins(plugin ...common.Accessor) {
	for _, plugin := range plugin {
		plugin.Init()
		common.Log.Infof("插件: [%s] 已加载", plugin.GetName())
		common.Plugins = append(common.Plugins, plugin)
	}
}
