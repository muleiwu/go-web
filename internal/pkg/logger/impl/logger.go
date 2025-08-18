package impl

import (
	"go.uber.org/zap"
)

type Logger struct {
	logger *zap.Logger
}

func NewLogger() *Logger {
	logger := &Logger{}
	logger.logger = zap.NewExample()
	return logger
}

func (receiver *Logger) Debug(format string, args ...interface{}) {
	receiver.logger.Debug(format)
}

func (receiver *Logger) Info(format string, args ...interface{}) {
	receiver.logger.Info(format)
}

func (receiver *Logger) Notice(format string, args ...interface{}) {
	receiver.logger.Info(format)
}

func (receiver *Logger) Error(format string, args ...interface{}) {
	receiver.logger.Error(format)
}

func (receiver *Logger) Warn(format string, args ...interface{}) {
	receiver.logger.Warn(format)
}

func (receiver *Logger) Fatal(format string, args ...interface{}) {
	receiver.logger.Fatal(format)
}
