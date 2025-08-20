package assembly

import (
	"sync"

	"cnb.cool/mliev/examples/go-web/internal/interfaces"
	"cnb.cool/mliev/examples/go-web/internal/pkg/env/impl"
)

type Env struct {
	Helper interfaces.HelperInterface
}

var (
	envOnce sync.Once
)

func (receiver *Env) Assembly() {
	envOnce.Do(func() {
		receiver.Helper.SetEnv(impl.NewEnv())
	})
}
