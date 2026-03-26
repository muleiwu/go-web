package main

import (
	"database/sql"
	"fmt"
	"os"

	"cnb.cool/mliev/open/go-web/config"
	"cnb.cool/mliev/open/go-web/pkg/container"
	"cnb.cool/mliev/open/go-web/pkg/interfaces"
	configAssembly "cnb.cool/mliev/open/go-web/pkg/server/config/assembly"
	databaseAssembly "cnb.cool/mliev/open/go-web/pkg/server/database/assembly"
	envAssembly "cnb.cool/mliev/open/go-web/pkg/server/env/assembly"
	loggerAssembly "cnb.cool/mliev/open/go-web/pkg/server/logger/assembly"
	"github.com/muleiwu/gsr"
	"github.com/pressly/goose/v3"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

var dialectMap = map[string]string{
	"mysql":      "mysql",
	"postgresql": "postgres",
	"sqlite":     "sqlite3",
}

func main() {
	rootCmd := &cobra.Command{
		Use:   "migrate",
		Short: "数据库迁移工具",
	}

	rootCmd.AddCommand(upCmd(), downCmd(), statusCmd(), createCmd(), redoCmd())

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func upCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "up",
		Short: "执行所有未应用的迁移",
		RunE: func(cmd *cobra.Command, args []string) error {
			db, dialect, dir, err := bootstrap()
			if err != nil {
				return err
			}
			if err := goose.SetDialect(dialect); err != nil {
				return err
			}
			return goose.Up(db, dir)
		},
	}
}

func downCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "down",
		Short: "回滚最近一次迁移",
		RunE: func(cmd *cobra.Command, args []string) error {
			db, dialect, dir, err := bootstrap()
			if err != nil {
				return err
			}
			if err := goose.SetDialect(dialect); err != nil {
				return err
			}
			return goose.Down(db, dir)
		},
	}
}

func statusCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "查看迁移状态",
		RunE: func(cmd *cobra.Command, args []string) error {
			db, dialect, dir, err := bootstrap()
			if err != nil {
				return err
			}
			if err := goose.SetDialect(dialect); err != nil {
				return err
			}
			return goose.Status(db, dir)
		},
	}
}

func createCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "create [name]",
		Short: "创建新的迁移文件",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			_, _, dir, err := bootstrap()
			if err != nil {
				return err
			}
			return goose.Create(nil, dir, args[0], "sql")
		},
	}
}

func redoCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "redo",
		Short: "回滚并重新执行最近一次迁移",
		RunE: func(cmd *cobra.Command, args []string) error {
			db, dialect, dir, err := bootstrap()
			if err != nil {
				return err
			}
			if err := goose.SetDialect(dialect); err != nil {
				return err
			}
			return goose.Redo(db, dir)
		},
	}
}

// bootstrap 初始化 Env、Config、Logger、Database 组件，返回 sql.DB、方言和迁移目录
func bootstrap() (*sql.DB, string, string, error) {
	assemblies := []interfaces.AssemblyInterface{
		&envAssembly.Env{},
		&configAssembly.Config{
			DefaultConfigs: config.Config{}.Get(),
		},
		&loggerAssembly.Logger{},
		&databaseAssembly.Database{},
	}

	if err := container.RegisterAssemblies(assemblies); err != nil {
		return nil, "", "", err
	}

	cfg := container.MustGet[gsr.Provider]()
	driver := cfg.GetString("database.driver", "mysql")
	dir := cfg.GetString("database.migration.dir", "migrations")

	if driver == "memory" {
		return nil, "", "", fmt.Errorf("memory 驱动不支持版本化迁移")
	}

	dialect, ok := dialectMap[driver]
	if !ok {
		return nil, "", "", fmt.Errorf("不支持的数据库驱动: %s", driver)
	}

	gormDB := container.MustGet[*gorm.DB]()
	sqlDB, err := gormDB.DB()
	if err != nil {
		return nil, "", "", fmt.Errorf("获取 sql.DB 失败: %w", err)
	}

	return sqlDB, dialect, dir, nil
}
