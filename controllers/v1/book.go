package v1

import (
	"demo/models"
	service "demo/service/v1"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/iiinsomnia/yiigo"
)

func BookIndex(c *gin.Context) {
	data, err := service.GetAllBooks()

	if err != nil {
		yiigo.Error(c, 500, "internal server error")

		return
	}

	yiigo.OK(c, data)
}

func BookView(c *gin.Context) {
	id := c.Param("id")

	_id, err := strconv.Atoi(id)

	if err != nil {
		yiigo.Error(c, -1, "param error")

		return
	}

	data, err := service.GetBookById(_id)

	if err != nil {
		yiigo.Error(c, 500, "internal server error")

		return
	}

	yiigo.OK(c, data)
}

func BookAdd(c *gin.Context) {
	form := &models.BookAdd{}

	if validate := c.ShouldBindWith(form, binding.Form); validate != nil {
		yiigo.Error(c, -1, validate.Error())

		return
	}

	id, err := service.AddNewBook(form)

	if err != nil {
		yiigo.Error(c, 500, "internal server error")

		return
	}

	yiigo.OK(c, gin.H{"id": id})
}

func BookEdit(c *gin.Context) {
	id := c.Param("id")

	_id, err := strconv.Atoi(id)

	if err != nil {
		yiigo.Error(c, -1, "param error")

		return
	}

	form := &models.BookEdit{}

	if validate := c.ShouldBindWith(form, binding.Form); validate != nil {
		yiigo.Error(c, -1, validate.Error())

		return
	}

	err = service.UpdateBookById(_id, form)

	if err != nil {
		yiigo.Error(c, 500, "internal server error")

		return
	}

	yiigo.OK(c)
}

func BookDelete(c *gin.Context) {
	id := c.Param("id")

	_id, err := strconv.Atoi(id)

	if err != nil {
		yiigo.Error(c, -1, "param error")

		return
	}

	err = service.DeleteBookById(_id)

	if err != nil {
		yiigo.Error(c, 500, "internal server error")

		return
	}

	yiigo.OK(c)
}
