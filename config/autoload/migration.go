package autoload

import (
	"cnb.cool/mliev/open/go-web/pkg/helper"
)

type Migration struct {
}

func (receiver Migration) Get() []any {
	return nil
}

func (receiver Migration) InitConfig() map[string]any {
	return map[string]any{
		"database.migration.dir": helper.GetEnv().GetString("database.migration.dir", "migrations"),
	}
}
