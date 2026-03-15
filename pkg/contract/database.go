package contract

import "gorm.io/gorm"

// Database 数据库服务契约（后续可抽象为接口）
type Database = gorm.DB
