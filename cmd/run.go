package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"cnb.cool/mliev/open/go-web/pkg/container"
	"cnb.cool/mliev/open/go-web/pkg/interfaces"
	"cnb.cool/mliev/open/go-web/pkg/server/reload"
	"github.com/muleiwu/gsr"
)

// Start 启动应用程序。
// 通过 functional options 模式传入配置，例如：
//
//	cmd.Start(
//	    cmd.WithStaticFs(thatFs),
//	    cmd.WithApp(config.App{}),
//	)
func Start(opts ...Option) {
	o := &Options{}
	for _, fn := range opts {
		fn(o)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	for {
		servers := initializeServices(o)

		select {
		case sig := <-sigChan:
			switch sig {
			case syscall.SIGHUP:
				getLogger().Info("收到 SIGHUP 信号，开始重启服务...")
				stopServices(servers)
				reloadConfiguration(o)
				getLogger().Info("正在重新启动服务...")
				time.Sleep(100 * time.Millisecond)
				continue

			case syscall.SIGINT, syscall.SIGTERM:
				getLogger().Info(fmt.Sprintf("收到 %s 信号，开始关闭服务...", sig))
				stopServices(servers)
				getLogger().Info("服务已全部关闭，程序退出")
				return
			}

		case <-reload.GetReloadChan():
			getLogger().Info("收到重启请求，开始重启服务...")
			stopServices(servers)
			reloadConfiguration(o)
			getLogger().Info("正在重新启动服务...")
			time.Sleep(100 * time.Millisecond)
			continue
		}
	}
}

func initializeServices(o *Options) []interfaces.ServerInterface {
	if err := container.RegisterAssemblies(o.App.Assemblies()); err != nil {
		panic(fmt.Sprintf("Assembly dependency error: %v", err))
	}

	// 将静态资源 FS 注入到 config 中
	config, err := container.Get[gsr.Provider]("config")
	if err == nil {
		config.Set("static.fs", o.StaticFs)
	}

	servers := o.App.Servers()
	for _, s := range servers {
		if err := s.Run(); err != nil {
			getLogger().Error(err.Error())
		}
	}

	return servers
}

func stopServices(servers []interfaces.ServerInterface) {
	getLogger().Info("正在停止所有服务...")
	for _, s := range servers {
		if err := s.Stop(); err != nil {
			getLogger().Error(fmt.Sprintf("停止服务失败: %v", err))
		}
	}
	getLogger().Info("所有服务已停止")
}

func reloadConfiguration(o *Options) {
	getLogger().Info("正在重新加载配置...")

	if err := container.ReloadAssemblies(o.App.Assemblies()); err != nil {
		getLogger().Error(fmt.Sprintf("Assembly dependency error: %v", err))
	}
}

// getLogger 从 container 获取 logger，容错处理
func getLogger() gsr.Logger {
	l, err := container.Get[gsr.Logger]("logger")
	if err != nil {
		return nil
	}
	return l
}
