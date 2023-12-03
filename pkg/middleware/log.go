package middleware

import (
	"bytes"
	"io"
	"net/http"
	"time"

	"github.com/pkg/errors"
	"github.com/tidwall/pretty"
	"go.uber.org/zap"

	libhttp "api/lib/http"
	"api/lib/log"
	"api/lib/util"
	"api/pkg/auth"
	"api/pkg/result"
)

const ContentJSON = "application/json"

// Log 日志中间件
func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		body := "<nil>"

		// 请求包含body
		if r.Body != nil && r.Body != http.NoBody {
			switch util.ContentType(r) {
			case libhttp.ContentForm:
				if err := r.ParseForm(); err != nil {
					result.ErrSystem(result.E(errors.WithMessage(err, "表单解析失败"))).JSON(w, r)
					return
				}

				body = r.Form.Encode()
			case libhttp.MultipartForm:
				if err := r.ParseMultipartForm(libhttp.MaxFormMemory); err != nil {
					if err != http.ErrNotMultipart {
						result.ErrSystem(result.E(errors.WithMessage(err, "表单解析失败"))).JSON(w, r)
						return
					}
				}

				body = r.Form.Encode()
			case ContentJSON:
				b, err := io.ReadAll(r.Body) // 取出Body
				if err != nil {
					result.ErrSystem(result.E(errors.WithMessage(err, "请求Body读取失败"))).JSON(w, r)
					return
				}
				r.Body.Close() // 关闭原Body

				body = string(pretty.Ugly(b))
				r.Body = io.NopCloser(bytes.NewReader(b)) // 重新赋值Body
			}
		}

		next.ServeHTTP(w, r)

		log.Info(r.Context(), "request info",
			zap.String("method", r.Method),
			zap.String("uri", r.URL.String()),
			zap.String("ip", r.RemoteAddr),
			zap.String("body", body),
			zap.String("identity", auth.GetIdentity(r.Context()).String()),
			zap.String("duration", time.Since(now).String()),
		)
	})
}
