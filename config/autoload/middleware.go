package autoload

import (
	"cnb.cool/mliev/examples/go-web/app/middleware"
	"github.com/gin-gonic/gin"
)

type Middleware struct {
}

func (receiver Middleware) InitConfig() map[string]any {
	return map[string]any{
		"http.middleware": []gin.HandlerFunc{
			middleware.CorsMiddleware(),
		},
	}
}
