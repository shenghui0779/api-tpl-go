package routes

import (
	"demo/controllers/v1"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RouteRegister register routes
func RouteRegister(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "☺︎ welcome to golang app!")
	})

	_v1 := r.Group("/v1")
	// _v1.Use(middlewares.Auth())
	{
		// book
		_v1.GET("books", v1.BookIndex)
		_v1.GET("books/:id", v1.BookView)
		_v1.POST("books", v1.BookAdd)
		_v1.PUT("books/:id", v1.BookEdit)
		_v1.DELETE("books/:id", v1.BookDelete)

		// student
		_v1.GET("students", v1.StudentIndex)
		_v1.GET("students/:id", v1.StudentView)
		_v1.POST("students", v1.StudentAdd)
		_v1.PUT("students/:id", v1.StudentEdit)
		_v1.DELETE("students/:id", v1.StudentDelete)
	}
}
