package autoload

import envInterface "cnb.cool/mliev/examples/go-web/pkg/interfaces"

type App struct {
}

func (receiver App) InitConfig(helper envInterface.HelperInterface) map[string]any {
	return map[string]any{
		"app.app_name": helper.GetEnv().GetString("app.app_name", "go-web-app"),
		"app.mode":     helper.GetEnv().GetString("app.mode", "debug"),
	}
}
