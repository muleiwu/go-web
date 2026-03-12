package driver

import (
	"cnb.cool/mliev/open/go-web/pkg/driver"
	"gorm.io/gorm"
)

// DatabaseDriverManager 数据库驱动管理器（全局单例）
var DatabaseDriverManager = driver.NewManager[*gorm.DB]()

func init() {
	DatabaseDriverManager.Extend("mysql", MysqlFactory)
	DatabaseDriverManager.Extend("postgresql", PostgresqlFactory)
	DatabaseDriverManager.Extend("sqlite", SqliteFactory)
	DatabaseDriverManager.Extend("memory", MemoryFactory)
	DatabaseDriverManager.SetDefault("mysql")
}
