package serializer

import (
	"github.com/JohnKucharsky/echo_gorm/models"
)

type UserBody struct {
	FirstName string  `json:"first_name" validate:"required"`
	LastName  *string `json:"last_name"`
}

func UserBodyToUser(userBody UserBody) models.User {
	return models.User{
		FirstName: userBody.FirstName,
		LastName:  userBody.LastName,
	}
}
