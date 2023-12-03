package config

import (
	"context"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Environment struct {
	AppDebug  bool
	AppSecret string
}

var ENV = new(Environment)

func refresh() {
	ENV.AppDebug = viper.GetBool("app.debug")
	ENV.AppSecret = viper.GetString("app.secret")
}

func Init(path string) {
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		log.Panic(context.Background(), "err read config file", zap.Error(err))
	}

	refresh()
	viper.OnConfigChange(func(e fsnotify.Event) {
		refresh()
	})
	viper.WatchConfig()
}
