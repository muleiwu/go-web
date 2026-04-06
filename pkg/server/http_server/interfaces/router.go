package interfaces

import (
	"io"
	"mime/multipart"
	"net/http"
)

// RouterContextInterface 抽象 HTTP 请求/响应上下文，隐藏底层框架细节
type RouterContextInterface interface {
	// 响应
	JSON(code int, obj any)
	HTML(code int, name string, obj any)
	String(code int, format string, values ...any)
	Data(code int, contentType string, data []byte)
	Redirect(code int, location string)
	File(filepath string)
	Stream(step func(w io.Writer) bool)
	Status(code int)
	// 请求参数
	Query(key string) string
	DefaultQuery(key, defaultValue string) string
	QueryArray(key string) []string
	PostForm(key string) string
	DefaultPostForm(key, defaultValue string) string
	ShouldBind(obj any) error
	ShouldBindJSON(obj any) error
	ShouldBindQuery(obj any) error
	GetRawData() ([]byte, error)
	// 文件上传
	FormFile(name string) (*multipart.FileHeader, error)
	SaveUploadedFile(file *multipart.FileHeader, dst string) error
	// 值存取
	Set(key string, value any)
	Get(key string) any
	GetString(key string) string
	// 路径参数（含正则命名捕获组）
	Param(key string) string
	// 请求信息
	Path() string
	FullPath() string
	Method() string
	ClientIP() string
	RemoteAddr() string
	Host() string
	Hostname() string
	Scheme() string
	URL() string
	ContentType() string
	UserAgent() string
	Referer() string
	IsWebsocket() bool
	// HTTP 头部
	GetHeader(key string) string
	SetHeader(key, value string)
	// Cookie
	Cookie(name string) (string, error)
	SetCookie(name, value string, maxAge int, path, domain string, secure, httpOnly bool)
	SetSameSite(mode http.SameSite)
	// 错误处理
	Error(err error)
	IsAborted() bool
	// 响应状态
	Written() bool
	GetStatus() int
	// 流程控制
	Next()
	Abort()
	AbortWithStatus(code int)
}

// HandlerFunc 是业务 handler 的统一签名
type HandlerFunc func(RouterContextInterface)

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
	Use(middleware ...HandlerFunc)
	RegexGroup(prefix string, middleware ...HandlerFunc) RegexRouterInterface
}
