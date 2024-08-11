package books

import (
	"golang.org/x/net/websocket"
	"log"
	"xiaoma-bot/config"
)

/**
认证鉴权请求
@MYT 20240810
*/

// Auth 发送认证
func Auth() {
	load := make(map[string]interface{})
	data := make(map[string]interface{})
	data["token"] = "Bot " + config.Conf.AppID + "." + config.Conf.Token
	data["intents"] = 1<<0 | 1<<1 | 1<<30
	if sessionId != "" && sequence != 0 {
		data["session_id"] = sessionId
		data["seq"] = sequence
	}
	load["op"] = 2
	load["d"] = data
	if err := websocket.JSON.Send(socket, &load); err != nil {
		log.Printf("auth fail. %v", err)
		return
	}
}
