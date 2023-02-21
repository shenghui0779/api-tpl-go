package middlewares

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/shenghui0779/yiigo"
	"github.com/tidwall/pretty"
	"go.uber.org/zap"

	"tplgo/pkg/consts"
	"tplgo/pkg/lib"
	"tplgo/pkg/logger"
	"tplgo/pkg/result"
)

// Log 日志中间件
func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		body := "<nil>"

		// 请求包含body
		if r.Body != nil && r.Body != http.NoBody {
			switch consts.ContentType(yiigo.ContentType(r)) {
			case consts.MIMEForm:
				if err := r.ParseForm(); err != nil {
					result.ErrSystem(result.Err(err)).JSON(w, r)

					return
				}

				body = r.Form.Encode()
			case consts.MIMEMultipartForm:
				if err := r.ParseMultipartForm(consts.MaxFormMemory); err != nil {
					if err != http.ErrNotMultipart {
						result.ErrSystem(result.Err(err)).JSON(w, r)

						return
					}
				}

				body = r.Form.Encode()
			case consts.ContentJSON:
				// 取出Body
				b, err := ioutil.ReadAll(r.Body)

				if err != nil {
					result.ErrSystem(result.Err(err)).JSON(w, r)

					return
				}

				// 关闭原Body
				r.Body.Close()

				body = string(pretty.Ugly(b))

				r.Body = ioutil.NopCloser(bytes.NewReader(b))
			}
		}

		next.ServeHTTP(w, r)

		logger.Info(r.Context(), "request info",
			zap.String("method", r.Method),
			zap.String("uri", r.URL.String()),
			zap.String("ip", r.RemoteAddr),
			zap.String("body", body),
			zap.String("identity", lib.GetIdentity(r.Context()).String()),
			zap.String("duration", time.Since(now).String()),
		)
	})
}
