package helper

import (
	"cnb.cool/mliev/open/go-web/pkg/container"
	"gorm.io/gorm"
)

func GetDatabase() *gorm.DB {
	return container.MustGet[*gorm.DB]("database")
}
