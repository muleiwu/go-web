package helper

import (
	"sync"

	"go.uber.org/zap"
)

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
