package impl

import (
	"github.com/muleiwu/gsr/logger_interface"
)

type HttpLogger struct {
	logger  logger_interface.LoggerInterface
	traceId string
}

func NewHttpLogger(logger logger_interface.LoggerInterface, traceId string) logger_interface.LoggerInterface {
	l := &HttpLogger{
		logger:  logger,
		traceId: traceId,
	}
	return l
}

func (receiver *HttpLogger) Debug(format string, args ...logger_interface.LoggerFieldInterface) {
	args = append(args, NewLoggerField("traceId", receiver.traceId))
	receiver.logger.Debug(format, args...)
}

func (receiver *HttpLogger) Info(format string, args ...logger_interface.LoggerFieldInterface) {
	args = append(args, NewLoggerField("traceId", receiver.traceId))
	receiver.logger.Info(format, args...)
}

func (receiver *HttpLogger) Notice(format string, args ...logger_interface.LoggerFieldInterface) {
	args = append(args, NewLoggerField("traceId", receiver.traceId))
	receiver.logger.Info(format, args...)
}

func (receiver *HttpLogger) Error(format string, args ...logger_interface.LoggerFieldInterface) {
	args = append(args, NewLoggerField("traceId", receiver.traceId))
	receiver.logger.Error(format, args...)
}

func (receiver *HttpLogger) Warn(format string, args ...logger_interface.LoggerFieldInterface) {
	args = append(args, NewLoggerField("traceId", receiver.traceId))
	receiver.logger.Warn(format, args...)
}

func (receiver *HttpLogger) Fatal(format string, args ...logger_interface.LoggerFieldInterface) {
	args = append(args, NewLoggerField("traceId", receiver.traceId))
	receiver.logger.Fatal(format, args...)
}

type LoggerFieldInterface struct {
	Key   string
	Value string
}

func NewLoggerField(key string, value string) logger_interface.LoggerFieldInterface {
	return &LoggerFieldInterface{
		Key:   key,
		Value: value,
	}
}

func (receiver *LoggerFieldInterface) GetKey() string {
	return receiver.Key
}

func (receiver *LoggerFieldInterface) GetValue() any {
	return receiver.Value
}
