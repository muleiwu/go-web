package autoload

import envInterface "cnb.cool/mliev/examples/go-web/internal/interfaces"

type Redis struct {
	env envInterface.EnvInterface
}

func (receiver Redis) InitConfig() map[string]any {
	return map[string]any{
		"redis.host":     receiver.env.GetString("redis.host", "localhost"),
		"redis.port":     receiver.env.GetInt("redis.port", 6379),
		"redis.password": receiver.env.GetString("redis.password", ""),
		"redis.db":       receiver.env.GetInt("redis.db", 0),
	}
}
