package models

type AdminModel struct {
	Id       int    `sql:"AUTO_INCREMENT" gorm:"column:id;primary_key"`
	Name     string `gorm:"column:name"`
	Password string `gorm:"column:password"`
	Salt     string `gorm:"column:salt"`
	Role     int    `gorm:"column:role"`
	Memo     string `gorm:"column:memo"`
	Status   int    `gorm:"column:status"`
}
