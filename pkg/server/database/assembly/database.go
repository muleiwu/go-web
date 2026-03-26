package assembly

import (
	"cnb.cool/mliev/open/go-web/pkg/container"
	"cnb.cool/mliev/open/go-web/pkg/server/database/config"
	dbDriver "cnb.cool/mliev/open/go-web/pkg/server/database/driver"
	"github.com/muleiwu/gsr"
)

type Database struct {
}

func (receiver *Database) Name() string        { return "database" }
func (receiver *Database) DependsOn() []string { return []string{"config"} }

func (receiver *Database) Assembly() (any, error) {
	cfg := container.MustGet[gsr.Provider]("config")
	databaseConfig := config.NewConfig(cfg)

	database, err := dbDriver.DatabaseDriverManager.Make(databaseConfig.Driver, databaseConfig)
	if err != nil {
		return nil, err
	}

	return database, nil
}
