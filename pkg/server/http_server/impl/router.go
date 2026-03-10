package impl

import (
	httpInterfaces "cnb.cool/mliev/open/go-web/pkg/server/http_server/interfaces"
	"github.com/gin-gonic/gin"
)

type Router struct {
	group gin.IRouter
	deps  *HttpDeps
}

func NewRouter(group gin.IRouter, deps *HttpDeps) *Router {
	return &Router{group: group, deps: deps}
}

func (r *Router) GET(path string, handler httpInterfaces.HandlerFunc) {
	r.group.GET(path, r.deps.WrapHandler(handler))
}

func (r *Router) POST(path string, handler httpInterfaces.HandlerFunc) {
	r.group.POST(path, r.deps.WrapHandler(handler))
}

func (r *Router) PUT(path string, handler httpInterfaces.HandlerFunc) {
	r.group.PUT(path, r.deps.WrapHandler(handler))
}

func (r *Router) DELETE(path string, handler httpInterfaces.HandlerFunc) {
	r.group.DELETE(path, r.deps.WrapHandler(handler))
}

func (r *Router) PATCH(path string, handler httpInterfaces.HandlerFunc) {
	r.group.PATCH(path, r.deps.WrapHandler(handler))
}

func (r *Router) HEAD(path string, handler httpInterfaces.HandlerFunc) {
	r.group.HEAD(path, r.deps.WrapHandler(handler))
}

func (r *Router) OPTIONS(path string, handler httpInterfaces.HandlerFunc) {
	r.group.OPTIONS(path, r.deps.WrapHandler(handler))
}

func (r *Router) Group(path string) httpInterfaces.RouterInterface {
	return &Router{
		group: r.group.Group(path),
		deps:  r.deps,
	}
}

func (r *Router) Use(middleware ...httpInterfaces.MiddlewareFunc) {
	handlers := make([]gin.HandlerFunc, len(middleware))
	for i, m := range middleware {
		fn := m
		handlers[i] = func(c *gin.Context) { fn(c) }
	}
	r.group.Use(handlers...)
}
