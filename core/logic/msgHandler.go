package logic

import (
	"encoding/json"

	"github.com/NothAmor/ZeroBot/core/common"
	"github.com/NothAmor/ZeroBot/core/proto"
	"github.com/gorilla/websocket"
)

// HandleMsg 处理消息
func HandleMsg(conn *websocket.Conn, msg []byte) func() {
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
			MsgHandler(conn, &commonMsg, msg)
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
			common.Log.Errorf("Unknown message type: %s", msg)
			return
		}
	}
}
