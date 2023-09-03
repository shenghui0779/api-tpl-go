package cmd

import (
	"api/config"
	"api/ent"
	"api/pkg/middlewares"
	"api/pkg/router"

	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/shenghui0779/yiigo"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var envFile string

var root = &cobra.Command{
	Use:         "api",
	Short:       "应用服务API",
	Long:        "应用服务API(build with yiigo & chi)",
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
	yiigo.LoadEnv(yiigo.WithEnvFile(envFile), yiigo.WithEnvWatcher(func(e fsnotify.Event) {
		yiigo.Logger().Info("env change ok", zap.String("event", e.String()))
		config.Refresh()
	}))

	yiigo.Init(
		yiigo.WithMySQL(yiigo.Default, config.DB()),
		yiigo.WithLogger(yiigo.Default, config.Logger()),
	)

	config.Refresh()
	ent.InitDB()
}

func serving() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID, middlewares.Cors, middlewares.Recovery)
	r.Mount("/debug", middleware.Profiler())

	router.App(r)

	srv := &http.Server{
		Addr:         ":8000",
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
