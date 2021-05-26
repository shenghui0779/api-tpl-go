package middlewares

import (
	"net/http"
	"runtime/debug"
	"tplgo/internal/result"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/shenghui0779/yiigo"
	"go.uber.org/zap"
)

// Recovery panic recover middleware
func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			// panic 捕获
			if err := recover(); err != nil {
				yiigo.Logger().Error("server panic",
					zap.String("request_id", middleware.GetReqID(r.Context())),
					zap.String("method", r.Method),
					zap.String("url", r.URL.String()),
					zap.Any("error", err),
					zap.ByteString("stack", debug.Stack()),
				)

				result.ErrSystem.JSON(w, r)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
