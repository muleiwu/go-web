package assembly

import (
	"cnb.cool/mliev/examples/go-web/internal/interfaces"
	configImpl "cnb.cool/mliev/examples/go-web/internal/pkg/config/impl"
)

type Config struct {
	Helper         interfaces.HelperInterface
	DefaultConfigs []interfaces.InitConfig
}

func (receiver *Config) Assembly() error {
	configHelper := configImpl.NewConfig()
	for _, defaultConfig := range receiver.DefaultConfigs {
		initConfigs := defaultConfig.InitConfig(receiver.Helper)
		for key, val := range initConfigs {
			configHelper.Set(key, val)
		}
	}

	receiver.Helper.SetConfig(configHelper)

	return nil
}
