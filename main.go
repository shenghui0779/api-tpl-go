package main

import (
	"demo/middlewares"
	"demo/routes"
	"fmt"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/iiinsomnia/yiigo"
	"go.uber.org/zap"
)

func main() {
	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)

	// 注册日志
	yiigo.RegisterLogger(yiigo.AsDefault, "app.log")

	// 使用配置文件
	if err := yiigo.UseEnv("env.toml"); err != nil {
		yiigo.Logger.Panic("use env error", zap.String("error", err.Error()))
	}

	// 注册DB
	if err := yiigo.RegisterDB(yiigo.AsDefault, yiigo.MySQL, fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?timeout=10s&charset=%s&collation=%s&parseTime=True&loc=Local",
		yiigo.Env.String("mysql.username"),
		yiigo.Env.String("mysql.password"),
		yiigo.Env.String("mysql.host"),
		yiigo.Env.Int("mysql.port"),
		yiigo.Env.String("mysql.database"),
		yiigo.Env.String("mysql.charset"),
		yiigo.Env.String("mysql.collection"),
	)); err != nil {
		yiigo.Logger.Panic("register db error", zap.String("error", err.Error()))
	}

	// 注册Redis
	yiigo.RegisterRedis(yiigo.AsDefault, "127.0.0.1:6379")

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
