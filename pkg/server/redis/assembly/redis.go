package assembly

import (
	"cnb.cool/mliev/open/go-web/pkg/container"
	redisConfig "cnb.cool/mliev/open/go-web/pkg/server/redis/config"
	redisDriver "cnb.cool/mliev/open/go-web/pkg/server/redis/driver"
	"github.com/muleiwu/gsr"
)

type Redis struct {
}

func (receiver *Redis) Assembly() error {
	cfg := container.MustGet[gsr.Provider]("config")
	rc := redisConfig.NewRedis(cfg)

	// 使用 DriverManager 创建 Redis 连接
	client, err := redisDriver.RedisDriverManager.Make("redis", rc)
	if err != nil {
		return err
	}

	container.Register(container.NewSimpleProvider("redis", client))
	return nil
}
