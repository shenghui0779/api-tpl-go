package service

import (
	"dao/mysql"
	"models"
)

func GetAdminInfo(id int) (*models.AdminModel, error) {
	adminDao := mysql.NewAdminDao()

	return adminDao.GetAdminById(id)
}

func GetAllAdmin() ([]models.AdminModel, error) {
	adminDao := mysql.NewAdminDao()

	return adminDao.GetAllAdmin()
}
