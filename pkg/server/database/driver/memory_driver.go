package driver

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// MemoryFactory 创建内存 SQLite 数据库连接
// config 参数被忽略
func MemoryFactory(_ any) (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
}
