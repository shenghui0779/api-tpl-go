package routes

import (
	"net/http"

	"github.com/iiinsomnia/yiigo_demo/controllers"
	"github.com/iiinsomnia/yiigo_demo/middlewares"

	"github.com/gin-gonic/gin"
)

// RouteRegister register routes
func RouteRegister(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "☺ welcome to golang app!")
	})

	root := r.Group("/")
	root.Use(middlewares.Logger())
	{
		root.POST("book/info", controllers.GetBookInfo)
		root.POST("book/add", controllers.AddBook)
		root.POST("book/edit", controllers.EditBook)
		root.POST("book/delete", controllers.DeleteBook)
	}
}
