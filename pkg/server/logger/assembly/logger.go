package assembly

import (
	"fmt"

	"cnb.cool/mliev/open/go-web/pkg/container"
	loggerDriver "cnb.cool/mliev/open/go-web/pkg/server/logger/driver"
	"github.com/muleiwu/gsr"
)

// Logger is the assembly component responsible for initializing the logger driver.
type Logger struct {
}

func (receiver *Logger) Assembly() error {
	config := container.MustGet[gsr.Provider]("config")
	mode := config.GetString("app.mode", "debug")

	// 使用 DriverManager 创建日志驱动
	logger, err := loggerDriver.LoggerDriverManager.Make(mode, nil)
	if err != nil {
		return fmt.Errorf("failed to initialize logger: %w", err)
	}

	container.Register(container.NewSimpleProvider("logger", logger))
	return nil
}
