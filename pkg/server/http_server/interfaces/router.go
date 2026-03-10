package interfaces

import (
	"cnb.cool/mliev/open/go-web/pkg/interfaces"
)

// Context 抽象 HTTP 请求/响应上下文，隐藏底层框架细节
type Context interface {
	// 响应
	JSON(code int, obj any)
	String(code int, format string, values ...any)
	// 值存取
	Set(key any, value any)
	GetString(key any) string
	// 流程控制
	Next()
	Abort()
	AbortWithStatus(code int)
}

// HandlerFunc 是业务 handler 的统一签名，隐藏了 WrapHandler 细节
type HandlerFunc func(Context, interfaces.HelperInterface)

// MiddlewareFunc 是中间件的统一签名
type MiddlewareFunc func(Context)

// RouterInterface 仿 Gin 路由风格，内部自动 WrapHandler
type RouterInterface interface {
	GET(path string, handler HandlerFunc)
	POST(path string, handler HandlerFunc)
	PUT(path string, handler HandlerFunc)
	DELETE(path string, handler HandlerFunc)
	PATCH(path string, handler HandlerFunc)
	HEAD(path string, handler HandlerFunc)
	OPTIONS(path string, handler HandlerFunc)
	Group(path string) RouterInterface
	Use(middleware ...MiddlewareFunc)
}
