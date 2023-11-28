package middleware

import (
	"net/http"

	"api/pkg/auth"
	"api/pkg/result"
)

// Auth App授权中间件
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		identity := auth.GetIdentity(ctx)
		if err := identity.Check(ctx); err != nil {
			result.ErrAuth(result.M(err.Error())).JSON(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}
