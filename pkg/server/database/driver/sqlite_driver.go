package driver

import (
	"fmt"

	"cnb.cool/mliev/open/go-web/pkg/server/database/config"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// SqliteFactory 创建 SQLite 数据库连接
// config 应为 *config.DatabaseConfig
func SqliteFactory(cfg any) (*gorm.DB, error) {
	dc, ok := cfg.(*config.DatabaseConfig)
	if !ok {
		return nil, fmt.Errorf("database sqlite driver: config must be *DatabaseConfig, got %T", cfg)
	}
	return gorm.Open(sqlite.Open(dc.Host), &gorm.Config{})
}
