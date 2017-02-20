package models

import "time"

type UserModel struct {
	ID        int       `sql:"AUTO_INCREMENT" gorm:"column:id;primary_key"`
	Name      string    `gorm:"column:name"`
	Email     string    `gorm:"column:email"`
	Password  string    `gorm:"column:password"`
	Salt      string    `gorm:"column:salt"`
	Role      string    `gorm:"column:role"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
