package logic

import (
	"encoding/json"

	"github.com/NothAmor/ZeroBot/core/common"
	"github.com/NothAmor/ZeroBot/core/proto"
)

// MsgHandler 处理消息
func MsgHandler(rcvMsg *proto.Msg, commonMsg *proto.CommonMsg, msg []byte) (err error) {
	// 获取消息类型
	var msgContent proto.Message

	err = json.Unmarshal(msg, &msgContent)
	if err != nil {
		common.Log.Errorf("Failed to unmarshal private message: %v", err)
		return
	}

	rcvMsg.MessageType = msgContent.MessageType
	rcvMsg.Sender = msgContent.Sender
	rcvMsg.RawMessage = msgContent.RawMessage

	switch msgContent.MessageType {
	case "private":
		// 私聊消息
		PrivateMsgHandler(rcvMsg, commonMsg, msg)

	case "group":
		// 群消息
		GroupMsgHandler(rcvMsg, commonMsg, msg)

	default:
		common.Log.Errorf("Unknown message type: %s", msgContent.MessageType)
		return
	}

	return
}

// PrivateMsgHandler 处理私聊消息
func PrivateMsgHandler(rcvMsg *proto.Msg, commonMsg *proto.CommonMsg, msg []byte) (err error) {
	var privateMsg proto.Message

	err = json.Unmarshal(msg, &privateMsg)
	if err != nil {
		common.Log.Errorf("Failed to unmarshal private message: %v", err)
		return
	}

	rcvMsg.UserID = privateMsg.UserID
	rcvMsg.Message = privateMsg.Message
	common.Log.Infof("收到私聊消息, 发送人: [%d], 消息内容: [%s]", privateMsg.UserID, privateMsg.RawMessage)

	PluginHandler(rcvMsg)
	return
}

// GroupMsgHandler 处理群消息
func GroupMsgHandler(rcvMsg *proto.Msg, commonMsg *proto.CommonMsg, msg []byte) (err error) {
	var groupMsg proto.Message

	err = json.Unmarshal(msg, &groupMsg)
	if err != nil {
		common.Log.Errorf("Failed to unmarshal group message: %v", err)
		return
	}

	rcvMsg.GroupID = groupMsg.GroupID
	rcvMsg.Message = groupMsg.Message
	common.Log.Infof("收到群消息, 群号: [%d], 发送人: [%s(%d)], 消息内容: [%s]", groupMsg.GroupID, groupMsg.Sender.Nickname, groupMsg.Sender.UserID, groupMsg.RawMessage)

	PluginHandler(rcvMsg)
	return
}
