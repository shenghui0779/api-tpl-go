package main

import (
	"demo/middlewares"
	"demo/routes"
	"runtime"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
	"github.com/iiinsomnia/yiigo"
)

func main() {
	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)

	run()
}

func run() {
	debug := yiigo.Env.Bool("app.debug", false)

	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	if debug {
		r.Use(gin.Logger())
	}

	r.Use(middlewares.Recovery())

	routes.RouteRegister(r)

	if err := r.Run(":8000"); err != nil {
		yiigo.Logger.Panic("yiigo demo run error", zap.String("error", err.Error()))
	}
}
