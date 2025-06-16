package logger

import (
	"go.uber.org/zap"
	"sync"
)

var logger *zap.Logger
var once sync.Once

func InitLogger() *zap.Logger {
	once.Do(func() {
		logger = zap.NewExample()
	})
	return logger
}

func Get() *zap.Logger {
	return InitLogger()
}
