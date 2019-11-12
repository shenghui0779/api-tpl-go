package middlewares

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/iiinsomnia/yiigo/v4"
	"github.com/iiinsomnia/yiigo_demo/helpers"
	"go.uber.org/zap"
)

var bufPool = sync.Pool{
	New: func() interface{} {
		return bytes.NewBuffer(make([]byte, 0, 16<<10)) // 16KB
	},
}

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

		requestID := yiigo.MD5(strconv.FormatInt(helpers.UniqueSerial(), 10))

		c.Request.Header.Set("Request-ID", requestID)

		defer func() {
			endTime := time.Now().UnixNano()

			var response interface{}

			if v, ok := c.Get("response"); ok {
				response = v
			}

			yiigo.Logger().Debug(fmt.Sprintf("[%s] %v", c.Request.Method, c.Request.URL),
				zap.String("request_id", requestID),
				zap.String("ip", c.ClientIP()),
				zap.ByteString("params", body),
				zap.Any("response", response),
				zap.String("duration", fmt.Sprintf("%f ms", float64(endTime-startTime)/1e6)),
			)

		}()

		c.Next()
	}
}

func drainBody(c *gin.Context) ([]byte, error) {
	buf := bufPool.Get().(*bytes.Buffer)
	buf.Reset()

	defer bufPool.Put(buf)

	if c.Request.Body == nil || c.Request.Body == http.NoBody {
		return nil, nil
	}

	if _, err := buf.ReadFrom(c.Request.Body); err != nil {
		yiigo.Logger().Error("drain request body error", zap.Error(err))

		return nil, err
	}

	if err := c.Request.Body.Close(); err != nil {
		yiigo.Logger().Error("drain request body error", zap.Error(err))

		return nil, err
	}

	body := buf.Bytes()

	c.Request.Body = ioutil.NopCloser(bytes.NewReader(body))

	return body, nil
}
