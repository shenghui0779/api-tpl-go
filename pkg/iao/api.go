package iao

import "tplgo/pkg/iao/user"

func NewUserIao() user.UserIao {
	return user.New()
}
