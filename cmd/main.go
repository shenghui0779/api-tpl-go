package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/shenghui0779/yiigo"
	"go.uber.org/zap"

	"tplgo/pkg/config"
	"tplgo/pkg/ent"
	"tplgo/pkg/middlewares"
	"tplgo/pkg/router"
)

var envFile string

func main() {
	flag.StringVar(&envFile, "envfile", ".env", "设置ENV配置文件")

	flag.Parse()

	yiigo.LoadEnv(yiigo.WithEnvFile(envFile), yiigo.WithEnvWatcher(func(e fsnotify.Event) {
		yiigo.Logger().Info("env change ok", zap.String("event", e.String()))
		config.RefreshENV()
	}))

	yiigo.Init(
		yiigo.WithMySQL(yiigo.Default, config.DB()),
		yiigo.WithLogger(yiigo.Default, config.Logger()),
	)

	config.RefreshENV()
	ent.InitDB()

	// make sure we have a working tempdir in minimal containers, because:
	// os.TempDir(): The directory is neither guaranteed to exist nor have accessible permissions.
	if err := os.MkdirAll(os.TempDir(), 0775); err != nil {
		yiigo.Logger().Error("err create temp dir", zap.Error(err))
	}

	serving()
}

func serving() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID, middlewares.Cors, middlewares.Recovery)

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
