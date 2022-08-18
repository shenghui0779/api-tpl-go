package middlewares

import (
	"context"
	"net/http"
	"runtime/debug"

	"go.uber.org/zap"

	"tplgo/pkg/lib"
	"tplgo/pkg/logger"
	"tplgo/pkg/result"
)

// Recovery panic recover middleware
func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			// panic 捕获
			if err := recover(); err != nil {
				logger.Err(r.Context(), "Server Panic", zap.Any("error", err), zap.ByteString("stack", debug.Stack()))
				result.ErrSystem().JSON(w, r)
			}
		}()

		token := r.Header.Get("Authorization")

		if len(token) != 0 {
			ctx := r.Context()

			identity, err := lib.ParseAuthToken(token)

			if err != nil {
				logger.Err(ctx, "err middleware recovery (parse auth_token)", zap.Error(err))
				next.ServeHTTP(w, r)

				return
			}

			logger.Info(ctx, "[AUTH] identity", zap.Int64("id", identity.ID()), zap.String("token", identity.Token()))

			next.ServeHTTP(w, r.WithContext(context.WithValue(ctx, lib.AuthIdentityKey, identity)))

			return
		}

		next.ServeHTTP(w, r)
	})
}
