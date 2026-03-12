package driver

import (
	"fmt"

	gocache "github.com/muleiwu/go-cache"
	"github.com/muleiwu/gsr"
	"github.com/redis/go-redis/v9"
)

// RedisFactory 创建基于 Redis 的缓存驱动
// config 应为 *redis.Client
func RedisFactory(config any) (gsr.Cacher, error) {
	client, ok := config.(*redis.Client)
	if !ok {
		return nil, fmt.Errorf("cache redis driver: config must be *redis.Client, got %T", config)
	}
	if client == nil {
		return nil, fmt.Errorf("cache redis driver: redis client is nil")
	}
	return gocache.NewRedis(client), nil
}
