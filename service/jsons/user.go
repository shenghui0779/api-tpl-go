package jsons

import (
	"models"
	"time"
)

type UserJson struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Salt      string    `json:"salt"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FormatUser2Json(data *models.UserModel) *UserJson {
	result := &UserJson{}

	result.ID = data.ID
	result.Name = data.Name
	result.Email = data.Email
	result.Password = data.Password
	result.Salt = data.Salt
	result.Role = data.Role
	result.CreatedAt = data.CreatedAt
	result.UpdatedAt = data.UpdatedAt

	return result
}

func FormatUserList2Json(data []models.UserModel) []UserJson {
	result := []UserJson{}

	for _, model := range data {
		temp := UserJson{}

		temp.ID = model.ID
		temp.Name = model.Name
		temp.Email = model.Email
		temp.Password = model.Password
		temp.Salt = model.Salt
		temp.Role = model.Role
		temp.CreatedAt = model.CreatedAt
		temp.UpdatedAt = model.UpdatedAt

		result = append(result, temp)
	}

	return result
}
