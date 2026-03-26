package config

import (
	"cnb.cool/mliev/open/go-web/pkg/interfaces"
	configAssembly "cnb.cool/mliev/open/go-web/pkg/server/config/assembly"
	envAssembly "cnb.cool/mliev/open/go-web/pkg/server/env/assembly"
	"cnb.cool/mliev/open/go-web/pkg/server/http_server/service"
	loggerAssembly "cnb.cool/mliev/open/go-web/pkg/server/logger/assembly"
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
		//&databaseAssembly.Database{},
		//&redisAssembly.Redis{},
		//&cacheAssembly.Cache{},
	}
}

// Servers 返回标准服务链（migration + http_server）。
func (a App) Servers() []interfaces.ServerInterface {
	return []interfaces.ServerInterface{
		//&migration.Migration{},
		&service.HttpServer{},
	}
}
