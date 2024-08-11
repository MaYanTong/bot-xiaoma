package dto

/**
频道传输数据实体
@MYT 20240810
*/

type Data struct {
	Id                string `json:"id"`
	Author            Author `json:"author"`
	ChannelId         string `json:"channel_id"`
	Content           string `json:"content"`
	GuildId           string `json:"guild_id"`
	Seq               int64  `json:"seq"`
	HeartbeatInterval int64  `json:"heartbeat_interval"`
	SessionId         string `json:"session_id"`
}
