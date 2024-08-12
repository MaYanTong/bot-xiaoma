package service

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"xiaoma-bot/config"
)

/**
db操作逻辑
@MYT 20240810
*/

var (
	// 定义一个全局数据库连接对象
	db *sql.DB
)

// DbInit 初始化数据库连接
func DbInit() {
	// 连接数据库的逻辑
	var err error
	dsn := config.Conf.UserName + ":" + config.Conf.PassWord + "@tcp(" + config.Conf.Url + ")/" + config.Conf.DbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error opening connection:", err)
	}
	// 检查连接是否有效
	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging the database:", err)
	}
	fmt.Println("Database connected.")
}

// 直接使用  原生增删改查API  就不封装一次了
