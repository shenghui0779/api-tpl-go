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

func (b *BookDao) GetById(id int, data interface{}) error {
	query := yiigo.X{
		"where": "id = ?",
		"binds": []interface{}{id},
	}

	err := b.MySQL.FindOne(query, data)

	if err != nil {
		if err.Error() != "not found" {
			yiigo.LogError(err.Error())
		}

		return err
	}

	return nil
}

func (b *BookDao) GetAll(data interface{}) error {
	err := b.MySQL.FindAll(data)

	if err != nil {
		yiigo.LogError(err.Error())
		return err
	}

	return nil
}

func (b *BookDao) AddNewRecord(data yiigo.X) (int64, error) {
	id, err := b.MySQL.Insert(data)

	if err != nil {
		yiigo.LogError(err.Error())
		return 0, err
	}

	return id, nil
}

func (b *BookDao) UpdateById(id int, data yiigo.X) error {
	query := yiigo.X{
		"where": "id = ?",
		"binds": []interface{}{id},
	}

	_, err := b.MySQL.Update(query, data)

	if err != nil {
		yiigo.LogError(err.Error())
		return err
	}

	return nil
}

func (b *BookDao) DeleteById(id int) error {
	query := yiigo.X{
		"where": "id = ?",
		"binds": []interface{}{id},
	}

	_, err := b.MySQL.Delete(query)

	if err != nil {
		yiigo.LogError(err.Error())
		return err
	}

	return nil
}
