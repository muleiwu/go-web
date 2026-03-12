package autoload

type Http struct {
}

func (receiver Http) InitConfig() map[string]any {
	return map[string]any{
		"http.load_static": false,
		"http.static_mode": "embed",
		"http.static_dir":  []string{},
	}
}
