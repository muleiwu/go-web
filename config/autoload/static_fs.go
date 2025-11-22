package autoload

import (
	"embed"

	envInterface "cnb.cool/mliev/examples/go-web/internal/interfaces"
)

type StaticFs struct {
}

func (receiver StaticFs) InitConfig(helper envInterface.HelperInterface) map[string]any {
	return map[string]any{
		"static.fs": map[string]embed.FS{
			// 这里会在启动的时候在 main.go 注入静态资源进来，请在 main.go 添加静态资源
			// 注意：实际的静态资源会在 cmd/run.go 中通过 helper.GetConfig().Set("static.fs", staticFs) 覆盖
		},
	}
}
