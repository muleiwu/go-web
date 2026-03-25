package autoload

import (
	"cnb.cool/mliev/open/go-web/pkg/container"
	"github.com/muleiwu/gsr"
)

type Migration struct {
}

func (receiver Migration) Get() []any {
	return nil
}

func (receiver Migration) InitConfig() map[string]any {
	env := container.MustGet[gsr.Enver]("env")
	return map[string]any{
		"database.migration.dir": env.GetString("database.migration.dir", "migrations"),
	}
}
