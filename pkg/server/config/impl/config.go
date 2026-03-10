package impl

import (
	"time"

	"github.com/muleiwu/anyto"
)

type Config struct {
	data map[string]any
}

func NewConfig() *Config {
	return &Config{data: map[string]any{}}
}

func (c *Config) Set(key string, value any) {
	c.data[key] = value
}

func (c *Config) Get(key string, defaultValue any) any {

	data, ok := c.data[key]

	if !ok {
		return defaultValue
	}

	return data
}

func (c *Config) GetBool(key string, defaultValue bool) bool {
	val := c.Get(key, defaultValue)
	result, err := anyto.Anyto(val).To().ValueE().Bool()
	if err != nil {
		return defaultValue
	}
	return result
}

func (c *Config) GetInt(key string, defaultValue int) int {
	val := c.Get(key, defaultValue)
	result, err := anyto.Anyto(val).To().ValueE().Int()
	if err != nil {
		return defaultValue
	}
	return result
}

func (c *Config) GetInt32(key string, defaultValue int32) int32 {
	val := c.Get(key, defaultValue)
	result, err := anyto.Anyto(val).To().ValueE().Int32()
	if err != nil {
		return defaultValue
	}
	return result
}

func (c *Config) GetInt64(key string, defaultValue int64) int64 {
	val := c.Get(key, defaultValue)
	result, err := anyto.Anyto(val).To().ValueE().Int64()
	if err != nil {
		return defaultValue
	}
	return result
}

func (c *Config) GetFloat64(key string, defaultValue float64) float64 {
	val := c.Get(key, defaultValue)
	result, err := anyto.Anyto(val).To().ValueE().Float64()
	if err != nil {
		return defaultValue
	}
	return result
}

func (c *Config) GetStringSlice(key string, defaultValue []string) []string {
	val := c.Get(key, defaultValue)
	result, err := anyto.Anyto(val).To().ValueE().StringSlice()
	if err != nil {
		return defaultValue
	}
	return result
}

func (c *Config) GetString(key string, defaultValue string) string {
	val := c.Get(key, defaultValue)
	result, err := anyto.Anyto(val).To().ValueE().String()
	if err != nil {
		return defaultValue
	}
	return result
}

func (c *Config) GetStringMapString(key string, defaultValue map[string]string) map[string]string {
	val := c.Get(key, defaultValue)
	result, err := anyto.Anyto(val).To().ValueE().StringMapString()
	if err != nil {
		return defaultValue
	}
	return result
}

func (c *Config) GetStringMapStringSlice(key string, defaultValue map[string][]string) map[string][]string {
	val := c.Get(key, defaultValue)
	result, err := anyto.Anyto(val).To().ValueE().StringMapStringSlice()
	if err != nil {
		return defaultValue
	}
	return result
}

func (c *Config) GetTime(key string, defaultValue time.Time) time.Time {
	val := c.Get(key, defaultValue)
	result, err := anyto.Anyto(val).To().ValueE().Time()
	if err != nil {
		return defaultValue
	}
	return result
}
