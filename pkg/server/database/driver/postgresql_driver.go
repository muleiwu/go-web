package driver

import (
	"fmt"

	"cnb.cool/mliev/open/go-web/pkg/server/database/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// PostgresqlFactory 创建 PostgreSQL 数据库连接
// config 应为 *config.DatabaseConfig
func PostgresqlFactory(cfg any) (*gorm.DB, error) {
	dc, ok := cfg.(*config.DatabaseConfig)
	if !ok {
		return nil, fmt.Errorf("database postgresql driver: config must be *DatabaseConfig, got %T", cfg)
	}
	dsn := dc.GetPostgreSQLDSN()
	if dsn == "" {
		return nil, fmt.Errorf("database postgresql driver: invalid database config")
	}
	return gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
}
