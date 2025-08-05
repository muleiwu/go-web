package assembly

import (
	"cnb.cool/mliev/examples/go-web/helper/env"
	"cnb.cool/mliev/examples/go-web/internal/impl"
)

type Env struct {
}

func (receiver Env) Assembly() {
	env.EnvHelper = impl.NewEnv()
}
