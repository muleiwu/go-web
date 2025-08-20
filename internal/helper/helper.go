package helper

import (
	"cnb.cool/mliev/examples/go-web/internal/interfaces"
)

type Helper struct {
	env      interfaces.EnvInterface
	config   interfaces.ConfigInterface
	logger   interfaces.LoggerInterface
	redis    interfaces.RedisInterface
	database interfaces.DatabaseInterface
}

func NewHelper() interfaces.HelperInterface {
	return &Helper{}
}

func (receiver *Helper) GetEnv() interfaces.EnvInterface {
	return receiver.env
}

func (receiver *Helper) GetConfig() interfaces.ConfigInterface {
	return receiver.config
}

func (receiver *Helper) GetLogger() interfaces.LoggerInterface {
	return receiver.logger
}

func (receiver *Helper) GetRedis() interfaces.RedisInterface {
	return receiver.redis
}

func (receiver *Helper) GetDatabase() interfaces.DatabaseInterface {
	return receiver.database
}

func (receiver *Helper) SetEnv(env interfaces.EnvInterface) {
	receiver.env = env
}

func (receiver *Helper) SetConfig(config interfaces.ConfigInterface) {
	receiver.config = config
}

func (receiver *Helper) SetLogger(logger interfaces.LoggerInterface) {
	receiver.logger = logger
}

func (receiver *Helper) SetRedis(redis interfaces.RedisInterface) {
	receiver.redis = redis
}

func (receiver *Helper) SetDatabase(database interfaces.DatabaseInterface) {
	receiver.database = database
}
