package assembly

import (
	"errors"
	"fmt"
	"reflect"

	"cnb.cool/mliev/open/go-web/pkg/container"
	cacheDriver "cnb.cool/mliev/open/go-web/pkg/server/cache/driver"
	"github.com/muleiwu/gsr"
	"github.com/redis/go-redis/v9"
)

type Cache struct {
}

func (receiver *Cache) Type() reflect.Type { return reflect.TypeFor[gsr.Cacher]() }
func (receiver *Cache) DependsOn() []reflect.Type {
	return []reflect.Type{
		reflect.TypeFor[gsr.Provider](),
		reflect.TypeFor[gsr.Logger](),
		reflect.TypeFor[*redis.Client](),
	}
}

func (receiver *Cache) Assembly() (any, error) {
	cfg := container.MustGet[gsr.Provider]()
	logger := container.MustGet[gsr.Logger]()

	driverName := cfg.GetString("cache.driver", "redis")
	logger.Debug("加载缓存驱动" + driverName)

	if driverName == "redis" {
		if _, err := container.Get[*redis.Client](); err != nil {
			return nil, errors.New("缓存服务驱动配置为：redis，但Redis服务不可用，拒绝启动")
		}
	}

	// 对于 redis 驱动，传递 redis client 作为 config
	var config any
	if driverName == "redis" {
		config = container.MustGet[*redis.Client]()
	}

	cacheInstance, err := cacheDriver.CacheDriverManager.Make(driverName, config)
	if err != nil {
		return nil, fmt.Errorf("加载缓存驱动失败: %w", err)
	}

	return cacheInstance, nil
}
