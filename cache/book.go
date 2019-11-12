package cache

import (
	"encoding/json"

	"github.com/gomodule/redigo/redis"
	"github.com/iiinsomnia/yiigo/v4"
	"github.com/iiinsomnia/yiigo_demo/models"
	"go.uber.org/zap"
)

type Book struct {
	pool *yiigo.RedisPoolResource
	key  string
}

func NewBook() *Book {
	return &Book{
		pool: yiigo.Redis(),
		key:  "yiigo.books",
	}
}

// Get 获取缓存
func (b *Book) Get(id int64) (*models.Book, bool) {
	conn, err := b.pool.Get()

	if err != nil {
		yiigo.Logger().Error("get book cache error", zap.Error(err))

		return nil, false
	}

	defer b.pool.Put(conn)

	bs, err := redis.Bytes(conn.Do("HGET", b.key, id))

	if err != nil {
		if err != redis.ErrNil {
			yiigo.Logger().Error("get book cache error", zap.Error(err))
		}

		return nil, false
	}

	data := new(models.Book)

	if err := json.Unmarshal(bs, data); err != nil {
		yiigo.Logger().Error("get book cache error", zap.Error(err))

		return nil, false
	}

	return data, true
}

// Set 设置缓存
func (b *Book) Set(id int64, data *models.Book) bool {
	conn, err := b.pool.Get()

	if err != nil {
		yiigo.Logger().Error("set book cache error", zap.Error(err))

		return false
	}

	defer b.pool.Put(conn)

	bs, err := json.Marshal(data)

	if err != nil {
		yiigo.Logger().Error("set book cache error", zap.Error(err))

		return false
	}

	if _, err := conn.Do("HSET", b.key, id, bs); err != nil {
		yiigo.Logger().Error("set book cache error", zap.Error(err))

		return false
	}

	return true
}

// Del 删除缓存
func (b *Book) Del(id int64) bool {
	conn, err := b.pool.Get()

	if err != nil {
		yiigo.Logger().Error("delete book cache error", zap.Error(err))

		return false
	}

	defer b.pool.Put(conn)

	if _, err := conn.Do("HDEL", b.key, id); err != nil {
		yiigo.Logger().Error("delete book cache error", zap.Error(err))

		return false
	}

	return true
}
