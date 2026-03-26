package assembly

import (
	"cnb.cool/mliev/open/go-web/pkg/server/env/impl"
)

type Env struct {
}

func (receiver *Env) Name() string        { return "env" }
func (receiver *Env) DependsOn() []string { return nil }

func (receiver *Env) Assembly() (any, error) {
	return impl.NewEnv(), nil
}
