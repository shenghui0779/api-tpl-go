package routes

import (
	"controllers/v1"

	"github.com/gin-gonic/gin"
)

func LoadArticleRoutes(router *gin.Engine) {
	_v1 := router.Group("/v1")
	// _v1.Use(middlewares.Auth())
	{
		_v1.GET("articles", v1.GetArticleList)
		_v1.GET("articles/:id", v1.GetArticleDetail)
		_v1.POST("articles", v1.AddNewArticle)
		_v1.PUT("articles/:id", v1.UpdateArticle)
		_v1.DELETE("articles/:id", v1.DeleteArticle)
	}
}
