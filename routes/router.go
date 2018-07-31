package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"demo/controllers"
)

// RouteRegister register routes
func RouteRegister(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "☺ welcome to golang app!")
	})

	// book
	r.GET("books", controllers.BookIndex)
	r.GET("books/:id", controllers.BookView)
	r.POST("books", controllers.BookAdd)
	r.PUT("books/:id", controllers.BookEdit)
	r.DELETE("books/:id", controllers.BookDelete)

	// student
	r.GET("students", controllers.StudentIndex)
	r.GET("students/:id", controllers.StudentView)
	r.POST("students", controllers.StudentAdd)
	r.PUT("students/:id", controllers.StudentEdit)
	r.DELETE("students/:id", controllers.StudentDelete)
}
