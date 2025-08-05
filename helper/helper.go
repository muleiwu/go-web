package helper

import (
	"cnb.cool/mliev/examples/go-web/helper/env"
	"cnb.cool/mliev/examples/go-web/helper/logger"
	"cnb.cool/mliev/examples/go-web/internal/interfaces"
)

type Helper struct {
}

func (receiver Helper) Env() interfaces.EnvInterface {
	return env.EnvHelper
}

func (receiver Helper) Logger() interfaces.LoggerInterface {
	return logger.LoggerHelper
}
