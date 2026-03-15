package autoload

import (
	"cnb.cool/mliev/open/go-web/app/middleware"
	httpInterfaces "cnb.cool/mliev/open/go-web/pkg/server/http_server/interfaces"
)

type Middleware struct {
}

func (receiver Middleware) InitConfig() map[string]any {
	return map[string]any{
		"http.middleware": []httpInterfaces.MiddlewareFunc{
			middleware.CorsMiddleware(),
		},
	}
}
