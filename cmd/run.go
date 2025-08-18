package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"cnb.cool/mliev/examples/go-web/config"
	"cnb.cool/mliev/examples/go-web/helper"
	"cnb.cool/mliev/examples/go-web/helper/migration"
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

	for _, assemblyInterface := range (config.Assembly{}).Get() {
		assemblyInterface.Assembly()
	}

	// 自动迁移数据库表结构
	err := migration.AutoMigrate()
	haltOnMigrationFailure := helper.Env().GetBool("database.halt_on_migration_failure", true)
	helper.Logger().Error(fmt.Sprintf("数据库迁移失败: %v", err))

	if haltOnMigrationFailure && err != nil {
		panic(err)
	}
}
