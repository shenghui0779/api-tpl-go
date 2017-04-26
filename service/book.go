package service

import (
	"dao/mongo"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/iiinsomnia/yiigo"
)

type Book struct {
	ID          int       `bson:"_id"`
	Title       string    `bson:"title"`
	Author      string    `bson:"author"`
	Version     string    `bson:"version"`
	Publisher   string    `bson:"publisher"`
	PublishTime string    `bson:"publishtime"`
	CreatedAt   time.Time `bson:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at"`
}

func GetBookById(id int) (yiigo.X, error) {
	book := &Book{}

	bookDao := mongo.NewBookDao()
	err := bookDao.FindById(id, book)

	if err != nil {
		if msg := err.Error(); msg != "sql: no rows in result set" {
			return nil, err
		}

		return yiigo.X{}, nil
	}

	data := yiigo.X{
		"id":          book.ID,
		"title":       book.Title,
		"author":      book.Author,
		"version":     book.Version,
		"publisher":   book.Publisher,
		"publishtime": book.PublishTime,
	}

	return data, nil

	return nil, nil
}

func GetAllBooks() ([]yiigo.X, error) {
	books := []Book{}

	bookDao := mongo.NewBookDao()
	err := bookDao.FindAll(&books)

	if err != nil {
		if msg := err.Error(); msg != "sql: no rows in result set" {
			return nil, err
		}

		return []yiigo.X{}, nil
	}

	data := formatBookList(books)

	return data, err
}

func AddNewBook(data bson.M) (int, error) {
	bookDao := mongo.NewBookDao()
	id, err := bookDao.Add(data)

	return id, err
}

func UpdateBookById(id int, data bson.M) error {
	bookDao := mongo.NewBookDao()
	err := bookDao.UpdateById(id, data)

	return err
}

func DeleteBookById(id int) error {
	bookDao := mongo.NewBookDao()
	err := bookDao.DeleteById(id)

	return err
}

func formatBookList(books []Book) []yiigo.X {
	data := []yiigo.X{}

	for _, v := range books {
		item := map[string]interface{}{
			"id":          v.ID,
			"title":       v.Title,
			"author":      v.Author,
			"version":     v.Version,
			"publisher":   v.Publisher,
			"publishtime": v.PublishTime,
		}

		data = append(data, item)
	}

	return data
}
