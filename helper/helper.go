package helper

import (
	"cnb.cool/mliev/examples/go-web/helper/config"
	"cnb.cool/mliev/examples/go-web/helper/database"
	"cnb.cool/mliev/examples/go-web/helper/env"
	"cnb.cool/mliev/examples/go-web/helper/logger"
	"cnb.cool/mliev/examples/go-web/helper/redis"
	configInterface "cnb.cool/mliev/examples/go-web/internal/pkg/config/interfaces"
	databaseInterface "cnb.cool/mliev/examples/go-web/internal/pkg/database/interfaces"
	envInterface "cnb.cool/mliev/examples/go-web/internal/pkg/env/interfaces"
	loggerInterface "cnb.cool/mliev/examples/go-web/internal/pkg/logger/interfaces"
	redisInterface "cnb.cool/mliev/examples/go-web/internal/pkg/redis/interfaces"
)

type Helper struct {
}

func Env() envInterface.EnvInterface {
	return env.EnvHelper
}

func Config() configInterface.ConfigInterface {
	return config.ConfigHelper
}

func Logger() loggerInterface.LoggerInterface {
	return logger.LoggerHelper
}

func Redis() redisInterface.RedisInterface {
	return redis.RedisHelper
}

func Database() databaseInterface.DatabaseInterface {
	return database.DatabaseHelper
}
