package middlewares

import "github.com/gin-gonic/gin"

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("access-token")
		// fmt.Println(c.Request.URL.RawQuery)

		if token != "123456789" {
			c.JSON(200, gin.H{"code": 1002, "errmsg": "access token fail"})
			c.Abort()

			return
		}

		c.Next()
	}
}
