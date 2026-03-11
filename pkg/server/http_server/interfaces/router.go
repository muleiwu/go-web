package interfaces

import (
	"cnb.cool/mliev/open/go-web/pkg/interfaces"
)

// RouterContextInterface 抽象 HTTP 请求/响应上下文，隐藏底层框架细节
type RouterContextInterface interface {
	// 响应
	JSON(code int, obj any)
	String(code int, format string, values ...any)
	// 值存取
	Set(key any, value any)
	GetString(key any) string
	// 路径参数（含正则命名捕获组）
	Param(key string) string
	// 流程控制
	Next()
	Abort()
	AbortWithStatus(code int)
}

// HandlerFunc 是业务 handler 的统一签名，隐藏了 WrapHandler 细节
type HandlerFunc func(RouterContextInterface, interfaces.HelperInterface)

// MiddlewareFunc 是中间件的统一签名
type MiddlewareFunc func(RouterContextInterface)

// RegexRouterInterface 正则路由接口，支持正则表达式模式匹配
type RegexRouterInterface interface {
	GET(pattern string, handlers ...HandlerFunc)
	POST(pattern string, handlers ...HandlerFunc)
	PUT(pattern string, handlers ...HandlerFunc)
	DELETE(pattern string, handlers ...HandlerFunc)
	PATCH(pattern string, handlers ...HandlerFunc)
	HEAD(pattern string, handlers ...HandlerFunc)
	OPTIONS(pattern string, handlers ...HandlerFunc)
	Any(pattern string, handlers ...HandlerFunc)
}

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
	RegexGroup(prefix string) RegexRouterInterface
}
