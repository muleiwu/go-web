package interfaces

type GetHelperInterface interface {
	GetEnv() EnvInterface
	GetConfig() ConfigInterface
	GetLogger() LoggerInterface
	GetRedis() RedisInterface
	GetDatabase() DatabaseInterface
}

type SetHelperInterface interface {
	SetEnv(env EnvInterface)
	SetConfig(config ConfigInterface)
	SetLogger(logger LoggerInterface)
	SetRedis(redis RedisInterface)
	SetDatabase(database DatabaseInterface)
}

type HelperInterface interface {
	GetHelperInterface
	SetHelperInterface
}
