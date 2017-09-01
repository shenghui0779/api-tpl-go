package service

import (
	"demo/models"
	"time"

	"gopkg.in/mgo.v2"

	"github.com/iiinsomnia/yiigo"
	"gopkg.in/mgo.v2/bson"
)

func GetStudentByID(id int) (yiigo.X, error) {
	defer yiigo.Flush()

	student := &models.Student{}

	mongo := yiigo.Mongo()

	err := mongo.DB("demo").C("student").FindId(id).One(student)

	mongo.Close()

	if err != nil {
		if err != mgo.ErrNotFound {
			yiigo.Err(err.Error())

			return nil, err
		}

		return yiigo.X{}, nil
	}

	data := yiigo.X{
		"id":     student.ID,
		"name":   student.Name,
		"sex":    student.Sex,
		"age":    student.Age,
		"school": student.School,
		"grade":  student.Grade,
		"class":  student.Class,
	}

	return data, nil
}

func GetAllStudents() ([]yiigo.X, error) {
	defer yiigo.Flush()

	students := []models.Student{}

	mongo := yiigo.Mongo()

	err := mongo.DB("demo").C("student").Find(bson.M{}).All(&students)

	mongo.Close()

	if err != nil {
		if err != mgo.ErrNotFound {
			return nil, err
		}

		return []yiigo.X{}, nil
	}

	data := formatStudentList(students)

	return data, nil
}

func AddNewStudent(data bson.M) (int, error) {
	defer yiigo.Flush()

	id, err := yiigo.Seq("demo", "student")

	if err != nil {
		yiigo.Err(err.Error())

		return 0, err
	}

	data["_id"] = id
	data["created_at"] = time.Now()
	data["updated_at"] = time.Now()

	mongo := yiigo.Mongo()

	err = mongo.DB("demo").C("student").Insert(data)

	if err != nil {
		yiigo.Err(err.Error())

		return 0, err
	}

	mongo.Close()

	return id, nil
}

func UpdateStudentByID(id int, data bson.M) error {
	defer yiigo.Flush()

	data["updated_at"] = time.Now()

	mongo := yiigo.Mongo()

	err := mongo.DB("demo").C("student").UpdateId(id, bson.M{"$set": data})

	if err != nil {
		yiigo.Err(err.Error())

		return err
	}

	mongo.Close()

	return nil
}

func DeleteStudentByID(id int) error {
	defer yiigo.Flush()

	mongo := yiigo.Mongo()

	err := mongo.DB("demo").C("student").RemoveId(id)

	if err != nil {
		yiigo.Err(err.Error())

		return err
	}

	mongo.Close()

	return nil
}

func formatStudentList(students []models.Student) []yiigo.X {
	data := []yiigo.X{}

	for _, v := range students {
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
