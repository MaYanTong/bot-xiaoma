package dto

/**
消息负载实体
@MYT 20240810
*/

type LoadMsg struct {
	Operate  int64  `json:"op"`
	Sequence int64  `json:"s"`
	Type     string `json:"t"`
	Data     Data   `json:"d"`
}
