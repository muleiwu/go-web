package driver

import (
	"context"
	"fmt"
	"time"

	redisConfig "cnb.cool/mliev/open/go-web/pkg/server/redis/config"
	"github.com/redis/go-redis/v9"
)

// RedisFactory 创建 Redis 客户端连接
// config 应为 *config.RedisConfig
func RedisFactory(cfg any) (*redis.Client, error) {
	rc, ok := cfg.(*redisConfig.RedisConfig)
	if !ok {
		return nil, fmt.Errorf("redis driver: config must be *RedisConfig, got %T", cfg)
	}

	client := redis.NewClient(rc.GetOptions())

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return client, nil
}
