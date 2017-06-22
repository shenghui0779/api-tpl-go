package routes

import (
	"demo/controllers/v1"

	"github.com/gin-gonic/gin"
)

func LoadBookRoutes(router *gin.Engine) {
	_v1 := router.Group("/v1")
	// _v1.Use(middlewares.Auth())
	{
		_v1.GET("books", v1.GetBookList)
		_v1.GET("books/:id", v1.GetBookDetail)
		_v1.POST("books", v1.AddNewBook)
		_v1.PUT("books/:id", v1.UpdateBook)
		_v1.DELETE("books/:id", v1.DeleteBook)
	}
}
