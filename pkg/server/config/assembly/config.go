package assembly

import (
	"cnb.cool/mliev/open/go-web/pkg/container"
	"cnb.cool/mliev/open/go-web/pkg/interfaces"
	configImpl "cnb.cool/mliev/open/go-web/pkg/server/config/impl"
)

type Config struct {
	DefaultConfigs []interfaces.InitConfig
}

func (receiver *Config) Assembly() error {
	configHelper := configImpl.NewConfig()
	for _, defaultConfig := range receiver.DefaultConfigs {
		initConfigs := defaultConfig.InitConfig()
		for key, val := range initConfigs {
			configHelper.Set(key, val)
		}
	}

	container.Register(container.NewSimpleProvider("config", configHelper))
	return nil
}
