package assembly

import (
	"cnb.cool/mliev/examples/go-web/helper/logger"
	"cnb.cool/mliev/examples/go-web/internal/impl"
)

type Logger struct {
}

func (receiver Logger) Assembly() {
	logger.LoggerHelper = impl.NewLogger()
}
