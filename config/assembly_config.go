package config

import (
	"cnb.cool/mliev/examples/go-web/internal/interfaces"
	database "cnb.cool/mliev/examples/go-web/internal/pkg/database/assembly"
	env "cnb.cool/mliev/examples/go-web/internal/pkg/env/assembly"
	logger "cnb.cool/mliev/examples/go-web/internal/pkg/logger/assembly"
	redis "cnb.cool/mliev/examples/go-web/internal/pkg/redis/assembly"
)

type AssemblyConfig struct {
}

func (receiver AssemblyConfig) Get() []interfaces.AssemblyInterface {
	return []interfaces.AssemblyInterface{
		env.Env{},
		logger.Logger{},
		database.Database{},
		redis.Redis{},
	}
}
