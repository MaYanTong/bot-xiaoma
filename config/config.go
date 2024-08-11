package config

import (
	"github.com/BurntSushi/toml"
)

/**
动态变量配置
@MYT 20240810
*/

var (
	confPath = "./config/config.toml"
	Conf     *Config
)

// Init 初始化配置文件中的参数
func Init() (err error) {
	_, err = toml.DecodeFile(confPath, &Conf)
	return
}

// Config 文件动态配置 appId、token
type Config struct {
	AppID string `toml:"appid"`
	Token string `toml:"token"`
}
