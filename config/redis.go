package config

import (
	"os"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

func Redis() *redis.UniversalOptions {
	cfg := &redis.UniversalOptions{
		Addrs: []string{os.Getenv("REDIS_ADDR")},
	}

	if v := os.Getenv("REDIS_USERNAME"); len(v) != 0 {
		cfg.Username = v
	}

	if v := os.Getenv("REDIS_PASSWORD"); len(v) != 0 {
		cfg.Password = v
	}

	if v := os.Getenv("REDIS_DB"); len(v) != 0 {
		if i, err := strconv.Atoi(v); err == nil && i >= 0 {
			cfg.DB = i
		}
	}

	if v := os.Getenv("REDIS_CONN_TIMEOUT"); len(v) != 0 {
		if i, err := strconv.Atoi(v); err == nil && i > 0 {
			cfg.DialTimeout = time.Second * time.Duration(i)
		}
	}

	if v := os.Getenv("REDIS_READ_TIMEOUT"); len(v) != 0 {
		if i, err := strconv.Atoi(v); err == nil && i > 0 {
			cfg.ReadTimeout = time.Second * time.Duration(i)
		}
	}

	if v := os.Getenv("REDIS_WRITE_TIMEOUT"); len(v) != 0 {
		if i, err := strconv.Atoi(v); err == nil && i > 0 {
			cfg.WriteTimeout = time.Second * time.Duration(i)
		}
	}

	if v := os.Getenv("REDIS_POOL_SIZE"); len(v) != 0 {
		if i, err := strconv.Atoi(v); err == nil && i > 0 {
			cfg.PoolSize = i
		}
	}

	if v := os.Getenv("REDIS_POOL_TIMEOUT"); len(v) != 0 {
		if i, err := strconv.Atoi(v); err == nil && i > 0 {
			cfg.PoolTimeout = time.Second * time.Duration(i)
		}
	}

	if v := os.Getenv("REDIS_MIN_IDLE_CONNS"); len(v) != 0 {
		if i, err := strconv.Atoi(v); err == nil && i > 0 {
			cfg.MinIdleConns = i
		}
	}

	if v := os.Getenv("REDIS_MAX_IDLE_CONNS"); len(v) != 0 {
		if i, err := strconv.Atoi(v); err == nil && i > 0 {
			cfg.MaxIdleConns = i
		}
	}

	if v := os.Getenv("REDIS_MAX_ACTIVE_CONNS"); len(v) != 0 {
		if i, err := strconv.Atoi(v); err == nil && i > 0 {
			cfg.MaxActiveConns = i
		}
	}

	if v := os.Getenv("REDIS_CONN_MAX_LIFETIME"); len(v) != 0 {
		if i, err := strconv.Atoi(v); err == nil && i > 0 {
			cfg.ConnMaxLifetime = time.Second * time.Duration(i)
		}
	}

	if v := os.Getenv("REDIS_CONN_MAX_IDLETIME"); len(v) != 0 {
		if i, err := strconv.Atoi(v); err == nil && i > 0 {
			cfg.ConnMaxIdleTime = time.Second * time.Duration(i)
		}
	}

	return cfg
}
