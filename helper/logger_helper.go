package helper

import (
	"sync"

	"go.uber.org/zap"
)

var (
	logger     *zap.Logger
	loggerOnce sync.Once
)

func Logger() *zap.Logger {
	loggerOnce.Do(func() {
		logger = zap.NewExample()
	})
	return logger
}
