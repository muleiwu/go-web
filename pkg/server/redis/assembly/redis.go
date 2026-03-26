package assembly

import (
	"cnb.cool/mliev/open/go-web/pkg/container"
	redisConfig "cnb.cool/mliev/open/go-web/pkg/server/redis/config"
	redisDriver "cnb.cool/mliev/open/go-web/pkg/server/redis/driver"
	"github.com/muleiwu/gsr"
)

type Redis struct {
}

func (receiver *Redis) Name() string        { return "redis" }
func (receiver *Redis) DependsOn() []string { return []string{"config"} }

func (receiver *Redis) Assembly() (any, error) {
	cfg := container.MustGet[gsr.Provider]("config")
	rc := redisConfig.NewRedis(cfg)

	client, err := redisDriver.RedisDriverManager.Make("redis", rc)
	if err != nil {
		return nil, err
	}

	return client, nil
}
