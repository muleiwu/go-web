package helper

import (
	"cnb.cool/mliev/open/go-web/pkg/container"
	"github.com/redis/go-redis/v9"
)

func GetRedis() *redis.Client {
	return container.MustGet[*redis.Client]()
}
