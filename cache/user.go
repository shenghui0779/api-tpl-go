package cache

import (
	"models"

	"github.com/iiinsomnia/yiigo"
)

type UserCache struct {
	yiigo.RedisBase
}

func NewUserCache() *UserCache {
	return &UserCache{yiigo.RedisBase{CacheName: "user"}}
}

func (a *UserCache) GetUserDetailCache(field string, data *models.UserModel) bool {
	return a.RedisBase.HGet("detail", field, data)
}

func (a *UserCache) SetUserDetailCache(field string, data *models.UserModel) bool {
	return a.RedisBase.HSet("detail", field, data)
}

func (a *UserCache) DelUserDetailCache(field string) bool {
	return a.RedisBase.HDel("detail", field)
}
