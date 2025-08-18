package assembly

import (
	"sync"

	"cnb.cool/mliev/examples/go-web/helper/env"
	"cnb.cool/mliev/examples/go-web/internal/pkg/env/impl"
)

type Env struct {
}

var (
	envOnce sync.Once
)

func (receiver *Env) Assembly() {
	envOnce.Do(func() {
		env.EnvHelper = impl.NewEnv()
	})
}
