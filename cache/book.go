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

func (a *BookCache) GetBook(field string, data interface{}) bool {
	return a.Redis.HGet("detail", field, data)
}

func (a *BookCache) SetBook(field string, data interface{}) bool {
	return a.Redis.HSet("detail", field, data)
}

func (a *BookCache) DelBook(field string) bool {
	return a.Redis.HDel("detail", field)
}
