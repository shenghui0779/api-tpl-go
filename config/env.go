package config

import (
	"api/logger"
	"context"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Environment struct {
	Debug     bool
	APISecret string
}

var ENV = new(Environment)

func refresh() {
	ENV.Debug = viper.GetBool("app.debug")
	ENV.APISecret = viper.GetString("app.secret")
}

func Init() {
	viper.SetConfigFile(".yml")
	if err := viper.ReadInConfig(); err != nil {
		logger.Panic(context.Background(), "err read config file", zap.Error(err))
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		refresh()
	})
	viper.WatchConfig()
}
