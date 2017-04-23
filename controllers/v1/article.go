package v1

import (
	"service"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/iiinsomnia/yiigo"
)

type ArticleAddForm struct {
	Title    string `form:"title" binding:"required"`
	AuthorID int    `form:"author_id" binding:"required"`
	Content  string `form:"content" binding:"required"`
}

type ArticleUpdateForm struct {
	Title    string `form:"title" binding:"required"`
	AuthorID int    `form:"author_id" binding:"required"`
	Content  string `form:"content" binding:"required"`
	Status   int    `form:"status" binding:"required"`
}

func GetArticleList(c *gin.Context) {
	data, err := service.GetAllArticles()

	if err != nil {
		yiigo.ReturnJson(c, -1, "server internal error")
		return
	}

	yiigo.ReturnSuccess(c, data)
}

func GetArticleDetail(c *gin.Context) {
	id := c.Param("id")

	_id, err := strconv.Atoi(id)

	if err != nil {
		yiigo.ReturnJson(c, -1, "param error")
		return
	}

	data, err := service.GetArticleById(_id)

	if err != nil {
		yiigo.ReturnJson(c, -1, "server internal error")
		return
	}

	yiigo.ReturnSuccess(c, data)
}

func AddNewArticle(c *gin.Context) {
	var form ArticleAddForm

	if validate := c.Bind(&form); validate != nil {
		yiigo.ReturnJson(c, -1, validate.Error())
		return
	}

	data := yiigo.X{
		"title":     c.PostForm("title"),
		"author_id": c.PostForm("author_id"),
		"content":   c.PostForm("content"),
	}

	id, err := service.AddNewArticle(data)

	if err != nil {
		yiigo.ReturnJson(c, -1, "server internal error")
		return
	}

	yiigo.ReturnSuccess(c, id)
}

func UpdateArticle(c *gin.Context) {
	id := c.Param("id")

	_id, err := strconv.Atoi(id)

	if err != nil {
		yiigo.ReturnJson(c, -1, "param error")
		return
	}

	var form ArticleUpdateForm

	if validate := c.Bind(&form); validate != nil {
		yiigo.ReturnJson(c, -1, validate.Error())
		return
	}

	data := yiigo.X{
		"title":     c.PostForm("title"),
		"author_id": c.PostForm("author_id"),
		"content":   c.PostForm("content"),
		"status":    c.PostForm("status"),
	}

	err = service.UpdateArticleById(_id, data)

	if err != nil {
		yiigo.ReturnJson(c, -1, "server internal error")
		return
	}

	yiigo.ReturnSuccess(c)
}

func DeleteArticle(c *gin.Context) {
	id := c.Param("id")

	_id, err := strconv.Atoi(id)

	if err != nil {
		yiigo.ReturnJson(c, -1, "param error")
		return
	}

	err = service.DeleteArticleById(_id)

	if err != nil {
		yiigo.ReturnJson(c, -1, "server internal error")
		return
	}

	yiigo.ReturnSuccess(c)
}
