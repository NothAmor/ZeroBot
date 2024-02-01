package handler

import (
	"net/http"

	"github.com/NothAmor/ZeroBot/core/common"
	"github.com/NothAmor/ZeroBot/core/logic"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

// WsHandler 处理 WebSocket 连接
func WsHandler(c *gin.Context) {
	// 升级为 WebSocket 连接
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		common.Log.Errorf("Failed to upgrade WebSocket: %v", err)
		return
	}

	common.Conn = conn
	logic.HandleWebSocket(c)
}

// Health 健康检查
func Health(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
