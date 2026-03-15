package helper

import (
	"cnb.cool/mliev/open/go-web/pkg/container"
	"github.com/muleiwu/gsr"
)

func GetLogger() gsr.Logger {
	return container.MustGet[gsr.Logger]("logger")
}
