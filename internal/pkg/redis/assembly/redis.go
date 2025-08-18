package assembly

import (
	"sync"

	"cnb.cool/mliev/examples/go-web/internal/helper"
	"cnb.cool/mliev/examples/go-web/internal/pkg/redis/impl"
)

type Redis struct {
	Helper *helper.Helper
}

var (
	redisOnce sync.Once
)

func (receiver *Redis) Assembly() {
	redisOnce.Do(func() {
		receiver.Helper.SetRedis(impl.NewRedis(receiver.Helper.GetConfig()))
	})
}
