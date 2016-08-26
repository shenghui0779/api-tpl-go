package routes

import (
	"github.com/gin-gonic/gin"
)

func InitHomeRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code":    1001,
			"success": true,
			"msg":     "welcome to yiigo service",
		})
	})
}
