package impl

import "github.com/muleiwu/golog"

// NewLogger creates a golog.Logger based on the server mode.
// In "release" mode a production (JSON, optimized) logger is used;
// otherwise a development (console, human-readable) logger is created.
func NewLogger(mode string) (*golog.Logger, error) {
	if mode == "release" {
		return golog.NewProductionLogger()
	}
	return golog.NewDevelopmentLogger()
}
