package common

import "github.com/gorilla/websocket"

// Accessor 插件接口
type Accessor interface {
	GetName() string                         // 插件名称
	GetDescription() string                  // 插件描述
	GetUsage() string                        // 插件用法
	GetUsageType() string                    // 插件用法类型
	GetRule() string                         // 插件规则
	Init()                                   // 初始化
	Matcher(rule, msg string) bool           // 匹配器
	Handle(conn *websocket.Conn, msg string) // 处理器
}
