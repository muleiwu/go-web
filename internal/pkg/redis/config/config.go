package config

import (
	envInterface "cnb.cool/mliev/examples/go-web/internal/pkg/env/interfaces"
)

type Config struct {
	env envInterface.EnvInterface
}

func (receiver Config) InitConfig() map[string]any {
	return map[string]any{
		"redis.host":     receiver.env.GetString("redis.host", "localhost"),
		"redis.port":     receiver.env.GetInt("redis.port", 6379),
		"redis.password": receiver.env.GetString("redis.password", ""),
		"redis.db":       receiver.env.GetInt("redis.db", 0),
	}
}
