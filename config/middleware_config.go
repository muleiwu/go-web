package config

import (
	"cnb.cool/mliev/examples/go-web/app/middleware"
	"github.com/gin-gonic/gin"
)

type MiddlewareConfig struct {
}

func (m MiddlewareConfig) Get() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		middleware.CorsMiddleware(),
	}
}
