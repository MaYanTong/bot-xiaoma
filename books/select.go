package books

import (
	"encoding/json"
	"fmt"
	"log"
	"xiaoma-bot/config"
	"xiaoma-bot/dto"
)

/**
事件操作
@MYT 20240810
*/

// Select 事件选择
func Select(loadMsg *dto.LoadMsg) {
	// 操作类型
	operate := loadMsg.Operate
	// 0 分发操作
	if operate == 0 {
		log.Println("opDispatch")
		// 消息序列号
		sequence = loadMsg.Sequence
		// 事件类型
		ty := loadMsg.Type
		if ty == "READY" {
			// 验证成功
			sessionId = loadMsg.Data.SessionId
		} else if ty == "AT_MESSAGE_CREATE" {
			// 走业务逻辑
			res := Compute(loadMsg)
			resp(loadMsg.Data, res)
		}
		// 7 重连操作
	} else if operate == 7 {
		log.Println("opReconnect")
		StartConn()
		// 11 心跳操作
	} else if operate == 11 {
		log.Println("opHeartbeatACK")
	}
}

type msgEntity struct {
	Content string `json:"content"`
	MsgID   string `json:"msg_id"`
}

// resp 响应信息
func resp(data dto.Data, res string) {
	header := make(map[string]string)
	header["authorization"] = "Bot " + config.Conf.AppID + "." + config.Conf.Token
	header["Content-Type"] = "application/json; charset=utf-8"

	var body msgEntity
	body.Content = fmt.Sprintf("<@!%v>\n", data.Author.Id) + res
	body.MsgID = data.Id
	msg, _ := json.Marshal(body)
	ExecPost("https://sandbox.api.sgroup.qq.com/channels/"+data.ChannelId+"/messages", msg, header)
}
