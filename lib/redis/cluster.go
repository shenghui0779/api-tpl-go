package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

var cluster *redis.ClusterClient

// Init 初始化Redis集群
func InitCluster() error {
	cluster = redis.NewClusterClient(buildClusterOpts(viper.GetStringSlice("redis-cluster.addrs"), viper.GetStringMap("redis-cluster.options")))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// verify connection
	if err := cluster.Ping(ctx).Err(); err != nil {
		cluster.Close()
		return err
	}

	return nil
}

func Cluster() *redis.ClusterClient {
	return cluster
}

func buildClusterOpts(addrs []string, opts map[string]any) *redis.ClusterOptions {
	cfg := &redis.ClusterOptions{
		Addrs: addrs,
	}

	if len(opts) != 0 {
		cfg.Username = cast.ToString(opts["username"])
		cfg.Password = cast.ToString(opts["password"])

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

		cfg.MaxRedirects = cast.ToInt(opts["max_redirects"])
		cfg.ReadOnly = cast.ToBool(opts["read_only"])
		cfg.RouteByLatency = cast.ToBool(opts["route_by_latency"])
		cfg.RouteRandomly = cast.ToBool(opts["route_randomly"])

		cfg.DisableIndentity = cast.ToBool(opts["disable_indentity"])
	}

	return cfg
}
