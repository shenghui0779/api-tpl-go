package user

import (
	"api/ent"
	"api/ent/user"
	"api/lib/db"
	"api/lib/log"
	"api/pkg/internal"
	"api/pkg/result"

	"context"
	"net/url"
	"time"

	"github.com/pkg/errors"
	"github.com/shenghui0779/yiigo"
	"go.uber.org/zap"
)

type RespList struct {
	Total int         `json:"total"`
	List  []*UserInfo `json:"list"`
}

type UserInfo struct {
	ID           int64  `json:"id"`
	Username     string `json:"username"`
	LoginAt      int64  `json:"login_at"`
	LoginAtStr   string `json:"login_at_str"`
	CreatedAt    int64  `json:"created_at"`
	CreatedAtStr string `json:"created_at_str"`
}

func List(ctx context.Context, query url.Values) result.Result {
	builder := db.Client().User.Query()
	if v, ok := internal.URLQuery(query, "username"); ok && len(v) != 0 {
		builder.Where(user.UsernameContains(v))
	}

	total := 0
	offset, limit := internal.QueryPage(ctx, query)

	// 仅第一页返回数量并由前端保存
	if offset == 0 {
		var err error

		total, err = builder.Clone().Unique(false).Count(ctx)
		if err != nil {
			log.Error(ctx, "error count user", zap.Error(err))
			return result.ErrSystem(result.E(errors.WithMessage(err, "用户Count失败")))
		}
	}

	records, err := builder.Unique(false).Select(
		user.FieldID,
		user.FieldUsername,
		user.FieldLoginAt,
		user.FieldCreatedAt,
	).Order(ent.Desc(user.FieldID)).Offset(offset).Limit(limit).All(ctx)
	if err != nil {
		log.Error(ctx, "error query user", zap.Error(err))
		return result.ErrSystem(result.E(errors.WithMessage(err, "用户查询失败")))
	}

	resp := &RespList{
		Total: total,
		List:  make([]*UserInfo, 0, len(records)),
	}

	for _, v := range records {
		data := &UserInfo{
			ID:           v.ID,
			Username:     v.Username,
			LoginAt:      v.LoginAt,
			CreatedAt:    v.CreatedAt,
			CreatedAtStr: yiigo.TimeToStr(v.CreatedAt, time.DateTime),
		}

		if v.LoginAt != 0 {
			data.LoginAtStr = yiigo.TimeToStr(v.LoginAt, time.DateTime)
		}

		resp.List = append(resp.List, data)
	}

	return result.OK(result.V(resp))
}
