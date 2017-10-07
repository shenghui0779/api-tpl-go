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

	err := yiigo.Bootstrap(true, true, true)

	if err != nil {
		yiigo.Error(err.Error())
		yiigo.Flush()
	}

	fmt.Println("app start, version", yiigo.EnvString("app", "version", "1.0.0"))

	run()
}

func run() {
	debug := yiigo.EnvBool("app", "debug", true)

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
