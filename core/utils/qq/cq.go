package qq

import (
	"encoding/json"
	"log"

	"github.com/NothAmor/ZeroBot/core/common"
	"github.com/NothAmor/ZeroBot/core/proto"
	"github.com/gorilla/websocket"
)

// SendPrivateMessage 发送私聊消息
func SendPrivateMessage(conn *websocket.Conn, userID int64, message string) (err error) {
	reply := proto.MessageTemplate{
		Action: proto.SendPrivateMessageAction,
		Params: proto.Params{
			UserID:  userID,
			Message: message,
		},
	}

	err = send(conn, reply)
	if err != nil {
		common.Log.Errorf("Failed to send private message: %v", err)
		return
	}

	return
}

// SendGroupMessage 发送群消息
func SendGroupMessage(conn *websocket.Conn, groupID int64, message string) (err error) {
	reply := proto.MessageTemplate{
		Action: proto.SendGroupMessageAction,
		Params: proto.Params{
			GroupID: groupID,
			Message: message,
		},
	}

	err = send(conn, reply)
	if err != nil {
		common.Log.Errorf("Failed to send group message: %v", err)
		return
	}

	return
}

func send(conn *websocket.Conn, data interface{}) (err error) {
	replyJSON, err := json.Marshal(data)
	if err != nil {
		log.Fatal("Failed to marshal reply message:", err)
	}

	err = conn.WriteMessage(websocket.TextMessage, replyJSON)
	if err != nil {
		log.Fatal("Failed to send reply message:", err)
	}
	return
}
