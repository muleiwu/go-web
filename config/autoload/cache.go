package autoload

import (
	"cnb.cool/mliev/open/go-web/pkg/container"
	"github.com/muleiwu/gsr"
)

type Cache struct {
}

func (receiver Cache) InitConfig() map[string]any {
	env := container.MustGet[gsr.Enver]("env")
	return map[string]any{
		"cache.driver": env.GetString("cache.driver", "none"), // memory,redis,none
	}
}
