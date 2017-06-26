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

func (s *StudentDao) GetById(id int, data interface{}) error {
	query := bson.M{"_id": id}
	err := s.Mongo.FindOne(query, data)

	if err != nil {
		yiigo.LogError(err.Error())
		return err
	}

	return nil
}

func (s *StudentDao) GetAll(data interface{}) error {
	err := s.Mongo.FindAll(data)

	if err != nil {
		yiigo.LogError(err.Error())
		return err
	}

	return nil
}

func (s *StudentDao) AddNewRecord(data bson.M) (int, error) {
	id, err := s.Mongo.Insert(data)

	if err != nil {
		yiigo.LogError(err.Error())
		return 0, err
	}

	return id, nil
}

func (s *StudentDao) UpdateById(id int, data bson.M) error {
	query := bson.M{"_id": id}
	err := s.Mongo.Update(query, data)

	if err != nil {
		yiigo.LogError(err.Error())
		return err
	}

	return nil
}

func (s *StudentDao) DeleteById(id int) error {
	query := bson.M{"_id": id}
	err := s.Mongo.Delete(query)

	if err != nil {
		yiigo.LogError(err.Error())
		return err
	}

	return nil
}
