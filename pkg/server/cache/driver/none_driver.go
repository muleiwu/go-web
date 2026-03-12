package driver

import (
	gocache "github.com/muleiwu/go-cache"
	"github.com/muleiwu/gsr"
)

// NoneFactory 创建空操作缓存驱动
// config 参数被忽略
func NoneFactory(_ any) (gsr.Cacher, error) {
	return gocache.NewNone(), nil
}
