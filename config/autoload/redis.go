package autoload

import (
	"cnb.cool/mliev/open/go-web/pkg/container"
	"github.com/muleiwu/gsr"
)

type Redis struct {
}

func (receiver Redis) InitConfig() map[string]any {
	env := container.MustGet[gsr.Enver]("env")
	return map[string]any{
		"redis.host":     env.GetString("redis.host", "localhost"),
		"redis.port":     env.GetInt("redis.port", 6379),
		"redis.password": env.GetString("redis.password", ""),
		"redis.db":       env.GetInt("redis.db", 0),
	}
}
