package main

import (
	"fmt"
	"routes"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/iiinsomnia/yiigo"
)

func main() {
	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)

	yiigo.LoadEnvConfig()
	yiigo.InitLogger()
	yiigo.InitDB()
	// yiigo.InitRedis()
	// yiigo.InitMongo()

	version := yiigo.GetEnvString("app", "version", "1.0")
	fmt.Println("server started, api version", version)

	run()
}

// load routes
func loadRoutes(router *gin.Engine) {
	routes.LoadAdminRoutes(router)
}

func run() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	loadRoutes(router)
	router.Run(":8000")
}
