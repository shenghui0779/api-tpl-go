package cmd

import (
	"api/lib/config"
	"api/lib/db"
	"api/lib/log"
	"api/lib/redis"
	"api/pkg/middleware"
	"api/pkg/router"

	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	chi_middleware "github.com/go-chi/chi/v5/middleware"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:         "api",
	Short:       "Go应用服务API",
	Long:        "Go应用服务API(build with ent & chi)",
	Annotations: map[string]string{},
	Version:     "v1.0.0",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		preInit(context.Background())
	},
	Run: func(cmd *cobra.Command, args []string) {
		// make sure we have a working tempdir in minimal containers, because:
		// os.TempDir(): The directory is neither guaranteed to exist nor have accessible permissions.
		if err := os.MkdirAll(os.TempDir(), 0775); err != nil {
			log.Error(context.Background(), "err create temp dir", zap.Error(err))
		}

		serving()
	},
}

func preInit(ctx context.Context) {
	// 初始化配置
	config.Init(cfgFile)
	// 初始化日志
	log.Init()
	// 初始化数据库
	if err := db.Init(); err != nil {
		log.Panic(ctx, "数据库初始化失败", zap.Error(err))
	}
	// 初始化Redis
	if err := redis.Init(); err != nil {
		log.Panic(ctx, "Redis初始化失败", zap.Error(err))
	}
}

func serving() {
	r := chi.NewRouter()

	r.Use(chi_middleware.RequestID, middleware.Cors, middleware.Recovery)
	r.Mount("/debug", chi_middleware.Profiler())

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
		log.Fatal(context.Background(), "serving error", zap.Error(err))
	}
}
