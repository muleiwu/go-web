package assembly

import (
	"sync"

	"cnb.cool/mliev/examples/go-web/helper/redis"
	"cnb.cool/mliev/examples/go-web/internal/pkg/redis/impl"
)

type Redis struct {
}

var (
	redisOnce sync.Once
)

func (receiver Redis) Assembly() {
	redisOnce.Do(func() {
		redis.RedisHelper = impl.NewRedis()
	})
}
