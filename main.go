package main

import (
	"demo/routes"
	"fmt"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/iiinsomnia/yiigo"
)

func main() {
	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)

	err := yiigo.Bootstrap(true, false, true)

	if err != nil {
		yiigo.Logger.Panic(err.Error())
	}

	fmt.Println("app start, version", yiigo.EnvString("app", "version", "1.0.0"))

	run()
}

func run() {
	debug := yiigo.EnvBool("app", "debug", false)

	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	if debug {
		r.Use(gin.Logger())
	}

	r.Use(gin.Recovery())

	routes.RouteRegister(r)
	r.Run(":8000")
}
