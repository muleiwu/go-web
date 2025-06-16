package router

import (
	"github.com/gin-gonic/gin"
	"mliev.com/template/go-web/app/Controller"
)

// InitRouter 初始化路由 路由目录 /api/、/sapi/、/v1/、/v2/
func InitRouter(router *gin.Engine) {

	//regexRouter := ginregex.New(router, nil)

	// 健康检查接口
	router.GET("/health", controller.GetHealth)
	router.GET("/health/simple", controller.GetHealthSimple)

	// 首页
	router.GET("/", controller.GetIndex)

	// API路由组
	v1 := router.Group("/api/v1")
	{
		// 这里添加v1版本的API路由
		_ = v1 // 暂时避免未使用变量的警告
	}
}
