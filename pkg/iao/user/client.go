package user

import (
	"context"

	"github.com/shenghui0779/yiigo"

	"tplgo/pkg/iao/base"
	"tplgo/pkg/result"
)

type UserIao interface {
	UserInfo(ctx context.Context, params *ParamsUserInfo) (*UserInfo, result.Result)
}

func New() UserIao {
	return &apiUser{
		client: base.NewClient(yiigo.Env("iao.user").String()),
	}
}

type apiUser struct {
	client base.Client
}

func (u *apiUser) UserInfo(ctx context.Context, params *ParamsUserInfo) (*UserInfo, result.Result) {
	data := new(UserInfo)

	if er := u.client.Post(ctx, "/users/info", base.NewJSONBody(params), base.NewResponse(data)); er != nil {
		return nil, er
	}

	return data, nil
}
