package driver

import (
	"github.com/muleiwu/golog"
	"github.com/muleiwu/gsr"
)

// DevelopmentFactory 创建开发环境日志驱动（控制台，人类可读）
// config 参数被忽略
func DevelopmentFactory(_ any) (gsr.Logger, error) {
	return golog.NewDevelopmentLogger()
}
