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

	yiigo.Bootstrap(true, true, true)

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
	mode := gin.ReleaseMode

	if yiigo.EnvBool("app", "debug", false) {
		mode = gin.DebugMode
	}

	gin.SetMode(mode)

	r := gin.New()
	loadRoutes(r)
	r.Run(":8000")
}
