package service

import (
	"database/sql"
	"demo/cache"
	"demo/models"
	"time"

	"github.com/iiinsomnia/yiigo"
)

func GetBookById(id int) (yiigo.X, error) {
	book := &models.Book{}

	ok := cache.GetBookCache(id, book)

	if !ok {
		err := yiigo.DB.Get(book, "SELECT * FROM book WHERE id = ?", id)

		if err != nil {
			if err != sql.ErrNoRows {
				yiigo.Logger.Error(err.Error())

				return nil, err
			}

			return yiigo.X{}, nil
		}

		cache.SetBookCache(id, book)
	}

	data := yiigo.X{
		"id":           book.ID,
		"title":        book.Title,
		"subtitle":     book.SubTitle,
		"author":       book.Author,
		"version":      book.Version,
		"price":        book.Price,
		"publisher":    book.Publisher,
		"publish_date": book.PublishDate,
	}

	return data, nil
}

func GetAllBooks() ([]models.Book, error) {
	var books []models.Book

	err := yiigo.DB.Select(&books, "SELECT * FROM book WHERE 1=1")

	if err != nil && err != sql.ErrNoRows {
		yiigo.Logger.Error(err.Error())

		return nil, err
	}

	return books, nil
}

func AddNewBook(data *models.BookAdd) (int64, error) {
	data.CreatedAt = time.Now()

	sql, binds := yiigo.InsertSQL("book", data)
	r, err := yiigo.DB.Exec(sql, binds...)

	if err != nil {
		yiigo.Logger.Error(err.Error())

		return 0, err
	}

	id, err := r.LastInsertId()

	return id, err
}

func UpdateBookById(id int, data *models.BookEdit) error {
	data.UpdatedAt = time.Now()

	sql, binds := yiigo.UpdateSQL("UPDATE book SET ? WHERE id = ?", data, id)
	_, err := yiigo.DB.Exec(sql, binds...)

	if err != nil {
		yiigo.Logger.Error(err.Error())

		return err
	}

	return nil
}

func DeleteBookById(id int) error {
	_, err := yiigo.DB.Exec("DELETE FROM book WHERE id = ?", id)

	if err != nil {
		yiigo.Logger.Error(err.Error())

		return err
	}

	return nil
}
