package impl

import (
	"cnb.cool/mliev/examples/go-web/internal/helper"
	"cnb.cool/mliev/examples/go-web/internal/interfaces"
	"github.com/gin-gonic/gin"
)

type HttpDeps struct {
	helper *helper.Helper
}

func NewHttpDeps(helper interfaces.HelperInterface) *HttpDeps {
	return &HttpDeps{}
}

// 使用闭包包装处理函数
func (d *HttpDeps) WrapHandler(handler func(*gin.Context, interfaces.HelperInterface)) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(c, d.helper)
	}
}
