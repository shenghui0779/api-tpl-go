package service

import (
	"demo/cache"
	"demo/dao/mysql"
	"time"

	"github.com/iiinsomnia/yiigo"
)

type Book struct {
	ID          int       `db:"id" json:"id"`
	Title       string    `db:"title" json:"title"`
	SubTitle    string    `db:"subtitle" json:"subtitle"`
	Author      string    `db:"author" json:"author"`
	Version     string    `db:"version" json:"version"`
	Price       string    `db:"price" json:"price"`
	Publisher   string    `db:"publisher" json:"publisher"`
	PublishDate string    `db:"publish_date" json:"publish_date"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

func GetBookById(id int) (yiigo.X, error) {
	book := &Book{}

	cache := cache.NewBookCache()
	ok := cache.GetBook(id, book)

	if !ok {
		bookDao := mysql.NewBookDao()
		err := bookDao.GetById(id, book)

		if err != nil {
			if err.Error() != "not found" {
				return nil, err
			}

			return yiigo.X{}, nil
		}

		cache.SetBook(id, book)
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
	books := []Book{}

	bookDao := mysql.NewBookDao()
	err := bookDao.GetAll(&books)

	if err != nil {
		return nil, err
	}

	data := formatBookList(books)

	return data, nil
}

func AddNewBook(data yiigo.X) (int64, error) {
	bookDao := mysql.NewBookDao()
	id, err := bookDao.AddNewRecord(data)

	return id, err
}

func UpdateBookById(id int, data yiigo.X) error {
	bookDao := mysql.NewBookDao()
	err := bookDao.UpdateById(id, data)

	return err
}

func DeleteBookById(id int) error {
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
