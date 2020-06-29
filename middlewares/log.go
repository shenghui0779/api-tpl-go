package middlewares

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shenghui0779/yiigo"
	"go.uber.org/zap"
)

// Logger log middleware
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now().UnixNano()

		body, err := drainBody(c)

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"code":    50000,
				"msg":     "服务器错误，请稍后重试",
			})

			c.Abort()

			return
		}

		defer func() {
			endTime := time.Now().UnixNano()

			var response interface{}

			if v, ok := c.Get("response"); ok {
				response = v
			}

			yiigo.Logger().Debug(fmt.Sprintf("[%s] %v", c.Request.Method, c.Request.URL),
				zap.String("ip", c.ClientIP()),
				zap.String("params", body),
				zap.Any("response", response),
				zap.String("duration", fmt.Sprintf("%f ms", float64(endTime-startTime)/1e6)),
			)

		}()

		c.Next()
	}
}

func drainBody(c *gin.Context) (string, error) {
	buf := yiigo.BufPool.Get()
	defer yiigo.BufPool.Put(buf)

	if c.Request.Body == nil || c.Request.Body == http.NoBody {
		return "", nil
	}

	if _, err := buf.ReadFrom(c.Request.Body); err != nil {
		yiigo.Logger().Error("drain request body error", zap.Error(err))

		return "", err
	}

	if err := c.Request.Body.Close(); err != nil {
		yiigo.Logger().Error("drain request body error", zap.Error(err))

		return "", err
	}

	bodyStr := buf.String()

	c.Request.Body = ioutil.NopCloser(bytes.NewReader([]byte(bodyStr)))

	return bodyStr, nil
}
