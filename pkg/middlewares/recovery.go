package middlewares

import (
	"context"
	"net/http"
	"runtime/debug"

	"go.uber.org/zap"

	"api/lib"
	"api/logger"
	"api/pkg/result"
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

		if token := r.Header.Get("Authorization"); len(token) != 0 {
			ctx := r.Context()
			identity := lib.AuthTokenToIdentity(ctx, token)

			next.ServeHTTP(w, r.WithContext(context.WithValue(ctx, lib.AuthIdentityKey, identity)))

			return
		}

		next.ServeHTTP(w, r)
	})
}
