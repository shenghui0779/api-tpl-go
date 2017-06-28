package routes

import (
	"demo/controllers/v1"

	"github.com/gin-gonic/gin"
)

func LoadBookRoutes(router *gin.Engine) {
	_v1 := router.Group("/v1")
	// _v1.Use(middlewares.Auth())
	{
		_v1.GET("books", v1.BookIndex)
		_v1.GET("books/:id", v1.BookView)
		_v1.POST("books", v1.BookAdd)
		_v1.PUT("books/:id", v1.BookEdit)
		_v1.DELETE("books/:id", v1.BookDelete)
	}
}
