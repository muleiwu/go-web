package impl

import (
	"cnb.cool/mliev/examples/go-web/internal/helper"
	"github.com/gin-gonic/gin"
)

type HttpDeps struct {
	helper *helper.Helper
}

func NewHttpDeps(helper *helper.Helper) *HttpDeps {
	return &HttpDeps{}
}

// 使用闭包包装处理函数
func (d *HttpDeps) WrapHandler(handler func(*gin.Context, *helper.Helper)) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(c, d.helper)
	}
}
