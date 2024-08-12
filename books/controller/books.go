package controller

import (
	"xiaoma-bot/books/wss"
)

/**
调用层
@MYT 20240810
*/

// Initiate 启动websocket连接
func Initiate() {
	// 开启连接
	wss.StartConn()
	// 心跳检测
	go wss.HeartBeat()
	// 监听websocket
	go wss.Listen()

	select {}
}
