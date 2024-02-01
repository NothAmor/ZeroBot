package proto

const (
	RuleTypePrivate = MessageTypePrivate // 私聊消息
	RuleTypeGroup   = MessageTypeGroup   // 群聊消息
	RuleTypeAll     = "all"              // 所有消息
)

const (
	UsageTypeCommand = "command" // 命令
	UsageTypeRegex   = "regex"   // 正则
	UsageTypePrefix  = "prefix"  // 前缀
)
