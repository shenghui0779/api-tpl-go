package v1

import (
	service "demo/service/v1"
	"strconv"

	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/iiinsomnia/yiigo"
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
		yiigo.Failed(c, "server internal error")
		return
	}

	yiigo.Success(c, data)
}

func StudentView(c *gin.Context) {
	id := c.Param("id")

	_id, err := strconv.Atoi(id)

	if err != nil {
		yiigo.Failed(c, "param error")
		return
	}

	data, err := service.GetStudentByID(_id)

	if err != nil {
		yiigo.Failed(c, "server internal error")
		return
	}

	yiigo.Success(c, data)
}

func StudentAdd(c *gin.Context) {
	form := &StudentForm{}

	if validate := c.ShouldBindWith(form, binding.Form); validate != nil {
		yiigo.Failed(c, validate.Error())
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
		yiigo.Failed(c, "server internal error")
		return
	}

	yiigo.Success(c, id)
}

func StudentEdit(c *gin.Context) {
	id := c.Param("id")

	_id, err := strconv.Atoi(id)

	if err != nil {
		yiigo.Failed(c, "param error")
		return
	}

	form := &StudentForm{}

	if validate := c.ShouldBindWith(form, binding.Form); validate != nil {
		yiigo.Failed(c, validate.Error())
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
		yiigo.Failed(c, "server internal error")
		return
	}

	yiigo.Success(c)
}

func StudentDelete(c *gin.Context) {
	id := c.Param("id")

	_id, err := strconv.Atoi(id)

	if err != nil {
		yiigo.Failed(c, "param error")
		return
	}

	err = service.DeleteStudentByID(_id)

	if err != nil {
		yiigo.Failed(c, "server internal error")
		return
	}

	yiigo.Success(c)
}
