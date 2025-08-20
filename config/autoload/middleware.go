package autoload

import (
	"cnb.cool/mliev/examples/go-web/app/middleware"
	envInterface "cnb.cool/mliev/examples/go-web/internal/interfaces"
	"github.com/gin-gonic/gin"
)

type Middleware struct {
}

func (receiver Middleware) InitConfig(env envInterface.EnvInterface) map[string]any {
	return map[string]any{
		"http.middleware": []gin.HandlerFunc{
			middleware.CorsMiddleware(),
		},
	}
}
