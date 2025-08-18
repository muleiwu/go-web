package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IndexController struct {
	BaseResponse
}

func (receiver IndexController) GetIndex(c *gin.Context) {
	c.String(http.StatusOK, "你好, 世界")
}
