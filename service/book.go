package service

import (
	"github.com/iiinsomnia/demo/cache"
	"github.com/iiinsomnia/demo/dao"
	"github.com/iiinsomnia/demo/helpers"
	"github.com/iiinsomnia/demo/reply"
)

type BookInfo struct {
	ID int64 `json:"id" valid:"required"`
}

func (b *BookInfo) Do() (*reply.BookInfoReply, error) {
	bookCache := cache.NewBook(b.ID)

	book, ok := bookCache.Get()

	if !ok {
		var err error

		bookDao := dao.NewBook()

		book, err = bookDao.FindByID(b.ID)

		if err != nil {
			return nil, helpers.Error(50000, err)
		}

		if book == nil {
			return nil, helpers.Error(10100)
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
