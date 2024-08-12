package wss

import (
	"golang.org/x/net/websocket"
	"log"
	"time"
	"xiaoma-bot/books/service"
	"xiaoma-bot/books/utils"
	"xiaoma-bot/dto"
)

/**
websocket监听
@MYT 20240810
*/

const key = "limit_rate"

// Listen 监听websocket
func Listen() {
	log.Printf("start listen...............")
	for true {
		log.Printf(".............listen...............")
		var loadMsg dto.LoadMsg
		if err := websocket.JSON.Receive(socket, &loadMsg); err != nil {
			log.Printf("reset conn. %v", err)
			// 失败重连
			// StartConn(socket, heartBeatInterval, sessionId, sequence)
			return
		}
		log.Printf("receive data: %v", loadMsg)
		client := utils.GetRedisClient()
		userId := loadMsg.Data.Author.Id
		rKey := key + userId
		limiter := service.RateLimiter(client, rKey, 2, 60*time.Second)
		if !limiter {
			// 限流
			resp(loadMsg.Data, "发送过于频繁,请稍后再试")
		} else {
			Select(&loadMsg)
		}
	}
}
