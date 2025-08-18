package config

import (
	envInterface "cnb.cool/mliev/examples/go-web/internal/pkg/env/interfaces"
)

type Config struct {
	env    envInterface.EnvInterface
	Driver string `json:"driver"`
}

func (receiver Config) InitConfig() map[string]any {
	return map[string]any{
		"database.driver":   receiver.env.GetString("database.driver", "mysql"),
		"database.host":     receiver.env.GetString("database.host", "127.0.0.1"),
		"database.port":     receiver.env.GetInt("database.port", 3306),
		"database.dbname":   receiver.env.GetString("database.dbname", "test"),
		"database.username": receiver.env.GetString("database.username", "test"),
		"database.password": receiver.env.GetString("database.password", "123456"),
	}
}
