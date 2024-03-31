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
	"go.uber.org/zap"
)

type RespList struct {
	Total int         `json:"total"`
	List  []*UserInfo `json:"list"`
}

type UserInfo struct {
	ID        int64  `json:"id"`
	Phone     string `json:"phone"`
	Nickname  string `json:"nickname"`
	LoginAt   string `json:"login_at"`
	CreatedAt string `json:"created_at"`
}

func List(ctx context.Context, query url.Values) result.Result {
	builder := db.Client().User.Query()
	if v, ok := internal.URLQuery(query, "phone"); ok && len(v) != 0 {
		builder.Where(user.Phone(v))
	}

	total := 0
	offset, limit := internal.QueryPage(ctx, query)

	// 仅第一页返回数量并由前端保存
	if offset == 0 {
		var err error

		total, err = builder.Clone().Unique(false).Count(ctx)
		if err != nil {
			log.Error(ctx, "Error count user", zap.Error(err))
			return result.ErrSystem(result.E(errors.WithMessage(err, "用户Count失败")))
		}
	}

	records, err := builder.Unique(false).Select(
		user.FieldID,
		user.FieldPhone,
		user.FieldNickname,
		user.FieldLoginAt,
		user.FieldCreatedAt,
	).Order(ent.Desc(user.FieldID)).Offset(offset).Limit(limit).All(ctx)
	if err != nil {
		log.Error(ctx, "Error query user", zap.Error(err))
		return result.ErrSystem(result.E(errors.WithMessage(err, "用户查询失败")))
	}

	resp := &RespList{
		Total: total,
		List:  make([]*UserInfo, 0, len(records)),
	}

	for _, v := range records {
		data := &UserInfo{
			ID:        v.ID,
			Phone:     v.Phone,
			Nickname:  v.Nickname,
			LoginAt:   v.LoginAt.Format(time.DateTime),
			CreatedAt: v.CreatedAt.Format(time.DateTime),
		}

		resp.List = append(resp.List, data)
	}

	return result.OK(result.V(resp))
}
