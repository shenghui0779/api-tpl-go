package service

import (
	"net/http"

	"tplgo/pkg/ent"
	"tplgo/pkg/ent/user"
	"tplgo/pkg/helpers"
	"tplgo/pkg/logger"
	"tplgo/pkg/result"

	"go.uber.org/zap"
)

type User interface {
	Info(w http.ResponseWriter, r *http.Request)
}

func NewUser() User {
	return new(users)
}

type users struct{}

type RespUserInfo struct {
	ID       int64  `json:"id"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

func (u *users) Info(w http.ResponseWriter, r *http.Request) {
	uid := helpers.URLParamInt(r, "id")

	record, err := ent.DB.User.Query().Where(user.ID(uid)).First(r.Context())

	if err != nil {
		if ent.IsNotFound(err) {
			result.ErrNotFound().JSON(w, r)

			return
		}

		logger.Err(r.Context(), "err query user", zap.Error(err))

		result.ErrSystem(result.Err(err)).JSON(w, r)

		return
	}

	result.OK(result.Data(&RespUserInfo{
		ID:       uid,
		Nickname: record.Nickname,
		Avatar:   record.Avatar,
	})).JSON(w, r)
}
