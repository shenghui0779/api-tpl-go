package main

import (
	"fmt"
	"routes"
	"runtime"

	"github.com/gin-gonic/gin"
)

func main() {
	configRuntime()
	startListening()
}

func configRuntime() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	fmt.Printf("Running with %d CPUs\n", nuCPU)
}

// init routes
func initRoutes(router *gin.Engine) {
	routes.InitHomeRoutes(router)
	routes.InitTestRoutes(router)
}

func startListening() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	initRoutes(router)
	fmt.Println("Server Start Successful")
	router.Run(":8000")
}
