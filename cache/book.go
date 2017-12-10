package cache

import (
	"demo/models"
	"encoding/json"

	"github.com/iiinsomnia/yiigo"
)

// GetBookCache 获取缓存
func GetBookCache(id int, data *models.Book) bool {
	redis, err := yiigo.RedisPool.Get()

	if err != nil {
		yiigo.Logger.Error(err.Error())

		return false
	}

	defer yiigo.RedisPool.Put(redis)

	r, err := redis.Do("HGET", "slim:books", id)

	if err != nil {
		yiigo.Logger.Error(err.Error())

		return false
	}

	if r == nil {
		return false
	}

	err = yiigo.ScanJSON(r, data)

	if err != nil {
		yiigo.Logger.Error(err.Error())

		return false
	}

	return true
}

// SetBookCache 设置缓存
func SetBookCache(id int, data *models.Book) bool {
	redis, err := yiigo.RedisPool.Get()

	if err != nil {
		yiigo.Logger.Error(err.Error())

		return false
	}

	defer yiigo.RedisPool.Put(redis)

	cache, err := json.Marshal(data)

	if err != nil {
		yiigo.Logger.Error(err.Error())

		return false
	}

	_, err = redis.Do("HSET", "slim:books", id, cache)

	if err != nil {
		yiigo.Logger.Error(err.Error())

		return false
	}

	return true
}

// DelBookCache 删除缓存
func DelBookCache(id int) bool {
	redis, err := yiigo.RedisPool.Get()

	if err != nil {
		yiigo.Logger.Error(err.Error())

		return false
	}

	defer yiigo.RedisPool.Put(redis)

	_, err = redis.Do("HDEL", "slim:books", id)

	if err != nil {
		yiigo.Logger.Error(err.Error())

		return false
	}

	return true
}
