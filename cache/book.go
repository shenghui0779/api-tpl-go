package cache

import (
	"demo/models"
	"encoding/json"

	"github.com/iiinsomnia/yiigo"
)

// GetBook 获取缓存
func GetBookCache(id int, data *models.Book) bool {
	redis, err := yiigo.RedisConn()

	if err != nil {
		yiigo.Err(err.Error())

		return false
	}

	r, err := redis.Do("HGET", "slim:books", id)

	if err != nil {
		yiigo.Err(err.Error())

		return false
	}

	if r == nil {
		return false
	}

	err = yiigo.ScanJSON(r, data)

	if err != nil {
		yiigo.Err(err.Error())

		return false
	}

	return true
}

// SetBook 设置缓存
func SetBookCache(id int, data *models.Book) bool {
	redis, err := yiigo.RedisConn()

	if err != nil {
		yiigo.Err(err.Error())

		return false
	}

	cache, err := json.Marshal(data)

	if err != nil {
		yiigo.Err(err.Error())

		return false
	}

	_, err = redis.Do("HSET", "slim:books", id, cache)

	if err != nil {
		yiigo.Err(err.Error())

		return false
	}

	return true
}

// DelBook 删除缓存
func DelBookCache(id int) bool {
	redis, err := yiigo.RedisConn()

	if err != nil {
		yiigo.Err(err.Error())

		return false
	}

	_, err = redis.Do("HDEL", "slim:books", id)

	if err != nil {
		yiigo.Err(err.Error())

		return false
	}

	return true
}
