package main

import (
	"github.com/go-redis/redis/v8"
	"xiaoma-bot/books/controller"
	"xiaoma-bot/books/service"
	"xiaoma-bot/books/utils"
	"xiaoma-bot/config"
)

func main() {
	// 初始化配置文件
	_ = config.Init()
	// 初始化数据库连接
	service.DbInit()
	// 初始化redis
	options := &redis.Options{
		Addr:     config.Conf.RedisAddr, // Redis地址
		Password: "",                    // 密码
		DB:       0,                     // 使用默认DB
	}
	utils.InitRedisClient(options)
	// 启动链接
	controller.Initiate()
}

//func main() {
//	test.TestCal()
//}
