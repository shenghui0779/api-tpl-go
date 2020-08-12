package service

import (
	"github.com/gin-gonic/gin"
	"github.com/shenghui0779/demo/cache"
	"github.com/shenghui0779/demo/dao"
	"github.com/shenghui0779/demo/helpers"
	"github.com/shenghui0779/demo/reply"
)

type BookInfo struct {
	ID int64 `json:"id" valid:"required"`
}

func (b *BookInfo) Do(ctx *gin.Context) (*reply.BookInfoReply, error) {
	bookCache := cache.NewBook(b.ID)

	book, ok := bookCache.Get()

	if !ok {
		var err error

		bookDao := dao.NewBook()

		book, err = bookDao.FindByID(b.ID)

		if err != nil {
			return nil, helpers.Error(ctx, helpers.ErrSystem, err)
		}

		if book == nil {
			return nil, helpers.Error(ctx, helpers.ErrBookNotFound)
		}

		bookCache.Set(book)
	}

	resp := &reply.BookInfoReply{
		Title:       book.Title,
		SubTitle:    book.SubTitle,
		Author:      book.Author,
		Version:     book.Version,
		Price:       book.Price,
		Publisher:   book.Publisher,
		PublishDate: book.PublishDate,
		CreatedAt:   book.CreatedAt,
		UpdatedAt:   book.UpdatedAt,
	}

	return resp, nil
}
