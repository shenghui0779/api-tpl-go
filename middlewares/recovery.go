package middlewares

import (
	"fmt"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/iiinsomnia/yiigo"
)

// Recovery panic recover middleware
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				yiigo.Logger.Error(fmt.Sprintf("%v\n%s", err, string(debug.Stack())))
				yiigo.Error(c, 5000, "Internal server error")
				c.Abort()

				return
			}
		}()

		c.Next()
	}
}
