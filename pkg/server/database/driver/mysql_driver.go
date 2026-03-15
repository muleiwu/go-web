package driver

import (
	"fmt"

	"cnb.cool/mliev/open/go-web/pkg/server/database/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// MysqlFactory 创建 MySQL 数据库连接
// config 应为 *config.DatabaseConfig
func MysqlFactory(cfg any) (*gorm.DB, error) {
	dc, ok := cfg.(*config.DatabaseConfig)
	if !ok {
		return nil, fmt.Errorf("database mysql driver: config must be *DatabaseConfig, got %T", cfg)
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dc.Username, dc.Password, dc.Host, dc.Port, dc.DBName)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
