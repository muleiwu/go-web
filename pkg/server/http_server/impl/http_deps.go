package impl

import (
	"reflect"
	"runtime"

	"cnb.cool/mliev/open/go-web/pkg/container"
	"cnb.cool/mliev/open/go-web/pkg/helper"
	httpInterfaces "cnb.cool/mliev/open/go-web/pkg/server/http_server/interfaces"
	"github.com/gin-gonic/gin"
	"github.com/muleiwu/gsr"
)

type HttpDeps struct {
	LastHandlerName string
}

func NewHttpDeps() *HttpDeps {
	return &HttpDeps{}
}

// WrapHandler 使用闭包包装处理函数，同时将真实 handler 名称存入 lastHandlerName
// 供 DebugPrintRouteFunc 使用。
// 在调用 handler 之前，将请求级 logger（带 traceId）写入上下文。
func (d *HttpDeps) WrapHandler(handler httpInterfaces.HandlerFunc) gin.HandlerFunc {
	d.LastHandlerName = runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name()
	return func(c *gin.Context) {
		// 将请求级 logger 存入上下文
		traceId := c.GetString("traceId")
		baseLogger := container.MustGet[gsr.Logger]()
		c.Set(helper.RequestLoggerKey, NewHttpLogger(baseLogger, traceId))

		handler(&routerContext{c})
	}
}
