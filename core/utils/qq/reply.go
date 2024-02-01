package qq

import (
	"fmt"
	"net/url"

	"github.com/NothAmor/ZeroBot/core/proto"
)

// ReplyText 回复文本消息
func ReplyText(msg *proto.Msg, text string) (err error) {
	if msg.GroupID != nil {
		err = SendGroupMessage(*msg.GroupID, text)
	} else {
		err = SendPrivateMessage(*msg.UserID, text)
	}

	return
}

// ReplyImage 回复图片消息, image为图片路径或URL
func ReplyImage(msg *proto.Msg, image string) {
	image = fmt.Sprintf("[CQ:image,file=%s]", url.QueryEscape(image))

	// for _, v := range []string{"&", "[", "]", ","} {
	// 	image = strings.Replace(image, v, fmt.Sprintf("&#%d;", v[0]), -1)
	// }

	if msg.GroupID != nil {
		SendGroupMessage(*msg.GroupID, image)
	} else {
		SendPrivateMessage(*msg.UserID, image)
	}
}
