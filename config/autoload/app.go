package autoload

import (
	"cnb.cool/mliev/open/go-web/pkg/container"
	"github.com/muleiwu/gsr"
)

type App struct {
}

func (receiver App) InitConfig() map[string]any {
	env := container.MustGet[gsr.Enver]("env")
	return map[string]any{
		"app.app_name": env.GetString("app.app_name", "go-web-app"),
		"app.mode":     env.GetString("app.mode", "debug"),
	}
}
