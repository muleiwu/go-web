package autoload

import (
	"cnb.cool/mliev/open/go-web/pkg/helper"
)

type App struct {
}

func (receiver App) InitConfig() map[string]any {
	return map[string]any{
		"app.app_name": helper.GetEnv().GetString("app.app_name", "go-web-app"),
		"app.mode":     helper.GetEnv().GetString("app.mode", "debug"),
	}
}
