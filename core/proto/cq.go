package proto

// cq-http 通用消息post_type类型
const (
	PostTypeMessage          = "message"      // 消息
	PostTypeMessageSent      = "message_sent" // 消息发送
	PostTypeMessageRequest   = "request"      // 请求
	PostTypeMessageNotice    = "notice"       // 通知
	PostTypeMessageMetaEvent = "meta_event"   // 元事件
)

// cq-http 元事件类型
const (
	MetaEventTypeHeartbeat = "heartbeat" // 心跳
	MetaEventTypeLifecycle = "lifecycle" // 生命周期
)

// cq-http 通用消息
type CommonMsg struct {
	PostType string `json:"post_type"` // 消息类型
	SelfID   int64  `json:"self_id"`   // 机器人QQ号
	Time     int64  `json:"time"`      // 时间戳
}

const (
	MessageTypePrivate = "private" // 私聊消息
	MessageTypeGroup   = "group"   // 群消息
)

const (
	SendPrivateMessageAction = "send_private_msg" // 发送私聊消息
	SendGroupMessageAction   = "send_group_msg"   // 发送群消息
)

// cq-http 通用消息模版
type (
	Params struct {
		UserID  int64  `json:"user_id,omitempty"`  // 目标用户的 QQ 号
		GroupID int64  `json:"group_id,omitempty"` // 目标群的群号
		Message string `json:"message"`            // 回复的消息内容
	}
	MessageTemplate struct {
		Action string `json:"action"` // API名称
		Params Params `json:"params"` // API参数
	}
)

// cq-http 发送消息响应
type MessageSentResp struct {
	Data struct {
		MessageID int `json:"message_id"`
	} `json:"data"`
	Message string `json:"message"`
	Retcode int    `json:"retcode"`
	Status  string `json:"status"`
}

// 元事件类型
type MetaEvent struct {
	MetaEventType string `json:"meta_event_type"`
}

// 心跳元事件
type HeartbeatMetaEvent struct {
	Status struct {
		AppEnabled     bool  `json:"app_enabled"`     // 程序是否可用
		AppGood        bool  `json:"app_good"`        // 程序正常
		AppInitialized bool  `json:"app_initialized"` // 程序是否初始化完毕
		Online         bool  `json:"online"`          // 是否在线
		PluginsGood    *bool `json:"plugins_good"`    // 插件正常(可能为 null)
		Stat           struct {
			PacketReceived  int `json:"packet_received"`   // 接收到的数据包数量
			PacketSent      int `json:"packet_sent"`       // 发送的数据包数量
			PacketLost      int `json:"packet_lost"`       // 丢弃的数据包数量
			MessageReceived int `json:"message_received"`  // 接收到的消息数量
			MessageSent     int `json:"message_sent"`      // 发送的消息数量
			DisconnectTimes int `json:"disconnect_times"`  // 断开连接的次数
			LostTimes       int `json:"lost_times"`        // 丢失连接的次数
			LastMessageTime int `json:"last_message_time"` // 最后一条消息的时间戳
		} `json:"stat"` // 统计信息
	} `json:"status"` // 状态
}

// 生命周期元事件
type LifecycleMetaEvent struct {
	SubType string `json:"sub_type"` // 生命周期子类型, 有 enable(启用), disable(停用), connect(连接)
}

// 发送人信息
type Sender struct {
	Age      int     `json:"age"`      // 年龄
	Area     *string `json:"area"`     // 地区
	Card     *string `json:"card"`     // 群名片
	Level    *string `json:"level"`    // 等级
	Nickname string  `json:"nickname"` // 昵称
	Role     *string `json:"role"`     // 角色
	Sex      string  `json:"sex"`      // 性别
	Title    *string `json:"title"`    // 头衔
	UserID   int64   `json:"user_id"`  // QQ号
}

// 消息数据
type MessageData struct {
	Type string `json:"type"`
	Data struct {
		ID   string `json:"id"`   // 消息ID
		Type string `json:"type"` // 消息类型
		Text string `json:"text"` // 消息内容
	} `json:"data"`
}

// 消息
// {"post_type":"message","message_type":"private","time":1706691414,"self_id":3626129355,"sub_type":"friend","sender":{"age":0,"nickname":"空虚公子","sex":"unknown","user_id":1565481748},"message_id":15217893,"user_id":1565481748,"target_id":3626129355,"message":[{"type":"text","data":{"text":"你好呀"}}],"raw_message":"你好呀","font":0}
type Message struct {
	MessageType string        `json:"message_type"` // 消息类型(private, group)
	SubType     string        `json:"sub_type"`     // 消息子类型(group, public)
	Message     []MessageData `json:"message"`      //
	RawMessage  string        `json:"raw_message"`  // CQ 码格式的消息
	Font        int           `json:"font"`         // 字体(0)
	Sender      Sender        `json:"sender"`       // 发送者信息
	MessageID   int           `json:"message_id"`   // 消息ID
	UserID      *int64        `json:"user_id"`      // 发送者QQ号
	GroupID     *int64        `json:"group_id"`     // 群号
	MessageSeq  int           `json:"message_seq"`  //
	Anonymous   interface{}   `json:"anonymous"`    // 匿名信息(id匿名用户 ID, name匿名用户名称, flag匿名用户 flag, 在调用禁言 API 时需要传入)
}
