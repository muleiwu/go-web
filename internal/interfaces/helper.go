package interfaces

import (
	"github.com/muleiwu/gsr/config_interface"
	"github.com/muleiwu/gsr/env_interface"
	"github.com/muleiwu/gsr/logger_interface"
)

type GetHelperInterface interface {
	GetEnv() env_interface.EnvInterface
	GetConfig() config_interface.ConfigInterface
	GetLogger() logger_interface.LoggerInterface
	GetRedis() RedisInterface
	GetDatabase() DatabaseInterface
}

type SetHelperInterface interface {
	SetEnv(env env_interface.EnvInterface)
	SetConfig(config config_interface.ConfigInterface)
	SetLogger(logger logger_interface.LoggerInterface)
	SetRedis(redis RedisInterface)
	SetDatabase(database DatabaseInterface)
}

type HelperInterface interface {
	GetHelperInterface
	SetHelperInterface
}
