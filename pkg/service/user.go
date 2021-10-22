package service

import (
	"context"

	"go.uber.org/zap"

	"tplgo/pkg/dao"
	"tplgo/pkg/logger"
	"tplgo/pkg/response"
	"tplgo/pkg/result"
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
	record, err := u.userdao.FindByID(ctx, id)

	if err != nil {
		logger.Err(ctx, "Service.User.Info error", zap.Error(err))

		return result.ErrSystem.Wrap(result.WithErr(err))
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
