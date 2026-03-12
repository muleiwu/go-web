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

// App 是框架默认的 AppProvider 实现。
// Configs 为 nil 时使用 go-web 内置的默认配置列表；
// Migrations 为 nil 时不执行任何迁移。
// 子项目可通过填充这两个字段来定制行为，无需重写启动逻辑。
type App struct {
	Configs    []interfaces.InitConfig
	Migrations []any
}

func (a App) configs() []interfaces.InitConfig {
	if a.Configs != nil {
		return a.Configs
	}
	return Config{}.Get()
}

func (a App) migrations() []any {
	if a.Migrations != nil {
		return a.Migrations
	}
	return autoload.Migration{}.Get()
}

// Assemblies 返回标准装配链（env -> config -> logger -> database -> redis -> cache）。
func (a App) Assemblies() []interfaces.AssemblyInterface {
	return []interfaces.AssemblyInterface{
		&envAssembly.Env{},
		&configAssembly.Config{
			DefaultConfigs: a.configs(),
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
			Migration: a.migrations(),
		},
		&service.HttpServer{},
	}
}
