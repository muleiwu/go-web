package migration

import (
	"fmt"

	"cnb.cool/mliev/examples/go-web/internal/helper"
)

type Migration struct {
	Helper    *helper.Helper
	Migration []any
}

func (receiver *Migration) Run() error {

	if len(receiver.Migration) > 0 {
		err := receiver.Helper.GetDatabase().AutoMigrate(receiver.Migration...)
		if err != nil {
			return fmt.Errorf("[db migration err:%s]", err.Error())
		}

		receiver.Helper.GetLogger().Info(fmt.Sprintf("[db migration success: %d models migrated]", len(receiver.Migration)))
	}
	return nil
}
