package helper

import (
	httpInterfaces "cnb.cool/mliev/open/go-web/pkg/server/http_server/interfaces"
	"github.com/muleiwu/gsr"
)

const RequestLoggerKey = "requestLogger"

// GetRequestLogger 从请求上下文中获取带 traceId 的请求级 Logger。
// 若上下文中未设置，则回退到全局 Logger。
func GetRequestLogger(c httpInterfaces.RouterContextInterface) gsr.Logger {
	if v := c.Get(RequestLoggerKey); v != nil {
		if l, ok := v.(gsr.Logger); ok {
			return l
		}
	}
	return GetLogger()
}
