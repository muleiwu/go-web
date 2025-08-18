package assembly

import (
	"sync"

	"cnb.cool/mliev/examples/go-web/helper/env"
	"cnb.cool/mliev/examples/go-web/internal/pkg/env/impl"
)

type HttpServer struct {
}

var (
	httpServerOnce sync.Once
)

func (receiver HttpServer) Assembly() {
	httpServerOnce.Do(func() {
		env.EnvHelper = impl.NewEnv()
	})
}
