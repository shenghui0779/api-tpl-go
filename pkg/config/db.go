package config

import (
	"os"
	"strconv"
	"time"

	"github.com/shenghui0779/yiigo"
)

func DB() *yiigo.DBConfig {
	cfg := &yiigo.DBConfig{
		DSN: os.Getenv("DB_DSN"),
		Options: &yiigo.DBOptions{
			MaxOpenConns:    20,
			MaxIdleConns:    10,
			ConnMaxLifetime: 10 * time.Minute,
			ConnMaxIdleTime: 5 * time.Minute,
		},
	}

	if v := os.Getenv("DB_MAX_OPEN_CONNS"); len(v) != 0 {
		if i, err := strconv.Atoi(v); err == nil && i > 0 {
			cfg.Options.MaxOpenConns = i
		}
	}

	if v := os.Getenv("DB_MAX_IDLE_CONNS"); len(v) != 0 {
		if i, err := strconv.Atoi(v); err == nil && i > 0 {
			cfg.Options.MaxIdleConns = i
		}
	}

	if v := os.Getenv("DB_CONN_MAX_LIFE_TIME"); len(v) != 0 {
		if i, err := strconv.Atoi(v); err == nil && i > 0 {
			cfg.Options.ConnMaxLifetime = time.Duration(i) * time.Second
		}
	}

	if v := os.Getenv("DB_CONN_MAX_IDLE_TIME"); len(v) != 0 {
		if i, err := strconv.Atoi(v); err == nil && i > 0 {
			cfg.Options.ConnMaxIdleTime = time.Duration(i) * time.Second
		}
	}

	return cfg
}
