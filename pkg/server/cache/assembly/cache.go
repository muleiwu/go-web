package assembly

import (
	"errors"
	"fmt"

	"cnb.cool/mliev/open/go-web/pkg/container"
	cacheDriver "cnb.cool/mliev/open/go-web/pkg/server/cache/driver"
	"github.com/muleiwu/gsr"
	"github.com/redis/go-redis/v9"
)

type Cache struct {
}

func (receiver *Cache) Name() string        { return "cache" }
func (receiver *Cache) DependsOn() []string { return []string{"config", "logger", "redis"} }

func (receiver *Cache) Assembly() (any, error) {
	cfg := container.MustGet[gsr.Provider]("config")
	logger := container.MustGet[gsr.Logger]("logger")

	driverName := cfg.GetString("cache.driver", "redis")
	logger.Debug("加载缓存驱动" + driverName)

	if driverName == "redis" {
		if _, err := container.Get[*redis.Client]("redis"); err != nil {
			return nil, errors.New("缓存服务驱动配置为：redis，但Redis服务不可用，拒绝启动")
		}
	}

	// 对于 redis 驱动，传递 redis client 作为 config
	var config any
	if driverName == "redis" {
		config = container.MustGet[*redis.Client]("redis")
	}

	cacheInstance, err := cacheDriver.CacheDriverManager.Make(driverName, config)
	if err != nil {
		return nil, fmt.Errorf("加载缓存驱动失败: %w", err)
	}

	return cacheInstance, nil
}
