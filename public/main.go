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

// load routes
func loadRoutes(r *gin.Engine) {
	routes.LoadWelcomeRoutes(r)
	routes.LoadBookRoutes(r)
	routes.LoadStudentRoutes(r)
}

func run() {
	debug := yiigo.EnvBool("app", "debug", true)

	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	if debug {
		r.Use(gin.Logger(), gin.Recovery())
	}

	loadRoutes(r)
	r.Run(":8000")
}
