package service

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

/**
redis限流工具
@MYT 20240812
*/

// RateLimiter 使用Redis进行固定窗口限流  可以升级使用滑动窗口 lua脚本 分布式锁等
func RateLimiter(rdb *redis.Client, key string, limit int64, interval time.Duration) bool {
	ctx := context.Background()
	// 对Key进行INCR操作，并返回当前值
	result, err := rdb.Incr(ctx, key).Result()
	if err != nil {
		fmt.Println("Error incrementing Redis key:", err)
		return false
	}
	// 检查是否超过限制
	if result > limit {
		// 超过限制，拒绝请求
		return false
	}
	// 设置或更新过期时间
	if err := rdb.Expire(ctx, key, interval).Err(); err != nil {
		fmt.Println("Error setting Redis key expiration:", err)
		return false
	}
	// 请求被允许
	log.Println("pass request.........")
	return true
}
