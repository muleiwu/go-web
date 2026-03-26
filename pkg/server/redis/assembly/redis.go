package assembly

import (
	"reflect"

	"cnb.cool/mliev/open/go-web/pkg/container"
	redisConfig "cnb.cool/mliev/open/go-web/pkg/server/redis/config"
	redisDriver "cnb.cool/mliev/open/go-web/pkg/server/redis/driver"
	"github.com/muleiwu/gsr"
	"github.com/redis/go-redis/v9"
)

type Redis struct {
}

func (receiver *Redis) Type() reflect.Type { return reflect.TypeFor[*redis.Client]() }
func (receiver *Redis) DependsOn() []reflect.Type {
	return []reflect.Type{reflect.TypeFor[gsr.Provider]()}
}

func (receiver *Redis) Assembly() (any, error) {
	cfg := container.MustGet[gsr.Provider]()
	rc := redisConfig.NewRedis(cfg)

	client, err := redisDriver.RedisDriverManager.Make("redis", rc)
	if err != nil {
		return nil, err
	}

	return client, nil
}
