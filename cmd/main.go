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

	"tplgo/pkg/config"
	"tplgo/pkg/console"
	"tplgo/pkg/ent"
	"tplgo/pkg/iao"
	"tplgo/pkg/middlewares"
	"tplgo/pkg/routes"
)

var envFile string

func main() {
	app := &cli.App{
		Name:     "tplgo",
		Usage:    "go web project template",
		Commands: console.Commands,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "envfile",
				Aliases:     []string{"E"},
				Value:       ".env",
				Usage:       "设置配置文件，默认：.env",
				Destination: &envFile,
			},
		},
		Before: func(c *cli.Context) error {
			yiigo.LoadEnv(yiigo.WithEnvFile(envFile))

			yiigo.Init(
				yiigo.WithMySQL(yiigo.Default, config.DB()),
				yiigo.WithLogger(yiigo.Default, config.Logger()),
			)

			return nil
		},
		Action: func(c *cli.Context) error {
			ent.InitDB()
			iao.InitClient()

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

	r.Use(middleware.RequestID, middlewares.Recovery)

	routes.Register(r)

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
