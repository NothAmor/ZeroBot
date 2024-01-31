package notify

import "github.com/NothAmor/ZeroBot/core/common"

// Notify 推送通知
func Notify(title, content string) (err error) {
	notifyConfig := common.Config.Notify

	// 检查是否启用推送通知
	if !notifyConfig.Enable {
		return
	}
	if len(notifyConfig.Use) == 0 {
		return
	}

	// 推送通知
	for _, v := range notifyConfig.Use {
		switch v {
		case "server-chan":
			err = ServerChanNotify(title, content)
			if err != nil {
				common.Log.Errorf("Failed to notify via server-chan: %v", err)
				return
			}
		case "email":
			err = EmailNotify(title, content)
			if err != nil {
				common.Log.Errorf("Failed to notify via email: %v", err)
				return
			}
		default:
			common.Log.Errorf("Unknown notify method: %s", v)
			return
		}
	}

	return
}
