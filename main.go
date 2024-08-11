package main

import (
	"xiaoma-bot/books"
	"xiaoma-bot/config"
)

func main() {
	// 初始化配置文件
	_ = config.Init()
	// 启动链接
	books.Initiate()
}

//func main() {
//	test.TestCal()
//}
