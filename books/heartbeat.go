package books

import (
	"golang.org/x/net/websocket"
	"log"
	"time"
)

/**
心跳检测
@MYT 20240810
*/

// HeartBeat 心跳检测
func HeartBeat() {
	tickerTime := time.NewTicker(time.Millisecond * time.Duration(heartBeatInterval))
	defer tickerTime.Stop()
	LoadMsg := make(map[string]int64)
	LoadMsg["op"] = 1
	for range tickerTime.C {
		LoadMsg["d"] = sequence
		if err := websocket.JSON.Send(socket, LoadMsg); err != nil {
			log.Printf("heartBeat test fail. %v", err)
			return
		}
	}
}
