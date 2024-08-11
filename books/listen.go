package books

import (
	"golang.org/x/net/websocket"
	"log"
	"xiaoma-bot/dto"
)

/**
websocket
@MYT 20240810
*/

// Listen 监听websocket
func Listen() {
	for true {
		var loadMsg dto.LoadMsg
		if err := websocket.JSON.Receive(socket, &loadMsg); err != nil {
			log.Printf("reset conn. %v", err)
			// 失败重连
			// StartConn(socket, heartBeatInterval, sessionId, sequence)
			return
		}
		Select(&loadMsg)
	}
}
