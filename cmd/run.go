package cmd

import (
	"embed"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	helper2 "cnb.cool/mliev/open/go-web/pkg/helper"
	"cnb.cool/mliev/open/go-web/pkg/interfaces"
	"cnb.cool/mliev/open/go-web/pkg/server/reload"
)

// Start 启动应用程序。
// staticFs 为嵌入的静态资源文件系统（templates、static 等）。
// app 为 AppProvider 实现，由调用方（go-web 自身或子项目）传入，
// 用于提供自定义的装配链和服务链。
func Start(staticFs map[string]embed.FS, app interfaces.AppProvider) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	for {
		helper := helper2.GetHelper()
		servers := initializeServices(staticFs, helper, app)

		select {
		case sig := <-sigChan:
			switch sig {
			case syscall.SIGHUP:
				helper.GetLogger().Info("收到 SIGHUP 信号，开始重启服务...")
				stopServices(servers, helper)
				reloadConfiguration(helper, app)
				helper.GetLogger().Info("正在重新启动服务...")
				time.Sleep(100 * time.Millisecond)
				continue

			case syscall.SIGINT, syscall.SIGTERM:
				helper.GetLogger().Info(fmt.Sprintf("收到 %s 信号，开始关闭服务...", sig))
				stopServices(servers, helper)
				helper.GetLogger().Info("服务已全部关闭，程序退出")
				return
			}

		case <-reload.GetReloadChan():
			helper.GetLogger().Info("收到重启请求，开始重启服务...")
			stopServices(servers, helper)
			reloadConfiguration(helper, app)
			helper.GetLogger().Info("正在重新启动服务...")
			time.Sleep(100 * time.Millisecond)
			continue
		}
	}
}

func initializeServices(staticFs map[string]embed.FS, helper interfaces.HelperInterface, app interfaces.AppProvider) []interfaces.ServerInterface {
	for _, a := range app.Assemblies(helper) {
		if err := a.Assembly(); err != nil {
			fmt.Printf("Error assembling: %v\n", err)
		}
	}

	helper.GetConfig().Set("static.fs", staticFs)

	servers := app.Servers(helper)
	for _, s := range servers {
		if err := s.Run(); err != nil {
			helper.GetLogger().Error(err.Error())
		}
	}

	return servers
}

func stopServices(servers []interfaces.ServerInterface, helper interfaces.HelperInterface) {
	helper.GetLogger().Info("正在停止所有服务...")
	for _, s := range servers {
		if err := s.Stop(); err != nil {
			helper.GetLogger().Error(fmt.Sprintf("停止服务失败: %v", err))
		}
	}
	helper.GetLogger().Info("所有服务已停止")
}

func reloadConfiguration(helper interfaces.HelperInterface, app interfaces.AppProvider) {
	helper.GetLogger().Info("正在重新加载配置...")

	if env := helper.GetEnv(); env != nil {
		type Reloader interface {
			Reload() error
		}
		if reloader, ok := env.(Reloader); ok {
			if err := reloader.Reload(); err != nil {
				helper.GetLogger().Error(fmt.Sprintf("重新加载配置失败: %v", err))
			} else {
				helper.GetLogger().Info("配置已成功重新加载")
			}
		}
	}

	for _, a := range app.Assemblies(helper) {
		if err := a.Assembly(); err != nil {
			helper.GetLogger().Error(fmt.Sprintf("重新装配服务失败: %v", err))
		}
	}
}
