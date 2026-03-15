package impl

import "github.com/gin-gonic/gin"

// routerContext 包装 *gin.Context，使其满足 RouterContextInterface。
// gin.Context.Get 返回 (any, bool)，此适配器将其简化为 Get(key) any。
type routerContext struct {
	*gin.Context
}

func (rc *routerContext) Set(key string, value any) {
	rc.Context.Set(key, value)
}

func (rc *routerContext) Get(key string) any {
	v, _ := rc.Context.Get(key)
	return v
}

func (rc *routerContext) GetString(key string) string {
	return rc.Context.GetString(key)
}

func (rc *routerContext) GetHeader(key string) string {
	return rc.Context.GetHeader(key)
}

func (rc *routerContext) SetHeader(key, value string) {
	rc.Context.Writer.Header().Set(key, value)
}

func (rc *routerContext) Method() string {
	return rc.Context.Request.Method
}
