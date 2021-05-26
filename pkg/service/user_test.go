package service

import (
	"context"
	"testing"
	"tplgo/internal/result"
	"tplgo/pkg/consts"
	"tplgo/pkg/mock"
	"tplgo/pkg/models"
	"tplgo/pkg/response"

	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
)

func TestUserInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userDao := mock.NewMockUserDao(ctrl)

	userDao.EXPECT().FindByID(int64(1)).Return(&models.User{
		ID:           1,
		Nickname:     "shenghui",
		Avatar:       "avatar.jpg",
		Gender:       consts.Male,
		Phone:        "13912999014",
		RegisteredAt: 1616979600,
	}, nil)

	s := &user{
		userdao: userDao,
	}

	r := s.Info(context.TODO(), 1)

	assert.Equal(t, result.OK.Wrap(result.WithData(&response.UserInfo{
		ID:     1,
		Name:   "shenghui",
		Avatar: "avatar.jpg",
	})), r)
}
