package wss

import (
	"encoding/json"
	"golang.org/x/net/websocket"
	"log"
	"xiaoma-bot/books/utils"
	"xiaoma-bot/config"
	"xiaoma-bot/dto"
)

/**
链接启动操作
@MYT 20240810
*/

var (
	socket            *websocket.Conn
	heartBeatInterval int64 = 15000
	sequence          int64
	sessionId         string
)

// StartConn 启动链接
func StartConn() {
	// 初始化链接
	wsUrl := GetWebSocketUrl()

	// 创建链接
	cf, err := websocket.NewConfig(wsUrl, "http://127.0.0.1")
	socket, err = websocket.DialConfig(cf)
	if err != nil {
		log.Printf("init wss conn error. %v", err)
		return
	}

	// 链接状态
	var res dto.LoadMsg
	if err = websocket.JSON.Receive(socket, &res); err != nil {
		log.Printf("wss receive error. %v", err)
		return
	}
	// 重置心跳时间
	heartBeatInterval = res.Data.HeartbeatInterval
	log.Printf("current conn heart beat tiem: %v", heartBeatInterval)
	// 认证
	Auth()
}

type Result struct {
	URL string `json:"url"`
}

// GetWebSocketUrl 获取websocket地址url
func GetWebSocketUrl() string {
	// 构建请求头参数
	header := make(map[string]string)
	header["authorization"] = "Bot " + config.Conf.AppID + "." + config.Conf.Token
	// 远程调用
	log.Printf("start exec http get url.................")
	res := utils.ExecGet("https://sandbox.api.sgroup.qq.com/gateway", nil, header)
	log.Printf("exec http get url result: %v", res)
	// 处理结果
	var result Result
	err := json.Unmarshal(res, &result)
	if err != nil {
		log.Printf("get wss url error. %v", err)
		return ""
	}
	log.Printf("get wss url: %v", result.URL)
	return result.URL
}
