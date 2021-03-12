package service

import (
	"context"

	"github.com/pkg/errors"
	"github.com/shenghui0779/demo/cache"
	"github.com/shenghui0779/demo/dao"
	"github.com/shenghui0779/demo/helpers"
	"github.com/shenghui0779/demo/response"
)

type BookService interface {
	Info(ctx context.Context, id int64) (*response.BookInfo, error)
}

func NewBookService() BookService {
	return new(book)
}

type book struct{}

func (b *book) Info(ctx context.Context, id int64) (*response.BookInfo, error) {
	bookCache := cache.NewBookCache(id)

	record, ok := bookCache.Get(ctx)

	if !ok {
		var err error

		bookDao := dao.NewBookDao()

		record, err = bookDao.FindByID(id)

		if err != nil {
			return nil, errors.Wrap(err, "书籍查询失败")
		}

		if record == nil {
			return nil, helpers.Err(helpers.ErrInvalidBook)
		}

		bookCache.Set(ctx, record)
	}

	resp := &response.BookInfo{
		Title:       record.Title,
		SubTitle:    record.SubTitle,
		Author:      record.Author,
		Version:     record.Version,
		Price:       record.Price,
		Publisher:   record.Publisher,
		PublishDate: record.PublishDate,
		CreatedAt:   record.CreatedAt,
		UpdatedAt:   record.UpdatedAt,
	}

	return resp, nil
}
