package assembly

import (
	"reflect"

	"cnb.cool/mliev/open/go-web/pkg/server/env/impl"
	"github.com/muleiwu/gsr"
)

type Env struct {
}

func (receiver *Env) Type() reflect.Type { return reflect.TypeFor[gsr.Enver]() }
func (receiver *Env) DependsOn() []reflect.Type {
	return []reflect.Type{}
}

func (receiver *Env) Assembly() (any, error) {
	return impl.NewEnv(), nil
}
