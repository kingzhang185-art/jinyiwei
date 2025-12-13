package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"sentinel-opinion-monitor/internal/config"
	appLogger "sentinel-opinion-monitor/internal/pkg/logger"
)

var rdb *redis.Client
var ctx = context.Background()

// Init 初始化 Redis 连接
func Init(cfg *config.RedisConfig) error {
	rdb = redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	// 测试连接
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return err
	}

	appLogger.Get().Info("Redis 连接成功",
		zap.String("addr", cfg.Addr),
		zap.Int("db", cfg.DB),
	)

	return nil
}

// GetClient 获取 Redis 客户端
func GetClient() *redis.Client {
	return rdb
}

// GetContext 获取 Redis 上下文
func GetContext() context.Context {
	return ctx
}

// Close 关闭 Redis 连接
func Close() error {
	if rdb != nil {
		return rdb.Close()
	}
	return nil
}

// Set 设置键值对
func Set(key string, value interface{}, expiration time.Duration) error {
	return rdb.Set(ctx, key, value, expiration).Err()
}

// Get 获取值
func Get(key string) (string, error) {
	return rdb.Get(ctx, key).Result()
}

// Delete 删除键
func Delete(key string) error {
	return rdb.Del(ctx, key).Err()
}

