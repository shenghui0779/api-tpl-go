package config

import (
	"os"

	"github.com/shenghui0779/yiigo"
)

func Logger() *yiigo.LoggerConfig {
	cfg := &yiigo.LoggerConfig{
		Filename: os.Getenv("LOG_PATH"),
		Options:  new(yiigo.LoggerOptions),
	}

	// 开发环境允许终端输出日志
	if os.Getenv("ENV") == "dev" {
		cfg.Options.Stderr = true
	}

	return cfg
}
