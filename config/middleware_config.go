package config

import (
	"github.com/gin-gonic/gin"
	"mliev.com/template/go-web/app/middleware"
)

type MiddlewareConfig struct {
}

func (m MiddlewareConfig) Get() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		middleware.CorsMiddleware(),
	}
}
