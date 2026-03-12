package driver

import (
	"github.com/muleiwu/golog"
	"github.com/muleiwu/gsr"
)

// ProductionFactory 创建生产环境日志驱动（JSON，优化性能）
// config 参数被忽略
func ProductionFactory(_ any) (gsr.Logger, error) {
	return golog.NewProductionLogger()
}
