package middlewares

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shenghui0779/yiigo"
	"go.uber.org/zap"
)

// Logger log middleware
func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		now := time.Now().Local()

		body, err := drainBody(ctx)

		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"success": false,
				"code":    50000,
				"msg":     "服务器错误，请稍后重试",
			})

			ctx.Abort()

			return
		}

		requestID := strconv.FormatInt(now.UnixNano(), 36)

		defer func() {
			var response interface{}

			if v, ok := ctx.Get("response"); ok {
				response = v
			}

			yiigo.Logger().Debug(fmt.Sprintf("[%s] %v", ctx.Request.Method, ctx.Request.URL),
				zap.String("request_id", requestID),
				zap.String("ip", ctx.ClientIP()),
				zap.String("params", body),
				zap.Any("response", response),
				zap.String("duration", time.Since(now).String()),
			)
		}()

		ctx.Request.Header.Set("request_id", requestID)
		ctx.Next()
	}
}

func drainBody(ctx *gin.Context) (string, error) {
	buf := yiigo.BufPool.Get()
	defer yiigo.BufPool.Put(buf)

	if ctx.Request.Body == nil || ctx.Request.Body == http.NoBody {
		return "", nil
	}

	if _, err := buf.ReadFrom(ctx.Request.Body); err != nil {
		yiigo.Logger().Error("drain request body error", zap.Error(err))

		return "", err
	}

	if err := ctx.Request.Body.Close(); err != nil {
		yiigo.Logger().Error("drain request body error", zap.Error(err))

		return "", err
	}

	bodyStr := buf.String()

	ctx.Request.Body = ioutil.NopCloser(bytes.NewReader([]byte(bodyStr)))

	return bodyStr, nil
}
