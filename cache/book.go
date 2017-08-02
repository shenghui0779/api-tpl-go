package cache

import (
	"encoding/json"

	"github.com/iiinsomnia/yiigo"
)

type BookCache struct {
	yiigo.Redis
}

func NewBookCache() *BookCache {
	return &BookCache{
		yiigo.Redis{},
	}
}

// GetBook 获取缓存
func (a *BookCache) GetBook(id int, data interface{}) bool {
	reply, err := a.Redis.Do("HGET", "slim:book:detail", id)

	if err != nil {
		yiigo.LogError(err.Error())
		return false
	}

	if reply == nil {
		return false
	}

	err = a.Redis.ScanJSON(reply, data)

	if err != nil {
		yiigo.LogError(err.Error())
		return false
	}

	return true
}

// SetBook 设置缓存
func (a *BookCache) SetBook(id int, data interface{}) bool {
	cache, err := json.Marshal(data)

	if err != nil {
		yiigo.LogError(err.Error())
		return false
	}

	_, err = a.Redis.Do("HSET", "slim:book:detail", id, cache)

	if err != nil {
		yiigo.LogError(err.Error())
		return false
	}

	return true
}

// DelBook 删除缓存
func (a *BookCache) DelBook(id int) bool {
	_, err := a.Redis.Do("HDEL", "slim:book:detail", id)

	if err != nil {
		yiigo.LogError(err.Error())
		return false
	}

	return true
}
