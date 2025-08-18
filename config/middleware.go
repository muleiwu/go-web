package config

import (
	"cnb.cool/mliev/examples/go-web/app/middleware"
	"github.com/gin-gonic/gin"
)

type Middleware struct {
}

func (m Middleware) Get() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		middleware.CorsMiddleware(),
	}
}
