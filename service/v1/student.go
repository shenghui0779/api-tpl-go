package service

import (
	"demo/dao/mongo"
	"time"

	"github.com/iiinsomnia/yiigo"
	"gopkg.in/mgo.v2/bson"
)

type Student struct {
	ID        int       `bson:"_id"`
	Name      string    `bson:"name"`
	Sex       string    `bson:"sex"`
	Age       int       `bson:"age"`
	School    string    `bson:"school"`
	Grade     string    `bson:"grade"`
	Class     string    `bson:"class"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

func GetStudentByID(id int) (yiigo.X, error) {
	Student := &Student{}

	StudentDao := mongo.NewStudentDao()
	err := StudentDao.GetById(id, Student)

	if err != nil {
		if msg := err.Error(); msg != "not found" {
			return nil, err
		}

		return yiigo.X{}, nil
	}

	data := yiigo.X{
		"id":     Student.ID,
		"name":   Student.Name,
		"sex":    Student.Sex,
		"age":    Student.Age,
		"school": Student.School,
		"grade":  Student.Grade,
		"class":  Student.Class,
	}

	return data, nil
}

func GetAllStudents() ([]yiigo.X, error) {
	Students := []Student{}

	StudentDao := mongo.NewStudentDao()
	err := StudentDao.GetAll(&Students)

	if err != nil {
		if msg := err.Error(); msg != "not found" {
			return nil, err
		}

		return []yiigo.X{}, nil
	}

	data := formatStudentList(Students)

	return data, err
}

func AddNewStudent(data bson.M) (int, error) {
	StudentDao := mongo.NewStudentDao()

	data["created_at"] = time.Now()
	data["updated_at"] = time.Now()

	id, err := StudentDao.AddNewRecord(data)

	return id, err
}

func UpdateStudentByID(id int, data bson.M) error {
	StudentDao := mongo.NewStudentDao()
	err := StudentDao.UpdateById(id, data)

	return err
}

func DeleteStudentByID(id int) error {
	StudentDao := mongo.NewStudentDao()
	err := StudentDao.DeleteById(id)

	return err
}

func formatStudentList(Students []Student) []yiigo.X {
	data := []yiigo.X{}

	for _, v := range Students {
		item := map[string]interface{}{
			"id":     v.ID,
			"name":   v.Name,
			"sex":    v.Sex,
			"age":    v.Age,
			"school": v.School,
			"grade":  v.Grade,
			"class":  v.Class,
		}

		data = append(data, item)
	}

	return data
}
