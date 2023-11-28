package cmd

import (
	"api/config"
	"api/db"
	"api/pkg/middleware"
	"api/pkg/router"
	"api/redis"

	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	chimidware "github.com/go-chi/chi/v5/middleware"
	goredis "github.com/redis/go-redis/v9"
	"github.com/shenghui0779/yiigo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var envFile string

var rootCmd = &cobra.Command{
	Use:         "api",
	Short:       "Go应用服务API",
	Long:        "Go应用服务API(build with yiigo & chi)",
	Annotations: map[string]string{},
	Version:     "v1.0.0",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		preInit()
	},
	Run: func(cmd *cobra.Command, args []string) {
		// make sure we have a working tempdir in minimal containers, because:
		// os.TempDir(): The directory is neither guaranteed to exist nor have accessible permissions.
		if err := os.MkdirAll(os.TempDir(), 0775); err != nil {
			yiigo.Logger().Error("err create temp dir", zap.Error(err))
		}

		serving()
	},
}

func preInit() {
	config.Init()

	db.Init(&db.Config{
		Driver: viper.GetString("db.driver"),
		DSN:    viper.GetString("db.dsn"),
		Options: &db.Options{
			MaxOpenConns:    viper.GetInt("db.max_open_conns"),
			MaxIdleConns:    viper.GetInt("db.max_idle_conns"),
			ConnMaxLifetime: viper.GetDuration("db.conn_max_lifetime") * time.Second,
			ConnMaxIdleTime: viper.GetDuration("db.conn_max_idle_time") * time.Second,
		},
	})

	redis.Init(&goredis.UniversalOptions{
		Addrs:           []string{viper.GetString("redis.addr")},
		DB:              viper.GetInt("redis.db"),
		Username:        viper.GetString("redis.username"),
		Password:        viper.GetString("redis.password"),
		DialTimeout:     viper.GetDuration("redis.conn_timeout") * time.Second,
		ReadTimeout:     viper.GetDuration("redis.read_timeout") * time.Second,
		WriteTimeout:    viper.GetDuration("redis.write_timeout") * time.Second,
		PoolSize:        viper.GetInt("redis.pool_size"),
		PoolTimeout:     viper.GetDuration("redis.pool_timeout") * time.Second,
		MinIdleConns:    viper.GetInt("redis.min_idle_conns"),
		MaxIdleConns:    viper.GetInt("redis.max_idle_conns"),
		MaxActiveConns:  viper.GetInt("max_active_conns"),
		ConnMaxIdleTime: viper.GetDuration("redis.conn_max_idle_time") * time.Second,
		ConnMaxLifetime: viper.GetDuration("redis.conn_max_lifetime") * time.Second,
	})
}

func serving() {
	r := chi.NewRouter()

	r.Use(chimidware.RequestID, middleware.Cors, middleware.Recovery)
	r.Mount("/debug", chimidware.Profiler())

	router.App(r)

	srv := &http.Server{
		Addr:         ":" + viper.GetString("app.port"),
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  10 * time.Second,
	}

	fmt.Println("listening on", srv.Addr)

	if err := srv.ListenAndServe(); err != nil {
		yiigo.Logger().Fatal("serving error", zap.Error(err))
	}
}
