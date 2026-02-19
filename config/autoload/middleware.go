package autoload

import (
	"cnb.cool/mliev/open/go-web/app/middleware"
	envInterface "cnb.cool/mliev/open/go-web/pkg/interfaces"
	"github.com/gin-gonic/gin"
)

type Middleware struct {
}

func (receiver Middleware) InitConfig(helper envInterface.HelperInterface) map[string]any {
	return map[string]any{
		"http.middleware": []gin.HandlerFunc{
			middleware.CorsMiddleware(helper),
		},
	}
}
