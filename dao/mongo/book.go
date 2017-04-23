package mongo

import (
	"github.com/iiinsomnia/yiigo"
	"gopkg.in/mgo.v2/bson"
)

type BookDao struct {
	yiigo.Mongo
}

func NewBookDao() *BookDao {
	return &BookDao{
		yiigo.Mongo{
			DB:         "library",
			Collection: "book",
		},
	}
}

func (b *BookDao) GetBookById(id int, data interface{}) error {
	query := bson.M{"_id": id}
	err := b.Mongo.FindOne(query, data)

	return err
}

func (b *BookDao) GetAllBooks(data interface{}) error {
	err := b.Mongo.FindAll(data)

	return err
}

func (b *BookDao) AddNewBook(data bson.M) (int, error) {
	id, err := b.Mongo.Insert(data)

	return id, err
}

func (b *BookDao) UpdateBookById(id int, data bson.M) error {
	query := bson.M{"_id": id}
	err := b.Mongo.Update(query, data)

	return err
}

func (b *BookDao) DeleteBookById(id int) error {
	query := bson.M{"_id": id}
	err := b.Mongo.Delete(query)

	return err
}
