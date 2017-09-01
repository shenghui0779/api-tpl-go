package service

import (
	"database/sql"
	"demo/cache"
	"demo/models"

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

func GetAllBooks() ([]yiigo.X, error) {
	defer yiigo.Flush()

	books := []models.Book{}

	query := "SELECT * FROM book"
	err := yiigo.DB.Select(&books, query)

	if err != nil && err != sql.ErrNoRows {
		yiigo.Errf("%s, SQL: %s", err.Error(), query)

		return nil, err
	}

	data := formatBookList(books)

	return data, nil
}

func AddNewBook(data yiigo.X) (int64, error) {
	defer yiigo.Flush()

	sql, binds := yiigo.InsertSQL("book", data)
	r, err := yiigo.DB.Exec(sql, binds...)

	if err != nil {
		yiigo.Errf("%s, SQL: %s, Args: %v", err.Error(), sql, binds)

		return 0, err
	}

	id, err := r.LastInsertId()

	return id, err
}

func UpdateBookById(id int, data yiigo.X) error {
	defer yiigo.Flush()

	sql, binds := yiigo.UpdateSQL("UPDATE book SET ? WHERE id = ?", data, id)
	_, err := yiigo.DB.Exec(sql, binds)

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

func formatBookList(books []models.Book) []yiigo.X {
	data := []yiigo.X{}

	for _, v := range books {
		item := yiigo.X{
			"id":           v.ID,
			"title":        v.Title,
			"subtitle":     v.SubTitle,
			"author":       v.Author,
			"version":      v.Version,
			"price":        v.Price,
			"publisher":    v.Publisher,
			"publish_date": v.PublishDate,
		}

		data = append(data, item)
	}

	return data
}
