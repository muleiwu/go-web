package config

import (
	"cnb.cool/mliev/open/go-web/config/autoload"
	"cnb.cool/mliev/open/go-web/pkg/interfaces"
	cacheAssembly "cnb.cool/mliev/open/go-web/pkg/server/cache/assembly"
	configAssembly "cnb.cool/mliev/open/go-web/pkg/server/config/assembly"
	databaseAssembly "cnb.cool/mliev/open/go-web/pkg/server/database/assembly"
	envAssembly "cnb.cool/mliev/open/go-web/pkg/server/env/assembly"
	"cnb.cool/mliev/open/go-web/pkg/server/http_server/service"
	loggerAssembly "cnb.cool/mliev/open/go-web/pkg/server/logger/assembly"
	"cnb.cool/mliev/open/go-web/pkg/server/migration"
	redisAssembly "cnb.cool/mliev/open/go-web/pkg/server/redis/assembly"
)

// App 是框架默认的 AppProvider 实现，纯声明式配置。
type App struct{}

// Assemblies 返回标准装配链（env -> config -> logger -> database -> redis -> cache）。
func (a App) Assemblies() []interfaces.AssemblyInterface {
	return []interfaces.AssemblyInterface{
		&envAssembly.Env{},
		&configAssembly.Config{
			DefaultConfigs: Config{}.Get(),
		},
		&loggerAssembly.Logger{},
		&databaseAssembly.Database{},
		&redisAssembly.Redis{},
		&cacheAssembly.Cache{},
	}
}

// Servers 返回标准服务链（migration + http_server）。
func (a App) Servers() []interfaces.ServerInterface {
	return []interfaces.ServerInterface{
		&migration.Migration{
			Migration: autoload.Migration{}.Get(),
		},
		&service.HttpServer{},
	}
}
