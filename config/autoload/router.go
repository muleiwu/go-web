package autoload

import (
	"cnb.cool/mliev/open/go-web/app/controller"
	httpInterfaces "cnb.cool/mliev/open/go-web/pkg/server/http_server/interfaces"
)

type Router struct {
}

func (receiver Router) InitConfig() map[string]any {
	return map[string]any{
		"http.router": func(router httpInterfaces.RouterInterface) {
			// 首页
			router.GET("/", controller.IndexController{}.GetIndex)

			health := router.Group("/health")
			// 健康检查接口
			health.GET("", controller.HealthController{}.GetHealth)
			health.GET("/simple", controller.HealthController{}.GetHealthSimple)
		},
	}
}
