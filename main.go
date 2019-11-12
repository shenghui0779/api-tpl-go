package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/iiinsomnia/yiigo/v4"
	"github.com/iiinsomnia/yiigo_demo/middlewares"
	"github.com/iiinsomnia/yiigo_demo/routes"
	"go.uber.org/zap"
)

func main() {
	run()
}

func run() {
	// 弃用Gin内置验证器
	binding.Validator = nil

	debug := yiigo.Env("app.debug").Bool(false)

	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	r.Use(middlewares.Recovery())

	routes.RouteRegister(r)

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
		yiigo.Logger().Fatal("yiigo-demo server start error", zap.Error(err))
	}
}
