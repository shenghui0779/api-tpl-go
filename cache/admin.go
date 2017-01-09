package cache

import (
	"models"

	"github.com/iiinsomnia/yiigo"
)

type AdminCache struct {
	yiigo.RedisBase
}

func NewAdminCache() *AdminCache {
	return &AdminCache{yiigo.RedisBase{CacheName: "admin"}}
}

func (a *AdminCache) GetAdminDetailCache(field string, data *models.AdminModel) bool {
	return a.RedisBase.HGet("detail", field, data)
}

func (a *AdminCache) SetAdminDetailCache(field string, data *models.AdminModel) bool {
	return a.RedisBase.HSet("detail", field, data)
}

func (a *AdminCache) DelAdminDetailCache(field string) bool {
	return a.RedisBase.HDel("detail", field)
}
