package service

import (
	"dao"
	"service/jsons"
)

func GetUserById(id int) (*jsons.UserJson, error) {
	userDao := dao.NewUserDao()
	user, err := userDao.GetUserById(id)

	if err != nil {
		return nil, err
	}

	data := jsons.FormatUser2Json(user)

	return data, nil
}

func GetUserList() ([]jsons.UserJson, int, error) {
	userDao := dao.NewUserDao()
	userArr, count, err := userDao.GetUserList()

	if err != nil {
		return nil, 0, err
	}

	data := jsons.FormatUserList2Json(userArr)

	return data, count, nil
}
