package controllers

import (
	"demo/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/mgo.v2/bson"
)

type StudentForm struct {
	Name   string `form:"name" binding:"required"`
	Sex    string `form:"sex" binding:"required"`
	Age    int    `form:"age" binding:"required"`
	School string `form:"school" binding:"required"`
	Grade  string `form:"grade" binding:"required"`
	Class  string `form:"class" binding:"required"`
}

func StudentIndex(c *gin.Context) {
	data, err := service.GetAllStudents()

	if err != nil {
		Err(c, 500, "internal server error")

		return
	}

	OK(c, data)
}

func StudentView(c *gin.Context) {
	id := c.Param("id")

	_id, err := strconv.Atoi(id)

	if err != nil {
		Err(c, -1, "param error")

		return
	}

	data, err := service.GetStudentByID(_id)

	if err != nil {
		Err(c, 500, "internal server error")

		return
	}

	OK(c, data)
}

func StudentAdd(c *gin.Context) {
	form := &StudentForm{}

	if validate := c.ShouldBindWith(form, binding.Form); validate != nil {
		Err(c, -1, validate.Error())

		return
	}

	data := bson.M{
		"name":   form.Name,
		"sex":    form.Sex,
		"age":    form.Age,
		"school": form.School,
		"grade":  form.Grade,
		"class":  form.Class,
	}

	id, err := service.AddNewStudent(data)

	if err != nil {
		Err(c, 500, "internal server error")

		return
	}

	OK(c, gin.H{"id": id})
}

func StudentEdit(c *gin.Context) {
	id := c.Param("id")

	_id, err := strconv.Atoi(id)

	if err != nil {
		Err(c, -1, "param error")

		return
	}

	form := &StudentForm{}

	if validate := c.ShouldBindWith(form, binding.Form); validate != nil {
		Err(c, -1, validate.Error())

		return
	}

	data := bson.M{
		"name":   form.Name,
		"sex":    form.Sex,
		"age":    form.Age,
		"school": form.School,
		"grade":  form.Grade,
		"class":  form.Class,
	}

	err = service.UpdateStudentByID(_id, data)

	if err != nil {
		Err(c, 500, "internal server error")

		return
	}

	OK(c)
}

func StudentDelete(c *gin.Context) {
	id := c.Param("id")

	_id, err := strconv.Atoi(id)

	if err != nil {
		Err(c, -1, "param error")

		return
	}

	err = service.DeleteStudentByID(_id)

	if err != nil {
		Err(c, 500, "internal server error")

		return
	}

	OK(c)
}
