package impl

import (
	"fmt"
	"strings"
	"time"

	"github.com/muleiwu/anyto"
	"github.com/spf13/viper"
)

type Env struct{}

func NewEnv() *Env {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	replace := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replace)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("警告: 配置文件未找到，将使用默认配置和环境变量: %v\n", err)
	}

	return &Env{}
}

// GetEnvWithDefault 获取环境变量，如果不存在则返回默认值
func (receiver *Env) GetEnvWithDefault(name string, def any) any {
	val := viper.Get(name)
	if val == nil {
		return def
	}
	return val
}

func (receiver *Env) Get(key string, defaultValue any) any {
	return receiver.GetEnvWithDefault(key, defaultValue)
}

func (receiver *Env) GetBool(key string, defaultValue bool) bool {
	val := receiver.Get(key, defaultValue)
	result, err := anyto.Anyto(val).To().ValueE().Bool()
	if err != nil {
		return defaultValue
	}
	return result
}

func (receiver *Env) GetInt(key string, defaultValue int) int {
	val := receiver.Get(key, defaultValue)
	result, err := anyto.Anyto(val).To().ValueE().Int()
	if err != nil {
		return defaultValue
	}
	return result
}

func (receiver *Env) GetInt32(key string, defaultValue int32) int32 {
	val := receiver.Get(key, defaultValue)
	result, err := anyto.Anyto(val).To().ValueE().Int32()
	if err != nil {
		return defaultValue
	}
	return result
}

func (receiver *Env) GetInt64(key string, defaultValue int64) int64 {
	val := receiver.Get(key, defaultValue)
	result, err := anyto.Anyto(val).To().ValueE().Int64()
	if err != nil {
		return defaultValue
	}
	return result
}

func (receiver *Env) GetFloat64(key string, defaultValue float64) float64 {
	val := receiver.Get(key, defaultValue)
	result, err := anyto.Anyto(val).To().ValueE().Float64()
	if err != nil {
		return defaultValue
	}
	return result
}

func (receiver *Env) GetStringSlice(key string, defaultValue []string) []string {
	val := receiver.Get(key, defaultValue)
	result, err := anyto.Anyto(val).To().ValueE().StringSlice()
	if err != nil {
		return defaultValue
	}
	return result
}

func (receiver *Env) GetString(key string, defaultValue string) string {
	val := receiver.Get(key, defaultValue)
	result, err := anyto.Anyto(val).To().ValueE().String()
	if err != nil {
		return defaultValue
	}
	return result
}

func (receiver *Env) GetStringMapString(key string, defaultValue map[string]string) map[string]string {
	val := receiver.Get(key, defaultValue)
	result, err := anyto.Anyto(val).To().ValueE().StringMapString()
	if err != nil {
		return defaultValue
	}
	return result
}

func (receiver *Env) GetStringMapStringSlice(key string, defaultValue map[string][]string) map[string][]string {
	val := receiver.Get(key, defaultValue)
	result, err := anyto.Anyto(val).To().ValueE().StringMapStringSlice()
	if err != nil {
		return defaultValue
	}
	return result
}

func (receiver *Env) GetTime(key string, defaultValue time.Time) time.Time {
	val := receiver.Get(key, defaultValue)
	result, err := anyto.Anyto(val).To().ValueE().Time()
	if err != nil {
		return defaultValue
	}
	return result
}
