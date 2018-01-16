package cache

import (
	"demo/models"
	"encoding/json"

	"github.com/iiinsomnia/yiigo"
)

// GetBookCache 获取缓存
func GetBookCache(id int, data *models.Book) bool {
	r := yiigo.Redis.Cmd("HGET", "slim:books", id)

	if r.Err != nil {
		yiigo.Logger.Error(r.Err.Error())

		return false
	}

	if r == nil {
		return false
	}

	err := yiigo.ScanJSON(r, data)

	if err != nil {
		yiigo.Logger.Error(err.Error())

		return false
	}

	return true
}

// SetBookCache 设置缓存
func SetBookCache(id int, data *models.Book) bool {
	cache, err := json.Marshal(data)

	if err != nil {
		yiigo.Logger.Error(err.Error())

		return false
	}

	r := yiigo.Redis.Cmd("HSET", "slim:books", id, cache)

	if r.Err != nil {
		yiigo.Logger.Error(r.Err.Error())

		return false
	}

	return true
}

// DelBookCache 删除缓存
func DelBookCache(id int) bool {
	r := yiigo.Redis.Cmd("HDEL", "slim:books", id)

	if r.Err != nil {
		yiigo.Logger.Error(r.Err.Error())

		return false
	}

	return true
}
