package interfaces

type HelperInterface interface {
	GetEnv() EnvInterface
	GetConfig() ConfigInterface
	GetLogger() LoggerInterface
	GetRedis() RedisInterface
	GetDatabase() DatabaseInterface
	SetEnv(env EnvInterface)
	SetConfig(config ConfigInterface)
	SetLogger(logger LoggerInterface)
	SetRedis(redis RedisInterface)
	SetDatabase(database DatabaseInterface)
}
