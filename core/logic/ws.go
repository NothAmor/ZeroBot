package logic

import (
	"github.com/NothAmor/ZeroBot/core/common"
	"github.com/NothAmor/ZeroBot/core/utils/safe"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// HandleWebSocket 处理 WebSocket 连接
func HandleWebSocket(c *gin.Context, conn *websocket.Conn) (err error) {
	var (
		wsMsg []byte
	)
	for {
		// 读取 WebSocket 消息
		_, wsMsg, err = conn.ReadMessage()
		if err != nil {
			common.Log.Errorf("Failed to read message from WebSocket: %v", err)
			return
		}

		// 异步处理消息
		safe.Go(HandleMsg(conn, wsMsg))
	}
}
