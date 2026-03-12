package driver

import (
	"cnb.cool/mliev/open/go-web/pkg/driver"
	"github.com/redis/go-redis/v9"
)

// RedisDriverManager Redis 驱动管理器（全局单例）
var RedisDriverManager = driver.NewManager[*redis.Client]()

func init() {
	RedisDriverManager.Extend("redis", RedisFactory)
	RedisDriverManager.SetDefault("redis")
}
