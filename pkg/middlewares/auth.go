package middlewares

import (
	"net/http"

	"tplgo/pkg/lib"
	"tplgo/pkg/result"
)

// Auth App授权中间件
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		identity := lib.GetIdentity(ctx)

		if err := identity.Check(ctx); err != nil {
			result.ErrAuth(result.Err(err)).JSON(w, r)

			return
		}

		next.ServeHTTP(w, r)
	})
}
