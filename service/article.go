package service

import (
	"cache"
	"dao/mysql"
	"strconv"
	"time"

	"github.com/iiinsomnia/yiigo"
)

type Article struct {
	ID        int       `db:"id"`
	Title     string    `db:"title"`
	AuthorID  int       `db:"author_id"`
	Content   string    `db:"content"`
	Status    int       `db:"status"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func GetArticleById(id int) (yiigo.X, error) {
	article := &Article{}

	cache := cache.NewArticleCache()

	cacheField := strconv.Itoa(id)
	ok := cache.GetArticleDetail(cacheField, article)

	if !ok {
		articleDao := mysql.NewArticleDao()
		err := articleDao.GetArticleById(id, article)

		if err != nil {
			if msg := err.Error(); msg != "sql: no rows in result set" {
				return nil, err
			}

			return yiigo.X{}, nil
		}

		cache.SetArticleDetail(cacheField, article)
	}

	data := yiigo.X{
		"id":     article.ID,
		"title":  article.Title,
		"author": article.AuthorID,
		"status": article.Status,
	}

	return data, nil
}

func GetAllArticles() ([]yiigo.X, error) {
	articles := []Article{}

	articleDao := mysql.NewArticleDao()
	err := articleDao.GetAllArticles(&articles)

	if err != nil {
		if msg := err.Error(); msg != "sql: no rows in result set" {
			return nil, err
		}

		return []yiigo.X{}, nil
	}

	data := formatArticleList(articles)

	return data, err
}

func AddNewArticle(data yiigo.X) (int64, error) {
	articleDao := mysql.NewArticleDao()
	id, err := articleDao.AddNewArticle(data)

	return id, err
}

func UpdateArticleById(id int, data yiigo.X) error {
	articleDao := mysql.NewArticleDao()
	err := articleDao.UpdateArticleById(id, data)

	return err
}

func DeleteArticleById(id int) error {
	articleDao := mysql.NewArticleDao()
	err := articleDao.DeleteArticleById(id)

	return err
}

func formatArticleList(articles []Article) []yiigo.X {
	data := []yiigo.X{}

	for _, v := range articles {
		item := yiigo.X{
			"id":        v.ID,
			"title":     v.Title,
			"author_id": v.AuthorID,
			"content":   v.Content,
			"status":    v.Status,
		}

		data = append(data, item)
	}

	return data
}
