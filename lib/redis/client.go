package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

var Client *redis.Client

// Init 初始化Redis(如有多个实例，在此方法中初始化)
func Init() error {
	cfg := &redis.Options{
		Addr: viper.GetString("redis.addr"),
	}

	opts := viper.GetStringMap("redis.options")
	if len(opts) != 0 {
		cfg.DB = cast.ToInt(opts["db"])
		cfg.Username = cast.ToString(opts["username"])
		cfg.Password = cast.ToString(opts["password"])
		cfg.MaxRetries = cast.ToInt(opts["max_retries"])
		cfg.MinRetryBackoff = cast.ToDuration(opts["min_retry_backoff"])
		cfg.MaxRetryBackoff = cast.ToDuration(opts["max_retry_backoff"])
		cfg.DialTimeout = cast.ToDuration(opts["dial_timeout"]) * time.Second
		cfg.ReadTimeout = cast.ToDuration(opts["read_timeout"]) * time.Second
		cfg.WriteTimeout = cast.ToDuration(opts["write_timeout"]) * time.Second
		cfg.ContextTimeoutEnabled = cast.ToBool(opts["context_timeout_enabled"])
		cfg.PoolFIFO = cast.ToBool(opts["pool_fifo"]) // PoolFIFO uses FIFO mode for each node connection pool GET/PUT (default LIFO).
		cfg.PoolSize = cast.ToInt(opts["pool_size"])
		cfg.PoolTimeout = cast.ToDuration(opts["pool_timeout"]) * time.Second
		cfg.MinIdleConns = cast.ToInt(opts["min_idle_conns"])
		cfg.MaxIdleConns = cast.ToInt(opts["max_idle_conns"])
		cfg.MaxActiveConns = cast.ToInt(opts["max_active_conns"])
		cfg.ConnMaxIdleTime = cast.ToDuration(opts["conn_max_idle_time"]) * time.Second
		cfg.ConnMaxLifetime = cast.ToDuration(opts["conn_max_lifetime"]) * time.Second
		cfg.DisableIndentity = cast.ToBool(opts["disable_indentity"])
	}

	Client = redis.NewClient(cfg)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// verify connection
	if err := Client.Ping(ctx).Err(); err != nil {
		Client.Close()
		return err
	}
	return nil
}
