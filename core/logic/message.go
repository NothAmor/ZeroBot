package logic

import (
	"encoding/json"

	"github.com/NothAmor/ZeroBot/core/common"
	"github.com/NothAmor/ZeroBot/core/proto"
)

// MsgHandler 处理消息
func MsgHandler(commonMsg *proto.CommonMsg, msg []byte) (err error) {
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
		PrivateMsgHandler(commonMsg, msg)

	case "group":
		// 群消息
		GroupMsgHandler(commonMsg, msg)

	default:
		common.Log.Errorf("Unknown message type: %s", msgContent.MessageType)
		return
	}

	return
}

// PrivateMsgHandler 处理私聊消息
func PrivateMsgHandler(commonMsg *proto.CommonMsg, msg []byte) (err error) {
	var privateMsg proto.Message

	err = json.Unmarshal(msg, &privateMsg)
	if err != nil {
		common.Log.Errorf("Failed to unmarshal private message: %v", err)
		return
	}

	common.Log.Infof("Received private message: %s", privateMsg.Message)

	PluginHandler(privateMsg.MessageType, msg)
	return
}

// GroupMsgHandler 处理群消息
func GroupMsgHandler(commonMsg *proto.CommonMsg, msg []byte) (err error) {
	var groupMsg proto.Message

	err = json.Unmarshal(msg, &groupMsg)
	if err != nil {
		common.Log.Errorf("Failed to unmarshal group message: %v", err)
		return
	}

	common.Log.Infof("Received group message: %s", groupMsg.Message)

	PluginHandler(groupMsg.MessageType, msg)
	return
}
