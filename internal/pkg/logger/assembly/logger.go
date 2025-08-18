package assembly

import (
	"sync"

	"cnb.cool/mliev/examples/go-web/internal/helper"
	"cnb.cool/mliev/examples/go-web/internal/pkg/logger/impl"
)

type Logger struct {
	Helper *helper.Helper
}

var (
	loggerOnce sync.Once
)

func (receiver *Logger) Assembly() {
	loggerOnce.Do(func() {
		receiver.Helper.SetLogger(impl.NewLogger())
	})
}
