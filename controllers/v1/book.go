package v1

import (
	service "demo/service/v1"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/iiinsomnia/yiigo"
)

type BookForm struct {
	Title       string `form:"title" binding:"required"`
	SubTitle    string `form:"subtitle" binding:"required"`
	Author      string `form:"author" binding:"required"`
	Version     string `form:"version" binding:"required"`
	Price       string `form:"price" binding:"required"`
	Publisher   string `form:"publisher" binding:"required"`
	PublishDate string `form:"publish_date" binding:"required"`
}

func BookIndex(c *gin.Context) {
	data, err := service.GetAllBooks()

	if err != nil {
		yiigo.Failed(c, "server internal error")
		return
	}

	yiigo.Success(c, data)
}

func BookView(c *gin.Context) {
	id := c.Param("id")

	_id, err := strconv.Atoi(id)

	if err != nil {
		yiigo.Failed(c, "param error")
		return
	}

	data, err := service.GetBookById(_id)

	if err != nil {
		yiigo.Failed(c, "server internal error")
		return
	}

	yiigo.Success(c, data)
}

func BookAdd(c *gin.Context) {
	form := &BookForm{}

	if validate := c.ShouldBindWith(form, binding.Form); validate != nil {
		yiigo.Failed(c, validate.Error())
		return
	}

	data := yiigo.X{
		"title":        form.Title,
		"subtitle":     form.SubTitle,
		"author":       form.Author,
		"version":      form.Version,
		"price":        form.Price,
		"publisher":    form.Publisher,
		"publish_date": form.PublishDate,
	}

	id, err := service.AddNewBook(data)

	if err != nil {
		yiigo.Failed(c, "server internal error")
		return
	}

	yiigo.Success(c, id)
}

func BookEdit(c *gin.Context) {
	id := c.Param("id")

	_id, err := strconv.Atoi(id)

	if err != nil {
		yiigo.Failed(c, "param error")
		return
	}

	form := &BookForm{}

	if validate := c.ShouldBindWith(form, binding.Form); validate != nil {
		yiigo.Failed(c, validate.Error())
		return
	}

	data := yiigo.X{
		"title":        form.Title,
		"subtitle":     form.SubTitle,
		"author":       form.Author,
		"version":      form.Version,
		"price":        form.Price,
		"publisher":    form.Publisher,
		"publish_date": form.PublishDate,
	}

	err = service.UpdateBookById(_id, data)

	if err != nil {
		yiigo.Failed(c, "server internal error")
		return
	}

	yiigo.Success(c)
}

func BookDelete(c *gin.Context) {
	id := c.Param("id")

	_id, err := strconv.Atoi(id)

	if err != nil {
		yiigo.Failed(c, "param error")
		return
	}

	err = service.DeleteBookById(_id)

	if err != nil {
		yiigo.Failed(c, "server internal error")
		return
	}

	yiigo.Success(c)
}
