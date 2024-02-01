package matcher

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/NothAmor/ZeroBot/core/common"
	"github.com/NothAmor/ZeroBot/core/common/plugin"
	"github.com/NothAmor/ZeroBot/core/proto"
)

// OnCommand 命令匹配器
func OnCommand(msg *proto.Msg, pluginInfo plugin.PluginInfo) (match bool, args []string) {
	pluginConfig := common.Config.Plugins
	starter := pluginConfig.Command
	command := pluginInfo.Usage
	separator := pluginConfig.Separator

	commandList := strings.Split(command, ",")
	for _, cmd := range commandList {
		if strings.HasPrefix(msg.RawMessage, starter+cmd) {
			match = true
			args = strings.Split(strings.TrimLeft(msg.RawMessage, fmt.Sprintf("%s%s ", starter, command)), separator)
			return
		}
	}
	return false, nil
}

// OnPrefix 前缀匹配器
func OnPrefix(msg *proto.Msg, pluginInfo plugin.PluginInfo) (match bool) {
	pluginConfig := common.Config.Plugins
	starter := pluginConfig.Command

	commandList := strings.Split(pluginInfo.Usage, ",")
	for _, command := range commandList {
		prefix := fmt.Sprintf("%s%s", starter, command)
		if strings.HasPrefix(msg.RawMessage, prefix) {
			return true
		}
	}
	return false
}

// OnRegex 正则匹配器
func OnRegex(msg *proto.Msg, pluginInfo plugin.PluginInfo) (match bool) {
	regex := regexp.MustCompile(pluginInfo.Usage)
	match = regex.MatchString(msg.RawMessage)
	return
}

// OnEqual 相等匹配器
func OnEqual(msg *proto.Msg, pluginInfo plugin.PluginInfo) (match bool) {
	commandList := strings.Split(pluginInfo.Usage, ",")
	for _, command := range commandList {
		if msg.RawMessage == command {
			return true
		}
	}
	return
}

// Matcher 匹配器
func Matcher(msg *proto.Msg, pluginInfo plugin.PluginInfo) (match bool, args []string, err error) {

	// 匹配消息类型
	if pluginInfo.Rule == msg.MessageType || pluginInfo.Rule == proto.RuleTypeAll {
		match = true
	} else {
		match = false
		return
	}

	switch pluginInfo.UsageType {
	case proto.UsageTypeCommand:
		match, args = OnCommand(msg, pluginInfo)
	case proto.UsageTypePrefix:
		match = OnPrefix(msg, pluginInfo)
		args = nil
	case proto.UsageTypeRegex:
		match = OnRegex(msg, pluginInfo)
		args = nil
	default:
		err = fmt.Errorf("unknown usage type: %s", pluginInfo.UsageType)
		return
	}

	return
}
