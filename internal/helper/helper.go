package helper

import (
	envInterface "cnb.cool/mliev/examples/go-web/internal/interfaces"
	configInterface "cnb.cool/mliev/examples/go-web/internal/pkg/config/interfaces"
	databaseInterface "cnb.cool/mliev/examples/go-web/internal/pkg/database/interfaces"
	loggerInterface "cnb.cool/mliev/examples/go-web/internal/pkg/logger/interfaces"
	redisInterface "cnb.cool/mliev/examples/go-web/internal/pkg/redis/interfaces"
)

type Helper struct {
	env      envInterface.EnvInterface
	config   configInterface.ConfigInterface
	logger   loggerInterface.LoggerInterface
	redis    redisInterface.RedisInterface
	database databaseInterface.DatabaseInterface
}

func NewHelper() *Helper {
	return &Helper{}
}

func (receiver *Helper) GetEnv() envInterface.EnvInterface {
	return receiver.env
}

func (receiver *Helper) GetConfig() configInterface.ConfigInterface {
	return receiver.config
}

func (receiver *Helper) GetLogger() loggerInterface.LoggerInterface {
	return receiver.logger
}

func (receiver *Helper) GetRedis() redisInterface.RedisInterface {
	return receiver.redis
}

func (receiver *Helper) GetDatabase() databaseInterface.DatabaseInterface {
	return receiver.database
}

func (receiver *Helper) SetEnv(env envInterface.EnvInterface) {
	receiver.env = env
}

func (receiver *Helper) SetConfig(config configInterface.ConfigInterface) {
	receiver.config = config
}

func (receiver *Helper) SetLogger(logger loggerInterface.LoggerInterface) {
	receiver.logger = logger
}

func (receiver *Helper) SetRedis(redis redisInterface.RedisInterface) {
	receiver.redis = redis
}

func (receiver *Helper) SetDatabase(database databaseInterface.DatabaseInterface) {
	receiver.database = database
}
