package helper

import (
	"cnb.cool/mliev/examples/go-web/helper/database"
	"cnb.cool/mliev/examples/go-web/helper/env"
	"cnb.cool/mliev/examples/go-web/helper/logger"
	"cnb.cool/mliev/examples/go-web/helper/redis"
	interfaces2 "cnb.cool/mliev/examples/go-web/internal/pkg/database/interfaces"
	interfaces3 "cnb.cool/mliev/examples/go-web/internal/pkg/env/interfaces"
	interfaces4 "cnb.cool/mliev/examples/go-web/internal/pkg/logger/interfaces"
	"cnb.cool/mliev/examples/go-web/internal/pkg/redis/interfaces"
)

type Helper struct {
}

func Env() interfaces3.EnvInterface {
	return env.EnvHelper
}

func Logger() interfaces4.LoggerInterface {
	return logger.LoggerHelper
}

func Redis() interfaces.RedisInterface {
	return redis.RedisHelper
}

func Database() interfaces2.DatabaseInterface {
	return database.DatabaseHelper
}
