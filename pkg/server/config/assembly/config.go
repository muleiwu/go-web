package assembly

import (
	"reflect"

	"cnb.cool/mliev/open/go-web/pkg/interfaces"
	configImpl "cnb.cool/mliev/open/go-web/pkg/server/config/impl"
	"github.com/muleiwu/gsr"
)

type Config struct {
	DefaultConfigs []interfaces.InitConfig
}

func (receiver *Config) Type() reflect.Type { return reflect.TypeFor[gsr.Provider]() }
func (receiver *Config) DependsOn() []reflect.Type {
	return []reflect.Type{reflect.TypeFor[gsr.Enver]()}
}

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
