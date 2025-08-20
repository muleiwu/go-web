package autoload

import (
	envInterface "cnb.cool/mliev/examples/go-web/internal/interfaces"
)

type Redis struct {
	env envInterface.EnvInterface
}

func (receiver Redis) InitConfig(env envInterface.EnvInterface) map[string]any {
	return map[string]any{
		"redis.host":     env.GetString("redis.host", "localhost"),
		"redis.port":     env.GetInt("redis.port", 6379),
		"redis.password": env.GetString("redis.password", ""),
		"redis.db":       env.GetInt("redis.db", 0),
	}
}
