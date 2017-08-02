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

	bootstrap()

	version := yiigo.GetEnvString("app", "version", "1.0.0")
	fmt.Println("app start, version", version)

	run()
}

// load routes
func loadRoutes(r *gin.Engine) {
	routes.LoadWelcomeRoutes(r)
	routes.LoadBookRoutes(r)
	routes.LoadStudentRoutes(r)
}

func bootstrap() {
	b := yiigo.New()

	// b.EnableMongo()
	b.EnableRedis()

	err := b.Bootstrap()

	if err != nil {
		yiigo.LogError(err.Error())
	}
}

func run() {
	debug := yiigo.GetEnvBool("app", "debug", false)
	mode := gin.ReleaseMode

	if debug {
		mode = gin.DebugMode
	}

	gin.SetMode(mode)

	r := gin.New()
	loadRoutes(r)
	r.Run(":8000")
}
