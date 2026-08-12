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
	dialector := newPostgresqlDialector(dc)
	if dialector == nil {
		return nil, fmt.Errorf("database postgresql driver: invalid database config")
	}
	return gorm.Open(dialector, &gorm.Config{})
}

func newPostgresqlDialector(dc *config.DatabaseConfig) gorm.Dialector {
	dsn := dc.GetPostgreSQLDSN()
	if dsn == "" {
		return nil
	}
	return postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: dc.PreferSimpleProtocol,
	})
}
