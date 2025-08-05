package config

type BaseConfig struct {
}

func (receiver BaseConfig) Get() map[string]string {
	return map[string]string{
		"app_name": "go-web-app",
	}
}
