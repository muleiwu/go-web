package assembly

import (
	"cnb.cool/mliev/open/go-web/pkg/container"
	"cnb.cool/mliev/open/go-web/pkg/server/env/impl"
)

type Env struct {
}

func (receiver *Env) Assembly() error {
	env := impl.NewEnv()
	container.Register(container.NewSimpleProvider("env", env))
	return nil
}
