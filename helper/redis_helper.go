package helper

import (
	"cnb.cool/mliev/examples/go-web/config"
	"github.com/redis/go-redis/v9"
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
