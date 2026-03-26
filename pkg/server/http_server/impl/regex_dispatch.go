package impl

import (
	"net/http"
	"reflect"
	"regexp"
	"runtime"
	"strings"

	"cnb.cool/mliev/open/go-web/pkg/container"
	httpInterfaces "cnb.cool/mliev/open/go-web/pkg/server/http_server/interfaces"
	"github.com/gin-gonic/gin"
	"github.com/muleiwu/golog"
	"github.com/muleiwu/gsr"
)

// regexMatcher 存储一条编译后的正则路由规则
type regexMatcher struct {
	method   string
	re       *regexp.Regexp
	handlers []httpInterfaces.HandlerFunc
}

// RegexRouter 收集正则路由并生成 Gin 通配路由的分发 handler
type RegexRouter struct {
	prefix   string
	router   *Router
	matchers []regexMatcher
	mounted  bool
}

func (rr *RegexRouter) add(method, pattern string, handlers []httpInterfaces.HandlerFunc) {
	re := regexp.MustCompile(pattern)
	rr.matchers = append(rr.matchers, regexMatcher{
		method:   method,
		re:       re,
		handlers: handlers,
	})

	// 打印正则路由注册日志，与 Gin 标准路由日志风格一致
	rr.logRouteRegistration(method, pattern, handlers)

	if !rr.mounted {
		rr.mount()
		rr.mounted = true
	}
}

func (rr *RegexRouter) logRouteRegistration(method, pattern string, handlers []httpInterfaces.HandlerFunc) {
	logger := container.MustGet[gsr.Logger]()
	displayMethod := method
	if displayMethod == "" {
		displayMethod = "ANY"
	}
	handlerName := ""
	if len(handlers) > 0 {
		handlerName = runtime.FuncForPC(reflect.ValueOf(handlers[len(handlers)-1]).Pointer()).Name()
	}
	logger.Info("路由注册",
		golog.Field("method", displayMethod),
		golog.Field("path", pattern),
		golog.Field("handler", handlerName),
		golog.Field("handlers", len(handlers)),
	)
}

// mount 注册 Gin 通配路由，将请求分发到匹配的正则 handler
func (rr *RegexRouter) mount() {
	// 确保 prefix 以 / 结尾后拼接 *any
	wildcard := strings.TrimRight(rr.prefix, "/") + "/*any"

	deps := rr.router.deps
	rr.router.group.Any(wildcard, func(c *gin.Context) {
		path := c.Request.URL.Path
		method := c.Request.Method

		for _, m := range rr.matchers {
			if m.method != "" && m.method != method {
				continue
			}
			matches := m.re.FindStringSubmatch(path)
			if matches == nil {
				continue
			}
			// 将命名捕获组写入 gin.Params
			for i, name := range m.re.SubexpNames() {
				if i == 0 || name == "" {
					continue
				}
				c.Params = append(c.Params, gin.Param{
					Key:   name,
					Value: matches[i],
				})
			}
			// 依次执行 handler 链（通过 WrapHandler 注入请求级 logger）
			for _, h := range m.handlers {
				deps.WrapHandler(h)(c)
			}
			return
		}
		// 无匹配，返回 404
		c.String(http.StatusNotFound, "404 not found")
	})
}

func (rr *RegexRouter) GET(pattern string, handlers ...httpInterfaces.HandlerFunc) {
	rr.add(http.MethodGet, pattern, handlers)
}

func (rr *RegexRouter) POST(pattern string, handlers ...httpInterfaces.HandlerFunc) {
	rr.add(http.MethodPost, pattern, handlers)
}

func (rr *RegexRouter) PUT(pattern string, handlers ...httpInterfaces.HandlerFunc) {
	rr.add(http.MethodPut, pattern, handlers)
}

func (rr *RegexRouter) DELETE(pattern string, handlers ...httpInterfaces.HandlerFunc) {
	rr.add(http.MethodDelete, pattern, handlers)
}

func (rr *RegexRouter) PATCH(pattern string, handlers ...httpInterfaces.HandlerFunc) {
	rr.add(http.MethodPatch, pattern, handlers)
}

func (rr *RegexRouter) HEAD(pattern string, handlers ...httpInterfaces.HandlerFunc) {
	rr.add(http.MethodHead, pattern, handlers)
}

func (rr *RegexRouter) OPTIONS(pattern string, handlers ...httpInterfaces.HandlerFunc) {
	rr.add(http.MethodOptions, pattern, handlers)
}

func (rr *RegexRouter) Any(pattern string, handlers ...httpInterfaces.HandlerFunc) {
	rr.add("", pattern, handlers)
}
