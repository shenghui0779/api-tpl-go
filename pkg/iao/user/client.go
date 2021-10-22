package user

import (
	"context"
	"os"

	"tplgo/pkg/iao/base"

	"github.com/pkg/errors"
)

type UserIao interface {
	UserInfo(ctx context.Context, params *ParamsUserInfo) (*UserInfo, error)
}

func New() UserIao {
	return &apiUser{
		client: base.NewClient(os.Getenv("IAO_USER")),
	}
}

type apiUser struct {
	client base.Client
}

func (u *apiUser) UserInfo(ctx context.Context, params *ParamsUserInfo) (*UserInfo, error) {
	data := new(UserInfo)
	resp := base.NewResponse(data)

	if err := u.client.Post(ctx, "/users/info", base.NewJSONBody(params), resp); err != nil {
		return nil, errors.Wrap(err, "Iao.UserInfo error")
	}

	return data, nil
}
