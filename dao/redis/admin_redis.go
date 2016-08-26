package redis

import (
	"models"

	"github.com/iiinsomnia/yiigo"
)

type AdminRedis struct {
	yiigo.RedisBase
}

func NewAdminRedis() *AdminRedis {
	return &AdminRedis{yiigo.RedisBase{CacheName: "test"}}
}

func (a *AdminRedis) SetAdminCache(id int, data models.AdminModel) bool {
	return a.RedisBase.HSet("admin", id, data)
}

func (a *AdminRedis) GetAdminCache(data *models.AdminModel, id int) bool {
	return a.RedisBase.HGet(data, "admin", id)
}
