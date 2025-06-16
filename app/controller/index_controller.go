package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type IndexController struct {
	BaseResponse
}

func (receiver IndexController) GetIndex(c *gin.Context) {
	c.String(http.StatusOK, "你好, 世界")
}
