package routes

import (
	"controllers/v1"
	"middlewares"

	"github.com/gin-gonic/gin"
)

func LoadUserRoutes(router *gin.Engine) {
	api := router.Group("/api")
	api.Use(middlewares.Auth())
	{
		gv1 := api.Group("/v1")
		{
			gv1.GET("user/list", v1.GetUserList)
			gv1.GET("user/detail", v1.GetUserDetail)
		}
	}
}
