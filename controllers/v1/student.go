package v1

import (
	"demo/service"
	"strconv"

	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"
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
		yiigo.ReturnJson(c, -1, "server internal error")
		return
	}

	yiigo.ReturnSuccess(c, data)
}

func StudentView(c *gin.Context) {
	id := c.Param("id")

	_id, err := strconv.Atoi(id)

	if err != nil {
		yiigo.ReturnJson(c, -1, "param error")
		return
	}

	data, err := service.GetStudentByID(_id)

	if err != nil {
		yiigo.ReturnJson(c, -1, "server internal error")
		return
	}

	yiigo.ReturnSuccess(c, data)
}

func StudentAdd(c *gin.Context) {
	var form StudentForm

	if validate := c.Bind(&form); validate != nil {
		yiigo.ReturnJson(c, -1, validate.Error())
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
		yiigo.ReturnJson(c, -1, "server internal error")
		return
	}

	yiigo.ReturnSuccess(c, id)
}

func StudentEdit(c *gin.Context) {
	id := c.Param("id")

	_id, err := strconv.Atoi(id)

	if err != nil {
		yiigo.ReturnJson(c, -1, "param error")
		return
	}

	var form StudentForm

	if validate := c.Bind(&form); validate != nil {
		yiigo.ReturnJson(c, -1, validate.Error())
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
		yiigo.ReturnJson(c, -1, "server internal error")
		return
	}

	yiigo.ReturnSuccess(c)
}

func StudentDelete(c *gin.Context) {
	id := c.Param("id")

	_id, err := strconv.Atoi(id)

	if err != nil {
		yiigo.ReturnJson(c, -1, "param error")
		return
	}

	err = service.DeleteStudentByID(_id)

	if err != nil {
		yiigo.ReturnJson(c, -1, "server internal error")
		return
	}

	yiigo.ReturnSuccess(c)
}
