package dao

import (
	"database/sql"
	"time"

	"github.com/iiinsomnia/yiigo/v4"
	"github.com/iiinsomnia/yiigo_demo/models"
	"github.com/jinzhu/gorm"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Book struct {
	db  *sqlx.DB
	orm *gorm.DB
}

func NewBook() *Book {
	return &Book{
		db:  yiigo.DB(),
		orm: yiigo.Orm(),
	}
}

func (b *Book) FindByID(id int64) (*models.Book, error) {
	record := new(models.Book)

	if err := b.db.Get(record, "SELECT * FROM `book` WHERE `id` = ? LIMIT 1", id); err != nil {
		if err != sql.ErrNoRows {
			return nil, errors.Wrap(err, "find book by id error")
		}

		return nil, nil
	}

	return record, nil
}

type BookCreateData struct {
	Title       string
	SubTitle    string
	Author      string
	Version     string
	Price       float64
	Publisher   string
	PublishDate string
}

func (b *Book) Create(data *BookCreateData) error {
	now := time.Now().Unix()

	createData := yiigo.X{
		"title":        data.Title,
		"subtitle":     data.SubTitle,
		"author":       data.Author,
		"version":      data.Version,
		"price":        data.Price,
		"publisher":    data.Publisher,
		"publish_date": data.PublishDate,
		"created_at":   now,
		"updated_at":   now,
	}

	query, binds := yiigo.InsertSQL("book", createData)

	if _, err := b.db.Exec(query, binds...); err != nil {
		return errors.Wrap(err, "create new book error")
	}

	return nil
}

type BookUpdateData struct {
	Title       string
	SubTitle    string
	Author      string
	Version     string
	Price       float64
	Publisher   string
	PublishDate string
}

func (b *Book) UpdateByID(id int64, data *BookUpdateData) error {
	now := time.Now().Unix()

	updateData := yiigo.X{
		"title":        data.Title,
		"subtitle":     data.SubTitle,
		"author":       data.Author,
		"version":      data.Version,
		"price":        data.Price,
		"publisher":    data.Publisher,
		"publish_date": data.PublishDate,
		"updated_at":   now,
	}

	query, binds := yiigo.UpdateSQL("UPDATE `book` SET ? WHERE `id` = ?", updateData, id)

	if _, err := b.db.Exec(query, binds...); err != nil {
		return errors.Wrap(err, "update book by id error")
	}

	return nil
}

func (b *Book) DeleteByID(id int64) error {
	if _, err := b.db.Exec("DELETE FROM book WHERE id = ?", id); err != nil {
		return errors.Wrap(err, "delete book by id error")
	}

	return nil
}
