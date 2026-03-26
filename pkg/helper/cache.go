package helper

import (
	"cnb.cool/mliev/open/go-web/pkg/container"
	"github.com/muleiwu/gsr"
)

func GetCache() gsr.Cacher {
	return container.MustGet[gsr.Cacher]()
}
