package config

import (
	"cnb.cool/mliev/examples/go-web/internal/interfaces"
	configAssembly "cnb.cool/mliev/examples/go-web/internal/pkg/config/assembly"
	databaseAssembly "cnb.cool/mliev/examples/go-web/internal/pkg/database/assembly"
	envAssembly "cnb.cool/mliev/examples/go-web/internal/pkg/env/assembly"
	loggerAssembly "cnb.cool/mliev/examples/go-web/internal/pkg/logger/assembly"
	redisAssembly "cnb.cool/mliev/examples/go-web/internal/pkg/redis/assembly"
)

type Assembly struct {
}

func (receiver Assembly) Get() []interfaces.AssemblyInterface {
	return []interfaces.AssemblyInterface{
		configAssembly.Config{},
		envAssembly.Env{},
		loggerAssembly.Logger{},
		databaseAssembly.Database{},
		redisAssembly.Redis{},
	}
}
