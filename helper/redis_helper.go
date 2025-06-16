package helper

import (
	"github.com/redis/go-redis/v9"
	"mliev.com/template/go-web/config"
	"sync"
)

var (
	rdb     *redis.Client
	rdbOnce sync.Once
)

// GetRedis initializes and returns a Redis client.
func GetRedis() *redis.Client {
	rdbOnce.Do(func() {
		redisConfig := config.GetRedisConfig()
		rdb = redis.NewClient(redisConfig.GetOptions())
	})

	return rdb
}
