package logger

import (
	"cnb.cool/mliev/examples/go-web/internal/interfaces"
	"go.uber.org/zap"
	"sync"
)

var LoggerHelper interfaces.LoggerInterface

var (
	logger     *zap.Logger
	loggerOnce sync.Once
)

// Logger initLogger initializes the logger only once.
func Logger() *zap.Logger {
	loggerOnce.Do(func() {
		logger = zap.NewExample()
	})
	return logger
}
