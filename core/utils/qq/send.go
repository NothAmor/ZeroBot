package qq

import (
	"encoding/json"

	"github.com/NothAmor/ZeroBot/core/common"
	"github.com/NothAmor/ZeroBot/core/proto"
	"github.com/gorilla/websocket"
)

// SendPrivateMessage 发送私聊消息
func SendPrivateMessage(userID int64, message string) (err error) {
	if common.Conn == nil {
		common.Log.Error("Connection is nil")
		return
	}

	reply := proto.MessageTemplate{
		Action: proto.SendPrivateMessageAction,
		Params: proto.Params{
			UserID:  userID,
			Message: message,
		},
	}

	err = send(reply)
	if err != nil {
		common.Log.Errorf("Failed to send private message: %v", err)
		return
	}

	return
}

// SendGroupMessage 发送群消息
func SendGroupMessage(groupID int64, message string) (err error) {
	if common.Conn == nil {
		common.Log.Error("Connection is nil")
		return
	}

	reply := proto.MessageTemplate{
		Action: proto.SendGroupMessageAction,
		Params: proto.Params{
			GroupID: groupID,
			Message: message,
		},
	}

	err = send(reply)
	if err != nil {
		common.Log.Errorf("Failed to send group message: %v", err)
		return
	}

	return
}

func send(data interface{}) (err error) {
	if common.Conn == nil {
		common.Log.Error("Connection is nil")
		return
	}

	replyJSON, err := json.Marshal(data)
	if err != nil {
		common.Log.Errorf("Failed to marshal message: %v", err)
		return
	}

	common.Log.Infof("发送消息: [%s]", replyJSON)

	err = common.Conn.WriteMessage(websocket.TextMessage, replyJSON)
	if err != nil {
		common.Log.Errorf("Failed to send message: %v", err)
		return
	}

	return
}
