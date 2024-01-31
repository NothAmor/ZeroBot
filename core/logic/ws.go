package logic

import (
	"github.com/NothAmor/ZeroBot/core/common"
	"github.com/NothAmor/ZeroBot/core/utils/safe"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// HandleWebSocket 处理 WebSocket 连接
func HandleWebSocket(c *gin.Context, conn *websocket.Conn) (err error) {
	closeNotify := c.Writer.CloseNotify()

	safe.Go(HandleMsg(conn, closeNotify))
	return
}

// HandleMsg 处理 WebSocket 消息
func HandleMsg(conn *websocket.Conn, closeNotify <-chan bool) func() {
	return func() {
		defer conn.Close()

		select {
		case <-closeNotify:
			common.Log.Debug("WebSocket closed")
			return
		default:
			_, msg, err := conn.ReadMessage()
			if err != nil {
				common.Log.Errorf("Failed to read message from WebSocket: %v", err)
				return
			}

			common.Log.Debugf("Received message from WebSocket: %s", msg)
		}
	}
}
