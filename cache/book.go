package cache

import "github.com/iiinsomnia/yiigo"

type BookCache struct {
	yiigo.Redis
}

func NewBookCache() *BookCache {
	return &BookCache{
		yiigo.Redis{
			CacheName: "book",
		},
	}
}

// GetBook 获取缓存
func (a *BookCache) GetBook(field string, data interface{}) bool {
	err := a.Redis.HGet("detail", field, data)

	if err != nil {
		if err.Error() != "not found" {
			yiigo.LogError(err.Error())
		}

		return false
	}

	return true
}

// SetBook 设置缓存
func (a *BookCache) SetBook(field string, data interface{}) bool {
	err := a.Redis.HSet("detail", field, data)

	if err != nil {
		yiigo.LogError(err.Error())
		return false
	}

	return true
}

// DelBook 删除缓存
func (a *BookCache) DelBook(field string) bool {
	err := a.Redis.HDel("detail", field)

	if err != nil {
		yiigo.LogError(err.Error())
		return false
	}

	return true
}
