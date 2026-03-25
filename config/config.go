package config

import (
	"cnb.cool/mliev/open/go-web/config/autoload"
	"cnb.cool/mliev/open/go-web/pkg/interfaces"
)

type Config struct {
}

func (receiver Config) Get() []interfaces.InitConfig {
	return []interfaces.InitConfig{
		autoload.App{},
		autoload.Cache{},
		autoload.Http{},
		autoload.StaticFs{},
		autoload.Database{},
		autoload.Redis{},
		autoload.Migration{},
		autoload.Middleware{},
		autoload.Router{},
	}
}
