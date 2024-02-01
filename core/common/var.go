package common

import (
	"github.com/NothAmor/ZeroBot/core/proto"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

var (
	Log    *logrus.Logger  // 日志
	Config *proto.Config   // 系统配置
	Web    *gin.Engine     // Gin Web
	Cron   *cron.Cron      // 定时任务
	Conn   *websocket.Conn // websocket连接

	Plugins []Accessor // 插件列表
)

var VERSION string // 版本号
