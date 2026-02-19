package assembly

import (
	"fmt"

	"cnb.cool/mliev/open/go-web/pkg/interfaces"
	"cnb.cool/mliev/open/go-web/pkg/server/logger/impl"
)

// Logger is the assembly component responsible for initializing the logger driver.
type Logger struct {
	Helper interfaces.HelperInterface
}

func (receiver *Logger) Assembly() error {
	mode := receiver.Helper.GetConfig().GetString("app.mode", "debug")
	logger, err := impl.NewLogger(mode)
	if err != nil {
		return fmt.Errorf("failed to initialize logger: %w", err)
	}
	receiver.Helper.SetLogger(logger)
	return nil
}
