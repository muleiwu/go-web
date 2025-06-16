package config

import (
	"fmt"
	"github.com/spf13/viper"
	"mliev.com/template/go-web/support/logger"
	"strings"
)

func InitViper() {
	viper.SetConfigName("config") // 配置文件名称(无扩展名)
	viper.SetConfigType("yaml")   // 或viper.SetConfigType("YAML")
	viper.AddConfigPath(".")      // 配置文件路径
	//支持读取环境变量
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	//读取配置文件
	err := viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {             // 处理读取配置文件的错误
		logger.Get().Error(fmt.Sprintf("Fatal error config file: %v \n", err))
	}
}
