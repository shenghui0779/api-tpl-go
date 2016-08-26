package mysql

import (
	"dao/redis"
	"models"

	"github.com/iiinsomnia/yiigo"
)

type AdminDao struct {
	yiigo.MysqlBase
}

func NewAdminDao() *AdminDao {
	return &AdminDao{yiigo.MysqlBase{TableName: "admin"}}
}

func (a *AdminDao) GetAdminById(id int) (*models.AdminModel, error) {
	adminRedis := redis.NewAdminRedis()
	adminModel := models.AdminModel{}

	ok := adminRedis.GetAdminCache(&adminModel, id)

	if !ok {
		query := map[string]interface{}{"id": id}

		err := a.MysqlBase.FindOne(&adminModel, query)

		if err != nil {
			return nil, err
		}

		adminRedis.SetAdminCache(id, adminModel)
	}

	return &adminModel, nil
}

func (a *AdminDao) GetAdminByName(name string) (*models.AdminModel, error) {
	adminModel := models.AdminModel{}
	query := map[string]interface{}{"where": "name = ?"}

	err := a.MysqlBase.FindOneBySql(&adminModel, query, name)

	if err != nil {
		return nil, err
	}

	return &adminModel, nil
}

func (a *AdminDao) GetAllAdmin() ([]models.AdminModel, error) {
	adminModelArr := []models.AdminModel{}

	query := map[string]interface{}{}

	err := a.MysqlBase.Find(&adminModelArr, query)

	return adminModelArr, err
}

func (a *AdminDao) AddAdmin(name string, password string, salt string, role int, memo string, status int) error {
	adminModel := models.AdminModel{
		Name:     name,
		Password: password,
		Salt:     salt,
		Role:     role,
		Memo:     memo,
		Status:   status,
	}

	err := a.MysqlBase.Insert(&adminModel)

	return err
}
