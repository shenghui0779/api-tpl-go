package config

import (
	"os"
	"strconv"
	"time"

	"github.com/shenghui0779/yiigo"
)

func Redis() *yiigo.RedisConfig {
	cfg := &yiigo.RedisConfig{
		Addr: os.Getenv("REDIS_ADDR"),
		Options: &yiigo.RedisOptions{
			ConnTimeout:  10 * time.Second,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
			PoolSize:     20,
			IdleTimeout:  time.Minute,
		},
	}

	if v := os.Getenv("REDIS_USERNAME"); len(v) != 0 {
		cfg.Options.Username = v
	}

	if v := os.Getenv("REDIS_PASSWORD"); len(v) != 0 {
		cfg.Options.Password = v
	}

	if v := os.Getenv("REDIS_DB"); len(v) != 0 {
		if i, err := strconv.Atoi(v); err == nil && i >= 0 {
			cfg.Options.Database = i
		}
	}

	if v := os.Getenv("REDIS_CONN_TIMEOUT"); len(v) != 0 {
		if i, err := strconv.Atoi(v); err == nil && i > 0 {
			cfg.Options.ConnTimeout = time.Second * time.Duration(i)
		}
	}

	if v := os.Getenv("REDIS_READ_TIMEOUT"); len(v) != 0 {
		if i, err := strconv.Atoi(v); err == nil && i > 0 {
			cfg.Options.ReadTimeout = time.Second * time.Duration(i)
		}
	}

	if v := os.Getenv("REDIS_WRITE_TIMEOUT"); len(v) != 0 {
		if i, err := strconv.Atoi(v); err == nil && i > 0 {
			cfg.Options.WriteTimeout = time.Second * time.Duration(i)
		}
	}

	if v := os.Getenv("REDIS_POOL_SIZE"); len(v) != 0 {
		if i, err := strconv.Atoi(v); err == nil && i > 0 {
			cfg.Options.PoolSize = i
		}
	}

	if v := os.Getenv("REDIS_POOL_PREFILL"); len(v) != 0 {
		if i, err := strconv.Atoi(v); err == nil && i >= 0 {
			cfg.Options.PoolPrefill = i
		}
	}

	if v := os.Getenv("REDIS_IDLE_TIMEOUT"); len(v) != 0 {
		if i, err := strconv.Atoi(v); err != nil && i > 0 {
			cfg.Options.IdleTimeout = time.Second * time.Duration(i)
		}
	}

	return cfg
}
