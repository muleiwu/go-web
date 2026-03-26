package helper

import (
	"cnb.cool/mliev/open/go-web/pkg/container"
	"github.com/muleiwu/gsr"
)

func GetEnv() gsr.Enver {
	return container.MustGet[gsr.Enver]()
}
