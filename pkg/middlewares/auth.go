package middlewares

import (
	"context"
	"errors"
	"net/http"

	"tplgo/pkg/lib"
	"tplgo/pkg/result"
)

// Auth App授权中间件
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		token := r.Header.Get("Authorization")

		if len(token) == 0 {
			result.ErrAuth(result.Err(errors.New("未授权，请先登录"))).JSON(w, r)
		}

		identity := lib.AuthTokenToIdentity(ctx, token)

		if err := identity.Check(ctx); err != nil {
			result.ErrAuth(result.Err(err)).JSON(w, r)

			return
		}

		// 注入授权身份
		next.ServeHTTP(w, r.WithContext(context.WithValue(ctx, lib.AuthIdentityKey, identity)))
	})
}
