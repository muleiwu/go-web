package driver

import (
	"database/sql"
	"fmt"

	"cnb.cool/mliev/open/go-web/pkg/server/database/config"
	mysqlDriver "github.com/go-sql-driver/mysql"
	gormmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// MysqlFactory 创建 MySQL 数据库连接
// config 应为 *config.DatabaseConfig
func MysqlFactory(cfg any) (*gorm.DB, error) {
	dc, ok := cfg.(*config.DatabaseConfig)
	if !ok {
		return nil, fmt.Errorf("database mysql driver: config must be *DatabaseConfig, got %T", cfg)
	}
	driverCfg, err := dc.MySQLDriverConfig()
	if err != nil {
		return nil, err
	}
	connector, err := mysqlDriver.NewConnector(driverCfg)
	if err != nil {
		return nil, err
	}
	sqlDB := sql.OpenDB(connector)
	db, err := gorm.Open(gormmysql.New(gormmysql.Config{
		Conn:      sqlDB,
		DSNConfig: driverCfg,
	}), &gorm.Config{})
	if err != nil {
		_ = sqlDB.Close()
		return nil, err
	}
	return db, nil
}
