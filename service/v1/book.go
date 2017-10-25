package service

import (
	"database/sql"
	"demo/cache"
	"demo/models"
	"time"

	"github.com/iiinsomnia/yiigo"
)

func GetBookById(id int) (yiigo.X, error) {
	defer yiigo.Flush()

	book := &models.Book{}

	ok := cache.GetBookCache(id, book)

	if !ok {
		query := "SELECT * FROM book WHERE id = ?"
		err := yiigo.DB.Get(book, query, id)

		if err != nil {
			if err != sql.ErrNoRows {
				yiigo.Errf("%s, SQL: %s, Args: [%d]", err.Error(), query, id)

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
	defer yiigo.Flush()

	books := []models.Book{}

	query := "SELECT * FROM book WHERE 1=1"
	err := yiigo.DB.Select(&books, query)

	if err != nil && err != sql.ErrNoRows {
		yiigo.Errf("%s, SQL: %s", err.Error(), query)

		return nil, err
	}

	return books, nil
}

func AddNewBook(data *models.BookAdd) (int64, error) {
	defer yiigo.Flush()

	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()

	sql, binds := yiigo.InsertSQL("book", data)
	r, err := yiigo.DB.Exec(sql, binds...)

	if err != nil {
		yiigo.Errf("%s, SQL: %s, Args: %v", err.Error(), sql, binds)

		return 0, err
	}

	id, err := r.LastInsertId()

	return id, err
}

func UpdateBookById(id int, data *models.BookEdit) error {
	defer yiigo.Flush()

	data.UpdatedAt = time.Now()

	sql, binds := yiigo.UpdateSQL("UPDATE book SET ? WHERE id = ?", data, id)
	_, err := yiigo.DB.Exec(sql, binds...)

	if err != nil {
		yiigo.Errf("%s, SQL: %s, Args: %v", err.Error(), sql, binds)

		return err
	}

	return nil
}

func DeleteBookById(id int) error {
	defer yiigo.Flush()

	query := "DELETE FROM book WHERE id = ?"
	_, err := yiigo.DB.Exec(query, id)

	if err != nil {
		yiigo.Errf("%s, SQL: %s, Args: [%d]", err.Error(), query, id)

		return err
	}

	return nil
}
