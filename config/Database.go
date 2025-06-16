package config

import "mliev.com/template/go-web/support"

type DatabaseConfig struct {
	Driver   string `json:"driver"`
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	DBName   string `json:"dbname"`
}

func GetDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		Driver:   support.Env("db.driver", "postgresql").(string),
		Host:     support.Env("db.host", "127.0.0.1").(string),
		Port:     support.Env("db.port", 5432).(int),
		DBName:   support.Env("db.dbname", "test").(string),
		Username: support.Env("db.username", "test").(string),
		Password: support.Env("db.password", "123456").(string),
	}
}
