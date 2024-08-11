package books

import (
	"golang.org/x/net/websocket"
)

/**
链接初始化
@MYT 20240810
*/

var (
	socket            *websocket.Conn
	heartBeatInterval int64 = 15000
	sequence          int64
	sessionId         string
)

// Initiate 启动websocket连接
func Initiate() {
	// 开启连接
	StartConn()
	// 心跳检测
	go HeartBeat()
	// 监听websocket
	go Listen()

	select {}
}
