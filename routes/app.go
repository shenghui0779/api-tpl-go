package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/iiinsomnia/demo/controllers"
	"github.com/iiinsomnia/demo/middlewares"
)

// RegisterApp register app routes
func RegisterApp(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "☺ welcome to golang app")
	})

	// 探侦地址，用于健康检查
	r.HEAD("/listen", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	root := r.Group("/")
	root.Use(middlewares.Logger())
	{
		root.POST("/book/info", controllers.GetBookInfo)
	}
}
