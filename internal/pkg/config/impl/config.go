package impl

import (
	"errors"
	"strings"
)

type Config struct {
	data map[string]any
}

func NewConfig() *Config {
	return &Config{data: map[string]any{}}
}

func (c *Config) Get(key string) (any, error) {
	if !strings.Contains(key, ".") {
		return c.data, nil
	}

	data, ok := c.data[key]

	if !ok {
		return nil, errors.New("config setting not found")
	}

	return data, nil
}

func (c *Config) Set(key string, value any) {

}
