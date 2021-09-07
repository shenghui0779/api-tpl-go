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

	"tplgo/pkg/console"
	"tplgo/pkg/middlewares"
	"tplgo/pkg/routes"
)

var envFile string

func main() {
	app := &cli.App{
		Name:     "tplgo",
		Usage:    "go project template",
		Commands: console.Commands,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "envfile",
				Aliases:     []string{"E"},
				Value:       "yiigo.toml",
				Usage:       "设置配置文件，默认：yiigo.toml",
				Destination: &envFile,
			},
		},
		Before: func(c *cli.Context) error {
			yiigo.LoadEnvFromFile(envFile)
			yiigo.Init(
				yiigo.WithDB(yiigo.Default, yiigo.MySQL, yiigo.Env("db.dsn").String()),
				yiigo.WithLogger(yiigo.Default, "logs/app.log"),
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
