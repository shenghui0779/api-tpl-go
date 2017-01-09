package jsons

import (
	"models"
	"time"
)

type AdminJson struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Salt      string    `json:"salt"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FormatAdmin2Json(data *models.AdminModel) *AdminJson {
	result := &AdminJson{}

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

func FormatAdminList2Json(data []models.AdminModel) []AdminJson {
	result := []AdminJson{}

	for _, model := range data {
		temp := AdminJson{}

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
