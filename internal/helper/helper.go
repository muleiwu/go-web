package helper

import (
	"cnb.cool/mliev/examples/go-web/internal/interfaces"
	"github.com/muleiwu/gsr/config_interface"
	"github.com/muleiwu/gsr/env_interface"
	"github.com/muleiwu/gsr/logger_interface"
)

type Helper struct {
	env      env_interface.EnvInterface
	config   config_interface.ConfigInterface
	logger   logger_interface.LoggerInterface
	redis    interfaces.RedisInterface
	database interfaces.DatabaseInterface
}

func NewHelper() interfaces.HelperInterface {
	return &Helper{}
}

func (receiver *Helper) GetEnv() env_interface.EnvInterface {
	return receiver.env
}

func (receiver *Helper) GetConfig() config_interface.ConfigInterface {
	return receiver.config
}

func (receiver *Helper) GetLogger() logger_interface.LoggerInterface {
	return receiver.logger
}

func (receiver *Helper) GetRedis() interfaces.RedisInterface {
	return receiver.redis
}

func (receiver *Helper) GetDatabase() interfaces.DatabaseInterface {
	return receiver.database
}

func (receiver *Helper) SetEnv(env env_interface.EnvInterface) {
	receiver.env = env
}

func (receiver *Helper) SetConfig(config config_interface.ConfigInterface) {
	receiver.config = config
}

func (receiver *Helper) SetLogger(logger logger_interface.LoggerInterface) {
	receiver.logger = logger
}

func (receiver *Helper) SetRedis(redis interfaces.RedisInterface) {
	receiver.redis = redis
}

func (receiver *Helper) SetDatabase(database interfaces.DatabaseInterface) {
	receiver.database = database
}
