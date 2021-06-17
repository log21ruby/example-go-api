package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitViper(path string) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatalf("cannot read in viper config:%s", err)
	}
	viper.AutomaticEnv()
}
