package config

import (
	"os"

	"github.com/shenghui0779/yiigo"
)

func Logger() *yiigo.LoggerConfig {
	return &yiigo.LoggerConfig{
		Filename: os.Getenv("LOG_PATH"),
		Options: &yiigo.LoggerOptions{
			Stderr: true,
		},
	}
}
