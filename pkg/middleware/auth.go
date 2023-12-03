package middleware

import (
	"net/http"

	"api/pkg/auth"
	"api/pkg/result"

	"github.com/pkg/errors"
)

// Auth App授权中间件
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		identity := auth.GetIdentity(ctx)
		if err := identity.Check(ctx); err != nil {
			result.ErrAuth(result.E(errors.WithMessage(err, "授权校验失败"))).JSON(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}
