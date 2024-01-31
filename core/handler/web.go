package handler

import (
	"github.com/NothAmor/ZeroBot/core/common"
	"github.com/NothAmor/ZeroBot/core/logic"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	err error
)

func WsHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}

	logic.HandleWebSocket(c, conn)
}

func Health(c *gin.Context) {
	common.Log.Println("pong")
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
