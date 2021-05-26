package models

import "tplgo/pkg/consts"

type User struct {
	ID           int64         `db:"id"`
	Nickname     string        `db:"nickname"`
	Avatar       string        `db:"avatar"`
	Gender       consts.Gender `db:"gender"`
	Phone        string        `db:"phone"`
	RegisteredAt int64         `db:"registered_at"`
	LastLoginAt  int64         `db:"last_login_at"`
	LastLoginIP  int           `db:"last_login_ip"`
	CreatedAt    int64         `db:"created_at"`
	UpdatedAt    int64         `db:"updated_at"`
}

func (u *User) IsMan() bool {
	return u.Gender == consts.Male
}

func (u *User) IsWoman() bool {
	return u.Gender == consts.Female
}
