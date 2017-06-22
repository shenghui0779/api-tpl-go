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

	runYiigo()

	version := yiigo.GetEnvString("app", "version", "1.0.0")
	fmt.Println("app start, version", version)

	runServer()
}

// load routes
func loadRoutes(r *gin.Engine) {
	routes.LoadWelcomeRoutes(r)
	routes.LoadBookRoutes(r)
	routes.LoadStudentRoutes(r)
}

func runYiigo() {
	b := yiigo.New()
	b.EnableMongo()
	b.EnableRedis()
	b.Run()
}

func runServer() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	loadRoutes(r)
	r.Run(":8000")
}
