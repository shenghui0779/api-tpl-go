package middlewares

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/shenghui0779/yiigo"
	"github.com/tidwall/pretty"
	"go.uber.org/zap"

	"tplgo/pkg/consts"
	"tplgo/pkg/logger"
	"tplgo/pkg/result"
)

var (
	bufPool = sync.Pool{
		New: func() interface{} {
			return bytes.NewBuffer(make([]byte, 0, 4<<10)) // 4KB
		},
	}
)

type loggercfg struct {
	nobody bool
	noresp bool
}

type LoggerOption func(cfg *loggercfg)

func NoBody() LoggerOption {
	return func(cfg *loggercfg) {
		cfg.nobody = true
	}
}

func NoResp() LoggerOption {
	return func(cfg *loggercfg) {
		cfg.noresp = true
	}
}

// Logger 日志中间件
func Logger(options ...LoggerOption) func(next http.Handler) http.Handler {
	cfg := new(loggercfg)

	for _, f := range options {
		f(cfg)
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			now := time.Now()

			var params string

			// 需要记录参数 且 请求包含body
			if !cfg.nobody && r.Body != nil && r.Body != http.NoBody {
				switch yiigo.ContentType(r) {
				case consts.MIMEForm:
					if err := r.ParseForm(); err != nil {
						result.ErrSystem(result.Err(err)).JSON(w, r)

						return
					}

					params = r.Form.Encode()
				case consts.MIMEMultipartForm:
					if err := r.ParseMultipartForm(consts.MaxFormMemory); err != nil {
						if err != http.ErrNotMultipart {
							result.ErrSystem(result.Err(err)).JSON(w, r)

							return
						}
					}

					params = r.Form.Encode()
				default:
					// 取出Body
					body, err := ioutil.ReadAll(r.Body)

					if err != nil {
						result.ErrSystem(result.Err(err)).JSON(w, r)

						return
					}

					// 关闭原Body
					r.Body.Close()

					params = string(pretty.Ugly(body))

					r.Body = ioutil.NopCloser(bytes.NewReader(body))
				}
			}

			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			var buf *bytes.Buffer

			// 需要记录返回
			if !cfg.noresp {
				buf = bufPool.Get().(*bytes.Buffer)
				buf.Reset()

				defer bufPool.Put(buf)

				ww.Tee(buf)
			}

			next.ServeHTTP(ww, r)

			logger.Info(r.Context(), fmt.Sprintf("[%s] %s", r.Method, r.URL.String()),
				zap.String("params", params),
				zap.String("response", buf.String()),
				zap.Int("status", ww.Status()),
				zap.String("duration", time.Since(now).String()),
			)
		})
	}
}
