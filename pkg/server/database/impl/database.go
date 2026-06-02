package impl

import (
	"database/sql"
	"fmt"

	"cnb.cool/mliev/open/go-web/pkg/server/database/config"
	"github.com/glebarez/sqlite"
	mysqlDriver "github.com/go-sql-driver/mysql"
	gormmysql "gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getMySQLDSN(host string, port int, username string, password string, dbName string) string {
	return (&config.DatabaseConfig{
		Host:     host,
		Port:     port,
		DBName:   dbName,
		Username: username,
		Password: password,
	}).GetMySQLDSN()
}

func getPostgreSQLDSN(host string, port int, username string, password string, dbName string) string {
	return (&config.DatabaseConfig{
		Host:     host,
		Port:     port,
		DBName:   dbName,
		Username: username,
		Password: password,
	}).GetPostgreSQLDSN()
}

func getSqliteSQLDSN(host string) string {
	return host
}

// NewDatabase 创建数据库连接（保留供兼容，推荐使用 driver 包）
func NewDatabase(driver string, host string, port int, dbName string, username string, password string) (*gorm.DB, error) {
	var dialector gorm.Dialector
	if driver == "postgresql" {
		dialector = postgres.New(postgres.Config{
			DSN:                  getPostgreSQLDSN(host, port, username, password, dbName),
			PreferSimpleProtocol: true,
		})
	} else if driver == "mysql" {
		mysqlCfg := &config.DatabaseConfig{
			Host:     host,
			Port:     port,
			DBName:   dbName,
			Username: username,
			Password: password,
		}
		driverCfg, err := mysqlCfg.MySQLDriverConfig()
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
	} else if driver == "sqlite" {
		dialector = sqlite.Open(getSqliteSQLDSN(host))
	} else if driver == "memory" {
		dialector = sqlite.Open(":memory:")
	} else {
		return nil, fmt.Errorf("invalid driver: %s", driver)
	}

	return gorm.Open(dialector, &gorm.Config{})
}
