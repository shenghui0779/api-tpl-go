package middlewares

import (
	"net/http"
	"runtime/debug"

	"go.uber.org/zap"

	"tplgo/pkg/logger"
	"tplgo/pkg/result"
)

// Recovery panic recover middleware
func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			// panic 捕获
			if err := recover(); err != nil {
				logger.Err(r.Context(), "Server Panic",
					zap.Any("error", err),
					zap.ByteString("stack", debug.Stack()),
				)

				result.ErrSystem().JSON(w, r)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
