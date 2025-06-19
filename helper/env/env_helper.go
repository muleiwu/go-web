package env

import (
	"cnb.cool/mliev/examples/go-web/config"
)

// Env 获取环境变量
func Env(name string, def any) any {
	return config.GetEnvWithDefault(name, def)
}

// EnvString 获取字符串环境变量
func EnvString(name string, def string) string {
	return config.GetString(name, def)
}

// EnvInt 获取整数环境变量
func EnvInt(name string, def int) int {
	return config.GetInt(name, def)
}

// EnvBool 获取布尔环境变量
func EnvBool(name string, def bool) bool {
	return config.GetBool(name, def)
}
