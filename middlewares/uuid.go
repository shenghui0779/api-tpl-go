package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/iiinsomnia/yiigo"
)

func UUID() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := c.Request.Header.Get("Access-UUID")

		if strings.TrimSpace(uuid) == "" {
			yiigo.Error(c, "Invalid token, access failed!")
			c.Abort()

			return
		}

		c.Next()
	}
}
