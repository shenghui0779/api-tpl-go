package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/iiinsomnia/yiigo"
)

func LoadWelcomeRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		yiigo.JSON(c, 0, "welcome to golang app!")
	})
}
