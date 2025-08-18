package assembly

import (
	"sync"

	"cnb.cool/mliev/examples/go-web/helper/config"
	configImpl "cnb.cool/mliev/examples/go-web/internal/pkg/config/impl"
	"cnb.cool/mliev/examples/go-web/internal/pkg/config/interfaces"
)

type Config struct {
	DefaultConfigs []interfaces.InitConfig
}

var (
	configOnce sync.Once
)

func (receiver *Config) Assembly() {
	configOnce.Do(func() {
		config.ConfigHelper = configImpl.NewConfig()
		for _, defaultConfig := range receiver.DefaultConfigs {
			initConfigs := defaultConfig.InitConfig()
			for key, val := range initConfigs {
				config.ConfigHelper.Set(key, val)
			}
		}
	})
}
