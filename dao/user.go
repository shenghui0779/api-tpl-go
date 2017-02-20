package dao

import (
	"cache"
	"models"
	"strconv"

	"github.com/iiinsomnia/yiigo"
)

type UserDao struct {
	yiigo.MysqlBase
}

func NewUserDao() *UserDao {
	return &UserDao{yiigo.MysqlBase{Table: "user"}}
}

func (a *UserDao) GetUserById(id int) (*models.UserModel, error) {
	model := &models.UserModel{}

	cache := cache.NewUserCache()

	cacheField := strconv.Itoa(id)
	ok := cache.GetUserDetailCache(cacheField, model)

	if ok {
		return model, nil
	}

	query := map[string]interface{}{
		"where": "id = ?",
		"bind":  []interface{}{id},
	}

	err := a.MysqlBase.FindOne(query, model)

	if err == nil {
		cache.SetUserDetailCache(cacheField, model)
	}

	return model, err
}

func (a *UserDao) GetUserList() ([]models.UserModel, int, error) {
	modelArr := []models.UserModel{}

	count := 0
	query := map[string]interface{}{"count": &count}

	err := a.MysqlBase.Find(query, &modelArr)

	return modelArr, count, err
}
