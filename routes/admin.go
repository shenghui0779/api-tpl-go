package routes

import (
	"controllers/v1"
	"middlewares"

	"github.com/gin-gonic/gin"
)

func LoadAdminRoutes(router *gin.Engine) {
	api := router.Group("/api")
	api.Use(middlewares.Auth())
	{
		gv1 := api.Group("/v1")
		{
			gv1.GET("admin/list", v1.GetAdminList)
			gv1.GET("admin/detail", v1.GetAdminDetail)
		}
	}
}
