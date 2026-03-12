package assembly

import (
	"cnb.cool/mliev/open/go-web/pkg/container"
	"cnb.cool/mliev/open/go-web/pkg/server/database/config"
	dbDriver "cnb.cool/mliev/open/go-web/pkg/server/database/driver"
	"github.com/muleiwu/gsr"
)

type Database struct {
}

func (receiver *Database) Assembly() error {
	cfg := container.MustGet[gsr.Provider]("config")
	databaseConfig := config.NewConfig(cfg)

	// 使用 DriverManager 创建数据库连接
	database, err := dbDriver.DatabaseDriverManager.Make(databaseConfig.Driver, databaseConfig)
	if err != nil {
		return err
	}

	container.Register(container.NewSimpleProvider("database", database))
	return nil
}
