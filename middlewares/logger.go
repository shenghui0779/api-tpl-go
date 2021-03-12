package middlewares

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/shenghui0779/demo/helpers"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/shenghui0779/yiigo"
	"go.uber.org/zap"
)

var (
	bufPool = sync.Pool{
		New: func() interface{} {
			return bytes.NewBuffer(make([]byte, 0, 2<<10)) // 2KB
		},
	}

	replacer = strings.NewReplacer("\n", "", "\t", "", "\r", "#")
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := time.Now().Local()

		var body []byte

		// 取出请求Body
		if r.Body != nil && r.Body != http.NoBody {
			var err error

			body, err = ioutil.ReadAll(r.Body)

			if err != nil {
				yiigo.Logger().Error("Middleware.Logger error", zap.Error(err), zap.String("url", r.URL.String()))

				helpers.JSON(w, yiigo.X{
					"success": false,
					"code":    helpers.ErrSystem,
					"msg":     "服务器错误，请稍后重试",
				})

				return
			}

			// 关闭原Body
			r.Body.Close()

			r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		}

		// 存储返回结果
		buf := bufPool.Get().(*bytes.Buffer)
		buf.Reset()

		defer bufPool.Put(buf)

		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		ww.Tee(buf)

		next.ServeHTTP(ww, r)

		yiigo.Logger("request").Info(r.URL.String(),
			zap.String("request_id", middleware.GetReqID(r.Context())),
			zap.String("method", r.Method),
			zap.String("body", replacer.Replace(string(body))),
			zap.String("response", buf.String()),
			zap.String("duration", time.Since(now).String()),
			zap.Int("status", ww.Status()),
		)
	})
}
