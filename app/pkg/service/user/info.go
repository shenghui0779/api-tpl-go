package user

import (
	"context"
	"time"

	"api/app/ent"
	"api/app/ent/user"
	"api/lib/result"
)

type UserInfo struct {
	// 自增ID
	ID int64 `json:"id,omitempty"`
	// 创建时间
	CreatedAt string `json:"created_at,omitempty"`
	// 更新时间
	UpdatedAt string `json:"updated_at,omitempty"`
	// 账号
	Account string `json:"account,omitempty"`
	// 用户名
	Username string `json:"username,omitempty"`
	// 密码
	Password string `json:"password,omitempty"`
	// 加密盐
	Salt string `json:"salt,omitempty"`
	// 登录时间
	LoginAt string `json:"login_at,omitempty"`
	// 登录Token
	LoginToken string `json:"login_token,omitempty"`
}

func EntToUserInfo(e *ent.User) *UserInfo {
	info := &UserInfo{
		ID:         e.ID,
		Account:    e.Account,
		Password:   e.Password,
		Salt:       e.Salt,
		LoginToken: e.LoginToken,
	}
	if !e.CreatedAt.IsZero() {
		info.CreatedAt = e.CreatedAt.In(time.Local).Format(time.DateTime)
	}
	if !e.UpdatedAt.IsZero() {
		info.UpdatedAt = e.UpdatedAt.In(time.Local).Format(time.DateTime)
	}
	if e.LoginAt.Valid && !e.LoginAt.Time.IsZero() {
		info.LoginAt = e.LoginAt.Time.In(time.Local).Format(time.DateTime)
	}
	return info
}

type ReqInfo struct {
	IDs    []int64  `json:"ids"`
	Fields []string `json:"fields"`
}

func Info(ctx context.Context, req *ReqInfo) result.Result {
	builder := ent.DB.User.Query().Unique(false)
	if len(req.IDs) != 0 {
		builder.Where(user.IDIn(req.IDs...))
	}
	if len(req.Fields) != 0 {
		builder.Select(req.Fields...)
	}
	records, err := builder.All(ctx)
	if err != nil {
		return result.ErrSystem(result.E(err))
	}
	resp := make(map[int64]*UserInfo, len(records))
	for _, v := range records {
		resp[v.ID] = EntToUserInfo(v)
	}
	return result.OK(result.V(resp))
}
