package driver

import (
	"time"

	gocache "github.com/muleiwu/go-cache"
	"github.com/muleiwu/gsr"
)

// MemoryFactory 创建基于内存的缓存驱动
// config 参数被忽略
func MemoryFactory(_ any) (gsr.Cacher, error) {
	return gocache.NewMemory(5*time.Minute, 10*time.Minute), nil
}
