package migration

import (
	"fmt"

	"cnb.cool/mliev/open/go-web/pkg/container"
	"github.com/muleiwu/gsr"
	"github.com/pressly/goose/v3"
	"gorm.io/gorm"
)

// dialectMap 将框架数据库驱动名映射到 goose 方言
var dialectMap = map[string]string{
	"mysql":      "mysql",
	"postgresql": "postgres",
	"sqlite":     "sqlite3",
}

type Migration struct {
	Dir string // SQL 迁移文件目录，默认 "migrations"
}

func (m *Migration) Run() error {
	dir := m.getDir()

	db := container.MustGet[*gorm.DB]()
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("[migration] 获取 sql.DB 失败: %w", err)
	}

	config := container.MustGet[gsr.Provider]()
	driver := config.GetString("database.driver", "mysql")

	// memory 驱动不支持 goose（内存数据库无法持久化迁移记录）
	if driver == "memory" {
		logger := container.MustGet[gsr.Logger]()
		logger.Warn("[migration] memory 驱动不支持版本化迁移，已跳过")
		return nil
	}

	dialect, ok := dialectMap[driver]
	if !ok {
		return fmt.Errorf("[migration] 不支持的数据库驱动: %s", driver)
	}

	if err := goose.SetDialect(dialect); err != nil {
		return fmt.Errorf("[migration] 设置方言失败: %w", err)
	}

	if err := goose.Up(sqlDB, dir); err != nil {
		return fmt.Errorf("[migration] 执行迁移失败: %w", err)
	}

	logger := container.MustGet[gsr.Logger]()
	logger.Info(fmt.Sprintf("[migration] 迁移完成 (dir=%s, dialect=%s)", dir, dialect))
	return nil
}

func (m *Migration) Stop() error {
	return nil
}

func (m *Migration) getDir() string {
	if m.Dir != "" {
		return m.Dir
	}

	config, err := container.Get[gsr.Provider]()
	if err == nil {
		if dir := config.GetString("database.migration.dir", ""); dir != "" {
			return dir
		}
	}

	return "migrations"
}
