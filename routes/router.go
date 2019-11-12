package routes

import (
	"net/http"

	"github.com/iiinsomnia/yiigo_demo/controllers"

	"github.com/gin-gonic/gin"
)

// RouteRegister register routes
func RouteRegister(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "☺ welcome to golang app!")
	})

	r.GET("books", controllers.BookIndex)
	r.GET("books/:id", controllers.BookView)
	r.POST("books", controllers.BookAdd)
	r.PUT("books/:id", controllers.BookEdit)
	r.DELETE("books/:id", controllers.BookDelete)
}
