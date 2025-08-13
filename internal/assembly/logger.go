package assembly

import (
	"sync"

	"cnb.cool/mliev/examples/go-web/helper/logger"
	"cnb.cool/mliev/examples/go-web/internal/impl"
)

type Logger struct {
}

var (
	loggerOnce sync.Once
)

func (receiver Logger) Assembly() {
	loggerOnce.Do(func() {
		logger.LoggerHelper = impl.NewLogger()
	})
}
