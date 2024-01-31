package notify

import (
	GoPusher "github.com/NothAmor/GoPusher"
	structs "github.com/NothAmor/GoPusher/structs"
	"github.com/NothAmor/ZeroBot/core/common"
)

// ServerChanNotify
func ServerChanNotify(title, content string) (err error) {
	serverChanParams := structs.ServerChanRequestStruct{
		Key:   common.Config.Notify.ServerChan.Key,
		Title: title,
		Desp:  content,
	}

	_, err = GoPusher.ServerChan(serverChanParams)
	if err != nil {
		common.Log.Errorf("Failed to create serverChan pusher: %v", err)
		return
	}

	return
}
