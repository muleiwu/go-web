package assembly

import (
	"fmt"
	"reflect"

	"cnb.cool/mliev/open/go-web/pkg/container"
	loggerDriver "cnb.cool/mliev/open/go-web/pkg/server/logger/driver"
	"github.com/muleiwu/gsr"
)

// Logger is the assembly component responsible for initializing the logger driver.
type Logger struct {
}

func (receiver *Logger) Type() reflect.Type { return reflect.TypeFor[gsr.Logger]() }
func (receiver *Logger) DependsOn() []reflect.Type {
	return []reflect.Type{reflect.TypeFor[gsr.Provider]()}
}

func (receiver *Logger) Assembly() (any, error) {
	config := container.MustGet[gsr.Provider]()
	mode := config.GetString("app.mode", "debug")

	logger, err := loggerDriver.LoggerDriverManager.Make(mode, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize logger: %w", err)
	}

	return logger, nil
}
