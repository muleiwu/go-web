package impl

import "github.com/gin-gonic/gin"

// routerContext 包装 *gin.Context，使其满足 RouterContextInterface。
// gin.Context.Get 返回 (any, bool)，此适配器将其简化为 Get(key) any。
type routerContext struct {
	*gin.Context
}

func (rc *routerContext) Set(key any, value any) {
	rc.Context.Set(key.(string), value)
}

func (rc *routerContext) Get(key any) any {
	v, _ := rc.Context.Get(key.(string))
	return v
}

func (rc *routerContext) GetString(key any) string {
	return rc.Context.GetString(key.(string))
}
