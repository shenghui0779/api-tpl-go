package middlewares

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/shenghui0779/demo/helpers"

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
				yiigo.Logger().Error(fmt.Sprintf("server panic: %v", err),
					zap.String("request_id", middleware.GetReqID(r.Context())),
					zap.String("url", r.URL.String()),
					zap.String("stack", string(debug.Stack())),
				)

				helpers.JSON(w, yiigo.X{
					"success": false,
					"code":    helpers.ErrSystem,
					"msg":     "服务器错误，请稍后重试",
				})
			}
		}()

		next.ServeHTTP(w, r)
	})
}
