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

func buildOpts(addrs []string, opts map[string]any) *redis.UniversalOptions {
	cfg := &redis.UniversalOptions{
		Addrs: addrs,
	}

	if len(opts) != 0 {
		// Database to be selected after connecting to the server.
		// Only single-node and failover clients.
		cfg.DB = cast.ToInt(opts["db"])

		// Common options.

		cfg.Username = cast.ToString(opts["username"])
		cfg.Password = cast.ToString(opts["password"])
		cfg.SentinelUsername = cast.ToString(opts["sentinel_username"])
		cfg.SentinelPassword = cast.ToString(opts["sentinel_password"])

		cfg.MaxRetries = cast.ToInt(opts["max_retries"])
		cfg.MinRetryBackoff = cast.ToDuration(opts["min_retry_backoff"])
		cfg.MaxRetryBackoff = cast.ToDuration(opts["max_retry_backoff"])

		cfg.DialTimeout = cast.ToDuration(opts["dial_timeout"]) * time.Second
		cfg.ReadTimeout = cast.ToDuration(opts["read_timeout"]) * time.Second
		cfg.WriteTimeout = cast.ToDuration(opts["write_timeout"]) * time.Second
		cfg.ContextTimeoutEnabled = cast.ToBool(opts["context_timeout_enabled"])

		// PoolFIFO uses FIFO mode for each node connection pool GET/PUT (default LIFO).
		cfg.PoolFIFO = cast.ToBool(opts["pool_fifo"])

		cfg.PoolSize = cast.ToInt(opts["pool_size"])
		cfg.PoolTimeout = cast.ToDuration(opts["pool_timeout"]) * time.Second
		cfg.MinIdleConns = cast.ToInt(opts["min_idle_conns"])
		cfg.MaxIdleConns = cast.ToInt(opts["max_idle_conns"])
		cfg.MaxActiveConns = cast.ToInt(opts["max_active_conns"])
		cfg.ConnMaxIdleTime = cast.ToDuration(opts["conn_max_idle_time"]) * time.Second
		cfg.ConnMaxLifetime = cast.ToDuration(opts["conn_max_lifetime"]) * time.Second

		// Only cluster clients.

		cfg.MaxRedirects = cast.ToInt(opts["max_redirects"])
		cfg.ReadOnly = cast.ToBool(opts["read_only"])
		cfg.RouteByLatency = cast.ToBool(opts["route_by_latency"])
		cfg.RouteRandomly = cast.ToBool(opts["route_randomly"])

		// The sentinel master name.
		// Only failover clients.
		cfg.MasterName = cast.ToString(opts["master_name"])

		cfg.DisableIndentity = cast.ToBool(opts["disable_indentity"])
	}

	return cfg
}
