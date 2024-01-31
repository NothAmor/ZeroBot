package logic

import (
	"encoding/json"

	"github.com/NothAmor/ZeroBot/core/common"
	"github.com/NothAmor/ZeroBot/core/proto"
	"github.com/spf13/cast"

	"github.com/gorilla/websocket"
)

// MsgHandler 处理消息
func MsgHandler(conn *websocket.Conn, commonMsg *proto.CommonMsg, msg []byte) (err error) {
	// 获取消息类型
	var msgContent proto.Message

	err = json.Unmarshal(msg, &msgContent)
	if err != nil {
		common.Log.Errorf("Failed to unmarshal private message: %v", err)
		return
	}

	switch msgContent.MessageType {
	case "private":
		// 私聊消息
		PrivateMsgHandler(conn, commonMsg, msg)

	case "group":
		// TODO: 群消息
		// GroupMsgHandler(commonMsg, msg)

	default:
		common.Log.Errorf("Unknown message type: %s", msgContent.MessageType)
		return
	}

	return
}

// PrivateMsgHandler TODO: 处理私聊消息
func PrivateMsgHandler(conn *websocket.Conn, commonMsg *proto.CommonMsg, msg []byte) (err error) {
	var privateMsg proto.Message

	err = json.Unmarshal(msg, &privateMsg)
	if err != nil {
		common.Log.Errorf("Failed to unmarshal private message: %v", err)
		return
	}

	for _, plugin := range common.Plugins {
		if plugin.GetRule() == "private" {
			if plugin.Matcher(privateMsg.MessageType, cast.ToString(msg)) {
				plugin.Handle(conn, cast.ToString(msg))
			}
		}
	}

	return
}
