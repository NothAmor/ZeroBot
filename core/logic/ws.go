package logic

import (
	"github.com/NothAmor/ZeroBot/core/common"
	"github.com/gorilla/websocket"
)

func HandleWebSocket(conn *websocket.Conn) func() {
	return func() {
		defer conn.Close()

		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				common.Log.Errorf("Failed to read message from WebSocket: %v", err)
				break
			}

			common.Log.Debugf("Received message from WebSocket: %s", msg)

			// err = conn.WriteMessage(websocket.TextMessage, []byte("Hello from server"))
			// if err != nil {
			// 	log.Println("Failed to write message to WebSocket:", err)
			// 	break
			// }
		}
	}
}
