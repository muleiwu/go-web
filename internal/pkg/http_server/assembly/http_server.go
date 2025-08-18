package assembly

import (
	"sync"
)

type HttpServer struct {
}

var (
	httpServerOnce sync.Once
)

func (receiver HttpServer) Assembly() {
	httpServerOnce.Do(func() {

	})
}
