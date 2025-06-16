package config

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

type RedisConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

func GetRedisConfig() RedisConfig {
	return RedisConfig{
		Host:     GetString("redis.host", "localhost"),
		Port:     GetInt("redis.port", 6379),
		Password: GetString("redis.password", ""),
		DB:       GetInt("redis.db", 0),
	}
}

func (rc RedisConfig) GetOptions() *redis.Options {
	return &redis.Options{
		Addr:     fmt.Sprintf("%s:%d", rc.Host, rc.Port),
		Password: rc.Password,
		DB:       rc.DB,
	}
}
