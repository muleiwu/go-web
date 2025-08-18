package assembly

import (
	"sync"

	"cnb.cool/mliev/examples/go-web/internal/helper"
	configImpl "cnb.cool/mliev/examples/go-web/internal/pkg/config/impl"
	"cnb.cool/mliev/examples/go-web/internal/pkg/config/interfaces"
)

type Config struct {
	Helper         *helper.Helper
	DefaultConfigs []interfaces.InitConfig
}

var (
	configOnce sync.Once
)

func (receiver *Config) Assembly() {
	configOnce.Do(func() {
		configHelper := configImpl.NewConfig()
		for _, defaultConfig := range receiver.DefaultConfigs {
			initConfigs := defaultConfig.InitConfig()
			for key, val := range initConfigs {
				configHelper.Set(key, val)
			}
		}

		receiver.Helper.SetConfig(configHelper)
	})
}
