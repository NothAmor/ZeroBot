package proto

// cq-http 通用消息
type CommonMsg struct {
	PostType string `json:"post_type"`
	SelfID   int64  `json:"self_id"`
	Time     int64  `json:"time"`
}
