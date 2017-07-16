package middlewares

import (
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/iiinsomnia/yiigo"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := c.Request.Header.Get("Access-UUID")
		accessTime := c.Request.Header.Get("Access-Time")
		accessSign := c.Request.Header.Get("Access-Sign")

		if strings.TrimSpace(uuid) == "" || strings.TrimSpace(accessTime) == "" || strings.TrimSpace(accessSign) == "" {
			yiigo.ReturnJSON(c, -1, "Invalid token, access failed!")
			c.Abort()

			return
		}

		// 验证登录
		code, msg := validateLogin()

		if code != 0 {
			yiigo.ReturnJSON(c, code, msg)
			c.Abort()

			return
		}

		// 验签
		code, msg = validateSign(c, accessTime, accessSign)

		if code != 0 {
			yiigo.ReturnJSON(c, code, msg)
			c.Abort()

			return
		}

		c.Next()
	}
}

// 验证登录
func validateLogin() (int, string) {
	return 0, "success"
}

// 验签
func validateSign(c *gin.Context, accessTime string, accessSign string) (int, string) {
	accessExpire := yiigo.GetEnvInt64("app", "accessExpire", 0)
	now := time.Now().Unix()
	timestamp, _ := strconv.ParseInt(accessTime, 10, 64)

	if accessExpire > 0 && (now-timestamp >= accessExpire) {
		return -1, "request expired!"
	}

	// uri := c.Request.RequestURI
	// path, _ := url.QueryUnescape(uri)

	return 0, "success"
}
