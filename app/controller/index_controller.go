package controller

import (
	"net/http"

	"cnb.cool/mliev/open/go-web/pkg/interfaces"
	httpInterfaces "cnb.cool/mliev/open/go-web/pkg/server/http_server/interfaces"
)

type IndexController struct {
	BaseResponse
}

func (receiver IndexController) GetIndex(c httpInterfaces.RouterContextInterface, helper interfaces.HelperInterface) {
	helper.GetLogger().Info("hello world")
	c.String(http.StatusOK, "你好, 世界")
}
