package middleware

import (
	"context"
	"net/http"
	"runtime/debug"

	"go.uber.org/zap"

	"api/lib/identity"
	"api/lib/log"
	"api/lib/result"
)

// Recovery panic recover middleware
func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			// panic 捕获
			if err := recover(); err != nil {
				log.Error(r.Context(), "Server Panic", zap.Any("error", err), zap.ByteString("stack", debug.Stack()))
				result.ErrSystem().JSON(w, r)
			}
		}()

		if token := r.Header.Get("Authorization"); len(token) != 0 {
			ctx := r.Context()
			id := identity.FromAuthToken(ctx, token)

			next.ServeHTTP(w, r.WithContext(context.WithValue(ctx, identity.IdentityKey, id)))

			return
		}

		next.ServeHTTP(w, r)
	})
}
