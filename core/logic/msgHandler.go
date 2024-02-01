package logic

import (
	"encoding/json"

	"github.com/NothAmor/ZeroBot/core/common"
	"github.com/NothAmor/ZeroBot/core/proto"
)

// HandleMsg 处理消息
func HandleMsg(msg []byte) func() {
	return func() {
		// 通用消息
		var commonMsg proto.CommonMsg

		err := json.Unmarshal(msg, &commonMsg)
		if err != nil {
			common.Log.Errorf("Failed to unmarshal message: %v", err)
			return
		}

		switch commonMsg.PostType {
		case proto.PostTypeMessage:
			// 消息
			MsgHandler(&commonMsg, msg)
		case proto.PostTypeMessageSent:
			// 消息发送
		case proto.PostTypeMessageRequest:
			// 请求
		case proto.PostTypeMessageNotice:
			// 通知
		case proto.PostTypeMessageMetaEvent:
			// 元事件
			MetaEventHandler(&commonMsg, msg)

		default:
			// 有可能是消息发送响应, 不是的话则是未知消息
			var messageResp proto.MessageSentResp
			err = json.Unmarshal(msg, &messageResp)
			if err != nil {
				common.Log.Errorf("Failed to unmarshal message: %v", err)
				return
			}

			if messageResp.Data.MessageID != 0 {
				return
			}

			common.Log.Errorf("Unknown message type: %s", msg)
			return
		}
	}
}
