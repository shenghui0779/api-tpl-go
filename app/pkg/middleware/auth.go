package middleware

import (
	"context"
	"net/http"

	"api/app/ent"
	"api/app/ent/user"
	"api/lib/identity"
	"api/lib/log"
	"api/lib/result"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// Auth App授权中间件
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		if err := check(ctx); err != nil {
			result.ErrAuth(result.E(errors.WithMessage(err, "授权校验失败"))).JSON(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func check(ctx context.Context) error {
	id := identity.FromContext(ctx)
	if id.ID() == 0 {
		return errors.New("未授权，请先登录")
	}

	recordAccount, err := ent.DB.User.Query().Unique(false).Select(user.FieldID, user.FieldLoginToken).Where(user.ID(id.ID())).First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return errors.New("无效的账号")
		}
		log.Error(ctx, "Error auth check", zap.Error(err))
		return errors.New("授权校验失败")
	}
	if len(recordAccount.LoginToken) == 0 || recordAccount.LoginToken != id.Token() {
		return errors.New("授权已失效")
	}
	return nil
}
