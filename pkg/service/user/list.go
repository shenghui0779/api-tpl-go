package user

import (
	"api/db"
	"api/db/ent"
	"api/db/ent/user"
	"api/lib/util"
	"api/logger"
	"api/pkg/result"
	"api/pkg/service/internal"

	"net/http"
	"time"

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

func List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	builder := db.Client().User.Query()

	if v, ok := internal.URLQuery(r, "username"); ok && len(v) != 0 {
		builder.Where(user.UsernameContains(v))
	}

	total := 0
	offset, limit := internal.QueryPage(r)

	// 仅第一页返回数量并由前端保存
	if offset == 0 {
		var err error

		total, err = builder.Clone().Unique(false).Count(ctx)
		if err != nil {
			logger.Err(ctx, "err count user", zap.Error(err))
			result.ErrSystem().JSON(w, r)

			return
		}
	}

	records, err := builder.Unique(false).Select(
		user.FieldID,
		user.FieldUsername,
		user.FieldLoginAt,
		user.FieldCreatedAt,
	).Order(ent.Desc(user.FieldID)).Offset(offset).Limit(limit).All(ctx)
	if err != nil {
		logger.Err(ctx, "err query user", zap.Error(err))
		result.ErrSystem().JSON(w, r)

		return
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
			CreatedAtStr: util.TimeToStr(v.CreatedAt, time.DateTime),
		}

		if v.LoginAt != 0 {
			data.LoginAtStr = util.TimeToStr(v.LoginAt, time.DateTime)
		}

		resp.List = append(resp.List, data)
	}

	result.OK(result.V(resp)).JSON(w, r)
}
