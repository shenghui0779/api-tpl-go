package main

import (
	"demo/middlewares"
	"demo/routes"
	"fmt"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/iiinsomnia/yiigo"
)

func main() {
	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)

	err := yiigo.Bootstrap(true, false, false)

	if err != nil {
		yiigo.Logger.Panic(err.Error())
	}

	fmt.Println("app start, version", yiigo.Env.String("app.version", "1.0.0"))

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
	r.Run(":8000")
}
