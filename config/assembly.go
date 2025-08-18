package config

import (
	"cnb.cool/mliev/examples/go-web/config/autoload"
	"cnb.cool/mliev/examples/go-web/internal/interfaces"
	configAssembly "cnb.cool/mliev/examples/go-web/internal/pkg/config/assembly"
	configInterface "cnb.cool/mliev/examples/go-web/internal/pkg/config/interfaces"
	databaseAssembly "cnb.cool/mliev/examples/go-web/internal/pkg/database/assembly"
	envAssembly "cnb.cool/mliev/examples/go-web/internal/pkg/env/assembly"
	loggerAssembly "cnb.cool/mliev/examples/go-web/internal/pkg/logger/assembly"
	redisAssembly "cnb.cool/mliev/examples/go-web/internal/pkg/redis/assembly"
)

type Assembly struct {
}

// Get 注入反转
func (receiver Assembly) Get() []interfaces.AssemblyInterface {
	return []interfaces.AssemblyInterface{
		configAssembly.Config{
			DefaultConfigs: []configInterface.InitConfig{
				autoload.Base{},
			},
		},
		envAssembly.Env{},
		loggerAssembly.Logger{},
		databaseAssembly.Database{},
		redisAssembly.Redis{},
	}
}
