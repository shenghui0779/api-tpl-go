package dao

import (
	"cache"
	"models"
	"strconv"

	"github.com/iiinsomnia/yiigo"
)

type AdminDao struct {
	yiigo.MysqlBase
}

func NewAdminDao() *AdminDao {
	return &AdminDao{yiigo.MysqlBase{Table: "admin"}}
}

func (a *AdminDao) GetAdminById(id int) (*models.AdminModel, error) {
	model := &models.AdminModel{}

	cache := cache.NewAdminCache()

	cacheField := strconv.Itoa(id)
	ok := cache.GetAdminDetailCache(cacheField, model)

	if ok {
		return model, nil
	}

	query := map[string]interface{}{
		"where": "id = ?",
		"bind":  []interface{}{id},
	}

	err := a.MysqlBase.FindOne(query, model)

	if err == nil {
		cache.SetAdminDetailCache(cacheField, model)
	}

	return model, err
}

func (a *AdminDao) GetAdminList() ([]models.AdminModel, int, error) {
	modelArr := []models.AdminModel{}

	count := 0
	query := map[string]interface{}{"count": &count}

	err := a.MysqlBase.Find(query, &modelArr)

	return modelArr, count, err
}
