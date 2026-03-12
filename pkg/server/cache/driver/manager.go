package driver

import (
	"cnb.cool/mliev/open/go-web/pkg/driver"
	"github.com/muleiwu/gsr"
)

// CacheDriverManager 缓存驱动管理器（全局单例）
var CacheDriverManager = driver.NewManager[gsr.Cacher]()

func init() {
	CacheDriverManager.Extend("redis", RedisFactory)
	CacheDriverManager.Extend("memory", MemoryFactory)
	CacheDriverManager.Extend("local", MemoryFactory)
	CacheDriverManager.Extend("none", NoneFactory)
	CacheDriverManager.SetDefault("redis")
}
