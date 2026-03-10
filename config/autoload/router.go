package autoload

import (
	"cnb.cool/mliev/open/go-web/app/controller"
	envInterface "cnb.cool/mliev/open/go-web/pkg/interfaces"
	httpInterfaces "cnb.cool/mliev/open/go-web/pkg/server/http_server/interfaces"
)

type Router struct {
}

func (receiver Router) InitConfig(helper envInterface.HelperInterface) map[string]any {
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
