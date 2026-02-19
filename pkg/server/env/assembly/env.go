package assembly

import (
	"cnb.cool/mliev/examples/go-web/pkg/interfaces"
	"cnb.cool/mliev/examples/go-web/pkg/server/env/impl"
)

type Env struct {
	Helper interfaces.HelperInterface
}

func (receiver *Env) Assembly() error {
	receiver.Helper.SetEnv(impl.NewEnv())

	return nil
}
