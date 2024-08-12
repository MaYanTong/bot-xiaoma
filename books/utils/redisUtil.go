package utils

import (
	"github.com/go-redis/redis/v8"
	"log"
	"sync"
)

/**
redis连接工具
@MYT 20240810
*/

// RedisClient 是全局Redis客户端的单例
var RedisClient *redis.Client
var once sync.Once

// InitRedisClient 初始化Redis客户端，使用单例模式
func InitRedisClient(options *redis.Options) {
	once.Do(func() {
		log.Println("init redis client...........")
		RedisClient = redis.NewClient(options)
	})
}

// GetRedisClient 返回Redis客户端实例，使用单例模式确保全局唯一
func GetRedisClient() *redis.Client {
	// 如果InitRedisClient没有被调用，或者RedisClient没有被初始化，这里将不会做任何事情
	log.Println("get redis client...........")
	return RedisClient
}
