package assembly

import (
	"cnb.cool/mliev/open/go-web/pkg/interfaces"
	configImpl "cnb.cool/mliev/open/go-web/pkg/server/config/impl"
)

type Config struct {
	DefaultConfigs []interfaces.InitConfig
}

func (receiver *Config) Name() string        { return "config" }
func (receiver *Config) DependsOn() []string { return []string{"env"} }

func (receiver *Config) Assembly() (any, error) {
	configHelper := configImpl.NewConfig()
	for _, defaultConfig := range receiver.DefaultConfigs {
		initConfigs := defaultConfig.InitConfig()
		for key, val := range initConfigs {
			configHelper.Set(key, val)
		}
	}

	return configHelper, nil
}
