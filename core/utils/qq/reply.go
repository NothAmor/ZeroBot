package qq

import "github.com/NothAmor/ZeroBot/core/proto"

// ReplyText 回复文本消息
func ReplyText(msg *proto.Msg, text string) (err error) {
	if msg.GroupID != nil {
		err = SendGroupMessage(*msg.GroupID, text)
	} else {
		err = SendPrivateMessage(*msg.UserID, text)
	}

	return
}

// ReplyImage 回复图片消息
func ReplyImage(msg *proto.Msg, image string) {
}
