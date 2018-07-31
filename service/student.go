package service

import (
	"demo/models"
	"time"

	"github.com/iiinsomnia/yiigo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func GetStudentByID(id int) (yiigo.X, error) {
	student := &models.Student{}

	session := yiigo.Mongo.Clone()

	defer session.Close()

	err := session.DB("demo").C("student").FindId(id).One(student)

	if err != nil {
		if err != mgo.ErrNotFound {
			yiigo.Logger.Error(err.Error())

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
	var students []models.Student

	session := yiigo.Mongo.Clone()

	defer session.Close()

	err := session.DB("demo").C("student").Find(bson.M{}).All(&students)

	if err != nil {
		if err != mgo.ErrNotFound {
			yiigo.Logger.Error(err.Error())

			return nil, err
		}

		return []yiigo.X{}, nil
	}

	data := formatStudentList(students)

	return data, nil
}

func AddNewStudent(data bson.M) (int, error) {
	session := yiigo.Mongo.Clone()

	defer session.Close()

	id, err := yiigo.Seq(session, "demo", "student")

	if err != nil {
		yiigo.Logger.Error(err.Error())

		return 0, err
	}

	data["_id"] = id
	data["created_at"] = time.Now()
	data["updated_at"] = time.Now()

	err = session.DB("demo").C("student").Insert(data)

	if err != nil {
		yiigo.Logger.Error(err.Error())

		return 0, err
	}

	return id, nil
}

func UpdateStudentByID(id int, data bson.M) error {
	data["updated_at"] = time.Now()

	session := yiigo.Mongo.Clone()

	err := session.DB("demo").C("student").UpdateId(id, bson.M{"$set": data})

	if err != nil {
		yiigo.Logger.Error(err.Error())

		return err
	}

	session.Close()

	return nil
}

func DeleteStudentByID(id int) error {
	session := yiigo.Mongo.Clone()

	err := session.DB("demo").C("student").RemoveId(id)

	if err != nil {
		yiigo.Logger.Error(err.Error())

		return err
	}

	session.Close()

	return nil
}

func formatStudentList(students []models.Student) []yiigo.X {
	var data []yiigo.X

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
