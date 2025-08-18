package config

import (
	"cnb.cool/mliev/examples/go-web/config/autoload"
	"cnb.cool/mliev/examples/go-web/internal/helper"
	"cnb.cool/mliev/examples/go-web/internal/interfaces"
	configAssembly "cnb.cool/mliev/examples/go-web/internal/pkg/config/assembly"
	configInterface "cnb.cool/mliev/examples/go-web/internal/pkg/config/interfaces"
	databaseAssembly "cnb.cool/mliev/examples/go-web/internal/pkg/database/assembly"
	envAssembly "cnb.cool/mliev/examples/go-web/internal/pkg/env/assembly"
	loggerAssembly "cnb.cool/mliev/examples/go-web/internal/pkg/logger/assembly"
	redisAssembly "cnb.cool/mliev/examples/go-web/internal/pkg/redis/assembly"
)

type Assembly struct {
	Helper *helper.Helper
}

// Get 注入反转(确保注入顺序，防止依赖为空或者循环依赖)
func (receiver *Assembly) Get() []interfaces.AssemblyInterface {

	return []interfaces.AssemblyInterface{
		&envAssembly.Env{Helper: receiver.Helper}, // 环境变量
		&configAssembly.Config{Helper: receiver.Helper, DefaultConfigs: []configInterface.InitConfig{autoload.Base{}}}, // 代码中的配置(可使用环境变量)
		&loggerAssembly.Logger{Helper: receiver.Helper},                                                                // 日志驱动
		&databaseAssembly.Database{Helper: receiver.Helper},                                                            // 数据库配置
		&redisAssembly.Redis{Helper: receiver.Helper},                                                                  // redis 配置
	}
}
