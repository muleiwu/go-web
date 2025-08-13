package helper

import (
	"cnb.cool/mliev/examples/go-web/helper/database"
	"cnb.cool/mliev/examples/go-web/helper/env"
	"cnb.cool/mliev/examples/go-web/helper/logger"
	"cnb.cool/mliev/examples/go-web/helper/redis"
	"cnb.cool/mliev/examples/go-web/internal/interfaces"
)

type Helper struct {
}

func Env() interfaces.EnvInterface {
	return env.EnvHelper
}

func Logger() interfaces.LoggerInterface {
	return logger.LoggerHelper
}

func Redis() interfaces.RedisInterface {
	return redis.RedisHelper
}

func Database() interfaces.DatabaseInterface {
	return database.DatabaseHelper
}
