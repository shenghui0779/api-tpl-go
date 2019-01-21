package cache

import (
	"demo/models"
	"encoding/json"

	"github.com/gomodule/redigo/redis"
	"github.com/iiinsomnia/yiigo"
)

// GetBookCache 获取缓存
func GetBookCache(id int, data *models.Book) bool {
	conn, err := yiigo.Redis.Get()

	if err != nil {
		yiigo.Logger.Error(err.Error())

		return false
	}

	defer yiigo.Redis.Put(conn)

	b, err := redis.Bytes(conn.Do("HGET", "yiigo:books", id))

	if err != nil {
		if err != redis.ErrNil {
			yiigo.Logger.Error(err.Error())
		}

		return false
	}

	if err := json.Unmarshal(b, data); err != nil {
		yiigo.Logger.Error(err.Error())

		return false
	}

	return true
}

// SetBookCache 设置缓存
func SetBookCache(id int, data *models.Book) bool {
	conn, err := yiigo.Redis.Get()

	if err != nil {
		yiigo.Logger.Error(err.Error())

		return false
	}

	defer yiigo.Redis.Put(conn)

	b, err := json.Marshal(data)

	if err != nil {
		yiigo.Logger.Error(err.Error())

		return false
	}

	if _, err := conn.Do("HSET", "yiigo:books", id, b); err != nil {
		yiigo.Logger.Error(err.Error())

		return false
	}

	return true
}

// DelBookCache 删除缓存
func DelBookCache(id int) bool {
	conn, err := yiigo.Redis.Get()

	if err != nil {
		yiigo.Logger.Error(err.Error())

		return false
	}

	defer yiigo.Redis.Put(conn)

	if _, err := conn.Do("HDEL", "yiigo:books", id); err != nil {
		yiigo.Logger.Error(err.Error())

		return false
	}

	return true
}
