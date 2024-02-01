package proto

// Msg 消息
type Msg struct {
	PostType    string        `json:"post_type"`    // 请求类型
	MessageType string        `json:"message_type"` // 消息类型
	RawMessage  string        `json:"raw_message"`  // 原始消息
	Message     []MessageData `json:"message"`      // 消息
	Sender      Sender        `json:"sender"`       // 发送者
	UserID      *int64        `json:"user_id"`      // 用户ID
	GroupID     *int64        `json:"group_id"`     // 群ID
}
