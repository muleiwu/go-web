package config

type Base struct {
}

func (receiver Base) InitConfig() map[string]any {
	return map[string]any{
		"app.base": map[string]any{
			"app_name": "go-web-app",
		},
		"app.assembly":   Assembly{}.Get(),
		"app.middleware": Middleware{}.Get(),
		"app.migration":  Migration{}.Get(),
	}
}
