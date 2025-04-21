package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/xiao-en-5970/Goodminton/backend/app/global"
)

func InitRedis() {
	// Redis连接配置
	redisClient := redis.NewClient(&redis.Options{
		Addr:     global.Cfg.Redis.Addr, // Redis服务器地址
		Password: "",               // 密码，没有则为空
		DB:       global.Cfg.Redis.DB,                // 使用默认DB
		PoolSize: global.Cfg.Redis.PoolSize,              // 连接池大小
	})
	
	// 测试连接
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		global.Logger.Fatalf("Failed to connect to Redis: %v", err)
		panic(fmt.Errorf("Failed to connect to Redis: %v", err))
	}
	global.RedisClient = redisClient
	global.Logger.Infoln("Redis connected successfully")
}