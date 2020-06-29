package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/shenghui0779/yiigo"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"

	"github.com/iiinsomnia/demo/cmd"
	"github.com/iiinsomnia/demo/middlewares"
	"github.com/iiinsomnia/demo/routes"
)

func main() {
	run()
}

func run() {
	app := cli.NewApp()
	app.Name = "demo"
	app.Usage = "yiigo demo."
	app.Commands = cmd.Commands
	app.Action = func(c *cli.Context) error {
		serving()

		return nil
	}

	if err := app.Run(os.Args); err != nil {
		yiigo.Logger().Fatal("yiigo-demo run error", zap.Error(err))
	}
}

func serving() {
	// 弃用Gin内置验证器
	binding.Validator = yiigo.NewGinValidator()

	debug := yiigo.Env("app.debug").Bool(false)

	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	r.Use(middlewares.Recovery())

	routes.RegisterApp(r)

	// go tool pprof ...
	pprof.Register(r)

	// Graceful restart & zero downtime deploy for Go servers.
	// Use `kill -USR2 pid` to restart.
	if err := gracehttp.Serve(
		&http.Server{
			Addr:         fmt.Sprintf(":%d", yiigo.Env("app.port").Int(8000)),
			Handler:      r,
			IdleTimeout:  10 * time.Second,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
		}); err != nil {
		yiigo.Logger().Fatal("yiigo-demo serving error", zap.Error(err))
	}
}
