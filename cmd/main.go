package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/shenghui0779/yiigo"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"

	"tplgo/internal/middlewares"
	"tplgo/pkg/console"
	"tplgo/pkg/routes"
)

var envDir string

func main() {
	app := &cli.App{
		Name:     "tplgo",
		Usage:    "go project template",
		Commands: console.Commands,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "env-dir",
				Aliases:     []string{"E"},
				Value:       "",
				Usage:       "配置文件所在目录，默认当前目录",
				Destination: &envDir,
			},
		},
		Before: func(c *cli.Context) error {
			yiigo.Init(
				yiigo.WithEnvDir(envDir),
				yiigo.WithEnvWatcher(),
			)

			return nil
		},
		Action: func(c *cli.Context) error {
			serving()

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		yiigo.Logger().Fatal("app running error", zap.Error(err))
	}
}

func serving() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middlewares.Recovery)

	routes.Register(r)

	srv := &http.Server{
		Addr:         ":10086",
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
