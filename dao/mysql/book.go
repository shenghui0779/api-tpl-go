package mysql

import "github.com/iiinsomnia/yiigo"

type BookDao struct {
	yiigo.MySQL
}

func NewBookDao() *BookDao {
	return &BookDao{
		yiigo.MySQL{
			Table: "book",
		},
	}
}

func (a *BookDao) GetById(id int, data interface{}) error {
	query := yiigo.X{
		"where": "id = ?",
		"binds": []interface{}{id},
	}

	err := a.MySQL.FindOne(query, data)

	return err
}

func (a *BookDao) GetAll(data interface{}) error {
	err := a.MySQL.FindAll(data)

	return err
}

func (a *BookDao) AddNewRecord(data yiigo.X) (int64, error) {
	id, err := a.MySQL.Insert(data)

	return id, err
}

func (a *BookDao) UpdateById(id int, data yiigo.X) error {
	query := yiigo.X{
		"where": "id = ?",
		"binds": []interface{}{id},
	}

	_, err := a.MySQL.Update(query, data)

	return err
}

func (a *BookDao) DeleteById(id int) error {
	query := yiigo.X{
		"where": "id = ?",
		"binds": []interface{}{id},
	}

	_, err := a.MySQL.Delete(query)

	return err
}
