package middlewares

import (
	"context"
	"net/http"

	"go.uber.org/zap"

	"tplgo/pkg/lib"
	"tplgo/pkg/logger"
	"tplgo/pkg/result"
)

// Auth App授权中间件
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if len(token) == 0 {
			result.ErrAuth().JSON(w, r)

			return
		}

		ctx := r.Context()

		identity, err := lib.VerifyAuthToken(ctx, token)

		if err != nil {
			result.ErrAuth(result.Err(err)).JSON(w, r)

			return
		}

		logger.Info(ctx, "req identity", zap.Int64("id", identity.ID()), zap.String("token", identity.Token()))

		next.ServeHTTP(w, r.WithContext(context.WithValue(ctx, lib.AuthIdentityKey, identity)))
	})
}
