package helper

import (
	"cnb.cool/mliev/open/go-web/pkg/container"
	"github.com/muleiwu/gsr"
)

func GetConfig() gsr.Provider {
	return container.MustGet[gsr.Provider]()
}
