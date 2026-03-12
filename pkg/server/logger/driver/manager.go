package driver

import (
	"cnb.cool/mliev/open/go-web/pkg/driver"
	"github.com/muleiwu/gsr"
)

// LoggerDriverManager 日志驱动管理器（全局单例）
var LoggerDriverManager = driver.NewManager[gsr.Logger]()

func init() {
	LoggerDriverManager.Extend("development", DevelopmentFactory)
	LoggerDriverManager.Extend("debug", DevelopmentFactory)
	LoggerDriverManager.Extend("production", ProductionFactory)
	LoggerDriverManager.Extend("release", ProductionFactory)
	LoggerDriverManager.SetDefault("development")
}
