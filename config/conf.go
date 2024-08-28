package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetConfigName("settings")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config/")

	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Sprintf("读取配置文件失败: %s\n", err))
	}

}
