package support

import "github.com/spf13/viper"

var configEnv = make(map[string]interface{})

func Env(name string, def any) any {
	val, ok := configEnv[name]
	if ok {
		return val
	}

	val = viper.Get(name)

	if val == nil {
		val = def
	}

	configEnv[name] = val
	return val
}
