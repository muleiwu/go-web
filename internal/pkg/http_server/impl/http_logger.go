package impl

import (
	"cnb.cool/mliev/examples/go-web/internal/interfaces"
)

type HttpLogger struct {
	logger  interfaces.LoggerInterface
	traceId string
}

func NewHttpLogger(logger interfaces.LoggerInterface, traceId string) *HttpLogger {
	l := &HttpLogger{
		logger:  logger,
		traceId: traceId,
	}
	return l
}

func (receiver *HttpLogger) Debug(format string, args ...any) {
	receiver.logger.Debug(format, args)
}

func (receiver *HttpLogger) Info(format string, args ...any) {
	receiver.logger.Info(format, args)
}

func (receiver *HttpLogger) Notice(format string, args ...any) {
	receiver.logger.Info(format, args)
}

func (receiver *HttpLogger) Error(format string, args ...any) {
	receiver.logger.Error(format, args)
}

func (receiver *HttpLogger) Warn(format string, args ...any) {
	receiver.logger.Warn(format, args)
}

func (receiver *HttpLogger) Fatal(format string, args ...any) {
	receiver.logger.Fatal(format, args)
}
