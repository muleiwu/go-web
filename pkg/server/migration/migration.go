package migration

import (
	"fmt"

	"cnb.cool/mliev/open/go-web/pkg/container"
	"github.com/muleiwu/gsr"
	"gorm.io/gorm"
)

type Migration struct {
	Migration []any
}

func (receiver *Migration) Run() error {
	if len(receiver.Migration) > 0 {
		db := container.MustGet[*gorm.DB]("database")
		err := db.AutoMigrate(receiver.Migration...)
		if err != nil {
			return fmt.Errorf("[db migration err:%s]", err.Error())
		}

		logger := container.MustGet[gsr.Logger]("logger")
		logger.Info(fmt.Sprintf("[db migration success: %d models migrated]", len(receiver.Migration)))
	}
	return nil
}

// Stop Migration 服务不需要停止操作，空实现
func (receiver *Migration) Stop() error {
	return nil
}
