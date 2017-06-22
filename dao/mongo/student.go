package mongo

import (
	"github.com/iiinsomnia/yiigo"
	"gopkg.in/mgo.v2/bson"
)

type StudentDao struct {
	yiigo.Mongo
}

func NewStudentDao() *StudentDao {
	return &StudentDao{
		yiigo.Mongo{
			DB:         "demo",
			Collection: "student",
		},
	}
}

func (b *StudentDao) GetById(id int, data interface{}) error {
	query := bson.M{"_id": id}
	err := b.Mongo.FindOne(query, data)

	return err
}

func (b *StudentDao) GetAll(data interface{}) error {
	err := b.Mongo.FindAll(data)

	return err
}

func (b *StudentDao) AddNewRecord(data bson.M) (int, error) {
	id, err := b.Mongo.Insert(data)

	return id, err
}

func (b *StudentDao) UpdateById(id int, data bson.M) error {
	query := bson.M{"_id": id}
	err := b.Mongo.Update(query, data)

	return err
}

func (b *StudentDao) DeleteById(id int) error {
	query := bson.M{"_id": id}
	err := b.Mongo.Delete(query)

	return err
}
