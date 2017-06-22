package v1

import (
	"demo/service"
	"strconv"

	"github.com/gin-gonic/gin"
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

func GetBookList(c *gin.Context) {
	data, err := service.GetAllBooks()

	if err != nil {
		yiigo.ReturnJson(c, -1, "server internal error")
		return
	}

	yiigo.ReturnSuccess(c, data)
}

func GetBookDetail(c *gin.Context) {
	id := c.Param("id")

	_id, err := strconv.Atoi(id)

	if err != nil {
		yiigo.ReturnJson(c, -1, "param error")
		return
	}

	data, err := service.GetBookById(_id)

	if err != nil {
		yiigo.ReturnJson(c, -1, "server internal error")
		return
	}

	yiigo.ReturnSuccess(c, data)
}

func AddNewBook(c *gin.Context) {
	var form BookForm

	if validate := c.Bind(&form); validate != nil {
		yiigo.ReturnJson(c, -1, validate.Error())
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
		yiigo.ReturnJson(c, -1, "server internal error")
		return
	}

	yiigo.ReturnSuccess(c, id)
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")

	_id, err := strconv.Atoi(id)

	if err != nil {
		yiigo.ReturnJson(c, -1, "param error")
		return
	}

	var form BookForm

	if validate := c.Bind(&form); validate != nil {
		yiigo.ReturnJson(c, -1, validate.Error())
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
		yiigo.ReturnJson(c, -1, "server internal error")
		return
	}

	yiigo.ReturnSuccess(c)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	_id, err := strconv.Atoi(id)

	if err != nil {
		yiigo.ReturnJson(c, -1, "param error")
		return
	}

	err = service.DeleteBookById(_id)

	if err != nil {
		yiigo.ReturnJson(c, -1, "server internal error")
		return
	}

	yiigo.ReturnSuccess(c)
}
