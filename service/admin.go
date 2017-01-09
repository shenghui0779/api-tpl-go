package service

import (
	"dao"
	"service/jsons"
)

func GetAdminById(id int) (*jsons.AdminJson, error) {
	adminDao := dao.NewAdminDao()
	admin, err := adminDao.GetAdminById(id)

	if err != nil {
		return nil, err
	}

	data := jsons.FormatAdmin2Json(admin)

	return data, nil
}

func GetAdminList() ([]jsons.AdminJson, int, error) {
	adminDao := dao.NewAdminDao()
	adminArr, count, err := adminDao.GetAdminList()

	if err != nil {
		return nil, 0, err
	}

	data := jsons.FormatAdminList2Json(adminArr)

	return data, count, nil
}
