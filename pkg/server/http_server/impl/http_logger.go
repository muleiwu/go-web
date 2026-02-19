package impl

import (
	"github.com/muleiwu/golog"
	"github.com/muleiwu/gsr"
)

// NewHttpLogger wraps logger with a pre-populated traceId field for request-scoped logging.
// When the underlying logger is a *golog.Logger, golog's With() is used to create an
// efficient child logger; otherwise the original logger is returned unchanged.
func NewHttpLogger(logger gsr.Logger, traceId string) gsr.Logger {
	if gl, ok := logger.(*golog.Logger); ok {
		return gl.With(golog.Field("traceId", traceId))
	}
	return logger
}
