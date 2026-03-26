package autoload

import (
	"cnb.cool/mliev/open/go-web/pkg/helper"
)

type Redis struct {
}

func (receiver Redis) InitConfig() map[string]any {
	return map[string]any{
		"redis.host":     helper.GetEnv().GetString("redis.host", "localhost"),
		"redis.port":     helper.GetEnv().GetInt("redis.port", 6379),
		"redis.password": helper.GetEnv().GetString("redis.password", ""),
		"redis.db":       helper.GetEnv().GetInt("redis.db", 0),
	}
}
