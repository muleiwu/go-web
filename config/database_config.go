package config

import (
	"fmt"
)

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
		Driver:   GetString("db.driver", "postgresql"),
		Host:     GetString("db.host", "127.0.0.1"),
		Port:     GetInt("db.port", 5432),
		DBName:   GetString("db.dbname", "test"),
		Username: GetString("db.username", "test"),
		Password: GetString("db.password", "123456"),
	}
}

func (dc DatabaseConfig) GetMySQLDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dc.Username,
		dc.Password,
		dc.Host,
		dc.Port,
		dc.DBName)
}

func (dc DatabaseConfig) GetPostgreSQLDSN() string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
		dc.Username,
		dc.Password,
		dc.Host,
		dc.Port,
		dc.DBName)
}
