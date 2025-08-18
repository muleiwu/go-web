package assembly

import (
	"sync"

	"cnb.cool/mliev/examples/go-web/helper/redis"
	"cnb.cool/mliev/examples/go-web/internal/pkg/config/interfaces"
	"cnb.cool/mliev/examples/go-web/internal/pkg/redis/impl"
)

type Redis struct {
	Config interfaces.ConfigInterface
}

var (
	redisOnce sync.Once
)

func (receiver *Redis) Assembly() {
	redisOnce.Do(func() {
		redis.RedisHelper = impl.NewRedis(receiver.Config)
	})
}
