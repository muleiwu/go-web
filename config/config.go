package config

import (
	"cnb.cool/mliev/examples/go-web/internal/interfaces"
	databaseConfig "cnb.cool/mliev/examples/go-web/internal/pkg/database/config"
	redisConfig "cnb.cool/mliev/examples/go-web/internal/pkg/redis/config"
)

type Config struct {
}

func (receiver Config) Get() []interfaces.InitConfig {
	return []interfaces.InitConfig{
		Base{},
		databaseConfig.Config{},
		redisConfig.Config{},
	}
}
