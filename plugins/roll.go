package plugins

import (
	"math/rand"
	"time"

	"github.com/NothAmor/ZeroBot/core/common/plugin"
	"github.com/NothAmor/ZeroBot/core/proto"
	"github.com/NothAmor/ZeroBot/core/utils/qq"
)

type Roll struct {
	plugin.PluginInfo
}

func (r *Roll) GetPluginInfo() *plugin.PluginInfo {
	return &r.PluginInfo
}

func (r *Roll) Init() {
	r.Name = "Roll"
	r.Description = "Roll"
	r.Usage = "roll"
	r.UsageType = proto.UsageTypeCommand
	r.Rule = proto.RuleTypeAll
}

func (r *Roll) Handle(msg *proto.Msg, args []string) {
	if len(args) == 0 {
		qq.ReplyText(msg, "参数为空")
		return
	}

	rand.New(rand.NewSource(time.Now().UnixNano()))
	qq.ReplyText(msg, args[rand.Intn(len(args))])
}
