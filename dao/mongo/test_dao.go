package mongo

import (
	"fmt"
	"models"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/iiinsomnia/yiigo"
)

type TestMongo struct {
	yiigo.MongoBase
}

func NewTestMongo() *TestMongo {
	return &TestMongo{yiigo.MongoBase{CollectionName: "test"}}
}

func (t *TestMongo) GetPersonList() ([]models.PersonModel, error) {
	personModelArr := []models.PersonModel{}
	query := bson.M{}
	var count int
	options := map[string]interface{}{"count": &count, "order": "-_id"}

	err := t.MongoBase.Find(&personModelArr, query, options)

	fmt.Println(count)

	if err != nil {
		return nil, err
	}

	return personModelArr, nil
}

func (t *TestMongo) AddPerson(name string, sex int, age int) error {
	personModel := models.PersonModel{Name: name, Sex: sex, Age: age, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	return t.MongoBase.Insert(&personModel)
}

func (t *TestMongo) UpdatePersonAgeById(id int64, age int) error {
	query := bson.M{"_id": id}
	return t.MongoBase.Increment(query, "age", age)
}

func (t *TestMongo) GetPersonById(id int64) (*models.PersonModel, error) {
	personModel := models.PersonModel{}
	query := bson.M{"_id": id}
	err := t.MongoBase.FindOne(&personModel, query)

	if err != nil {
		return nil, err
	}

	return &personModel, nil
}

func (t *TestMongo) DeletePersonById(id int64) error {
	query := bson.M{"_id": id}
	return t.MongoBase.Delete(query)
}
