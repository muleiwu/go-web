package impl

import (
	"strings"
	"time"
)

type Config struct {
	data map[string]any
}

func NewConfig() *Config {
	return &Config{data: map[string]any{}}
}

func (c *Config) Set(key string, value any) {

}

func (c *Config) Get(key string, defaultValue any) any {
	if !strings.Contains(key, ".") {
		return c.data
	}

	data, ok := c.data[key]

	if !ok {
		return defaultValue
	}

	return data
}

func (c *Config) GetBool(key string, defaultValue bool) bool {
	return c.Get(key, defaultValue).(bool)
}

func (c *Config) GetInt(key string, defaultValue int) int {
	return c.Get(key, defaultValue).(int)
}

func (c *Config) GetInt32(key string, defaultValue int32) int32 {
	return c.Get(key, defaultValue).(int32)
}

func (c *Config) GetInt64(key string, defaultValue int64) int64 {
	return c.Get(key, defaultValue).(int64)
}

func (c *Config) GetFloat64(key string, defaultValue float64) float64 {
	return c.Get(key, defaultValue).(float64)
}

func (c *Config) GetStringSlice(key string, defaultValue []string) []string {
	return c.Get(key, defaultValue).([]string)
}

func (c *Config) GetString(key string, defaultValue string) string {
	return c.Get(key, defaultValue).(string)
}

func (c *Config) GetStringMapString(key string, defaultValue map[string]string) map[string]string {
	return c.Get(key, defaultValue).(map[string]string)
}

func (c *Config) GetStringMapStringSlice(key string, defaultValue map[string][]string) map[string][]string {
	return c.Get(key, defaultValue).(map[string][]string)
}

func (c *Config) GetTime(key string, defaultValue time.Time) time.Time {
	return c.Get(key, defaultValue).(time.Time)
}
