package notify

import (
	GoPusher "github.com/NothAmor/GoPusher"
	structs "github.com/NothAmor/GoPusher/structs"
	"github.com/NothAmor/ZeroBot/core/common"
)

// EmailNotify
func EmailNotify(title, content string) (err error) {
	emailConfig := common.Config.Notify.Email
	mailParams := structs.SmtpRequestStruct{
		Host:     emailConfig.Host,
		Account:  emailConfig.Account,
		Password: emailConfig.Password,
		Port:     emailConfig.Port,
		MailType: "html",
		Sender:   emailConfig.Account,
		SendTo:   emailConfig.Receivers,
		Title:    title,
		Content:  content,
	}

	_, err = GoPusher.Smtp(mailParams)
	if err != nil {
		common.Log.Errorf("Failed to create smtp pusher: %v", err)
		return
	}

	return
}
