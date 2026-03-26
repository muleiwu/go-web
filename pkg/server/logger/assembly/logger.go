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

func (receiver *Logger) Name() string        { return "logger" }
func (receiver *Logger) DependsOn() []string { return []string{"config"} }

func (receiver *Logger) Assembly() (any, error) {
	config := container.MustGet[gsr.Provider]("config")
	mode := config.GetString("app.mode", "debug")

	logger, err := loggerDriver.LoggerDriverManager.Make(mode, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize logger: %w", err)
	}

	return logger, nil
}
