package interfaces

type LoggerInterface interface {
	Debug(format string, args ...any)
	Info(format string, args ...any)
	Notice(format string, args ...any)
	Error(format string, args ...any)
	Warn(format string, args ...any)
	Fatal(format string, args ...any)
}
