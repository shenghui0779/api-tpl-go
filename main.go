package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/shenghui0779/demo/cmd"
	"github.com/shenghui0779/demo/middlewares"
	"github.com/shenghui0779/demo/routes"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/shenghui0779/yiigo"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

func main() {
	app := cli.NewApp()
	app.Name = "yiigo.demo"
	app.Usage = "a demo for yiigo"
	app.Commands = cmd.Commands
	app.Action = func(c *cli.Context) error {
		serving()

		return nil
	}

	if err := app.Run(os.Args); err != nil {
		yiigo.Logger().Fatal("app running error", zap.Error(err))
	}
}

func serving() {
	r := chi.NewRouter()

	r.Use(middlewares.Recovery)
	r.Use(middleware.RequestID)

	routes.RegisterApp(r)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", yiigo.Env("app.port").Int(8000)),
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
