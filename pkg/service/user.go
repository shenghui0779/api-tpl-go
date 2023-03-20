package service

import (
	"errors"
	"net/http"
	"time"

	"github.com/shenghui0779/yiigo"
	"go.uber.org/zap"

	"tplgo/pkg/ent"
	"tplgo/pkg/ent/user"
	"tplgo/pkg/lib"
	"tplgo/pkg/logger"
	"tplgo/pkg/result"
)

type User interface {
	Create(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
}

func NewUser() User {
	return new(users)
}

type users struct{}

type ParamsUserCreate struct {
	Username string `json:"username" valid:"required"`
	Password string `json:"password" valid:"required"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}

func (u *users) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	params := new(ParamsUserCreate)

	if err := lib.BindJSON(r, params); err != nil {
		logger.Err(ctx, "err params", zap.Error(err))
		result.ErrParams().JSON(w, r)

		return
	}

	userRecords, err := ent.DB.User.Query().Unique(false).Select(user.FieldID).Where(user.Username(params.Username)).All(ctx)

	if err != nil {
		logger.Err(ctx, "err query user", zap.Error(err))
		result.ErrParams().JSON(w, r)

		return
	}

	if len(userRecords) != 0 {
		result.ErrParams(result.Err(errors.New("该用户名已被使用"))).JSON(w, r)

		return
	}

	now := time.Now().Unix()
	salt := lib.Nonce(16)

	_, err = ent.DB.User.Create().
		SetUsername(params.Username).
		SetPassword(yiigo.MD5(params.Password + salt)).
		SetSalt(salt).
		SetPhone(params.Phone).
		SetEmail(params.Email).
		SetRegistedAt(now).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Save(ctx)

	if err != nil {
		logger.Err(ctx, "err create user", zap.Error(err))
		result.ErrSystem().JSON(w, r)

		return
	}

	result.OK().JSON(w, r)
}

type RespUserList struct {
	Total int             `json:"total"`
	List  []*UserListData `json:"list"`
}

type UserListData struct {
	ID             int64  `json:"id"`
	Username       string `json:"username"`
	Phone          string `json:"phone"`
	Email          string `json:"email"`
	LastLoginAt    int64  `json:"last_login_at"`
	LastLoginAtStr string `json:"last_login_at_str"`
	CreatedAt      int64  `json:"created_at"`
	CreatedAtStr   string `json:"created_at_str"`
}

func (u *users) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	builder := ent.DB.User.Query()

	if v, ok := lib.URLQuery(r, "username"); ok && len(v) != 0 {
		builder.Where(user.UsernameContains(v))
	}

	total := 0

	offset, limit := lib.QueryPage(r)

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
		user.FieldPhone,
		user.FieldEmail,
		user.FieldLastLoginAt,
		user.FieldCreatedAt,
	).Order(ent.Desc(user.FieldID)).Offset(offset).Limit(limit).All(ctx)

	if err != nil {
		logger.Err(ctx, "err query user", zap.Error(err))
		result.ErrSystem().JSON(w, r)

		return
	}

	resp := &RespUserList{
		Total: total,
		List:  make([]*UserListData, 0, len(records)),
	}

	for _, v := range records {
		data := &UserListData{
			ID:           v.ID,
			Username:     v.Username,
			Phone:        v.Phone,
			Email:        v.Email,
			LastLoginAt:  v.LastLoginAt,
			CreatedAt:    v.CreatedAt,
			CreatedAtStr: yiigo.Date(v.CreatedAt),
		}

		if v.LastLoginAt != 0 {
			data.LastLoginAtStr = yiigo.Date(v.LastLoginAt)
		}

		resp.List = append(resp.List, data)
	}

	result.OK(result.Data(resp)).JSON(w, r)
}
