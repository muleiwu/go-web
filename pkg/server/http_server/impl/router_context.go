package impl

import (
	"io"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// routerContext 包装 *gin.Context，使其满足 RouterContextInterface。
// gin.Context.Get 返回 (any, bool)，此适配器将其简化为 Get(key) any。
type routerContext struct {
	*gin.Context
}

// ── 响应 ──

func (rc *routerContext) File(filepath string) {
	rc.Context.File(filepath)
}

func (rc *routerContext) Stream(step func(w io.Writer) bool) {
	rc.Context.Stream(step)
}

func (rc *routerContext) Status(code int) {
	rc.Context.Status(code)
}

// ── 请求参数 ──

func (rc *routerContext) QueryArray(key string) []string {
	return rc.Context.QueryArray(key)
}

func (rc *routerContext) DefaultPostForm(key, defaultValue string) string {
	return rc.Context.DefaultPostForm(key, defaultValue)
}

func (rc *routerContext) ShouldBind(obj any) error {
	return rc.Context.ShouldBind(obj)
}

func (rc *routerContext) ShouldBindQuery(obj any) error {
	return rc.Context.ShouldBindQuery(obj)
}

func (rc *routerContext) GetRawData() ([]byte, error) {
	return rc.Context.GetRawData()
}

// ── 文件上传 ──

func (rc *routerContext) FormFile(name string) (*multipart.FileHeader, error) {
	return rc.Context.FormFile(name)
}

func (rc *routerContext) SaveUploadedFile(file *multipart.FileHeader, dst string) error {
	return rc.Context.SaveUploadedFile(file, dst)
}

// ── 值存取 ──

func (rc *routerContext) Set(key string, value any) {
	rc.Context.Set(key, value)
}

func (rc *routerContext) Get(key string) any {
	v, _ := rc.Context.Get(key)
	return v
}

func (rc *routerContext) GetString(key string) string {
	return rc.Context.GetString(key)
}

// ── HTTP 头部 ──

func (rc *routerContext) GetHeader(key string) string {
	return rc.Context.GetHeader(key)
}

func (rc *routerContext) SetHeader(key, value string) {
	rc.Context.Writer.Header().Set(key, value)
}

// ── 请求信息 ──

func (rc *routerContext) Method() string {
	return rc.Context.Request.Method
}

func (rc *routerContext) Path() string {
	return rc.Context.Request.URL.Path
}

func (rc *routerContext) Host() string {
	return rc.Context.Request.Host
}

func (rc *routerContext) Hostname() string {
	host := rc.Context.Request.Host
	if idx := strings.LastIndex(host, ":"); idx != -1 {
		return host[:idx]
	}
	return host
}

func (rc *routerContext) Scheme() string {
	if rc.Context.Request.TLS != nil {
		return "https"
	}
	if scheme := rc.Context.GetHeader("X-Forwarded-Proto"); scheme != "" {
		return scheme
	}
	return "http"
}

func (rc *routerContext) URL() string {
	return rc.Context.Request.RequestURI
}

func (rc *routerContext) ContentType() string {
	return rc.Context.ContentType()
}

func (rc *routerContext) UserAgent() string {
	return rc.Context.Request.UserAgent()
}

func (rc *routerContext) Referer() string {
	return rc.Context.Request.Referer()
}

func (rc *routerContext) RemoteAddr() string {
	return rc.Context.Request.RemoteAddr
}

func (rc *routerContext) IsWebsocket() bool {
	return rc.Context.IsWebsocket()
}

// ── Cookie ──

func (rc *routerContext) SetSameSite(mode http.SameSite) {
	rc.Context.SetSameSite(mode)
}

// ── 错误处理 ──

func (rc *routerContext) Error(err error) {
	_ = rc.Context.Error(err)
}

// ── 响应状态 ──

func (rc *routerContext) Written() bool {
	return rc.Context.Writer.Written()
}

func (rc *routerContext) GetStatus() int {
	return rc.Context.Writer.Status()
}

// ── 流程控制扩展 ──

func (rc *routerContext) AbortWithStatusJSON(code int, obj any) {
	rc.Context.AbortWithStatusJSON(code, obj)
}

// ── 底层原语 ──

func (rc *routerContext) Request() *http.Request {
	return rc.Context.Request
}

func (rc *routerContext) ResponseWriter() http.ResponseWriter {
	return rc.Context.Writer
}
