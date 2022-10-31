package config

import (
	"os"
	"strconv"

	"github.com/shenghui0779/yiigo"
)

func Logger() *yiigo.LoggerConfig {
	cfg := &yiigo.LoggerConfig{
		Filename: os.Getenv("LOG_PATH"),
		Options:  new(yiigo.LoggerOptions),
	}

	if ok, _ := strconv.ParseBool(os.Getenv("DEBUG")); ok {
		cfg.Options.Stderr = true
	}

	return cfg
}
