package impl

import (
	"reflect"
	"runtime"

	helper2 "cnb.cool/mliev/open/go-web/pkg/helper"
	"cnb.cool/mliev/open/go-web/pkg/interfaces"
	"github.com/gin-gonic/gin"
	"github.com/muleiwu/gsr"
)

// lastHandlerName holds the real controller method name resolved just before
// WrapHandler returns its closure. Route registration is synchronous so this
// is safe without a mutex.
var lastHandlerName string

type HttpDeps struct {
	helper interfaces.HelperInterface
	engine *gin.Engine
}

func NewHttpDeps(helper interfaces.HelperInterface, engine *gin.Engine) *HttpDeps {
	return &HttpDeps{
		helper: helper,
	}
}

// WrapHandler 使用闭包包装处理函数，同时将真实 handler 名称存入 lastHandlerName
// 供 DebugPrintRouteFunc 使用。
func (d *HttpDeps) WrapHandler(handler func(*gin.Context, interfaces.HelperInterface)) gin.HandlerFunc {
	lastHandlerName = runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name()
	return func(c *gin.Context) {
		handler(c, d.getHttpDeps(d.getTraceId(c)))
	}
}

func (d *HttpDeps) getTraceId(c *gin.Context) string {
	return c.GetString("traceId")
}

func (d *HttpDeps) getHttpDeps(traceId string) interfaces.HelperInterface {
	h := &helper2.Helper{}
	h.SetLogger(d.getLogger(d.helper.GetLogger(), traceId))
	h.SetDatabase(d.helper.GetDatabase())
	h.SetRedis(d.helper.GetRedis())
	h.SetConfig(d.helper.GetConfig())
	h.SetEnv(d.helper.GetEnv())
	return h
}

func (d *HttpDeps) getLogger(logger gsr.Logger, traceId string) gsr.Logger {
	return NewHttpLogger(logger, traceId)
}
