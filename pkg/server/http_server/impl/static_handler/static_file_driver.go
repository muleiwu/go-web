package static_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// StaticFileDriver 静态文件驱动接口
type StaticFileDriver interface {
	// FileExists 检查文件是否存在
	FileExists(path string) bool

	// GetFS 获取文件系统（用于 gin.StaticFS）
	GetFS(dir string) (http.FileSystem, error)

	// ServeFile 直接提供文件到 gin.Context
	ServeFile(c *gin.Context, dir string, relativePath string) error

	// ServeSPAFallback 把指定目录下的 index.html 作为 SPA fallback 响应给客户端。
	// 内部需直接写出字节流，避免 http.FileServer / http.ServeFile 对 index.html
	// 触发 "./" 301 重定向，从而引发 SPA history 路由的死循环。
	ServeSPAFallback(c *gin.Context, dir string) error

	// GetDriverName 获取驱动名称（用于日志）
	GetDriverName() string
}
