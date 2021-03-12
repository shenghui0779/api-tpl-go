package cache

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gomodule/redigo/redis"
	"github.com/shenghui0779/yiigo"
	"go.uber.org/zap"

	"github.com/shenghui0779/demo/models"
)

type BookCache interface {
	Get(ctx context.Context) (*models.Book, bool)
	Set(ctx context.Context, data *models.Book) bool
	Del(ctx context.Context) bool
}

type book struct {
	pool *yiigo.RedisPoolResource
	key  string
}

func NewBookCache(id int64) BookCache {
	return &book{
		pool: yiigo.Redis(),
		key:  fmt.Sprintf("yiigo:books:%d", id),
	}
}

// Get 获取缓存
func (b *book) Get(ctx context.Context) (*models.Book, bool) {
	conn, err := b.pool.Get(ctx)

	if err != nil {
		yiigo.Logger().Error("get book cache error", zap.Error(err))

		return nil, false
	}

	defer b.pool.Put(conn)

	bs, err := redis.Bytes(conn.Do("GET", b.key))

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
func (b *book) Set(ctx context.Context, data *models.Book) bool {
	conn, err := b.pool.Get(ctx)

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

	if _, err := conn.Do("SET", b.key, bs); err != nil {
		yiigo.Logger().Error("set book cache error", zap.Error(err))

		return false
	}

	return true
}

// Del 删除缓存
func (b *book) Del(ctx context.Context) bool {
	conn, err := b.pool.Get(ctx)

	if err != nil {
		yiigo.Logger().Error("delete book cache error", zap.Error(err))

		return false
	}

	defer b.pool.Put(conn)

	if _, err := conn.Do("DEL", b.key); err != nil {
		yiigo.Logger().Error("delete book cache error", zap.Error(err))

		return false
	}

	return true
}
