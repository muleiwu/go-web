package controller

import (
	"net/http"

	"cnb.cool/mliev/open/go-web/pkg/helper"
	httpInterfaces "cnb.cool/mliev/open/go-web/pkg/server/http_server/interfaces"
	"github.com/muleiwu/gsr"
)

type IndexController struct {
	BaseResponse

	cache gsr.Cacher
}

func (receiver IndexController) GetIndex(c httpInterfaces.RouterContextInterface) {
	helper.GetRequestLogger(c).Info("hello world")
	c.String(http.StatusOK, "你好, 世界")
}
