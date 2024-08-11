package dto

/**
用户实体
@MYT 20240810
*/

type Author struct {
	Avatar   string `json:"avatar"`
	Bot      bool   `json:"bot"`
	Id       string `json:"id"`
	Username string `json:"username"`
}
