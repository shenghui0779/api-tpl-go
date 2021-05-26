package service

import (
	"context"
	"tplgo/internal/result"
	"tplgo/pkg/dao"
	"tplgo/pkg/response"

	"github.com/pkg/errors"
)

type UserService interface {
	Info(ctx context.Context, id int64) result.Result
}

func NewUserService() UserService {
	return &user{
		userdao: dao.NewUserDao(),
	}
}

type user struct {
	userdao dao.UserDao
}

func (u *user) Info(ctx context.Context, id int64) result.Result {
	record, err := u.userdao.FindByID(id)

	if err != nil {
		return result.ErrSystem.Wrap(result.WithErr(errors.Wrap(err, "Service.User.Info 用户查询失败")))
	}

	if record == nil {
		return result.ErrUserNotFound
	}

	return result.OK.Wrap(result.WithData(&response.UserInfo{
		ID:     id,
		Name:   record.Nickname,
		Avatar: record.Avatar,
	}))
}
