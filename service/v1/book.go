package service

import (
	"database/sql"
	"demo/cache"
	"demo/dao/mysql"
	"demo/models"

	"github.com/iiinsomnia/yiigo"
)

func GetBookById(id int) (yiigo.X, error) {
	defer yiigo.Flush()

	book := &models.Book{}

	ok := cache.GetBookCache(id, book)

	if !ok {
		query := "SELECT * FROM yii_book WHERE id = ?"
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

	books := []Book{}

	query := "SELECT * FROM yii_book"
	err := yiigo.DB.Get(book, query, id)

	if err != nil && err != sql.ErrNoRows {
		yiigo.Errf("%s, SQL: %s, Args: [%d]", err.Error(), query, id)

		return nil, err
	}

	data := formatBookList(books)

	return data, nil
}

func AddNewBook(data yiigo.X) (int64, error) {
	defer yiigo.Flush()

	bookDao := mysql.NewBookDao()
	id, err := bookDao.AddNewRecord(data)

	return id, err
}

func UpdateBookById(id int, data yiigo.X) error {
	defer yiigo.Flush()

	bookDao := mysql.NewBookDao()
	err := bookDao.UpdateById(id, data)

	return err
}

func DeleteBookById(id int) error {
	defer yiigo.Flush()

	bookDao := mysql.NewBookDao()
	err := bookDao.DeleteById(id)

	return err
}

func formatBookList(books []Book) []yiigo.X {
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
