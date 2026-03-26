package assembly

import (
	"reflect"

	"cnb.cool/mliev/open/go-web/pkg/container"
	"cnb.cool/mliev/open/go-web/pkg/server/database/config"
	dbDriver "cnb.cool/mliev/open/go-web/pkg/server/database/driver"
	"github.com/muleiwu/gsr"
	"gorm.io/gorm"
)

type Database struct {
}

func (receiver *Database) Type() reflect.Type { return reflect.TypeFor[*gorm.DB]() }
func (receiver *Database) DependsOn() []reflect.Type {
	return []reflect.Type{reflect.TypeFor[gsr.Provider]()}
}

func (receiver *Database) Assembly() (any, error) {
	cfg := container.MustGet[gsr.Provider]()
	databaseConfig := config.NewConfig(cfg)

	database, err := dbDriver.DatabaseDriverManager.Make(databaseConfig.Driver, databaseConfig)
	if err != nil {
		return nil, err
	}

	return database, nil
}
