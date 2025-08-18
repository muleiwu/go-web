package assembly

import (
	"sync"

	"cnb.cool/mliev/examples/go-web/helper/config"
	configImpl "cnb.cool/mliev/examples/go-web/internal/pkg/config/impl"
)

type Config struct {
	DefaultConfig map[string]any
}

var (
	configOnce sync.Once
)

func (receiver Config) Assembly() {
	configOnce.Do(func() {
		config.ConfigHelper = configImpl.NewConfig()
	})
}
