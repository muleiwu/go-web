package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"cnb.cool/mliev/examples/go-web/config"
	"cnb.cool/mliev/examples/go-web/helper/migration"
	helper2 "cnb.cool/mliev/examples/go-web/internal/helper"
)

// Start 启动应用程序
func Start() {
	initializeServices()
	// 添加阻塞以保持主程序运行
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c
}

// initializeServices 初始化所有服务
func initializeServices() {

	helper := helper2.NewHelper()

	assembly := config.NewAssembly(helper)
	for _, assemblyInterface := range assembly.Get() {
		assemblyInterface.Assembly()
	}

	// 自动迁移数据库表结构
	err := migration.AutoMigrate(helper)
	haltOnMigrationFailure := helper.GetEnv().GetBool("database.halt_on_migration_failure", true)

	if haltOnMigrationFailure && err != nil {
		helper.GetLogger().Error(fmt.Sprintf("数据库迁移失败: %v", err))
		panic(err)
	}
}
