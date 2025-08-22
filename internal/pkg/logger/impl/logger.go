package impl

import (
	"fmt"

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

func (receiver *Logger) Debug(format string, args ...any) {
	receiver.logger.Debug(fmt.Sprintf(format, args))
}

func (receiver *Logger) Info(format string, args ...any) {
	receiver.logger.Info(fmt.Sprintf(format, args))
}

func (receiver *Logger) Notice(format string, args ...any) {
	receiver.logger.Info(fmt.Sprintf(format, args))
}

func (receiver *Logger) Error(format string, args ...any) {
	receiver.logger.Error(fmt.Sprintf(format, args))
}

func (receiver *Logger) Warn(format string, args ...any) {
	receiver.logger.Warn(fmt.Sprintf(format, args))
}

func (receiver *Logger) Fatal(format string, args ...any) {
	receiver.logger.Fatal(fmt.Sprintf(format, args))
}
