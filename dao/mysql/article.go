package mysql

import "github.com/iiinsomnia/yiigo"

type ArticleDao struct {
	yiigo.MySQL
}

func NewArticleDao() *ArticleDao {
	return &ArticleDao{
		yiigo.MySQL{
			Table: "article",
		},
	}
}

func (a *ArticleDao) FindById(id int, data interface{}) error {
	query := yiigo.X{
		"where": "id = ?",
		"binds": []interface{}{id},
	}

	err := a.MySQL.FindOne(query, data)

	return err
}

func (a *ArticleDao) FindAll(data interface{}) error {
	err := a.MySQL.FindAll(data)

	return err
}

func (a *ArticleDao) Add(data yiigo.X) (int64, error) {
	id, err := a.MySQL.Insert(data)

	return id, err
}

func (a *ArticleDao) UpdateById(id int, data yiigo.X) error {
	query := yiigo.X{
		"where": "id = ?",
		"binds": []interface{}{id},
	}

	_, err := a.MySQL.Update(query, data)

	return err
}

func (a *ArticleDao) DeleteById(id int) error {
	query := yiigo.X{
		"where": "id = ?",
		"binds": []interface{}{id},
	}

	_, err := a.MySQL.Delete(query)

	return err
}
