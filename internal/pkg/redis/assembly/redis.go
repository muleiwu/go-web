package assembly

import (
	"sync"

	"cnb.cool/mliev/examples/go-web/internal/interfaces"
	"cnb.cool/mliev/examples/go-web/internal/pkg/redis/impl"
)

type Redis struct {
	Helper interfaces.HelperInterface
}

var (
	redisOnce sync.Once
)

func (receiver *Redis) Assembly() {
	redisOnce.Do(func() {
		receiver.Helper.SetRedis(impl.NewRedis(receiver.Helper.GetConfig()))
	})
}
