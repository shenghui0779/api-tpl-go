package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

var cli redis.UniversalClient

// Init 初始化Redis(如有多个实例，在此方法中初始化)
func Init() error {
	cli = redis.NewUniversalClient(buildOpts(viper.GetStringSlice("redis.addrs"), viper.GetStringMap("redis.options")))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// verify connection
	if err := cli.Ping(ctx).Err(); err != nil {
		cli.Close()
		return err
	}

	return nil
}

func Client() redis.UniversalClient {
	return cli
}

func buildOpts(addrs []string, options map[string]any) *redis.UniversalOptions {
	cfg := &redis.UniversalOptions{
		Addrs: addrs,
	}

	if len(options) != 0 {
		// Database to be selected after connecting to the server.
		// Only single-node and failover clients.
		cfg.DB = cast.ToInt(options["db"])

		// Common options.

		cfg.Username = cast.ToString(options["username"])
		cfg.Password = cast.ToString(options["password"])
		cfg.SentinelUsername = cast.ToString(options["sentinel_username"])
		cfg.SentinelPassword = cast.ToString(options["sentinel_password"])

		cfg.MaxRetries = cast.ToInt(options["max_retries"])
		cfg.MinRetryBackoff = cast.ToDuration(options["min_retry_backoff"])
		cfg.MaxRetryBackoff = cast.ToDuration(options["max_retry_backoff"])

		cfg.DialTimeout = cast.ToDuration(options["dial_timeout"]) * time.Second
		cfg.ReadTimeout = cast.ToDuration(options["read_timeout"]) * time.Second
		cfg.WriteTimeout = cast.ToDuration(options["write_timeout"]) * time.Second
		cfg.ContextTimeoutEnabled = cast.ToBool(options["context_timeout_enabled"])

		// PoolFIFO uses FIFO mode for each node connection pool GET/PUT (default LIFO).
		cfg.PoolFIFO = cast.ToBool(options["pool_fifo"])

		cfg.PoolSize = cast.ToInt(options["pool_size"])
		cfg.PoolTimeout = cast.ToDuration(options["pool_timeout"]) * time.Second
		cfg.MinIdleConns = cast.ToInt(options["min_idle_conns"])
		cfg.MaxIdleConns = cast.ToInt(options["max_idle_conns"])
		cfg.MaxActiveConns = cast.ToInt(options["max_active_conns"])
		cfg.ConnMaxIdleTime = cast.ToDuration(options["conn_max_idle_time"]) * time.Second
		cfg.ConnMaxLifetime = cast.ToDuration(options["conn_max_lifetime"]) * time.Second

		// Only cluster clients.

		cfg.MaxRedirects = cast.ToInt(options["max_redirects"])
		cfg.ReadOnly = cast.ToBool(options["read_only"])
		cfg.RouteByLatency = cast.ToBool(options["route_by_latency"])
		cfg.RouteRandomly = cast.ToBool(options["route_randomly"])

		// The sentinel master name.
		// Only failover clients.
		cfg.MasterName = cast.ToString(options["master_name"])

		cfg.DisableIndentity = cast.ToBool(options["disable_indentity"])
	}

	return cfg
}
