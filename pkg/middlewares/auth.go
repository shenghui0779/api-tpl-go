package middlewares

import (
	"errors"
	"net/http"

	"tplgo/pkg/ent"
	"tplgo/pkg/ent/user"
	"tplgo/pkg/lib"
	"tplgo/pkg/logger"
	"tplgo/pkg/result"

	"go.uber.org/zap"
)

// Auth App授权中间件
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		identity := lib.GetIdentity(ctx)

		if identity.ID() == 0 {
			result.ErrAuth(result.Err(errors.New("未授权，请先登录"))).JSON(w, r)

			return
		}

		record, err := ent.DB.User.Query().Unique(false).Select(user.FieldID, user.FieldLoginToken).Where(user.ID(identity.ID())).First(ctx)

		if err != nil {
			logger.Err(ctx, "err middleware auth (query user)", zap.Error(err))
			result.ErrAuth(result.Err(errors.New("内部服务器错误"))).JSON(w, r)

			return
		}

		if len(record.LoginToken) == 0 || record.LoginToken != identity.Token() {
			result.ErrAuth(result.Err(errors.New("授权已失效"))).JSON(w, r)

			return
		}

		next.ServeHTTP(w, r)
	})
}
