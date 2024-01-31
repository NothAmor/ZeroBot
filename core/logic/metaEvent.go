package logic

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/NothAmor/ZeroBot/core/common"
	"github.com/NothAmor/ZeroBot/core/proto"
	"github.com/NothAmor/ZeroBot/core/utils/notify"
)

var (
	lastHeartBeatTime *time.Time // 最后一次心跳时间
)

// MetaEventHandler 处理元事件
func MetaEventHandler(commonMsg *proto.CommonMsg, msg []byte) (err error) {
	// 获取元事件类型
	var metaEvent proto.MetaEvent
	err = json.Unmarshal(msg, &metaEvent)
	if err != nil {
		common.Log.Errorf("Failed to unmarshal meta event: %v", err)
		return
	}

	switch metaEvent.MetaEventType {
	case proto.MetaEventTypeHeartbeat:
		// 心跳
		HeartBeatHandler(commonMsg, msg)

	case proto.MetaEventTypeLifecycle:
		// 生命周期
		LifeCycleHandler(commonMsg, msg)

	default:
		common.Log.Errorf("Unknown meta event type: %s", metaEvent.MetaEventType)
		return
	}
	return
}

// LifeCycleHandler 处理生命周期
func LifeCycleHandler(commonMsg *proto.CommonMsg, msg []byte) (err error) {
	var lifeCycle proto.LifecycleMetaEvent

	err = json.Unmarshal(msg, &lifeCycle)
	if err != nil {
		common.Log.Errorf("Failed to unmarshal lifecycle meta event: %v", err)
		return
	}

	switch lifeCycle.SubType {
	case "enable":
		common.Log.Infof("ZeroBot enabled")
	case "disable":
		common.Log.Infof("ZeroBot disabled")
	case "connect":
		common.Log.Infof("cq-http connected")
	default:
		common.Log.Errorf("Unknown lifecycle meta event subtype: %s", lifeCycle.SubType)
		return
	}

	return
}

// HeartBeatHandler 处理心跳
func HeartBeatHandler(commonMsg *proto.CommonMsg, msg []byte) (err error) {
	var heartBeat proto.HeartbeatMetaEvent

	if lastHeartBeatTime == nil {
		timeNow := time.Now()
		lastHeartBeatTime = &timeNow
	}

	err = json.Unmarshal(msg, &heartBeat)
	if err != nil {
		common.Log.Errorf("Failed to unmarshal heartbeat meta event: %v", err)
		return
	}

	hbTime := time.Unix(commonMsg.Time, 0)

	interval := hbTime.Sub(*lastHeartBeatTime)
	if interval > 6*time.Second {
		common.Log.Warnf("Heartbeat interval too long: %v", interval)
	}

	lastHeartBeatTime = &hbTime

	// 检查 cq-http 状态
	if !heartBeat.Status.AppEnabled || !heartBeat.Status.AppGood {
		common.Log.Warnf("cq-http is not enabled or not good")

		if !common.Config.Notify.Enable {
			return
		}
		err = notify.Notify("[Warning] cq-http状态异常", fmt.Sprintf("心跳信息: %+v", heartBeat))
		if err != nil {
			common.Log.Errorf("Failed to send notification: %v", err)
			return
		}
	}

	// 检查 机器人 在线状态
	if !heartBeat.Status.Online {
		common.Log.Warnf("cq-http is offline")

		if !common.Config.Notify.Enable {
			return
		}
		err = notify.Notify("[Warning] cq-http离线", fmt.Sprintf("心跳信息: %+v", heartBeat))
		if err != nil {
			common.Log.Errorf("Failed to send notification: %v", err)
			return
		}
	}

	return
}
