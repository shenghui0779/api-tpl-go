package service

import (
	"github.com/iiinsomnia/yiigo_demo/cache"
	"github.com/iiinsomnia/yiigo_demo/dao"
	"github.com/iiinsomnia/yiigo_demo/helpers"
	"github.com/iiinsomnia/yiigo_demo/reply"
)

type BookAdd struct {
	Title       string  `json:"title" valid:"required~title必填"`
	SubTitle    string  `json:"subtitle" valid:"required~subtitle必填"`
	Author      string  `json:"author" valid:"required~author必填"`
	Version     string  `json:"version" valid:"required~version必填"`
	Price       float64 `json:"price" valid:"required~price必填"`
	Publisher   string  `json:"publisher" valid:"required~publisher必填"`
	PublishDate string  `json:"publish_date" valid:"required~publish_date必填"`
}

func (b *BookAdd) Do() error {
	createData := &dao.BookCreateData{
		Title:       b.Title,
		SubTitle:    b.SubTitle,
		Author:      b.Author,
		Version:     b.Version,
		Price:       b.Price,
		Publisher:   b.Publisher,
		PublishDate: b.PublishDate,
	}

	bookDao := dao.NewBook()

	if err := bookDao.Create(createData); err != nil {
		return helpers.Error(10101, err)
	}

	return nil
}

type BookInfo struct {
	ID int64 `json:"id" valid:"required~id必填"`
}

func (b *BookInfo) Do() (*reply.BookInfoReply, error) {
	bookCache := cache.NewBook()

	book, ok := bookCache.Get(b.ID)

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

		bookCache.Set(b.ID, book)
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

type BookEdit struct {
	ID          int64   `json:"id" valid:"required~id必填"`
	Title       string  `json:"title" valid:"required~title必填"`
	SubTitle    string  `json:"subtitle" valid:"required~subtitle必填"`
	Author      string  `json:"author" valid:"required~author必填"`
	Version     string  `json:"version" valid:"required~version必填"`
	Price       float64 `json:"price" valid:"required~price必填"`
	Publisher   string  `json:"publisher" valid:"required~publisher必填"`
	PublishDate string  `json:"publish_date" valid:"required~publish_date必填"`
}

func (b *BookEdit) Do() error {
	updateData := &dao.BookUpdateData{
		Title:       b.Title,
		SubTitle:    b.SubTitle,
		Author:      b.Author,
		Version:     b.Version,
		Price:       b.Price,
		Publisher:   b.Publisher,
		PublishDate: b.PublishDate,
	}

	bookDao := dao.NewBook()

	if err := bookDao.UpdateByID(b.ID, updateData); err != nil {
		return helpers.Error(10102, err)
	}

	return nil
}

type BookDelete struct {
	ID int64 `json:"id" valid:"required~id必填"`
}

func (b *BookDelete) Do() error {
	bookDao := dao.NewBook()

	if err := bookDao.DeleteByID(b.ID); err != nil {
		return helpers.Error(10103, err)
	}

	return nil
}
