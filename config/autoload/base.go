package autoload

type Base struct {
}

func (receiver Base) InitConfig() map[string]any {
	return map[string]any{
		"app.base.app_name": "go-web-app",
	}
}
