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

func (receiver *Cache) Assembly() error {
	cfg := container.MustGet[gsr.Provider]("config")
	logger := container.MustGet[gsr.Logger]("logger")

	driverName := cfg.GetString("cache.driver", "redis")
	logger.Debug("加载缓存驱动" + driverName)

	if driverName == "redis" {
		if _, err := container.Get[*redis.Client]("redis"); err != nil {
			panic(errors.New("缓存服务驱动配置为：redis，但Redis服务不可用，拒绝启动"))
		}
	}

	// 对于 redis 驱动，传递 redis client 作为 config
	var config any
	if driverName == "redis" {
		config = container.MustGet[*redis.Client]("redis")
	}

	cacheInstance, err := cacheDriver.CacheDriverManager.Make(driverName, config)
	if err != nil {
		fmt.Printf("[cache] 加载缓存驱动失败: %s\n", err.Error())
		return nil
	}

	container.Register(container.NewSimpleProvider("cache", cacheInstance))
	return nil
}
