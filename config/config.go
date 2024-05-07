package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var ConfigInfo config

func Init() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config.yaml")
	viper.AddConfigPath("../../../../config") //这个路径为调用Init的路径为始点进行查找起点
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logrus.Info("config file not found: ", err)
		} else {
			logrus.Info("config error :", err)
		}
	}
	if err := viper.Unmarshal(&ConfigInfo); err != nil {
		logrus.Info("config decode error: ", err)
	}
}
