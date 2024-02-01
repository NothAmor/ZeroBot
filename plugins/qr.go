package plugins

import (
	"fmt"

	"github.com/NothAmor/ZeroBot/core/common/plugin"
	"github.com/NothAmor/ZeroBot/core/proto"
	"github.com/NothAmor/ZeroBot/core/utils/qq"
)

type QR struct {
	plugin.PluginInfo
}

func (r *QR) GetPluginInfo() *plugin.PluginInfo {
	return &r.PluginInfo
}

func (r *QR) Init() {
	r.Name = "二维码生成"
	r.Description = "二维码生成"
	r.Usage = "qr,二维码生成"
	r.UsageType = proto.UsageTypeCommand
	r.Rule = proto.RuleTypeAll
}

func (r *QR) Handle(msg *proto.Msg, args []string) {
	if len(args) == 0 {
		qq.ReplyText(msg, "参数为空")
		return
	}

	qq.ReplyImage(msg, fmt.Sprintf("http://qr.nothamor.com/qr?data=%s&type=image&size=150", args[0]))
}
